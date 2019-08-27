package entity

//玩家神域数据
type PlayerShenYuEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	KeyNum     int32  `gorm:"column:keyNum"`
	Round      int32  `gorm:"column:round"`
	Exp        int64  `gorm:"column:exp"`
	ItemInfo   string `gorm:"column:itemInfo"`
	EndTime    int64  `gorm:"column:endTime"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *PlayerShenYuEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerShenYuEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerShenYuEntity) TableName() string {
	return "t_player_shenyu"
}
