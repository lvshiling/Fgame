package entity

type PlayerChuangShiGuanZhiEntity struct {
	Id              int64 `gorm:"column:id"`
	PlayerId        int64 `gorm:"column:playerId"`
	ReceiveRewLevel int32 `gorm:"column:receiveRewLevel"` //领取的奖励等级
	Level           int32 `gorm:"column:level"`
	Times           int32 `gorm:"column:times"`
	WeiWang         int32 `gorm:"column:weiWang"`
	UpdateTime      int64 `gorm:"column:updateTime"`
	CreateTime      int64 `gorm:"column:createTime"`
	DeleteTime      int64 `gorm:"column:deleteTime"`
}

func (e *PlayerChuangShiGuanZhiEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerChuangShiGuanZhiEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerChuangShiGuanZhiEntity) TableName() string {
	return "t_player_chuangshi_guanzhi"
}
