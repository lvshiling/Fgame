package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/lineup/lineup"
	"fgame/fgame/cross/lineup/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	crosstypes "fgame/fgame/game/cross/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_LINEUP_ATTEND_TYPE), dispatch.HandlerFunc(handleLineupAttend))
}

func handleLineupAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("lineup:处理跨服通用排队，参与")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siMsg := msg.(*crosspb.SILineupAttend)
	typeInt := siMsg.GetCrossType()
	crossType := crosstypes.CrossType(typeInt)
	sceneId := siMsg.GetSceneId()

	err = lineupAttend(tpl, crossType, sceneId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("lineup:处理跨服通用排队,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lineup:处理跨服通用排队,完成")

	return nil
}

//跨服排队
func lineupAttend(pl *player.Player, crossType crosstypes.CrossType, sceneId int64) (err error) {
	pos, flag := lineup.GetLineupService().GetHasLineUp(crossType, sceneId, pl.GetId())
	if flag {
		isMsg := pbutil.BuildISLineupAttend(pos, int32(crossType), sceneId)
		pl.SendMsg(isMsg)
		return
	}

	pos = lineup.GetLineupService().Attend(crossType, sceneId, pl.GetId())
	isMsg := pbutil.BuildISLineupAttend(pos, int32(crossType), sceneId)
	pl.SendMsg(isMsg)
	return
}
