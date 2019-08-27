package entity

//玩家属性数据
type PlayerPropertyEntity struct {
	Id            int64 `gorm:"primary_key;column:id"`
	PlayerId      int64 `gorm:"column:playerId"`
	Level         int32 `gorm:"column:level"`
	Exp           int64 `gorm:"column:exp"`
	Silver        int64 `gorm:"column:silver"`
	Gold          int64 `gorm:"column:gold"`
	BindGold      int64 `gorm:"column:bindGold"`
	Evil          int32 `gorm:"column:evil"`
	CurrentHP     int64 `gorm:"column:currentHP"`
	CurrentTP     int64 `gorm:"column:currentTP"`
	ZhuanSheng    int32 `gorm:"column:zhuanSheng"`
	Power         int64 `gorm:"column:power"`
	Charm         int32 `gorm:"column:charm"`
	GoldYuanLevel int32 `gorm:"column:goldYuanLevel"`
	GoldYuanExp   int64 `gorm:"column:goldYuanExp"`
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (psm *PlayerPropertyEntity) GetId() int64 {
	return psm.Id
}

func (psm *PlayerPropertyEntity) GetPlayerId() int64 {
	return psm.PlayerId
}

func (psm *PlayerPropertyEntity) TableName() string {
	return "t_player_property"
}
