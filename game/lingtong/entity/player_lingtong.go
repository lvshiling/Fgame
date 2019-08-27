package entity

//玩家灵童信息数据
type PlayerLingTongEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	LingTongId int32 `gorm:"column:lingTongId"`
	Level      int32 `gorm:"column:level"`
	BasePower  int64 `gorm:"column:basePower"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerLingTongEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerLingTongEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerLingTongEntity) TableName() string {
	return "t_player_lingtong"
}
