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
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_ATTEND_TYPE), dispatch.HandlerFunc(handleArenapvpAttend))
}

//处理跨服参加pvp
func handleArenapvpAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("pvp:处理跨服参加pvp")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isMsg := msg.(*crosspb.ISArenapvpAttend)
	isLineUp := isMsg.GetIsLineUp()
	sceneId := isMsg.GetSceneId()
	err = arenapvpAttend(tpl, isLineUp, sceneId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"isLineUp": isLineUp,
				"err":      err,
			}).Error("pvp:处理跨服参加pvp,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("pvp:处理跨服参加pvp,完成")
	return nil

}

//参加pvp
func arenapvpAttend(pl player.Player, isLineUp bool, sceneId int64) (err error) {
	if !isLineUp {
		//进入跨服pvp
		crosslogic.CrossPlayerDataLogin(pl)
	} else {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"isLineUp": isLineUp,
			}).Infoln("pvp:处理跨服参加比武大会，排队")
		lineuplogic.SendCrossLineup(pl, crosstypes.CrossTypeArenapvp, sceneId)
	}
	return
}
