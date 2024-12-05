package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"test/pkg/cart"
	"test/pkg/docs"
	"test/pkg/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcPort := ":50051"

	go func() {
		listener, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		server := grpc.NewServer()

		pb.RegisterCartServiceServer(server, cart.NewHandler())
		// pb.RegisterAuthServiceServer(server, auth.NewHandler())
		// pb.RegisterCheckoutServiceServer(server, checkout.NewHandler())

		reflection.Register(server)

		log.Printf("start gRPC server on port %s\n", grpcPort)
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	ctx := context.Background()

	router := http.NewServeMux()
	mux := runtime.NewServeMux()

	router.Handle("/", mux)
	router.Handle("/docs/", http.StripPrefix("/docs", docs.NewHandler()))

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterCartServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts); err != nil {
		log.Fatalf("Failed to start REST gateway: %v", err)
	}

	httpPort := ":8080"
	log.Printf("start HTTP server on port %s\n", httpPort)
	if err := http.ListenAndServe(httpPort, router); err != nil {
		log.Fatalf("Failed to serve REST gateway: %v", err)
	}
}
