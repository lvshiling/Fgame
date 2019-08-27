package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	weektypes "fgame/fgame/game/week/types"
	"sync"
)

type WeekTemplateService interface {
	//获取周卡配置
	GetWeekTemplate(typ weektypes.WeekType) *gametemplate.WeekTemplate
}

type weekTemplateService struct {
	//周卡特权配置
	weekMap map[weektypes.WeekType]*gametemplate.WeekTemplate
}

func (st *weekTemplateService) init() (err error) {
	// 周卡特权配置
	st.weekMap = make(map[weektypes.WeekType]*gametemplate.WeekTemplate)
	tempMap := template.GetTemplateService().GetAll((*gametemplate.WeekTemplate)(nil))
	for _, temp := range tempMap {
		weekTemp, _ := temp.(*gametemplate.WeekTemplate)
		st.weekMap[weekTemp.GetWeekType()] = weekTemp
	}

	return
}

func (st *weekTemplateService) GetWeekTemplate(typ weektypes.WeekType) *gametemplate.WeekTemplate {
	temp, ok := st.weekMap[typ]
	if !ok {
		return nil
	}

	return temp
}

var (
	once sync.Once
	st   *weekTemplateService
)

func Init() (err error) {
	once.Do(func() {
		st = &weekTemplateService{}
		err = st.init()
	})

	return
}

func GetWeekTemplateService() WeekTemplateService {
	return st
}
