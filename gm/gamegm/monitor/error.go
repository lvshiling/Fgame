package monitor

import (
	"errors"
	"fgame/fgame/gm/gamegm/basic/pb"
	commonpb "fgame/fgame/gm/gamegm/monitor/common/pb"
	messagetypepb "fgame/fgame/gm/gamegm/monitor/messagetype/pb"
	"fgame/fgame/gm/gamegm/session"

	"github.com/golang/protobuf/proto"
)

var (
	ErrorPlayerLoginSameTime = errors.New("玩家多个设备同时登录")
	ErrorPlayerAuthTimeout   = errors.New("认证超时")
	ErrorPlayerNoAuth        = errors.New("玩家还没验证")
)

var (
	errorMap = map[error]commonpb.ErrorCode{
		ErrorPlayerAuthTimeout: commonpb.ErrorCode_AuthTimeout,
	}
)

func CodeForError(err error) commonpb.ErrorCode {
	code, flag := errorMap[err]
	if !flag {
		return commonpb.ErrorCode_Unknown
	}
	return code
}

//发送错误代码
func CloseWithError(s session.Session, terr error) error {
	errorCode := CodeForError(terr)
	gcError := &commonpb.GCError{}
	gcError.ErrorCode = &errorCode
	msg := &pb.Message{}
	msgType := int32(messagetypepb.QiPaiMessageType_GCError)
	msg.MessageType = &msgType
	err := proto.SetExtension(msg, commonpb.E_GcError, gcError)
	if err != nil {
		return err
	}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	err = s.Send(msgBytes)
	if err != nil {
		return err
	}
	return s.Close()
}
