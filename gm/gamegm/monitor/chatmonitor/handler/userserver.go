package handler

import (
	"fgame/fgame/gm/gamegm/basic/pb"
	"fgame/fgame/gm/gamegm/monitor"
	chatpb "fgame/fgame/gm/gamegm/monitor/chatmonitor/pb/chat"
	messagetypepb "fgame/fgame/gm/gamegm/monitor/chatmonitor/pb/messagetype"
	"fgame/fgame/gm/gamegm/session"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func HandleUserServer(s session.Session, msg *pb.Message) error {
	log.Debug("玩家注册服务器", msg)
	pl := monitor.PlayerInContext(s.Context())
	qs := monitor.MonitorServiceInContext(s.Context())
	playerid := pl.Id()
	log.Debug("玩家注册服务器，玩家ID：", playerid)
	userMsg, err := proto.GetExtension(msg, chatpb.E_CgChatMinitor)
	if err != nil {
		//TODO 发送异常信息
		log.WithFields(
			log.Fields{
				"sessionId": s.Id(),
				"error":     err,
			}).Error("玩家设置服务器消息解析错误")
		monitor.CloseWithError(s, err)
		return err
	}

	cgUserMsg, ok := userMsg.(*chatpb.CGChatMinitor)
	if ok {
		// log.Debug("转换ok")
		// log.Debug("服务器个数", len(cgUserMsg.GetServerlist()))
	}

	service := qs.GetUserServerManage()
	serverArray := cgUserMsg.GetServerlist()
	service.SetUserServer(playerid, serverArray)

	gcUserServer := &chatpb.GCChatMinitor{}
	gcUserServer.PlayerId = &playerid
	gcMsg := &pb.Message{}
	gcMsgType := int32(messagetypepb.ChatMonitorMessageType_GCChatMinitorType)
	gcMsg.MessageType = &gcMsgType
	err = proto.SetExtension(gcMsg, chatpb.E_GcChatMinitor, gcUserServer)
	if err != nil {
		log.Debug("在这里异常：", err)
		monitor.CloseWithError(s, err)
		return err
	}

	gcMsgB, err := proto.Marshal(gcMsg)
	if err != nil {
		monitor.CloseWithError(s, err)
		return err
	}
	pl.Send(gcMsgB)

	return nil
}
