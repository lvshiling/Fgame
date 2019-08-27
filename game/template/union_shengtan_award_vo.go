/*此类自动生成,请勿修改*/
package template

/*仙盟圣坛奖励模板配置*/
type UnionShengTanAwardTemplateVO struct {

	//id
	Id int `json:"id"`

	//最小等级
	MinLev int32 `json:"min_lev"`

	//奖励经验
	RewExp int32 `json:"rew_exp"`

	//奖励经验点
	RewExpPoint int32 `json:"rew_exp_point"`

	//银两奖励
	RewSilver int32 `json:"rew_silver"`

	//元宝奖励
	RewGold int32 `json:"rew_gold"`

	//绑元奖励
	RewBindGold int32 `json:"rew_bind_gold"`

	//奖励物品id
	RewItemId string `json:"rew_item_id"`

	//奖励物品数量
	RewItemCount string `json:"rew_item_count"`
}
