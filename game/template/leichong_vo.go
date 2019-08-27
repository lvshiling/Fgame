/*此类自动生成,请勿修改*/
package template

/*运营活动次数奖励模板配置*/
type TimesRewTemplateVO struct {

	//id
	Id int `json:"id"`

	//活动id
	Group int32 `json:"group"`

	//抽奖次数
	DrawTimes int32 `json:"draw_times"`

	//vip等级
	VipLevel int32 `json:"vip_level"`

	//奖励id
	RawId string `json:"raw_id"`

	//奖励数量
	RawCount string `json:"raw_count"`
}
