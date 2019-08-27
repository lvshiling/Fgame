package codec

import (
	"fgame/fgame/common/codec"

	"github.com/golang/protobuf/proto"
)

var (
	c = codec.NewCodec()
)

func Encode(msg proto.Message) (msgBytes []byte, err error) {
	return c.Encode(msg)
}

func Decode(msg []byte) (m *codec.Message, err error) {
	return c.Decode(msg)
}

func RegisterMsg(msgType codec.MessageType, msg proto.Message) {
	c.Register(msgType, msg)
}

func GetCodec() *codec.Codec {
	return c
}
