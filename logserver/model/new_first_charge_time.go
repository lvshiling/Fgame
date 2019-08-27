/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*NewFirstChargeTime)(nil))
}

/*新首充活动时间改变日志*/
type NewFirstChargeTime struct {
	SystemLogMsg `bson:",inline"`

	//新首充活动开始时间
	StartTime int64 `json:"startTime"`

	//持续时间
	Duration int32 `json:"duration"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *NewFirstChargeTime) LogName() string {
	return "new_first_charge_time"
}
