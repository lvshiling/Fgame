package entity

//玩家淬炼槽位数据
type PlayerShenQiSmeltEntity struct {
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

func (e *PlayerShenQiSmeltEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerShenQiSmeltEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerShenQiSmeltEntity) TableName() string {
	return "t_player_shenqi_smelt"
}
