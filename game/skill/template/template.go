package template

import (
	"fgame/fgame/core/template"
	skilltypes "fgame/fgame/game/skill/types"
	gametemplate "fgame/fgame/game/template"
	"math/rand"
	"sync"
)

type SkillTemplateService interface {
	GetSkillTemplate(id int32) *gametemplate.SkillTemplate
	GetSkillTemplateByTypeAndLevel(typ int32, lev int32) *gametemplate.SkillTemplate
	GetSkillTemplateByType(typ int32) *gametemplate.SkillTemplate
	GetAllSkillTemplates() map[int32]*gametemplate.SkillTemplate
	GetMaxLevel(typeId int32) int32
	RandomSkillTemplate(firstType skilltypes.SkillFirstType) *gametemplate.SkillTemplate
	GetSkillTianFuTemplate(skillId int32, tianFuId int32) *gametemplate.TianFuTemplate
	GetSkillParentTianFuTemplate(skillId int32, tianFuId int32) *gametemplate.TianFuTemplate
}
type skillTemplateService struct {
	skillMap map[int32]*gametemplate.SkillTemplate
	//静态
	skillTypeMap         map[int32]map[int32]*gametemplate.SkillTemplate
	allSkillTemplateList map[int32]*gametemplate.SkillTemplate
	maxLevelMap          map[int32]int32
	skillFirstTypeMap    map[skilltypes.SkillFirstType][]*gametemplate.SkillTemplate
}

func (ss *skillTemplateService) init() error {
	ss.skillMap = make(map[int32]*gametemplate.SkillTemplate)
	ss.skillTypeMap = make(map[int32]map[int32]*gametemplate.SkillTemplate)
	ss.allSkillTemplateList = make(map[int32]*gametemplate.SkillTemplate)
	ss.maxLevelMap = make(map[int32]int32)
	ss.skillFirstTypeMap = make(map[skilltypes.SkillFirstType][]*gametemplate.SkillTemplate)

	for _, tempSkillTemplate := range template.GetTemplateService().GetAll((*gametemplate.SkillTemplate)(nil)) {
		skillTemplate := tempSkillTemplate.(*gametemplate.SkillTemplate)
		ss.skillMap[int32(skillTemplate.TemplateId())] = skillTemplate

		tempSkillMap, ok := ss.skillTypeMap[skillTemplate.TypeId]
		if !ok {
			tempSkillMap = make(map[int32]*gametemplate.SkillTemplate)
			ss.skillTypeMap[skillTemplate.TypeId] = tempSkillMap
		}
		tempSkillMap[skillTemplate.Lev] = skillTemplate
		if skillTemplate.Lev == 1 {
			ss.allSkillTemplateList[skillTemplate.TypeId] = skillTemplate
			skillFirstList := ss.skillFirstTypeMap[skillTemplate.GetSkillFirstType()]
			skillFirstList = append(skillFirstList, skillTemplate)
			ss.skillFirstTypeMap[skillTemplate.GetSkillFirstType()] = skillFirstList
		}
		if skillTemplate.IsStatic() {
			//获取最高等级
			lev := ss.maxLevelMap[skillTemplate.TypeId]
			if skillTemplate.Lev > lev {
				ss.maxLevelMap[skillTemplate.TypeId] = lev
			}
		} else {
			ss.maxLevelMap[skillTemplate.TypeId] = skillTemplate.GetMaxLevel()
		}

	}
	return nil
}

func (ss *skillTemplateService) GetSkillTemplate(id int32) *gametemplate.SkillTemplate {
	skillTemplate, ok := ss.skillMap[id]
	if !ok {
		return nil
	}
	return skillTemplate
}

func (ss *skillTemplateService) GetSkillTemplateByTypeAndLevel(typ int32, lev int32) *gametemplate.SkillTemplate {
	tempSkillMap, ok := ss.skillTypeMap[typ]
	if !ok {
		return nil
	}
	firstSkillTemplate, ok := tempSkillMap[1]
	if !ok {
		return nil
	}
	if firstSkillTemplate.IsStatic() {
		return tempSkillMap[lev]
	}
	return firstSkillTemplate
}

func (ss *skillTemplateService) GetSkillTemplateByType(typ int32) *gametemplate.SkillTemplate {
	tempSkillMap, ok := ss.skillTypeMap[typ]
	if !ok {
		return nil
	}
	firstSkillTemplate, ok := tempSkillMap[1]
	if !ok {
		return nil
	}

	return firstSkillTemplate
}

func (ss *skillTemplateService) GetAllSkillTemplates() map[int32]*gametemplate.SkillTemplate {
	return ss.allSkillTemplateList
}

func (ss *skillTemplateService) GetMaxLevel(typeId int32) int32 {
	return ss.maxLevelMap[typeId]
}

func (ss *skillTemplateService) RandomSkillTemplate(firstType skilltypes.SkillFirstType) *gametemplate.SkillTemplate {
	skillTemplateList := ss.skillFirstTypeMap[firstType]
	num := len(skillTemplateList)
	index := rand.Intn(num)
	return skillTemplateList[index]
}

func (ss *skillTemplateService) GetSkillTianFuTemplate(skillId int32, tianFuId int32) *gametemplate.TianFuTemplate {
	skillTemplate, ok := ss.skillMap[skillId]
	if !ok {
		return nil
	}
	tianFuTemplate := skillTemplate.GetTianFuTemplate(tianFuId)
	if tianFuTemplate == nil {
		return nil
	}
	return tianFuTemplate
}

func (ss *skillTemplateService) GetSkillParentTianFuTemplate(skillId int32, tianFuId int32) *gametemplate.TianFuTemplate {
	skillTemplate, ok := ss.skillMap[skillId]
	if !ok {
		return nil
	}

	tianFuTempalte := skillTemplate.GetTianFuTemplate(tianFuId)
	if tianFuTempalte != nil && tianFuTempalte.ParentId != 0 {
		return skillTemplate.GetTianFuTemplate(tianFuTempalte.ParentId)
	}
	return nil
}

var (
	once sync.Once
	cs   *skillTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &skillTemplateService{}
		err = cs.init()
	})
	return err
}

func GetSkillTemplateService() SkillTemplateService {
	return cs
}
