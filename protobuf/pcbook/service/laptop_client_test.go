package service

import (
	"context"
	"io"
	"net"
	"testing"

	"github.com/betNevS/code-examples/protobuf/pcbook/serializer"

	"github.com/betNevS/code-examples/protobuf/pcbook/sample"

	"github.com/betNevS/code-examples/protobuf/pcbook/pb"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopStore := NewInMemoryLaptopStore()

	serverAddress := startTestLaptopServer(t, laptopStore, nil)
	laptopClient := startTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	other, err := laptopStore.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	requireSameLaptop(t, laptop, other)
}

func TestClientSearchLaptop(t *testing.T) {
	t.Parallel()

	filter := &pb.Filter{
		MaxPriceUsd: 2000,
		MinCpuCores: 4,
		MinCpuGhz:   2.2,
		MinRam: &pb.Memory{
			Value: 8,
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	store := NewInMemoryLaptopStore()
	expectedID := make(map[string]bool)

	for i := 0; i < 6; i++ {
		laptop := sample.NewLaptop()
		switch i {
		case 0:
			laptop.PriceUsd = 2500
		case 1:
			laptop.Cpu.NumberCores = 2
		case 2:
			laptop.Cpu.MinGhz = 2.0
		case 3:
			laptop.Memory = &pb.Memory{
				Value: 4600,
				Unit:  pb.Memory_MEGABYTE,
			}
		case 4:
			laptop.PriceUsd = 1999
			laptop.Cpu.NumberCores = 4
			laptop.Cpu.MinGhz = 2.5
			laptop.Cpu.MaxGhz = 4.5
			laptop.Memory = &pb.Memory{
				Value: 16,
				Unit:  pb.Memory_GIGABYTE,
			}
			expectedID[laptop.Id] = true
		case 5:
			laptop.PriceUsd = 2000
			laptop.Cpu.NumberCores = 6
			laptop.Cpu.MinGhz = 2.8
			laptop.Cpu.MaxGhz = 5.0
			laptop.Memory = &pb.Memory{
				Value: 64,
				Unit:  pb.Memory_GIGABYTE,
			}
			expectedID[laptop.Id] = true
		}

		err := store.Save(laptop)
		require.NoError(t, err)
	}

	serverAddress := startTestLaptopServer(t, store, nil)
	laptopClient := startTestLaptopClient(t, serverAddress)

	req := &pb.SearchLaptopRequest{Filter: filter}
	stream, err := laptopClient.SearchLaptop(context.Background(), req)
	require.NoError(t, err)

	found := 0
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		require.Contains(t, expectedID, res.GetLaptop().GetId())
		found++
	}
	require.Equal(t, len(expectedID), found)
}

func requireSameLaptop(t *testing.T, laptop *pb.Laptop, other *pb.Laptop) {
	json1, err := serializer.ProbufToJson(laptop)
	require.NoError(t, err)

	json2, err := serializer.ProbufToJson(other)
	require.NoError(t, err)

	require.Equal(t, json1, json2)

}

func startTestLaptopClient(t *testing.T, address string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}

func startTestLaptopServer(t *testing.T, laptopStore LaptopStore, imageStore ImageStore) string {
	laptopServer := NewLaptopServer(laptopStore, imageStore)

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}
