/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerChess)(nil))
}

/*苍龙棋局*/
type PlayerChess struct {
	PlayerLogMsg `bson:",inline"`

	//抽奖次数
	AttendTimes int32 `json:"attendTimes"`

	//原因编号
	Reason int32 `json:"reason"`

	//日志原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerChess) LogName() string {
	return "player_chess"
}
