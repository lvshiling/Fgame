/*此类自动生成,请勿修改*/
package template

/*仙术模板配置*/
type AllianceSkillTemplateVO struct {

	//id
	Id int `json:"id"`

	//仙术名称
	Name string `json:"name"`

	//仙术类型
	Type int32 `json:"type"`

	//仙法等级
	Level int32 `json:"level"`

	//仙法描述
	Des string `json:"des"`

	//仙法描述（扩展）
	DesInfo string `json:"des_info"`

	//升级所需贡献度
	NeedContribution int64 `json:"need_contribution"`

	//属性id
	SkillId int32 `json:"skill_id"`

	//下级id
	NextId int32 `json:"next_id"`

	//前置等级id
	need_level_id int32 `json:"need_level_id"`

	//需要的仙盟等级
	NeedUnionLevel int32 `json:"need_union_level"`

	//仙术开启所需仙盟等级
	OpenNeedLevel int32 `json:"open_need_level"`

	//前端技能顺序
	shunxu int32 `json:"shunxu"`
}
