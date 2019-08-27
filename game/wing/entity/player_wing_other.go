package entity

//玩家非进阶战翼数据
type PlayerWingOtherEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Typ        int32 `gorm:"column:typ"`
	WingId     int32 `gorm:"column:wingId"`
	Level      int32 `gorm:"column:level"`
	UpNum      int32 `gorm:"column:upNum"`
	UpPro      int32 `gorm:"column:upPro"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pwoe *PlayerWingOtherEntity) GetId() int64 {
	return pwoe.Id
}

func (pwoe *PlayerWingOtherEntity) GetPlayerId() int64 {
	return pwoe.PlayerId
}

func (pwoe *PlayerWingOtherEntity) TableName() string {
	return "t_player_wing_other"
}
