/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerLingTongDev)(nil))
}

/*灵童养成系统*/
type PlayerLingTongDev struct {
	PlayerLogMsg `bson:",inline"`

	//系统类型
	ClassType int32 `json:"classType"`

	//当前阶数
	CurAdvancedNum int32 `json:"curAdvancedNum"`

	//变化前阶数
	BeforeAdvancedNum int32 `json:"beforeAdvancedNum"`

	//进阶阶数
	AdvancedNum int32 `json:"advancedNum"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerLingTongDev) LogName() string {
	return "player_lingtongdev"
}
