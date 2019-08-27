package xinfa

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	xinfatypes "fgame/fgame/game/xinfa/types"
	"fmt"
	"sync"
)

//心法接口处理
type XinFaService interface {
	//获取心法模板
	GetXinFaByTypeAndLevel(typ xinfatypes.XinFaType, level int32) *gametemplate.XinFaTemplate
	GetXinFaTemplate(id int32) *gametemplate.XinFaTemplate
	GetSkillId(typ xinfatypes.XinFaType, level int32) (skillId int32)
	//获取最大等级
	GetXinFaMaxLevel(typ xinfatypes.XinFaType) *gametemplate.XinFaTemplate
}

type xinFaService struct {
	//心法模板
	xinFaTemplateMap map[int32]*gametemplate.XinFaTemplate
	//心法配置
	xinFaTypeTemplateMap map[xinfatypes.XinFaType]map[int32]*gametemplate.XinFaTemplate
	//心法最大等级
	xinFaMaxLevelMap map[xinfatypes.XinFaType]*gametemplate.XinFaTemplate
}

//初始化
func (xfs *xinFaService) init() error {
	xfs.xinFaTemplateMap = make(map[int32]*gametemplate.XinFaTemplate)
	xfs.xinFaTypeTemplateMap = make(map[xinfatypes.XinFaType]map[int32]*gametemplate.XinFaTemplate)
	xfs.xinFaMaxLevelMap = make(map[xinfatypes.XinFaType]*gametemplate.XinFaTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.XinFaTemplate)(nil))
	for _, templateObject := range templateMap {
		xfTemplate, _ := templateObject.(*gametemplate.XinFaTemplate)

		xfs.xinFaTemplateMap[int32(xfTemplate.TemplateId())] = xfTemplate

		xinFaLevelTemplateMap, ok := xfs.xinFaTypeTemplateMap[xfTemplate.GetType()]
		if !ok {
			xinFaLevelTemplateMap = make(map[int32]*gametemplate.XinFaTemplate)
			xfs.xinFaTypeTemplateMap[xfTemplate.GetType()] = xinFaLevelTemplateMap
		}
		xinFaLevelTemplateMap[xfTemplate.Level] = xfTemplate
		xfs.initXinFaMaxLevel(xfTemplate)
	}

	//校验level=1
	for _, xinFaLevelTemplateMap := range xfs.xinFaTypeTemplateMap {
		_, ok := xinFaLevelTemplateMap[1]
		if !ok {
			return fmt.Errorf("xinFaService: level exist 1 should be ok")
		}

	}

	return nil
}

func (xfs *xinFaService) initXinFaMaxLevel(temp *gametemplate.XinFaTemplate) {
	levelTemplate, ok := xfs.xinFaMaxLevelMap[temp.GetType()]
	if !ok {
		xfs.xinFaMaxLevelMap[temp.GetType()] = temp
		return
	}

	if temp.Level > levelTemplate.Level {
		xfs.xinFaMaxLevelMap[temp.GetType()] = temp
	}
	return
}

//获取最大等级
func (xfs *xinFaService) GetXinFaMaxLevel(typ xinfatypes.XinFaType) *gametemplate.XinFaTemplate {
	return xfs.xinFaMaxLevelMap[typ]
}

func (xfs *xinFaService) GetSkillId(typ xinfatypes.XinFaType, level int32) (skillId int32) {
	skillId = int32(0)
	to := xfs.GetXinFaByTypeAndLevel(typ, level)
	if to != nil {
		skillId = to.SkillId
	}
	return skillId
}

//获取心法配置
func (xfs *xinFaService) GetXinFaByTypeAndLevel(typ xinfatypes.XinFaType, level int32) *gametemplate.XinFaTemplate {
	xinFaLevelTemplateMap, ok := xfs.xinFaTypeTemplateMap[typ]
	if !ok {
		return nil
	}
	to, ok := xinFaLevelTemplateMap[level]
	if !ok {
		return nil
	}
	return to
}

func (xfs *xinFaService) GetXinFaTemplate(id int32) *gametemplate.XinFaTemplate {
	to, ok := xfs.xinFaTemplateMap[id]
	if !ok {
		return nil
	}
	return to
}

var (
	once sync.Once
	cs   *xinFaService
)

func Init() (err error) {
	once.Do(func() {
		cs = &xinFaService{}
		err = cs.init()
	})
	return err
}

func GetXinFaService() XinFaService {
	return cs
}
