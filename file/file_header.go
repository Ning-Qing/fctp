package file

import (
	"io"
	"os"
)

type FileHeader interface {
	Open() *os.File
	Size() int64
	Name() string
	io.Writer
}

type fileHeader struct {
	file *os.File
	size int64
	name string
}

func NewFileHeader(name string) (FileHeader, error) {
	file, err := os.CreateTemp(os.TempDir(), "fctp-")
	if err != nil {
		return nil, err
	}
	return &fileHeader{
		file: file,
		size: 0,
		name: name,
	}, nil
}

func (h *fileHeader) Open() *os.File {
	return h.file
}

func (h *fileHeader) Size() int64 {
	return h.size
}

func (h *fileHeader) Name() string {
	return h.name
}

func (h *fileHeader) Write(p []byte) (n int, err error) {
	n, err = h.file.Write(p)
	h.size += int64(n)
	return n, err
}
