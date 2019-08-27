package entity

//创世成员信息
type ChuangShiMemberEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	Platform       int32  `gorm:"column:platform"`
	ServerId       int32  `gorm:"column:serverId"`
	PlayerPlatform int32  `gorm:"column:playerPlatform"` //玩家平台
	PlayerServerId int32  `gorm:"column:playerServerId"` //玩家服务器
	PlayerId       int64  `gorm:"column:playerId"`
	PlayerName     string `gorm:"column:playerName"`
	AllianceId     int64  `gorm:"column:allianceId"`
	AllianceName   string `gorm:"column:allianceName"`
	Force          int64  `gorm:"column:force"`    //战力
	Pos            int32  `gorm:"column:pos"`      //职位
	CampType       int32  `gorm:"column:campType"` //阵营
	AlPos          int32  `gorm:"column:alPos"`    //仙盟职位
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (e *ChuangShiMemberEntity) GetId() int64 {
	return e.Id
}

func (e *ChuangShiMemberEntity) TableName() string {
	return "t_chuangshi_member"
}
