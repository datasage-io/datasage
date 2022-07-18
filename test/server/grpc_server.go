package main

import (
	"context"
	"fmt"
	"log"
	"net"

	config "github.com/datasage-io/datasage/src/integrations/grpc_config"
	"google.golang.org/grpc"
)

type server struct {
	config.UnimplementedDataSageServerServer
}

func (s *server) LogSend(ctx context.Context, in *config.Log) (*config.Null, error) {
	log.Printf("GRPC: Receive message body from client: %s", in.Body)
	return &config.Null{}, nil
}

func RunGRPCServer() {
	listen, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	fmt.Println("gRPC Server listening on :2222")
	config.RegisterDataSageServerServer(grpcServer, &server{})
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func main() {
	RunGRPCServer()
}
