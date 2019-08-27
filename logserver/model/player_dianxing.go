/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerDianXing)(nil))
}

/*点星系统升级*/
type PlayerDianXing struct {
	PlayerLogMsg `bson:",inline"`

	//当前星谱
	CurXingPu int32 `json:"curXingPu"`

	//变化前星谱
	BeforeXingPu int32 `json:"beforeXingPu"`

	//当前等级
	CurLev int32 `json:"curLev"`

	//变化前等级
	BeforeLev int32 `json:"beforeLev"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerDianXing) LogName() string {
	return "player_dianxing"
}
