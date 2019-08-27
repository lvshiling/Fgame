package entity

type DaLiWanEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Typ        int32 `gorm:"column:typ"`
	StartTime  int64 `gorm:"column:startTime"`
	Duration   int64 `gorm:"column:duration"`
	Expired    int32 `gorm:"column:expired"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *DaLiWanEntity) GetId() int64 {
	return e.Id
}

func (e *DaLiWanEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *DaLiWanEntity) TableName() string {
	return "t_player_daliwan"
}
