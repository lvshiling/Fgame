package listener

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/event"
	dropeventtypes "fgame/fgame/game/drop/event/types"
	droplogic "fgame/fgame/game/drop/logic"
	gameevent "fgame/fgame/game/event"
	fourgodtemplate "fgame/fgame/game/fourgod/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	propertylogic "fgame/fgame/game/property/logic"
	"fgame/fgame/game/scene/scene"
)

//物品获得
func dropItemGet(target event.EventTarget, data event.EventData) (err error) {
	dropItem := target.(scene.DropItem)
	pl := data.(player.Player)
	//发送事件
	s := pl.GetScene()
	itemId := dropItem.GetItemId()
	num := dropItem.GetItemNum()
	level := dropItem.GetLevel()
	bind := dropItem.GetBindType()
	upstar := dropItem.GetUpstar()
	attrList := dropItem.GetAttrList()

	itemTempalte := item.GetItemService().GetItem(int(itemId))
	if itemTempalte.GetItemSubType() == itemtypes.ItemAutoUseResSubTypeKey {
		keyNum := pl.GetFourGodKey()
		maxNum := fourgodtemplate.GetFourGodTemplateService().GetFourGodConstTemplate().KeyMax
		if keyNum >= maxNum {
			playerlogic.SendSystemMessage(pl, lang.FourGodPickUpKeyReachLimit)
			return
		}
	}

	goldLog := commonlog.GoldLogReasonMonsterKilled
	goldReasonText := goldLog.String()
	silverLog := commonlog.SilverLogReasonMonsterKilled
	silverReasonText := silverLog.String()
	inventoryLog := commonlog.InventoryLogReasonMonsterKilled
	inventoryReasonText := inventoryLog.String()
	levelLog := commonlog.LevelLogReasonMonsterKilled
	levelReasonText := levelLog.String()

	flag, err := droplogic.AddItemWithProperty(pl, itemId, num, level, upstar, attrList, bind, goldLog, goldReasonText, silverLog, silverReasonText, inventoryLog, inventoryReasonText, levelLog, levelReasonText)
	if err != nil {
		return
	}

	if !flag {
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//同步属性
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	itemData := droplogic.SceneDropConvertToDropItemData(dropItem)
	s.OnPlayerGetItem(pl, itemData)
	dropItem.GetScene().RemoveSceneObject(dropItem, false)
	return
}

func init() {
	gameevent.AddEventListener(dropeventtypes.EventTypeDropItemGet, event.EventListenerFunc(dropItemGet))
}
