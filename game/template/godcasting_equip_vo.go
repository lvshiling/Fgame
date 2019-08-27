/*此类自动生成,请勿修改*/
package template

/*神铸装备配置*/
type GodCastingEquipTemplateVO struct {

	//id
	Id int `json:"id"`

	//物品Id
	ItemId int32 `json:"item_id"`

	//可用来升级的物品
	UseItemId string `json:"use_item_id"`

	//可用来升级的物品对应数量
	UseItemCount string `json:"use_item_count"`

	//神铸成功率万分比
	UpdateWfb int32 `json:"update_wfb"`

	//神铸最小次数
	TimesMin int32 `json:"times_min"`

	//神铸最大次数
	TimesMax int32 `json:"times_max"`
}
