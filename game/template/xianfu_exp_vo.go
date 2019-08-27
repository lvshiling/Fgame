/*此类自动生成,请勿修改*/
package template

/*经验副本配置*/
type XianFuExpTemplateVO struct {

	//id
	Id int `json:"id"`

	//名字
	Name string `json:"name"`

	//地图ID
	MapId int32 `json:"map_id"`

	//天藏老人ID
	BossId int32 `json:"boss_id"`

	//升级所需银两
	UpgradeYinliang int64 `json:"upgrade_yinliang"`

	//升级所需元宝
	UpgradeGold int32 `json:"upgrade_gold"`

	//升级所需绑元
	UpgradeBindGold int32 `json:"upgrade_bind_gold"`

	//升级所需物品ID
	UpgradeItemId int32 `json:"upgrade_item_id"`

	//升级所需物品数量
	UpgradeItemNum int32 `json:"upgrade_item_num"`

	//升级所需时间（秒）
	UpgradeTime int32 `json:"upgrade_time"`

	//加速所需元宝（元宝/小时）
	SpeedUpNeedGold float64 `json:"speed_up_need_gold"`

	//扫荡所需元宝
	SaodangNeedGold int32 `json:"saodang_need_gold"`

	//扫荡所需的物品，支持多个，字符串
	SaodangNeedItemId string `json:"saodang_need_item_id"`

	//扫荡所需的物品的数量，无物品相对应
	SaodangNeedItemCount string `json:"saodang_need_item_count"`

	//扫荡奖励经验值
	RawExp int64 `json:"raw_exp"`

	//扫荡奖励经验点
	RawExpPoint int64 `json:"raw_exp_point"`

	//扫荡奖励银两
	RawSilver int64 `json:"raw_silver"`

	//扫荡奖励元宝
	RawGold int32 `json:"raw_gold"`

	//扫荡奖励绑定元宝
	RawBindGold int32 `json:"raw_bind_gold"`

	//扫荡奖励物品ID，用逗号隔开
	RawItemId string `json:"raw_item_id"`

	//扫荡奖励物品数量，用逗号隔开
	RawItemCount string `json:"raw_item_count"`

	//扫荡奖励掉落包ID，用逗号隔开
	RawDropId string `json:"raw_drop_id"`

	//通关奖励物品ID，用逗号隔开
	GetItemId string `json:"get_item_id"`

	//通关奖励物品数量，用逗号隔开
	GetItemCount string `json:"get_item_count"`

	//奖励预览物品ID，字符串，支持多个
	RewardPreviewId string `json:"reward_preview_id"`

	//奖励预览的个数，字符串，与奖励预览ID相对应
	RewardPreviewCount string `json:"reward_preview_count"`

	//进入副本所需物品ID
	NeedItemId int32 `json:"need_item_id"`

	//进入副本所需物品数量
	NeedItemCount int32 `json:"need_item_count"`

	//下一级ID
	NextId int `json:"next_id"`

	//每日免费次数
	Free int32 `json:"free"`

	//扫荡波数限制
	GroupLimit int32 `json:"saodang_need_boshu"`
}
