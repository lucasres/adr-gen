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

type GetPagesInput struct {
	SpaceKey *string
	Title    *string
	Type     *string
	Status   *string
	Start    *uint
	Limit    *uint
}

type PageListResult struct {
	Limit int          `json:"limit"`
	Size  int          `json:"size"`
	Start int          `json:"start"`
	Pages []PageResult `json:"results"`
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

// GetPages search for pages in confluence
// https://docs.atlassian.com/ConfluenceServer/rest/7.11.6/#api/content-getContent
func (c *Client) GetPages(ctx context.Context, input *GetPagesInput) (*PageListResult, error) {
	req, err := c.createRequest(ctx, http.MethodGet, "content", nil)
	if err != nil {
		return nil, fmt.Errorf("can't create get pages request: %w", err)
	}

	q := req.URL.Query()

	if input.SpaceKey != nil {
		q.Add("spaceKey", *input.SpaceKey)
	}

	if input.Title != nil {
		q.Add("title", *input.Title)
	}

	if input.Type != nil {
		q.Add("type", *input.Type)
	}

	if input.Status != nil {
		q.Add("status", *input.Status)
	}

	if input.Limit != nil {
		q.Add("limit", fmt.Sprint(*input.Limit))
	}

	if input.Start != nil {
		q.Add("start", fmt.Sprint(*input.Start))
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't request get pages: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"confluence don't return HTTP Status Code 200. Returned: [%d] - \"%s\"",
			res.StatusCode,
			res.Body,
		)
	}

	var result *PageListResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("can't decode get pages response: %w", err)
	}

	return result, nil
}
