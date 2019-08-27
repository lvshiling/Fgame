package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	crosslogic "fgame/fgame/game/cross/logic"
	"fgame/fgame/game/lineup/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_LINEUP_SUCCESS_TYPE), dispatch.HandlerFunc(handleLineUpSuccess))
}

//处理排队成功
func handleLineUpSuccess(s session.Session, msg interface{}) (err error) {
	log.Debug("lianyu:处理排队成功")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isMsg := msg.(*crosspb.ISLineupSuccess)
	crossType := isMsg.GetCrossType()

	err = lineupLineUpSuccess(tpl, crossType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("lianyu:处理排队成功,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lianyu:处理排队成功,完成")
	return nil

}

//排队成功
func lineupLineUpSuccess(pl player.Player, crossType int32) (err error) {
	pl.SetLineUp(false)

	scMsg := pbutil.BuildSCLineupSuccess(crossType)
	pl.SendMsg(scMsg)

	//进入跨服
	crosslogic.CrossPlayerDataLogin(pl)
	return
}
