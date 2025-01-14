package main

// Автор: Александр Каюков Игоревич, ИВТ-23, 2024 год
// номер зачетки: 220161
// В-1

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/XoneRush/gRPCmessengerGolang/Server/AuthService/model"
	pb "github.com/XoneRush/gRPCmessengerGolang/Server/AuthService/protos"
	"github.com/XoneRush/gRPCmessengerGolang/Server/AuthService/web"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Путь к свойствам для подключения к бд
const properties_path = "./properties/properties.json"

// Порт, на котором стартует сервер
const port = ":8090"

func main() {
	log := hclog.Default()

	//Получение необходимых для приложения свойств
	props := getPropertesFromJson(properties_path)
	duration, err := time.ParseDuration(props.Duration)
	if err != nil {
		log.Error("Error in parsing duration", "error", err)
	}
	//Инициализация базы данных
	db := model.Connect(props)

	//Инициализация сервера
	gs := grpc.NewServer()
	cs := web.NewAuth(log, db, props.Secret, duration)
	pb.RegisterAuthServiceServer(gs, cs)

	//Функция позволяющая клиенту видеть функции сервера
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
