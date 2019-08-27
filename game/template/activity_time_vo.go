/*此类自动生成,请勿修改*/
package template

/*活动时间模板配置*/
type ActivityTimeTemplateVO struct {

	//id
	Id int `json:"id"`

	//活动ID
	ActivityId int `json:"activity_id"`

	//开服几天后开始第一次活动
	AfterOpensvrDays int32 `json:"after_opensvr_days"`

	//开服几天后必定开启活动
	KaiQiOpensvrDay int32 `json:"kaiqi_opensvr_day"`

	//合服几天后必定开启活动
	KaiQiHeFuDay int32 `json:"kaiqi_hefu_day"`

	//开服几天后结束开启活动
	AfterKaifuDayOver int32 `json:"after_kaifu_day_over"`

	//合服几天后必定开启活动
	AfterHefuDays int32 `json:"after_hefu_days"`

	//合服几天后结束开启活动
	AfterHefuDayOver int32 `json:"after_hefu_day_over"`

	//时间段
	TimeQuantum string `json:"time_quantum"`

	//每周几开
	Weekday string `json:"weekday"`

	//每月几号开
	Monthday string `json:"monthday"`

	//活动开始时间
	BeginTime string `json:"begin_time"`

	//活动结束时间
	EndTime string `json:"end_time"`

	//活动持续时间
	ActivityTimes int64 `json:"activity_times"`

	//活动弹窗时间
	PopTimes int64 `json:"pop_times"`

	//活动开始前多久开始刷公告
	BeginNoticeTime int64 `json:"begin_notice_time"`

	//活动结束前多久开始刷公告
	EndNoticeTime int64 `json:"end_notice_time"`

	//距离活动开始前多久弹提示
	IconRemindTime int64 `json:"icon_remind_time"`

	//每日可获得的奖励次数
	AwardTimes int32 `json:"award_times"`

	//奖励倍数
	BeiShu int32 `json:"beishu"`
}
