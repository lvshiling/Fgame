/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerMarryDevelop)(nil))
}

/*表白等级*/
type PlayerMarryDevelop struct {
	PlayerLogMsg `bson:",inline"`

	//当前等级
	CurLevel int32 `json:"curLevel"`

	//变化前等级
	BeforeLevel int32 `json:"beforeLevel"`

	//变化等级
	ChangedLevel int32 `json:"changedLevel"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerMarryDevelop) LogName() string {
	return "player_marry_develop"
}
