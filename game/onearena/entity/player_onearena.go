package entity

//玩家灵池抢夺数据
type PlayerOneArenaEntity struct {
	Id          int64 `gorm:"primary_key;column:id"`
	PlayerId    int64 `gorm:"column:playerId"`
	Level       int32 `gorm:"column:level"`
	Pos         int32 `gorm:"column:pos"`
	RobTime     int64 `gorm:"column:robTime"`
	KunSilver   int64 `gorm:"column:kunSilver"`
	KunBindGold int64 `gorm:"column:kunBindGold"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (pme *PlayerOneArenaEntity) GetId() int64 {
	return pme.Id
}

func (pme *PlayerOneArenaEntity) GetPlayerId() int64 {
	return pme.PlayerId
}

func (pme *PlayerOneArenaEntity) TableName() string {
	return "t_player_onearena"
}
