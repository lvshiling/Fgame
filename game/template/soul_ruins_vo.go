/*此类自动生成,请勿修改*/
package template

/*帝魂遗迹*/
type SoulRuinsTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//副本名称
	Name string `json:"name"`

	//关联地图id
	MapId int32 `json:"map_id"`

	//章节数
	Chapter int32 `json:"chapter"`

	//难度
	Type int32 `json:"type"`

	//关卡等级
	Level int32 `json:"level"`

	//奖励银两
	RewYinliang int32 `json:"rew_yinliang"`

	//奖励经验固定值
	RewExp int32 `json:"rew_exp"`

	//奖励经验
	RewUplev int32 `json:"rew_uplev"`

	//奖励元宝
	RewGold int32 `json:"rew_gold"`

	//首次通关奖励挑战次数
	RewTime int32 `json:"rew_time"`

	//前置副本ID
	FrontId int32 `json:"front_id"`

	//特殊事件-隐藏BOSS的怪物ID
	SpecialBossId int32 `json:"special_boss_id"`

	//特殊事件-隐藏BOSS给予的BUFFID
	SpecialBossBuff int32 `json:"special_boss_buff"`

	//特殊事件-隐藏BOSS的出现概率
	SpecialBossRate int32 `json:"special_boss_rate"`

	//特殊事件,隐藏BOSS刷的组
	SpecialBossGroup int32 `json:"special_boss_group"`

	//特殊事件-帝魂降临 出现的帝魂id
	SoulId int32 `json:"soul_id"`

	//特殊事件-帝魂降临 给予的物品ID
	SoulItemId int32 `json:"soul_item_id"`

	//特殊事件-帝魂降临 给予的物品ID
	SoulItemCount int32 `json:"soul_item_count"`

	//特殊事件-帝魂降临 出现的概率
	SoulRate int32 `json:"soul_rate"`

	//特殊事件,神秘高人的组
	SoulGroup int32 `json:"soul_group"`

	//特殊事件-马贼 刷新的怪物ID
	RobberId int32 `json:"robber_id"`

	//特殊事件-马贼 刷新的怪物数量
	RobberCount int32 `json:"robber_count"`

	//特殊事件-马贼 出现的概率
	RobberRate int32 `json:"robber_rate"`

	//特殊事件,马贼的刷怪组
	RobberGroup int32 `json:"robber_group"`

	//马贼索要的银两
	RobberSilver int32 `json:"robber_silver"`

	//首次打必出的随机事件
	FirstEvent int32 `json:"first_event"`

	//星数减少时限1
	Time1 int32 `json:"time_1"`

	//星数减少时限2
	Time2 int32 `json:"time_2"`

	//星数减少时限3
	Time3 int32 `json:"time_3"`

	//星数减少时限2
	SweepItemId int32 `json:"sweep_item_id"`

	//扫荡所需的物品数量
	SweepItemCount int32 `json:"sweep_item_count"`

	//扫荡所需背包空位
	SweepNeedBag int32 `json:"sweep_need_bag"`

	//掉落包1id
	SweepDrop1 string `json:"sweep_drop1"`

	//掉落包2id
	SweepDrop2 int32 `json:"sweep_drop2"`

	//掉落包3id
	SweepDrop3 int32 `json:"sweep_drop3"`

	//掉落包4id
	SweepDrop4 int32 `json:"sweep_drop4"`

	//隐藏BOSS掉落包ID
	SpecialBossDrop int32 `json:"special_boss_drop"`

	//帝魂降临掉落包
	SoulDrop int32 `json:"soul_drop"`

	//马贼掉落包ID
	RobberDrop int32 `json:"robber_drop"`

	//奖励预览
	RewPre string `json:"rew_pre"`

	//推荐战力
	RecForce int32 `json:"rec_force"`
}
