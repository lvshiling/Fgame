package entity

//玩家阵旗数据
type PlayerZhenQiEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Type       int32 `gorm:"column:type"`
	ZhenQiPos  int32 `gorm:"column:zhenQiPos"`
	Number     int32 `gorm:"column:number"`
	NumberNum  int32 `gorm:"column:numberNum"`
	NumberPro  int32 `gorm:"column:numberPro"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pe *PlayerZhenQiEntity) GetId() int64 {
	return pe.Id
}

func (pe *PlayerZhenQiEntity) GetPlayerId() int64 {
	return pe.PlayerId
}

func (pe *PlayerZhenQiEntity) TableName() string {
	return "t_player_zhenqi"
}
