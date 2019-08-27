package entity

//资源找回数据
type PlayerFoundBackEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	ResType    int32 `gorm:"column:resType"`
	ResLevel   int32 `gorm:"column:resLevel"`
	Status     int32 `gorm:"column:isReceive"`
	FoundTimes int32 `gorm:"column:foundTimes"`
	Group      int32 `gorm:"column:group"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerFoundBackEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFoundBackEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFoundBackEntity) TableName() string {
	return "t_player_found_back"
}
