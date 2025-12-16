package telegrambotapi

import "net/http"

type client struct {
	baseURL    string
	httpClient *http.Client

	token string
}

type NewClientOptions struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

func New(opts NewClientOptions) Client {
	if opts.BaseURL == "" {
		opts.BaseURL = baseURL
	}
	if opts.HTTPClient == nil {
		opts.HTTPClient = &http.Client{}
	}
	return &client{
		baseURL:    opts.BaseURL,
		httpClient: opts.HTTPClient,
		token:      opts.Token,
	}
}
