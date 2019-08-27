package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/soul/pbutil"
	playersoul "fgame/fgame/game/soul/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SOUL_GET_TYPE), dispatch.HandlerFunc(handleSoulGet))
}

//处理帝魂信息
func handleSoulGet(s session.Session, msg interface{}) (err error) {
	log.Debug("soul:处理获取帝魂消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = soulGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("soul:处理获取帝魂消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("soul:处理获取帝魂消息完成")
	return nil

}

//获取帝魂界面信息的逻辑
func soulGet(pl player.Player) (err error) {
	soulManager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	embedList := soulManager.GetSoulEmbed()
	soulInfo := soulManager.GetSoulInfoAll()
	scSoulGet := pbutil.BuildSCSoulGet(embedList, soulInfo)
	pl.SendMsg(scSoulGet)
	return
}
