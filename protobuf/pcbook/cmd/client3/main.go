package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/betNevS/code-examples/protobuf/pcbook/sample"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/betNevS/code-examples/protobuf/pcbook/pb"
	"google.golang.org/grpc"
)

func main() {
	serverAddr := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddr)

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	laptopClient := pb.NewLaptopServiceClient(conn)

	testRateLaptop(laptopClient)
}

func testUploadImage(laptopClient pb.LaptopServiceClient) {
	laptop := sample.NewLaptop()
	createLaptop(laptopClient, laptop)
	uploadImage(laptopClient, laptop.GetId(), "tmp/yj1.jpeg")
}

func uploadImage(laptopClient pb.LaptopServiceClient, laptopID string, imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := laptopClient.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptopId:  laptopID,
				ImageType: filepath.Ext(imagePath),
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send image info: ", err, stream.RecvMsg(nil))
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			err2 := stream.RecvMsg(nil)
			log.Fatal("cannot send chunk to server: ", err, err2)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	log.Printf("image uploaded with id: %s, size: %d", res.GetId(), res.GetSize())

}

func testSearchLaptop(laptopClient pb.LaptopServiceClient) {
	for i := 0; i < 10; i++ {
		createLaptop(laptopClient, sample.NewLaptop())
	}

	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam: &pb.Memory{
			Value: 8,
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	searchLaptop(laptopClient, filter)
}

func searchLaptop(laptopClient pb.LaptopServiceClient, filter *pb.Filter) {
	log.Println("search filter: ", filter)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SearchLaptopRequest{
		Filter: filter,
	}
	stream, err := laptopClient.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatal("cannot search laptop: ", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}
		laptop := res.GetLaptop()
		log.Println("- found: ", laptop.GetId())
		log.Println("  + brand: ", laptop.GetBrand())
		log.Println("  + name: ", laptop.GetName())
		log.Println("  + cpu cores: ", laptop.GetCpu().GetNumberCores())
		log.Println("  + cpu min ghz: ", laptop.GetCpu().GetMinGhz())
		log.Println("  + ram: ", laptop.GetMemory().GetValue(), laptop.GetMemory().GetUnit())
		log.Println("  + price: ", laptop.GetPriceUsd(), "usd")
	}
}

func createLaptop(client pb.LaptopServiceClient, laptop *pb.Laptop) {
	//laptop.Id = ""
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Println("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
	}
	log.Printf("created laptop with id: %v", res.Id)
}

func testRateLaptop(laptopClient pb.LaptopServiceClient) {
	n := 3
	laptopIDs := make([]string, n)

	for i := 0; i < n; i++ {
		laptop := sample.NewLaptop()
		laptopIDs[i] = laptop.GetId()
		createLaptop(laptopClient, laptop)
	}

	scores := make([]float64, n)
	for {
		fmt.Println("rate laptop (y/n)?")
		var answer string
		fmt.Scan(&answer)

		if strings.ToLower(answer) != "y" {
			break
		}

		for i := 0; i < n; i++ {
			scores[i] = sample.RandomLaptopScore()
		}

		err := rateLaptop(laptopClient, laptopIDs, scores)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func rateLaptop(laptopClient pb.LaptopServiceClient, laptopIDs []string, scores []float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := laptopClient.RateLaptop(ctx)
	if err != nil {
		return fmt.Errorf("cannot rate laptop: %v", err)
	}

	waitResponse := make(chan error)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("no more response")
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("cannot receive stream response: %v", err)
				return
			}

			log.Println("received response: ", res)
		}
	}()

	for i, laptopID := range laptopIDs {
		req := &pb.RateLaptopRequest{
			LaptopId: laptopID,
			Score:    scores[i],
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send stream request: %v - %v", err, stream.RecvMsg(nil))
		}
		log.Println("sent request: ", req)
	}
	err = stream.CloseSend()

	if err != nil {
		return fmt.Errorf("cannot close send: %v", err)
	}
	err = <-waitResponse
	return err
}
