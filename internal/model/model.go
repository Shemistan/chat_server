package model

// Chat - модель чата
type Chat struct {
	ID    int64    `db:"id"`
	Name  string   `db:"name"`
	Users []string `db:"users"`
}

// Message - модель сообщения
type Message struct {
	ChatName  string
	UserLogin string
	Message   string
}
