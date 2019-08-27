package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//至尊称号接口处理
type TitleDingZhiTemplateService interface {
	//获取至尊称号模板
	GetTitleDingZhiTempalte(titleId int32) *gametemplate.TitleDingZhiTemplate
}

type supremeTitleTemplateService struct {
	//至尊称号模板
	supremeTitleTemplateMap map[int32]*gametemplate.TitleDingZhiTemplate
}

//初始化
func (ts *supremeTitleTemplateService) init() error {
	ts.supremeTitleTemplateMap = make(map[int32]*gametemplate.TitleDingZhiTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.TitleDingZhiTemplate)(nil))
	for _, templateObject := range templateMap {
		supremeTitleTemplate, _ := templateObject.(*gametemplate.TitleDingZhiTemplate)
		ts.supremeTitleTemplateMap[int32(supremeTitleTemplate.TemplateId())] = supremeTitleTemplate
	}
	return nil
}

//获取至尊称号模板
func (ts *supremeTitleTemplateService) GetTitleDingZhiTempalte(titleId int32) *gametemplate.TitleDingZhiTemplate {
	titleTemplate, ok := ts.supremeTitleTemplateMap[titleId]
	if !ok {
		return nil
	}
	return titleTemplate
}

var (
	once sync.Once
	cs   *supremeTitleTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &supremeTitleTemplateService{}
		err = cs.init()
	})
	return err
}

func GetTitleDingZhiTemplateService() TitleDingZhiTemplateService {
	return cs
}
