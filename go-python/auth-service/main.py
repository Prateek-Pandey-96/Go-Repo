from fastapi import FastAPI, Depends, status, HTTPException
from sqlalchemy.orm import Session
from sqlalchemy.exc import IntegrityError
from redis import Redis
import bcrypt

from database.database import get_db
from tokenHelper import createToken
from cache.cache import getClient, set
from response import Response
from userReq import UserReq


from database import user
# necessary to create tables in db
# from database.database import engine, Base
# Base.metadata.create_all(engine)

app = FastAPI()

@app.post('/auth/register')
def register_user(userReq: UserReq, db:Session = Depends(get_db)):
    hashed_password = bcrypt.hashpw(userReq.password.encode('utf-8'), bcrypt.gensalt())
    userReq.password = hashed_password.decode('utf-8')
    
    try:
        new_user = user.User(**userReq.model_dump())
        
        db.add(new_user)
        db.commit()
        
        return Response(status.HTTP_200_OK, "User registered successfully!")
    except IntegrityError:
        db.rollback()
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Username already exists!"
        )

@app.post('/auth/login')
def login(userReq: UserReq, db:Session = Depends(get_db), cache:Redis = Depends(getClient)):
    fetched_user = db.query(user.User).filter_by(username = userReq.username).first()
    
    if fetched_user is None:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="User not found"
        )
    
    if not bcrypt.checkpw(userReq.password.encode('utf-8'), fetched_user.password.encode('utf-8')):
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Invalid password"
        )
    else:
        token = createToken({"username": userReq.username}, expires_delta=15)
        set(cache, f'{userReq.username}_token', token)
        return Response(status.HTTP_200_OK, token)


if __name__ == '__main__':
    import uvicorn
    uvicorn.run(app, host="localhost", port=8000)