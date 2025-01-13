package web

import (
	"ChatService/model"
	pb "ChatService/protos"
	"database/sql"
	"errors"
)

func (a *App) CreateChatInDB(name string, members []model.Member_model) error {
	//Создать в базе данных чат
	stmt := "INSERT INTO chats(chat_name) VALUES($1)"

	_, err := a.DB.Exec(stmt, name)
	if err != nil {
		a.log.Error(err.Error())
		return err
	}

	//Добавить каждого участника в базу данных
	//Осуществляется провека на наличие пользователя
	stmt = "SELECT login FROM users WHERE userid = $1"
	stmt1 := "INSERT INTO members(userid, chatid) VALUES($1, $2)"
	for _, m := range members {
		//проверка на существование пользователя,
		//если его не существует, то такой пользователь просто не добавляется в базу данных \ чат
		var login string
		row := a.DB.QueryRow(stmt, m.UserID)
		err := row.Scan(&login)
		if err == sql.ErrNoRows {
			a.log.Error(err.Error())
			continue
		}
		if err != nil {
			return err
		}

		//Добавление мембера в базу данных
		_, err = a.DB.Exec(stmt1, m.UserID, m.ChatID)
		if err != nil {
			return err
		}
	}

	//При успешном добавлении в базу данных:
	return nil
}

// Получить список всех пользователей чата
func (a *App) GetChatMembers(chatID int) ([]int, error) {
	stmt := "SELECT * FROM members WHERE chatid = $1"

	rows, err := a.DB.Query(stmt, chatID)
	if err != nil {
		return nil, err
	}

	var ids []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}
	return ids, nil
}

// Добавляет участника в базу данных
func (a *App) AddMemberToDB(m model.Member_model) error {
	stmt := "INSERT INTO members(chatid, userid) VALUES($1,$2)"

	err := a.checkExistance(m)
	if err != nil {
		return err
	}

	_, err = a.DB.Exec(stmt, m.ChatID, m.UserID)
	if err != nil {
		return err
	}

	return nil
}

// Удаляет участника из бд
func (a *App) RemoveFromDB(m model.Member_model) error {
	stmt := "DELETE FROM members WHERE userid = $1"

	err := a.checkExistance(m)
	if err != nil {
		return err
	}

	_, err = a.DB.Exec(stmt, m.UserID)
	if err != nil {
		return err
	}
	return nil
}

func ConvertPBtoModel(m *pb.Member) model.Member_model {
	mbr := model.Member_model{
		UserID: m.GetUserID(),
		ChatID: m.GetChatID(),
		Role:   m.GetRole(),
	}
	return mbr
}

// Проверка существует ли пользователь
func (a *App) checkExistance(m model.Member_model) error {
	stmt := "SELECT login FROM users WHERE userid = $1"

	//Делается один запрос к БД, если по id ничего не возвращается, делается вывод, что юзера нет
	var login string
	row := a.DB.QueryRow(stmt, m.UserID)
	err := row.Scan(&login)
	if err == sql.ErrNoRows {
		a.log.Error(err.Error())
		return errors.New("user does not exist")
	}
	if err != nil {
		return err
	}
	return nil
}
