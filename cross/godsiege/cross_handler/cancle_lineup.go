package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/godsiege/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/game/godsiege/godsiege"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_GODSIEGE_CANCLE_LINEUP_TYPE), dispatch.HandlerFunc(handleGodSiegeCancleLineUp))
}

//处理神兽攻城取消排队
func handleGodSiegeCancleLineUp(s session.Session, msg interface{}) (err error) {
	log.Debug("godsiege:处理神兽攻城取消排队")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siGodSiegeCancleLineUp := msg.(*crosspb.SIGodSiegeCancleLineUp)
	godType := siGodSiegeCancleLineUp.GetGodType()

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
		}).Debug("godsiege:处理神兽攻城取消排队,完成")
	return nil

}

//神兽攻城取消排队
func godSiegeCancleLineUp(pl *player.Player, godType int32) (err error) {
	godSiegeType := godsiegetypes.GodSiegeType(godType)
	if !godSiegeType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"godType":  godType,
			}).Info("godsiege:处理参加神兽攻城,godType类型错误")
		pl.Close(nil)
		return
	}
	flag := godsiege.GetGodSiegeService().CancleLineUp(godSiegeType, pl.GetId())
	if !flag {
		return
	}
	isGodSiegeCancleUp := pbutil.BuildISGodSiegeCancleUp(godType)
	pl.SendMsg(isGodSiegeCancleUp)
	return
}
