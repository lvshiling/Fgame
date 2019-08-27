package entity

type ChuangShiChengFangJianSheEntity struct {
	Id          int64 `gorm:"column:id"`
	ServerId    int32 `gorm:"column:serverId"`
	PlayerId    int64 `gorm:"column:playerId"`
	CityId      int64 `gorm:"column:cityId"`
	JianSheType int32 `gorm:"column:jianSheType"`
	Num         int32 `gorm:"column:num"`
	Status      int32 `gorm:"column:status"` //建设状态
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (e *ChuangShiChengFangJianSheEntity) GetId() int64 {
	return e.Id
}

func (e *ChuangShiChengFangJianSheEntity) TableName() string {
	return "t_chuangshi_chengfang_jianshe"
}
