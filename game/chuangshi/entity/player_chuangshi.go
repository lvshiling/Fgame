package entity

type PlayerChuangShiEntity struct {
	Id            int64 `gorm:"column:id"`
	PlayerId      int64 `gorm:"column:playerId"`
	CampType      int32 `gorm:"column:campType"`
	Pos           int32 `gorm:"column:pos"`
	Jifen         int64 `gorm:"column:jifen"`
	Diamonds      int64 `gorm:"column:diamonds"`
	WeiWang       int64 `gorm:"column:weiWang"`
	LastMyPayTime int64 `gorm:"column:lastMyPayTime"`
	JoinCampTime  int64 `gorm:"column:joinCampTime"`
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (e *PlayerChuangShiEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerChuangShiEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerChuangShiEntity) TableName() string {
	return "t_player_chuangshi"
}
