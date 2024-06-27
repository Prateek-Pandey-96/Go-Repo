import jwt, datetime
from datetime import datetime, timedelta, timezone
from fastapi import HTTPException, status, Depends
from config import get_settings 
from cache.cache import getClient, get

settings = get_settings()

def createToken(data: dict, expires_delta: int):
    to_encode = data.copy()
    expire = datetime.now(timezone.utc) + timedelta(minutes=expires_delta)
    to_encode.update({"exp": expire})
    encoded_jwt = jwt.encode(
        to_encode, 
        settings.SECRET_KEY, 
        algorithm=settings.ALGORITHM
    )
    return encoded_jwt

def verify_token(token: str) -> bool:
    try:
        # decoded_token = jwt.decode(
        #     token,
        #     settings.SECRET_KEY,
        #     algorithms=[settings.ALGORITHM]
        # )
        # print(decoded_token)
        generator = getClient()
        print(get(next(generator), "kamlessh"))
        # if (get(getClient(), "kamlesh")) is not None:
        #     return True
        # else:
        #     return False
        return True
    except jwt.PyJWTError as e:
        raise Exception(e)
        return False