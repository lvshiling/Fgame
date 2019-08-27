package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/wushuangweapon/pbutil"
	playerwushuangweapon "fgame/fgame/game/wushuangweapon/player"
	wushuangweapontemplate "fgame/fgame/game/wushuangweapon/template"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	wushuangDataManager := pl.GetPlayerDataManager(playertypes.PlayerWushuangWeaponDataManagerType).(*playerwushuangweapon.PlayerWushuangWeaponDataManager)
	var bodyposinfoList []*playerwushuangweapon.PlayerWushuangWeaponSlotObject
	for _, slotObj := range wushuangDataManager.GetSlotObjectMap() {
		if slotObj.IsEquip() {
			bodyposinfoList = append(bodyposinfoList, slotObj)
		}
	}
	scWushuangWeaponInfo := pbutil.BuildSCWushuangWeaponInfo(bodyposinfoList)
	pl.SendMsg(scWushuangWeaponInfo)

	// 神甲认主补偿
	now := global.GetGame().GetTimeService().Now()
	buchangTemp := wushuangweapontemplate.GetWushuangWeaponTemplateService().GetWushuangWeaponBuchangTemplate()
	rewTime := buchangTemp.GetRewTime()
	isSameDay, _ := timeutils.IsSameDay(rewTime, now)
	if isSameDay && !wushuangDataManager.IsSendBuchangEmail() {
		//判断身上是否有无双神器
		for _, slotObj := range wushuangDataManager.GetSlotObjectMap() {
			if !slotObj.IsEquip() {
				continue
			}
			if slotObj.GetItemId() == buchangTemp.ItemId {
				sendBuchangEmail(pl, buchangTemp.GetRewItemMap())
				wushuangDataManager.SendBuchangEmail()
				return
			}
		}
		//判断背包是否有无双神器
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		if inventoryManager.NumOfItems(buchangTemp.ItemId) != int32(0) {
			sendBuchangEmail(pl, buchangTemp.GetRewItemMap())
			wushuangDataManager.SendBuchangEmail()
			return
		}
	}

	return
}

func sendBuchangEmail(pl player.Player, rewMap map[int32]int32) {
	mailTitle := fmt.Sprintf(lang.GetLangService().ReadLang(lang.WushuangWeaponBuchangEmailTitle))
	mailContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.WushuangWeaponBuchangEmailContent))
	emaillogic.AddEmail(pl, mailTitle, mailContent, rewMap)
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
