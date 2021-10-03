package confluence

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lucasres/adr-gen/pkg/service/confluence/types"
)

// CreatePageInput define request struct
// https://docs.atlassian.com/ConfluenceServer/rest/7.11.6/#api/content-createContent
type CreatePageInput struct {
	Type  string      `json:"type"`
	Title string      `json:"title"`
	Space types.Space `json:"space"`
	Body  types.Body  `json:"body"`
}

// PageResult define request response struct
// https://developer.atlassian.com/server/confluence/confluence-rest-api-examples/#create-a-new-page
type PageResult struct {
	ID int `json:"id"`
	CreatePageInput
	types.IdentifiableEntity
}

func (c *Client) CreatePage(ctx context.Context, page *CreatePageInput) (*PageResult, error) {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(page); err != nil {
		return nil, fmt.Errorf("can't encode create page payload to bytes buffer: %w", err)
	}

	req, err := c.createRequest(ctx, http.MethodPost, "content", body)
	if err != nil {
		return nil, fmt.Errorf("can't create create page request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't request create page: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf(
			"confluence don't return HTTP Status Code 201. Returned: [%d] - \"%s\"",
			res.StatusCode,
			res.Body,
		)
	}

	var result *PageResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("can't decode create page response: %w", err)
	}

	return result, nil
}
