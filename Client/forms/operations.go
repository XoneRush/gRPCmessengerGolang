package forms

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	auth "github.com/XoneRush/gRPCmessengerGolang/Server/AuthService/protos"
	ChatService_proto "github.com/XoneRush/gRPCmessengerGolang/Server/ChatService/protos"
	"github.com/golang-jwt/jwt/v5"
	//chat "github.com/XoneRush/gRPCmessengerGolang/Server/ChatService/protos"
)

type UserClaims struct {
	ID       int    `json:"uid"`
	login    string `json:"ulogin"`
	nickname string `json:"nickname"`
	jwt.Claims
}

type Properties struct {
	Secret string `json:"secret"`
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
		return []byte(c.Properties.Secret), nil
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

// Заполняет список чатов для юзера
func (c *Client) GetChatList() error {
	if c.chat != nil {
		c.ChatList.Clear()
	}

	var chats []ChatData
	stream, err := c.ChatClient.GetChatList(context.Background(), &ChatService_proto.Member{
		UserID: int32(c.UserData.ID),
	})
	if err != nil {
		return err
	}

	for {
		chat, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		chats = append(chats, ChatData{
			Name: chat.GetName(),
			ID:   chat.GetChatID(),
		})
	}

	c.Chats = chats
	return nil
}

func (c *Client) refresh() {
	c.AddChatList()
	c.AddFlex()
}

func (c *Client) ClearMsgs() {
	c.chat.Clear()
	c.time.Clear()
}

func (c *Client) ParseProperties(path string) Properties {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Cant open prop file")
	}
	reader := bufio.NewReader(file)

	decoder := json.NewDecoder(reader)
	var props Properties
	err = decoder.Decode(&props)
	if err != nil {
		fmt.Println("error in decoding")
	}
	return props
}
