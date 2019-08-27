package entity

//玩家神龙数据
type PlayerDragonEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	StageId    int32  `gorm:"column:stageId"`
	ItemInfo   string `gorm:"column:itemInfo"`
	Status     int32  `gorm:"column:status"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (pdm *PlayerDragonEntity) GetId() int64 {
	return pdm.Id
}

func (pdm *PlayerDragonEntity) GetPlayerId() int64 {
	return pdm.PlayerId
}

func (pdm *PlayerDragonEntity) TableName() string {
	return "t_player_dragon"
}
