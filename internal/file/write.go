package file

import "io/ioutil"

type Writer interface {
	Write(path, content string) error
}

type LocalWriter struct{}

func (w *LocalWriter) Write(path, content string) error {
	err := ioutil.WriteFile(path, []byte(content), 0644)

	if err != nil {
		return err
	}

	return nil
}

func NewLocalWrite() *LocalWriter {
	return &LocalWriter{}
}
