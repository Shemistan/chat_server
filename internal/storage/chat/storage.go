package chat

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Shemistan/chat_server/internal/client/db"
	"github.com/Shemistan/chat_server/internal/model"
	def "github.com/Shemistan/chat_server/internal/storage"
)

type storage struct {
	db        db.Client
	txManager db.TxManager
}

// NewStorage - новый storage
func NewStorage(db db.Client, txManager db.TxManager) def.Chat {
	return &storage{
		db:        db,
		txManager: txManager,
	}
}

// CreateChat - создать чата
func (s *storage) CreateChat(ctx context.Context, req model.Chat) (int64, error) {
	var chatID int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		chatID, errTx = s.createChat(ctx, req.Name)
		if errTx != nil {
			return errTx
		}

		errTx = s.addUsers(ctx, req.Users)
		if errTx != nil {
			return errTx
		}

		errTx = s.addChatUserList(ctx, chatID, req.Users)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return chatID, nil
}

func (s *storage) createChat(ctx context.Context, chatName string) (int64, error) {
	query := `INSERT INTO chat_v1( name) VALUES ( $1) RETURNING(id);`

	var id int64
	err := s.db.DB().QueryRowContext(ctx, db.Query{
		Name:     "create_chat",
		QueryRaw: query,
	}, chatName).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *storage) addUsers(ctx context.Context, userLogins []string) error {
	query := `INSERT INTO users(login) SELECT login FROM unnest($1::varchar[]) AS login 
              WHERE login NOT IN (SELECT login FROM users) RETURNING id;`

	_, err := s.db.DB().ExecContext(ctx, db.Query{
		Name:     "add_user",
		QueryRaw: query,
	}, userLogins)
	if err != nil {
		return err
	}

	return err
}

func (s *storage) addChatUserList(ctx context.Context, chatID int64, userLogins []string) error {
	query := `
        INSERT INTO chat_users_list (chat_id, user_id)
        SELECT $1, id
        FROM users
        WHERE login IN (%s);
    `

	placeholders := make([]string, len(userLogins))
	for i := range userLogins {
		placeholders[i] = "$" + strconv.Itoa(i+2) // chatID использует плейсхолдер $1
	}

	placeholdersQuery := strings.Join(placeholders, ", ")

	query = fmt.Sprintf(query, placeholdersQuery)

	args := make([]interface{}, len(userLogins)+1)
	args[0] = chatID
	for i, login := range userLogins {
		args[i+1] = login
	}

	_, err := s.db.DB().ExecContext(ctx, db.Query{
		Name:     "add_chat_user_list",
		QueryRaw: query,
	}, args...)
	return err
}

// AddMessage - создать сообщение
func (s *storage) AddMessage(ctx context.Context, req model.Message) error {
	query := `
INSERT INTO messages (chat_id, user_id, message)
select c.id, u.id, $1 from chat_v1 as c
                           left join chat_users_list cul on c.id = cul.chat_id
                           left join users u on cul.user_id = u.id
WHERE u.login=$2 AND c.name = $3;
`

	var id int64
	err := s.db.DB().QueryRowContext(ctx, db.Query{
		Name:     "add_message",
		QueryRaw: query,
	}, req.Message, req.UserLogin, req.ChatName).Scan(&id)
	if err != nil {
		return err
	}

	log.Printf("created message(%d) for user(%s) in to chat_v1(%s)", id, req.UserLogin, req.ChatName)

	return nil
}

// DeactivateChat - деактивировать чат
func (s *storage) DeactivateChat(ctx context.Context, chatID int64) error {
	query := `UPDATE chat_v1 SET is_active=false WHERE chat_v1.id = $1`

	_, err := s.db.DB().ExecContext(ctx, db.Query{
		Name:     "deactivate_chat",
		QueryRaw: query,
	}, chatID)
	if err != nil {
		return err
	}
	return nil
}
