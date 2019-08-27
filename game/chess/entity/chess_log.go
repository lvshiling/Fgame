package entity

//玩家棋局日志数据
type ChessLogEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	ServerId   int32  `gorm:"column:serverId"`
	PlayerName string `gorm:"column:playerName"`
	ItemId     int32  `gorm:"column:itemId"`
	ItemNum    int32  `gorm:"column:itemNum"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *ChessLogEntity) GetId() int64 {
	return e.Id
}

func (e *ChessLogEntity) TableName() string {
	return "t_chess_log"
}
