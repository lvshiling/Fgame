package handler

import (
	"fgame/fgame/client/player/player"
	"fgame/fgame/client/processor"
	clientsession "fgame/fgame/client/session"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_SC_SELECT_JOB_TYPE), dispatch.HandlerFunc(handleSelectJob))
}

//处理角色消息
func handleSelectJob(s session.Session, msg interface{}) (err error) {
	log.Debug("login:选择角色成功")
	cs := clientsession.SessionInContext(s.Context())
	pl := cs.Player().(*player.Player)
	err = playerSelectJob(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.Id(),
				"err":      err,
			}).Error("login:选择角色成功完成")
		return
	}
	log.Debug("login:选择角色成功完成")
	return
}

func playerSelectJob(pl *player.Player) (err error) {

	log.WithFields(
		log.Fields{
			"playerId": pl.Id(),
		}).Info("login:进入选择角色界面")

	return nil
}
