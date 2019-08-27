package equipbaoku

import (
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
)

//装备宝库日志列表对象
type EquipBaoKuLogObject struct {
	playerName string
	typ        equipbaokutypes.BaoKuType
	itemId     int32
	itemNum    int32
	updateTime int64
	createTime int64
}

func NewEquipBaoKuLogObject() *EquipBaoKuLogObject {
	o := &EquipBaoKuLogObject{}
	return o
}

func (o *EquipBaoKuLogObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *EquipBaoKuLogObject) GetUpdateTime() int64 {
	return o.updateTime
}

func (o *EquipBaoKuLogObject) GetPlayerName() string {
	return o.playerName
}

func (o *EquipBaoKuLogObject) GetItemId() int32 {
	return o.itemId
}

func (o *EquipBaoKuLogObject) GetItemNum() int32 {
	return o.itemNum
}

func (o *EquipBaoKuLogObject) GetBaoKuType() equipbaokutypes.BaoKuType {
	return o.typ
}
