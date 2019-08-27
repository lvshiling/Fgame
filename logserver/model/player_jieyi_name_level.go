/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerJieYiNameLevel)(nil))
}

/*结义威名等级改变*/
type PlayerJieYiNameLevel struct {
	PlayerLogMsg `bson:",inline"`

	//之前威名等级
	BeforeNameLevel int32 `json:"beforeNameLevel"`

	//当前威名等级
	CurNameLevel int32 `json:"curNameLevel"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerJieYiNameLevel) LogName() string {
	return "player_jieyi_name_level"
}
