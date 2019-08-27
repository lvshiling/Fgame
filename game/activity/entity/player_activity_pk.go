package entity

type PlayerActivityPkEntity struct {
	Id             int64 `gorm:"primary_key;column:id"` //Id
	PlayerId       int64 `gorm:"column:playerId"`       //玩家Id
	ActivityType   int32 `gorm:"column:activityType"`   //活动类型
	KilledNum      int32 `gorm:"column:killedNum"`      //被杀数
	LastKilledTime int64 `gorm:"column:lastKilledTime"` //上次被杀时间
	UpdateTime     int64 `gorm:"column:updateTime"`     //更新时间
	CreateTime     int64 `gorm:"column:createTime"`     //创建时间
	DeleteTime     int64 `gorm:"column:deleteTime"`     //删除时间

}

func (e *PlayerActivityPkEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerActivityPkEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerActivityPkEntity) TableName() string {
	return "t_player_activity_pk"
}
