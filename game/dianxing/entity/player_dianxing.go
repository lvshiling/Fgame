package entity

//玩家点星数据
type PlayerDianXingEntity struct {
	Id                int64 `gorm:"primary_key;column:id"`
	PlayerId          int64 `gorm:"column:playerId"`
	CurrType          int32 `gorm:"column:currType"`
	CurrLevel         int32 `gorm:"column:currLevel"`
	DianXingTimes     int32 `gorm:"column:dianXingTimes"`
	DianXingBless     int32 `gorm:"column:dianXingBless"`
	DianXingBlessTime int64 `gorm:"column:dianXingBlessTime"`
	XingChenNum       int64 `gorm:"column:xingChenNum"`
	JieFengLev        int32 `gorm:"column:jieFengLev"`
	JieFengTimes      int32 `gorm:"column:jieFengTimes"`
	JieFengBless      int32 `gorm:"column:jieFengBless"`
	Power             int64 `gorm:"column:power"`
	UpdateTime        int64 `gorm:"column:updateTime"`
	CreateTime        int64 `gorm:"column:createTime"`
	DeleteTime        int64 `gorm:"column:deleteTime"`
}

func (e *PlayerDianXingEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerDianXingEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerDianXingEntity) TableName() string {
	return "t_player_dianxing"
}
