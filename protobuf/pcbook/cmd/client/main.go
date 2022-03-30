package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/betNevS/code-examples/protobuf/pcbook/client"
	"google.golang.org/grpc/credentials"

	"github.com/betNevS/code-examples/protobuf/pcbook/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	username        = "admin1"
	password        = "123"
	refreshDuration = 30 * time.Second
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server'CA certificate")
	}
	config := &tls.Config{
		RootCAs: certPool,
	}
	return credentials.NewTLS(config), nil
}

func main() {
	serverAddr := flag.String("address", "", "the server address")
	enableTLS := flag.Bool("tls", false, "enable TLS/SSL")

	flag.Parse()
	log.Printf("dial server %s", *serverAddr)

	transportOption := grpc.WithInsecure()

	if *enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		transportOption = grpc.WithTransportCredentials(tlsCredentials)
	}

	cc1, err := grpc.Dial(*serverAddr, transportOption)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	authClient := client.NewAuthClient(cc1, username, password)
	interceptor, err := client.NewAuthInterceptor(authClient, authMethod(), refreshDuration)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	cc2, err := grpc.Dial(
		*serverAddr,
		transportOption,
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := client.NewLaptopClient(cc2)

	for {
		laptop := sample.NewLaptop()
		laptopClient.CreateLaptop(laptop)
		if err != nil {
			st, ok := status.FromError(err)
			if ok && st.Code() == codes.AlreadyExists {
				log.Println("laptop already exists")
			} else {
				log.Println("cannot create laptop: ", err)
			}
		}
		time.Sleep(time.Second)
	}
}

func authMethod() map[string]bool {
	const laptopServicePath = "/techschool.pcbook.LaptopService/"
	return map[string]bool{
		laptopServicePath + "CreateLaptop": true,
		laptopServicePath + "UploadImage":  true,
		laptopServicePath + "RateLaptop":   true,
	}
}
