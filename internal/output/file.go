package output

import (
	"bytes"
	"html/template"
	"path/filepath"

	"github.com/lucasres/adr-gen/internal/engine"
	"github.com/lucasres/adr-gen/internal/file"
)

type FileOutput struct{}

func (o *FileOutput) BuildOutput(hit engine.ADRHit) (string, error) {
	templateContent := `
# {{.ID}}: {{.Description}}

## Status

What is the status, such as proposed, accepted, rejected, deprecated, superseded, etc.?

## Context

What is the issue that we're seeing that is motivating this decision or change?

## Decision

What is the change that we're proposing and/or doing?

## Consequences

What becomes easier or more difficult to do because of this change?	
	`

	tmplt := template.Must(template.New("template").Parse(templateContent))

	var out bytes.Buffer

	tmplt.Execute(&out, hit)

	return out.String(), nil
}

func (o *FileOutput) Write(path string, hits map[string]engine.ADRHit) error {
	w := file.NewLocalWrite()

	for _, hit := range hits {
		content, err := o.BuildOutput(hit)

		if err != nil {
			return err
		}

		err = w.Write(filepath.Join(path, hit.ID+".md"), content)

		if err != nil {
			return err
		}
	}

	return nil
}

func NewFileOutput() *FileOutput {
	return &FileOutput{}
}
