package common

import (
	"fgame/fgame/game/global"
	scenetemplate "fgame/fgame/game/scene/template"
)

//cd组
type CDGroupManager struct {
	cdGroupMap map[int32]int64
}

func (m *CDGroupManager) UseCDGroup(group int32) bool {
	if m.IsInCD(group) {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	m.cdGroupMap[group] = now
	return true
}

func (m *CDGroupManager) IsInCD(group int32) bool {
	now := global.GetGame().GetTimeService().Now()
	tempCdGroupTemplate := scenetemplate.GetSceneTemplateService().GetCdGroup(group)
	if tempCdGroupTemplate == nil {
		return true
	}

	lastTime, ok := m.cdGroupMap[group]
	if !ok {
		return false
	}
	cdGroupElapse := now - lastTime
	if cdGroupElapse < int64(tempCdGroupTemplate.Time) {
		return true
	}
	return false
}

//cd组
func NewCDGroupManager() *CDGroupManager {
	m := &CDGroupManager{}
	m.cdGroupMap = make(map[int32]int64)
	return m
}
