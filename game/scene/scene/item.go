package scene

import (
	itemtypes "fgame/fgame/game/item/types"
	scenetypes "fgame/fgame/game/scene/types"
)

//掉落物品
type DropItem interface {
	SceneObject
	//获取拥有者类型
	GetOwnerType() scenetypes.DropOwnerType
	//物品id
	GetItemId() int32
	//物品数量
	GetItemNum() int32
	//拥有者id
	GetOwnerId() int64
	//获取等级
	GetLevel() int32
	//获取金装强化星级
	GetUpstar() int32
	//获取金装附件属性
	GetAttrList() []int32
	//获取绑定类型
	GetBindType() itemtypes.ItemBindType
	Tick()
}
