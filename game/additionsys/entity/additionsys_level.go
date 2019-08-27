package entity

//玩家附加系统升级数据
type PlayerAdditionSysLevelEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	SysType    int32 `gorm:"column:sysType"`
	Level      int32 `gorm:"column:level"`
	UpNum      int32 `gorm:"column:upNum"`
	UpPro      int32 `gorm:"column:upPro"`
	LingLevel  int32 `gorm:"column:lingLevel"`
	LingNum    int32 `gorm:"column:lingNum"`
	LingPro    int32 `gorm:"column:lingPro"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pmoe *PlayerAdditionSysLevelEntity) GetId() int64 {
	return pmoe.Id
}

func (pmoe *PlayerAdditionSysLevelEntity) GetPlayerId() int64 {
	return pmoe.PlayerId
}

func (pmoe *PlayerAdditionSysLevelEntity) TableName() string {
	return "t_player_addition_sys_level"
}
