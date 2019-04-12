# Service Service

This is the Service service

Generated with

```
micro new github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service --namespace=mu.micro.book.user --alias=service --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.user.srv.service
- Type: srv
- Alias: service

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
./service-srv
```

Build a docker image
```
make docker
```