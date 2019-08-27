package grpc

import (
	"errors"
	grpcpb "fgame/fgame/core/session/grpc/pb"
)

var (
	ErrNotSupported    = errors.New("not supported")
	ErrPayloadTooLarge = errors.New("payload too large")
)

func serverReceive(conn grpcpb.Connection_ConnectServer, v interface{}) (err error) {
	msg, err := conn.Recv()
	if err != nil {
		return
	}
	switch data := v.(type) {
	case *[]byte:
		*data = msg.Body
		return nil
	}
	return ErrNotSupported
}

func clientReceive(conn grpcpb.Connection_ConnectClient, v interface{}) (err error) {
	msg, err := conn.Recv()
	if err != nil {
		return
	}
	switch data := v.(type) {
	case *[]byte:
		*data = msg.Body
		return nil
	}
	return ErrNotSupported
}

type Handler interface {
	Handle(conn *grpcpb.Connection_ConnectServer)
}

type HandlerFunc func(*grpcpb.Connection_ConnectServer)

func (hf HandlerFunc) Handle(conn *grpcpb.Connection_ConnectServer) {
	hf(conn)
}
