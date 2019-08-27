package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/xuechi/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XUECHI_GET_TYPE), dispatch.HandlerFunc(handleXueChiGet))
}

//处理血池信息
func handleXueChiGet(s session.Session, msg interface{}) (err error) {
	log.Debug("xuechi:处理获取血池消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = xueChiGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("xuechi:处理获取血池消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("xuechi:处理获取血池消息完成")
	return nil

}

//获取血池界面信息的逻辑
func xueChiGet(pl scene.Player) (err error) {
	bloodLine := pl.GetBloodLine()
	blood := pl.GetBlood()
	scXueChiGet := pbutil.BuildSCXueChiGet(bloodLine, blood)
	pl.SendMsg(scXueChiGet)
	return
}
