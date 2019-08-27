package entity

type ChuangShiShenWangVoteEntity struct {
	Id             int64 `gorm:"column:id"`
	Platform       int32 `gorm:"column:platform"`
	ServerId       int32 `gorm:"column:serverId"`
	CampType       int32 `gorm:"column:campType"` //阵营
	PlayerServerId int32 `gorm:"column:playerServerId"`
	PlayerId       int64 `gorm:"column:playerId"`
	TicketNum      int32 `gorm:"column:ticketNum"` //
	UpdateTime     int64 `gorm:"column:updateTime"`
	CreateTime     int64 `gorm:"column:createTime"`
	DeleteTime     int64 `gorm:"column:deleteTime"`
}

func (e *ChuangShiShenWangVoteEntity) GetId() int64 {
	return e.Id
}

func (e *ChuangShiShenWangVoteEntity) TableName() string {
	return "t_chuangshi_shenwang_vote"
}
