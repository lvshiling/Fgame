package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//vip配置服务
type VipTemplaterService interface {
	// VIP模板
	GetVipTemplate(level, star int32) *gametemplate.VipTemplate
	GetVipTemplateById(id int32) *gametemplate.VipTemplate
	// 消费等级模板
	GetCostLevelTemplate(level int32) *gametemplate.CostLevelTemplate
}

type vipTemplaterService struct {
	vipMap       map[int32]map[int32]*gametemplate.VipTemplate
	vipByIdMap   map[int32]*gametemplate.VipTemplate
	costLevelMap map[int32]*gametemplate.CostLevelTemplate
}

//初始化
func (ts *vipTemplaterService) init() error {
	ts.vipMap = make(map[int32]map[int32]*gametemplate.VipTemplate)
	ts.vipByIdMap = make(map[int32]*gametemplate.VipTemplate)
	ts.costLevelMap = make(map[int32]*gametemplate.CostLevelTemplate)
	//vip
	templateMap := template.GetTemplateService().GetAll((*gametemplate.VipTemplate)(nil))
	for _, temp := range templateMap {
		vipTemplate, _ := temp.(*gametemplate.VipTemplate)
		starMap, ok := ts.vipMap[vipTemplate.Level]
		if !ok {
			starMap = make(map[int32]*gametemplate.VipTemplate)
			ts.vipMap[vipTemplate.Level] = starMap
		}
		starMap[vipTemplate.Star] = vipTemplate

		ts.vipByIdMap[int32(vipTemplate.Id)] = vipTemplate
	}

	// 付费等级
	costTemplateMap := template.GetTemplateService().GetAll((*gametemplate.CostLevelTemplate)(nil))
	for _, temp := range costTemplateMap {
		costLevelTemplate, _ := temp.(*gametemplate.CostLevelTemplate)
		_, ok := ts.costLevelMap[costLevelTemplate.Level]
		if !ok {
			ts.costLevelMap[costLevelTemplate.Level] = costLevelTemplate
		}
	}

	return nil
}

func (ts *vipTemplaterService) GetVipTemplate(level, star int32) *gametemplate.VipTemplate {
	starMap, ok := ts.vipMap[level]
	if !ok {
		return nil
	}
	temp, ok := starMap[star]
	if !ok {
		return nil
	}

	return temp
}

func (ts *vipTemplaterService) GetVipTemplateById(id int32) *gametemplate.VipTemplate {
	temp, ok := ts.vipByIdMap[id]
	if !ok {
		return nil
	}

	return temp
}

func (ts *vipTemplaterService) GetCostLevelTemplate(level int32) *gametemplate.CostLevelTemplate {
	temp, ok := ts.costLevelMap[level]
	if !ok {
		return nil
	}

	return temp
}

var (
	once sync.Once
	cs   *vipTemplaterService
)

func Init() (err error) {
	once.Do(func() {
		cs = &vipTemplaterService{}
		err = cs.init()
	})
	return err
}

func GetVipTemplateService() VipTemplaterService {
	return cs
}
