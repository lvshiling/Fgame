package entity

type PlayerArenapvpGuessLogEntity struct {
	Id         int64 `gorm:"primary_key;column:id"` //Id
	PlayerId   int64 `gorm:"column:playerId"`       //玩家Id
	RaceNum    int32 `gorm:"column:raceNum"`        //届数
	GuessId    int64 `gorm:"column:guessId"`        //竞猜id
	WinnerId   int64 `gorm:"column:winnerId"`       //获胜id
	GuessType  int32 `gorm:"column:guessType"`      //竞猜类型
	UpdateTime int64 `gorm:"column:updateTime"`     //更新时间
	CreateTime int64 `gorm:"column:createTime"`     //创建时间
	DeleteTime int64 `gorm:"column:deleteTime"`     //删除时间
}

func (e *PlayerArenapvpGuessLogEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerArenapvpGuessLogEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerArenapvpGuessLogEntity) TableName() string {
	return "t_player_arenapvp_guess_log"
}
