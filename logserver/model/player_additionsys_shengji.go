/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerAdditionSysShengJi)(nil))
}

/*附加系统等级升级*/
type PlayerAdditionSysShengJi struct {
	PlayerLogMsg `bson:",inline"`

	//当前等级
	CurLev int32 `json:"curLev"`

	//变化前等级
	BeforeLev int32 `json:"beforeLev"`

	//当前升级次数
	CurUpNum int32 `json:"curUpNum"`

	//变化前升级次数
	BeforeUpNum int32 `json:"beforeUpNum"`

	//当前升级进度
	CurUpPro int32 `json:"curUpPro"`

	//变化前升级进度
	BeforeUpPro int32 `json:"beforeUpPro"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerAdditionSysShengJi) LogName() string {
	return "player_additionsys_shengji"
}
