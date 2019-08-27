/*此类自动生成,请勿修改*/
package template

/*活动模板配置*/
type ActivityTemplateVO struct {

	//id
	Id int `json:"id"`

	//活动名称
	Name string `json:"name"`

	//活动类型
	ActiveType int32 `json:"active_type"`

	//活动地图ID
	Mapid int32 `json:"mapid"`

	//子活动地图ID
	SubMapid int32 `json:"sub_mapid"`

	//参与最小等级限制
	LevMin int32 `json:"lev_min"`

	//参与最大等级限制
	LevMax int32 `json:"lev_max"`

	//功能开启ID
	KaiqiId int32 `json:"kaiqi_id"`

	//进入需要的功能开启ID
	NeedKaiqiId int32 `json:"need_kaiqi_id"`

	//每天限制参与次数
	JoinCount int32 `json:"join_count"`

	//公告间隔时间
	NoticeInterval int64 `json:"notice_interval"`

	//活动开始时公告内容
	NoticeContentBegin string `json:"notice_content_begin"`

	//活动结束时公告内容
	NoticeContentEnd string `json:"notice_content_end"`

	//即将开始时的公告内容
	NoticeContentSoonbegin string `json:"notice_content_soonbegin"`

	//即将结束时公告的内容
	NoticeContentSoonend string `json:"notice_content_soonend"`
}
