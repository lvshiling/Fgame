/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerMassacre)(nil))
}

/*戮仙刃*/
type PlayerMassacre struct {
	PlayerLogMsg `bson:",inline"`

	//当前阶数
	CurAdvancedNum int32 `json:"curAdvancedNum"`

	//变化前阶数
	BeforeAdvancedNum int32 `json:"beforeAdvancedNum"`

	//变化星数
	ChangedNum int32 `json:"changedNum"`

	//变化前杀气数量
	BefShaQiNum int64 `json:"befShaQiNum"`

	//当前杀气数量
	CurShaQiNum int64 `json:"curShaQiNum"`

	//当前地图
	CurMapId int32 `json:"curMapId"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerMassacre) LogName() string {
	return "player_massacre"
}
