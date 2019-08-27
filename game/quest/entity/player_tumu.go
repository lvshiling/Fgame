package entity

//玩家屠魔次数数据
type PlayerTuMoEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Num        int32 `gorm:"column:num"`
	ExtraNum   int32 `gorm:"column:extraNum"`
	UsedNum    int32 `gorm:"column:usedNum"`
	UsedBuyNum int32 `gorm:"column:usedBuyNum"`
	BuyNum     int32 `gorm:"column:buyNum"`
	LastTime   int64 `gorm:"column:lastTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (p *PlayerTuMoEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerTuMoEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerTuMoEntity) TableName() string {
	return "t_player_tumo"
}
