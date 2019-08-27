package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	lineuplogic "fgame/fgame/game/lineup/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_LINEUP_SCENE_FINISH_TO_CANCEL_TYPE), dispatch.HandlerFunc(handleLineupFinishLineUpCancle))
}

//处理场景结束通知排队人员
func handleLineupFinishLineUpCancle(s session.Session, msg interface{}) (err error) {
	log.Debug("lianyu:处理场景结束通知排队人员")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = lineupFinishLineUpCancle(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("lianyu:处理场景结束通知排队人员,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lianyu:处理场景结束通知排队人员,完成")
	return nil

}

//场景结束通知排队人员
func lineupFinishLineUpCancle(pl player.Player) (err error) {
	lineuplogic.CancelCrossLineup(pl)
	return
}
