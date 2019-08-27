package model

type AllianceQuery struct {
	Id            int    `gorm:"column:id"`
	AllianceName  string `gorm:"column:allianceName"`
	AllianceLevel int    `gorm:"column:allianceLevel"`
	TotalForce    int    `gorm:"column:totalForce"`
	CreateTime    int    `gorm:"column:createTime"`
	PlayerCount   int    `gorm:"column:playerCount"`
	Notice        string `gorm:"column:notice"`
}
