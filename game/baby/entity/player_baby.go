package entity

//玩家宝宝数据
type PlayerBabyEntity struct {
	Id                 int64  `gorm:"primary_key;column:id"`
	PlayerId           int64  `gorm:"column:playerId"`
	Name               string `gorm:"column:name"`
	Sex                int32  `gorm:"column:sex"`
	Quality            int32  `gorm:"column:quality"`
	SkillList          string `gorm:"column:skillList"`
	LearnLevel         int32  `gorm:"column:learnLevel"`
	LearnExp           int32  `gorm:"column:learnExp"`
	AttrBeiShu         int32  `gorm:"column:attrBeiShu"`
	ActivateTimes      int32  `gorm:"column:activateTimes"`
	LockTimes          int32  `gorm:"column:lockTimes"`
	RefreshTimes       int32  `gorm:"column:refreshTimes"`
	RefreshCostItemNum int32  `gorm:"column:costItemNum"`
	UpdateTime         int64  `gorm:"column:updateTime"`
	CreateTime         int64  `gorm:"column:createTime"`
	DeleteTime         int64  `gorm:"column:deleteTime"`
}

func (e *PlayerBabyEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerBabyEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerBabyEntity) TableName() string {
	return "t_player_baby"
}
