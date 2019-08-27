package logic

import (
	dropeventtypes "fgame/fgame/game/drop/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//自动获取
func AutoGetDropItem(dropItem scene.DropItem) (flag bool, err error) {
	ownerId := dropItem.GetOwnerId()
	//拥有者是空的
	if ownerId == 0 {
		return
	}
	//TODO做组队掉落
	//TODO 玩家不在线 panic
	p := dropItem.GetScene().GetPlayer(ownerId)
	if p == nil {
		return
	}

	return GetDropItem(p, dropItem)
}

//手动获取
func GetDropItem(pl scene.Player, dropItem scene.DropItem) (flag bool, err error) {
	if pl.GetScene() != dropItem.GetScene() {
		return
	}
	//修改在物品里面发送
	gameevent.Emit(dropeventtypes.EventTypeDropItemGet, dropItem, pl)
	return
}

//副本自动捡所有东西
func FuBenGetAllItems(pl scene.Player) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	for _, dropItem := range s.GetAllItems() {
		gameevent.Emit(dropeventtypes.EventTypeDropItemGet, dropItem, pl)
	}
	//修改在物品里面发送
	// gameevent.Emit(dropeventtypes.EventTypeFubenDropItemsGet, pl, nil)
	return

}
