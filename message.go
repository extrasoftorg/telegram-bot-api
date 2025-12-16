package telegrambotapi

import (
	"bytes"
	"context"
	"encoding/json"
)

type SendMessageRequest struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func (c *client) SendMessage(ctx context.Context, req SendMessageRequest) (*Message, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	msg, err := makeRequest[Message](ctx, c.token, "/sendMessage", bytes.NewReader(body), makeRequestOptions{
		baseURL:    c.baseURL,
		httpClient: c.httpClient,
	})
	if err != nil {
		return nil, err
	}
	return msg, nil
}
