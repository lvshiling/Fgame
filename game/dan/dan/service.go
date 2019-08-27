package dan

import (
	"fgame/fgame/core/template"
	"sync"

	gametemplate "fgame/fgame/game/template"
)

type DanService interface {
	//获取食丹配置
	GetEatDan(id int) *gametemplate.EatDanTemplate
	//获取炼丹配置
	GetAlchemy(kindId int) *gametemplate.AlchemyTemplate
}

type danService struct {
	//食丹配置
	eatDanMap map[int]*gametemplate.EatDanTemplate
	//炼丹配置
	alchemyMap map[int]*gametemplate.AlchemyTemplate
}

//初始化
func (ds *danService) init() error {
	ds.eatDanMap = make(map[int]*gametemplate.EatDanTemplate)
	ds.alchemyMap = make(map[int]*gametemplate.AlchemyTemplate)
	//食丹
	templateMap := template.GetTemplateService().GetAll((*gametemplate.EatDanTemplate)(nil))
	for _, templateObject := range templateMap {
		eatDanTemplate, _ := templateObject.(*gametemplate.EatDanTemplate)
		ds.eatDanMap[eatDanTemplate.TemplateId()] = eatDanTemplate
	}
	//炼丹
	atemplateMap := template.GetTemplateService().GetAll((*gametemplate.AlchemyTemplate)(nil))

	for _, atemplateObject := range atemplateMap {
		achemyTemplate, _ := atemplateObject.(*gametemplate.AlchemyTemplate)
		ds.alchemyMap[achemyTemplate.TemplateId()] = achemyTemplate
	}

	return nil
}

//获取食丹配置
func (ds *danService) GetEatDan(id int) *gametemplate.EatDanTemplate {
	to, ok := ds.eatDanMap[id]
	if !ok {
		return nil
	}
	return to
}

//获取炼丹配置
func (ds *danService) GetAlchemy(kindId int) *gametemplate.AlchemyTemplate {
	to, ok := ds.alchemyMap[kindId]
	if !ok {
		return nil
	}
	return to
}

var (
	once sync.Once
	cs   *danService
)

func Init() (err error) {
	once.Do(func() {
		cs = &danService{}
		err = cs.init()
	})
	return err
}

func GetDanService() DanService {
	return cs
}
