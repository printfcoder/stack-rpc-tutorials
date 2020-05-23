# gRPC 风格Server Codec

go-micro的gRPC Server中的默认编码器是私有类，不能直接覆盖。不过可以使用grpc.Codec方法将默认的编码器覆盖掉