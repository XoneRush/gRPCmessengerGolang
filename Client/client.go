package main

import (
	"Client/forms"
	"Client/web"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	//pb "github.com/XoneRush/gRPCmessengerGolang/Server/AuthService"
)

var cli forms.Client

func main() {
	cli = forms.Client{
		App:          tview.NewApplication(),
		IndexForm:    tview.NewForm(),
		RegisterForm: tview.NewForm(),
		LoginForm:    tview.NewForm(),
		ChatForm:     tview.NewForm(),
		Pages:        tview.NewPages(),
		UserData:     forms.UserData{},
	}

	Startup(&cli)

	if err := cli.App.SetRoot(cli.Pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	cli.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 123 {
			cli.App.Stop()
		}
		return event
	})

}

func Startup(c *forms.Client) {
	c.AuthClient = web.ConnectWithAuth()

	c.AddIndexForm()
	c.AddRegisterForm()
	c.AddLoginForm()
	c.AddChatForm()
}
