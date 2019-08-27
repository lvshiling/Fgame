package entity

//玩家活跃度跨5点记录数据
type PlayerLivenessCrossFiveEntity struct {
	Id            int64 `gorm:"primary_key;column:id"`
	PlayerId      int64 `gorm:"column:playerId"`
	CrossFiveTime int64 `gorm:"column:crossDayTime"`
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (p *PlayerLivenessCrossFiveEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerLivenessCrossFiveEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerLivenessCrossFiveEntity) TableName() string {
	return "t_player_liveness_cross_five"
}
