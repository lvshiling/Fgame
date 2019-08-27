package dragon

import (
	"fgame/fgame/core/template"
	"sync"

	gametemplate "fgame/fgame/game/template"
)

type DragonService interface {
	//获取神龙现世配置
	GetDragonTemplate(id int32) *gametemplate.DragonTemplate
}

type dragonService struct {
	//神龙现世配置
	dragonMap map[int32]*gametemplate.DragonTemplate
}

//初始化
func (ds *dragonService) init() error {
	ds.dragonMap = make(map[int32]*gametemplate.DragonTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.DragonTemplate)(nil))
	for _, templateObject := range templateMap {
		dragonTemplate, _ := templateObject.(*gametemplate.DragonTemplate)
		ds.dragonMap[int32(dragonTemplate.TemplateId())] = dragonTemplate
	}

	return nil
}

//获取神龙现世配置
func (ds *dragonService) GetDragonTemplate(id int32) *gametemplate.DragonTemplate {
	to, ok := ds.dragonMap[id]
	if !ok {
		return nil
	}
	return to

}

var (
	once sync.Once
	cs   *dragonService
)

func Init() (err error) {
	once.Do(func() {
		cs = &dragonService{}
		err = cs.init()
	})
	return err
}

func GetDragonService() DragonService {
	return cs
}
