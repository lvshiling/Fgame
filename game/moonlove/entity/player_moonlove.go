package entity

type PlayerMoonloveEntity struct {
	Id              int64 `gorm:"primary_key;column:id"`  //Id
	PlayerId        int64 `gorm:"column:playerId"`        //玩家Id
	CharmNum        int32 `gorm:"column:charmNum"`        //魅力值
	GenerousNum     int32 `gorm:"column:generousNum"`     //豪气值
	PreActivityTime int64 `gorm:"column:preActivityTime"` //上次活动时间
	UpdateTime      int64 `gorm:"column:updateTime"`      //更新时间
	CreateTime      int64 `gorm:"column:createTime"`      //创建时间
	DeleteTime      int64 `gorm:"column:deleteTime"`      //删除时间

}

func (e *PlayerMoonloveEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerMoonloveEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerMoonloveEntity) TableName() string {
	return "t_player_moonlove"
}
