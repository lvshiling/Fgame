package entity

//回收
type TradeRecycleEntity struct {
	Id                int64 `gorm:"primary_key;column:id"`
	ServerId          int32 `gorm:"column:serverId"`
	RecycleGold       int64 `gorm:"column:recycleGold"`
	RecycleTime       int64 `gorm:"column:recycleTime"`
	CustomRecycleGold int64 `gorm:"column:customRecycleGold"`
	UpdateTime        int64 `gorm:"column:updateTime"`
	CreateTime        int64 `gorm:"column:createTime"`
	DeleteTime        int64 `gorm:"column:deleteTime"`
}

func (e *TradeRecycleEntity) GetId() int64 {
	return e.Id
}

func (e *TradeRecycleEntity) TableName() string {
	return "t_trade_recycle"
}
