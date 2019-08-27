package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	godsiegelogic "fgame/fgame/game/godsiege/logic"
	"fgame/fgame/game/godsiege/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GODSIEGE_CANCLE_LINEUP_TYPE), dispatch.HandlerFunc(handleGodSiegeCancleLineUp))
}

//处理神兽攻城取消排队
func handleGodSiegeCancleLineUp(s session.Session, msg interface{}) (err error) {
	log.Debug("godsiege:处理神兽攻城取消排队")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csGodSiegeCancleLineUp := msg.(*uipb.CSGodSiegeCancleLineUp)
	godType := csGodSiegeCancleLineUp.GetGodType()

	err = godSiegeCancleLineUp(tpl, godType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"godType":  godType,
			}).Error("godsiege:处理神兽攻城取消排队,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"godType":  godType,
		}).Debug("godsiege:处理神兽攻城取消排队")
	return nil

}

//处理神兽攻城取消排队
func godSiegeCancleLineUp(pl player.Player, godType int32) (err error) {
	flag := pl.IsGodSiegeLineUp()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"godType":  godType,
		}).Warn("godsiege:您当前未在排队中")
		playerlogic.SendSystemMessage(pl, lang.GodSiegeCancleLineUpNoExist)
		return
	}

	_, curGodType := pl.GetGodSiegeLineUp()
	if int32(curGodType) != godType {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"godType":  godType,
		}).Warn("godsiege:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	godsiegelogic.GodSiegeCancleLineUpSend(pl, int32(godType))
	scGodSiegeCancleLineUp := pbutil.BuildSCGodSiegeCancleLineUp(godType)
	pl.SendMsg(scGodSiegeCancleLineUp)
	return
}
