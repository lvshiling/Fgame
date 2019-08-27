package template

import (
	"fgame/fgame/core/template"
	itemtypes "fgame/fgame/game/item/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//红包接口处理
type HongBaoTemplateService interface {
	//获取红包配置
	GetHongBaoByTemplateType(typ itemtypes.ItemHongBaoSubType) *gametemplate.HongBaoTemplate
}

type hongBaoTemplateService struct {
	//红包map
	hongBaoMap map[itemtypes.ItemHongBaoSubType]*gametemplate.HongBaoTemplate
}

//初始化
func (cs *hongBaoTemplateService) init() error {
	cs.hongBaoMap = make(map[itemtypes.ItemHongBaoSubType]*gametemplate.HongBaoTemplate)
	//红包
	templateMap := template.GetTemplateService().GetAll((*gametemplate.HongBaoTemplate)(nil))
	for _, templateObject := range templateMap {
		hongBaoTemplate, _ := templateObject.(*gametemplate.HongBaoTemplate)
		cs.hongBaoMap[hongBaoTemplate.GetHongBaoType()] = hongBaoTemplate
	}

	return nil
}

//获取红包配置
func (cs *hongBaoTemplateService) GetHongBaoByTemplateType(typ itemtypes.ItemHongBaoSubType) *gametemplate.HongBaoTemplate {
	hongBao, ok := cs.hongBaoMap[typ]
	if !ok {
		return nil
	}

	return hongBao
}

var (
	once sync.Once
	cs   *hongBaoTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &hongBaoTemplateService{}
		err = cs.init()
	})
	return err
}

func GetHongBaoTemplateService() HongBaoTemplateService {
	return cs
}
