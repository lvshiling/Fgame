package entity

type PlayerMaterialEntity struct {
	Id           int64 `gorm:"primary_key;column:id"`
	PlayerId     int64 `gorm:"column:playerId"`
	MaterialType int32 `gorm:"column:materialType"`
	UseTimes     int32 `gorm:"column:useTimes"`
	Group        int32 `gorm:"column:group"`
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (e *PlayerMaterialEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerMaterialEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerMaterialEntity) TableName() string {
	return "t_player_material"
}
