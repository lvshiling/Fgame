package entity

//玩家仙桃大会数据
type PlayerXianTaoEntity struct {
	Id               int64 `gorm:"primary_key;column:id"`
	PlayerId         int64 `gorm:"column:playerId"`
	JuniorPeachCount int32 `gorm:"column:juniorPeachCount"`
	HighPeachCount   int32 `gorm:"column:highPeachCount"`
	RobCount         int32 `gorm:"column:robCount"`
	BeRobCount       int32 `gorm:"column:beRobCount"`
	EndTime          int64 `gorm:"column:endTime"`
	UpdateTime       int64 `gorm:"column:updateTime"`
	CreateTime       int64 `gorm:"column:createTime"`
	DeleteTime       int64 `gorm:"column:deleteTime"`
}

func (p *PlayerXianTaoEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerXianTaoEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerXianTaoEntity) TableName() string {
	return "t_player_xian_tao"
}
