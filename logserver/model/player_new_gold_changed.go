/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerNewGoldChanged)(nil))
}

/*新元宝变化*/
type PlayerNewGoldChanged struct {
	PlayerLogMsg `bson:",inline"`

	//变化元宝数
	ChangedNum int64 `json:"changedNum"`

	//变化前的元宝数
	BeforeGold int64 `json:"beforeGold"`

	//当前的元宝数
	CurGold int64 `json:"curGold"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerNewGoldChanged) LogName() string {
	return "player_new_glod_changed"
}
