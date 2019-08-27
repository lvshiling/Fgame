package entity

//回收
type PlayerTradeRecycleEntity struct {
	Id          int64 `gorm:"primary_key;column:id"`
	PlayerId    int64 `gorm:"column:playerId"`
	RecycleGold int64 `gorm:"column:recycleGold"`
	RecycleTime int64 `gorm:"column:recycleTime"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (e *PlayerTradeRecycleEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerTradeRecycleEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerTradeRecycleEntity) TableName() string {
	return "t_player_trade_recycle"
}
