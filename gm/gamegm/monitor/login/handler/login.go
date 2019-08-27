package handler

import (
	"fgame/fgame/gm/gamegm/basic/pb"
	monitor "fgame/fgame/gm/gamegm/monitor"
	loginpb "fgame/fgame/gm/gamegm/monitor/login/pb"
	"fgame/fgame/gm/gamegm/session"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

//处理登录
func HandleLogin(s session.Session, msg *pb.Message) error {
	log.Debug("处理登陆消息,", msg)
	pl := monitor.PlayerInContext(s.Context())
	//玩家重复登录
	if !pl.IsInit() {
		//TODO 发送错误信息
		log.WithFields(
			log.Fields{
				"playerId": pl.Id(),
			}).Warn("玩家重复登陆")
		return nil
	}

	loginMsg, err := proto.GetExtension(msg, loginpb.E_CgLogin)
	if err != nil {
		//TODO 发送异常信息
		log.WithFields(
			log.Fields{
				"sessionId": s.Id(),
				"error":     err,
			}).Error("玩家登陆消息解析错误")
		return err
	}

	cgLogin := loginMsg.(*loginpb.CGLogin)
	token := cgLogin.GetToken()
	qs := monitor.MonitorServiceInContext(s.Context())
	//登录
	playerId, err := qs.Login(pl, token)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"error":    err,
			}).Error("玩家登录失败")
		monitor.CloseWithError(s, err)
		return err
	}
	return nil

}
