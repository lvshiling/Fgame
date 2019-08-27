/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerGoldChanged)(nil))
}

/*元宝变化*/
type PlayerGoldChanged struct {
	PlayerLogMsg `bson:",inline"`

	//变化元宝数
	ChangedNum int64 `json:"changedNum"`

	//变化前的元宝数
	BeforeGold int64 `json:"beforeGold"`

	//变化前的绑元数
	BeforeBindGold int64 `json:"beforeBindGold"`

	//当前的元宝数
	CurGold int64 `json:"curGold"`

	//当前的绑元数
	CurBindGold int64 `json:"curBindGold"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerGoldChanged) LogName() string {
	return "player_glod_changed"
}
