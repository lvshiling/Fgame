package funcopen

import (
	"fgame/fgame/core/template"
	funcopentypes "fgame/fgame/game/funcopen/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type FuncOpenService interface {
	GetFuncOpenTemplateById(funcId int32) *gametemplate.ModuleOpenedTemplate
	GetFuncOpenTemplate(funcOpenType funcopentypes.FuncOpenType) *gametemplate.ModuleOpenedTemplate
	GetAll() map[funcopentypes.FuncOpenType]*gametemplate.ModuleOpenedTemplate
	GetManualFuncOpenMap() map[funcopentypes.FuncOpenType]*gametemplate.ModuleOpenedTemplate
}

//快捷缓存
//物品配置的整合
type funcOpenService struct {
	funcOpenMap       map[funcopentypes.FuncOpenType]*gametemplate.ModuleOpenedTemplate
	funcMap           map[int32]*gametemplate.ModuleOpenedTemplate
	manualFuncOpenMap map[funcopentypes.FuncOpenType]*gametemplate.ModuleOpenedTemplate
}

func (fos *funcOpenService) init() error {
	fos.funcOpenMap = make(map[funcopentypes.FuncOpenType]*gametemplate.ModuleOpenedTemplate)
	fos.funcMap = make(map[int32]*gametemplate.ModuleOpenedTemplate)
	fos.manualFuncOpenMap = make(map[funcopentypes.FuncOpenType]*gametemplate.ModuleOpenedTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.ModuleOpenedTemplate)(nil))
	for _, templateObject := range templateMap {
		moduleOpenedTemplate, _ := templateObject.(*gametemplate.ModuleOpenedTemplate)
		fos.funcOpenMap[moduleOpenedTemplate.GetFuncOpenType()] = moduleOpenedTemplate
		fos.funcMap[int32(moduleOpenedTemplate.Id)] = moduleOpenedTemplate
		if moduleOpenedTemplate.GetFuncOpenCheckType() == funcopentypes.FuncOpenCheckTypeManual {
			fos.manualFuncOpenMap[moduleOpenedTemplate.GetFuncOpenType()] = moduleOpenedTemplate
		}
	}

	return nil
}

func (fos *funcOpenService) GetAll() map[funcopentypes.FuncOpenType]*gametemplate.ModuleOpenedTemplate {
	return fos.funcOpenMap
}

func (fos *funcOpenService) GetFuncOpenTemplate(funcOpenType funcopentypes.FuncOpenType) *gametemplate.ModuleOpenedTemplate {
	funcOpenTemplate, ok := fos.funcOpenMap[funcOpenType]
	if !ok {
		return nil
	}
	return funcOpenTemplate
}

func (fos *funcOpenService) GetManualFuncOpenMap() map[funcopentypes.FuncOpenType]*gametemplate.ModuleOpenedTemplate {

	return fos.manualFuncOpenMap
}

func (fos *funcOpenService) GetFuncOpenTemplateById(funcId int32) *gametemplate.ModuleOpenedTemplate {
	funcOpenTemplate, ok := fos.funcMap[funcId]
	if !ok {
		return nil
	}
	return funcOpenTemplate
}

var (
	once sync.Once
	cs   *funcOpenService
)

func Init() (err error) {
	once.Do(func() {
		cs = &funcOpenService{}
		err = cs.init()
	})
	return err
}

func GetFuncOpenService() FuncOpenService {
	return cs
}
