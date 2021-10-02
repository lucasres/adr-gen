package file

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
)

type Walker interface {
	Walk(ctx context.Context, dirPath string) error
	Out() <-chan string
}

type LocalWalk struct {
	ignorePatterns []*regexp.Regexp
	out            chan string
}

func NewLocalWalk(bufferSize int, ignorePatterns []string) (*LocalWalk, error) {
	w := &LocalWalk{
		out: make(chan string, bufferSize),
	}

	for i := range ignorePatterns {
		if reg, err := regexp.Compile(ignorePatterns[i]); err != nil {
			return nil, fmt.Errorf("can't create regex pattern \"%s\" to ignore files: %w", ignorePatterns[i], err)
		} else {
			w.ignorePatterns = append(w.ignorePatterns, reg)
		}
	}

	return w, nil
}

func (l *LocalWalk) Walk(ctx context.Context, dirPath string) error {
	defer close(l.out)

	return filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("can't walk on specified dir \"%s\": %w", dirPath, err)
		}

		if d.IsDir() {
			return nil
		}

		for i := range l.ignorePatterns {
			if l.ignorePatterns[i].MatchString(path) {
				return nil
			}
		}

		fmt.Println(path, d, err)

		l.out <- path
		return nil
	})
}

func (l *LocalWalk) Out() <-chan string {
	return l.out
}
