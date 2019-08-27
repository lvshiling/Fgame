package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/lineup/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_LINEUP_ATTEND_TYPE), dispatch.HandlerFunc(handleLineupAttend))
}

//处理跨服参加排队
func handleLineupAttend(s session.Session, msg interface{}) (err error) {
	log.Debug("lianyu:处理跨服参加排队")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	isMsg := msg.(*crosspb.ISLineupAttend)
	beforeNum := isMsg.GetBeforeNum()
	typeInt := isMsg.GetCrossType()
	crossType := crosstypes.CrossType(typeInt)
	sceneId := isMsg.GetSceneId()

	err = lineupAttend(tpl, crossType, beforeNum, sceneId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"beforeNum": beforeNum,
				"err":       err,
			}).Error("lianyu:处理跨服参加排队,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"beforeNum": beforeNum,
		}).Debug("lianyu:处理跨服参加排队,完成")
	return nil

}

//参加排队
func lineupAttend(pl player.Player, crossType crosstypes.CrossType, beforeNum int32, sceneId int64) (err error) {

	//排队通知
	pl.SetLineUp(true)
	scMsg := pbutil.BuildSCLineupNotice(beforeNum, int32(crossType))
	pl.SendMsg(scMsg)
	return
}
