/*此类自动生成,请勿修改*/
package template

/*帝魂遗迹星级奖励配置*/
type SoulRuinsStarTemplateVO struct {

	//id
	Id int `json:"id"`

	//对应章节
	Chapter int32 `json:"chapter"`

	//类型
	Type int32 `json:"type"`

	//领取星数
	NeedStar int32 `json:"need_star"`

	//奖励元宝
	RewGold int32 `json:"rew_gold"`

	//奖励银两
	RewYinliang int32 `json:"rew_yinliang"`

	//奖励经验固定值
	RewExp int32 `json:"rew_exp"`

	//奖励经验uplev点
	RewUplev int32 `json:"rew_uplev"`

	//奖励物品1 ID
	RewItemId1 int32 `json:"rew_item_id1"`

	//奖励物品1 数量
	RewItemCount1 int32 `json:"rew_item_count1"`

	//奖励物品2 ID
	RewItemId2 int32 `json:"rew_item_id2"`

	//奖励物品2 数量
	RewItemCount2 int32 `json:"rew_item_count2"`

	//奖励物品3 ID
	RewItemId3 int32 `json:"rew_item_id3"`

	//奖励物品3 数量
	RewItemCount3 int32 `json:"rew_item_count3"`

	//关联到物品表,点击查看奖励tips
	RewardShowItemId int32 `json:"reward_show_item_id"`

	//章节界面奖励描述
	RewardShowResult string `json:"reward_show_result"`
}
