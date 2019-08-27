package juexue

import (
	"fgame/fgame/core/template"
	jxtypes "fgame/fgame/game/juexue/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sync"
)

//绝学接口处理
type JueXueService interface {
	//获取绝学模板
	GetJueXueByTypeAndLevel(typ jxtypes.JueXueType, insight jxtypes.JueXueStageType, level int32) *gametemplate.JueXueTemplate
	GetTypeIsValid(typ jxtypes.JueXueType) bool
	GetJueXueTemplate(id int32) *gametemplate.JueXueTemplate
	GetSkillId(tage jxtypes.JueXueStageType, typ jxtypes.JueXueType, level int32) (skillId int32)
	//获取最大等级
	GetJueXueMaxLevel(typ jxtypes.JueXueType, stage jxtypes.JueXueStageType) *gametemplate.JueXueTemplate
}

type jueXueService struct {
	//绝学模板
	jueXueTemplateMap map[int32]*gametemplate.JueXueTemplate
	//绝学配置
	jueXueLevelTemplateMap map[jxtypes.JueXueType]map[jxtypes.JueXueStageType]map[int32]*gametemplate.JueXueTemplate
	//绝学最大等级
	jueXueMaxLevelMap map[jxtypes.JueXueType]map[jxtypes.JueXueStageType]*gametemplate.JueXueTemplate
}

//初始化
func (jxs *jueXueService) init() error {
	jxs.jueXueTemplateMap = make(map[int32]*gametemplate.JueXueTemplate)
	jxs.jueXueLevelTemplateMap = make(map[jxtypes.JueXueType]map[jxtypes.JueXueStageType]map[int32]*gametemplate.JueXueTemplate)
	jxs.jueXueMaxLevelMap = make(map[jxtypes.JueXueType]map[jxtypes.JueXueStageType]*gametemplate.JueXueTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.JueXueTemplate)(nil))
	for _, templateObject := range templateMap {
		jxTemplate, _ := templateObject.(*gametemplate.JueXueTemplate)

		jxs.jueXueTemplateMap[int32(jxTemplate.TemplateId())] = jxTemplate

		jueXueTemplateTypMap, ok := jxs.jueXueLevelTemplateMap[jxTemplate.GetType()]
		if !ok {
			jueXueTemplateTypMap = make(map[jxtypes.JueXueStageType]map[int32]*gametemplate.JueXueTemplate)
			jxs.jueXueLevelTemplateMap[jxTemplate.GetType()] = jueXueTemplateTypMap
		}

		jueXueTemplateInsightMap, ok := jueXueTemplateTypMap[jxTemplate.GetInsight()]
		if !ok {
			jueXueTemplateInsightMap = make(map[int32]*gametemplate.JueXueTemplate)
			jueXueTemplateTypMap[jxTemplate.GetInsight()] = jueXueTemplateInsightMap
		}
		jueXueTemplateInsightMap[jxTemplate.Level] = jxTemplate

		jxs.initJueXueMaxLevel(jxTemplate)
	}

	//校验level=1
	for _, jueXueTypeTemplateMap := range jxs.jueXueLevelTemplateMap {
		for _, jueXueStageTemplateMap := range jueXueTypeTemplateMap {
			_, ok := jueXueStageTemplateMap[1]
			if !ok {
				return fmt.Errorf("juexueService: level exist 1 should be ok")
			}
		}
	}
	return nil
}

func (jxs *jueXueService) initJueXueMaxLevel(temp *gametemplate.JueXueTemplate) {
	stageTemplateMap, ok := jxs.jueXueMaxLevelMap[temp.GetType()]
	if !ok {
		stageTemplateMap = make(map[jxtypes.JueXueStageType]*gametemplate.JueXueTemplate)
		jxs.jueXueMaxLevelMap[temp.GetType()] = stageTemplateMap
	}
	levelTemplate, ok := stageTemplateMap[temp.GetInsight()]
	if !ok {
		stageTemplateMap[temp.GetInsight()] = temp
		return
	}

	if temp.Level > levelTemplate.Level {
		stageTemplateMap[temp.GetInsight()] = temp
	}
	return
}

//获取最大等级
func (jxs *jueXueService) GetJueXueMaxLevel(typ jxtypes.JueXueType, stage jxtypes.JueXueStageType) *gametemplate.JueXueTemplate {
	stageTemplateMap, ok := jxs.jueXueMaxLevelMap[typ]
	if !ok {
		return nil
	}
	levelTemplate, ok := stageTemplateMap[stage]
	if !ok {
		return nil
	}
	return levelTemplate
}

//获取技能
func (jxs *jueXueService) GetSkillId(stage jxtypes.JueXueStageType, typ jxtypes.JueXueType, level int32) (skillId int32) {
	skillId = int32(0)
	to := jxs.GetJueXueByTypeAndLevel(typ, stage, level)
	if to != nil {
		skillId = to.Skill
	}
	return
}

//获取绝学配置
func (jxs *jueXueService) GetJueXueByTypeAndLevel(typ jxtypes.JueXueType, insight jxtypes.JueXueStageType, level int32) *gametemplate.JueXueTemplate {
	jueXueTemplateTypMap, ok := jxs.jueXueLevelTemplateMap[typ]
	if !ok {
		return nil
	}
	jueXueTemplateInsightMap, ok := jueXueTemplateTypMap[insight]
	if !ok {
		return nil
	}
	to, ok := jueXueTemplateInsightMap[level]
	if !ok {
		return nil
	}

	return to
}

func (jxs *jueXueService) GetJueXueTemplate(id int32) *gametemplate.JueXueTemplate {
	to, ok := jxs.jueXueTemplateMap[id]
	if !ok {
		return nil
	}
	return to
}

func (jxs *jueXueService) GetTypeIsValid(typ jxtypes.JueXueType) bool {
	_, ok := jxs.jueXueLevelTemplateMap[typ]
	if !ok {
		return false
	}
	return true
}

var (
	once sync.Once
	cs   *jueXueService
)

func Init() (err error) {
	once.Do(func() {
		cs = &jueXueService{}
		err = cs.init()
	})
	return err
}

func GetJueXueService() JueXueService {
	return cs
}
