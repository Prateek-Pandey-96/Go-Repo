from tokenHelper import verify_token

token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InByYXRlZWsiLCJleHAiOjE3MTk1MTIzNDB9.esml8T4jkrdoLXdGNEP6utCBZ-szMEghEbqnAwURuZM"
print(verify_token(token))