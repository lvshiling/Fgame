package entity

//婚姻安排数据
type MarryWedEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	ServerId    int32  `gorm:"column:serverId"`
	Period      int32  `gorm:"column:period"`
	Grade       int32  `gorm:"column:grade"`
	HunCheGrade int32  `gorm:"column:hunCheGrade"`
	SugarGrade  int32  `gorm:"column:sugarGrade"`
	Status      int32  `gorm:"column:status"`
	PlayerId    int64  `gorm:"column:playerId"`
	SpouseId    int64  `gorm:"column:spouseId"`
	Name        string `gorm:"column:name"`
	SpouseName  string `gorm:"column:spouseName"`
	HTime       int64  `gorm:"column:hTime"`
	LastTime    int64  `gorm:"column:lastTime"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
	IsFirst     int32  `gorm:"column:isFirst"`
}

func (p *MarryWedEntity) GetId() int64 {
	return p.Id
}

func (p *MarryWedEntity) TableName() string {
	return "t_wedding"
}
