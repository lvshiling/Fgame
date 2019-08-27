/*此类自动生成,请勿修改*/
package template

/*周卡每日奖励模板配置*/
type WeekDayTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一天
	NextId int32 `json:"next_id"`

	//奖励天数
	DayInt int32 `json:"day"`

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
}
