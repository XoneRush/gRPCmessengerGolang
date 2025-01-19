package web

import (
	"log"

	auth "github.com/XoneRush/gRPCmessengerGolang/Server/AuthService/protos"
	chat "github.com/XoneRush/gRPCmessengerGolang/Server/ChatService/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

// Подключение к сервису Аутенфикации
//
// Возвращает объект клиента, из которого можно делать запросы
func ConnectWithAuth() auth.AuthServiceClient {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("localhost:8090", opts...)
	if err != nil {
		grpclog.Fatalf("failed to connect %v", err)
	}

	log.Print("Succses!")

	AuthClient := auth.NewAuthServiceClient(conn)

	return AuthClient
}

// Подключение к сервису чатов
//
// Возвращает объект клиента, из которого можно делать запросы
func ConnectWithChats() chat.ChatServiceClient {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("localhost:8091", opts...)
	if err != nil {
		grpclog.Fatalf("failed to connect %v", err)
	}

	log.Print("Succses!")

	ChatClient := chat.NewChatServiceClient(conn)

	return ChatClient
}
