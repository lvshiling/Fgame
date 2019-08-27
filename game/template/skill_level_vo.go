/*此类自动生成,请勿修改*/
package template

/*技能等级配置*/
type SkillLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个等级id
	NextId int32 `json:"next_id"`

	//类型
	Typ int32 `json:"type"`

	//等级
	Level int32 `json:"level"`

	//银两
	CostSilver int32 `json:"cost_yinliang"`

	//元宝
	CostGold int32 `json:"cost_yuanbao"`

	//物品
	CostItemId string `json:"cost_item_id"`

	//物品数量
	CostItemCount string `json:"cost_item_count"`

	//战力
	Force int32 `json:"force"`

	//伤害基础值
	DamageValueBase int32 `json:"damage_value_base"`

	//伤害值
	DamageValue int32 `json:"damage_value"`

	//技能伤害
	SpellDamage int32 `json:"spell_damage"`

	//技能强度
	SpellPower int32 `json:"spell_power"`
}
