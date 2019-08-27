package entity

//玩家飞升数据
type PlayerFeiShengEntity struct {
	Id            int64 `gorm:"primary_key;column:id"`
	PlayerId      int64 `gorm:"column:playerId"`
	FeiLevel      int32 `gorm:"column:feiLevel"`
	AddRate       int32 `gorm:"column:addRate"`
	GongDeNum     int64 `gorm:"column:gongDeNum"`
	LeftPotential int32 `gorm:"column:leftPotential"`
	TiZhi         int32 `gorm:"column:tiZhi"`
	LiDao         int32 `gorm:"column:liDao"`
	JinGu         int32 `gorm:"column:jinGu"`
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (e *PlayerFeiShengEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFeiShengEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFeiShengEntity) TableName() string {
	return "t_player_fei_sheng"
}
