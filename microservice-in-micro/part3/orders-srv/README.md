# Orders Service

This is the Orders service

Generated with

```
micro new github.com/micro-in-cn/tutorials/microservice-in-micro/part3/orders-srv --namespace=mu.micro.book --alias=orders --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.srv.orders
- Type: srv
- Alias: orders

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
./orders-srv
```

Build a docker image
```
make docker
```