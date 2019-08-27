package entity

//玩家血盾数据
type PlayerXueDunEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Blood      int64 `gorm:"column:blood"`
	Number     int32 `gorm:"column:number"`
	Star       int32 `gorm:"column:star"`
	StarNum    int32 `gorm:"column:starNum"`
	StarPro    int32 `gorm:"column:starPro"`
	CulLevel   int32 `gorm:"column:culLevel"`
	CulNum     int32 `gorm:"column:culNum"`
	CulPro     int32 `gorm:"column:culPro"`
	IsActive   int32 `gorm:"column:isActive"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pme *PlayerXueDunEntity) GetId() int64 {
	return pme.Id
}

func (pme *PlayerXueDunEntity) GetPlayerId() int64 {
	return pme.PlayerId
}

func (pme *PlayerXueDunEntity) TableName() string {
	return "t_player_xuedun"
}
