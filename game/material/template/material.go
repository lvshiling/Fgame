package material

import (
	"fgame/fgame/core/template"
	materialtypes "fgame/fgame/game/material/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

//模板配置获取接口
type MaterialTemplateService interface {
	GetMaterialTemplate(typ materialtypes.MaterialType) *gametemplate.MaterialTemplate
	//材料副本免费挑战次数
	GetFreePlayTimes(typ materialtypes.MaterialType) int32
}

type materialTemplateService struct {
	materialMap map[materialtypes.MaterialType]*gametemplate.MaterialTemplate
}

//初始化材料副本配置
func (s *materialTemplateService) init() (err error) {
	//银两副本
	s.materialMap = make(map[materialtypes.MaterialType]*gametemplate.MaterialTemplate)

	materialTempMap := template.GetTemplateService().GetAll((*gametemplate.MaterialTemplate)(nil))
	for _, temp := range materialTempMap {
		materialTemp, _ := temp.(*gametemplate.MaterialTemplate)
		s.materialMap[materialTemp.GetMaterialType()] = materialTemp
	}

	return
}

// //获取初始秘境材料副本配置
// func (s *materialTemplateService) GetFirstMaterial(xfType materialtypes.MaterialType) *gametemplate.MaterialTemplate {
// 	return getMaterialMap(xfType)[int32(1)]
// }

//获取秘境材料副本配置
func (s *materialTemplateService) GetMaterialTemplate(typ materialtypes.MaterialType) *gametemplate.MaterialTemplate {
	return s.materialMap[typ]
}

// //获取基础次数
// func (s *materialTemplateService) GetBasicPlayTimes(xfType materialtypes.MaterialType) int32 {
// 	constantType := xfType.GetChallengeNumConstantType()
// 	maxChallengeTimes := constant.GetConstantService().GetConstant(constantType)

// 	return maxChallengeTimes
// }

//材料副本免费挑战次数
func (s *materialTemplateService) GetFreePlayTimes(typ materialtypes.MaterialType) int32 {
	temp := s.GetMaterialTemplate(typ)
	if temp == nil {
		return 0
	}

	return temp.Free
}

var (
	once sync.Once
	s    *materialTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &materialTemplateService{}
		err = s.init()
	})
	return
}

func GetMaterialTemplateService() MaterialTemplateService {
	return s
}
