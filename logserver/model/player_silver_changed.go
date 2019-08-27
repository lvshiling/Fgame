/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerSilverChanged)(nil))
}

/*银两变化*/
type PlayerSilverChanged struct {
	PlayerLogMsg `bson:",inline"`

	//变化银两数
	ChangedNum int64 `json:"changedNum"`

	//变化前的银两数
	BeforeSilver int64 `json:"beforeSilver"`

	//当前的银两数
	CurSilver int64 `json:"curSilver"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerSilverChanged) LogName() string {
	return "player_silver_changed"
}
