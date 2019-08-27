/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerBabyLearn)(nil))
}

/*宝宝读书*/
type PlayerBabyLearn struct {
	PlayerLogMsg `bson:",inline"`

	//当前等级
	CurBabyLevel int32 `json:"curBabyLevel"`

	//之前等级
	BeforeBabyLevel int32 `json:"beforeBabyLevel"`

	//变化等级
	ChangedLevel int32 `json:"changedLevel"`

	//原因编号
	Reason int32 `json:"reason"`

	//原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerBabyLearn) LogName() string {
	return "player_baby_learn"
}
