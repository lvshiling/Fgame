package entity

//玩家非进阶坐骑数据
type PlayerMountOtherEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Typ        int32 `gorm:"column:typ"`
	MountId    int32 `gorm:"column:mountId"`
	Level      int32 `gorm:"column:level"`
	UpNum      int32 `gorm:"column:upNum"`
	UpPro      int32 `gorm:"column:upPro"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pmoe *PlayerMountOtherEntity) GetId() int64 {
	return pmoe.Id
}

func (pmoe *PlayerMountOtherEntity) GetPlayerId() int64 {
	return pmoe.PlayerId
}

func (pmoe *PlayerMountOtherEntity) TableName() string {
	return "t_player_mount_other"
}
