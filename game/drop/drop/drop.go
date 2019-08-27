package drop

import (
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/common/common"
	dropeventtypes "fgame/fgame/game/drop/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/idutil"
)

//掉落
type dropItem struct {
	*scene.SceneObjectBase
	id            int64
	dropOwnerType scenetypes.DropOwnerType
	itemId        int32
	level         int32
	itemNum       int32
	bindType      itemtypes.ItemBindType
	upstar        int32
	attrList      []int32
	protectTime   int64
	existTime     int64
	ownerId       int64
	createTime    int64
	autoTime      int64
}

func (di *dropItem) GetOwnerType() scenetypes.DropOwnerType {
	return di.dropOwnerType
}

func (di *dropItem) GetId() int64 {
	return di.id
}

func (di *dropItem) GetItemId() int32 {
	return di.itemId
}

func (di *dropItem) GetLevel() int32 {
	return di.level
}

func (di *dropItem) GetUpstar() int32 {
	return di.upstar
}

func (di *dropItem) GetAttrList() []int32 {
	return di.attrList
}

func (di *dropItem) GetItemNum() int32 {
	return di.itemNum
}

func (di *dropItem) GetBindType() itemtypes.ItemBindType {
	return di.bindType
}

func (di *dropItem) GetOwnerId() int64 {
	return di.ownerId
}

const (
	autoTime = int64(5 * common.SECOND)
)

func (di *dropItem) Tick() {
	//定时
	now := global.GetGame().GetTimeService().Now()
	elapseTime := now - di.createTime
	if elapseTime >= di.existTime {
		gameevent.Emit(dropeventtypes.EventTypeDropItemRemove, di, nil)
		return
	}
	if di.ownerId != 0 {
		//过了保护时间
		if elapseTime >= di.protectTime {
			di.ownerId = 0
			//自动获得
			gameevent.Emit(dropeventtypes.EventTypeDropItemOwnerChanged, di, nil)

		}
	}

	if di.GetScene().MapTemplate().IsFuBen() {
		if elapseTime >= di.autoTime {
			//自动获得
			gameevent.Emit(dropeventtypes.EventTypeDropItemAuto, di, nil)
			di.autoTime += autoTime
		}
	}

}

func CreateDropItem(ownerType scenetypes.DropOwnerType, ownerId int64, itemId, level int32, itemNum int32, bindType itemtypes.ItemBindType, pos coretypes.Position, protectTime int32, existTime int32) scene.DropItem {
	emptyAttrList := []int32{}
	return CreateDropItemPropertyData(ownerType, ownerId, itemId, level, 0, emptyAttrList, itemNum, bindType, pos, protectTime, existTime)
}

//TODO 添加保护时间
func CreateDropItemPropertyData(ownerType scenetypes.DropOwnerType, ownerId int64, itemId, level, upstar int32, attrList []int32, itemNum int32, bindType itemtypes.ItemBindType, pos coretypes.Position, protectTime int32, existTime int32) scene.DropItem {
	di := &dropItem{}
	id, _ := idutil.GetId()
	di.id = id
	angel := float64(0.0)
	di.dropOwnerType = ownerType
	di.itemId = itemId
	di.ownerId = ownerId
	di.itemNum = itemNum
	di.level = level
	di.upstar = upstar
	di.attrList = attrList
	di.bindType = bindType
	di.protectTime = int64(protectTime) * int64(common.SECOND)
	di.existTime = int64(existTime) * int64(common.SECOND)
	di.SceneObjectBase = scene.NewSceneObjectBase(di, pos, angel, scenetypes.BiologyTypeItem)
	now := global.GetGame().GetTimeService().Now()
	di.createTime = now
	di.autoTime = autoTime

	return di
}
