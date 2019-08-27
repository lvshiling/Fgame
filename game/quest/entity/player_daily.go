package entity

//玩家日环数据
type PlayerDailyEntity struct {
	Id            int64 `gorm:"primary_key;column:id"`
	PlayerId      int64 `gorm:"column:playerId"`
	DailyTag      int32 `gorm:"column:dailyTag"`
	SeqId         int32 `gorm:"column:seqId"`
	Times         int32 `gorm:"column:times"`
	LastTime      int64 `gorm:"column:lastTime"`
	CrossFiveTime int64 `gorm:"column:crossDayTime"`
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (p *PlayerDailyEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerDailyEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerDailyEntity) TableName() string {
	return "t_player_daily"
}
