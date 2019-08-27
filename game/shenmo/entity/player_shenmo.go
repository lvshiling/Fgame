package entity

//玩家神魔数据
type PlayerShenMoEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	GongXunNum int32 `gorm:"column:gongXunNum"`
	KillNum    int32 `gorm:"column:killNum"`
	EndTime    int64 `gorm:"column:endTime"`
	RewTime    int64 `gorm:"column:rewTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerShenMoEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerShenMoEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerShenMoEntity) TableName() string {
	return "t_player_shenmo"
}
