/*此类自动生成,请勿修改*/
package template

/*屠龙装备技能配置*/
type TuLongEquipSkillTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//技能等级
	Level int32 `json:"level"`

	//技能类型
	Type int32 `json:"type"`

	//需要装备阶数
	NeedJieShu int32 `json:"value1"`

	//需要装备数量
	NeedEquipNum int32 `json:"value2"`

	//需要装备总强化等级
	NeedStrengthenLevel int32 `json:"value3"`

	//技能id
	SkillId int32 `json:"skill_id"`

	//物品id
	UplevelItem string `json:"uplevel_item"`

	//物品数量
	UplevelItemCount string `json:"uplevel_item_count"`

	//升级成功率
	UplevelRate int32 `json:"uplevel_rate"`
}
