package template

import (
	"fgame/fgame/core/template"
	shangguzhilingtypes "fgame/fgame/game/shangguzhiling/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sort"
	"sync"
)

type ShangguzhilingTemplateService interface {
	//上古之灵基础模板
	GetLingShouTemplate(typ shangguzhilingtypes.LingshouType) *gametemplate.ShangguzhilingBaseTemplate
	//灵纹配置
	GetLingWenTemplate(typ shangguzhilingtypes.LingshouType, lingwentyp shangguzhilingtypes.LingwenType) *gametemplate.ShangguzhilingLingwenTemplate
	//灵炼配置
	GetLingLianTemplate(typ shangguzhilingtypes.LingshouType, lingliantyp shangguzhilingtypes.LinglianPosType) *gametemplate.ShangguzhilingLinglianTemplate
	//常量配置
	GetConstant() *gametemplate.ShangguzhilingConstantTemplate
	//灵炼套装(降序排序)
	GetLingLianTaoZhuangListTempLate() []*gametemplate.ShangguzhilingLinglianTaozhuangTemplate
	//灵炼锁定配置
	GetLingLianLockUseItemCount(times int32) int32
}

type shangguzhilingTemplateService struct {
	lingshouBaseMap map[shangguzhilingtypes.LingshouType]*gametemplate.ShangguzhilingBaseTemplate
	lingwenMap      map[shangguzhilingtypes.LingshouType]map[shangguzhilingtypes.LingwenType]*gametemplate.ShangguzhilingLingwenTemplate
	linglianMap     map[shangguzhilingtypes.LingshouType]map[shangguzhilingtypes.LinglianPosType]*gametemplate.ShangguzhilingLinglianTemplate
	constantTemp    *gametemplate.ShangguzhilingConstantTemplate
	taozhuangList   lingLianTaozhuangList
	lockMap         map[int32]*gametemplate.ShangguzhilingLinglianLockTemplate
	maxLockTemp     *gametemplate.ShangguzhilingLinglianLockTemplate
}

//灵炼套装分组模板排序类型
type lingLianTaozhuangList []*gametemplate.ShangguzhilingLinglianTaozhuangTemplate

func (s lingLianTaozhuangList) Len() int           { return len(s) }
func (s lingLianTaozhuangList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s lingLianTaozhuangList) Less(i, j int) bool { return s[i].NeedLevel > s[j].NeedLevel }

//获取常量模板
func (s *shangguzhilingTemplateService) GetConstant() *gametemplate.ShangguzhilingConstantTemplate {
	return s.constantTemp
}

func (s *shangguzhilingTemplateService) GetLingLianTaoZhuangListTempLate() []*gametemplate.ShangguzhilingLinglianTaozhuangTemplate {
	return s.taozhuangList
}

func (s *shangguzhilingTemplateService) GetLingLianLockUseItemCount(times int32) int32 {
	if times <= int32(0) {
		return int32(0)
	}
	temp, ok := s.lockMap[times]
	if !ok {
		return s.maxLockTemp.SuodingUseItemCount
	}
	return temp.SuodingUseItemCount
}

//上古之灵基础模板
func (s *shangguzhilingTemplateService) GetLingShouTemplate(typ shangguzhilingtypes.LingshouType) *gametemplate.ShangguzhilingBaseTemplate {
	temp, ok := s.lingshouBaseMap[typ]
	if !ok {
		return nil
	}
	return temp
}

//灵纹配置
func (s *shangguzhilingTemplateService) GetLingWenTemplate(typ shangguzhilingtypes.LingshouType, lingwentyp shangguzhilingtypes.LingwenType) *gametemplate.ShangguzhilingLingwenTemplate {
	tempM, ok := s.lingwenMap[typ]
	if !ok {
		return nil
	}
	temp, ok := tempM[lingwentyp]
	if !ok {
		return nil
	}
	return temp
}

//灵炼配置
func (s *shangguzhilingTemplateService) GetLingLianTemplate(typ shangguzhilingtypes.LingshouType, lingliantyp shangguzhilingtypes.LinglianPosType) *gametemplate.ShangguzhilingLinglianTemplate {
	tempM, ok := s.linglianMap[typ]
	if !ok {
		return nil
	}
	temp, ok := tempM[lingliantyp]
	if !ok {
		return nil
	}
	return temp
}

func (s *shangguzhilingTemplateService) init() error {
	//上古之灵基础配置
	s.lingshouBaseMap = make(map[shangguzhilingtypes.LingshouType]*gametemplate.ShangguzhilingBaseTemplate)
	templateMap := template.GetTemplateService().GetAll((*gametemplate.ShangguzhilingBaseTemplate)(nil))
	for _, templateObject := range templateMap {
		lingshouTemplate, _ := templateObject.(*gametemplate.ShangguzhilingBaseTemplate)
		s.lingshouBaseMap[lingshouTemplate.GetLingShouType()] = lingshouTemplate
	}

	//灵纹配置
	s.lingwenMap = make(map[shangguzhilingtypes.LingshouType]map[shangguzhilingtypes.LingwenType]*gametemplate.ShangguzhilingLingwenTemplate)
	templateMap = template.GetTemplateService().GetAll((*gametemplate.ShangguzhilingLingwenTemplate)(nil))
	for _, templateObject := range templateMap {
		temp, _ := templateObject.(*gametemplate.ShangguzhilingLingwenTemplate)
		tempM, ok := s.lingwenMap[temp.GetLingShouType()]
		if !ok {
			tempM = make(map[shangguzhilingtypes.LingwenType]*gametemplate.ShangguzhilingLingwenTemplate)
		}
		tempM[temp.GetLingWenType()] = temp
		s.lingwenMap[temp.GetLingShouType()] = tempM
	}

	//灵炼配置
	s.linglianMap = make(map[shangguzhilingtypes.LingshouType]map[shangguzhilingtypes.LinglianPosType]*gametemplate.ShangguzhilingLinglianTemplate)
	templateMap = template.GetTemplateService().GetAll((*gametemplate.ShangguzhilingLinglianTemplate)(nil))
	for _, templateObject := range templateMap {
		temp, _ := templateObject.(*gametemplate.ShangguzhilingLinglianTemplate)
		tempM, ok := s.linglianMap[temp.GetLingShouType()]
		if !ok {
			tempM = make(map[shangguzhilingtypes.LinglianPosType]*gametemplate.ShangguzhilingLinglianTemplate)
		}
		tempM[temp.GetLingLianPosType()] = temp
		s.linglianMap[temp.GetLingShouType()] = tempM
	}

	//上古之灵常量
	templateMap = template.GetTemplateService().GetAll((*gametemplate.ShangguzhilingConstantTemplate)(nil))
	if len(templateMap) != 1 {
		err := fmt.Errorf("上古之灵常量配置应该只有一条！")
		return err
	}
	for _, templateObject := range templateMap {
		constantTemplate, _ := templateObject.(*gametemplate.ShangguzhilingConstantTemplate)
		s.constantTemp = constantTemplate
	}

	//灵炼套装
	s.taozhuangList = make(lingLianTaozhuangList, 0, 1)
	templateMap = template.GetTemplateService().GetAll((*gametemplate.ShangguzhilingLinglianTaozhuangTemplate)(nil))
	for _, templateObject := range templateMap {
		temp, _ := templateObject.(*gametemplate.ShangguzhilingLinglianTaozhuangTemplate)
		s.taozhuangList = append(s.taozhuangList, temp)
	}
	sort.Sort(s.taozhuangList)

	//灵炼锁定配置
	s.lockMap = make(map[int32]*gametemplate.ShangguzhilingLinglianLockTemplate)
	templateMap = template.GetTemplateService().GetAll((*gametemplate.ShangguzhilingLinglianLockTemplate)(nil))
	var maxTemp *gametemplate.ShangguzhilingLinglianLockTemplate
	for _, templateObject := range templateMap {
		temp, _ := templateObject.(*gametemplate.ShangguzhilingLinglianLockTemplate)
		s.lockMap[temp.Times] = temp
		if temp.IsMaxTemp() {
			maxTemp = temp
		}
	}
	if maxTemp == nil {
		err := fmt.Errorf("上古之灵灵炼锁定配置应该有一个最大值模板！")
		return err
	}
	s.maxLockTemp = maxTemp

	return nil
}

var (
	once sync.Once
	cs   *shangguzhilingTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &shangguzhilingTemplateService{}
		err = cs.init()
	})
	return err
}

func GetShangguzhilingTemplateService() ShangguzhilingTemplateService {
	return cs
}
