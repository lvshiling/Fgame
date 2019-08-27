/*此类自动生成,请勿修改*/
package template

/*活动炼制模板配置*/
type MadeTemplateVO struct {

	//id
	Id int `json:"id"`

	//活动id
	GroupId int32 `json:"group"`

	//最小区间
	LevelMin int32 `json:"Level_min"`

	//最大区间
	LevelMax int32 `json:"Level_max"`

	//经验
	Exp int32 `json:"exp"`

	//经验点
	ExpPoint int32 `json:"uplev"`

	//所需消耗
	CostBase string `json:"cost_base"`
}
