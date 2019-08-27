/*此类自动生成,请勿修改*/
package template

/*任务配置*/
type QuestTemplateVO struct {

	//id
	Id int `json:"id"`

	//免费次数用完是否直接完成任务
	IsAutoFinishFree int `json:"is_auto_finish_free"`

	//任务名称
	Name string `json:"name"`

	//标签
	QuestTag int32 `json:"quest_tag"`

	//任务类型
	QuestType int32 `json:"quest_type"`

	//任务子类型
	QuestSubType int32 `json:"quest_sub_type"`

	//任务要求(开天)
	QuestDemand string `json:"quest_demand"`

	//任务要求数量(开天)
	QuestDemandCount string `json:"quest_demand_count"`

	//任务要求(奕剑)
	QuestDemand2 string `json:"quest_demand2"`

	//任务要求数量(奕剑)
	QuestDemandCount2 string `json:"quest_demand_count2"`

	//任务要求(破月)
	QuestDemand3 string `json:"quest_demand3"`

	//任务要求数量(破月)
	QuestDemandCount3 string `json:"quest_demand_count3"`

	//任务需要的转数
	ReqZhuanshu int32 `json:"req_zhuanshu"`

	//任务需要的等级
	ReqLevel int32 `json:"req_level"`

	//任务物品id(开天)
	ReqItemId string `json:"req_item_id"`

	//任务物品数量(开天)
	ReqItemCount string `json:"req_item_count"`

	//任务物品id(奕剑)
	ReqItemId2 string `json:"req_item_id2"`

	//任务物品数量(奕剑)
	ReqItemCount2 string `json:"req_item_count2"`

	//任务物品id(破月)
	ReqItemId3 string `json:"req_item_id3"`

	//任务物品数量(破月)
	ReqItemCount3 string `json:"req_item_count3"`

	//前置任务
	PrevQuest string `json:"prev_quest"`

	//后置任务
	NextQuest string `json:"next_quest"`

	//后续非主线任务id
	FollowId string `json:"follow_id"`

	//任务npcId
	AcceptCreature int32 `json:"accept_creature"`

	//执行任务npcid
	CommitCreature int32 `json:"commit_creature"`

	//是否自动完成
	IsAutoCommi int32 `json:"is_auto_commi"`

	//是否有额外的任务奖励
	IsPopReward int32 `json:"is_pop_reward"`

	//任务文本
	Objectives string `json:"objectives"`

	//结束对话
	EndText string `json:"end_text"`

	//接受任务对话
	AcceptTaskText string `json:"accept_task_text"`

	//完成任务时对话
	DeliverTaskText string `json:"deliver_task_text"`

	//任务品质
	QuestLevel int32 `json:"quest_level"`

	//任务标记,屠魔任务用到
	QuestTb int32 `json:"quest_tb"`

	//特殊条件
	SpecialConditions int32 `json:"special_conditions"`

	//最小等级
	MinLevel int32 `json:"min_level"`

	//最大等级
	MaxLevel int32 `json:"max_level"`

	//最小转数
	MinZhuanshu int32 `json:"min_zhuanshu"`

	//最大转数
	MaxZhuanshu int32 `json:"max_zhuanshu"`

	//消耗银两
	ConsumeSilver int32 `json:"consume_silver"`

	//消耗元宝
	ConsumeGold int32 `json:"consume_gold"`

	//消耗物品(开天)
	ConsumeItem string `json:"consume_item"`

	//消耗物品数量(开天)
	ConsumeItemCount string `json:"consume_item_count"`

	//消耗物品(奕剑)
	ConsumeItem2 string `json:"consume_item2"`

	//消耗物品数量(奕剑)
	ConsumeItemCount2 string `json:"consume_item_count2"`

	//消耗物品(破月)
	ConsumeItem3 string `json:"consume_item3"`

	//消耗物品数量(破月)
	ConsumeItemCount3 string `json:"consume_item_count3"`

	//奖励物品id1
	RewItemId1 string `json:"rew_item_id_1"`

	//奖励物品id1数量
	RewItemCount1 string `json:"rew_item_count_1"`

	//奖励物品id2
	RewItemId2 string `json:"rew_item_id_2"`

	//奖励物品id2数量
	RewItemCount2 string `json:"rew_item_count_2"`

	//奖励物品id3
	RewItemId3 string `json:"rew_item_id_3"`

	//奖励物品id3数量
	RewItemCount3 string `json:"rew_item_count_3"`

	//奖励经验
	RewXp int32 `json:"rew_xp"`

	//奖励经验点
	RewExpPoint int32 `json:"rew_exp_point"`

	//奖励银两
	RewSilver int32 `json:"rew_silver"`

	//奖励元宝
	RewGold int32 `json:"rew_gold"`

	//奖励绑定元宝
	RewBindGold int32 `json:"rew_bind_gold"`

	//奖励转数
	RewZhuanshu int32 `json:"rew_zhuanshu"`

	//接受声音
	AcceptSound string `json:"accept_sound"`

	//执行声音
	ExcutedSound string `json:"excuted_sound"`

	//自动寻路
	AutoTrack int32 `json:"auto_track"`

	//引导箭头
	IsShowArrow int32 `json:"is_show_arrow"`

	//优先级
	PriorityLevel int32 `json:"priority_level"`

	//窗提示玩家传送到任务目标所在的地图与坐标 0为否1为是
	IsDeliver int32 `json:"is_deliver"`

	//弹出相关id
	WindowId string `json:"window_id"`

	//特殊任务类型时使用 天机牌
	DirectNumber int32 `json:"direct_number"`

	//消耗的飞鞋数量
	IsFeiXie int32 `json:"is_feixie"`

	//跳跃点
	Tiaoyuedian int32 `json:"tiaoyuedian"`

	//副本id
	FubenId int32 `json:"fuben_id"`
}
