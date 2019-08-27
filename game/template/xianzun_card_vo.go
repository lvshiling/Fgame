/*此类自动生成,请勿修改*/
package template

/*仙尊特权卡配置*/
type XianZunCardTemplateVO struct {

	//id
	Id int `json:"id"`

	//仙尊卡类型
	Type int32 `json:"type"`

	//仙尊卡名称
	Name string `json:"name"`

	//所需元宝
	NeedGold int32 `json:"need_gold"`

	//持续时间
	Duration int64 `json:"duration"`

	//激活奖励银两
	JiHuoRewSilver int32 `json:"jihuo_rew_silver"`

	//激活奖励元宝
	JiHuoRewGold int32 `json:"jihuo_rew_gold"`

	//激活奖励绑元
	JiHuoRewBindGold int32 `json:"jihuo_rew_bind_gold"`

	//激活获得物品id
	ItemId string `json:"jihuo_get_item"`

	//激活获得物品数量
	ItemCount string `json:"jihuo_get_item_count"`

	//每日领取银两的数量
	DayRewSilver int32 `json:"day_rew_silver"`

	//每日领取元宝的数量
	DayRewGold int32 `json:"day_rew_gold"`

	//每日领取绑元的数量
	DayRewBindGold int32 `json:"day_rew_bind_gold"`

	//每日奖励物品id
	DayRewItem string `json:"day_rew_item"`

	//每日奖励物品数量
	DayRewItemCount string `json:"day_rew_item_count"`

	//银两副本增加的每日免费挑战次数
	SilverXianFuFreeAdd int32 `json:"silver_xianfu_free_add"`

	//经验副本增加的每日免费挑战次数
	ExpXianFuFreeAdd int32 `json:"exp_xianfu_free_add"`

	//获得经验加成的生物类型
	BiologySetType string `json:"biology_set_type"`

	//杀怪获得的经验加成比例
	ExpBiologyAddPercent int32 `json:"exp_biology_add_percent"`

	//坐骑基础属性增加的万分比
	MountAttrAddPercent int32 `json:"mount_attr_add_percent"`

	//天界Boss免费进入次数
	TianJieBossFreeAdd int32 `json:"tianjie_boss_free_add"`

	//3V3积分增加万分比
	JiFenAddPercent int32 `json:"arena_jifen_add_percent"`

	//每天获得积分上限增加的万分比
	JiFenMaxAddPercent int32 `json:"arena_jifen_max_add_percent"`

	//战翼基础属性增加的万分比
	WingAttrAddPercent int32 `json:"wing_attr_add_percent"`
}
