/*此类自动生成,请勿修改*/
package template

/*城战奖励配置*/
type WarAwardTemplateVO struct {

	//id
	Id int `json:"id"`

	//类型
	Type int32 `json:"type"`

	//银两
	WarAwardSilver int32 `json:"war_award_silver"`

	//元宝
	WarAwardGold int32 `json:"war_award_gold"`

	//绑定
	WarAwardBindgold int32 `json:"war_award_bindgold"`

	//物品
	WarAwardItemId string `json:"war_award_item_id"`

	//数量
	WarAwardItemIdCount string `json:"war_award_item_id_count"`
}
