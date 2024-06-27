import grpc
from concurrent import futures
from pb2.auth_pb2_grpc import add_AuthServiceServicer_to_server, AuthServiceServicer
from pb2.auth_pb2 import VerifyResponse
from tokenHelper import verify_token

class AuthService(AuthServiceServicer):
    def VerifyToken(self, request, context):
        validity = verify_token(request.token)
        return VerifyResponse(loggedIn=validity)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_AuthServiceServicer_to_server(AuthService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Server started on port 50051")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()