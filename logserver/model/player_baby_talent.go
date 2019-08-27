/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerBabyTalent)(nil))
}

/*宝宝天赋*/
type PlayerBabyTalent struct {
	PlayerLogMsg `bson:",inline"`

	//当前宝宝天赋
	CurBabyTalent []int32 `json:"curBabyTalent"`

	//之前宝宝天赋
	BeforeBabyTalent []int32 `json:"beforeBabyTalent"`

	//变化天赋
	ChangedTalent []int32 `json:"changedTalent"`

	//原因编号
	Reason int32 `json:"reason"`

	//原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerBabyTalent) LogName() string {
	return "player_baby_talent"
}
