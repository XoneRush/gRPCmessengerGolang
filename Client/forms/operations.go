package forms

import (
	"context"

	pb "github.com/XoneRush/gRPCmessengerGolang/Server/AuthService/protos"
)

// Запрос на регистрацию
func (c *Client) Register() string {
	//Из заполненных данных формируется структура для запроса
	user := pb.AuthRequest_User{
		Login:    c.UserData.Login,
		Password: c.UserData.Password,
		Details: &pb.AuthRequest_User_UserDetails{
			Nickname: c.UserData.Nickname,
		},
	}

	//Сам запрос, если все нормально, получаем токен для чата
	//Запрос делается из AuthServiceClient, который инициализируется с помощью Connect
	resp, err := c.AuthClient.Register(context.Background(), &pb.AuthRequest{User: &user})
	if err != nil {
		return err.Error()
	}

	c.token = resp.SessionToken
	return resp.ResponseMessage
	//fmt.Println(resp.ResponseMessage, " ", resp.Status)
}

func (c *Client) Login() string {
	//Из заполненных данных формируется структура для запроса
	user := pb.AuthRequest_User{
		Login:    c.UserData.Login,
		Password: c.UserData.Password,
		Details: &pb.AuthRequest_User_UserDetails{
			Nickname: c.UserData.Nickname,
		},
	}

	resp, err := c.AuthClient.Login(context.Background(), &pb.AuthRequest{User: &user})
	if err != nil {
		return err.Error()
	}

	c.token = resp.SessionToken
	return resp.ResponseMessage
	//fmt.Println(resp.ResponseMessage, " ", resp.Status)
}
