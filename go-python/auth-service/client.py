import grpc
import pb2.auth_pb2
import pb2.auth_pb2_grpc

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = pb2.auth_pb2_grpc.AuthServiceStub(channel)
        token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InByYXRlZWsiLCJleHAiOjE3MTk1MTMzNjd9.C-DYuE4bzJ5CuWvHH5E4_6KKPibpBaHjPGvvTFKN-Ek"
        response = stub.VerifyToken(pb2.auth_pb2.VerifyRequest(token=token))
        print("Validity: ", response.loggedIn)

if __name__ == '__main__':
    run()