import jwt
import datetime
from datetime import datetime, timedelta, timezone
from config import get_settings 
from cache.cache import getClient, get

settings = get_settings()


def create_token(data: dict, expires_delta: int):
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
        decoded_token = jwt.decode(
            token,
            settings.SECRET_KEY,
            algorithms=[settings.ALGORITHM]
        )
        generator = getClient()
        username = decoded_token['username']
        if (get(next(generator), f'{username}_token')) is not None:
            return True
        else:
            return False
    except jwt.PyJWTError:
        return False