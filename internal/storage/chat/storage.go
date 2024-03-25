package user

import (
	"context"
	"fmt"

	"log"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/Shemistan/chat_server/internal/model"
	def "github.com/Shemistan/chat_server/internal/storage"
)

type storage struct {
	db *pgxpool.Pool
}

// NewStorage - новый storage
func NewStorage(db *pgxpool.Pool) def.Chat {
	return &storage{db: db}
}

// CreateChat - создать чата
func (s *storage) CreateChat(ctx context.Context, req model.Chat) (int64, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx) //nolint:errcheck

	chatID, err := s.createChat(ctx, tx, req.Name)
	if err != nil {
		return 0, err
	}

	err = s.addUsers(ctx, tx, req.Users)
	if err != nil {
		return 0, err
	}

	err = s.addChatUserList(ctx, tx, chatID, req.Users)
	if err != nil {
		return 0, err
	}

	return chatID, tx.Commit(ctx)
}

func (s *storage) createChat(ctx context.Context, tx pgx.Tx, chatName string) (int64, error) {
	query := `INSERT INTO chat( name) VALUES ( $1) RETURNING(id);`

	var id int64
	err := tx.QueryRow(ctx, query, chatName).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *storage) addUsers(ctx context.Context, tx pgx.Tx, userLogins []string) error {
	query := `INSERT INTO users(login) SELECT login FROM unnest($1::varchar[]) AS login 
              WHERE login NOT IN (SELECT login FROM users) RETURNING id;`

	_, err := tx.Exec(ctx, query, userLogins)
	if err != nil {
		return err
	}

	return err
}

func (s *storage) addChatUserList(ctx context.Context, tx pgx.Tx, chatID int64, userLogins []string) error {
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

	_, err := tx.Exec(ctx, query, args...)
	return err
}

// AddMessage - создать сообщение
func (s *storage) AddMessage(ctx context.Context, req model.Message) error {
	query := `
INSERT INTO messages (chat_id, user_id, message)
select c.id, u.id, $1 from chat as c
                           left join chat_users_list cul on c.id = cul.chat_id
                           left join users u on cul.user_id = u.id
WHERE u.login=$2 AND c.name = $3;
`

	var id int64
	err := s.db.QueryRow(ctx, query, req.Message, req.UserLogin, req.ChatName).Scan(&id)
	if err != nil {
		return err
	}

	log.Printf("created message(%d) for user(%s) in to chat(%s)", id, req.UserLogin, req.ChatName)

	return nil
}

// DeactivateChat - деактивировать чат
func (s *storage) DeactivateChat(ctx context.Context, chatID int64) error {
	query := `UPDATE chat SET is_active=false WHERE chat.id = $1`

	_, err := s.db.Exec(ctx, query, chatID)
	if err != nil {
		return err
	}
	return nil
}
