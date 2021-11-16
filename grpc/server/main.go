package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/betNevS/code-examples/grpc/helloworld"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Recv: %v", req.GetName())
	return &pb.HelloReply{Message: "hello" + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen", port)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("serve failed, err:", err)
	}
}
