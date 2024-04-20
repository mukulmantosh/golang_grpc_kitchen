package orders

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *GRPCServer {
	return &GRPCServer{addr: addr}
}

func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Failed to listen on port %s", s.addr)
	}
	server := grpc.NewServer()

	// Register gRPC Services
	//server.RegisterService()
	return server.Serve(lis)
}
