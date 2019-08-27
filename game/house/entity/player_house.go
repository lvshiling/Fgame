package entity

//玩家房子数据
type PlayerHouseEntity struct {
	Id                int64 `gorm:"primary_key;column:id"`
	PlayerId          int64 `gorm:"column:playerId"`
	HouseIndex        int32 `gorm:"column:houseIndex"`
	HouseType         int32 `gorm:"column:houseType"`
	Level             int32 `gorm:"column:level"`
	MaxLevel          int32 `gorm:"column:maxLevel"`
	DayTimes          int32 `gorm:"column:dayTimes"`
	IsBroken          int32 `gorm:"column:isBroken"`
	LastBrokenTime    int64 `gorm:"column:lastBrokenTime"`
	IsRent            int32 `gorm:"column:isRent"`
	RefreshUpdateTime int64 `gorm:"column:rentUpdateTime"`
	UpdateTime        int64 `gorm:"column:updateTime"`
	CreateTime        int64 `gorm:"column:createTime"`
	DeleteTime        int64 `gorm:"column:deleteTime"`
}

func (e *PlayerHouseEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerHouseEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerHouseEntity) TableName() string {
	return "t_player_house"
}
