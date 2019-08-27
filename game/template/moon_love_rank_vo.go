/*此类自动生成,请勿修改*/
package template

/*月下情缘排行榜配置*/
type MoonloveRankTemplateVO struct {

	//id
	Id int `json:"id"`

	//0豪气榜1魅力榜
	Type int32 `json:"type"`

	//排名
	Rank int32 `json:"rank"`

	//奖励经验值
	RewExp int64 `json:"rew_exp"`

	//奖励经验点
	RewExpPoint int32 `json:"rew_exp_point"`

	//奖励银两
	RewSilver int64 `json:"rew_silver"`

	//奖励元宝
	RewGold int32 `json:"rew_gold"`

	//奖励绑定元宝
	RewBindGold int32 `json:"rew_bind_gold"`

	//奖励物品Id
	RewItemId string `json:"rew_item_id"`

	//奖励物品数量
	RewItemCount string `json:"rew_item_count"`
}
