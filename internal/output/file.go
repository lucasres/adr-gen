package output

import (
	"github.com/lucasres/adr-gen/internal/engine"
	"github.com/lucasres/adr-gen/internal/file"
)

type FileOutput struct{}

func (o *FileOutput) Write(path string, hits map[string]engine.ADRHit) error {
	w := file.NewLocalWrite()

	for _, hit := range hits {
		err := w.Write(path+hit.ID, hit.Description)

		if err != nil {
			return err
		}
	}

	return nil
}

func NewFileOutput() *FileOutput {
	return &FileOutput{}
}
