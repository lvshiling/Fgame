package entity

//玩家物品使用数据
type PlayerItemUseEntity struct {
	Id          int64 `gorm:"primary_key;column:id"`
	PlayerId    int64 `gorm:"column:playerId"`
	ItemId      int32 `gorm:"column:itemId"`
	TodayTimes  int32 `gorm:"column:todayTimes"`
	TotalTimes  int32 `gorm:"column:totalTimes"`
	LastUseTime int64 `gorm:"column:lastUseTime"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (e *PlayerItemUseEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerItemUseEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerItemUseEntity) TableName() string {
	return "t_player_item_use"
}
