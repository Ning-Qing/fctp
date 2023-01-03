package file

import (
	"io"
	"os"
)

type FileHeader interface {
	Open() *os.File
	Size() int64
	io.Writer
}

type fileHeader struct {
	file *os.File
	size int64
}

func NewFileHeader() (FileHeader, error) {
	file, err := os.CreateTemp(os.TempDir(), "fctp-")
	if err != nil {
		return nil, err
	}
	return &fileHeader{
		file: file,
		size: 0,
	}, nil
}

func (h *fileHeader) Open() *os.File {
	return h.file
}

func (h *fileHeader) Size() int64 {
	return h.size
}

func (h *fileHeader) Write(p []byte) (n int, err error) {
	n, err = h.file.Write(p)
	h.size += int64(n)
	return n, err
}
