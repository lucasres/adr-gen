package confluence

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lucasres/adr-gen/pkg/service/confluence/types"
)

// CreatePagePayload define request struct
// https://docs.atlassian.com/ConfluenceServer/rest/7.11.6/#api/content-createContent
type CreatePagePayload struct {
	Type  string      `json:"type"`
	Title string      `json:"title"`
	Space types.Space `json:"space"`
	Body  types.Body  `json:"body"`
}

// CreatePageResult define request response struct
// https://developer.atlassian.com/server/confluence/confluence-rest-api-examples/#create-a-new-page
type CreatePageResult struct {
	CreatePagePayload
	types.IdentifiableEntity
}

func (c *Client) CreatePage(ctx context.Context, page *CreatePagePayload) (*CreatePageResult, error) {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(page); err != nil {
		return nil, fmt.Errorf("can't encode CreatePgaePayload to bytes buffer: %w", err)
	}

	res, err := c.doRequest(ctx, http.MethodPost, "content", body)
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

	var result *CreatePageResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("can't decode create page response: %w", err)
	}

	return result, nil
}
