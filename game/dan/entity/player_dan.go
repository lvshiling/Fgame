package entity

//玩家食丹数据
type PlayerDanEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	LevelId    int32  `gorm:"column:levelId"`
	DanInfo    string `gorm:"column:danInfo"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (pdm *PlayerDanEntity) GetId() int64 {
	return pdm.Id
}

func (pdm *PlayerDanEntity) GetPlayerId() int64 {
	return pdm.PlayerId
}

func (pdm *PlayerDanEntity) TableName() string {
	return "t_player_dan"
}
