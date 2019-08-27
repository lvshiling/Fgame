package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	massacrelogic "fgame/fgame/game/massacre/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MASSACRE_ADVANCED_TYPE), dispatch.HandlerFunc(handleMassacreAdvanced))
}

//处理戮仙刃进阶信息
func handleMassacreAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("massacre:处理获取戮仙刃进阶消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMassacreAdvanced := msg.(*uipb.CSMassacreAdvanced)
	autoFlag := csMassacreAdvanced.GetAutoFlag()

	err = massacrelogic.HandleMassacreAdvanced(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
				"error":    err,
			}).Error("massacre:处理获取戮仙刃进阶消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"autoFlag": autoFlag,
		}).Debug("massacre:处理获取戮仙刃进阶消息完成")
	return nil

}
