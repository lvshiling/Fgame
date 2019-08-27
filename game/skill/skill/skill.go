package skill

import (
	"fgame/fgame/game/battle/battle"
	cdcommon "fgame/fgame/game/cd/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	skillcommon "fgame/fgame/game/skill/common"
	skilleventtypes "fgame/fgame/game/skill/event/types"
	skilltemplate "fgame/fgame/game/skill/template"
	skilltypes "fgame/fgame/game/skill/types"
)

//技能
type SkillCDManager struct {
	bo      scene.BattleObject
	cdMap   map[int32]int64
	cdGroup *cdcommon.CDGroupManager
}

func (m *SkillCDManager) UseSkill(skillId int32) (flag bool) {
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)
	if skillTemplate == nil {
		return
	}
	if m.IsSkillInCD(skillId) {
		return
	}
	group := int32(skillTemplate.GetCdGroup().TemplateId())
	flag = m.cdGroup.UseCDGroup(group)
	if !flag {
		panic("skill:使用cd组应该成功")
	}
	now := global.GetGame().GetTimeService().Now()
	m.cdMap[skillId] = now
	gameevent.Emit(skilleventtypes.EventTypeSkillUse, m.bo, skillTemplate.TypeId)
	return
}

func (m *SkillCDManager) IsSkillInCD(skillId int32) bool {
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)
	if skillTemplate == nil {
		return true
	}
	group := int32(skillTemplate.GetCdGroup().TemplateId())
	if m.cdGroup.IsInCD(group) {
		return true
	}
	now := global.GetGame().GetTimeService().Now()
	cdElapse := int64(0)
	lastCdTime, ok := m.cdMap[skillId]
	if !ok {
		return false
	}
	cdElapse = now - lastCdTime
	if cdElapse < int64(skillTemplate.CdTime) {
		return true
	}
	return false
}

//技能管理器
func NewSkillCDManager(bo scene.BattleObject, cdGroupManager *cdcommon.CDGroupManager) *SkillCDManager {
	m := &SkillCDManager{}
	m.cdGroup = cdGroupManager
	m.bo = bo
	m.cdMap = make(map[int32]int64)
	return m
}

type SkillManager struct {
	bo scene.BattleObject
	//技能动作管理器
	*battle.SkillActionManager
	skillCDManager *SkillCDManager
	//技能
	skillSecondTypeMap map[skilltypes.SkillSecondType]map[int32]skillcommon.SkillObject
	//所有技能
	skillMap map[int32]skillcommon.SkillObject
}

func (m *SkillManager) IsSkillInCd(skillId int32) bool {
	return m.skillCDManager.IsSkillInCD(skillId)
}

func (m *SkillManager) UseSkill(skillId int32) bool {
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)
	if skillTemplate == nil {
		return false
	}
	flag := m.skillCDManager.UseSkill(skillId)
	if !flag {
		return false
	}
	m.SkillActionManager.AddSkillAction(skillId)

	return true
}

func (m *SkillManager) GetAllSkills() map[int32]skillcommon.SkillObject {
	return m.skillMap
}

func (m *SkillManager) GetSkills(typ skilltypes.SkillSecondType) map[int32]skillcommon.SkillObject {
	skillMap, ok := m.skillSecondTypeMap[typ]
	if !ok {
		return nil
	}
	return skillMap
}

func (m *SkillManager) GetSkill(skillId int32) skillcommon.SkillObject {
	ski, ok := m.skillMap[skillId]
	if !ok {
		return nil
	}
	return ski
}

func (m *SkillManager) addSkill(s skillcommon.SkillObject) {
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(s.GetSkillId())
	if skillTemplate == nil {
		return
	}
	m.skillMap[s.GetSkillId()] = s
	skillMap, ok := m.skillSecondTypeMap[skillTemplate.GetSkillSecondType()]
	if !ok {
		skillMap = make(map[int32]skillcommon.SkillObject)
		m.skillSecondTypeMap[skillTemplate.GetSkillSecondType()] = skillMap
	}
	skillMap[s.GetSkillId()] = s
}

func (m *SkillManager) RemoveSkill(skillId int32) {
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(skillId)
	if skillTemplate == nil {
		return
	}
	skillTypeId := skillTemplate.TypeId
	skillMap, ok := m.skillSecondTypeMap[skillTemplate.GetSkillSecondType()]
	if ok {
		delete(skillMap, skillTypeId)
	}
	delete(m.skillMap, skillTypeId)
	gameevent.Emit(skilleventtypes.EventTypeSkillRemove, m.bo, skillTypeId)
}

func (m *SkillManager) Heartbeat() {
	m.SkillActionManager.Heartbeat()
}

func (m *SkillManager) isValid(skillId int32) bool {
	if skillId <= 0 {
		return false
	}
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(skillId)
	if skillTemplate == nil {
		return false
	}
	return true
}

//静态技能是否存在
func (m *SkillManager) ifSkillExist(skillId int32) bool {
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(skillId)
	if skillTemplate == nil {
		return false
	}
	so, ok := m.skillMap[skillTemplate.TypeId]
	if !ok {
		return false
	}
	if so.GetLevel() != skillTemplate.Lev {
		return false
	}
	return true
}

//静态技能改变判断
func (m *SkillManager) isChangeValid(oldSkillId int32, newSkillId int32) bool {
	if oldSkillId <= 0 || newSkillId <= 0 {
		return false
	}
	oldTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(oldSkillId)
	if oldTemplate == nil {
		return false
	}
	newTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(newSkillId)
	if newTemplate == nil {
		return false
	}

	//技能类型是否为职业技能
	if newTemplate.FirstType == int32(skilltypes.SkillFirstTypeRole) {
		return false
	}
	//技能类型是否相同
	if oldTemplate.FirstType != newTemplate.FirstType {
		return false
	}
	flag := m.ifSkillExist(oldSkillId)
	if !flag {
		return false
	}

	flag = m.ifSkillExist(newSkillId)
	if flag {
		return false
	}

	return true
}

//动态技能是否存在
func (m *SkillManager) ifDynamicSkillExist(skillId int32) bool {
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(skillId)
	if skillTemplate == nil {
		return false
	}
	_, ok := m.skillMap[skillTemplate.TypeId]
	if !ok {
		return false
	}
	return true
}

//添加静态技能
func (m *SkillManager) AddStaticSkill(skillId int32) bool {
	flag := m.isValid(skillId)
	if !flag {
		return false
	}
	flag = m.ifSkillExist(skillId)
	if flag {
		return false
	}
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(skillId)
	if skillTemplate == nil {
		return false
	}

	lev := skillTemplate.Lev
	skillType := skillTemplate.TypeId
	pts := skillcommon.CreateSkillObject(skillType, lev, nil)
	m.addSkill(pts)
	gameevent.Emit(skilleventtypes.EventTypeSkillAdd, m.bo, pts)
	return true
}

//改变动态技能
func (m *SkillManager) ChangeDynamicSkill(skillId int32, level int32) {
	flag := m.ifDynamicSkillExist(skillId)
	if !flag {
		pts := skillcommon.CreateSkillObject(skillId, level, nil)
		m.addSkill(pts)
		gameevent.Emit(skilleventtypes.EventTypeSkillAdd, m.bo, pts)
		return
	}

	skillObj := m.skillMap[skillId]
	skillObj.SetLevel(level)
}

//改变静态技能
func (m *SkillManager) ChangeSkill(oldSkillId int32, newSkillId int32) bool {
	//添加新技能
	if oldSkillId == 0 {
		return m.AddStaticSkill(newSkillId)
	}
	//卸下技能
	if newSkillId == 0 {
		//移除旧技能
		m.RemoveSkill(oldSkillId)
		return true
	}
	//更换技能
	flag := m.isChangeValid(oldSkillId, newSkillId)
	if !flag {
		return false
	}
	//移除旧技能
	m.RemoveSkill(oldSkillId)
	//添加新技能
	m.AddStaticSkill(newSkillId)
	return true
}

func (m *SkillManager) ToAllSkillInfo() (skillList []*skillcommon.SkillObjectImpl) {
	skillList = make([]*skillcommon.SkillObjectImpl, 0, 16)
	for _, skillObj := range m.GetAllSkills() {
		skillInfo := &skillcommon.SkillObjectImpl{
			SkillId: skillObj.GetSkillId(),
			Level:   skillObj.GetLevel(),
		}
		skillList = append(skillList, skillInfo)
	}
	return
}

func CreateSkillManager(bo scene.BattleObject, cdGroup *cdcommon.CDGroupManager, skillList []skillcommon.SkillObject) *SkillManager {
	skillManager := &SkillManager{}
	skillManager.bo = bo
	skillManager.skillCDManager = NewSkillCDManager(bo, cdGroup)
	skillManager.SkillActionManager = battle.CreateSkillActionManager(bo)
	skillManager.skillSecondTypeMap = make(map[skilltypes.SkillSecondType]map[int32]skillcommon.SkillObject)
	skillManager.skillMap = make(map[int32]skillcommon.SkillObject)

	for _, ski := range skillList {
		skillManager.addSkill(ski)
	}
	return skillManager
}
