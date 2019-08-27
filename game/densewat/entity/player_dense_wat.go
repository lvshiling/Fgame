package entity

//玩家金银密窟数据
type PlayerDenseWatEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Num        int32 `gorm:"column:num"`
	EndTime    int64 `gorm:"column:endTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pdm *PlayerDenseWatEntity) GetId() int64 {
	return pdm.Id
}

func (pdm *PlayerDenseWatEntity) GetPlayerId() int64 {
	return pdm.PlayerId
}

func (pdm *PlayerDenseWatEntity) TableName() string {
	return "t_player_dense_wat"
}
