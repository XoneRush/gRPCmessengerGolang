package web

import (
	"log"

	pb "github.com/XoneRush/gRPCmessengerGolang/Server/AuthService/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

// Подключение к сервису Аутенфикации
//
// Возвращает объект клиента, из которого можно делать запросы
func ConnectWithAuth() pb.AuthServiceClient {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("localhost:8090", opts...)
	if err != nil {
		grpclog.Fatalf("failed to connect %v", err)
	}

	log.Print("Succses!")

	AuthClient := pb.NewAuthServiceClient(conn)

	return AuthClient
}

func ConnectWithChats() {

}
