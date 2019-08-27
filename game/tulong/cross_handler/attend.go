package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	lineuplogic "fgame/fgame/game/lineup/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_TULONG_ATTEND_TYPE), dispatch.HandlerFunc(handleTuLongAttend))
}

//处理跨服参加仙盟屠龙
func handleTuLongAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理跨服参加仙盟屠龙")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isTuLongAttend := msg.(*crosspb.ISTuLongAttend)
	isLineUp := isTuLongAttend.GetIsLineUp()
	sceneId := isTuLongAttend.GetSceneId()
	err = tuLongAttend(tpl, isLineUp, sceneId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"isLineUp": isLineUp,
				"sceneId":  sceneId,
				"err":      err,
			}).Error("shenmo:处理跨服参加仙盟屠龙,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"sceneId":  sceneId,
		}).Debug("shenmo:处理跨服参加仙盟屠龙,完成")
	return nil

}

//参加仙盟屠龙
func tuLongAttend(pl player.Player, isLineUp bool, sceneId int64) (err error) {
	if !isLineUp {
		//进入跨服仙盟屠龙
		crosslogic.CrossPlayerDataLogin(pl)
	} else {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"isLineUp": isLineUp,
			}).Infoln("pvp:处理跨服参加屠龙，排队")
		lineuplogic.SendCrossLineup(pl, crosstypes.CrossTypeTuLong, sceneId)
	}
	return
}
