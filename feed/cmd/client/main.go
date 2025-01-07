package main

import (
	"context"
	"log"

	ssov1 "github.com/username/protos/gen/go/sso"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("sso:44044", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("init conn: %s", err)
	}

	defer conn.Close()

	client := ssov1.NewAuthClient(conn)

	resp, err := client.IsAdmin(context.Background(), &ssov1.IsAdminRequest{UserId: 1})
	if err != nil {
		log.Fatalf("error on query: %s", err)
	}

	log.Println(resp)
}