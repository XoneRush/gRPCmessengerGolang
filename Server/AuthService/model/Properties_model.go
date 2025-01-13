package model

type Properties struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Secret   string `json:"secret"`
	Duration string `json:"duration"`
}
