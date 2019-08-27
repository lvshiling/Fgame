/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerEquipBaoKuLuckyPoints)(nil))
}

/*装备宝库幸运值变化*/
type PlayerEquipBaoKuLuckyPoints struct {
	PlayerLogMsg `bson:",inline"`

	//当前幸运值
	CurNum int32 `json:"curNum"`

	//变化前幸运值
	BefNum int32 `json:"befNum"`

	//相关物品信息
	WithItems string `json:"withItems"`

	//变化原因编号
	Reason int32 `json:"reason"`

	//变化原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerEquipBaoKuLuckyPoints) LogName() string {
	return "player_equipbaoku_luckypoints"
}
