package entity

//玩家定情时装获取标志
type PlayerMarryJiNianSjEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	SjGetFlag  int   `gorm:"column:sjGetFlag"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (p *PlayerMarryJiNianSjEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerMarryJiNianSjEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerMarryJiNianSjEntity) TableName() string {
	return "t_player_marry_jinian_sj"
}
