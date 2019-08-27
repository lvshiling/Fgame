package entity

type YingLingPuSuiPianEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	TuJianId   int32 `gorm:"column:tuJianId"`
	TuJianType int32 `gorm:"column:tuJianType"`
	SuiPianId  int32 `gorm:"column:suiPianId"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *YingLingPuSuiPianEntity) GetId() int64 {
	return e.Id
}

func (e *YingLingPuSuiPianEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *YingLingPuSuiPianEntity) TableName() string {
	return "t_player_yinglingpu_suipian"
}
