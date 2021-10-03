package engine

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Sengine struct{}

func (e *Sengine) Analize(c ContentParsed) (map[string]ADRHit, error) {
	result := map[string]ADRHit{}

	hitRegex := regexp.MustCompile(`((\/\/|#)\s@ADR-\d):.*\n`)
	hits := hitRegex.FindAll(c.ByteContent, -1)

	for _, hit := range hits {
		// limpa o // @
		regexClear := regexp.MustCompile(`// @`)
		hit = regexClear.ReplaceAll(hit, []byte(""))
		//@todo: pode ser melhorado pegando por grupos direto do regex
		// ADR-x : Lorem ipsun
		splited := strings.Split(string(hit), ":")
		result[splited[0]] = ADRHit{
			ID:          splited[0],
			Description: splited[1],
		}
	}

	return result, nil
}

func (e *Sengine) Parse(content []byte) (*ContentParsed, error) {
	return &ContentParsed{
		RawContent:  string(content),
		ByteContent: content,
	}, nil
}

func (e *Sengine) PreProcess(content []byte) ([]byte, error) {
	// @todo precisa ser refatorado, estava bugando o codigo
	return content, nil
}

func (e *Sengine) Output(hits map[string]ADRHit) error {
	return nil
}

func (e *Sengine) Run(content []byte) error {
	processed, err := e.PreProcess(content)

	if err != nil {
		return err
	}

	parsed, err := e.Parse(processed)

	if err != nil {
		return err
	}

	hits, err := e.Analize(*parsed)

	if err != nil {
		return err
	}

	for a, b := range hits {
		fmt.Println(a)
		fmt.Println(b.Description, b.ID)
	}

	return errors.New("must be not implemented")
}

func NewSengine() *Sengine {
	return &Sengine{}
}
