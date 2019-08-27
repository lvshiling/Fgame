package entity

//玩家天劫塔数据
type PlayerTianJieTaEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	PlayerId       int64  `gorm:"column:playerId"`
	PlayerName     string `gorm:"column:playerName"`
	Level          int32  `gorm:"column:level"`
	UsedTime       int64  `gorm:"column:usedTime"`
	IsCheckReissue int32  `gorm:"column:isCheckReissue"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (ptjt *PlayerTianJieTaEntity) GetId() int64 {
	return ptjt.Id
}

func (ptjt *PlayerTianJieTaEntity) GetPlayerId() int64 {
	return ptjt.PlayerId
}

func (ptjt *PlayerTianJieTaEntity) TableName() string {
	return "t_player_tianjieta"
}
