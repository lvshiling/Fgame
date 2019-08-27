package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_BUFF_REMOVE_TYPE), dispatch.HandlerFunc(handlePlayerBuffRemove))
}

//玩家仙盟变化
func handlePlayerBuffRemove(s session.Session, msg interface{}) (err error) {
	log.Debug("buff:玩家buff移除")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)
	if tpl.GetScene() == nil {
		return
	}
	siBuffRemove := msg.(*crosspb.SIBuffRemove)
	buffId := siBuffRemove.GetBuffId()
	scenelogic.RemoveBuff(tpl, buffId)
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("buff:玩家buff移除,完成")
	return nil

}
