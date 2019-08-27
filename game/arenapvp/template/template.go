package template

import (
	"fgame/fgame/core/template"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sort"
	"sync"
)

type ArenapvpTemplateService interface {
	//pvp常量配置
	GetArenapvpConstantTemplate() *gametemplate.ArenapvpConstantTemplate
	//pvp配置
	GetArenapvpTemplate(pvpType arenapvptypes.ArenapvpType) *gametemplate.ArenapvpTemplate
	//获取排行榜根据名次
	GetArenapvpRankTemplateByRankNum(rankNum int32) *gametemplate.ArenapvpRankTemplate
	//淘汰配置
	GetArenapvpTaoTaiTemp(pvpType arenapvptypes.ArenapvpType) *gametemplate.ArenapvpTaoTaiTemplate
}

//排序
type ArenapvpRankTemplateList []*gametemplate.ArenapvpRankTemplate

func (s ArenapvpRankTemplateList) Len() int {
	return len(s)
}

func (s ArenapvpRankTemplateList) Less(i, j int) bool {
	return s[i].RankMin < s[j].RankMin
}

func (s ArenapvpRankTemplateList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//快捷缓存
//配置的整合
type arenapvpTemplateService struct {
	arenapvpContantTemplate *gametemplate.ArenapvpConstantTemplate
	arenapvpMap             map[arenapvptypes.ArenapvpType]*gametemplate.ArenapvpTemplate
	arenapvpRankList        []*gametemplate.ArenapvpRankTemplate //排行榜
	arenapvpTaoTaiMap       map[arenapvptypes.ArenapvpType]*gametemplate.ArenapvpTaoTaiTemplate
	//
}

func (s *arenapvpTemplateService) init() (err error) {
	constantTo := template.GetTemplateService().Get(1, (*gametemplate.ArenapvpConstantTemplate)(nil))
	if constantTo == nil {
		return fmt.Errorf("竞技场常量表不存在")
	}
	s.arenapvpContantTemplate = constantTo.(*gametemplate.ArenapvpConstantTemplate)

	s.arenapvpMap = make(map[arenapvptypes.ArenapvpType]*gametemplate.ArenapvpTemplate)
	arenapvpToMap := template.GetTemplateService().GetAll((*gametemplate.ArenapvpTemplate)(nil))
	for _, to := range arenapvpToMap {
		arenapvpTemplate := to.(*gametemplate.ArenapvpTemplate)
		s.arenapvpMap[arenapvpTemplate.GetArenapvpType()] = arenapvpTemplate
	}

	//排行榜
	s.arenapvpRankList = make([]*gametemplate.ArenapvpRankTemplate, 0, 8)
	rankToMap := template.GetTemplateService().GetAll((*gametemplate.ArenapvpRankTemplate)(nil))
	for _, to := range rankToMap {
		rankTemplate, _ := to.(*gametemplate.ArenapvpRankTemplate)
		s.arenapvpRankList = append(s.arenapvpRankList, rankTemplate)
	}
	sort.Sort(sort.Reverse(ArenapvpRankTemplateList(s.arenapvpRankList)))

	// 淘汰配置
	s.arenapvpTaoTaiMap = make(map[arenapvptypes.ArenapvpType]*gametemplate.ArenapvpTaoTaiTemplate)
	taotaiToMap := template.GetTemplateService().GetAll((*gametemplate.ArenapvpTaoTaiTemplate)(nil))
	for _, to := range taotaiToMap {
		taotaiTemplate := to.(*gametemplate.ArenapvpTaoTaiTemplate)
		s.arenapvpTaoTaiMap[taotaiTemplate.GetArenapvpType()] = taotaiTemplate
	}

	return nil
}

func (s *arenapvpTemplateService) GetArenapvpConstantTemplate() *gametemplate.ArenapvpConstantTemplate {
	return s.arenapvpContantTemplate
}

func (s *arenapvpTemplateService) GetArenapvpTemplate(pvpType arenapvptypes.ArenapvpType) *gametemplate.ArenapvpTemplate {
	temp, ok := s.arenapvpMap[pvpType]
	if !ok {
		return nil
	}
	return temp
}

//获取排行榜根据名次
func (s *arenapvpTemplateService) GetArenapvpRankTemplateByRankNum(rankNum int32) *gametemplate.ArenapvpRankTemplate {
	if rankNum <= 0 || rankNum > 100 {
		panic(fmt.Errorf("shenmo: rankNum 应该是大于0小于等于100的"))
	}
	for index, arenapvpRank := range s.arenapvpRankList {
		if rankNum >= arenapvpRank.RankMin && index == 0 {
			return arenapvpRank
		}
		if rankNum < arenapvpRank.RankMin {
			continue
		}
		if rankNum >= arenapvpRank.RankMin && rankNum <= arenapvpRank.RankMax {
			return arenapvpRank
		}
	}
	return nil
}

//获取排行榜根据名次
func (s *arenapvpTemplateService) GetArenapvpTaoTaiTemp(pvpType arenapvptypes.ArenapvpType) *gametemplate.ArenapvpTaoTaiTemplate {
	temp, ok := s.arenapvpTaoTaiMap[pvpType]
	if !ok {
		return nil
	}
	return temp
}

var (
	once sync.Once
	cs   *arenapvpTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &arenapvpTemplateService{}
		err = cs.init()
	})
	return err
}

func GetArenapvpTemplateService() ArenapvpTemplateService {
	return cs
}
