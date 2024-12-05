package cart

import (
	"context"
	"test/pkg/pb"
)

type Handler struct {
	pb.UnimplementedCartServiceServer
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) AddItem(ctx context.Context, in *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	return &pb.AddItemResponse{Message: "hello there"}, nil
}
