package entity

//玩家婚姻数据
type PlayerMarryJiNianEntity struct {
	Id          int64 `gorm:"primary_key;column:id"`
	PlayerId    int64 `gorm:"column:playerId"`
	JiNianType  int32 `gorm:"column:jiNianType"`
	JiNianCount int32 `gorm:"column:jiNianCount"`
	SendFlag    int32 `gorm:"column:sendFlag"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (p *PlayerMarryJiNianEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerMarryJiNianEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerMarryJiNianEntity) TableName() string {
	return "t_player_marry_jinian"
}
