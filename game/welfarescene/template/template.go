package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//运营活动副本配置服务
type WelfareSceneTemplateService interface {
	// 运营活动副本配置
	GetWelfareSceneTemplate(id int32) *gametemplate.WelfareSceneTemplate
}

type welfareSceneTemplateService struct {
	fubenMap map[int32]*gametemplate.WelfareSceneTemplate
}

//初始化
func (ts *welfareSceneTemplateService) init() error {
	ts.fubenMap = make(map[int32]*gametemplate.WelfareSceneTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.WelfareSceneTemplate)(nil))
	for _, to := range templateMap {
		fuBenTemp, _ := to.(*gametemplate.WelfareSceneTemplate)
		ts.fubenMap[int32(fuBenTemp.Id)] = fuBenTemp
	}

	return nil
}

func (ts *welfareSceneTemplateService) GetWelfareSceneTemplate(id int32) *gametemplate.WelfareSceneTemplate {
	temp, ok := ts.fubenMap[id]
	if !ok {
		return nil
	}

	return temp
}

var (
	once sync.Once
	cs   *welfareSceneTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &welfareSceneTemplateService{}
		err = cs.init()
	})
	return err
}

func GetWelfareSceneTemplateService() WelfareSceneTemplateService {
	return cs
}
