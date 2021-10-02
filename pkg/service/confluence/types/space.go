package types

// Space representes Confluence space json object
// https://docs.atlassian.com/ConfluenceServer/rest/7.11.6/#api/content-getContent
type Space struct {
	ID   *int    `json:"id,omitempty"`
	Key  string  `json:"key"`
	Name *string `json:"name,omitempty"`
}
