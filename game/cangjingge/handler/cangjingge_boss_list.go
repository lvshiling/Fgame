package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/cangjingge/cangjingge"
	"fgame/fgame/game/cangjingge/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_CANGJINGGE_BOSS_LIST_TYPE), dispatch.HandlerFunc(handlerCangJingGeBossList))
}

//藏经阁BOSS列表请求
func handlerCangJingGeBossList(s session.Session, msg interface{}) (err error) {
	log.Debug("cangjingge:处理藏经阁BOSS列表请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = cangJingGeBossList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("cangjingge:处理藏经阁BOSS列表请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("cangjingge:处理藏经阁BOSS列表请求完成")

	return
}

func cangJingGeBossList(pl player.Player) (err error) {
	bossList := cangjingge.GetCangJingGeService().GetCangJingGeBossList()
	scMsg := pbutil.BuildSCCangJingGeBossList(bossList)
	pl.SendMsg(scMsg)
	return
}
