package log

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type LogMsg interface {
	LogName() string
	// bson.Setter
}

var (
	logMsgMap map[string]reflect.Type
)

func init() {
	logMsgMap = make(map[string]reflect.Type)
}

func RegisterLogMsg(to LogMsg) {
	_, exist := logMsgMap[to.LogName()]
	if exist {
		panic(fmt.Sprintf("重复注册%s日志", to.LogName()))
	}
	logMsgMap[to.LogName()] = reflect.TypeOf(to)
}

func getLogMsg(logType string) LogMsg {
	typ, ok := logMsgMap[logType]
	if !ok {
		return nil
	}
	elem := typ.Elem()
	val := reflect.New(elem).Interface()
	msg, ok := val.(LogMsg)
	if !ok {
		return nil
	}
	return msg
}

func GetLogMsgList(logType string) interface{} {
	typ, ok := logMsgMap[logType]
	if !ok {
		return nil
	}
	sliceTyp := reflect.SliceOf(typ)
	sliceVal := reflect.MakeSlice(sliceTyp, 0, 16)
	x := reflect.New(sliceVal.Type())
	val := x.Interface()
	return val
}

func Decode(logType string, content []byte) (msg LogMsg, err error) {
	msg = getLogMsg(logType)
	if msg == nil {
		err = fmt.Errorf("不能创建消息%s", logType)
		return
	}
	err = json.Unmarshal(content, msg)
	if err != nil {
		return
	}

	return
}
