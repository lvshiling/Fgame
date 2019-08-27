package entity

//玩家元宝送不停
type PlayerSongBuTingEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	IsReceive  int32 `gorm:"column:isReceive"`
	Times      int32 `gorm:"column:times"`
	LastTime   int64 `gorm:"column:lastTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (psde *PlayerSongBuTingEntity) GetId() int64 {
	return psde.Id
}

func (psde *PlayerSongBuTingEntity) GetPlayerId() int64 {
	return psde.PlayerId
}

func (psde *PlayerSongBuTingEntity) TableName() string {
	return "t_player_yuanbao_songbuting"
}
