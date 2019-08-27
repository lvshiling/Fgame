package entity

type ChuangShiShenWangVoteEntity struct {
	Id         int64 `gorm:"column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	PlayerId   int64 `gorm:"column:playerId"`
	SupportId  int64 `gorm:"column:supportId"`
	Status     int32 `gorm:"column:status"` //投票状态
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *ChuangShiShenWangVoteEntity) GetId() int64 {
	return e.Id
}

func (e *ChuangShiShenWangVoteEntity) TableName() string {
	return "t_chuangshi_shenwang_vote"
}
