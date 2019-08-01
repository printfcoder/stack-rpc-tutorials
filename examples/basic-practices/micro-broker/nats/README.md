# go-micro-pubsub-with-nats

Original Repository: https://github.com/BruceWangNo1/go-micro-pubsub-with-nats

A PubSub microservice example with NATS as the broker using Go-Micro microservice framework

## Generating srv Service Template

```bash
micro new --namespace=go.micro.pubsub --type=srv --alias=pubsub github.com/brucewangno1/go-micro-pubsub-with-nats/srv
```

## Generating cli Service Template

```bash
micro new --namespace=go.micro.pubsub --type=srv --alias=pubsub github.com/brucewangno1/go-micro-pubsub-with-nats/cli
```

## Writing Code

Templates are just templates. Look through the code to learn how to write microservices.

## Getting NATS Started in The Background

```bash
nats-server
```

## Getting Consul Started in The Background

```bash
consul agent -dev -node localmachine -ui
```

## Run Services

In one terminal windows, do this:

```bash
cd $GOPATH/github.com/brucewangno1/go-micro-pubsub-with-nats/srv
go run main.go --broker=nats --broker_address=127.0.0.1:4222
```

In another one, do this:

```bash
cd $GOPATH/github.com/brucewangno1/go-micro-pubsub-with-nats/cli
go run main.go --broker=nats --broker_address=127.0.0.1:4222
```

All right. All right. All right. You can check whether PubSub is utilizing nats by stopping nats.
