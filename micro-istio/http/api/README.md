# Sample Api

This is the sample service with fqdn go.micro.api.sample.

## Getting Started

### Run Service

```
$ go run main.go
```

### Building a container
```
$ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o sample ./main.go
$ docker build -t hbchen/go-micro-istio-api-sample:v0.0.4 .
```