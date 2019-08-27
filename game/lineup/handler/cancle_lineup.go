package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/lineup/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINEUP_CANCEL_TYPE), dispatch.HandlerFunc(handleCancleLineUp))
}

//处理取消排队
func handleCancleLineUp(s session.Session, msg interface{}) (err error) {
	log.Debug("lineup:处理取消排队")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSLineupCancel)
	crossTypeInt := csMsg.GetCrossType()
	sceneId := csMsg.GetSceneId()

	crossType := crosstypes.CrossType(crossTypeInt)
	if !crossType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("lineup:处理取消排队,参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = lineupCancleLineUp(tpl, crossType, sceneId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("lineup:处理取消排队,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lineup:处理取消排队")
	return nil

}

//处理取消排队
func lineupCancleLineUp(pl player.Player, crossType crosstypes.CrossType, sceneId int64) (err error) {
	if !pl.IsLineUp() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("lineup:您当前未在排队中")
		playerlogic.SendSystemMessage(pl, lang.PlayerLineUpNoExist)
		return
	}

	isMsg := pbutil.BuildSILineupCancle(int32(crossType), sceneId)
	pl.SendCrossMsg(isMsg)

	scMSg := pbutil.BuildSCLineupCancel(int32(crossType))
	pl.SendMsg(scMSg)
	return
}
