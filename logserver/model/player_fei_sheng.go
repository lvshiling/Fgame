/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerFeiSheng)(nil))
}

/*飞升*/
type PlayerFeiSheng struct {
	PlayerLogMsg `bson:",inline"`

	//当前等级
	CurFeiShengLevel int32 `json:"curFeiShengLevel"`

	//变化前等级
	BeforeFeiShengLevel int32 `json:"beforeFeiShengLevel"`

	//提升的等级
	Uplevel int32 `json:"uplevel"`

	//当前功德
	CurGongDe int64 `json:"curGongDe"`

	//变化前功德
	BeforeGongDe int64 `json:"beforeGongDe"`

	//功德
	CostGongDe int64 `json:"costGongDe"`

	//升级原因编号
	Reason int32 `json:"reason"`

	//升级原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerFeiSheng) LogName() string {
	return "player_fei_sheng"
}
