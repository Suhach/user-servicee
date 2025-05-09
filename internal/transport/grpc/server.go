package grpc

import (
	"fmt"
	userpb "github.com/Suhach/protoc-cont/proto"
	"github.com/Suhach/user-servicee/internal/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RunGRPC(svc user.Service) error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer listener.Close()

	grpcSrv := grpc.NewServer()

	userHandler := NewHandler(svc)
	userpb.RegisterUserServiceServer(grpcSrv, userHandler)

	log.Printf("gRPC server listening at %v", listener.Addr())
	if err := grpcSrv.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
