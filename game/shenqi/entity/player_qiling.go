package entity

//玩家器灵槽位数据
type PlayerShenQiQiLingEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	ShenQiType int32 `gorm:"column:shenQiType"`
	QiLingType int32 `gorm:"column:qiLingType"`
	SlotId     int32 `gorm:"column:slotId"`
	Level      int32 `gorm:"column:level"`
	UpNum      int32 `gorm:"column:upNum"`
	UpPro      int32 `gorm:"column:upPro"`
	ItemId     int32 `gorm:"column:itemId"`
	BindType   int32 `gorm:"column:bindType"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerShenQiQiLingEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerShenQiQiLingEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerShenQiQiLingEntity) TableName() string {
	return "t_player_shenqi_qiling"
}
