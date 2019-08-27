package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/processor"
	"fgame/fgame/game/common/pbutil"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GET_TIME_TYPE), dispatch.HandlerFunc(handleGetTime))
}

//获取时间
func handleGetTime(s session.Session, msg interface{}) error {
	log.Debug("common:处理获取时间信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)
	now := global.GetGame().GetTimeService().Now()
	scGetTime := pbutil.BuildSCGetTime(now)
	err := tpl.SendMsg(scGetTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("common:处理获取时间信息,错误")
		return err
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("common:处理获取时间信息完成")
	return nil
}
