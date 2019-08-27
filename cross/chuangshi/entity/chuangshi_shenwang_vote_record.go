package entity

type ChuangShiShenWangVoteRecordEntity struct {
	Id                  int64  `gorm:"column:id"`
	Platform            int32  `gorm:"column:platform"`
	ServerId            int32  `gorm:"column:serverId"`
	CampType            int32  `gorm:"column:campType"` //阵营
	PlayerServerId      int32  `gorm:"column:playerServerId"`
	PlayerId            int64  `gorm:"column:playerId"`
	PlayerName          string `gorm:"column:playerName"`
	HouXuanPlatform     int32  `gorm:"column:houXuanPlatform"`
	HouXuanGameServerId int32  `gorm:"column:houXuanGameServerId"`
	HouXuanPlayerId     int64  `gorm:"column:houXuanPlayerId"`
	HouXuanPlayerName   string `gorm:"column:houXuanPlayerName"`
	LastVoteTime        int64  `gorm:"column:lastVoteTime"` //上次投票时间
	UpdateTime          int64  `gorm:"column:updateTime"`
	CreateTime          int64  `gorm:"column:createTime"`
	DeleteTime          int64  `gorm:"column:deleteTime"`
}

func (e *ChuangShiShenWangVoteRecordEntity) GetId() int64 {
	return e.Id
}

func (e *ChuangShiShenWangVoteRecordEntity) TableName() string {
	return "t_chuangshi_shenwang_vote_record"
}
