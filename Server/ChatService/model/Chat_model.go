package model

import "time"

type Chat_model struct {
	ChatID int32
	Name   string
	Date   time.Time
}

type Member_model struct {
	UserID int32
	ChatID int32
	Role   string
}

type Msg_model struct {
	Src  int32
	Dst  int32
	Data string
}
