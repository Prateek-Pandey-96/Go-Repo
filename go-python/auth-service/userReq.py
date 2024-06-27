from pydantic import BaseModel

class UserReq(BaseModel):
    username: str
    password: str