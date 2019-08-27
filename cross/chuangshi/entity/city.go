package entity

//创世城
type ChuangShiCityEntity struct {
	Id           int64 `gorm:"primary_key;column:id"`
	Platform     int32 `gorm:"column:platform"`
	ServerId     int32 `gorm:"column:serverId"`
	CampType     int32 `gorm:"column:campType"`     //阵营
	OriginalCamp int32 `gorm:"column:originalCamp"` //初始阵营
	Typ          int32 `gorm:"column:typ"`          //城市类型
	Index        int32 `gorm:"column:index"`        //
	OwnerId      int64 `gorm:"column:ownerId"`      //所属玩家id
	Jifen        int64 `gorm:"column:jifen"`        //积分
	Diamonds     int64 `gorm:"column:diamonds"`     //钻石
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (e *ChuangShiCityEntity) GetId() int64 {
	return e.Id
}

func (e *ChuangShiCityEntity) TableName() string {
	return "t_chuangshi_city"
}
