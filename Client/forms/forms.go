package forms

import "github.com/rivo/tview"

type Client struct {
	App          *tview.Application
	Pages        *tview.Pages
	IndexForm    *tview.Form
	RegisterForm *tview.Form
	UserData     UserData
	token        string
}

type UserData struct {
	Nickname string
	Login    string
	Password string
}

//Стартовая страничка
func (c *Client) AddIndexForm() {
	text := tview.NewTextView().SetText("Kayukov A.I. IVT-23 messenger. Choose option:")
	c.IndexForm.AddFormItem(text)

	c.IndexForm.AddButton("Register page", func() {
		c.Pages.SwitchToPage("Register page")
	})

	c.Pages.AddPage("Index", c.IndexForm, true, true)
}

//Форма для регистрации
func (c *Client) AddRegisterForm() {

	c.RegisterForm.AddInputField("Login", "", 20, nil, func(login string) {
		c.UserData.Login = login
	})
	c.RegisterForm.AddInputField("Password", "", 20, nil, func(password string) {
		c.UserData.Password = password
	})
	c.RegisterForm.AddInputField("Login", "", 20, nil, func(Nickname string) {
		c.UserData.Nickname = Nickname
	})

	c.RegisterForm.AddButton("Register", func() {
		//Логика регистрации
	})

	c.Pages.AddPage("Register", c.RegisterForm, true, false)
}

func ChatForm() {

}
