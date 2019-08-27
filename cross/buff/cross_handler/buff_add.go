package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	"fgame/fgame/game/common/common"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_BUFF_ADD_TYPE), dispatch.HandlerFunc(handlePlayerBuffAdd))
}

//玩家仙盟变化
func handlePlayerBuffAdd(s session.Session, msg interface{}) (err error) {
	log.Debug("buff:玩家buff添加")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)
	if tpl.GetScene() == nil{
		return
	}
	siBuffAdd := msg.(*crosspb.SIBuffAdd)
	buffId := siBuffAdd.GetBuffData().GetBuffId()
	scenelogic.AddBuff(tpl, buffId, 0, common.MAX_RATE)
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("buff:玩家buff添加,完成")
	return nil

}
