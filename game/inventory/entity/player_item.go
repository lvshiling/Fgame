package entity

//玩家物品数据
type PlayerItemEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	PlayerId     int64  `gorm:"column:playerId"`
	BagType      int32  `gorm:"column:bagType"`
	ItemId       int32  `gorm:"column:itemId"`
	Num          int32  `gorm:"column:num"`
	Index        int32  `grom:"column:index"`
	Level        int32  `grom:"column:level"`
	Used         int32  `gorm:"column:used"`
	IsDepot      int32  `gorm:"column:isDepot"`
	BindType     int32  `gorm:"column:bindType"`
	ItemGetTime  int64  `gorm:"column:itemGetTime"`
	PropertyData string `gorm:"column:porpertyData"`
	LastUseTime  int64  `gorm:"column:lastUseTime"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *PlayerItemEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerItemEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerItemEntity) TableName() string {
	return "t_player_item"
}
