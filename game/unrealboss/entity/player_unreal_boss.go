package entity

//玩家幻境BOSS数据
type PlayerUnrealBossEntity struct {
	Id            int64 `gorm:"primary_key;column:id"`
	PlayerId      int64 `gorm:"column:playerId"`
	PiLaoNum      int32 `gorm:"column:pilaoNum"`
	BuyPiLaoNum   int32 `gorm:"column:buyPiLaoNum"`
	BuyPiLaoTimes int32 `gorm:"column:buyPiLaoTimes"`
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (e *PlayerUnrealBossEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerUnrealBossEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerUnrealBossEntity) TableName() string {
	return "t_player_unreal_boss"
}
