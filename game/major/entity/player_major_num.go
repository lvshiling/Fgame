package entity

//玩家非进阶坐骑数据
type PlayerMajorNumEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Num        int32 `gorm:"column:num"`
	LastTime   int64 `gorm:"column:lastTime"`
	MajorType  int32 `gorm:"column:majorType"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pmne *PlayerMajorNumEntity) GetId() int64 {
	return pmne.Id
}

func (pmne *PlayerMajorNumEntity) GetPlayerId() int64 {
	return pmne.PlayerId
}

func (pmoe *PlayerMajorNumEntity) TableName() string {
	return "t_player_major_num"
}
