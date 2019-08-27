package entity

//玩家会员数据
type PlayerHuiYuanEntity struct {
	Id                     int64 `gorm:"primary_key;column:id"`
	PlayerId               int64 `gorm:"column:playerId"`
	Type                   int32 `gorm:"column:type"`
	Level                  int32 `gorm:"column:level"`
	LastReceiveTime        int64 `gorm:"column:lastReceiveTime"`
	LastInterimReceiveTime int64 `gorm:"column:lastInterimReceiveTime"`
	PlusBuyTime            int64 `gorm:"column:plusBuyTime"`
	InterimBuyTime         int64 `gorm:"column:interimBuyTime"`
	ExpireTime             int64 `gorm:"column:expireTime"`
	UpdateTime             int64 `gorm:"column:updateTime"`
	CreateTime             int64 `gorm:"column:createTime"`
	DeleteTime             int64 `gorm:"column:deleteTime"`
}

func (e *PlayerHuiYuanEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerHuiYuanEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerHuiYuanEntity) TableName() string {
	return "t_player_huiyuan"
}
