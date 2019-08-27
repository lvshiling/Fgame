package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	crosslogic "fgame/fgame/game/cross/logic"
	godsiegelogic "fgame/fgame/game/godsiege/logic"
	"fgame/fgame/game/godsiege/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_GODSIEGE_LINEUP_SUCCESS_TYPE), dispatch.HandlerFunc(handleGodSiegeLineUpSuccess))
}

//处理神兽攻城排队成功
func handleGodSiegeLineUpSuccess(s session.Session, msg interface{}) (err error) {
	log.Debug("godsiege:处理神兽攻城排队成功")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isGodSiegeLineUpSuccess := msg.(*crosspb.ISGodSiegeLineUpSuccess)
	godType := isGodSiegeLineUpSuccess.GetGodType()

	err = godSiegeLineUpSuccess(tpl, godType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"godType":  godType,
				"err":      err,
			}).Error("godsiege:处理神兽攻城排队成功,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("godsiege:处理神兽攻城排队成功,完成")
	return nil

}

//神兽攻城排队成功
func godSiegeLineUpSuccess(pl player.Player, godType int32) (err error) {
	pl.GodSiegeCancleLineUp()
	godsiegelogic.GodSiegeLineUpSuccess(pl, godType)
	scGodSiegeLineUpSuccess := pbutil.BuildSCGodSiegeLineUpSuccess(godType)
	pl.SendMsg(scGodSiegeLineUpSuccess)
	//进入跨服神兽攻城
	crosslogic.CrossPlayerDataLogin(pl)
	return
}
