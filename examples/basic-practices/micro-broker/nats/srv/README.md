# Pubsub Service

This is the Pubsub service

Generated with

```
micro new github.com/brucewangno1/go-micro-pubsub-with-nats/srv --namespace=go.micro --alias=pubsub --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.pubsub
- Type: srv
- Alias: pubsub

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./pubsub-srv
```

Build a docker image
```
make docker
```