/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerJieYiShengWeiZhi)(nil))
}

/*结义声威值改变*/
type PlayerJieYiShengWeiZhi struct {
	PlayerLogMsg `bson:",inline"`

	//之前声威值
	BeforeShengWeiZhi int32 `json:"beforeShengWeiZhi"`

	//当前声威值
	CurShengWeiZhi int32 `json:"curShengWeiZhi"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerJieYiShengWeiZhi) LogName() string {
	return "player_jieyi_shengweizhi"
}
