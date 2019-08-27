/*此类自动生成,请勿修改*/
package template

/*八卦符石等级配置*/
type FuShiLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一级id
	NextId int32 `json:"next_id"`

	//符石类型
	Type int32 `json:"type"`

	//对应开启的技能id
	SkillId int32 `json:"skill_id"`

	//符石等级
	Level int32 `json:"level"`

	//升级所需物品id
	UpLevelItem string `json:"uplevel_item"`

	//升级所需物品数量
	UpLevelItemCount string `json:"uplevel_item_count"`

	//升级概率
	UpLevelRate int32 `json:"uplevel_rate"`
}
