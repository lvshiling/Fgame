package entity

//婚宴档次预定数据
type MarryPreWedEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	ServerId    int32  `gorm:"column:serverId"`
	Period      int32  `gorm:"column:period"`
	PlayerId    int64  `gorm:"column:playerId"`
	PlayerName  string `gorm:"column:playerName"`
	PeerId      int64  `gorm:"column:peerId"`
	Grade       int32  `gorm:"column:grade"`
	HunCheGrade int32  `gorm:"column:hunCheGrade"`
	SugarGrade  int32  `gorm:"column:sugarGrade"`
	Status      int32  `gorm:"column:status"`
	HoldTime    int64  `gorm:"column:holdTime"`
	PreWedTime  int64  `gorm:"column:preWedTime"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (p *MarryPreWedEntity) GetId() int64 {
	return p.Id
}

func (p *MarryPreWedEntity) TableName() string {
	return "t_marry_pre_wed"
}
