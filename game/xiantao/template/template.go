package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	xiantaotypes "fgame/fgame/game/xiantao/types"
	"sync"
)

type XianTaoTemplateService interface {
	//获取仙桃大会常量模板
	GetXianTaoConstTemplate() *gametemplate.XianTaoConstantTemplate
	//根据次数获取仙桃大会次数模板
	GetXianTaoTimesTempByTimes(val int32) *gametemplate.XianTaoTimesTemplate
	//仙桃大会模板
	GetXianTaoTempByArg(typ xiantaotypes.XianTaoType, val int32) *gametemplate.XianTaoTemplate
}

type xianTaoTemplateService struct {
	//仙桃大会常量模板
	xianTaoConstTemplate *gametemplate.XianTaoConstantTemplate
	//仙桃大会次数模板
	xianTaoTimesTempMap map[int32]*gametemplate.XianTaoTimesTemplate
	//仙桃大会模板
	xianTaoTempMap map[xiantaotypes.XianTaoType]map[int]*gametemplate.XianTaoTemplate
}

//初始化
func (xtts *xianTaoTemplateService) init() error {
	//常量配置
	templateMap := template.GetTemplateService().GetAll((*gametemplate.XianTaoConstantTemplate)(nil))
	for _, templateObject := range templateMap {
		xtts.xianTaoConstTemplate, _ = templateObject.(*gametemplate.XianTaoConstantTemplate)
		break
	}

	//次数模板
	xtts.xianTaoTimesTempMap = make(map[int32]*gametemplate.XianTaoTimesTemplate)
	xianTaoTimesTemplateMap := template.GetTemplateService().GetAll((*gametemplate.XianTaoTimesTemplate)(nil))
	for _, templateObject := range xianTaoTimesTemplateMap {
		xianTaoTimesTemplate, _ := templateObject.(*gametemplate.XianTaoTimesTemplate)
		xtts.xianTaoTimesTempMap[xianTaoTimesTemplate.Times] = xianTaoTimesTemplate
	}

	//仙桃大会模板
	xtts.xianTaoTempMap = make(map[xiantaotypes.XianTaoType]map[int]*gametemplate.XianTaoTemplate)
	xianTaoTemplateMap := template.GetTemplateService().GetAll((*gametemplate.XianTaoTemplate)(nil))
	for _, templateObject := range xianTaoTemplateMap {
		xianTaoTemplate, _ := templateObject.(*gametemplate.XianTaoTemplate)
		tempM, ok := xtts.xianTaoTempMap[xianTaoTemplate.GetXianTaoType()]
		if !ok {
			tempM = make(map[int]*gametemplate.XianTaoTemplate)
			xtts.xianTaoTempMap[xianTaoTemplate.GetXianTaoType()] = tempM
		}
		tempM[xianTaoTemplate.TemplateId()] = xianTaoTemplate
	}

	return nil
}

//获取仙桃大会常量模板
func (xtts *xianTaoTemplateService) GetXianTaoConstTemplate() *gametemplate.XianTaoConstantTemplate {
	return xtts.xianTaoConstTemplate
}

//根据次数获取仙桃大会次数模板
func (xtts *xianTaoTemplateService) GetXianTaoTimesTempByTimes(val int32) *gametemplate.XianTaoTimesTemplate {
	temp, ok := xtts.xianTaoTimesTempMap[val]
	if !ok {
		return nil
	}
	return temp
}

//根据次数获取仙桃大会次数模板
func (xtts *xianTaoTemplateService) GetXianTaoTempByArg(typ xiantaotypes.XianTaoType, val int32) *gametemplate.XianTaoTemplate {
	if val <= 0 {
		return nil
	}
	var curTemp *gametemplate.XianTaoTemplate
	switch typ {
	case xiantaotypes.XianTaoTypeQianNian:
		curTemp = xtts.xianTaoTempMap[typ][1]
	case xiantaotypes.XianTaoTypeBaiNian:
		curTemp = xtts.xianTaoTempMap[typ][101]
	}
	for !curTemp.IsInXianTaoRange(val) {
		nextTemp := curTemp.GetNextTemplate()
		if nextTemp == nil {
			return curTemp
		}
		curTemp = nextTemp
	}
	return curTemp
}

var (
	once sync.Once
	cs   *xianTaoTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &xianTaoTemplateService{}
		err = cs.init()
	})
	return err
}

func GetXianTaoTemplateService() XianTaoTemplateService {
	return cs
}
