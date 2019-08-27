package entity

//外域Boss掉落记录数据
type OutlandBossDropRecordsEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	ServerId   int32  `gorm:"column:serverId"`
	KillerName string `gorm:"column:killerName"`
	BiologyId  int32  `gorm:"column:biologyId"`
	MapId      int32  `gorm:"column:mapId"`
	DropTime   int64  `gorm:"column:dropTime"`
	ItemInfo   string `gorm:"column:itemInfo"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (ere *OutlandBossDropRecordsEntity) GetId() int64 {
	return ere.Id
}

func (ere *OutlandBossDropRecordsEntity) TableName() string {
	return "t_outland_boss_drop_records"
}
