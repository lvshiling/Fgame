/*此类自动生成,请勿修改*/
package template

/*活动采集模板配置*/
type CollectActivityTemplateVO struct {

	//id
	Id int `json:"id"`

	//活动id
	Group int32 `json:"group"`

	//采集物id
	BiologyId int32 `json:"biology_id"`

	//刷新数量
	RebornCount int32 `json:"reborn_count"`

	//刷新地图id
	RebornMapId int32 `json:"reborn_map_id"`

	//刷新时间点
	RebornTime string `json:"reborn_time"`
}
