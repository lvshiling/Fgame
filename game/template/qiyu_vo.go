/*此类自动生成,请勿修改*/
package template

/*奇遇模板配置*/
type QiYuTemplateVO struct {

	//id
	Id int `json:"id"`

	//接取任务等级
	Level string `json:"level"`

	//接取转生等级
	ZhuanSheng string `json:"zhuansheng"`

	//接取飞升等级
	FeiSheng string `json:"feisheng"`

	//奇遇子任务
	QuestId string `json:"quest_id"`

	//经验
	RewExp int32 `json:"rew_exp"`

	//经验点
	RewExpPoint int32 `json:"rew_exp_point"`

	//奖励银两
	RewSilver int32 `json:"rew_silver"`

	//奖励绑定元宝
	RewBindGold int32 `json:"rew_bind_gold"`

	//奖励元宝
	RewGold int32 `json:"rew_gold"`

	//奖励物品Id
	RewItemId string `json:"rew_item_id"`

	//奖励物品数量
	RewItemCount string `json:"rew_item_count"`

	//过期时间
	GuoQiTime int32 `json:"guoqi_time"`
}
