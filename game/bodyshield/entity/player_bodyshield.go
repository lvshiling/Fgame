package entity

//玩家护体盾数据
type PlayerBodyShieldEntity struct {
	Id             int64 `gorm:"primary_key;column:id"`
	PlayerId       int64 `gorm:"column:playerId"`
	AdvancedId     int   `gorm:"column:advancedId"`
	JinjiadanLevel int32 `gorm:"column:jinjiadanLevel"`
	JinjiadanNum   int32 `gorm:"column:jinjiadanNum"`
	JinjiadanPro   int32 `gorm:"column:jinjiadanPro"`
	TimesNum       int32 `gorm:"column:timesNum"`
	Bless          int32 `gorm:"column:bless"`
	BlessTime      int64 `gorm:"column:blessTime"`
	ShieldId       int32 `gorm:"column:shieldId"`
	ShieldNum      int32 `gorm:"column:shieldNum"`
	ShieldPro      int32 `gorm:"column:shieldPro"`
	Power          int64 `gorm:"column:power"`
	SPower         int64 `gorm:"column:spower"`
	UpdateTime     int64 `gorm:"column:updateTime"`
	CreateTime     int64 `gorm:"column:createTime"`
	DeleteTime     int64 `gorm:"column:deleteTime"`
}

func (e *PlayerBodyShieldEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerBodyShieldEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerBodyShieldEntity) TableName() string {
	return "t_player_body_shield"
}
