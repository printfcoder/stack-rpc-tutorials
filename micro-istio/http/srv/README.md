# Sample Srv

This is the sample service with fqdn go.micro.srv.sample.

## Getting Started

### Run Service

```
$ go run main.go
```

### Building a container

If you would like to build the docker container do the following
```
$ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o sample ./main.go
$ docker build -t hbchen/go-micro-istio-srv-sample:v0.0.1 .
```