/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerEquipBaoKuAttendPoints)(nil))
}

/*装备宝库积分变化*/
type PlayerEquipBaoKuAttendPoints struct {
	PlayerLogMsg `bson:",inline"`

	//当前积分
	CurNum int32 `json:"curNum"`

	//变化前积分
	BefNum int32 `json:"befNum"`

	//相关物品id
	ItemId int32 `json:"itemId"`

	//相关物品数量
	ItemCount int32 `json:"itemCount"`

	//变化原因编号
	Reason int32 `json:"reason"`

	//变化原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerEquipBaoKuAttendPoints) LogName() string {
	return "player_equipbaoku_attendpoints"
}
