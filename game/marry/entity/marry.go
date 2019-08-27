package entity

//婚姻数据
type MarryEntity struct {
	Id                 int64  `gorm:"primary_key;column:id"`
	ServerId           int32  `gorm:"column:serverId"`
	PlayerId           int64  `gorm:"column:playerId"`
	SpouseId           int64  `gorm:"column:spouseId"`
	PlayerName         string `gorm:"column:playerName"`
	SpouseName         string `gorm:"column:spouseName"`
	PlayerRingLevel    int32  `gorm:"column:playerRingLevel"`
	SpouseRingLevel    int32  `gorm:"column:spouseRingLevel"`
	Role               int32  `gorm:"column:role"`
	SpouseRole         int32  `gorm:"column:spouseRole"`
	Sex                int32  `gorm:"column:sex"`
	SpouseSex          int32  `gorm:"column:spouseSex"`
	Point              int32  `gorm:"column:point"`
	Ring               int32  `gorm:"column:ring"`
	Status             int32  `gorm:"column:status"`
	DevelopLevel       int32  `gorm:"column:developLevel"`
	SpouseDevelopLevel int32  `gorm:"column:spouseDevelopLevel"`
	PlayerSuit         string `gorm:"column:playerSuit"`
	SpouseSuit         string `gorm:"column:spouseSuit"`
	UpdateTime         int64  `gorm:"column:updateTime"`
	CreateTime         int64  `gorm:"column:createTime"`
	DeleteTime         int64  `gorm:"column:deleteTime"`
}

func (p *MarryEntity) GetId() int64 {
	return p.Id
}

func (p *MarryEntity) TableName() string {
	return "t_marry"
}
