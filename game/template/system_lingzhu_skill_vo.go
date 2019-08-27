/*此类自动生成,请勿修改*/
package template

/*附加灵珠技能配置*/
type SystemLingZhuSkillTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//灵童等级
	LingtongId int32 `json:"lingtong_id"`

	//等级
	Level int32 `json:"level"`

	//需要灵珠的最低等级
	NeedLingzhuLevel int32 `json:"need_lingzhu_level"`

	//关联技能Id
	SkillId int32 `json:"skill_id"`
}
