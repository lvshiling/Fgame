package template

import (
	"fgame/fgame/core/template"
	activetypes "fgame/fgame/game/activity/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type ActiveTemplateService interface {
	//活动类型获取配置
	GetActiveByType(typ activetypes.ActivityType) *gametemplate.ActivityTemplate
	GetActiveTemplate(id int32) *gametemplate.ActivityTemplate
	//获取所有活动配置
	GetActiveAll() map[activetypes.ActivityType]*gametemplate.ActivityTemplate
	GetPkRewardCdByType(typ activetypes.ActivityType) *gametemplate.PkRewardCdTemplate
}

//服务配置
type activeTemplateService struct {
	activeMap       map[int32]*gametemplate.ActivityTemplate
	activeTypeMap   map[activetypes.ActivityType]*gametemplate.ActivityTemplate
	pkActiveTypeMap map[activetypes.ActivityType]*gametemplate.PkRewardCdTemplate
}

func (st *activeTemplateService) init() (err error) {
	//活动
	st.activeMap = make(map[int32]*gametemplate.ActivityTemplate)
	st.activeTypeMap = make(map[activetypes.ActivityType]*gametemplate.ActivityTemplate)
	st.pkActiveTypeMap = make(map[activetypes.ActivityType]*gametemplate.PkRewardCdTemplate)

	tempMap := template.GetTemplateService().GetAll((*gametemplate.ActivityTemplate)(nil))
	for _, temp := range tempMap {
		activityTmep, _ := temp.(*gametemplate.ActivityTemplate)
		st.activeMap[int32(activityTmep.TemplateId())] = activityTmep
		st.activeTypeMap[activityTmep.GetActivityType()] = activityTmep
	}

	tempPkRewardMap := template.GetTemplateService().GetAll((*gametemplate.PkRewardCdTemplate)(nil))
	for _, temp := range tempPkRewardMap {
		pkRewardTmep, _ := temp.(*gametemplate.PkRewardCdTemplate)
		st.pkActiveTypeMap[pkRewardTmep.GetActivityType()] = pkRewardTmep

	}

	return
}

func (st *activeTemplateService) GetActiveByType(typ activetypes.ActivityType) *gametemplate.ActivityTemplate {
	return st.activeTypeMap[typ]
}

func (st *activeTemplateService) GetActiveTemplate(id int32) *gametemplate.ActivityTemplate {
	return st.activeMap[id]
}

func (st *activeTemplateService) GetPkRewardCdByType(typ activetypes.ActivityType) *gametemplate.PkRewardCdTemplate {
	temp, ok := st.pkActiveTypeMap[typ]
	if !ok {
		return nil
	}
	return temp
}

func (st *activeTemplateService) GetActiveAll() map[activetypes.ActivityType]*gametemplate.ActivityTemplate {
	return st.activeTypeMap
}

var (
	once      sync.Once
	acService *activeTemplateService
)

func Init() (err error) {
	once.Do(func() {
		acService = &activeTemplateService{}
		err = acService.init()
	})
	return
}

func GetActivityTemplateService() ActiveTemplateService {
	return acService
}
