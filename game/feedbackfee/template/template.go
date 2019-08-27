package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type FeedbackfeeTemplateService interface {
	GetExchangeTemplate(money int32) *gametemplate.MoneyTemplate
}

type feedbackfeeTemplateService struct {
	moneyMap map[int32]*gametemplate.MoneyTemplate
}

func (t *feedbackfeeTemplateService) init() (err error) {
	t.moneyMap = make(map[int32]*gametemplate.MoneyTemplate)
	allMoneyTemplate := template.GetTemplateService().GetAll((*gametemplate.MoneyTemplate)(nil))
	for _, tempMoneyTemplate := range allMoneyTemplate {
		moneyTemplate := tempMoneyTemplate.(*gametemplate.MoneyTemplate)
		t.moneyMap[moneyTemplate.Money] = moneyTemplate
	}
	return
}

func (t *feedbackfeeTemplateService) GetExchangeTemplate(money int32) *gametemplate.MoneyTemplate {
	moneyTemplate, ok := t.moneyMap[money]
	if !ok {
		return nil
	}
	return moneyTemplate
}

var (
	once sync.Once
	st   *feedbackfeeTemplateService
)

func Init() (err error) {
	once.Do(func() {
		st = &feedbackfeeTemplateService{}
		err = st.init()
	})

	return
}

func GetFeedbackfeeTemplateService() FeedbackfeeTemplateService {
	return st
}
