package entity

type YingLingPuEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	TuJianId   int32 `gorm:"column:tuJianId"`
	TuJianType int32 `gorm:"column:tuJianType"`
	Level      int32 `gorm:"column:level"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *YingLingPuEntity) GetId() int64 {
	return e.Id
}

func (e *YingLingPuEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *YingLingPuEntity) TableName() string {
	return "t_player_yinglingpu"
}
