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

	for {
		laptop := sample.NewLaptop()
		req := &pb.CreateLaptopRequest{
			Laptop: laptop,
		}
		res, err := laptopClient.CreateLaptop(context.Background(), req)
		if err != nil {
			st, ok := status.FromError(err)
			if ok && st.Code() == codes.AlreadyExists {
				log.Println("laptop already exists")
			} else {
				log.Println("cannot create laptop: ", err)
			}
		}

		log.Printf("created laptop with id: %v", res)
		time.Sleep(time.Second)
	}

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
