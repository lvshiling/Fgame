package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	crosslogic "fgame/fgame/game/cross/logic"
	"fgame/fgame/game/godsiege/pbutil"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_GODSIEGE_ATTEND_TYPE), dispatch.HandlerFunc(handleGodSiegeAttend))
}

//处理跨服参加神兽攻城
func handleGodSiegeAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("godsiege:处理跨服参加神兽攻城")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isGodSiegeAttend := msg.(*crosspb.ISGodSiegeAttend)
	godType := isGodSiegeAttend.GetGodType()
	isLineUp := isGodSiegeAttend.GetIsLineUp()
	beforeNum := isGodSiegeAttend.GetBeforeNum()
	err = godSiegeAttend(tpl, godType, isLineUp, beforeNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"godType":   godType,
				"isLineUp":  isLineUp,
				"beforeNum": beforeNum,
				"err":       err,
			}).Error("godsiege:处理跨服参加神兽攻城,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"godType":   godType,
			"isLineUp":  isLineUp,
			"beforeNum": beforeNum,
		}).Debug("godsiege:处理跨服参加神兽攻城,完成")
	return nil

}

//参加神兽攻城
func godSiegeAttend(pl player.Player, godType int32, isLineUp bool, beforeNum int32) (err error) {
	if !isLineUp {
		//进入跨服神兽攻城
		crosslogic.CrossPlayerDataLogin(pl)
	} else {
		pl.GodSiegeLineUp(godsiegetypes.GodSiegeType(godType))
		scGodSiegeLineUp := pbutil.BuildSCGodSiegeLineUp(godType, beforeNum)
		pl.SendMsg(scGodSiegeLineUp)
	}
	return
}
