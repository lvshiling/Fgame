package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	tianshutypes "fgame/fgame/game/tianshu/types"
	"sync"
)

//天书配置服务
type TianShuTemplaterService interface {
	GetTianShuTemplate(typ tianshutypes.TianShuType, level int32) *gametemplate.TianShuTemplate
}

type tianshuTemplaterService struct {
	tianshuMap map[tianshutypes.TianShuType]map[int32]*gametemplate.TianShuTemplate
}

//初始化
func (ts *tianshuTemplaterService) init() error {
	ts.tianshuMap = make(map[tianshutypes.TianShuType]map[int32]*gametemplate.TianShuTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.TianShuTemplate)(nil))
	for _, temp := range templateMap {
		tianshuTemplate, _ := temp.(*gametemplate.TianShuTemplate)

		levelMap, ok := ts.tianshuMap[tianshuTemplate.GetTianShuType()]
		if !ok {
			levelMap = make(map[int32]*gametemplate.TianShuTemplate)
			ts.tianshuMap[tianshuTemplate.GetTianShuType()] = levelMap
		}
		_, ok = levelMap[tianshuTemplate.Level]
		if !ok {
			levelMap[tianshuTemplate.Level] = tianshuTemplate
		}
	}

	return nil
}

func (ts *tianshuTemplaterService) GetTianShuTemplate(typ tianshutypes.TianShuType, level int32) *gametemplate.TianShuTemplate {
	levelMap, ok := ts.tianshuMap[typ]
	if !ok {
		return nil
	}

	temp, ok := levelMap[level]
	if !ok {
		return nil
	}

	return temp
}

var (
	once sync.Once
	cs   *tianshuTemplaterService
)

func Init() (err error) {
	once.Do(func() {
		cs = &tianshuTemplaterService{}
		err = cs.init()
	})
	return err
}

func GetTianShuTemplateService() TianShuTemplaterService {
	return cs
}
