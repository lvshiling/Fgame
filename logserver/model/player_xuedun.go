/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerXueDun)(nil))
}

/*血盾*/
type PlayerXueDun struct {
	PlayerLogMsg `bson:",inline"`

	//当前阶数
	CurNumber int32 `json:"curNumber"`

	//当前星级
	CurStar int32 `json:"curStar"`

	//变化前阶数
	BeforeNumber int32 `json:"beforeNumber"`

	//变化前星级
	BeforeStar int32 `json:"beforeStar"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerXueDun) LogName() string {
	return "player_xuedun"
}
