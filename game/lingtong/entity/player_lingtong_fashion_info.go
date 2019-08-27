package entity

//玩家激活灵童信息数据
type PlayerLingTongFashionInfoEntity struct {
	Id           int64 `gorm:"primary_key;column:id"`
	PlayerId     int64 `gorm:"column:playerId"`
	FashionId    int32 `gorm:"column:fashionId"`
	UpgradeLevel int32 `gorm:"column:upgradeLevel"`
	UpgradeNum   int32 `gorm:"column:upgradeNum"`
	UpgradePro   int32 `gorm:"column:upgradePro"`
	IsExpire     int32 `gorm:"column:isExpire"`
	ActivateTime int64 `gorm:"column:activateTime"`
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (e *PlayerLingTongFashionInfoEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerLingTongFashionInfoEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerLingTongFashionInfoEntity) TableName() string {
	return "t_player_lingtong_fashion_info"
}
