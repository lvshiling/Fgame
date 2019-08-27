package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/equipbaoku/equipbaoku"
	"fgame/fgame/game/equipbaoku/pbutil"
	playerequipbaoku "fgame/fgame/game/equipbaoku/player"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeEquipBaoKu) {
		return
	}
	equipTyp := equipbaokutypes.BaoKuTypeEquip
	materialTyp := equipbaokutypes.BaoKuTypeMaterials

	manager := pl.GetPlayerDataManager(playertypes.PlayerEquipBaoKuDataManagerType).(*playerequipbaoku.PlayerEquipBaoKuDataManager)
	equipObj := manager.GetEquipBaoKuObj(equipTyp)
	materialObj := manager.GetEquipBaoKuObj(materialTyp)

	err = manager.RefreshEquipBaoKuShop()
	if err != nil {
		return
	}

	equipLogList := equipbaoku.GetEquipBaoKuService().GetLogByTime(0, equipTyp)
	materialLogList := equipbaoku.GetEquipBaoKuService().GetLogByTime(0, materialTyp)
	shopBuyCountMap := manager.GetEquipBaoKuShopBuyAll()

	// 装备宝库信息
	scEquipBaoKuInfoGet := pbutil.BuildSCEquipBaoKuInfoGet(equipObj, materialObj, equipLogList, materialLogList, shopBuyCountMap)
	pl.SendMsg(scEquipBaoKuInfoGet)

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
