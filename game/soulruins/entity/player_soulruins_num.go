package entity

//玩家帝陵遗迹挑战次数
type PlayerSoulRuinsNumEntity struct {
	Id          int64 `gorm:"primary_key;column:id"`
	PlayerId    int64 `gorm:"column:playerId"`
	Num         int32 `gorm:"column:num"`
	ExtraBuyNum int32 `gorm:"column:extraBuyNum"`
	RewNum      int32 `gorm:"column:rewNum"`
	UsedNum     int32 `gorm:"column:usedNum"`
	UsedBuyNum  int32 `gorm:"column:usedBuyNum"`
	UsedRewNum  int32 `gorm:"column:usedRewNum"`
	BuyNum      int32 `gorm:"column:buyNum"`
	LastTime    int64 `gorm:"column:lastTime"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (p *PlayerSoulRuinsNumEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerSoulRuinsNumEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerSoulRuinsNumEntity) TableName() string {
	return "t_player_soulruins_num"
}
