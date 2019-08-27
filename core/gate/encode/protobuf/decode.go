package protobuf

import (
	"fgame/fgame/core/gate/encode/protobuf/pb"
	"fgame/fgame/core/session"

	"github.com/golang/protobuf/proto"
)

func Unmarshal(session session.Session, msg interface{}) (pMsg *pb.Message, err error) {
	msgBytes, exist := msg.([]byte)
	if !exist {
		return
	}

	pMsg = &pb.Message{}
	err = proto.Unmarshal(msgBytes, pMsg)
	if err != nil {
		return
	}
	return
}
