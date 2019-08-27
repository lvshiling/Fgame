/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerItemChanged)(nil))
}

/*物品变化*/
type PlayerItemChanged struct {
	PlayerLogMsg `bson:",inline"`

	//变化的物品id
	ChangedItemId int32 `json:"changedItemId"`

	//变化的物品名字
	ChangedItemName string `json:"changedItemName"`

	//变化的物品数
	ChangedItemNum int32 `json:"changedItemNum"`

	//变化前的物品数量
	BeforeItemNum int32 `json:"beforeItemNum"`

	//当前的物品数
	CurItemNum int32 `json:"curItemNum"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerItemChanged) LogName() string {
	return "player_item_changed"
}
