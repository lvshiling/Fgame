package handler

import (
	"fgame/fgame/gm/gamegm/basic/pb"
	monitor "fgame/fgame/gm/gamegm/monitor"
	loginpb "fgame/fgame/gm/gamegm/monitor/login/pb"
	messagetypepb "fgame/fgame/gm/gamegm/monitor/messagetype/pb"
	"fgame/fgame/gm/gamegm/session"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

//处理ping
func HandlePing(s session.Session, msg *pb.Message) error {
	log.Debug("处理ping消息", msg)
	pl := monitor.PlayerInContext(s.Context())
	if pl == nil {
		log.WithFields(
			log.Fields{
				"sessionId": s.Id(),
			}).Warn("ping,玩家还没登陆")
		return nil
	}

	flag := pl.Ping()
	if !flag {
		log.Warn("ping,超时")
		return nil
	}

	m := &pb.Message{}
	pingMsgType := int32(messagetypepb.QiPaiMessageType_GCPingType)
	m.MessageType = &pingMsgType
	gcPing := &loginpb.GCPing{}
	now := time.Now().UnixNano() / int64(time.Millisecond)
	gcPing.Now = &now
	err := proto.SetExtension(m, loginpb.E_GcPing, gcPing)
	if err != nil {
		return err
	}
	msgBytes, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	pl.Send(msgBytes)
	return nil
}
