package micro

import (
	"bytes"

	"github.com/Ning-Qing/fctp"
	"github.com/Ning-Qing/fctp/pb"
	"github.com/Ning-Qing/fctp/stream/grpc"
)

type transfer struct {
	buf    []byte
	buffer *bytes.Buffer
	stream grpc.Stream
}

func NewTransfer(s grpc.Stream) fctp.Transfer {
	buf := make([]byte, 0, 1<<10)
	return &transfer{
		buf:    buf,
		buffer: bytes.NewBuffer(buf),
		stream: s,
	}
}

func (t *transfer) Write(p []byte) (n int, err error) {
	n, err = t.buffer.Write(p)
	if err != nil {
		return n, err
	}
	err = t.stream.SendMsg(&pb.Message{
		Data: t.buffer.Bytes(),
	})
	if err != nil {
		return 0, err
	}
	t.buffer.Reset()
	return n, nil
}

func (t *transfer) Read(p []byte) (n int, err error) {
	msg := new(pb.Message)
	err = t.stream.RecvMsg(msg)
	if err != nil {
		return 0, err
	}
	n = copy(p, msg.Data)
	return n, nil
}
