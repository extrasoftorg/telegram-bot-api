package telegrambotapi

import "context"

const baseURL = "https://api.telegram.org/bot"

type Client interface {
	SendMessage(ctx context.Context, req SendMessageRequest) (*Message, error)
}
