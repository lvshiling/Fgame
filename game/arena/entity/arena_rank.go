package entity

//3v3排行榜
type ArenaRankEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	ServerId     int32  `gorm:"column:serverId"`
	PlayerId     int64  `gorm:"column:playerId"`
	PlayerName   string `gorm:"column:playerName"`
	CurWinCount  int32  `gorm:"column:curWinCount"`
	WinCount     int32  `gorm:"column:winCount"`
	LastWinCount int32  `gorm:"column:lastWinCount"`
	LastTime     int64  `gorm:"column:lastTime"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (se *ArenaRankEntity) GetId() int64 {
	return se.Id
}

func (se *ArenaRankEntity) TableName() string {
	return "t_arena_rank"
}
