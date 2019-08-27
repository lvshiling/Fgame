package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/dragon/pbutil"
	playerdragon "fgame/fgame/game/dragon/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DRAGON_GET_TYPE), dispatch.HandlerFunc(handleDragonGet))
}

//处理神龙信息
func handleDragonGet(s session.Session, msg interface{}) (err error) {
	log.Debug("dragon:处理获取神龙消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = dragonGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("dragon:处理获取神龙消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("dragon:处理获取神龙消息完成")
	return nil
}

//处理神龙界面信息逻辑
func dragonGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerDragonDataManagerType).(*playerdragon.PlayerDragonDataManager)
	dragonInfo := manager.GetDragon()
	scDragonGet := pbuitl.BuildSCDragonGet(dragonInfo)
	pl.SendMsg(scDragonGet)
	return
}
