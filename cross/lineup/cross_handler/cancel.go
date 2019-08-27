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
	processor.Register(codec.MessageType(crosspb.MessageType_SI_LINEUP_CANCEL_TYPE), dispatch.HandlerFunc(handleLineupCancel))
}

func handleLineupCancel(s session.Session, msg interface{}) (err error) {
	log.Debug("lineup:处理跨服通用排队，参与")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siMsg := msg.(*crosspb.SILineupCancle)
	typeInt := siMsg.GetCrossType()
	crossType := crosstypes.CrossType(typeInt)
	sceneId := siMsg.GetSceneId()

	err = lineupCancel(tpl, crossType, sceneId)
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
func lineupCancel(pl *player.Player, crossType crosstypes.CrossType, sceneId int64) (err error) {
	flag := lineup.GetLineupService().CancleLineUp(crossType, sceneId, pl.GetId())
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"crossType": crossType,
			}).Infoln("lineup:处理跨服通用取消排队失败")
		return
	}

	isMsg := pbutil.BuildISLineupCancle(int32(crossType))
	pl.SendMsg(isMsg)
	return
}
