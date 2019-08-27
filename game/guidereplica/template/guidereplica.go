package template

import (
	"fgame/fgame/core/template"
	guidereplicatypes "fgame/fgame/game/guidereplica/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//引导副本配置服务
type GuideReplicaTemplateService interface {
	// 引导副本配置
	GetGuideReplicaTemplate(id int) *gametemplate.GuideReplicaTemplate
	// 引导副本配置
	GetGuideReplicaTemplateByArg(MapId int32, typ guidereplicatypes.GuideReplicaType) *gametemplate.GuideReplicaTemplate
}

type guideReplicaTemplateService struct {
	fubenMap      map[int]*gametemplate.GuideReplicaTemplate
	fubenByArgMap map[guidereplicatypes.GuideReplicaType]map[int32]*gametemplate.GuideReplicaTemplate
}

//初始化
func (ts *guideReplicaTemplateService) init() error {
	ts.fubenMap = make(map[int]*gametemplate.GuideReplicaTemplate)

	templateMap := template.GetTemplateService().GetAll((*gametemplate.GuideReplicaTemplate)(nil))
	for _, temp := range templateMap {
		guideReplicaTemplate, _ := temp.(*gametemplate.GuideReplicaTemplate)
		ts.fubenMap[guideReplicaTemplate.TemplateId()] = guideReplicaTemplate
	}

	ts.fubenByArgMap = make(map[guidereplicatypes.GuideReplicaType]map[int32]*gametemplate.GuideReplicaTemplate)
	for _, templateObject := range templateMap {
		to, _ := templateObject.(*gametemplate.GuideReplicaTemplate)
		tempM, ok := ts.fubenByArgMap[to.GetGuideType()]
		if !ok {
			tempM = make(map[int32]*gametemplate.GuideReplicaTemplate)
			ts.fubenByArgMap[to.GetGuideType()] = tempM
		}
		tempM[to.MapId] = to
	}
	return nil
}

func (ts *guideReplicaTemplateService) GetGuideReplicaTemplate(id int) *gametemplate.GuideReplicaTemplate {
	temp, ok := ts.fubenMap[id]
	if !ok {
		return nil
	}

	return temp
}

func (ts *guideReplicaTemplateService) GetGuideReplicaTemplateByArg(MapId int32, typ guidereplicatypes.GuideReplicaType) *gametemplate.GuideReplicaTemplate {
	subMap, ok := ts.fubenByArgMap[typ]
	if !ok {
		return nil
	}

	to, ok := subMap[MapId]
	if !ok {
		return nil
	}

	return to
}

var (
	once sync.Once
	cs   *guideReplicaTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &guideReplicaTemplateService{}
		err = cs.init()
	})
	return err
}

func GetGuideReplicaTemplateService() GuideReplicaTemplateService {
	return cs
}
