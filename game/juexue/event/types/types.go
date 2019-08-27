package types

import (
	juexuetypes "fgame/fgame/game/juexue/types"
)

type JueXueEventType string

const (
	//绝学升级
	EventTypeJueXueUpgrade JueXueEventType = "jueXueUpgrade"
	//绝学卸下
	EventTypeJueXueUnload JueXueEventType = "jueXueUnload"
	//绝学使用
	EventTypeJueXueUse JueXueEventType = "jueXueUse"
	//绝学顿悟
	EventTypeJueXueInsight JueXueEventType = "jueXueInsight"
	//绝学激活
	EventTypeJueXueAcitivate JueXueEventType = "jueXueActivate"
)

type JueXueInsightEventData struct {
	typ        juexuetypes.JueXueType
	oldSkillId int32
}

func (j *JueXueInsightEventData) GetType() juexuetypes.JueXueType {
	return j.typ
}

func (j *JueXueInsightEventData) GetOldSkillId() int32 {
	return j.oldSkillId
}

func CreateJueXueInsightEventData(typ juexuetypes.JueXueType, oldSkillId int32) *JueXueInsightEventData {
	d := &JueXueInsightEventData{
		typ:        typ,
		oldSkillId: oldSkillId,
	}
	return d
}
