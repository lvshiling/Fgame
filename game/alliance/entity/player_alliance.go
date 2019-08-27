package entity

//玩家仙盟数据
type PlayerAllianceEntity struct {
	Id                       int64  `gorm:"primary_key;column:id"`
	PlayerId                 int64  `gorm:"column:playerId"`
	AllianceId               int64  `gorm:"column:allianceId"`
	AllianceName             string `gorm:"column:allianceName"`
	AllianceLevel            int32  `gorm:"column:allianceLevel"`
	DonateMap                string `gorm:"column:donateMap"`
	CurrentGongXian          int64  `gorm:"column:currentGongXian"`
	LastJuanXuanTime         int64  `gorm:"column:lastJuanXuanTime"`
	SceneRewardMap           string `gorm:"column:sceneRewardMap"`
	WarPoint                 int32  `gorm:"column:warPoint"`
	LastAllianceSceneEndTime int64  `gorm:"column:lastAllianceSceneEndTime"`
	YaoPai                   int32  `gorm:"column:yaoPai"`
	LastYaoPaiUpdateTime     int64  `gorm:"column:lastYaoPaiUpdateTime"`
	ConvertTimes             int32  `gorm:"column:convertTimes"`
	LastConvertUpdateTime    int64  `gorm:"column:lastConvertUpdateTime"`
	ReliveTime               int32  `gorm:"column:reliveTime"`
	LastReliveTime           int64  `gorm:"column:lastReliveTime"`
	DepotPoint               int32  `gorm:"column:depotPoint"`
	LastMemberCallTime       int64  `gorm:"column:lastMemberCallTime"`
	LastYuXiMemberCallTime   int64  `gorm:"column:lastYuXiMemberCallTime"`
	TotalWinTime             int32  `gorm:"column:totalWinTime"`
	UpdateTime               int64  `gorm:"column:updateTime"`
	CreateTime               int64  `gorm:"column:createTime"`
	DeleteTime               int64  `gorm:"column:deleteTime"`
}

func (e *PlayerAllianceEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerAllianceEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerAllianceEntity) TableName() string {
	return "t_player_alliance"
}
