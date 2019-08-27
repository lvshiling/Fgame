package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	shihunfanlogic "fgame/fgame/game/shihunfan/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHIHUNFAN_ADVANCED_TYPE), dispatch.HandlerFunc(handleShiHunFanAdvanced))
}

//处理噬魂幡进阶信息
func handleShiHunFanAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("shihunfan:处理获取噬魂幡进阶消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csShiHunFanAdvanced := msg.(*uipb.CSShihunfanAdvanced)
	autoFlag := csShiHunFanAdvanced.GetBuyFlag()

	err = shihunfanlogic.HandleShiHunFanAdvanced(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
				"error":    err,
			}).Error("shihunfan:处理获取噬魂幡进阶消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"autoFlag": autoFlag,
		}).Debug("shihunfan:处理获取噬魂幡进阶消息完成")
	return nil

}
