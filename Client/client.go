package main

import (
	"Client/forms"
	"Client/web"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	//pb "github.com/XoneRush/gRPCmessengerGolang/Server/AuthService"
)

var cli forms.Client

const properties_path = "./properties/properties.json"

func main() {
	cli = forms.Client{
		App:          tview.NewApplication(),
		IndexForm:    tview.NewForm(),
		RegisterForm: tview.NewForm(),
		LoginForm:    tview.NewForm(),
		Pages:        tview.NewPages(),
		ChatList:     tview.NewList(),
		Chats:        []forms.ChatData{},
		UserData:     forms.UserData{},
	}

	Startup(&cli)

	if err := cli.App.SetRoot(cli.Pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	cli.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 123 {
			cli.Pages.SwitchToPage("Index")
		}
		return event
	})

}

func Startup(c *forms.Client) {
	c.ParseProperties(properties_path)
	ConnectWithServices(c)
	AddForms(c)
}

func ConnectWithServices(c *forms.Client) {
	c.AuthClient = web.ConnectWithAuth()
	c.ChatClient = web.ConnectWithChats()
}

func AddForms(c *forms.Client) {
	c.AddIndexForm()
	c.AddRegisterForm()
	c.AddLoginForm()
	c.AddChatList()
	c.AddFlex()
}
