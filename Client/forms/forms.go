package forms

import "github.com/rivo/tview"

type Client struct {
	App       *tview.Application
	IndexForm *tview.Form
	UserData  UserData
}

type UserData struct {
	Nickname string
	Login    string
	Password string
}

func (c *Client) AddIndexForm(form *tview.Form) {

	form.AddInputField("Login", "", 20, nil, func(login string) {
		c.UserData.Login = login
	})
}

func ChatForm() {

}
