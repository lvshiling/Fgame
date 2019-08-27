/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerExpChanged)(nil))
}

/*玩家经验变化*/
type PlayerExpChanged struct {
	PlayerLogMsg `bson:",inline"`

	//当前经验
	CurExp int64 `json:"curExp"`

	//变化前经验
	BeforeExp int64 `json:"beforeExp"`

	//当前等级
	CurLevel int32 `json:"curLevel"`

	//变化前等级
	BeforeLevel int32 `json:"beforeLevel"`

	//变化经验
	ChangedExp int64 `json:"changedExp"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerExpChanged) LogName() string {
	return "player_exp_changed"
}
