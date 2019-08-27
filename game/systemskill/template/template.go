package template

import (
	"fgame/fgame/core/template"
	sysskilltypes "fgame/fgame/game/systemskill/types"
	systemskilltypes "fgame/fgame/game/systemskill/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

// 系统技能通用接口
type SystemSkillCommonTemplate interface {
	TemplateId() int
	GetType() systemskilltypes.SystemSkillType
	GetSubType() systemskilltypes.SystemSkillSubType
	GetNeedItemMap() map[int32]int32
	// field
	GetSkillId() int32
	GetNumber() int32
	GetCostGold() int32
	GetCostSilver() int32
	GetNextId() int32
	GetLevel() int32
	GetNeedEquipQuality() int32
	GetNeedEquipCount() int32
}

//系统技能接口处理
type SystemSkillTemplateService interface {
	//获取系统技能模板
	GetSystemSkillTemplateByTypeAndLevel(typ sysskilltypes.SystemSkillType, subType sysskilltypes.SystemSkillSubType, level int32) SystemSkillCommonTemplate
	GetSystemSkillTemplate(id int32) SystemSkillCommonTemplate
	GetSkillId(typ sysskilltypes.SystemSkillType, subType sysskilltypes.SystemSkillSubType, level int32) (skillId int32)
}

type systemSkillTemplateService struct {
	//系统技能模板
	sysSkillTemplateMap map[int32]SystemSkillCommonTemplate
	//系统技能配置
	sysSkillTypeTemplateMap map[sysskilltypes.SystemSkillType]map[sysskilltypes.SystemSkillSubType]map[int32]SystemSkillCommonTemplate
}

//初始化
func (s *systemSkillTemplateService) init() error {
	s.sysSkillTemplateMap = make(map[int32]SystemSkillCommonTemplate)
	s.sysSkillTypeTemplateMap = make(map[sysskilltypes.SystemSkillType]map[sysskilltypes.SystemSkillSubType]map[int32]SystemSkillCommonTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.SystemSkillTemplate)(nil))
	for _, templateObject := range templateMap {
		skTemplate, _ := templateObject.(*gametemplate.SystemSkillTemplate)

		s.sysSkillTemplateMap[int32(skTemplate.TemplateId())] = skTemplate

		typ := skTemplate.GetType()
		subType := skTemplate.GetSubType()
		level := skTemplate.Level

		skTypeTemplateMap, ok := s.sysSkillTypeTemplateMap[typ]
		if !ok {
			skTypeTemplateMap = make(map[sysskilltypes.SystemSkillSubType]map[int32]SystemSkillCommonTemplate)
			s.sysSkillTypeTemplateMap[typ] = skTypeTemplateMap
		}
		skLevelTeamplateMap, ok := skTypeTemplateMap[subType]
		if !ok {
			skLevelTeamplateMap = make(map[int32]SystemSkillCommonTemplate)
			skTypeTemplateMap[subType] = skLevelTeamplateMap
		}
		skLevelTeamplateMap[level] = skTemplate
	}

	// 圣痕技能
	shenghenTempMap := template.GetTemplateService().GetAll((*gametemplate.SystemSkillShengHenTemplate)(nil))
	for _, templateObject := range shenghenTempMap {
		skTemplate, _ := templateObject.(*gametemplate.SystemSkillShengHenTemplate)

		s.sysSkillTemplateMap[int32(skTemplate.TemplateId())] = skTemplate

		typ := skTemplate.GetType()
		subType := skTemplate.GetSubType()
		level := skTemplate.Level

		skTypeTemplateMap, ok := s.sysSkillTypeTemplateMap[typ]
		if !ok {
			skTypeTemplateMap = make(map[sysskilltypes.SystemSkillSubType]map[int32]SystemSkillCommonTemplate)
			s.sysSkillTypeTemplateMap[typ] = skTypeTemplateMap
		}
		skLevelTeamplateMap, ok := skTypeTemplateMap[subType]
		if !ok {
			skLevelTeamplateMap = make(map[int32]SystemSkillCommonTemplate)
			skTypeTemplateMap[subType] = skLevelTeamplateMap
		}
		skLevelTeamplateMap[level] = skTemplate
	}

	return nil
}

func (s *systemSkillTemplateService) GetSkillId(typ sysskilltypes.SystemSkillType, subType sysskilltypes.SystemSkillSubType, level int32) (skillId int32) {
	skillId = int32(0)
	to := s.GetSystemSkillTemplateByTypeAndLevel(typ, subType, level)
	if to != nil {
		skillId = to.GetSkillId()
	}
	return skillId
}

//获取系统技能配置
func (s *systemSkillTemplateService) GetSystemSkillTemplateByTypeAndLevel(typ sysskilltypes.SystemSkillType, subType sysskilltypes.SystemSkillSubType, level int32) SystemSkillCommonTemplate {
	skTypeTemplateMap, ok := s.sysSkillTypeTemplateMap[typ]
	if !ok {
		return nil
	}
	skLevelMap, ok := skTypeTemplateMap[subType]
	if !ok {
		return nil
	}
	to, ok := skLevelMap[level]
	if !ok {
		return nil
	}
	return to
}

func (s *systemSkillTemplateService) GetSystemSkillTemplate(id int32) SystemSkillCommonTemplate {
	to, ok := s.sysSkillTemplateMap[id]
	if !ok {
		return nil
	}
	return to
}

var (
	once sync.Once
	cs   *systemSkillTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &systemSkillTemplateService{}
		err = cs.init()
	})
	return err
}

func GetSystemSkillTemplateService() SystemSkillTemplateService {
	return cs
}
