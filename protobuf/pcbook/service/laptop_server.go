package service

import (
	"context"
	"log"

	"github.com/pkg/errors"

	"google.golang.org/grpc/codes"

	"github.com/google/uuid"
	"google.golang.org/grpc/status"

	"github.com/betNevS/code-examples/protobuf/pcbook/pb"
)

type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	laptopStore LaptopStore
	imageStore  ImageStore
}

func NewLaptopServer(laptopStore LaptopStore, imageStore ImageStore) *LaptopServer {
	return &LaptopServer{
		laptopStore: laptopStore,
		imageStore:  imageStore,
	}
}

func (s *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.GetId())
	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}

	// some heavy processing
	//time.Sleep(6 * time.Second)
	// important
	if ctx.Err() == context.Canceled {
		log.Println("request is canceld")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}
	// important
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("deadline is exceeded - yangjie")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded - yangjie")
	}

	// save laptop to db
	err := s.laptopStore.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to store: %v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)

	return &pb.CreateLaptopResponse{Id: laptop.Id}, nil
}

func (s *LaptopServer) SearchLaptop(req *pb.SearchLaptopRequest, stream pb.LaptopService_SearchLaptopServer) error {
	filter := req.GetFilter()
	log.Printf("receive a search-laptop request with filter: %v", filter)

	err := s.laptopStore.Search(stream.Context(), filter, func(laptop *pb.Laptop) error {
		res := &pb.SearchLaptopResponse{
			Laptop: laptop,
		}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		log.Printf("sent laptop with id: %s", laptop.GetId())
		return nil
	})
	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}
	return nil
}

func (s *LaptopServer) UploadImage(stream *pb.LaptopService_UploadImageServer) error {

	return nil
}

func logError(err error) error {
	if err != nil {
		log.Println(err)
	}
	return err
}
