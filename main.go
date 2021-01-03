package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpcClient/services"
	"log"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "localhost")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	prodClient := services.NewProdServiceClient(conn)
	prodResponse, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 12})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(prodResponse)

}
