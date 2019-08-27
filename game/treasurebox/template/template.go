package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//宝箱模板接口处理
type BoxTemplateService interface {
	//获取宝箱模板
	GetBoxTemplate(id int32) *gametemplate.BoxTemplate
}

type boxTemplateService struct {
	boxTemplateMap map[int32]*gametemplate.BoxTemplate
}

//初始化
func (bts *boxTemplateService) init() (err error) {
	bts.boxTemplateMap = make(map[int32]*gametemplate.BoxTemplate)
	//赋值boxTemplateMap
	templateMap := template.GetTemplateService().GetAll((*gametemplate.BoxTemplate)(nil))
	for _, templateObject := range templateMap {
		boxTemplate, _ := templateObject.(*gametemplate.BoxTemplate)
		bts.boxTemplateMap[int32(boxTemplate.TemplateId())] = boxTemplate
	}

	return nil
}

//获取宝箱模板id
func (bts *boxTemplateService) GetBoxTemplate(id int32) *gametemplate.BoxTemplate {
	to, ok := bts.boxTemplateMap[id]
	if !ok {
		return nil
	}
	return to
}

var (
	once sync.Once
	cs   *boxTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &boxTemplateService{}
		err = cs.init()
	})
	return err
}

func GetBoxTemplateService() BoxTemplateService {
	return cs
}
