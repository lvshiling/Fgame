/*此类自动生成,请勿修改*/
package template

/*仙尊问答常量模板配置*/
type QuizConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//问答活动开始时间
	BeginTime string `json:"activity_begin_time"`

	//问答活动结束时间
	EndTime string `json:"activity_end_time"`

	//必定刷新的时间点配置
	RefreshTime string `json:"refresh_time"`

	//题目刷新间隔时间(毫秒)
	IntervalTime int32 `json:"interval_time"`

	//每题题目的答题时间
	DaTiTime int32 `json:"dati_time"`

	//假消息时间下限(毫秒)
	MsgTimeMin int32 `json:"msg_time_min"`

	//假消息时间上限(毫秒)
	MsgTimeMax int32 `json:"msg_time_max"`
}
