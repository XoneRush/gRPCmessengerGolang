package forms

import (
	"context"
	"io"
	"log"
	"strconv"
	"time"

	auth "github.com/XoneRush/gRPCmessengerGolang/Server/AuthService/protos"
	chat "github.com/XoneRush/gRPCmessengerGolang/Server/ChatService/protos"
	"github.com/rivo/tview"
)

type Client struct {
	App          *tview.Application
	AuthClient   auth.AuthServiceClient
	ChatClient   chat.ChatServiceClient
	Pages        *tview.Pages
	IndexForm    *tview.Form
	RegisterForm *tview.Form
	LoginForm    *tview.Form
	ChatForm     *tview.Form
	UserData     UserData
	token        string
}

type UserData struct {
	ID       int
	Nickname string
	Login    string
	Password string
}

// Стартовая страничка
func (c *Client) AddIndexForm() {
	text := tview.NewTextView().SetText("Kayukov A.I. IVT-23 messenger. Choose option:")
	c.IndexForm.AddFormItem(text)

	c.IndexForm.AddButton("Register page", func() {
		c.Pages.SwitchToPage("Register page")
	})

	c.IndexForm.AddButton("Login page", func() {
		c.Pages.SwitchToPage("Login page")
	})

	c.IndexForm.AddButton("Chat page", func() {
		c.Pages.SwitchToPage("Chat page")
	})

	c.Pages.AddPage("Index", c.IndexForm, true, true)
}

// Форма для регистрации
func (c *Client) AddRegisterForm() {
	textArea := tview.NewTextArea()
	textArea.SetSize(10, 50)
	textArea.SetLabel("Response")
	c.RegisterForm.AddFormItem(textArea)

	c.RegisterForm.AddInputField("Login", "", 20, nil, func(login string) {
		c.UserData.Login = login
	})
	c.RegisterForm.AddInputField("Password", "", 20, nil, func(password string) {
		c.UserData.Password = password
	})
	c.RegisterForm.AddInputField("Nickname", "", 20, nil, func(Nickname string) {
		c.UserData.Nickname = Nickname
	})

	c.RegisterForm.AddButton("Register", func() {
		//Логика регистрации
		resp := c.Register()
		textArea.SetText(resp, false)

	})

	c.RegisterForm.AddButton("Back", func() {
		c.Pages.SwitchToPage("Index")
	})

	c.Pages.AddPage("Register page", c.RegisterForm, true, false)
}

func (c *Client) AddLoginForm() {
	textArea := tview.NewTextArea()
	textArea.SetSize(10, 50)
	c.LoginForm.AddFormItem(textArea)

	c.LoginForm.AddInputField("Login", "", 20, nil, func(login string) {
		c.UserData.Login = login
	})

	c.LoginForm.AddInputField("Password", "", 20, nil, func(password string) {
		c.UserData.Password = password
	})

	c.LoginForm.AddButton("Login", func() {
		resp := c.Login()
		textArea.SetText(resp, false)
		c.UserData.ID, _ = strconv.Atoi(c.GetIdFromToken())
	})

	c.LoginForm.AddButton("Back", func() {
		c.Pages.SwitchToPage("Index")
	})

	c.Pages.AddPage("Login page", c.LoginForm, true, false)
}

func (c *Client) AddChatForm() {
	var msg string = ""
	var dst int = 1
	//var chatName string
	textArea := tview.NewTextArea()
	textArea.SetSize(10, 50)
	textArea.SetBorder(true)
	c.ChatForm.AddFormItem(textArea)

	stream, _ := c.ChatClient.SendMessage(context.Background())
	waitc := make(chan struct{})

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Println("Failed to receive a note", false)
			}

			//msgs
			currText := textArea.GetText()
			textArea.SetText(currText+in.Data+"       "+time.Now().Format(time.TimeOnly)+"\n", false)

		}
	}()

	//Выбрать имя чата (имя получателя)
	c.ChatForm.AddInputField("Which chat?", "", 20, nil, func(name string) {
		//chatName = name
		dst, _ = strconv.Atoi(name)
	})

	//Само сообщение
	c.ChatForm.AddInputField("Message", "", 20, nil, func(message string) {
		msg = message
	})

	//Отправить
	c.ChatForm.AddButton("Send!", func() {
		//логика отправки
		if err := stream.Send(&chat.Msg{Dst: int32(dst), Src: int32(c.UserData.ID), Data: msg}); err != nil {
			textArea.SetText("error!", false)
		}
	})

	c.ChatForm.AddButton("Back", func() {
		c.Pages.SwitchToPage("Index")
	})

	c.Pages.AddPage("Chat page", c.ChatForm, true, false)
}
