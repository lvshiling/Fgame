package protobuf_util

import (
	"errors"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/websocket"
)

var (
	TypeError = errors.New("data is not proto message type")
)

func protobufMarshal(v interface{}) (msg []byte, payloadType byte, err error) {
	message, exist := v.(proto.Message)
	if !exist {
		err = TypeError
		return
	}
	msg, err = proto.Marshal(message)
	return msg, websocket.BinaryFrame, err
}

func protobufUnmarshal(msg []byte, payloadType byte, v interface{}) (err error) {
	message, exist := v.(proto.Message)
	if !exist {
		err = TypeError
		return
	}
	err = proto.Unmarshal(msg, message)
	return err
}

var Protobuf = &websocket.Codec{
	Unmarshal: protobufUnmarshal,
	Marshal:   protobufMarshal,
}
