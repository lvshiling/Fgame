/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerAdditionSysAwake)(nil))
}

/*附加系统觉醒*/
type PlayerAdditionSysAwake struct {
	PlayerLogMsg `bson:",inline"`

	//当前是否觉醒
	CurIsAwake int32 `json:"curIsAwake"`

	//变化前是否觉醒
	BeforeIsAwake int32 `json:"beforeIsAwake"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerAdditionSysAwake) LogName() string {
	return "player_additionsys_awake"
}
