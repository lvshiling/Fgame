package entity

//玩家婚姻数据
type PlayerMarryEntity struct {
	Id                 int64  `gorm:"primary_key;column:id"`
	PlayerId           int64  `gorm:"column:playerId"`
	SpouseId           int64  `gorm:"column:spouseId"`
	SpouseName         string `gorm:"column:spouseName"`
	Status             int32  `gorm:"column:status"`
	Ring               int32  `gorm:"column:ring"`
	RingLevel          int32  `gorm:"column:ringLevel"`
	RingNum            int32  `gorm:"column:ringNum"`
	RingExp            int32  `gorm:"column:ringExp"`
	TreeLevel          int32  `gorm:"column:treeLevel"`
	TreeNum            int32  `gorm:"column:treeNum"`
	TreeExp            int32  `gorm:"column:treeExp"`
	IsProposal         int32  `gorm:"column:isProposal"`
	WedStatus          int32  `gorm:"column:wedStatus"`
	DevelopExp         int32  `gorm:"column:developExp"`
	DevelopLevel       int32  `gorm:"column:developLevel"`
	CoupleDevelopLevel int32  `gorm:"column:coupleDevelopLevel"`
	UpdateTime         int64  `gorm:"column:updateTime"`
	CreateTime         int64  `gorm:"column:createTime"`
	DeleteTime         int64  `gorm:"column:deleteTime"`
	MarryCount         int32  `gorm:"column:marryCount"`
}

func (p *PlayerMarryEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerMarryEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerMarryEntity) TableName() string {
	return "t_player_marry"
}
