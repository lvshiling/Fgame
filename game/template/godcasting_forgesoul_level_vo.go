/*此类自动生成,请勿修改*/
package template

/*神铸锻魂升级配置*/
type GodCastingForgeSoulLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一id
	NextId int32 `json:"next_id"`

	//等级
	Level int32 `json:"level"`

	//锻魂消耗的物品数量
	UseItemCount int32 `json:"use_item_count"`

	//触发几率万分比
	ChufaRate int32 `json:"chufa_rate"`

	//抵抗几率万分比
	DikangRate int32 `json:"dikang_rate"`

	//升级几率万分比
	UpdateWfb int32 `json:"update_wfb"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//战力增加
	AddPower int32 `json:"add_power"`
}
