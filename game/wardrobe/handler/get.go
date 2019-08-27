package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/wardrobe/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WARDROBE_GET_TYPE), dispatch.HandlerFunc(handleWardrobeGet))
}

//处理获取衣橱信息
func handleWardrobeGet(s session.Session, msg interface{}) (err error) {
	log.Debug("wardrobe:处理获取衣橱信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = wardrobeGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("wardrobe:处理获取衣橱信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("wardrobe:处理获取衣橱信息完成")
	return nil
}

//处理衣橱界面信息逻辑
func wardrobeGet(pl player.Player) (err error) {
	scWardrobeGet := pbutil.BuildSCWardrobeGet(pl)
	pl.SendMsg(scWardrobeGet)
	return
}
