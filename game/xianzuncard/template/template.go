package template

import (
	gametemplate "fgame/fgame/game/template"
	xianzuncardtypes "fgame/fgame/game/xianzuncard/types"
	"sync"

	"fgame/fgame/core/template"
)

type XianZunCardTemplateService interface {
	// 获取仙尊特权卡模板
	GetXianZunCardTemplate(typ xianzuncardtypes.XianZunCardType) *gametemplate.XianZunCardTemplate
}

type xianZunCardTemplateService struct {
	// 仙尊特权卡
	xianZunCardTempMap map[xianzuncardtypes.XianZunCardType]*gametemplate.XianZunCardTemplate
}

func (s *xianZunCardTemplateService) init() (err error) {
	s.xianZunCardTempMap = make(map[xianzuncardtypes.XianZunCardType]*gametemplate.XianZunCardTemplate)
	xianZunTempMap := template.GetTemplateService().GetAll((*gametemplate.XianZunCardTemplate)(nil))
	for _, temp := range xianZunTempMap {
		xianZunTemp, _ := temp.(*gametemplate.XianZunCardTemplate)
		s.xianZunCardTempMap[xianZunTemp.GetXianZunCardType()] = xianZunTemp
	}
	return
}

func (s *xianZunCardTemplateService) GetXianZunCardTemplate(typ xianzuncardtypes.XianZunCardType) *gametemplate.XianZunCardTemplate {
	temp, ok := s.xianZunCardTempMap[typ]
	if !ok {
		return nil
	}
	return temp
}

var (
	once    sync.Once
	xianZun *xianZunCardTemplateService
)

func Init() (err error) {
	once.Do(func() {
		xianZun = &xianZunCardTemplateService{}
		err = xianZun.init()
	})
	return
}

func GetXianZunCardTemplateService() XianZunCardTemplateService {
	return xianZun
}
