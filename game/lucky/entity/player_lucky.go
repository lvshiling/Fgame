package entity

//玩家幸运符数据
type PlayerLuckyEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Typ        int32 `gorm:"column:typ"`
	SubType    int32 `gorm:"column:subType"`
	ItemId     int32 `gorm:"column:itemId"`
	ExpireTime int64 `gorm:"column:expireTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pwe *PlayerLuckyEntity) GetId() int64 {
	return pwe.Id
}

func (pwe *PlayerLuckyEntity) GetPlayerId() int64 {
	return pwe.PlayerId
}

func (pwe *PlayerLuckyEntity) TableName() string {
	return "t_player_lucky"
}
