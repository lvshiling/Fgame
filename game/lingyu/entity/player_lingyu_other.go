package entity

//玩家非进阶领域数据
type PlayerLingyuOtherEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Typ        int32 `gorm:"column:typ"`
	LingYuId   int32 `gorm:"column:lingYuId"`
	Level      int32 `gorm:"column:level"`
	UpNum      int32 `gorm:"column:upNum"`
	UpPro      int32 `gorm:"column:upPro"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerLingyuOtherEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerLingyuOtherEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerLingyuOtherEntity) TableName() string {
	return "t_player_lingyu_other"
}
