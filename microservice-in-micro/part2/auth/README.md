# Auth Service

This is the Auth service

Generated with

```
micro new github.com/micro-in-cn/tutorials/microservice-in-micro/part2/auth --namespace=mu.micro.book --alias=auth --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.srv.auth
- Type: srv
- Alias: auth

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
./auth-srv
```

Build a docker image
```
make docker
```