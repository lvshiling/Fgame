package template

import (
	"fgame/fgame/core/template"
	fushitypes "fgame/fgame/game/fushi/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

// 八卦符石模板接口
type FuShiTemplateService interface {
	// 通过符石id和符石等级获取符石等级模板
	GetFuShiLevelByFuShiTypeAndLevel(index fushitypes.FuShiType, level int32) *gametemplate.FuShiLevelTemplate
	// 通过符石id获取符石模板
	GetFuShiByFuShiType(index fushitypes.FuShiType) *gametemplate.FuShiTemplate
}

type fushiTemplateService struct {
	// 八卦符石配置
	fushiLevelMap map[fushitypes.FuShiType]map[int32]*gametemplate.FuShiLevelTemplate
	// 八卦符石常量配置
	fushiMap map[fushitypes.FuShiType]*gametemplate.FuShiTemplate
}

func (f *fushiTemplateService) init() (err error) {
	f.fushiLevelMap = make(map[fushitypes.FuShiType]map[int32]*gametemplate.FuShiLevelTemplate)
	f.fushiMap = make(map[fushitypes.FuShiType]*gametemplate.FuShiTemplate)

	// 符石配置
	tempMap := template.GetTemplateService().GetAll((*gametemplate.FuShiLevelTemplate)(nil))
	for _, temp := range tempMap {
		fushiTemp, _ := temp.(*gametemplate.FuShiLevelTemplate)

		subMap, ok := f.fushiLevelMap[fushiTemp.GetFuShiType()]
		if !ok {
			subMap = make(map[int32]*gametemplate.FuShiLevelTemplate)
			f.fushiLevelMap[fushiTemp.GetFuShiType()] = subMap
		}

		subMap[fushiTemp.Level] = fushiTemp
	}

	// 符石常量配置
	tempConstantMap := template.GetTemplateService().GetAll((*gametemplate.FuShiTemplate)(nil))
	for _, temp := range tempConstantMap {
		constantTemp, _ := temp.(*gametemplate.FuShiTemplate)
		f.fushiMap[constantTemp.GetFuShiType()] = constantTemp
	}

	return
}

func (f *fushiTemplateService) GetFuShiLevelByFuShiTypeAndLevel(index fushitypes.FuShiType, level int32) *gametemplate.FuShiLevelTemplate {
	subMap, ok := f.fushiLevelMap[index]
	if !ok {
		return nil
	}
	temp, ok := subMap[level]
	if !ok {
		return nil
	}

	return temp
}

func (f *fushiTemplateService) GetFuShiByFuShiType(index fushitypes.FuShiType) (temp *gametemplate.FuShiTemplate) {
	temp, _ = f.fushiMap[index]
	return
}

var (
	once  sync.Once
	fushi *fushiTemplateService
)

func Init() (err error) {
	once.Do(func() {
		fushi = &fushiTemplateService{}
		err = fushi.init()
	})

	return err
}

func GetFuShiTemplateService() FuShiTemplateService {
	return fushi
}
