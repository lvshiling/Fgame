/*此类自动生成,请勿修改*/
package template

/*神魔称号配置*/
type ShenMoTitleTemplateVO struct {

	//id
	Id int `json:"id"`

	//最小击杀人数
	KillMin int32 `json:"kill_min"`

	//最大击杀人数
	KillMax int32 `json:"kill_max"`

	//被击杀后，击杀者获得的功勋
	GiveGongXun int32 `json:"give_gongxun"`

	//被击杀后，击杀者仙盟获得的积分
	GiveJiFen int32 `json:"give_jifen"`

	//称号
	Title int32 `json:"title"`
}
