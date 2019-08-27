package moonlove

import (
	"fgame/fgame/core/template"
	moonlovetypes "fgame/fgame/game/moonlove/types"
	gametemplate "fgame/fgame/game/template"
	"sync"
)

type MoonloveTemplateService interface {
	//获取月下情缘配置
	GetMoonloveMap() map[int32]*gametemplate.MoonloveTemplate
	//获取月下情缘配置
	GetMoonloveTemplate(level int32) *gametemplate.MoonloveTemplate
	//获取月下情缘豪气配置
	GetMoonloveGenerousRankMap(ranking int32) *gametemplate.MoonloveRankTemplate
	//获取月下情缘魅力配置
	GetMoonloveCharmRankMap(ranking int32) *gametemplate.MoonloveRankTemplate
}

//配置服务
type moonloveTemplateService struct {
	moonloveMap   map[int32]*gametemplate.MoonloveTemplate
	rankByTypeMap map[moonlovetypes.MoonloveRankType]map[int32]*gametemplate.MoonloveRankTemplate
}

func (s *moonloveTemplateService) init() (err error) {
	//月下情缘
	s.moonloveMap = make(map[int32]*gametemplate.MoonloveTemplate)

	mlTempMap := template.GetTemplateService().GetAll((*gametemplate.MoonloveTemplate)(nil))
	for _, temp := range mlTempMap {
		mlTemp, _ := temp.(*gametemplate.MoonloveTemplate)
		s.moonloveMap[int32(mlTemp.TemplateId())] = mlTemp
	}

	//月下情缘排行榜
	s.rankByTypeMap = make(map[moonlovetypes.MoonloveRankType]map[int32]*gametemplate.MoonloveRankTemplate)
	rankTempMap := template.GetTemplateService().GetAll((*gametemplate.MoonloveRankTemplate)(nil))
	for _, temp := range rankTempMap {
		rankTemp, _ := temp.(*gametemplate.MoonloveRankTemplate)
		rankMap, ok := s.rankByTypeMap[rankTemp.GetRankType()]
		if !ok {
			rankMap = make(map[int32]*gametemplate.MoonloveRankTemplate)
			s.rankByTypeMap[rankTemp.GetRankType()] = rankMap
		}

		rankMap[rankTemp.Rank] = rankTemp
	}

	return
}

func (s *moonloveTemplateService) GetMoonloveMap() map[int32]*gametemplate.MoonloveTemplate {
	return s.moonloveMap
}

func (s *moonloveTemplateService) GetMoonloveTemplate(level int32) *gametemplate.MoonloveTemplate {
	for _, temp := range s.moonloveMap {
		if level < temp.MinLev || level > temp.MaxLev {
			continue
		}

		return temp
	}

	return nil
}

func (s *moonloveTemplateService) GetMoonloveGenerousRankMap(ranking int32) *gametemplate.MoonloveRankTemplate {
	rankMap, ok := s.rankByTypeMap[moonlovetypes.MoonloveRankTypeGenerous]
	if !ok {
		return nil
	}

	temp, ok := rankMap[ranking]
	if !ok {
		return nil
	}

	return temp
}

func (s *moonloveTemplateService) GetMoonloveCharmRankMap(ranking int32) *gametemplate.MoonloveRankTemplate {
	rankMap, ok := s.rankByTypeMap[moonlovetypes.MoonloveRankTypeGenerousCharm]
	if !ok {
		return nil
	}

	temp, ok := rankMap[ranking]
	if !ok {
		return nil
	}

	return temp
}

var (
	once sync.Once
	s    *moonloveTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &moonloveTemplateService{}
		err = s.init()
	})
	return
}

func GetMoonloveTemplateService() MoonloveTemplateService {
	return s
}
