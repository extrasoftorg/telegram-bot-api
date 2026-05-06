package telegrambotapi

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type SendMessageRequest struct {
	ChatID      int64                 `json:"chat_id"`
	Text        string                `json:"text"`
	ParseMode   string                `json:"parse_mode,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (c *client) SendMessage(ctx context.Context, req SendMessageRequest) (*Message, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	msg, err := makeRequest[Message](ctx, http.MethodPost, "/sendMessage", bytes.NewReader(body), makeRequestOptions{
		baseURL:    c.baseURL,
		httpClient: c.httpClient,
		token:      c.token,
	})
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (c *client) AnswerCallbackQuery(ctx context.Context, req AnswerCallbackQueryRequest) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	_, err = makeRequest[bool](ctx, http.MethodPost, "/answerCallbackQuery", bytes.NewReader(body), makeRequestOptions{
		baseURL:    c.baseURL,
		httpClient: c.httpClient,
		token:      c.token,
	})
	return err
}
