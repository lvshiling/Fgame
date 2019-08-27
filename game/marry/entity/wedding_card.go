package entity

//喜帖数据
type MarryWedCardEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	ServerId   int32  `gorm:"column:serverId"`
	PlayerId   int64  `gorm:"column:playerId"`
	SpouseId   int64  `gorm:"column:spouseId"`
	PlayerName string `gorm:"column:playerName"`
	SpouseName string `gorm:"column:spouseName"`
	HoldTime   string `gorm:"column:holdTime"`
	OutOfTime  int64  `gorm:"column:outOfTime"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (p *MarryWedCardEntity) GetId() int64 {
	return p.Id
}

func (p *MarryWedCardEntity) TableName() string {
	return "t_wedding_card"
}
