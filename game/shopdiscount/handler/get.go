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
	"fgame/fgame/game/shopdiscount/pbutil"
	playershopdiscount "fgame/fgame/game/shopdiscount/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHOP_DISCOUNT_GET_TYPE), dispatch.HandlerFunc(handleShopDiscountGet))
}

//处理获取商城促销信息
func handleShopDiscountGet(s session.Session, msg interface{}) (err error) {
	log.Debug("shopdiscount:处理获取获取商城促销消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = shopDiscountGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("shopdiscount:处理获取获取商城促销消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shopdiscount:处理获取获取商城促销消息完成")
	return nil

}

//获取获取商城促销界面信息的逻辑
func shopDiscountGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerShopDiscountDataManagerType).(*playershopdiscount.PlayerShopDiscountDataManager)
	obj := manager.GetShopDiscountObj()
	scShopDiscountGet := pbutil.BuildSCShopDiscountGet(obj)
	pl.SendMsg(scShopDiscountGet)
	return
}
