package entity

//玩家命盘祭炼数据
type PlayerMingGeRefinedEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	SubType    int32 `gorm:"column:subType"`
	Number     int32 `gorm:"column:number"`
	Star       int32 `gorm:"column:star"`
	RefinedNum int32 `gorm:"column:refinedNum"`
	RefinedPro int32 `gorm:"column:refinedPro"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pmne *PlayerMingGeRefinedEntity) GetId() int64 {
	return pmne.Id
}

func (pmne *PlayerMingGeRefinedEntity) GetPlayerId() int64 {
	return pmne.PlayerId
}

func (pmoe *PlayerMingGeRefinedEntity) TableName() string {
	return "t_player_mingge_pan_refined"
}
