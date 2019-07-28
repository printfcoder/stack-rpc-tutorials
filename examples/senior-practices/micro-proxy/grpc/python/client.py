import grpc
import greeter_pb2
import greeter_pb2_grpc

def main():
    with grpc.insecure_channel("localhost:8081") as channel:
        stub = greeter_pb2_grpc.GreeterStub(channel)
        resp = stub.SayHello(greeter_pb2.HelloRequest(name="HoHoHo"))
        print(resp.message)

if __name__ == "__main__":
    main()