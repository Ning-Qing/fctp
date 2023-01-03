package fctp

import "io"

type Transfer interface {
	io.Writer
	io.Reader
}


