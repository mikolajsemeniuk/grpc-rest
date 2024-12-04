# X

* <https://google.github.io/proto-lens/installing-protoc.html>
* `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
* `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest`
* `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest`

## Test GRPC

```sh
grpcurl -plaintext -d '{"item_id":"123","name":"Smartphone","quantity":2,"price":499.99}' localhost:50051 cart.CartService/AddItem
```

## Test REST

```sh
curl -X POST http://localhost:8080/cart -H "Content-Type: application/json" -d '{"item_id":"123","name":"Smartphone","quantity":2,"price":499.99}'
```
