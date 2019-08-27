package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/tower/tower"
)

//物品掉落
func battleObjectDrop(target event.EventTarget, data event.EventData) (err error) {
	n, ok := target.(scene.NPC)
	if !ok {
		return
	}
	dropItem, ok := data.(scene.DropItem)
	if !ok {
		return
	}
	s := n.GetScene()

	if !s.MapTemplate().IsTower() {
		return
	}

	bilologyType := n.GetBiologyTemplate().GetBiologyScriptType()
	if bilologyType != scenetypes.BiologyScriptTypeTowerBoss && bilologyType != scenetypes.BiologyScriptTypeTowerMonster {
		return
	}

	p := s.GetPlayer(dropItem.GetOwnerId())
	if p == nil {
		return
	}

	itemId := dropItem.GetItemId()
	num := dropItem.GetItemNum()
	itemTemp := item.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {
		return
	}

	if itemTemp.GetQualityType() < itemtypes.ItemQualityTypePurple {
		return
	}

	sd := s.SceneDelegate().(tower.TowerSceneData)
	biologyName := n.GetBiologyTemplate().Name
	sd.AddLog(p.GetName(), biologyName, itemId, num)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDrop, event.EventListenerFunc(battleObjectDrop))
}
