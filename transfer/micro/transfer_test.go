package micro_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/Ning-Qing/fctp/pb"
	"github.com/Ning-Qing/fctp/stream/grpc"
	"github.com/Ning-Qing/fctp/transfer/micro"
	"github.com/bmizerany/assert"
	"google.golang.org/protobuf/proto"
)

type stream struct {
	ch chan []byte
}

func NewStream() grpc.Stream {
	return &stream{
		ch: make(chan []byte, 1),
	}
}

func (s *stream) SendMsg(m interface{}) error {
	buf, err := proto.Marshal(m.(proto.Message))
	if err != nil {
		return err
	}
	s.ch <- buf
	return nil
}

func (s *stream) RecvMsg(m interface{}) error {
	return proto.Unmarshal(<-s.ch, m.(proto.Message))
}

func TestStream(t *testing.T) {
	var err error
	want := []byte("hello world")
	s := NewStream()
	err = s.SendMsg(&pb.Message{
		Data: want,
	})
	assert.Equal(t, err, nil)
	msg := new(pb.Message)
	err = s.RecvMsg(msg)
	assert.Equal(t, err, nil)
	assert.Equal(t, msg.Data, want)
}

func TestTransfer(t *testing.T) {
	s := NewStream()
	buffer := bytes.NewBuffer(make([]byte, 1024<<20))
	go RecvServer(s)
	transfer := micro.NewTransfer(s)
	n, err := io.CopyBuffer(transfer, buffer, make([]byte, 1<<20))
	if err != nil {
		t.Fail()
	}
	t.Log(n >> 20)
}

func BenchmarkTransfer(b *testing.B) {
	s := NewStream()
	buffer := bytes.NewBuffer(make([]byte, 0, 1<<20))
	go RecvServer(s)
	transfer := micro.NewTransfer(s)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := buffer.Write(make([]byte, 1<<20))
		if err != nil {
			b.Fail()
		}
		io.CopyBuffer(transfer, buffer, make([]byte, 1<<10))
	}
}

func RecvServer(s grpc.Stream) {
	transfer := micro.NewTransfer(s)
	buffer := bytes.NewBuffer(make([]byte, 1<<20))
	n, err := io.CopyBuffer(buffer, transfer, make([]byte, 1<<10))
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
