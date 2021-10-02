package confluence

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	// pat is acronym for Personal Access Token
	authorization string
	urlPrefix     string
	httpClient    *http.Client
}

// NewClient create a new Confluence HTTP API client
// pat is acronym for Personal Access Token: https://confluence.atlassian.com/enterprise/using-personal-access-tokens-1026032365.html
// urlPrefix is the base URL of your instance of Confluence Server
func NewClient(pat, urlPrefix string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		authorization: fmt.Sprintf("Bearer: %s", pat),
		urlPrefix:     urlPrefix,
		httpClient:    httpClient,
	}
}

// doRequest Create and send HTTP request
// endpoint is concatened with Client.urlPrefix: "{Client.urlPrefix}/{endpoint}". So do not add "/" prefix do endpoint
func (c *Client) doRequest(ctx context.Context, method, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		method,
		fmt.Sprintf("%s/%s", c.urlPrefix, endpoint),
		body,
	)

	if err != nil {
		return nil, fmt.Errorf("can't create new request in confluence client: %w", err)
	}

	// https://confluence.atlassian.com/enterprise/using-personal-access-tokens-1026032365.html#UsingPersonalAccessTokens-UsingPATs
	req.Header.Add("Authorization", c.authorization)

	return c.httpClient.Do(req)
}
