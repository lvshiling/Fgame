/*此类自动生成,请勿修改*/
package template

/*好友推送模板配置*/
type FriendNoticeTemplateVO struct {

	//id
	Id int `json:"id"`

	//推送类型
	Type int32 `json:"type"`

	//推送条件
	TiaoJian int32 `json:"tiaojian"`

	//反馈奖励1
	JiDanRewardExp int32 `json:"jidan_reward_exp"`

	//反馈奖励2
	XianHuaRewardExp int32 `json:"xianhua_reward_exp"`

	//反馈奖励1经验点
	JiDanRewardExpPoint int32 `json:"jidan_reward_exp_point"`

	//反馈奖励2经验点
	XianHuaRewardExpPoint int32 `json:"xianhua_reward_exp_point"`

	//砸鸡蛋或送鲜花的玩家收到的物品
	ZhuheItem string `json:"zhuhe_item"`

	//砸鸡蛋或送鲜花的玩家收到的物品数量
	ZhuheItemCount string `json:"zhuhe_item_count"`
}
