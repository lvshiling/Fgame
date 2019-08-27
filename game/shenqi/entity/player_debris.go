package entity

//玩家碎片槽位数据
type PlayerShenQiDebrisEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	ShenQiType int32 `gorm:"column:shenQiType"`
	SlotId     int32 `gorm:"column:slotId"`
	Level      int32 `gorm:"column:level"`
	UpNum      int32 `gorm:"column:upNum"`
	UpPro      int32 `gorm:"column:upPro"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerShenQiDebrisEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerShenQiDebrisEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerShenQiDebrisEntity) TableName() string {
	return "t_player_shenqi_debris"
}
