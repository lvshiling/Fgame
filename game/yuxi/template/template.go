package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"fmt"

	"sync"
)

type YuXiTemplateService interface {
	//获取玉玺之战常量模板
	GetYuXiConstTemplate() *gametemplate.YuXiConstantTemplate
}

type yuXiTemplateService struct {
	//玉玺之战常量模板
	yuXiConstTemplate *gametemplate.YuXiConstantTemplate
}

//初始化
func (t *yuXiTemplateService) init() error {
	//常量配置
	templateMap := template.GetTemplateService().GetAll((*gametemplate.YuXiConstantTemplate)(nil))
	if len(templateMap) != 1 {
		return fmt.Errorf("玉玺之战常量配置应该有且只有一条")
	}
	for _, to := range templateMap {
		t.yuXiConstTemplate = to.(*gametemplate.YuXiConstantTemplate)
		break
	}

	return nil
}

//获取玉玺之战常量模板
func (t *yuXiTemplateService) GetYuXiConstTemplate() *gametemplate.YuXiConstantTemplate {
	return t.yuXiConstTemplate
}

var (
	once sync.Once
	cs   *yuXiTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &yuXiTemplateService{}
		err = cs.init()
	})
	return err
}

func GetYuXiTemplateService() YuXiTemplateService {
	return cs
}
