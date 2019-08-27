package entity

//创世城池建设
type ChuangShiCityJianSheEntity struct {
	Id            int64  `gorm:"primary_key;column:id"`
	Platform      int32  `gorm:"column:platform"`
	ServerId      int32  `gorm:"column:serverId"`
	CityId        int64  `gorm:"column:cityId"`        //城池id
	JianSheType   int32  `gorm:"column:jianSheType"`   //建设类型
	JianSheLevel  int32  `gorm:"column:jianSheLevel"`  //建设等级
	JianSheExp    int32  `gorm:"column:jianSheExp"`    //建设经验
	SkillLevelSet int32  `gorm:"column:skillLevelSet"` //当前使用技能（天气台专用）
	SkillMap      string `gorm:"column:skillMap"`      //技能激活记录（天气台专用）
	UpdateTime    int64  `gorm:"column:updateTime"`
	CreateTime    int64  `gorm:"column:createTime"`
	DeleteTime    int64  `gorm:"column:deleteTime"`
}

func (e *ChuangShiCityJianSheEntity) GetId() int64 {
	return e.Id
}

func (e *ChuangShiCityJianSheEntity) TableName() string {
	return "t_chuangshi_city_jianshe"
}
