package common

//系统技能信息
type SystemSkillInfo struct {
	Type         int32                     `json:"type"`
	SysSkillList []*SystemSkillSubTypeInfo `json:"systemSkillSubTypeInfo"`
}

type SystemSkillSubTypeInfo struct {
	SubType int32 `json:"subType"`
	Level   int32 `json:"level"`
}

type AllSystemSkillInfo struct {
	SystemSkillList []*SystemSkillInfo `json:"systemSkillInfo"`
}
