package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type CollectTemplateService interface {
	GetMiZang(id int32) *gametemplate.BossMiZangTemplate
}

type collectTemplateService struct {
	miZangMap map[int32]*gametemplate.BossMiZangTemplate
}

func (t *collectTemplateService) init() error {
	t.miZangMap = make(map[int32]*gametemplate.BossMiZangTemplate)
	tempMiZangMap := template.GetTemplateService().GetAll((*gametemplate.BossMiZangTemplate)(nil))
	for _, tempMiZangTemplate := range tempMiZangMap {
		miZangTemplate := tempMiZangTemplate.(*gametemplate.BossMiZangTemplate)
		t.miZangMap[int32(miZangTemplate.TemplateId())] = miZangTemplate
	}
	return nil
}

func (t *collectTemplateService) GetMiZang(id int32) *gametemplate.BossMiZangTemplate {
	temp, ok := t.miZangMap[id]
	if !ok {
		return nil
	}
	return temp
}

var (
	once sync.Once
	cs   *collectTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &collectTemplateService{}
		err = cs.init()
	})
	return err
}

func GetCollectTemplateService() CollectTemplateService {
	return cs
}
