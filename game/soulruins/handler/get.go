package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/soulruins/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SOULRUINS_GET_TYPE), dispatch.HandlerFunc(handleSoulRuinsGet))
}

//处理帝陵遗迹信息
func handleSoulRuinsGet(s session.Session, msg interface{}) (err error) {
	log.Debug("soulruins:处理获取帝陵遗迹消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = soulRuinsGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("soulruins:处理获取帝陵遗迹消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("soulruins:处理获取帝陵遗迹消息完成")
	return nil

}

//获取帝陵遗迹界面信息的逻辑
func soulRuinsGet(pl player.Player) (err error) {
	scSoulRuinsGet := pbutil.BuildSCSoulRuinsGet(pl)
	pl.SendMsg(scSoulRuinsGet)
	return
}
