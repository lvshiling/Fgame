package entity

//玩家附加系统觉醒数据
type PlayerAdditionSysAwakeEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	SysType    int32 `gorm:"column:sysType"`
	IsAwake    int32 `gorm:"column:isAwake"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerAdditionSysAwakeEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerAdditionSysAwakeEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerAdditionSysAwakeEntity) TableName() string {
	return "t_player_addition_sys_awake"
}
