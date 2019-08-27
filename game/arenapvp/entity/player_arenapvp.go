package entity

type PlayerArenapvpEntity struct {
	Id          int64 `gorm:"primary_key;column:id"` //Id
	PlayerId    int64 `gorm:"column:playerId"`       //玩家Id
	ReliveTimes int32 `gorm:"column:reliveTimes"`    //已复活次数
	OutStatus   int32 `gorm:"column:outStatus"`      //是否淘汰：0否1是
	JiFen       int32 `gorm:"column:jiFen"`          //积分
	GuessNotice int32 `gorm:"column:guessNotice"`    //竞猜提醒设置
	PvpRecord   int32 `gorm:"column:pvpRecord"`      //pvp成绩
	TicketFlag  int32 `gorm:"column:ticketFlag"`     //是否购买门票
	UpdateTime  int64 `gorm:"column:updateTime"`     //更新时间
	CreateTime  int64 `gorm:"column:createTime"`     //创建时间
	DeleteTime  int64 `gorm:"column:deleteTime"`     //删除时间
}

func (e *PlayerArenapvpEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerArenapvpEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerArenapvpEntity) TableName() string {
	return "t_player_arenapvp"
}
