/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerLaBa)(nil))
}

/*拉霸*/
type PlayerLaBa struct {
	PlayerLogMsg `bson:",inline"`

	//当前次数
	CurTimes int32 `json:"curTimes"`

	//花费元宝
	CostGold int32 `json:"costGold"`

	//奖励元宝
	RewGold int32 `json:"rewGold"`

	//升级原因编号
	Reason int32 `json:"reason"`

	//升级原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerLaBa) LogName() string {
	return "player_laba"
}
