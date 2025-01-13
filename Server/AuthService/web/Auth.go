package web

import (
	"AuthService/model"
	pb "AuthService/protos"
	"context"
	"database/sql"
	"time"

	"github.com/hashicorp/go-hclog"
)

type Auth struct {
	pb.UnimplementedAuthServiceServer
	log      hclog.Logger
	DB       *sql.DB
	Secret   string
	Duration time.Duration
}

func NewAuth(l hclog.Logger, db *sql.DB, secret string, duration time.Duration) *Auth {
	return &Auth{log: l, DB: db, Secret: secret, Duration: duration}
}

func (a *Auth) Register(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	a.log.Info("Handle register", "Login", req.GetUser().GetLogin())
	//logic
	var token string
	login := req.GetUser().GetLogin()
	password := req.GetUser().GetPassword()
	nickname := req.GetUser().GetDetails().GetNickname()
	u := model.NewUser(login, password, model.NewDetails(nickname))

	//Добавление пользователя в базу данных
	err := a.AddUser(login, password, nickname)
	if err != nil {
		return returnErr(err)
	}

	//Генерация jwt токена
	token, err = a.GenerateToken(u)
	if err != nil {
		return returnErr(err)
	}

	//end logic
	return returnSuccses(token), nil
}

func (a *Auth) Login(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	a.log.Info("Handle login", "Login", req.GetUser().GetLogin())

	var token string
	login := req.GetUser().GetLogin()
	password := req.GetUser().GetPassword()
	nickname := req.GetUser().GetDetails().GetNickname()
	u := model.NewUser(login, password, model.NewDetails(nickname))

	//Проверка на подлинность данных
	err := a.Authenticate(login, password)
	if err != nil {
		return returnErr(err)
	}

	//Генерация токена
	token, err = a.GenerateToken(u)
	if err != nil {
		return returnErr(err)
	}

	return returnSuccses(token), nil
}

func returnErr(err error) (*pb.AuthResponse, error) {
	return &pb.AuthResponse{Status: pb.Statuses_ERROR, SessionToken: "", ResponseMessage: err.Error()}, err
}
func returnSuccses(token string) *pb.AuthResponse {
	return &pb.AuthResponse{Status: pb.Statuses_SUCCESS, SessionToken: token, ResponseMessage: "Succsess!"}
}
