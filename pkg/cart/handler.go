package cart

import (
	context "context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

//go:generate protoc -I=. -I=google/api -I=protoc-gen-openapiv2/options --go_out=. --go-grpc_out=. --grpc-gateway_out=. --openapiv2_out=templates/. --openapiv2_opt=output_format=yaml cart.proto
type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) AddItem(ctx context.Context, in *AddItemRequest) (*AddItemResponse, error) {
	return &AddItemResponse{Message: "hello there"}, nil
}

func (h *Handler) HostSwagger(ctx context.Context, in *emptypb.Empty) (*TemplateResponse, error) {
	return &TemplateResponse{Content: swagger}, nil
}

func (h *Handler) HostElements(ctx context.Context, in *emptypb.Empty) (*TemplateResponse, error) {
	return &TemplateResponse{Content: elements}, nil
}

func (h *Handler) mustEmbedUnimplementedCartServiceServer() {}
