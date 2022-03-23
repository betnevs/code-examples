package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/betNevS/code-examples/protobuf/pcbook/pb"
	"github.com/betNevS/code-examples/protobuf/pcbook/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//var kacp = keepalive.ClientParameters{
//	Time:                10 * time.Second,
//	Timeout:             time.Second,
//	PermitWithoutStream: true,
//}

func main() {
	serverAddr := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddr)

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dail server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	laptop.Id = "9f3d84d5-8929-455b-9461-f130a277fe5b"
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Println("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}
	log.Printf("created laptop with id: %s", res.Id)

	//time.Sleep(20 * time.Second)

	//log.Println("second request")
	//
	//res, err = laptopClient.CreateLaptop(ctx, req)
	//if err != nil {
	//	st, ok := status.FromError(err)
	//	if ok && st.Code() == codes.AlreadyExists {
	//		log.Println("laptop already exists")
	//	} else {
	//		log.Println("cannot create laptop: ", err)
	//	}
	//}
	//log.Printf("created laptop with id: %s", res.Id)

	//fmt.Println("laptop: ", res)
	//
	//time.Sleep(time.Second)
	//log.Println("third request")
	//
	//res, err = laptopClient.CreateLaptop(ctx, req)
	//if err != nil {
	//	st, ok := status.FromError(err)
	//	if ok && st.Code() == codes.AlreadyExists {
	//		log.Println("laptop already exists")
	//	} else {
	//		log.Println("cannot create laptop: ", err)
	//	}
	//}
	//fmt.Println("laptop: ", res)
}
