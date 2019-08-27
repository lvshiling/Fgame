package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type DaLiWanTemplateService interface {
	GetLinShiTemplate(typ int32) *gametemplate.LinShiAttrTemplate
}

type daLiWanTemplateService struct {
	linShiMap map[int32]*gametemplate.LinShiAttrTemplate
}

func (t *daLiWanTemplateService) init() error {
	t.linShiMap = make(map[int32]*gametemplate.LinShiAttrTemplate)
	tempLinShiTemplateMap := template.GetTemplateService().GetAll((*gametemplate.LinShiAttrTemplate)(nil))
	for _, tempLinShiTemplate := range tempLinShiTemplateMap {
		linShiTemplate := tempLinShiTemplate.(*gametemplate.LinShiAttrTemplate)
		t.linShiMap[int32(linShiTemplate.Id)] = linShiTemplate
	}
	return nil
}

func (t *daLiWanTemplateService) GetLinShiTemplate(typ int32) *gametemplate.LinShiAttrTemplate {
	temp, ok := t.linShiMap[typ]
	if !ok {
		return nil
	}
	return temp
}

var (
	once sync.Once
	cs   *daLiWanTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &daLiWanTemplateService{}
		err = cs.init()
	})
	return err
}

func GetDaLiWanTemplateService() DaLiWanTemplateService {
	return cs
}
