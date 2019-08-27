/*此类自动生成,请勿修改*/
package template

/*材料副本配置*/
type MaterialTemplateVO struct {

	//id
	Id int `json:"id"`

	//副本类型
	Type int32 `json:"type"`

	//名字
	Name string `json:"name"`

	//地图ID
	MapId int32 `json:"map_id"`

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

	//进入副本所需物品ID
	NeedItemId int32 `json:"need_item_id"`

	//进入副本所需物品数量
	NeedItemCount int32 `json:"need_item_count"`

	//每日免费次数
	Free int32 `json:"free"`

	//总次数
	AllTimes int32 `json:"all_times"`

	//扫荡需要等级
	NeedLevel int32 `json:"need_level"`

	//扫荡波数限制
	GroupLimit int32 `json:"saodang_need_boshu"`
}
