/*此类自动生成,请勿修改*/
package template

/*系统圣痕技能配置*/
type SystemSkillShengHenTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//技能类型
	Type int32 `json:"type"`

	//技能子类型
	SubType int32 `json:"sub_type"`

	//需要系统等级
	Number int32 `json:"number"`

	//技能等级
	Level int32 `json:"level"`

	//升级消耗的银两
	CostSilver int32 `json:"cost_yinliang"`

	//升级消耗的元宝
	CostGold int32 `json:"cost_yuanbao"`

	//物品
	CostItemId string `json:"cost_item_id"`

	//物品数量
	CostItemCount string `json:"cost_item_count"`

	//技能id
	SkillId int32 `json:"skill_id"`

	//需要的装备品质
	NeedEquipQuality int32 `json:"need_equip_quality"`

	//需要的装备数量
	NeedEquipCount int32 `json:"need_equip_count"`
}
