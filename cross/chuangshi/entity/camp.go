package entity

//创世城
type ChuangShiCampEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	Platform       int32  `gorm:"column:platform"`
	ServerId       int32  `gorm:"column:serverId"`
	CampType       int32  `gorm:"column:campType"`       //阵营
	ShenWangStatus int32  `gorm:"column:shenWangStatus"` //神王竞选阶段
	Jifen          int64  `gorm:"column:jifen"`          //积分
	Diamonds       int64  `gorm:"column:diamonds"`       //钻石
	PayJifen       int64  `gorm:"column:payJifen"`       //阵营积分工资
	PayDiamonds    int64  `gorm:"column:payDiamonds"`    //阵营钻石工资
	LastShouYiTime int64  `gorm:"column:lastShouYiTime"` //上次工资时间
	TargetMap      string `gorm:"column:targetMap"`      //攻城目标
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (e *ChuangShiCampEntity) GetId() int64 {
	return e.Id
}

func (e *ChuangShiCampEntity) TableName() string {
	return "t_chuangshi_camp"
}
