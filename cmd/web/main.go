package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"test/pkg/cart"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcPort := ":50051"
	httpPort := ":8080"
	ctx := context.Background()

	go func() {
		listener, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		server := grpc.NewServer()

		cart.RegisterCartServiceServer(server, cart.NewHandler())

		reflection.Register(server)

		if err := server.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := cart.RegisterCartServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts); err != nil {
		log.Fatalf("Failed to start REST gateway: %v", err)
	}

	if err := http.ListenAndServe(httpPort, mux); err != nil {
		log.Fatalf("Failed to serve REST gateway: %v", err)
	}
}
