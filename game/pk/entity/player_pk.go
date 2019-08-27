package entity

//玩家pk数据
type PlayerPkEntity struct {
	Id           int64 `gorm:"primary_key;column:id"`
	PlayerId     int64 `gorm:"column:playerId"`
	PkValue      int32 `gorm:"column:pkValue"`
	KillNum      int32 `gorm:"column:killNum"`
	LastKillTime int64 `gorm:"column:lastKillTime"`
	OnlineTime   int64 `gorm:"column:onlineTime"`
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (p *PlayerPkEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerPkEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerPkEntity) TableName() string {
	return "t_player_pk"
}
