package cmd

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
)

type CmdType int32

var (
	cmdMap    = make(map[CmdType]reflect.Type)
	revCmdMap = make(map[reflect.Type]CmdType)
)

func RegisterCmd(cmdType CmdType, msg proto.Message) {
	_, exist := cmdMap[cmdType]
	if exist {
		panic(fmt.Errorf("repeat register cmd type %d", cmdType))
	}

	typ := reflect.TypeOf(msg).Elem()

	cmdMap[cmdType] = typ
	revCmdMap[typ] = cmdType
}

func GetTypeForCmd(msg proto.Message) (cmdType CmdType, exist bool) {
	typ := reflect.TypeOf(msg).Elem()
	cmdType, exist = revCmdMap[typ]
	return
}

func GetCmdForType(typ CmdType) proto.Message {
	rType, ok := cmdMap[typ]
	if !ok {
		return nil
	}
	bodyInter := reflect.New(rType).Interface()
	m, ok := bodyInter.(proto.Message)
	if !ok {
		return nil
	}
	return m
}

type CmdHandler interface {
	HandleCmd(msg proto.Message) (err error)
}

type CmdHandlerFunc func(msg proto.Message) (err error)

func (f CmdHandlerFunc) HandleCmd(msg proto.Message) (err error) {
	return f(msg)
}

var (
	cmdHandlerMap = make(map[CmdType]CmdHandler)
)

func RegisterCmdHandler(cmdType CmdType, h CmdHandler) {
	_, ok := cmdHandlerMap[cmdType]
	if ok {
		panic(fmt.Errorf("重复注册命令%d", cmdType))
	}
	cmdHandlerMap[cmdType] = h
}

func HandlerCmd(cmdType CmdType, msg proto.Message) (err error) {
	h, ok := cmdHandlerMap[cmdType]
	if !ok {
		err = ErrorCodeCommonCmdHandlerNoFound
		return
	}
	return h.HandleCmd(msg)
}
