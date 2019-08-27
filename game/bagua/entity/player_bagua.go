package entity

//玩家八卦数据
type PlayerBaGuaEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Level      int32 `gorm:"column:level"`
	IsBuChang  int32 `gorm:"column:isBuChang"`
	InviteTime int64 `gorm:"column:inviteTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (t *PlayerBaGuaEntity) GetId() int64 {
	return t.Id
}

func (t *PlayerBaGuaEntity) GetPlayerId() int64 {
	return t.PlayerId
}

func (t *PlayerBaGuaEntity) TableName() string {
	return "t_player_bagua"
}
