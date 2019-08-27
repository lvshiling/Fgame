package entity

//玩家附加系统槽位数据
type PlayerAdditionSysSlotEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	SysType    int32 `gorm:"column:sysType"`
	SlotId     int32 `gorm:"column:slotId"`
	Level      int32 `gorm:"column:level"`
	ShenZhuLev int32 `gorm:"column:shenZhuLev"`
	ShenZhuNum int32 `gorm:"column:shenZhuNum"`
	ShenZhuPro int32 `gorm:"column:shenZhuPro"`
	ItemId     int32 `gorm:"column:itemId"`
	BindType   int32 `gorm:"column:bindType"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerAdditionSysSlotEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerAdditionSysSlotEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerAdditionSysSlotEntity) TableName() string {
	return "t_player_addition_sys_slot"
}
