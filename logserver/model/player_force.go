/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerForce)(nil))
}

/*战力*/
type PlayerForce struct {
	PlayerLogMsg `bson:",inline"`

	//当前战力
	Force int64 `json:"force"`

	//变化前的战力
	BeforeForce int64 `json:"beforeForce"`

	//变化掩码
	Mask uint64 `json:"mask"`

	//原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerForce) LogName() string {
	return "player_force"
}
