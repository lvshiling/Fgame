package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/inventory/pbutil"
	dropeventtypes "fgame/fgame/game/drop/event/types"
	droplogic "fgame/fgame/game/drop/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//物品获得
func dropItemGet(target event.EventTarget, data event.EventData) (err error) {
	dropItem := target.(scene.DropItem)
	pl := data.(scene.Player)
	s := pl.GetScene()
	//发送物品获取消息
	isDropItemGet := pbutil.BuildISDropItemGet(dropItem.GetItemId(), dropItem.GetItemNum(), dropItem.GetLevel(), dropItem.GetAttrList(), dropItem.GetUpstar())
	pl.SendMsg(isDropItemGet)

	itemData := droplogic.SceneDropConvertToDropItemData(dropItem)
	s.OnPlayerGetItem(pl, itemData)
	dropItem.GetScene().RemoveSceneObject(dropItem, false)
	return
}

func init() {
	gameevent.AddEventListener(dropeventtypes.EventTypeDropItemGet, event.EventListenerFunc(dropItemGet))
}
