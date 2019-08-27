package template

// import (
// 	logintypes	"fgame/fgame/account/login/types"
// 	"fgame/fgame/core/template"
// 	ranktypes "fgame/fgame/game/rank/types"
// 	gametemplate "fgame/fgame/game/template"
// 	"math/rand"
// 	"sync"
// )

// type RankTemplateService interface {
// 	//获取渠道模板
// 	GetQuDaoTemplate(sdkType logintypes.SDKType) *gametemplate.RankTemplate
// }

// type rankTemplateService struct {
// 	qudaoTempMap map[logintypes.SDKType]*gametemplate.RankTemplate
// }

// //初始化
// func (s *rankTemplateService) init() error {
// 	s.qudaoTempMap = make(map[logintypes.SDKType]*gametemplate.RankTemplate)

// 	//渠道
// 	templateMap := template.GetTemplateService().GetAll((*gametemplate.RankTemplate)(nil))
// 	for _, temp := range templateMap {
// 		rankTemplate, _ := temp.(*gametemplate.RankTemplate)
// 		s.rankMap[rankTemplate.TemplateId()] = rankTemplate

// 		typ := rankTemplate.GetTyp()
// 		if typ == ranktypes.RankTypeAdvanced {
// 			s.rankNumberMap[rankTemplate.Number] = rankTemplate
// 		}
// 		s.shenFaList = append(s.shenFaList, rankTemplate)
// 	}

// 	return nil
// }

// //
// func (s *rankTemplateService) GetQuDaoTemplate(sdkType logintypes.SDKType) *gametemplate.RankTemplate {
// 	temp,ok := s.qudaoTempMap[sdkType]
// 	if !ok{
// 		return nil
// 	}

// 	return  temp
// }

// //吃幻化

// var (
// 	once sync.Once
// 	cs   *rankTemplateService
// )

// func Init() (err error) {
// 	once.Do(func() {
// 		cs = &rankTemplateService{}
// 		err = cs.init()
// 	})
// 	return err
// }

// func GetRankTemplateService() RankTemplateService {
// 	return cs
// }
