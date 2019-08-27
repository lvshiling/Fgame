package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/godsiege/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/game/godsiege/godsiege"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	playerlogic "fgame/fgame/game/player/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_GODSIEGE_ATTEND_TYPE), dispatch.HandlerFunc(handleGodSiegeAttend))
}

//处理参加神兽攻城
func handleGodSiegeAttend(s session.Session, msg interface{}) (err error) {
	log.Info("godsiege:处理参加神兽攻城")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siGodSiegeAttend := msg.(*crosspb.SIGodSiegeAttend)
	godType := siGodSiegeAttend.GetGodType()

	err = godSiegeAttend(tpl, godType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"godType":  godType,
				"err":      err,
			}).Error("godsiege:处理参加神兽攻城,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"godType":  godType,
		}).Info("godsiege:处理参加神兽攻城,完成")
	return nil

}

//参加神兽攻城
func godSiegeAttend(pl *player.Player, godType int32) (err error) {
	godSiegeType := godsiegetypes.GodSiegeType(godType)
	if !godSiegeType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"godType":  godType,
			}).Warn("godsiege:处理参加神兽攻城,godType类型错误")
		pl.Close(nil)
		return
	}

	flag := godsiege.GetGodSiegeService().IsGodSiegeActivityTime(godSiegeType)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"godType":  godType,
			}).Warn("godsiege:处理参加神兽攻城,活动未开始")
		//活动未开始
		playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
		pl.Close(nil)
		return
	}
	pos, isLineUp, flag := godsiege.GetGodSiegeService().GetHasLineUp(godSiegeType, pl.GetId())
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"godType":  godType,
			}).Warn("godsiege:处理参加神兽攻城,活动未开始")
		return
	}
	if isLineUp {
		isGodSiegeAttend := pbutil.BuildISGodSiegeAttend(godType, isLineUp, pos)
		pl.SendMsg(isGodSiegeAttend)
		return
	}
	pos, isLineUp, _ = godsiege.GetGodSiegeService().Attend(godSiegeType, pl.GetId())
	isGodSiegeAttend := pbutil.BuildISGodSiegeAttend(godType, isLineUp, pos)
	pl.SendMsg(isGodSiegeAttend)
	return
}
