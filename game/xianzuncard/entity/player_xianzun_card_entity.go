package entity

// 玩家仙尊特权卡数据
type PlayerXianZunCardEntity struct {
	Id           int64 `gorm:"primary_key;column:id"`
	PlayerId     int64 `gorm:"column:playerId"`
	Typ          int32 `gorm:"column:typ"`
	IsActivite   int32 `gorm:"column:IsActivite"`
	IsReceive    int32 `gorm:"column:isReceive"`
	ActiviteTime int64 `gorm:"column:activiteTime"`
	ReceiveTime  int64 `gorm:"column:receiveTime"`
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (e *PlayerXianZunCardEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerXianZunCardEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerXianZunCardEntity) TableName() string {
	return "t_player_xianzun_card"
}
