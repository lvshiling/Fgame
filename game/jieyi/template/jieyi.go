package template

import (
	jieyitypes "fgame/fgame/game/jieyi/types"
	gametemplate "fgame/fgame/game/template"
	"sync"

	"fgame/fgame/core/template"
)

type JieYiTemplateService interface {
	// 获取结义常量模板
	GetJieYiConstantTemplate() *gametemplate.JieYiConstantTemplate
	// 获取结义道具模板
	GetJieYiDaoJuTemplate(typ jieyitypes.JieYiDaoJuType) *gametemplate.JieYiDaoJuTemplate
	// 获取结义威名等级模板
	GetJieYiNameTemplate(level int32) *gametemplate.JieYiNameLevelTemplate
	// 获取结义信物模板
	GetJieYiTokenTemplate(typ jieyitypes.JieYiTokenType) *gametemplate.JieYiTokenTemplate
	// 获取结义信物等级模板
	GetJieYiTokenLevelTemplate(typ jieyitypes.JieYiTokenType, level int32) *gametemplate.JieYiTokenLevelTemplate
}

type jieYiTemplateService struct {
	// 结义常量
	jieYiConstantTemp *gametemplate.JieYiConstantTemplate
	// 结义道具
	jieYiDaoJuTempMap map[jieyitypes.JieYiDaoJuType]*gametemplate.JieYiDaoJuTemplate
	// 结义威名
	jieYiNameTempMap map[int32]*gametemplate.JieYiNameLevelTemplate
	// 结义信物
	jieYiTokenMap map[jieyitypes.JieYiTokenType]*gametemplate.JieYiTokenTemplate
}

func (s *jieYiTemplateService) init() (err error) {
	s.jieYiDaoJuTempMap = make(map[jieyitypes.JieYiDaoJuType]*gametemplate.JieYiDaoJuTemplate)
	s.jieYiNameTempMap = make(map[int32]*gametemplate.JieYiNameLevelTemplate)
	s.jieYiTokenMap = make(map[jieyitypes.JieYiTokenType]*gametemplate.JieYiTokenTemplate)

	// 常量
	constantTemp := template.GetTemplateService().GetAll((*gametemplate.JieYiConstantTemplate)(nil))
	if len(constantTemp) != 1 {
		panic("jieyi: 结义常量模板应该只有一个")
	}
	for _, temp := range constantTemp {
		constant, _ := temp.(*gametemplate.JieYiConstantTemplate)
		s.jieYiConstantTemp = constant
	}

	// 结义道具
	daoJuTempMap := template.GetTemplateService().GetAll((*gametemplate.JieYiDaoJuTemplate)(nil))
	for _, temp := range daoJuTempMap {
		daoJuTemp, _ := temp.(*gametemplate.JieYiDaoJuTemplate)
		s.jieYiDaoJuTempMap[daoJuTemp.GetDaoJuType()] = daoJuTemp
	}

	// 结义威名
	nameTempMap := template.GetTemplateService().GetAll((*gametemplate.JieYiNameLevelTemplate)(nil))
	for _, temp := range nameTempMap {
		nameTemp, _ := temp.(*gametemplate.JieYiNameLevelTemplate)
		s.jieYiNameTempMap[nameTemp.Level] = nameTemp
	}

	// 结义信物
	tokenTempMap := template.GetTemplateService().GetAll((*gametemplate.JieYiTokenTemplate)(nil))
	for _, temp := range tokenTempMap {
		tokenTemp, _ := temp.(*gametemplate.JieYiTokenTemplate)
		s.jieYiTokenMap[tokenTemp.GetTokenType()] = tokenTemp
	}

	return
}

func (s *jieYiTemplateService) GetJieYiConstantTemplate() *gametemplate.JieYiConstantTemplate {
	return s.jieYiConstantTemp
}

func (s *jieYiTemplateService) GetJieYiDaoJuTemplate(typ jieyitypes.JieYiDaoJuType) *gametemplate.JieYiDaoJuTemplate {
	temp, ok := s.jieYiDaoJuTempMap[typ]
	if !ok {
		return nil
	}
	return temp
}

func (s *jieYiTemplateService) GetJieYiNameTemplate(level int32) *gametemplate.JieYiNameLevelTemplate {
	temp, ok := s.jieYiNameTempMap[level]
	if !ok {
		return nil
	}
	return temp
}

func (s *jieYiTemplateService) GetJieYiTokenTemplate(typ jieyitypes.JieYiTokenType) *gametemplate.JieYiTokenTemplate {
	temp, ok := s.jieYiTokenMap[typ]
	if !ok {
		return nil
	}
	return temp
}

func (s *jieYiTemplateService) GetJieYiTokenLevelTemplate(typ jieyitypes.JieYiTokenType, level int32) *gametemplate.JieYiTokenLevelTemplate {
	temp, ok := s.jieYiTokenMap[typ]
	if !ok {
		return nil
	}
	levelTemp := temp.GetJieYiTokenLevelTemp(level)
	return levelTemp
}

var (
	once  sync.Once
	jieyi *jieYiTemplateService
)

func Init() (err error) {
	once.Do(func() {
		jieyi = &jieYiTemplateService{}
		err = jieyi.init()
	})
	return
}

func GetJieYiTemplateService() JieYiTemplateService {
	return jieyi
}
