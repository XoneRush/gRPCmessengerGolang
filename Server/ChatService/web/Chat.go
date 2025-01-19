package web

import (
	"context"
	"database/sql"
	"io"
	"log"
	"sync"

	"github.com/XoneRush/gRPCmessengerGolang/Server/ChatService/model"
	pb "github.com/XoneRush/gRPCmessengerGolang/Server/ChatService/protos"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

type App struct {
	pb.UnimplementedChatServiceServer
	log     hclog.Logger
	DB      *sql.DB
	mu      sync.Mutex
	clients map[int32]map[int32]pb.ChatService_SendMessageServer
}

func NewApp(l hclog.Logger, db *sql.DB) *App {
	return &App{log: l, DB: db, clients: make(map[int32]map[int32]grpc.BidiStreamingServer[pb.Msg, pb.Msg])}
}

// Создание чата
//
// На вход: структура чата
//
// Примечание: Получает данные каждого участника чата (из параметра), затем конвертирует *pb.Member -> model.Member,
// после чего помещает чат в базу данных
func (a *App) CreateChat(ctx context.Context, chat *pb.Chat) (*pb.Void, error) {
	chatName := chat.GetName()
	members := chat.GetListOfMembers()
	var chatMembers []model.Member_model

	for _, mmbr := range members {
		m := ConvertPBtoModel(mmbr)
		chatMembers = append(chatMembers, m)
	}
	err := a.CreateChatInDB(chatName, chatMembers)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (a *App) SendMessage(stream pb.ChatService_SendMessageServer) error {
	//Получение ID пользователя и чата
	msg, err := stream.Recv()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	userID := msg.GetSrc()
	chatID := msg.GetDst()

	//Создание мапы следующей структуры:
	//chatID : {
	// UserID1 : stream
	// UserID2 : stream
	// }
	// Т.е. мапа будет содержать активное подключение каждого клиента
	a.mu.Lock()
	if _, ok := a.clients[chatID]; !ok {
		a.clients[chatID] = make(map[int32]grpc.BidiStreamingServer[pb.Msg, pb.Msg])
	}
	a.clients[chatID][userID] = stream
	a.mu.Unlock()

	//Обработка входящих сообщений
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			// Удаление клиента из чата при ошибке
			a.mu.Lock()
			delete(a.clients[chatID], userID)
			a.mu.Unlock()
			return err
		}
		err = a.isMemberInChat(int(msg.Src), int(msg.Dst))
		if err != nil {
			a.log.Info("Error!", "err", err.Error())
			continue
		}

		a.log.Info("Message!", "data", msg.Data, "from", msg.Dst)

		//Пересылка сообщений другим клиентам
		a.mu.Lock()
		for _, clientStream := range a.clients[chatID] {
			nickname, err := a.GetNicknameFromDB(int(msg.Src))
			if err != nil {
				log.Printf("Error sending message to client: %v", err)
			}

			str := nickname + ": " + msg.Data
			msg.Data = str

			if err := clientStream.Send(msg); err != nil {
				log.Printf("Error sending message to client: %v", err)
			}
		}
		a.mu.Unlock()
	}
}

// Добавляет участника в чат (указанный в поле chatid структуры)
func (a *App) AddMember(ctx context.Context, member *pb.Member) (*pb.Msg, error) {
	m := ConvertPBtoModel(member)

	err := a.AddMemberToDB(m)
	if err != nil {
		a.log.Error(err.Error())
		return nil, err
	}

	a.log.Info("Added member", "chat", member.GetChatID(), "member", member.GetUserID())
	return &pb.Msg{Data: "Succses! Added member", Dst: member.GetChatID()}, nil
}

// Удаляет участника из чата
func (a *App) RemoveMember(ctx context.Context, member *pb.Member) (*pb.Msg, error) {
	m := ConvertPBtoModel(member)

	err := a.RemoveFromDB(m)
	if err != nil {
		a.log.Error(err.Error())
		return &pb.Msg{Data: "Error!", Dst: member.GetChatID()}, err
	}

	a.log.Info("Removed member", "chat", member.GetChatID(), "member", member.GetUserID())
	return &pb.Msg{Data: "Succses! Removed member ", Dst: member.GetChatID()}, nil
}

// Необязательно но желательно
// Stream, возвращают множество объектов в поток
func (a *App) ListMembers(chat *pb.Chat, listMembers grpc.ServerStreamingServer[pb.Member]) error {
	return nil
}

func (a *App) ListMsgs(chat *pb.Chat, listMsgs grpc.ServerStreamingServer[pb.Msg]) error {

	return nil
}
