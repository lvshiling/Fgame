/*此类自动生成,请勿修改*/
package template

/*附加灵珠升级配置*/
type SystemLingZhuUpLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//等级
	Level int32 `json:"level"`

	//该等级增加的生命
	Hp int32 `json:"hp"`

	//该等级增加的攻击
	Attack int32 `json:"attack"`

	//该等级增加的防御
	Defence int32 `json:"defence"`

	//百分比
	Percent int32 `json:"percent"`

	//升级消耗的物品数量
	UseItemCount int32 `json:"use_item_count"`

	//成功率万分比
	UpdateWfb int32 `json:"update_wfb"`

	//每次升级最小次数
	TimesMin int32 `json:"times_min"`

	//每次升级最大次数
	TimesMax int32 `json:"times_max"`

	//每次升级每次随机增加祝福值最小数目
	AddMin int32 `json:"add_min"`

	//每次升级每次随机增加祝福值最大数目
	AddMax int32 `json:"add_max"`

	//祝福值最大
	ZhufuMax int32 `json:"zhufu_max"`
}
