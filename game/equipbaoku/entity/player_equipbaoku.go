package entity

//玩家装备宝库数据
type PlayerEquipBaoKuEntity struct {
	Id                    int64 `gorm:"primary_key;column:id"`
	PlayerId              int64 `gorm:"column:playerId"`
	Typ                   int32 `gorm:"column:typ"`
	LuckyPoints           int32 `gorm:"column:luckyPoints"`
	AttendPoints          int32 `gorm:"column:attendPoints"`
	TotalAttendTimes      int32 `gorm:"column:TotalAttendTimes"`
	LastSystemRefreshTime int64 `gorm:"column:lastSystemRefreshTime"`
	UpdateTime            int64 `gorm:"column:updateTime"`
	CreateTime            int64 `gorm:"column:createTime"`
	DeleteTime            int64 `gorm:"column:deleteTime"`
}

func (e *PlayerEquipBaoKuEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerEquipBaoKuEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerEquipBaoKuEntity) TableName() string {
	return "t_player_equipbaoku"
}
