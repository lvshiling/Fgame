package entity

//玩家炼丹数据
type PlayerAlchemyEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	KindId     int   `gorm:"column:kindId"`
	Num        int   `gorm:"column:num"`
	StartTime  int64 `gorm:"column:startTime"`
	State      int   `gorm:"column:state"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pam *PlayerAlchemyEntity) GetId() int64 {
	return pam.Id
}

func (pam *PlayerAlchemyEntity) GetPlayerId() int64 {
	return pam.PlayerId
}

func (pam *PlayerAlchemyEntity) TableName() string {
	return "t_player_alchemy"
}
