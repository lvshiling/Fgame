/*此类自动生成,请勿修改*/
package template

/*vip模板配置*/
type HuiYuanTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一级
	NextId int32 `json:"next_id"`

	//会员等级
	Level int32 `json:"level"`

	//需要消费的元宝
	NeedGold int32 `json:"need_gold"`

	//持续时间
	Duration int64 `json:"duration"`

	//奖励银两
	RewSilver int32 `json:"rew_silver"`

	//奖励元宝
	RewGold int32 `json:"rew_gold"`

	//奖励绑定元宝
	RewBindGold int32 `json:"rew_bind_gold"`

	//首次奖励
	GetItem string `json:"get_item"`

	//首次奖励数量
	GetItemCount string `json:"get_item_count"`

	//每天奖励物品
	RewItem string `json:"rew_item"`

	//每天奖励数量
	RewItemCount string `json:"rew_item_count"`

	//是否屠魔第四个任务
	TumoFour int32 `json:"tumo_four"`

	//帝陵遗迹
	DilingCount int32 `json:"diling_count"`

	//秘境仙府
	MijingCount int32 `json:"mijing_count"`

	//屠魔任务次数
	TumoCount int32 `json:"tumo_count"`

	//天机牌任务次数
	TianjipaiCount int32 `json:"tianjipai_count"`

	//镖车次数
	BiaocheCount int32 `json:"biaoche_count"`

	//劫镖次数
	JiebiaoCount int32 `json:"jiebiao_count"`

	//银两棋局次数
	QijuSilverCount int32 `json:"qiju_silver_count"`

	//后台版本
	HoutaiType int32 `json:"houtai_type"`
}
