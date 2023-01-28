package http

import (
	"io"
)

type transfer struct {
	stream io.ReadCloser
}

// func NewTransfer(r *stdhttp.Request, w *stdhttp.Response) fctp.Transfer {
// 	return &transfer{
// 		stream: stream,
// 	}
// }
