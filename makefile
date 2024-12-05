
.PHONY: proto
proto:
	protoc -I=proto -I=proto/google/api -I=proto/protoc-gen-openapiv2/options --go_out=pkg/pb --go-grpc_out=pkg/pb --grpc-gateway_out=pkg/pb --openapiv2_out=cmd/web/. --openapiv2_opt=output_format=yaml proto/cart.proto