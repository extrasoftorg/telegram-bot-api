package telegrambotapi

import "errors"

type User struct {
	ID        int64   `json:"id"`
	IsBot     bool    `json:"is_bot"`
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
	Username  *string `json:"username"`
}

type ChatType string

func (t ChatType) String() string {
	return string(t)
}

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSuperGroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

type Chat struct {
	ID    int64    `json:"id"`
	Type  ChatType `json:"type"`
	Title *string  `json:"title"`
}

type Message struct {
	ID   int64   `json:"message_id"`
	From *User   `json:"from"`
	Chat Chat    `json:"chat"`
	Text *string `json:"text"`
}

var ErrUnexpectedStatusCode = errors.New("unexpected status code")
