package main

import (
	"github.com/XoneRush/gRPCmessengerGolang/Server/ChatService/model"
	"github.com/XoneRush/gRPCmessengerGolang/Server/ChatService/web"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/XoneRush/gRPCmessengerGolang/Server/ChatService/protos"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const properties_path = "./properties/properties.json"
const port = ":8091"

func main() {
	log := hclog.Default()

	//Получение необходимых для приложения свойств
	props := getPropertesFromJson(properties_path)

	//Инициализация базы данных
	db := model.Connect(props)

	//Инициализация сервера
	gs := grpc.NewServer()
	cs := web.NewApp(log, db)
	pb.RegisterChatServiceServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Error("Unable to listen", "error", err)
	} else {
		log.Info("Server started", "port", port, "time", time.Now().Format(time.RFC850))
	}

	gs.Serve(l)
}

func getPropertesFromJson(path string) model.Properties {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Cant open prop file")
	}
	reader := bufio.NewReader(file)

	decoder := json.NewDecoder(reader)
	var props model.Properties
	err = decoder.Decode(&props)
	if err != nil {
		fmt.Println("error in decoding")
	}
	return props
}
