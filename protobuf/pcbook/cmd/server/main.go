package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/betNevS/code-examples/protobuf/pcbook/pb"
	"github.com/betNevS/code-examples/protobuf/pcbook/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingStore()

	laptopServer := service.NewLaptopServer(laptopStore, imageStore, ratingStore)
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("127.0.0.1:%d", *port)
	log.Println("start server address:", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
