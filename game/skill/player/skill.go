package player

import (
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skillcommon "fgame/fgame/game/skill/common"
	"fgame/fgame/game/skill/dao"
	skilleventtypes "fgame/fgame/game/skill/event/types"
	skilltemplate "fgame/fgame/game/skill/template"
	skilltypes "fgame/fgame/game/skill/types"
	gametemplate "fgame/fgame/game/template"
	xiuxianbookeventtypes "fgame/fgame/game/welfare/xiuxianbook/event/types"
	"fgame/fgame/pkg/idutil"
)

//玩家技能数据管理器
type PlayerSkillDataManager struct {
	p player.Player
	//职业技能数据
	playerSkillObjectMap map[int32]*PlayerSkillObject
	//技能列表
	skillList []skillcommon.SkillObject
	//玩家技能cd
	playerSkillCdObjectMap map[int32]*PlayerSkillCdObject
}

func (m *PlayerSkillDataManager) GetAllSkill() []skillcommon.SkillObject {
	return m.skillList
}

//根据技能id获取职业技能信息
func (m *PlayerSkillDataManager) GetSkill(skillId int32) (pso *PlayerSkillObject) {
	skillObj, exist := m.playerSkillObjectMap[skillId]
	if !exist {
		return
	}
	pso = skillObj
	return
}

func (m *PlayerSkillDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerSkillDataManager) Load() (err error) {
	//TODO 数据加载封装
	m.playerSkillObjectMap = make(map[int32]*PlayerSkillObject)
	m.playerSkillCdObjectMap = make(map[int32]*PlayerSkillCdObject)
	pseList, err := dao.GetSkillDao().GetRoleSkillList(m.p.GetId())
	if err != nil {
		return
	}
	for _, pse := range pseList {
		pso := NewPlayerSkillObject(m.p)
		pso.FromEntity(pse)
		m.addPersistentSkill(pso)
	}

	skillCdList, err := dao.GetSkillDao().GetSkillCdList(m.p.GetId())
	if err != nil {
		return
	}
	for _, skillCd := range skillCdList {
		psco := NewPlayerSkillCdObject(m.p)
		psco.FromEntity(skillCd)
		m.playerSkillCdObjectMap[skillCd.SkillId] = psco
	}
	return nil
}

func (m *PlayerSkillDataManager) AfterLoad() (err error) {
	//刷新
	m.refresh()
	return
}

//刷新数据
func (m *PlayerSkillDataManager) refresh() {
	role := m.Player().GetRole()
	sex := m.Player().GetSex()
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(role, sex)
	normalSkillMap := playerCreateTemplate.GetSkillNormalMap()
	if len(normalSkillMap) > 0 {
		m.addNormalSkill(normalSkillMap)
	}
	m.addJumpSkill(playerCreateTemplate.GetJumpSkill())
	return
}

//普通攻击添加
func (m *PlayerSkillDataManager) addNormalSkill(normalSkillMap map[int32]int32) {
	for skillId, level := range normalSkillMap {
		pns := skillcommon.CreateSkillObject(skillId, level, nil)
		m.skillList = append(m.skillList, pns)
	}
}

//跳跃技能添加
func (m *PlayerSkillDataManager) addJumpSkill(jumpSkillTemplate *gametemplate.SkillTemplate) {
	pns := skillcommon.CreateSkillObject(int32(jumpSkillTemplate.TemplateId()), jumpSkillTemplate.Lev, nil)
	m.skillList = append(m.skillList, pns)

}

//永久技能添加
func (m *PlayerSkillDataManager) addPersistentSkill(pso *PlayerSkillObject) {
	to := skilltemplate.GetSkillTemplateService().GetSkillTemplate(pso.GetSkillId())
	//职业技能
	if to.FirstType == int32(skilltypes.SkillFirstTypeRole) {
		m.playerSkillObjectMap[pso.SkillId] = pso
	}
	tianFuList := convertTianFu(pso.TianFuMap)
	pns := skillcommon.CreateSkillObject(pso.SkillId, pso.Level, tianFuList)
	m.skillList = append(m.skillList, pns)
}

func (m *PlayerSkillDataManager) GetRoleSkillMap() map[int32]*PlayerSkillObject {
	return m.playerSkillObjectMap
}

//判断玩家能学习该职业技能
func (m *PlayerSkillDataManager) IsValid(skillId int32) bool {
	if skillId <= 0 {
		return false
	}
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(skillId)
	if skillTemplate == nil {
		return false
	}
	//技能类型是否为职业技能
	if skillTemplate.FirstType != int32(skilltypes.SkillFirstTypeRole) {
		return false
	}
	//该技能对玩家是否有效
	role := m.p.GetRole()
	if int32(role) != skillTemplate.ProNeed {
		return false
	}
	return true
}

//是否已拥有该职业技能
func (m *PlayerSkillDataManager) IfSkillExist(skillId int32) bool {
	//该技能是否已获得
	skill := m.GetSkill(skillId)
	if skill == nil {
		return false
	}
	return true
}

func (m *PlayerSkillDataManager) Heartbeat() {

}

//添加新职业技能
func (m *PlayerSkillDataManager) AddSkill(skillId int32) bool {
	flag := m.IsValid(skillId)
	if !flag {
		return false
	}
	flag = m.IfSkillExist(skillId)
	if flag {
		return false
	}
	id, err := idutil.GetId()
	if err != nil {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	pso := NewPlayerSkillObject(m.p)
	pso.Id = id
	pso.PlayerId = m.p.GetId()
	pso.SkillId = skillId
	pso.Level = int32(1)
	pso.TianFuMap = make(map[int32]int32)
	pso.CreateTime = now
	pso.SetModified()
	m.addPersistentSkill(pso)
	m.p.ChangeDynamicSkill(skillId, pso.Level)
	return true
}

//能升级的技能
func (m *PlayerSkillDataManager) CanUpgradeRoleSkills() map[int32]int32 {
	skillsMap := make(map[int32]int32)
	for _, skillInfo := range m.playerSkillObjectMap {
		skillId := skillInfo.SkillId
		level := skillInfo.Level
		pLevel := m.p.GetLevel()
		limit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeSkillLimit)
		limitLevel := pLevel * limit
		if level >= limitLevel {
			continue
		}
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(skillId)
		skillLevelTemplate := skillTemplate.GetSkillByLevel(level)
		if skillLevelTemplate.NextId == 0 {
			continue
		}
		skillsMap[skillId] = level
	}
	if len(skillsMap) <= 0 {
		return nil
	}
	return skillsMap
}

//能否升级
func (m *PlayerSkillDataManager) IfCanUpgradeSkill(skillId int32, level int32) bool {
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(skillId)
	if skillTemplate == nil {
		return false
	}
	skillInfo := m.GetSkill(skillId)
	if skillInfo == nil {
		return false
	}
	pLevel := m.p.GetLevel()
	limit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeSkillLimit)
	limitLevel := pLevel * limit
	if level >= limitLevel {
		return false
	}

	skillLevelTemplate := skillTemplate.GetSkillByLevel(level)
	if skillLevelTemplate.NextId == 0 {
		return false
	}

	return true
}

//职业技能的总等级
func (m *PlayerSkillDataManager) GetTotaolLevel() (totalLevel int32) {
	totalLevel = 0
	for _, skillObj := range m.playerSkillObjectMap {
		totalLevel += skillObj.Level
	}
	return
}

//升级技能
func (m *PlayerSkillDataManager) UpgradeSkill(skillId int32, level int32) bool {
	flag := m.IfCanUpgradeSkill(skillId, level)
	if !flag {
		return false
	}
	skillInfo := m.GetSkill(skillId)
	skillInfo.Level += 1
	now := global.GetGame().GetTimeService().Now()
	skillInfo.UpdateTime = now
	skillInfo.SetModified()
	m.p.ChangeDynamicSkill(skillId, skillInfo.Level)

	gameevent.Emit(skilleventtypes.EventTypeSkillUpgrade, m.p, skillId)
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, m.p, nil)
	return true
}

//升级所有能升级的技能
func (m *PlayerSkillDataManager) UpgradeSkillAll(skillMap map[int32]int32) bool {
	for skillId, level := range skillMap {
		flag := m.UpgradeSkill(skillId, level)
		if !flag {
			return false
		}
	}
	return true
}

func (m *PlayerSkillDataManager) SkillCdTime(skillTypeId int32) {
	now := global.GetGame().GetTimeService().Now()
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillTypeId)
	if skillTemplate == nil {
		return
	}
	skillCdObj, exist := m.playerSkillCdObjectMap[skillTemplate.TypeId]
	if exist {
		skillCdObj.LastTime = now
		skillCdObj.UpdateTime = now
		skillCdObj.SetModified()
		return
	}

	id, _ := idutil.GetId()
	psco := NewPlayerSkillCdObject(m.p)
	psco.Id = id
	psco.PlayerId = m.p.GetId()
	psco.SkillId = skillTemplate.TypeId
	psco.LastTime = now
	psco.CreateTime = now
	m.playerSkillCdObjectMap[skillTypeId] = psco
	psco.SetModified()
}

func (m *PlayerSkillDataManager) GetSkillCdMap() map[int32]*PlayerSkillCdObject {
	return m.playerSkillCdObjectMap
}

//玩家所有职业技能是否达到满级
func (m *PlayerSkillDataManager) IfAllRoleFull() (flag bool) {
	for _, skillInfo := range m.playerSkillObjectMap {
		skillId := skillInfo.SkillId
		level := skillInfo.Level
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(skillId)
		skillLevelTemplate := skillTemplate.GetSkillByLevel(level)
		if skillLevelTemplate.NextId != 0 {
			return
		}
	}
	flag = true
	return
}

func (m *PlayerSkillDataManager) GetSkillTianFuLevel(skillId int32, tianFuId int32) (level int32, flag bool) {
	skillInfo, ok := m.playerSkillObjectMap[skillId]
	if !ok {
		return
	}
	level, ok = skillInfo.TianFuMap[tianFuId]
	if ok {
		flag = true
		return
	}
	return
}

func (m *PlayerSkillDataManager) HasedTianFuAwaken(skillId int32, tianFuId int32) (flag bool) {
	skillInfo, ok := m.playerSkillObjectMap[skillId]
	if !ok {
		return
	}
	_, ok = skillInfo.TianFuMap[tianFuId]
	if ok {
		flag = true
		return
	}
	return
}

func (m *PlayerSkillDataManager) TianFuAwaken(skillId int32, tianFuId int32) (flag bool) {
	tianFuTemplate := skilltemplate.GetSkillTemplateService().GetSkillTianFuTemplate(skillId, tianFuId)
	if tianFuTemplate == nil {
		return
	}
	hasedFlag := m.HasedTianFuAwaken(skillId, tianFuId)
	if hasedFlag {
		return
	}
	skillInfo, ok := m.playerSkillObjectMap[skillId]
	if !ok {
		return
	}

	parentTianFuTemplate := skilltemplate.GetSkillTemplateService().GetSkillParentTianFuTemplate(skillId, tianFuId)
	if parentTianFuTemplate != nil &&
		!m.HasedTianFuAwaken(skillId, int32(parentTianFuTemplate.TemplateId())) {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	skillInfo.TianFuMap[tianFuId] = 1
	skillInfo.UpdateTime = now
	skillInfo.SetModified()
	skillObj := m.p.GetSkill(skillId)
	skillObj.SetTianFuLevel(tianFuId, 1)
	gameevent.Emit(skilleventtypes.EventTypeSkillTianFuAwaken, m.p, tianFuId)
	flag = true
	return
}

func (m *PlayerSkillDataManager) TianFuUpgrade(skillId int32, tianFuId int32) (flag bool) {
	tianFuTemplate := skilltemplate.GetSkillTemplateService().GetSkillTianFuTemplate(skillId, tianFuId)
	if tianFuTemplate == nil {
		return
	}
	hasedFlag := m.HasedTianFuAwaken(skillId, tianFuId)
	if !hasedFlag {
		return
	}
	skillInfo, ok := m.playerSkillObjectMap[skillId]
	if !ok {
		return
	}
	curLevel, ok := skillInfo.TianFuMap[tianFuId]
	if !ok {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	skillInfo.TianFuMap[tianFuId] = curLevel + 1
	skillInfo.UpdateTime = now
	skillInfo.SetModified()
	skillObj := m.p.GetSkill(skillId)
	skillObj.SetTianFuLevel(tianFuId, curLevel+1)
	eventData := skilleventtypes.CreateSkillTianFuUpgradeEventData(skillId, tianFuId)
	gameevent.Emit(skilleventtypes.EventTypeSkillTianFuUpgrade, m.p, eventData)
	flag = true
	return
}

//gm 改变职业技能等级 仅gm使用
func (m *PlayerSkillDataManager) ChangeDynamicLevel(skillId int32, level int32) bool {
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(skillId)
	if skillTemplate == nil {
		return false
	}
	skillInfo := m.GetSkill(skillId)
	if skillInfo == nil {
		return false
	}
	pLevel := m.p.GetLevel()
	limit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeSkillLimit)
	limitLevel := pLevel * limit
	if level >= limitLevel {
		return false
	}

	skillLevelTemplate := skillTemplate.GetSkillByLevel(level)
	if skillLevelTemplate == nil {
		return false
	}
	skillInfo.Level = level
	now := global.GetGame().GetTimeService().Now()
	skillInfo.UpdateTime = now
	skillInfo.SetModified()
	return true
}

//gm 技能等级 仅gm使用
func (m *PlayerSkillDataManager) GmSetSkillClear(obj *PlayerSkillObject) {
	now := global.GetGame().GetTimeService().Now()
	obj.Level = 1
	obj.UpdateTime = now
	obj.SetModified()
	return
}

func CreatePlayerSkillDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerSkillDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerSkillDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerSkillDataManager))
}
