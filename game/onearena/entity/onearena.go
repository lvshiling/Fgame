package entity

//灵池数据
type OneArenaEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	ServerId   int32  `gorm:"column:serverId"`
	Level      int32  `gorm:"column:level"`
	Pos        int32  `gorm:"column:pos"`
	OwnerId    int64  `gorm:"column:ownerId"`
	OwnerName  string `gorm:"column:ownerName"`
	LastTime   int64  `gorm:"column:lastTime"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (oae *OneArenaEntity) GetId() int64 {
	return oae.Id
}

func (oae *OneArenaEntity) TableName() string {
	return "t_onearena"
}
