/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*JieYiJoin)(nil))
}

/*加入结义*/
type JieYiJoin struct {
	JieYiLogMsg `bson:",inline"`

	//加入的玩家
	PlayerId int64 `json:"playerId"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *JieYiJoin) LogName() string {
	return "jieyi_join"
}
