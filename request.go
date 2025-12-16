package telegrambotapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type makeRequestOptions struct {
	baseURL    string
	httpClient *http.Client
}

func makeRequest[T any](
	ctx context.Context,
	token string,
	method string,
	body io.Reader,
	opts makeRequestOptions,
) (*T, error) {
	fullURL := fmt.Sprintf("%s%s%s", opts.baseURL, token, method)
	req, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		return nil, err
	}

	resp, err := opts.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result T
	if err = json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
