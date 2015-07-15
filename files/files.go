package files

import (
	"mime/multipart"
	"bytes"
	"io"
)

type File struct {
	Name string
	Body string
	Hash string
}

type MultiFile struct {
	Body *bytes.Buffer
	Writer *multipart.Writer
}

func NewMultiFile () *MultiFile {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	return &MultiFile{
		body, writer,
	}
}

func (m *MultiFile) AddFile(name string, reader io.Reader) error {
	part, err := m.Writer.CreateFormFile("file", name)

	if err != nil {
		return err
	}

	_, err = io.Copy(part, reader)

	return err
}
