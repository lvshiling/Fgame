package entity

// 玩家结义数据
type PlayerJieYiEntity struct {
	Id              int64  `gorm:"primary_key;column:id"`
	PlayerId        int64  `gorm:"column:playerId"`
	JieYiId         int64  `gorm:"column:jieYiId"`
	Name            string `gorm:"column:name"`
	Rank            int32  `gorm:"column:rank"`
	JieYiDaoJu      int32  `gorm:"column:jieYiDaoJu"`
	TokenType       int32  `gorm:"column:tokenType"`
	TokenLev        int32  `gorm:"column:tokenLev"`
	TokenPro        int32  `gorm:"column:tokenPro"`
	TokenNum        int32  `gorm:"column:tokenNum"`
	ShengWeiZhi     int32  `gorm:"column:shengWeiZhi"`
	NameLev         int32  `gorm:"column:nameLev"`
	NamePro         int32  `gorm:"column:namePro"`
	NameNum         int32  `gorm:"column:nameNum"`
	LastQiuYuanTime int64  `gorm:"column:lastQiuYuanTime"`
	LastDropTime    int64  `gorm:"column:lastDropTime"`
	LastInviteTime  int64  `gorm:"column:lastInviteTime"`
	LastPostTime    int64  `gorm:"column:lastPostTime"`
	LastLeaveTime   int64  `gorm:"column:lastLeaveTime"`
	UpdateTime      int64  `gorm:"column:updateTime"`
	CreateTime      int64  `gorm:"column:createTime"`
	DeleteTime      int64  `gorm:"column:deleteTime"`
}

func (e *PlayerJieYiEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerJieYiEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerJieYiEntity) TableName() string {
	return "t_player_jieyi"
}
