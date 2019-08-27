package entity

//玩家非进阶仙体数据
type PlayerXianTiOtherEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Typ        int32 `gorm:"column:typ"`
	XianTiId   int32 `gorm:"column:xianTiId"`
	Level      int32 `gorm:"column:level"`
	UpNum      int32 `gorm:"column:upNum"`
	UpPro      int32 `gorm:"column:upPro"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pmoe *PlayerXianTiOtherEntity) GetId() int64 {
	return pmoe.Id
}

func (pmoe *PlayerXianTiOtherEntity) GetPlayerId() int64 {
	return pmoe.PlayerId
}

func (pmoe *PlayerXianTiOtherEntity) TableName() string {
	return "t_player_xianti_other"
}
