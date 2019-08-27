package entity

//玩家场景数据
type PlayerSceneEntity struct {
	Id          int64   `gorm:"primary_key;column:id"`
	PlayerId    int64   `gorm:"column:playerId"`
	MapId       int32   `gorm:"column:mapId"`
	SceneId     int64   `gorm:"column:sceneId"`
	PosX        float64 `gorm:"column:posX"`
	PosY        float64 `gorm:"column:posY"`
	PosZ        float64 `gorm:"column:posZ"`
	LastSceneId int64   `gorm:"column:lastSceneId"`
	LastMapId   int32   `gorm:"column:lastMapId"`
	LastPosX    float64 `gorm:"column:lastPosX"`
	LastPosY    float64 `gorm:"column:lastPosY"`
	LastPosZ    float64 `gorm:"column:lastPosZ"`
	UpdateTime  int64   `gorm:"column:updateTime"`
	CreateTime  int64   `gorm:"column:createTime"`
	DeleteTime  int64   `gorm:"column:deleteTime"`
}

func (psm *PlayerSceneEntity) GetId() int64 {
	return psm.Id
}

func (psm *PlayerSceneEntity) GetPlayerId() int64 {
	return psm.PlayerId
}

func (psm *PlayerSceneEntity) TableName() string {
	return "t_player_scene"
}
