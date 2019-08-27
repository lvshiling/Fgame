package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/equipbaoku/equipbaoku"
	"fgame/fgame/game/equipbaoku/pbutil"
	playerequipbaoku "fgame/fgame/game/equipbaoku/player"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EQUIPBAOKU_INFO_GET_TYPE), dispatch.HandlerFunc(handleEquipBaoKuInfoGet))

}

//处理装备宝库信息
func handleEquipBaoKuInfoGet(s session.Session, msg interface{}) (err error) {
	log.Debug("equipbaoku:处理获取装备宝库消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = equipBaoKuInfoGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("equipbaoku:处理获取装备宝库消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("equipbaoku:处理获取装备宝库消息完成")
	return nil

}

//获取装备宝库界面信息逻辑
func equipBaoKuInfoGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerEquipBaoKuDataManagerType).(*playerequipbaoku.PlayerEquipBaoKuDataManager)
	equipType := equipbaokutypes.BaoKuTypeEquip
	materialType := equipbaokutypes.BaoKuTypeMaterials

	equipObj := manager.GetEquipBaoKuObj(equipType)
	equipLogList := equipbaoku.GetEquipBaoKuService().GetLogByTime(0, equipType)
	materialObj := manager.GetEquipBaoKuObj(materialType)
	materialLogList := equipbaoku.GetEquipBaoKuService().GetLogByTime(0, materialType)
	shopBuyCountMap := manager.GetEquipBaoKuShopBuyAll()

	scEquipBaoKuInfoGet := pbutil.BuildSCEquipBaoKuInfoGet(equipObj, materialObj, equipLogList, materialLogList, shopBuyCountMap)
	pl.SendMsg(scEquipBaoKuInfoGet)
	return
}
