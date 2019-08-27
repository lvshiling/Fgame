package codec

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
)

type MessageType uint16

type Message struct {
	MessageType MessageType
	Body        interface{}
}

var (
	//消息长度太小
	ErrorMessageTooSmall = errors.New("error message too small")
	//查找不到结构
	ErrorMessageNoExist = errors.New("error message no exist")
	//查找不到类型
	ErrorMessageNoType = errors.New("error message no type")
)

type Codec struct {
	protoMap    map[MessageType]reflect.Type
	revProtoMap map[reflect.Type]MessageType
}

//TODO 优化buff池 参考:https://github.com/grpc/grpc-go/blob/master/encoding/proto/proto.go
func (c *Codec) Encode(msg proto.Message) (msgBytes []byte, err error) {
	msgType, exist := c.getMsgType(msg)
	if !exist {
		err = fmt.Errorf("encode: msg no type %#v", msg)
		return
	}

	var outputBuffer bytes.Buffer
	msgTypeInt := uint16(msgType)
	if err = binary.Write(&outputBuffer, binary.LittleEndian, msgTypeInt); err != nil {
		return
	}
	bodyBytes, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	bodyBytesLen := uint16(len(bodyBytes))
	if err = binary.Write(&outputBuffer, binary.LittleEndian, bodyBytesLen); err != nil {
		return
	}

	if err = binary.Write(&outputBuffer, binary.LittleEndian, bodyBytes); err != nil {
		return
	}
	msgBytes = outputBuffer.Bytes()
	return
}

func (c *Codec) Decode(msg []byte) (m *Message, err error) {
	//TODO
	if len(msg) < 4 {
		err = ErrorMessageTooSmall
		return
	}

	msgTypeInt := binary.LittleEndian.Uint16(msg[:2])
	msgType := MessageType(msgTypeInt)
	body, ok := c.getMessage(msgType)
	if !ok {
		err = fmt.Errorf("msg [%d] no exist", msgType)
		return
	}
	//TODO
	err = proto.Unmarshal(msg[4:], body)
	if err != nil {
		err = fmt.Errorf("msg [%d] unmarshal error [%s]", msgType, err.Error())
		return
	}
	m = &Message{
		MessageType: MessageType(msgType),
		Body:        body,
	}

	return
}
func (c *Codec) Register(msgType MessageType, msg proto.Message) {

	_, exist := c.protoMap[msgType]
	if exist {
		panic(fmt.Errorf("repeat register message type %d", msgType))
	}

	typ := reflect.TypeOf(msg).Elem()

	c.protoMap[msgType] = typ
	c.revProtoMap[typ] = msgType
}

func (c *Codec) getMsgType(msg proto.Message) (msgType MessageType, exist bool) {
	typ := reflect.TypeOf(msg).Elem()
	msgType, exist = c.revProtoMap[typ]
	return
}

func (c *Codec) getMessage(msgType MessageType) (m proto.Message, ok bool) {
	typ, ok := c.protoMap[MessageType(msgType)]
	if !ok {
		return
	}

	bodyInter := reflect.New(typ).Interface()
	m, ok = bodyInter.(proto.Message)
	return
}

func NewCodec() *Codec {
	c := &Codec{}
	c.protoMap = make(map[MessageType]reflect.Type)
	c.revProtoMap = make(map[reflect.Type]MessageType)
	return c
}
