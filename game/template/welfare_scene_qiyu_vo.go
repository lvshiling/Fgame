/*此类自动生成,请勿修改*/
package template

/*运营活动奇遇岛副本配置*/
type WelfareSceneQiYuTemplateVO struct {

	//id
	Id int `json:"id"`

	//怪物起始配置
	BiologyBeginId int32 `json:"pos_begin_id"`

	//开始刷新时间
	RefreshBiologyBeginTime string `json:"refresh_biology_begin_time"`

	//刷新间隔时间
	RefreshBiologyTime int64 `json:"refresh_biology_time"`

	//刷新次数
	RefreshBiologyTimes int32 `json:"refresh_biology_times"`
}
