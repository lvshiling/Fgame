package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	yuxilogic "fgame/fgame/game/yuxi/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_YUXI_GET_IFNO_TYPE), dispatch.HandlerFunc(handleYuXiGetInfo))
}

//处理玉玺之战信息
func handleYuXiGetInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("yuxi:处理玉玺之战信息消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = yuXiGetInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("yuxi:处理玉玺之战信息消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("yuxi:处理玉玺之战信息消息完成")
	return nil
}

//处理玉玺之战信息信息逻辑
func yuXiGetInfo(pl player.Player) (err error) {
	yuxilogic.SendYuXiWarInfo(pl)
	return
}
