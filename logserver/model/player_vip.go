/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerVip)(nil))
}

/*VIP*/
type PlayerVip struct {
	PlayerLogMsg `bson:",inline"`

	//当前等级
	CurVipLevel int32 `json:"curVipLevel"`

	//变化前等级
	BeforeVipLevel int32 `json:"beforeVipLevel"`

	//提升的等级
	Uplevel int32 `json:"uplevel"`

	//当前充值额
	CurGold int64 `json:"curGold"`

	//变化前充值额
	BeforeGold int64 `json:"beforeGold"`

	//充值额
	AddGold int64 `json:"addGold"`

	//升级原因编号
	Reason int32 `json:"reason"`

	//升级原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerVip) LogName() string {
	return "player_vip"
}
