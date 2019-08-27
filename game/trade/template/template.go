package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sync"
)

//宝箱模板接口处理
type TradeTemplateService interface {
	//获取宝箱模板
	GetTradeConstantTemplate() *gametemplate.TradeConstantTemplate
}

type tradeTemplateService struct {
	tradeConstantTemplate *gametemplate.TradeConstantTemplate
}

//初始化
func (s *tradeTemplateService) init() (err error) {
	tradeConstantTemplate := template.GetTemplateService().Get(1, (*gametemplate.TradeConstantTemplate)(nil))
	if tradeConstantTemplate == nil {
		return fmt.Errorf("trade:常量表不能为空")
	}
	s.tradeConstantTemplate = tradeConstantTemplate.(*gametemplate.TradeConstantTemplate)
	return nil
}

//获取宝箱模板id
func (s *tradeTemplateService) GetTradeConstantTemplate() *gametemplate.TradeConstantTemplate {
	return s.tradeConstantTemplate
}

var (
	once sync.Once
	cs   *tradeTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &tradeTemplateService{}
		err = cs.init()
	})
	return err
}

func GetTradeTemplateService() TradeTemplateService {
	return cs
}
