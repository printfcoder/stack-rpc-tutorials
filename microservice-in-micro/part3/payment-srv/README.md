# Payment Service

This is the Payment service

Generated with

```
micro new github.com/micro-in-cn/tutorials/microservice-in-micro/part3/payment-srv --namespace=mu.micro.book --alias=payment --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.srv.payment
- Type: srv
- Alias: payment

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
./payment-srv
```

Build a docker image
```
make docker
```