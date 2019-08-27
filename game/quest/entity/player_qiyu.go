package entity

//玩家奇遇数据
type PlayerQiYuEntity struct {
	Id          int64 `gorm:"primary_key;column:id"`
	PlayerId    int64 `gorm:"column:playerId"`
	QiYuId      int32 `gorm:"column:qiyuId"`
	Level       int32 `gorm:"column:level"`
	Zhuan       int32 `gorm:"column:zhuan"`
	Fei         int32 `gorm:"column:fei"`
	IsFinish    int32 `gorm:"column:isFinish"`
	IsReceive   int32 `gorm:"column:isReceive"`
	IsHadNotice int32 `gorm:"column:isHadNotice"`
	EndTime     int64 `gorm:"column:endTime"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (p *PlayerQiYuEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerQiYuEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerQiYuEntity) TableName() string {
	return "t_player_qiyu"
}
