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
	"fgame/fgame/game/shop/pbutil"
	playershop "fgame/fgame/game/shop/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHOP_LIMIT_TYPE), dispatch.HandlerFunc(handleShopLimitGet))
}

//处理获取商铺限购次数信息
func handleShopLimitGet(s session.Session, msg interface{}) (err error) {
	log.Debug("shop:处理获取获取商铺限购次数消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = shopLimitGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("shop:处理获取获取商铺限购次数消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shop:处理获取获取商铺限购次数消息完成")
	return nil
}

//获取获取商铺限购次数信息的逻辑
func shopLimitGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerShopDataManagerType).(*playershop.PlayerShopDataManager)
	shops := manager.GetShopBuyAll()
	scShopLimit := pbutil.BuildSCShopLimit(shops)
	err = pl.SendMsg(scShopLimit)
	return
}
