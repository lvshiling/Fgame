package entity

//玩家灵池抢夺记录数据
type PlayerOneArenaRecordEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Level      int32 `gorm:"column:level"`
	Pos        int32 `gorm:"column:pos"`
	RobTime    int64 `gorm:"column:robTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pme *PlayerOneArenaRecordEntity) GetId() int64 {
	return pme.Id
}

func (pme *PlayerOneArenaRecordEntity) GetPlayerId() int64 {
	return pme.PlayerId
}

func (pme *PlayerOneArenaRecordEntity) TableName() string {
	return "t_player_onearena_record"
}
