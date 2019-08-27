package entity

type PlayerRingBaoKuEntity struct {
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

func (e *PlayerRingBaoKuEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerRingBaoKuEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerRingBaoKuEntity) TableName() string {
	return "t_player_ring_baoku"
}
