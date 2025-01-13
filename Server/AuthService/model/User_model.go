package model

type User_model struct {
	login    string
	password string
	details  User_details
}

type User_details struct {
	nickname string
}

func (u *User_model) GetLogin() string {
	return u.login
}

func (u *User_model) GetNickname() string {
	return u.details.nickname
}

func NewUser(login string, password string, details User_details) User_model {
	user := User_model{
		login:    login,
		password: password,
		details:  details,
	}
	return user
}

func NewDetails(nickname string) User_details {
	det := User_details{
		nickname: nickname,
	}
	return det
}
