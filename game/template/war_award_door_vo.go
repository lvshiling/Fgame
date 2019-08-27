/*此类自动生成,请勿修改*/
package template

/*城门奖励配置*/
type WarAwardDoorTemplateVO struct {

	//id
	Id int `json:"id"`

	//生物id
	BiologyId int32 `json:"biology_id"`

	//银两
	WarDoorSilver int32 `json:"war_door_silver"`

	//元宝
	WarDoorGold int32 `json:"war_door_gold"`

	//绑元
	WarDoorBindgold int32 `json:"war_door_bindgold"`

	//物品
	WarDoorItemId string `json:"war_door_item_id"`

	//物品数量
	WarDoorItemIdCount string `json:"war_door_item_id_count"`

	//所需积分
	NeedJiFen int32 `json:"need_jifen"`
}
