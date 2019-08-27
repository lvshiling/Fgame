package entity

//玩家附加系统通灵数据
type PlayerAdditionSysTongLingEntity struct {
	Id          int64 `gorm:"primary_key;column:id"`
	PlayerId    int64 `gorm:"column:playerId"`
	SysType     int32 `gorm:"column:sysType"`
	TongLingLev int32 `gorm:"column:tongLingLev"`
	TongLingNum int32 `gorm:"column:tongLingNum"`
	TongLingPro int32 `gorm:"column:tongLingPro"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (pmoe *PlayerAdditionSysTongLingEntity) GetId() int64 {
	return pmoe.Id
}

func (pmoe *PlayerAdditionSysTongLingEntity) GetPlayerId() int64 {
	return pmoe.PlayerId
}

func (pmoe *PlayerAdditionSysTongLingEntity) TableName() string {
	return "t_player_addition_sys_tongling"
}
