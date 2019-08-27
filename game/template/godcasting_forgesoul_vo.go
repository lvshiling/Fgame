/*此类自动生成,请勿修改*/
package template

/*神铸锻魂配置*/
type GodCastingForgeSoulTemplateVO struct {

	//id
	Id int `json:"id"`

	//锻魂类型
	Type int32 `json:"type"`

	//装备部位
	SubType int32 `json:"sub_type"`

	//关联的锻魂等级表起始ID
	LevelBeginId int32 `json:"level_begin_id"`

	//锻魂消耗的物品ID
	UseItemId int32 `json:"use_item_id"`

	//关联特殊技能表ID
	TeshuSkillId int32 `json:"teshu_skill_id"`
}
