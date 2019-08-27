package entity

//玩家非进阶身法数据
type PlayerShenfaOtherEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Typ        int32 `gorm:"column:typ"`
	ShenFaId   int32 `gorm:"column:shenFaId"`
	Level      int32 `gorm:"column:level"`
	UpNum      int32 `gorm:"column:upNum"`
	UpPro      int32 `gorm:"column:upPro"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerShenfaOtherEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerShenfaOtherEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerShenfaOtherEntity) TableName() string {
	return "t_player_shenfa_other"
}
