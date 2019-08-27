package entity

type ArenapvpGuessRecordEntity struct {
	Id         int64 `gorm:"primary_key;column:id"` //Id
	ServerId   int32 `gorm:"column:serverId"`       //玩家Id
	PlayerId   int64 `gorm:"column:playerId"`       //玩家Id
	RaceNumber int32 `gorm:"column:raceNumber"`     //届数
	GuessType  int32 `gorm:"column:guessType"`      //竞猜类型
	GuessId    int64 `gorm:"column:guessId"`        //竞猜id
	WinnerId   int64 `gorm:"column:winnerId"`       //竞猜id
	Status     int32 `gorm:"column:status"`         //竞猜类型
	UpdateTime int64 `gorm:"column:updateTime"`     //更新时间
	CreateTime int64 `gorm:"column:createTime"`     //创建时间
	DeleteTime int64 `gorm:"column:deleteTime"`     //删除时间
}

func (e *ArenapvpGuessRecordEntity) GetId() int64 {
	return e.Id
}

func (e *ArenapvpGuessRecordEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *ArenapvpGuessRecordEntity) TableName() string {
	return "t_arenapvp_guess_record"
}
