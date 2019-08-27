package entity

//玩家怀孕数据
type PlayerPregnantEntity struct {
	Id           int64 `gorm:"primary_key;column:id"`
	PlayerId     int64 `gorm:"column:playerId"`
	ChaoShengNum int32 `gorm:"column:chaoshengNum"`
	TonicPro     int32 `gorm:"column:tonicPro"`
	PregnantTime int64 `gorm:"column:pregnantTime"`
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (e *PlayerPregnantEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerPregnantEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerPregnantEntity) TableName() string {
	return "t_player_pregnant"
}
