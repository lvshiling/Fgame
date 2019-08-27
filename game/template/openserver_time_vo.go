/*此类自动生成,请勿修改*/
package template

/*开服活动时间配置*/
type OpenserverTimeTemplateVO struct {

	//id
	Id int `json:"id"`

	//活动标识
	Group int32 `json:"group"`

	//活动名称
	Name string `json:"name"`

	//活动类型
	Type int32 `json:"type"`

	//活动子类型
	SubType int32 `json:"subtype"`

	//活动时间类型
	TimeType int32 `json:"time_type"`

	//参数1
	Value1 int64 `json:"value_1"`

	//参数2
	Value2 int64 `json:"value_2"`

	//关联功能开启
	OpenId int32 `json:"open_id"`

	//功能开启邮件标题
	MailTitle string `json:"mail_title"`

	//功能开启邮件内容
	MailDes string `json:"mail_des"`

	//功能开启奖励
	MailRewItem string `json:"mail_rew_item"`

	//功能开启奖励数量
	MailRewItemCount string `json:"mail_rew_item_count"`

	//关联活动
	RelatedActivity string `json:"related_activity"`

	//是否循环活动
	IsCircle int32 `json:"is_circle"`

	//是否合服循环活动
	IsMergeCircle int32 `json:"is_hefu_circle"`

	//开服几天不开启
	CloseOpenDay int32 `json:"close_open_day"`

	//是否传音
	IsChuanyin int32 `json:"is_chuanyin"`

	//传音内容
	ChuanyinText string `json:"chuanyin_text"`
}
