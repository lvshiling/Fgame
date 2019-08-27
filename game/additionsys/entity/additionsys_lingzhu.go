package entity

//玩家附加系统五行灵珠数据
type PlayerAdditionSysLingZhuEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	SysType    int32 `gorm:"column:sysType"`
	LingZhuId  int32 `gorm:"column:lingZhuId"`
	Level      int32 `gorm:"column:level"`
	Times      int32 `gorm:"column:times"`
	Bless      int64 `gorm:"column:bless"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerAdditionSysLingZhuEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerAdditionSysLingZhuEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerAdditionSysLingZhuEntity) TableName() string {
	return "t_player_addition_sys_lingzhu"
}
