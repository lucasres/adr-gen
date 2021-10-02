package file

import (
	"context"
	"fmt"
	"io"
	"os"
)

type Reader interface {
	Read(ctx context.Context, walk Walker) error
	Out() <-chan []byte
}

type LocalReader struct {
	out chan []byte
}

func NewLocalReader(bufferSize int) *LocalReader {
	return &LocalReader{
		out: make(chan []byte, bufferSize),
	}
}

func (l *LocalReader) Read(ctx context.Context, walk Walker) error {
	defer close(l.out)

	for filePath := range walk.Out() {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("can't read file \"%s\": %w", filePath, err)
		}

		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			return fmt.Errorf("can't read all content of file \"%s\": %w", filePath, err)
		}

		l.out <- content
	}

	return nil
}

func (l *LocalReader) Out() <-chan []byte {
	return l.out
}
