package entity

//霸主
type ArenapvpBaZhuEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	RaceNumber     int32  `gorm:"column:raceNumber"`
	Platform       int32  `gorm:"column:platform"`
	ServerId       int32  `gorm:"column:serverId"`
	PlayerPlatform int32  `gorm:"column:playerPlatform"` //玩家所在平台
	PlayerServerId int32  `gorm:"column:playerServerId"` //玩家所在服
	PlayerId       int64  `gorm:"column:playerId"`       //玩家di
	PlayerName     string `gorm:"column:playerName"`     //玩家姓名
	Role           int32  `gorm:"column:role"`
	Sex            int32  `gorm:"column:sex"`
	WingId         int32  `gorm:"column:wingId"`
	WeaponId       int32  `gorm:"column:weaponId"`
	FashionId      int32  `gorm:"column:fashionId"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (se *ArenapvpBaZhuEntity) GetId() int64 {
	return se.Id
}

func (se *ArenapvpBaZhuEntity) TableName() string {
	return "t_arenapvp_bazhu"
}
