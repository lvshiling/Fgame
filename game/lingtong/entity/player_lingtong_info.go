package entity

//玩家激活灵童信息数据
type PlayerLingTongInfoEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	PlayerId     int64  `gorm:"column:playerId"`
	LingTongId   int32  `gorm:"column:lingTongId"`
	LingTongName string `gorm:"column:lingTongName"`
	UpgradeLevel int32  `gorm:"column:upgradeLevel"`
	UpgradeNum   int32  `gorm:"column:upgradeNum"`
	UpgradePro   int32  `gorm:"column:upgradePro"`
	PeiYangLevel int32  `gorm:"column:peiYangLevel"`
	PeiYangNum   int32  `gorm:"column:peiYangNum"`
	PeiYangPro   int32  `gorm:"column:peiYangPro"`
	StarLevel    int32  `gorm:"column:starLevel"`
	StarNum      int32  `gorm:"column:starNum"`
	StarPro      int32  `gorm:"column:starPro"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *PlayerLingTongInfoEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerLingTongInfoEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerLingTongInfoEntity) TableName() string {
	return "t_player_lingtong_info"
}
