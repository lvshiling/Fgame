package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_WED_TRANSFER_TYPE), dispatch.HandlerFunc(handleMarryTransfer))
}

//处理婚礼传送信息
func handleMarryTransfer(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理婚礼传送消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = marryTransfer(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("marry:处理婚礼传送消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理婚礼传送消息完成")
	return nil
}

//处理婚礼传送信息逻辑
func marryTransfer(pl player.Player) (err error) {
	err = marrylogic.PlayerEnterMarry(pl)
	return
}
