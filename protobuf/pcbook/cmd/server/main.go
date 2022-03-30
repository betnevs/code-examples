package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc/reflection"

	"github.com/betNevS/code-examples/protobuf/pcbook/pb"
	"github.com/betNevS/code-examples/protobuf/pcbook/service"
	"google.golang.org/grpc"
)

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

const (
	serverCertFile = "cert/server-cert.pem"
	serverKeyFile  = "cert/server-key.pem"
)

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/techschool.pcbook.LaptopService/"
	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin"},
		laptopServicePath + "UploadImage":  {"admin"},
		laptopServicePath + "RateLaptop":   {"admin", "user"},
	}
}

func seedUser(userStore service.UserStore) error {
	err := createUser(userStore, "admin1", "123", "admin")
	if err != nil {
		return err
	}
	return createUser(userStore, "user1", "123", "user")
}

func createUser(userStore service.UserStore, username, password, role string) error {
	user, err := service.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		return nil, err
	}
	config := tls.Config{
		Certificates: []tls.Certificate{serverCert},
	}
	return credentials.NewTLS(&config), nil
}

func runGRPCServer(authServer pb.AuthServiceServer, laptopServer pb.LaptopServiceServer,
	jwtManager *service.JWTManager, enableTLS bool, listener net.Listener) error {
	interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())

	serverOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}

	if enableTLS {
		tlsCredential, err := loadTLSCredentials()
		if err != nil {
			return fmt.Errorf("cannot load TLS credentials: %w", err)
		}
		serverOpts = append(serverOpts, grpc.Creds(tlsCredential))
	}

	grpcServer := grpc.NewServer(serverOpts...)

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	pb.RegisterAuthServiceServer(grpcServer, authServer)
	reflection.Register(grpcServer)

	log.Printf("Start gRPC server at %s, TLS = %t", listener.Addr().String(), enableTLS)
	return grpcServer.Serve(listener)
}

func runRESTServer(authServer pb.AuthServiceServer, laptopServer pb.LaptopServiceServer,
	jwtManager *service.JWTManager, enableTLS bool, listener net.Listener) error {
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := pb.RegisterAuthServiceHandlerServer(ctx, mux, authServer)
	if err != nil {
		return err
	}

	err = pb.RegisterLaptopServiceHandlerServer(ctx, mux, laptopServer)
	if err != nil {
		return err
	}

	log.Printf("Start REST server at %s, TLS = %t", listener.Addr().String(), enableTLS)
	if enableTLS {
		return http.ServeTLS(listener, mux, serverCertFile, serverKeyFile)
	}
	return http.Serve(listener, mux)
}

func main() {
	port := flag.Int("port", 0, "the server port")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	serveType := flag.String("type", "grpc", "type of server(grpc/rest)")
	flag.Parse()

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingStore()

	userStore := service.NewInMemoryUserStore()
	err2 := seedUser(userStore)
	if err2 != nil {
		log.Fatal(err2)
	}
	jwtManager := service.NewJWTManager(secretKey, tokenDuration)
	authServer := service.NewAuthServer(userStore, jwtManager)

	laptopServer := service.NewLaptopServer(laptopStore, imageStore, ratingStore)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	log.Println("start server address:", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	if *serveType == "grpc" {
		err = runGRPCServer(authServer, laptopServer, jwtManager, *enableTLS, listener)
	} else {
		err = runRESTServer(authServer, laptopServer, jwtManager, *enableTLS, listener)
	}
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
