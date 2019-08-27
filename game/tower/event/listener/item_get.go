package listener

import (
	"fgame/fgame/core/event"
	dropeventtypes "fgame/fgame/game/drop/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

//打宝塔物品获得统计
func dropItemGet(target event.EventTarget, data event.EventData) (err error) {
	dropItem := target.(scene.DropItem)
	pl := data.(player.Player)
	//发送事件
	s := pl.GetScene()
	itemId := dropItem.GetItemId()
	num := dropItem.GetItemNum()

	if !s.MapTemplate().IsTower() {
		return
	}

	pl.CountTowerItemMap(itemId, num)
	return
}

func init() {
	gameevent.AddEventListener(dropeventtypes.EventTypeDropItemGet, event.EventListenerFunc(dropItemGet))
}
