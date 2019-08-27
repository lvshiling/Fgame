/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerNewBindGoldChanged)(nil))
}

/*新绑元变化*/
type PlayerNewBindGoldChanged struct {
	PlayerLogMsg `bson:",inline"`

	//变化元宝数
	ChangedNum int64 `json:"changedNum"`

	//变化前的绑元数
	BeforeBindGold int64 `json:"beforeBindGold"`

	//当前的绑元数
	CurBindGold int64 `json:"curBindGold"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerNewBindGoldChanged) LogName() string {
	return "player_new_bind_glod_changed"
}
