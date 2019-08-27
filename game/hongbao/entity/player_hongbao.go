package entity

//玩家红包数据
type PlayerHongBaoEntity struct {
	Id          int64 `gorm:"primary_key;column:id"`
	PlayerId    int64 `gorm:"column:playerId"`
	SnatchCount int32 `gorm:"column:snatchCount"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (pjxe *PlayerHongBaoEntity) GetId() int64 {
	return pjxe.Id
}

func (pjxe *PlayerHongBaoEntity) GetPlayerId() int64 {
	return pjxe.PlayerId
}

func (pjxe *PlayerHongBaoEntity) TableName() string {
	return "t_player_hongbao"
}
