package entity

//玩家非进阶灵童升星数据
type PlayerLingTongOtherEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	ClassType  int32 `gorm:"column:classType"`
	Type       int32 `gorm:"column:type"`
	SeqId      int32 `gorm:"column:seqId"`
	Level      int32 `gorm:"column:level"`
	UpNum      int32 `gorm:"column:upNum"`
	UpPro      int32 `gorm:"column:upPro"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerLingTongOtherEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerLingTongOtherEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerLingTongOtherEntity) TableName() string {
	return "t_player_lingtong_other"
}
