/*此类自动生成,请勿修改*/
package template

/*开服活动配置*/
type OpenserverActivityTemplateVO struct {

	//id
	Id int `json:"id"`

	//活动标识
	Group int32 `json:"group"`

	//活动类型
	Type int32 `json:"type"`

	//活动子类型
	SubType int32 `json:"subtype"`

	//参数1
	Value1 int32 `json:"value_1"`

	//参数2
	Value2 int32 `json:"value_2"`

	//参数3
	Value3 int32 `json:"value_3"`

	//参数4
	Value4 int32 `json:"value_4"`

	//奖励银两
	RewSilver int32 `json:"rew_silver"`

	//奖励元宝
	RewGold int32 `json:"rew_gold"`

	//奖励绑元
	RewGoldBind int32 `json:"rew_gold_bind"`

	//奖励物品id
	AwardItemId string `json:"award_item_id"`

	//奖励物品数量
	AwardItemCount string `json:"award_item_count"`

	//活动名
	Label string `json:"label"`

	//邮件内容
	MailDes string `json:"mail_des"`

	//过期类型
	TimeType int32 `json:"time_type"`

	//过期时间
	LimitTime int32 `json:"limit_time"`

	//物品过期标识
	ItemExpireFlag string `json:"item_expire_flag"`
}
