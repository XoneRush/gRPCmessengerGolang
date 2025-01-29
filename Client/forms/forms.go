package forms

import (
	"strconv"

	auth "github.com/XoneRush/gRPCmessengerGolang/Server/AuthService/protos"
	chat "github.com/XoneRush/gRPCmessengerGolang/Server/ChatService/protos"
	"github.com/rivo/tview"
	"google.golang.org/grpc"
)

type Client struct {
	App          *tview.Application
	AuthClient   auth.AuthServiceClient
	ChatClient   chat.ChatServiceClient
	Pages        *tview.Pages
	IndexForm    *tview.Form
	RegisterForm *tview.Form
	LoginForm    *tview.Form
	ChatList     *tview.List
	UserData     UserData
	token        string

	Chats []ChatData

	MsgStream grpc.BidiStreamingClient[chat.Msg, chat.Msg]
	Waitc     chan struct{}
	chat      *tview.TextView
	dst       int
	time      *tview.TextView
}

type UserData struct {
	ID       int
	Nickname string
	Login    string
	Password string
}

type ChatData struct {
	ID   int32
	Name string
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

	c.IndexForm.AddButton("Chats", func() {
		c.StartMessaging()
		c.Pages.SwitchToPage("Chats")
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

func (c *Client) AddChatList() {
	c.ChatList.ShowSecondaryText(false)
	c.ChatList.SetBorder(true)
	c.ChatList.SetTitle("Chat list")

	c.ChatList.InsertItem(0, "FR, god i hate those lists", "", 0, func() {
		c.dst = 1
	})
	c.ChatList.InsertItem(0, "2", "", 0, func() {

	})
	c.ChatList.InsertItem(0, "3", "", 0, func() {

	})

	// flex.AddItem(c.ChatList, 0, 1, true)
	c.Pages.AddPage("Chat list", c.ChatList, true, false)
}

func (c *Client) AddFlex() {
	//flexes
	// flex & list -> subflex & Input -> chatflex & time
	flex := tview.NewFlex()
	subflex := tview.NewFlex()
	chatflex := tview.NewFlex()
	chatflex.SetBorder(true)
	chatflex.SetTitle("Chat")

	c.chat = c.AddChatArea(30, 50)
	c.time = tview.NewTextView()

	//chatflex
	chatflex.AddItem(c.chat, 0, 3, true)
	chatflex.AddItem(c.time, 0, 1, true)

	//subflex
	subflex.SetDirection(tview.FlexRow)
	subflex.AddItem(chatflex, 0, 4, true)
	subflex.AddItem(c.AddInput(), 0, 1, true)

	//flex
	flex.AddItem(subflex, 0, 3, true)
	flex.AddItem(c.ChatList, 0, 1, true)

	c.Pages.AddPage("Chats", flex, true, false)

}

func (c *Client) AddChatArea(rows int, cols int) *tview.TextView {
	chatArea := tview.NewTextView()
	chatArea.SetSize(rows, cols)

	return chatArea
}

func (c *Client) AddInput() *tview.Form {
	form := tview.NewForm()
	form.SetBorder(true)
	var msg string

	form.AddInputField("Message", "", 75, nil, func(text string) {
		msg = text
	})

	form.AddButton("->", func() {
		//Делается запрос на сервер с отправлением сообщения Dst - получаетль (id чата), src - отправитель, data - сообщение
		if err := c.MsgStream.Send(&chat.Msg{Dst: int32(c.dst), Src: int32(c.UserData.ID), Data: msg}); err != nil {
			c.chat.SetText("error!")
		}
	})

	//При выходе в меню поток обмена сообщениями закрывается
	form.AddButton("Back", func() {
		c.MsgStream.CloseSend()
		<-c.Waitc
		c.Pages.SwitchToPage("Index")
	})

	return form
}
