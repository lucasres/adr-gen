package engine

import (
	"errors"
	"regexp"
)

type Sengine struct{}

func (e *Sengine) Analize(c ContentParsed) (map[string]ADRHit, error) {

	return nil, nil
}

func (e *Sengine) Parse(content []byte) (*ContentParsed, error) {
	return &ContentParsed{
		RawContent:  string(content),
		ByteContent: content,
	}, nil
}

func (e *Sengine) PreProcess(content []byte) ([]byte, error) {
	// remove end line for ;
	regexEndLine := regexp.MustCompile(`\n`)
	contentWithSemicolon := regexEndLine.ReplaceAll([]byte(content), []byte(";"))
	// remove 2 spaces or more
	regexTwoSpaces := regexp.MustCompile(`\s{2,}`)
	contentWithSingleSpaces := regexTwoSpaces.ReplaceAll(contentWithSemicolon, []byte(""))

	return contentWithSingleSpaces, nil
}

func (e *Sengine) Output(hits map[string]ADRHit) error {
	return nil
}

func (e *Sengine) Run(content []byte) error {
	processed, err := e.PreProcess(content)

	if err != nil {
		return err
	}

	_, err = e.Parse(processed)

	if err != nil {
		return err
	}

	return errors.New("must be not implemented")
}

func NewSengine() *Sengine {
	return &Sengine{}
}
