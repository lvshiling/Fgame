/*此类自动生成,请勿修改*/
package template

/*仙尊问答配置*/
type QuizTemplateVO struct {

	//id
	Id int `json:"id"`

	//答题内容1
	AnswerA string `json:"answer_1"`

	//答题内容2
	AnswerB string `json:"answer_2"`

	//答题内容3
	AnswerC string `json:"answer_3"`

	//答题内容4
	AnswerD string `json:"answer_4"`

	//出现概率权重
	QuanZhong int32 `json:"quanzhong"`

	//正确答案
	RightAnswer int32 `json:"right_answer"`

	//答对奖励经验点
	RewExp int32 `json:"rew_exp"`

	//答对奖励经验点
	RewExpPoint int32 `json:"rew_exp_point"`

	//答对奖励银两
	RewardSilver int32 `json:"reward_silver"`

	//答对奖励绑元
	RewardBindGold int32 `json:"reward_bindgold"`

	//答对奖励元宝
	RewardGold int32 `json:"reward_gold"`

	//答对奖励物品的概率
	RewardItemRate int32 `json:"reward_item_rate"`

	//答对奖励物品
	RewItemId string `json:"rew_item_id"`

	//答对奖励物品数量
	RewItemCount string `json:"rew_item_count"`

	//答错奖励经验点
	ErrorRewExp int32 `json:"error_rew_exp"`

	//答错奖励经验点
	ErrorRewExpPoint int32 `json:"error_rew_exp_point"`

	//答错奖励银两
	ErrorRewardSilver int32 `json:"error_reward_silver"`

	//答错奖励绑元
	ErrorRewardBindGold int32 `json:"error_reward_bindgold"`

	//答错奖励元宝
	ErrorRewardGold int32 `json:"error_reward_gold"`

	//答错奖励物品的概率
	ErrorRewardItemRate int32 `json:"error_reward_item_rate"`

	//答错奖励物品
	ErrorRewItemId string `json:"error_rew_item_id"`

	//答错奖励物品数量
	ErrorRewItemCount string `json:"error_rew_item_count"`
}
