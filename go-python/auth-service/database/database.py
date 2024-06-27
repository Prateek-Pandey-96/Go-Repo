from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.declarative import declarative_base
from config import get_settings

settings = get_settings()

URL_DB = settings.DB_URL
engine = create_engine(URL_DB)
session_local = sessionmaker(autocommit=False, autoflush=False, bind=engine)
Base = declarative_base()

def get_db():
    db = session_local()
    try:
        yield db
    finally:
        db.close()