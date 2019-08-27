package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	tianmologic "fgame/fgame/game/tianmo/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TIANMOTI_ADVANCED_TYPE), dispatch.HandlerFunc(handleTianMoAdvanced))
}

//处理天魔进阶信息
func handleTianMoAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("tianMo:处理获取天魔进阶消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSTianMoAdvanced)
	autoFlag := csMsg.GetAutoFlag()

	err = tianmologic.HandleTianMoAdvanced(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
				"error":    err,
			}).Error("tianMo:处理获取天魔进阶消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"autoFlag": autoFlag,
		}).Debug("tianMo:处理获取天魔进阶消息完成")
	return nil

}
