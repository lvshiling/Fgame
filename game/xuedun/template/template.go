package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sync"
)

//血盾模板接口处理
type XueDunTemplateService interface {
	//获取血盾模板升阶配置
	GetXueDunNumber(number int32, star int32) *gametemplate.BloodShieldTemplate
	//获取血盾培养配置
	GetXueDunPeiYangTemplate(level int32) *gametemplate.BloodPeiYangTemplate
	//获取血盾激活模板
	GetXueDunActiveTemplate() *gametemplate.BloodShieldTemplate
	//血盾吞噬
	GetXueDunEatPeiYangTemplate(curCulLevel int32, num int32) (bloodPeiYangTemplate *gametemplate.BloodPeiYangTemplate, flag bool)
}

type xueDunTemplateService struct {
	//升阶map
	numberMap map[int32]map[int32]*gametemplate.BloodShieldTemplate
	//血盾培养模板配置
	peiYangMap map[int32]*gametemplate.BloodPeiYangTemplate
}

//初始化
func (ms *xueDunTemplateService) init() error {
	ms.numberMap = make(map[int32]map[int32]*gametemplate.BloodShieldTemplate)
	ms.peiYangMap = make(map[int32]*gametemplate.BloodPeiYangTemplate)
	//血盾模板
	templateMap := template.GetTemplateService().GetAll((*gametemplate.BloodShieldTemplate)(nil))
	for _, templateObject := range templateMap {
		bloodTemplate, _ := templateObject.(*gametemplate.BloodShieldTemplate)

		bloodTypeMap, exist := ms.numberMap[bloodTemplate.Type]
		if !exist {
			bloodTypeMap = make(map[int32]*gametemplate.BloodShieldTemplate)
			ms.numberMap[bloodTemplate.Type] = bloodTypeMap
		}
		bloodTypeMap[bloodTemplate.Star] = bloodTemplate
	}

	//血盾培养模板
	peiYangTemplateMap := template.GetTemplateService().GetAll((*gametemplate.BloodPeiYangTemplate)(nil))
	for _, templateObject := range peiYangTemplateMap {
		peiYangTemplate, _ := templateObject.(*gametemplate.BloodPeiYangTemplate)
		ms.peiYangMap[peiYangTemplate.Level] = peiYangTemplate
	}

	activeTemplate := ms.GetXueDunNumber(1, 1)
	if activeTemplate == nil {
		return fmt.Errorf("xuedun:血盾激活模板应该是存在的")
	}

	return nil
}

//获取血盾模板进阶配置
func (ms *xueDunTemplateService) GetXueDunNumber(number int32, star int32) *gametemplate.BloodShieldTemplate {
	bloodStarMap, ok := ms.numberMap[number]
	if !ok {
		return nil
	}
	to, ok := bloodStarMap[star]
	if !ok {
		return nil
	}
	return to
}

//获取血盾模板培养配置
func (ms *xueDunTemplateService) GetXueDunPeiYangTemplate(level int32) *gametemplate.BloodPeiYangTemplate {
	to, ok := ms.peiYangMap[level]
	if !ok {
		return nil
	}
	return to
}

//获取血盾激活模板
func (ms *xueDunTemplateService) GetXueDunActiveTemplate() *gametemplate.BloodShieldTemplate {
	return ms.GetXueDunNumber(1, 1)
}

//血盾吞噬
func (ms *xueDunTemplateService) GetXueDunEatPeiYangTemplate(curCulLevel int32, num int32) (bloodPeiYangTemplate *gametemplate.BloodPeiYangTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curCulLevel + 1; leftNum > 0; level++ {
		bloodPeiYangTemplate, flag = ms.peiYangMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= bloodPeiYangTemplate.ItemCount
	}
	//次数要满足刚好升级升级
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

var (
	once sync.Once
	cs   *xueDunTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &xueDunTemplateService{}
		err = cs.init()
	})
	return err
}

func GetXueDunTemplateService() XueDunTemplateService {
	return cs
}
