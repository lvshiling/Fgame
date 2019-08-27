/*此类自动生成,请勿修改*/
package template

/*好友添加模板配置*/
type FriendAddTemplateVO struct {

	//id
	Id int `json:"id"`

	//好友数量
	Num int32 `json:"num"`

	//奖励银两数量
	RewardSilver int32 `json:"reward_silver"`

	//奖励物品id
	RewardItemId string `json:"reward_item_id"`

	//奖励物品数量
	RewardItemCount string `json:"reward_item_count"`

	//奖励经验
	RewardExp int32 `json:"reward_exp"`

	//奖励经验点
	RewardExpPoint int32 `json:"reward_exp_point"`
}
