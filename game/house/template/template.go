package house

import (
	"fgame/fgame/core/template"
	housetypes "fgame/fgame/game/house/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"sync"
)

//房子接口处理
type HouseTemplateService interface {
	//获取升级房子配置
	GetHouseTemplate(houseIndex int32, houseType housetypes.HouseType, number int32) *gametemplate.HouseTemplate
	//获取房子常量配置
	GetHouseConstantTemplate() *gametemplate.HouseConstantTemplate
	//获取随机房子配置
	GetRandomHouseTemp() *gametemplate.HouseTemplate
}

type houseTemplateService struct {
	//房子配置
	houseNumberMap map[int32]map[housetypes.HouseType]map[int32]*gametemplate.HouseTemplate
	houseList      []*gametemplate.HouseTemplate
	//房子常量配置
	houseConstanTemp *gametemplate.HouseConstantTemplate
}

//初始化
func (s *houseTemplateService) init() error {
	s.houseNumberMap = make(map[int32]map[housetypes.HouseType]map[int32]*gametemplate.HouseTemplate)
	//房子
	templateMap := template.GetTemplateService().GetAll((*gametemplate.HouseTemplate)(nil))
	for _, templateObject := range templateMap {
		houseTemplate, _ := templateObject.(*gametemplate.HouseTemplate)
		subOfSubMap, ok := s.houseNumberMap[houseTemplate.HouseIndex]
		if !ok {
			subOfSubMap = make(map[housetypes.HouseType]map[int32]*gametemplate.HouseTemplate)
			s.houseNumberMap[houseTemplate.HouseIndex] = subOfSubMap
		}
		subMap, ok := subOfSubMap[houseTemplate.GetHouseType()]
		if !ok {
			subMap = make(map[int32]*gametemplate.HouseTemplate)
			subOfSubMap[houseTemplate.GetHouseType()] = subMap
		}
		subMap[houseTemplate.Level] = houseTemplate

		s.houseList = append(s.houseList, houseTemplate)
	}

	// 房子常量
	constantTemplateMap := template.GetTemplateService().GetAll((*gametemplate.HouseConstantTemplate)(nil))
	if len(constantTemplateMap) != 1 {
		return fmt.Errorf("house:房子常量配置只有一条")
	}
	for _, temp := range constantTemplateMap {
		constantTemp, _ := temp.(*gametemplate.HouseConstantTemplate)
		s.houseConstanTemp = constantTemp
	}

	return nil
}

//获取升级房子配置
func (s *houseTemplateService) GetHouseTemplate(houseIndex int32, houseType housetypes.HouseType, number int32) *gametemplate.HouseTemplate {
	subOfSubMap, ok := s.houseNumberMap[houseIndex]
	if !ok {
		return nil
	}
	subMap, ok := subOfSubMap[houseType]
	if !ok {
		return nil
	}

	to, ok := subMap[number]
	if !ok {
		return nil
	}
	return to
}

//获取随机房子配置
func (s *houseTemplateService) GetRandomHouseTemp() *gametemplate.HouseTemplate {
	if len(s.houseList) == 0 {
		return nil
	}

	randomIndex := mathutils.RandomRange(0, len(s.houseList))
	return s.houseList[randomIndex]
}

//获取房子常量配置
func (s *houseTemplateService) GetHouseConstantTemplate() *gametemplate.HouseConstantTemplate {
	return s.houseConstanTemp
}

var (
	once sync.Once
	cs   *houseTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &houseTemplateService{}
		err = cs.init()
	})
	return err
}

func GetHouseTemplateService() HouseTemplateService {
	return cs
}
