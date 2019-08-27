package entity

//玩家灵童各系统养成数据
type PlayerLingTongDevEntity struct {
	Id            int64  `gorm:"primary_key;column:id"`
	PlayerId      int64  `gorm:"column:playerId"`
	ClassType     int32  `gorm:"column:classType"`
	AdvancedId    int    `gorm:"column:advancedId"`
	SeqId         int32  `gorm:"column:seqId"`
	UnrealLevel   int32  `gorm:"column:unrealLevel"`
	UnrealNum     int32  `gorm:"column:unrealNum"`
	UnrealPro     int32  `gorm:"column:unrealPro"`
	UnrealInfo    string `gorm:"column:unrealInfo"`
	TimesNum      int32  `gorm:"column:timesNum"`
	Bless         int32  `gorm:"column:bless"`
	BlessTime     int64  `gorm:"column:blessTime"`
	CulLevel      int32  `gorm:"column:culLevel"`
	CulNum        int32  `gorm:"column:culNum"`
	CulPro        int32  `gorm:"column:culPro"`
	TongLingLevel int32  `gorm:"column:tongLingLevel"`
	TongLingNum   int32  `gorm:"column:tongLingNum"`
	TongLingPro   int32  `gorm:"column:tongLingPro"`
	Hidden        int32  `gorm:"column:hidden"`
	UpdateTime    int64  `gorm:"column:updateTime"`
	CreateTime    int64  `gorm:"column:createTime"`
	DeleteTime    int64  `gorm:"column:deleteTime"`
}

func (e *PlayerLingTongDevEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerLingTongDevEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerLingTongDevEntity) TableName() string {
	return "t_player_lingtong_develop"
}
