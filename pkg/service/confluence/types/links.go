package types

// Links given links for current and related objects
// https://developer.atlassian.com/server/confluence/confluence-rest-api-examples/
type Links struct {
	Self *string `json:"self,omitempty"`
}
