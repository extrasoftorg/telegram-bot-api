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
	token      string
}

type apiResponse[T any] struct {
	OK          bool   `json:"ok"`
	Description string `json:"description"`
	Result      *T     `json:"result"`
}

func makeRequest[T any](
	ctx context.Context,
	method string,
	path string,
	body io.Reader,
	opts makeRequestOptions,
) (*T, error) {
	fullURL := fmt.Sprintf("%s%s%s", opts.baseURL, opts.token, path)
	req, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := opts.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result apiResponse[T]
	if err = json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, err
	}

	if !result.OK {
		return nil, fmt.Errorf("telegram api error: %s", result.Description)
	}

	return result.Result, nil
}
