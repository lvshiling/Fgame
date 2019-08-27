/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerMarryDevelopExp)(nil))
}

/*表白经验*/
type PlayerMarryDevelopExp struct {
	PlayerLogMsg `bson:",inline"`

	//当前表白经验
	CurExp int32 `json:"curExp"`

	//变化前表白经验
	BeforeExp int32 `json:"beforeExp"`

	//当前表白等级
	CurLevel int32 `json:"curLevel"`

	//变化前表白等级
	BeforeLevel int32 `json:"beforeLevel"`

	//变化经验
	ChangedExp int32 `json:"changedExp"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerMarryDevelopExp) LogName() string {
	return "player_marry_develop_exp"
}
