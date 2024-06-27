from pydantic_settings import BaseSettings, SettingsConfigDict
from functools import lru_cache

class Settings(BaseSettings):
    DB_URL: str
    REDIS_HOST: str
    REDIS_PORT: int
    REDIS_DB: int
    SECRET_KEY: str
    ALGORITHM: str
    ACCESS_TOKEN_EXPIRE_MINUTES: int
    model_config = SettingsConfigDict(env_file=".env")

@lru_cache
def get_settings():
    return Settings()