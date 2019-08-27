package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/dan/pbutil"
	playerdan "fgame/fgame/game/dan/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DAN_GET_TYPE), dispatch.HandlerFunc(handleDanGet))
}

//处理食丹信息
func handleDanGet(s session.Session, msg interface{}) (err error) {
	log.Debug("dan:处理获取食丹消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = danGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("dan:处理获取食丹消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("dan:处理获取食丹消息完成")
	return nil
}

//处理食丹界面信息逻辑
func danGet(pl player.Player) (err error) {
	danManager := pl.GetPlayerDataManager(types.PlayerDanDataManagerType).(*playerdan.PlayerDanDataManager)
	danInfo := danManager.GetDanInfo()
	scDanGet := pbuitl.BuildSCDanGet(danInfo)
	pl.SendMsg(scDanGet)
	return
}
