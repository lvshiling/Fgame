/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerFeedbackfee)(nil))
}

/*兑换*/
type PlayerFeedbackfee struct {
	PlayerLogMsg `bson:",inline"`

	//当前金额(分)
	CurMoney int32 `json:"curMoney"`

	//变化前金额(分)
	BeforeMoney int32 `json:"beforeMoney"`

	//变化的金额(分)
	Changed int32 `json:"changed"`

	//升级原因编号
	Reason int32 `json:"reason"`

	//升级原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerFeedbackfee) LogName() string {
	return "player_feedbackfee"
}
