/*此类自动生成,请勿修改*/
package template

/*帝魂觉醒配置*/
type SoulAwakenTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//帝魂名称
	Name string `json:"name"`

	//帝魂大种类
	Type int32 `json:"type"`

	//帝魂类型(标签)
	SoulType int32 `json:"soul_type"`

	//觉醒阶别
	Order int32 `json:"order"`

	//觉醒所需要的物品id
	NeedItemId string `json:"need_item_id"`

	//觉醒所需要的物品数量
	NeedItemCount string `json:"need_item_count"`

	//该等级附带的技能id
	SkillId int32 `json:"skill_id"`

	//升级所需物品
	UplevelNeeditem string `json:"uplevel_needitem"`

	//升级所需物品数量
	UplevelItemCount string `json:"uplevel_item_count"`

	//升级帝魂技能id
	UplevelSkillId int32 `json:"uplevel_skill_id"`
}
