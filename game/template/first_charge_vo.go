/*此类自动生成,请勿修改*/
package template

/*首冲配置*/
type FirstChargeTemplateVO struct {

	//id
	Id int `json:"id"`

	//职业
	Profession int32 `json:"profession"`

	//性别
	Gender int32 `json:"gender"`

	//奖励银两
	RewSilver int32 `json:"rew_silver"`

	//奖励元宝
	RewGold int32 `json:"rew_gold"`

	//奖励绑元
	RewGoldBind int32 `json:"rew_gold_bind"`

	//奖励物品id
	RewItemId string `json:"rew_item_id"`

	//奖励物品数量
	RewItemCount string `json:"rew_item_count"`
}
