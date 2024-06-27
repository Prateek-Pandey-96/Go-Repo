from redis import Redis
from config import get_settings

settings = get_settings()

def getClient():
    yield Redis(
        host=settings.REDIS_HOST, 
        port=settings.REDIS_PORT, 
        db=settings.REDIS_DB
    )

def set(r: Redis, k: str, val: str):
    try:
        r.set(k, val, settings.ACCESS_TOKEN_EXPIRE_MINUTES)
    except Exception:
        raise Exception(f"An error occured while setting: {k}")
    
def get(r: Redis, k: str):
    try:
        if(r.exists(k)):
            return r.get(k)
        else:
            return None
    except Exception:
        raise Exception(f"An error occured while getting: {k}")