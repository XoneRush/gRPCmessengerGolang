package main

import (
	pb "AuthService/protos"
	"context"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	log := hclog.Default()

	conn, err := grpc.NewClient("localhost:8090", opts...)
	if err != nil {
		grpclog.Fatalf("failed to connect %v", err)
	}

	log.Info("Connected to server!")

	client := pb.NewAuthServiceClient(conn)

	user := pb.AuthRequest_User{
		Login:    "Alex2",
		Password: "123123",
		Details: &pb.AuthRequest_User_UserDetails{
			Nickname: "",
		},
	}

	resp, err := client.Login(context.Background(), &pb.AuthRequest{User: &user})

	token := resp.SessionToken

	log.Info("SUCCESS", "token", token)

}
