package entity

//玩家组队副本数据
type PlayerTeamCopyEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	PurPose    int32 `gorm:"column:purpose"`
	Num        int32 `gorm:"column:num"`
	RewTime    int64 `gorm:"column:rewTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerTeamCopyEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerTeamCopyEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerTeamCopyEntity) TableName() string {
	return "t_player_team_copy"
}
