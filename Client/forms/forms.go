package forms

import (
	pb "github.com/XoneRush/gRPCmessengerGolang/Server/AuthService/protos"
	"github.com/rivo/tview"
)

type Client struct {
	App        *tview.Application
	AuthClient pb.AuthServiceClient
	//ChatClient *pb.ChatServiceClient - доделать позже
	Pages        *tview.Pages
	IndexForm    *tview.Form
	RegisterForm *tview.Form
	LoginForm    *tview.Form
	ChatForm     *tview.Form
	UserData     UserData
	token        string
}

type UserData struct {
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
	})

	c.LoginForm.AddButton("Back", func() {
		c.Pages.SwitchToPage("Index")
	})

	c.Pages.AddPage("Login page", c.LoginForm, true, false)
}

func (c *Client) AddChatForm() {
	textArea := tview.NewTextArea()
	textArea.SetSize(10, 50)
	c.ChatForm.AddFormItem(textArea)

	//Выбрать имя чата (имя получателя)
	c.ChatForm.AddInputField("Which chat?", "", 20, nil, func(name string) {

	})

	//Само сообщение
	c.ChatForm.AddInputField("Message", "", 20, nil, func(message string) {

	})

	//Отправить
	c.ChatForm.AddButton("Send!", func() {
		//логика отправки
	})

	c.ChatForm.AddButton("Back", func() {
		c.Pages.SwitchToPage("Index")
	})

	c.Pages.AddPage("Chat page", c.ChatForm, true, false)
}
