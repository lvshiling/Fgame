package entity

//玩家棋局数据
type PlayerChessEntity struct {
	Id                    int64 `gorm:"primary_key;column:id"`
	PlayerId              int64 `gorm:"column:playerId"`
	ChessId               int32 `gorm:"column:chessId"`
	AttendTimes           int32 `gorm:"column:attendTimes"`
	TotalAttendTimes      int32 `gorm:"column:totalAttendTimes"`
	ChessType             int32 `gorm:"column:chessType"`
	LastSystemRefreshTime int64 `gorm:"column:lastSystemRefreshTime"`
	UpdateTime            int64 `gorm:"column:updateTime"`
	CreateTime            int64 `gorm:"column:createTime"`
	DeleteTime            int64 `gorm:"column:deleteTime"`
}

func (e *PlayerChessEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerChessEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerChessEntity) TableName() string {
	return "t_player_chess"
}
