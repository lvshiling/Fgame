package entity

//玩家背包数据
type PlayerInventoryEntity struct {
	Id            int64 `gorm:"primary_key;column:id"`
	PlayerId      int64 `gorm:"column:playerId"`
	SlotNum       int32 `gorm:"column:slotNum"`
	DepotNum      int32 `gorm:"column:depotNum"`
	MiBaoDepotNum int32 `gorm:"column:miBaoDepotNum"`
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (e *PlayerInventoryEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerInventoryEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerInventoryEntity) TableName() string {
	return "t_player_inventory"
}
