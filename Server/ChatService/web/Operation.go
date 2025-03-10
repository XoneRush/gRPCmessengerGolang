package web

import (
	"database/sql"
	"errors"

	"github.com/XoneRush/gRPCmessengerGolang/Server/ChatService/model"
	pb "github.com/XoneRush/gRPCmessengerGolang/Server/ChatService/protos"
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
	stmt := "DELETE FROM members WHERE userid = $1 AND chatid = $2"

	err := a.checkExistance(m)
	if err != nil {
		return err
	}

	_, err = a.DB.Exec(stmt, m.UserID, m.ChatID)
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

func (a *App) GetNicknameFromDB(id int) (string, error) {
	stmt := "SELECT nickname FROM users WHERE userid = $1"

	var nickname string
	row := a.DB.QueryRow(stmt, id)
	err := row.Scan(&nickname)
	if err == sql.ErrNoRows {
		a.log.Error(err.Error())
		return "", errors.New("user does not exist")
	}
	if err != nil {
		return "", err
	}
	return nickname, nil
}

func (a *App) isMemberInChat(uid int, chatid int) error {
	stmt := "SELECT chatid FROM members WHERE userid = $1 AND chatid = $2"

	var c int
	row := a.DB.QueryRow(stmt, uid, chatid)
	err := row.Scan(&c)
	if err == sql.ErrNoRows {
		a.log.Error(err.Error())
		return errors.New("User cant send msg to this chat")
	}
	if err != nil {
		return err
	}
	return nil
}

// Получить все чаты, приуореченные к конкретному участнику
func (a *App) GetChatsFromDB(memberID int32) ([]pb.Chat, error) {
	stmt := "SELECT c.chatid, c.chat_name FROM chats as c JOIN members as m ON c.chatid = m.chatid WHERE m.userid = $1"
	var chats []pb.Chat
	var tmp_name string
	var tmp_id int

	rows, err := a.DB.Query(stmt, memberID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&tmp_id, &tmp_name)
		if err != nil {
			return nil, err
		}
		chat := pb.Chat{
			Name:   tmp_name,
			ChatID: int32(tmp_id),
		}

		chats = append(chats, chat)
	}

	return chats, nil
}
