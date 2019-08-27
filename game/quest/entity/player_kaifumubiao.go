package entity

//玩家开服目标数据
type PlayerKaiFuMuBiaoEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	KaiFuDay   int32 `gorm:"column:kaiFuDay"`
	FinishNum  int32 `gorm:"column:finishNum"`
	IsReward   int32 `gorm:"column:isReward"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (p *PlayerKaiFuMuBiaoEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerKaiFuMuBiaoEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerKaiFuMuBiaoEntity) TableName() string {
	return "t_player_kaifumubiao"
}
