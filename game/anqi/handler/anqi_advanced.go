package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	anqilogic "fgame/fgame/game/anqi/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ANQI_ADVANCED_TYPE), dispatch.HandlerFunc(handleAnqiAdvanced))
}

//处理暗器进阶信息
func handleAnqiAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("anqi:处理获取暗器进阶消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAnqiAdvanced := msg.(*uipb.CSAnqiAdvanced)
	autoFlag := csAnqiAdvanced.GetAutoFlag()

	err = anqilogic.HandleAnqiAdvanced(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
				"error":    err,
			}).Error("anqi:处理获取暗器进阶消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"autoFlag": autoFlag,
		}).Debug("anqi:处理获取暗器进阶消息完成")
	return nil

}
