# Demo 

```bash
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/greeter.proto
```

```bash
micro call goland.greeter Greeter.Hello '{"name":"Goland"}' 
```