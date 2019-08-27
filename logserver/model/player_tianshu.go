/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerTianShu)(nil))
}

/*天书*/
type PlayerTianShu struct {
	PlayerLogMsg `bson:",inline"`

	//当前等级
	CurLevel int32 `json:"curLevel"`

	//变化前等级
	BeforeLevel int32 `json:"beforeLevel"`

	//提升等级
	Uplevel int32 `json:"uplevel"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerTianShu) LogName() string {
	return "player_tianshu"
}