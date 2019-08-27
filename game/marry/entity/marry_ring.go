package entity

//求婚婚戒数据
type MarryRingEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	ServerId     int32  `gorm:"column:serverId"`
	PlayerId     int64  `gorm:"column:playerId"`
	PeerId       int64  `gorm:"column:peerId"`
	PeerName     string `gorm:"column:peerName"`
	Ring         int32  `gorm:"column:ring"`
	Status       int32  `gorm:"column:status"`
	ProposalTime int64  `gorm:"column:proposalTime"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (p *MarryRingEntity) GetId() int64 {
	return p.Id
}

func (p *MarryRingEntity) TableName() string {
	return "t_marry_ring"
}
