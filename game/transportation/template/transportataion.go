package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	transportationtypes "fgame/fgame/game/transportation/types"
	"sync"
)

type TransportationTemplateService interface {
	//获取镖车配置
	GetTransportationTemplate(typ transportationtypes.TransportationType) *gametemplate.BiaocheTemplate
	//获取镖车路径初始配置
	GetTransportationMoveTemplateFirst() *gametemplate.BiaocheMoveTemplate
	//获取镖车路径配置
	GetTransportationMoveTemplate(tempId int32) *gametemplate.BiaocheMoveTemplate
}

type transportationTemplateService struct {
	transportMap     map[int32]*gametemplate.BiaocheTemplate
	transportTypeMap map[transportationtypes.TransportationType]*gametemplate.BiaocheTemplate
	transportMoveMap map[int32]*gametemplate.BiaocheMoveTemplate
}

func (st *transportationTemplateService) init() (err error) {
	st.transportMap = make(map[int32]*gametemplate.BiaocheTemplate)
	st.transportTypeMap = make(map[transportationtypes.TransportationType]*gametemplate.BiaocheTemplate)

	tempMap := template.GetTemplateService().GetAll((*gametemplate.BiaocheTemplate)(nil))
	for _, tem := range tempMap {
		ftem, _ := tem.(*gametemplate.BiaocheTemplate)
		st.transportMap[int32(ftem.TemplateId())] = ftem
		st.transportTypeMap[ftem.GetTransportType()] = ftem
	}

	st.transportMoveMap = make(map[int32]*gametemplate.BiaocheMoveTemplate)
	moveTempMap := template.GetTemplateService().GetAll((*gametemplate.BiaocheMoveTemplate)(nil))
	for _, tem := range moveTempMap {
		ftem, _ := tem.(*gametemplate.BiaocheMoveTemplate)
		st.transportMoveMap[int32(ftem.TemplateId())] = ftem
	}

	return
}

func (st *transportationTemplateService) GetTransportationTemplate(typ transportationtypes.TransportationType) *gametemplate.BiaocheTemplate {
	return st.transportTypeMap[typ]
}

func (st *transportationTemplateService) GetTransportationMoveTemplateFirst() *gametemplate.BiaocheMoveTemplate {
	return st.transportMoveMap[int32(1)]
}

func (st *transportationTemplateService) GetTransportationMoveTemplate(tempId int32) *gametemplate.BiaocheMoveTemplate {
	return st.transportMoveMap[tempId]
}

var (
	once sync.Once
	st    *transportationTemplateService
)

func Init() (err error) {
	once.Do(func() {
		st = &transportationTemplateService{}
		err = st.init()
	})

	return
}

func GetTransportationTemplateService() TransportationTemplateService {
	return st
}
