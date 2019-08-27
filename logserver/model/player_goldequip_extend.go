/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerGoldEquipExtend)(nil))
}

/*金装继承*/
type PlayerGoldEquipExtend struct {
	PlayerLogMsg `bson:",inline"`

	//继承前等级
	BeforeLevel int32 `json:"beforeLevel"`

	//继承后等级
	AfterLevel int32 `json:"afterLevel"`

	//原因编号
	Reason int32 `json:"reason"`

	//日志原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerGoldEquipExtend) LogName() string {
	return "player_goldequip_extend"
}
