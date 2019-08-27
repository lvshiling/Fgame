package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/transportation/pbutil"
	"fgame/fgame/game/transportation/transpotation"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TRANSPORTATION_PROTECT_TYPE), dispatch.HandlerFunc(handlerProtectTrasnportation))
}

//护镖请求
func handlerProtectTrasnportation(s session.Session, msg interface{}) (err error) {
	log.Debug("transport:处理护镖请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = protectTrasnportation(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("transport:处理护镖请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("transport:处理护镖请求完成")

	return
}

func protectTrasnportation(pl player.Player) (err error) {
	biaoChe := transpotation.GetTransportService().GetTransportation(pl.GetId())
	if biaoChe == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("transport:镖车不存在")
		return
	}

	//玩家跳转场景
	pos := biaoChe.GetPosition()
	s := biaoChe.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("transport:镖车不在场景中")
		return
	}

	playerScene := pl.GetScene()
	if playerScene == s {
		scenelogic.FixPosition(pl, pos)
	} else {
		scenelogic.PlayerEnterScene(pl, s, pos)
	}

	scTransportationProtect := pbutil.BuildSCTransportationProtect()
	pl.SendMsg(scTransportationProtect)

	return
}
