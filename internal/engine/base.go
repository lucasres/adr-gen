package engine

type ADRHit struct {
	Description string
	ID          string
}

type ContentParsed struct {
	RawContent  string
	ByteContent []byte
}

type Engine interface {
	PreProcess(content []byte) ([]byte, error)
	Parse([]byte) (*ContentParsed, error)
	Analize(content ContentParsed) (map[string]ADRHit, error)
	Run(content []byte) (map[string]ADRHit, error)
}
