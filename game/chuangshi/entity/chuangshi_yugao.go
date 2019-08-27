package entity

type ChuangShiYuGaoEntity struct {
	Id         int64 `gorm:"column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	Num        int64 `gorm:"column:num"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *ChuangShiYuGaoEntity) GetId() int64 {
	return e.Id
}

func (e *ChuangShiYuGaoEntity) TableName() string {
	return "t_chuangshi_yugao"
}
