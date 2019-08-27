package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/equipbaoku/pbutil"
	playerequipbaoku "fgame/fgame/game/equipbaoku/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_SHOP_LIMIT_TYPE), dispatch.HandlerFunc(handleEquipBaoKuShopLimitGet))
}

//处理获取宝库商店限购次数信息
func handleEquipBaoKuShopLimitGet(s session.Session, msg interface{}) (err error) {
	log.Debug("equipbaoku:处理获取获取宝库商店限购次数消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = equipBaoKuShopLimitGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("equipbaoku:处理获取获取宝库商店限购次数消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("equipbaoku:处理获取获取宝库商店限购次数消息完成")
	return nil
}

//获取获取宝库商店限购次数信息的逻辑
func equipBaoKuShopLimitGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerEquipBaoKuDataManagerType).(*playerequipbaoku.PlayerEquipBaoKuDataManager)
	shops := manager.GetEquipBaoKuShopBuyAll()
	scShopLimit := pbutil.BuildSCEquipBaoKuShopLimit(shops)
	err = pl.SendMsg(scShopLimit)
	return
}
