package entity

type PlayerArenaEntity struct {
	Id              int64 `gorm:"primary_key;column:id"`  //Id
	PlayerId        int64 `gorm:"column:playerId"`        //玩家Id
	EndTime         int64 `gorm:"column:endTime"`         //结束时间
	CulRewardTime   int32 `gorm:"column:culRewardTime"`   //已经奖励次数
	TotalRewardTime int32 `gorm:"column:totalRewardTime"` //累计获取次数
	JiFenCount      int32 `gorm:"column:jiFenCount"`      //累计积分
	JiFenDay        int32 `gorm:"column:jiFenDay"`        //每日积分
	ArenaTime       int64 `gorm:"column:arenaTime"`       //积分更新时间
	WinCount        int32 `gorm:"column:winCount"`        //连胜次数
	FailCount       int32 `gorm:"column:failCount"`       //连败次数
	DayWinCount     int32 `gorm:"column:dayWinCount"`     //当天连胜
	DayMaxWinCount  int32 `gorm:"column:dayMaxWinCount"`  //当天最高连胜
	ReliveTime      int32 `gorm:"column:reliveTime"`      //复活次数
	RankRewTime     int64 `gorm:"column:rankRewTime"`     //上次周榜奖励时间
	UpdateTime      int64 `gorm:"column:updateTime"`      //更新时间
	CreateTime      int64 `gorm:"column:createTime"`      //创建时间
	DeleteTime      int64 `gorm:"column:deleteTime"`      //删除时间

}

func (e *PlayerArenaEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerArenaEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerArenaEntity) TableName() string {
	return "t_player_arena"
}
