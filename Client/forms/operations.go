package forms

import (
	"context"
	"errors"
	"io"
	"log"
	"strconv"
	"time"

	auth "github.com/XoneRush/gRPCmessengerGolang/Server/AuthService/protos"
	"github.com/golang-jwt/jwt/v5"
	//chat "github.com/XoneRush/gRPCmessengerGolang/Server/ChatService/protos"
)

type UserClaims struct {
	ID       int    `json:"uid"`
	login    string `json:"ulogin"`
	nickname string `json:"nickname"`
	jwt.Claims
}

// Запрос на регистрацию на сервер
func (c *Client) Register() string {
	//Из заполненных данных формируется структура для запроса
	user := auth.AuthRequest_User{
		Login:    c.UserData.Login,
		Password: c.UserData.Password,
		Details: &auth.AuthRequest_User_UserDetails{
			Nickname: c.UserData.Nickname,
		},
	}

	//Сам запрос, если все нормально, получаем токен для чата
	//Запрос делается из AuthServiceClient, который инициализируется с помощью Connect
	resp, err := c.AuthClient.Register(context.Background(), &auth.AuthRequest{User: &user})
	if err != nil {
		return err.Error()
	}

	c.token = resp.SessionToken
	return resp.ResponseMessage
	//fmt.Println(resp.ResponseMessage, " ", resp.Status)
}

func (c *Client) Login() string {
	//Из заполненных данных формируется структура для запроса
	user := auth.AuthRequest_User{
		Login:    c.UserData.Login,
		Password: c.UserData.Password,
		Details: &auth.AuthRequest_User_UserDetails{
			Nickname: c.UserData.Nickname,
		},
	}

	resp, err := c.AuthClient.Login(context.Background(), &auth.AuthRequest{User: &user})
	if err != nil {
		return err.Error()
	}

	c.token = resp.SessionToken
	return resp.ResponseMessage
	//fmt.Println(resp.ResponseMessage, " ", resp.Status)
}

func (c *Client) GetIdFromToken() string {
	claims, err := c.parseToken()
	if err != nil {
		return err.Error()
	}

	id := claims["uid"].(float64)

	return strconv.FormatFloat(id, 'g', 2, 64)
}

func (c *Client) GetNicknameFromToken() string {
	claims, err := c.parseToken()
	if err != nil {
		return err.Error()
	}

	nick := claims["nickname"].(string)
	return nick
}

// Парсинг свойств сделать
func (c *Client) parseToken() (jwt.MapClaims, error) {
	token, err := jwt.Parse(c.token, func(token *jwt.Token) (interface{}, error) {
		return []byte("d19b71f8942a1d89b1bbb1481e4df6bb770d1c746d41fe61bbf870b8252cebb0bc2f9b3718a25ba5f68119e49bc0b24e8e4b56236b3bce53f7b138f63b1846f1"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("jwt token isnt valid")
}

// Старт потока обмена сообщениями
func (c *Client) StartMessaging() {
	c.MsgStream, _ = c.ChatClient.SendMessage(context.Background())
	c.Waitc = make(chan struct{})

	go func() {
		for {
			in, err := c.MsgStream.Recv()
			if err == io.EOF {
				close(c.Waitc)
				return
			}
			if err != nil {
				log.Println("Failed to receive a note", false)
			}

			//msgs
			currText := c.chat.GetText(false)
			c.chat.SetText(currText + in.Data + "\n")
			currTime := c.time.GetText(false)
			c.time.SetText(currTime + time.Now().Format(time.TimeOnly) + "\n")

		}
	}()
}

// func (c *Client) GetChatList() []ChatData{
// 	stream, err := c.ChatClient.

// 	for {
// 		chat, err := stream
// 	}
// }
