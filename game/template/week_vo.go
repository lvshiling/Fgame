/*此类自动生成,请勿修改*/
package template

/*周卡模板配置*/
type WeekTemplateVO struct {

	//id
	Id int `json:"id"`

	//周卡类型
	Type int32 `json:"type"`

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

	//购买奖励
	GetItem string `json:"get_item"`

	//购买奖励数量
	GetItemCount string `json:"get_item_count"`

	//每日奖励关联id
	EveryDayGetBegin int32 `json:"every_day_get_begin"`

	//额外奖励需要的天数
	EwaiNeedDay int32 `json:"ewai_get_day"`

	//额外奖励物品
	EwaiGetItem string `json:"ewai_get_item"`

	//额外奖励数量
	EwaiGetItemCount string `json:"ewai_get_item_count"`
}
