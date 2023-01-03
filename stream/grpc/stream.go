package grpc

type Stream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
}
