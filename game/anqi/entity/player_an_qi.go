package entity

//玩家暗器数据
type PlayerAnQiEntity struct {
	Id           int64 `gorm:"primary_key;column:id"`
	PlayerId     int64 `gorm:"column:playerId"`
	AdvancedId   int   `gorm:"column:advancedId"`
	AnqiDanLevel int32 `gorm:"column:anqiDanLevel"`
	AnqiDanNum   int32 `gorm:"column:anqiDanNum"`
	AnqiDanPro   int32 `gorm:"column:anqiDanPro"`
	TimesNum     int32 `gorm:"column:timesNum"`
	Bless        int32 `gorm:"column:bless"`
	BlessTime    int64 `gorm:"column:blessTime"`
	Power        int64 `gorm:"column:power"`
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (e *PlayerAnQiEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerAnQiEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerAnQiEntity) TableName() string {
	return "t_player_anqi"
}
