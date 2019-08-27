package template

import (
	"fgame/fgame/core/template"
	systemcompensatetypes "fgame/fgame/game/systemcompensate/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//系统补偿接口处理
type ReturnXiTongTemplateService interface {
	//获取系统补偿模板
	GetSystemCompensateTemplate(typ systemcompensatetypes.SystemCompensateType, number int32) *gametemplate.ReturnXiTongTemplate
}

type systemSkillTemplateService struct {
	//系统补偿配置
	sysCompensateTempMap map[systemcompensatetypes.SystemCompensateType]map[int32]*gametemplate.ReturnXiTongTemplate
}

//初始化
func (s *systemSkillTemplateService) init() error {
	s.sysCompensateTempMap = make(map[systemcompensatetypes.SystemCompensateType]map[int32]*gametemplate.ReturnXiTongTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.ReturnXiTongTemplate)(nil))
	for _, templateObject := range templateMap {
		skTemplate, _ := templateObject.(*gametemplate.ReturnXiTongTemplate)

		typ := skTemplate.GetSysCompensateType()
		number := skTemplate.Number

		skTypeTemplateMap, ok := s.sysCompensateTempMap[typ]
		if !ok {
			skTypeTemplateMap = make(map[int32]*gametemplate.ReturnXiTongTemplate)
			s.sysCompensateTempMap[typ] = skTypeTemplateMap
		}
		skTypeTemplateMap[number] = skTemplate
	}

	return nil
}

//获取系统补偿配置
func (s *systemSkillTemplateService) GetSystemCompensateTemplate(typ systemcompensatetypes.SystemCompensateType, number int32) *gametemplate.ReturnXiTongTemplate {
	skTypeTemplateMap, ok := s.sysCompensateTempMap[typ]
	if !ok {
		return nil
	}
	to, ok := skTypeTemplateMap[number]
	if !ok {
		return nil
	}
	return to
}

var (
	once sync.Once
	cs   *systemSkillTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &systemSkillTemplateService{}
		err = cs.init()
	})
	return err
}

func GetReturnXiTongTemplateService() ReturnXiTongTemplateService {
	return cs
}
