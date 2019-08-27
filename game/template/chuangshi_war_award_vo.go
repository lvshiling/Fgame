/*此类自动生成,请勿修改*/
package template

/*创世城战奖励配置*/
type ChuangShiWarAwardTemplateVO struct {

	//id
	Id int `json:"id"`

	//类型
	Type int32 `json:"type"`

	//获得的创世积分
	WarAwardJifen int64 `json:"war_award_jifen"`

	//奖励银两
	WarAwardSilver int64 `json:"war_award_silver"`

	//奖励元宝
	WarAwardGold int64 `json:"war_award_gold"`

	//奖励绑定元宝
	WarAwardBindgold int64 `json:"war_award_bindgold"`

	//奖励物品id
	WarAwardItemId string `json:"war_award_item_id"`

	//奖励物品数量
	WarAwardItemCount string `json:"war_award_item_id_count"`
}
