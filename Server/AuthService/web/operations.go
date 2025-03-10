package web

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/XoneRush/gRPCmessengerGolang/Server/AuthService/model"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AddUser добавление юзера в бд
//
// В базе данных пароль будет храниться в хэшированном виде
func (a *Auth) AddUser(login string, password string, nickname string) error {
	req := "INSERT INTO users(login, password, nickname) VALUES($1, $2, $3)"

	//Генерация хэша из пароля, чтобы в базе данных он не хранился в "голом" виде
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	_, err = a.DB.Exec(req, login, passHash, nickname)
	if err != nil {
		return err
	}
	return nil
}

// Аутентификация
//
// На вход: логин, пароль
func (a *Auth) Authenticate(login string, password string) error {
	req := "SELECT userid, password FROM users WHERE login = $1"
	var id int
	var passwordHash []byte

	row := a.DB.QueryRow(req, login)
	err := row.Scan(&id, &passwordHash)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(password))
	if err != nil {
		return errors.New("Password or login are wrong!")
	}

	return nil
}

// Генерация токена
//
// Параметры: пользователь
//
// На выход: jwt токен
func (a *Auth) GenerateToken(user model.User_model) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	uid, err := a.GetIdByLogin(user.GetLogin())
	if err != nil {
		return "error in generating token", err
	}

	//Помещение в токен информации о пользователи и времени пока токен будет валидным
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = uid
	claims["ulogin"] = user.GetLogin()
	claims["nickname"] = user.GetNickname()
	claims["exp"] = time.Now().Add(a.Duration).Unix()

	tokenString, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return "error in generating token", err
	}

	return tokenString, nil
}

func (a *Auth) GetIdByLogin(login string) (int, error) {
	stmt := "SELECT userid FROM users WHERE login = $1"
	var uid int

	row := a.DB.QueryRow(stmt, login)
	if row.Scan(&uid) == sql.ErrNoRows {
		return 0, errors.New("wrong login")
	}
	return uid, nil
}
