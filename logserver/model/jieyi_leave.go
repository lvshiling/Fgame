/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*JieYiLeave)(nil))
}

/*离开结义*/
type JieYiLeave struct {
	JieYiLogMsg `bson:",inline"`

	//离开的玩家
	PlayerId int64 `json:"playerId"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *JieYiLeave) LogName() string {
	return "jieyi_leave"
}
