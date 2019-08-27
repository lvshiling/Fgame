package entity

type PlayerChuangShiVoteEntity struct {
	Id           int64 `gorm:"column:id"`
	PlayerId     int64 `gorm:"column:playerId"`
	Status       int32 `gorm:"column:status"`       //投票状态：
	LastVoteTime int64 `gorm:"column:lastVoteTime"` //上次投票时间
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (e *PlayerChuangShiVoteEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerChuangShiVoteEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerChuangShiVoteEntity) TableName() string {
	return "t_player_chuangshi_vote"
}
