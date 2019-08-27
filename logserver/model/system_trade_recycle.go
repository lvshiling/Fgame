/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*SystemTradeRecycle)(nil))
}

/*系统交易回购池日志*/
type SystemTradeRecycle struct {
	SystemLogMsg `bson:",inline"`

	//回购池金额
	RecycleGold int64 `json:"recycleGold"`

	//原因编号
	Reason int32 `json:"reason"`

	//原因
	ReasonText string `json:"reasonText"`
}

func (c *SystemTradeRecycle) LogName() string {
	return "system_trade_recycle"
}
