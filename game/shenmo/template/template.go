package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sort"
	"sync"
)

//排序
type ShenMoTitleTemplateList []*gametemplate.ShenMoTitleTemplate

func (s ShenMoTitleTemplateList) Len() int {
	return len(s)
}

func (s ShenMoTitleTemplateList) Less(i, j int) bool {
	return s[i].KillMin < s[j].KillMin
}

func (s ShenMoTitleTemplateList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//排序
type ShenMoRankTemplateList []*gametemplate.ShenMoRankTemplate

func (s ShenMoRankTemplateList) Len() int {
	return len(s)
}

func (s ShenMoRankTemplateList) Less(i, j int) bool {
	return s[i].RankMin < s[j].RankMin
}

func (s ShenMoRankTemplateList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//神魔接口处理
type ShenMoTemplateService interface {
	//获取神魔常量模板
	GetShenMoConstantTemplate() *gametemplate.ShenMoConstantTemplate
	//获取神魔称号根据杀人数
	GetShenMoTitleTemplateByKillNum(killNum int32) *gametemplate.ShenMoTitleTemplate
	//获取神魔排行榜根据名次
	GetShenMoRankTemplateByRankNum(rankNum int32) *gametemplate.ShenMoRankTemplate
}

type shenMoTemplateService struct {
	//神魔模板
	shenMoConstantTemplate *gametemplate.ShenMoConstantTemplate
	//神魔称号
	shenMoTitleList []*gametemplate.ShenMoTitleTemplate
	//神魔排行
	shenMoRankList []*gametemplate.ShenMoRankTemplate
}

//初始化
func (s *shenMoTemplateService) init() error {
	s.shenMoTitleList = make([]*gametemplate.ShenMoTitleTemplate, 0, 8)
	s.shenMoRankList = make([]*gametemplate.ShenMoRankTemplate, 0, 8)

	//神魔常量
	templateMap := template.GetTemplateService().GetAll((*gametemplate.ShenMoConstantTemplate)(nil))
	for _, templateObject := range templateMap {
		s.shenMoConstantTemplate, _ = templateObject.(*gametemplate.ShenMoConstantTemplate)
		break
	}

	//神魔称号
	templateTitleMap := template.GetTemplateService().GetAll((*gametemplate.ShenMoTitleTemplate)(nil))
	for _, templateObject := range templateTitleMap {
		titleTemplate, _ := templateObject.(*gametemplate.ShenMoTitleTemplate)
		s.shenMoTitleList = append(s.shenMoTitleList, titleTemplate)
	}
	sort.Sort(sort.Reverse(ShenMoTitleTemplateList(s.shenMoTitleList)))

	//神魔排行
	templateRankMap := template.GetTemplateService().GetAll((*gametemplate.ShenMoRankTemplate)(nil))
	for _, templateObject := range templateRankMap {
		rankTemplate, _ := templateObject.(*gametemplate.ShenMoRankTemplate)
		s.shenMoRankList = append(s.shenMoRankList, rankTemplate)
	}
	sort.Sort(sort.Reverse(ShenMoRankTemplateList(s.shenMoRankList)))
	return nil
}

//获取神魔常量模板
func (s *shenMoTemplateService) GetShenMoConstantTemplate() *gametemplate.ShenMoConstantTemplate {
	return s.shenMoConstantTemplate
}

//获取神魔称号根据杀人数
func (s *shenMoTemplateService) GetShenMoTitleTemplateByKillNum(killNum int32) *gametemplate.ShenMoTitleTemplate {
	if killNum < 0 {
		panic(fmt.Errorf("shenmo: killNum 应该是大于0的"))
	}
	for index, shenMoTitle := range s.shenMoTitleList {
		if killNum >= shenMoTitle.KillMin && index == 0 {
			return shenMoTitle
		}
		if killNum < shenMoTitle.KillMin {
			continue
		}
		if killNum >= shenMoTitle.KillMin && killNum <= shenMoTitle.KillMax {
			return shenMoTitle
		}
	}
	return nil
}

//获取神魔排行榜根据名次
func (s *shenMoTemplateService) GetShenMoRankTemplateByRankNum(rankNum int32) *gametemplate.ShenMoRankTemplate {
	if rankNum <= 0 || rankNum > 100 {
		panic(fmt.Errorf("shenmo: rankNum 应该是大于0小于等于100的"))
	}
	for index, shenMoRank := range s.shenMoRankList {
		if rankNum >= shenMoRank.RankMin && index == 0 {
			return shenMoRank
		}
		if rankNum < shenMoRank.RankMin {
			continue
		}
		if rankNum >= shenMoRank.RankMin && rankNum <= shenMoRank.RankMax {
			return shenMoRank
		}
	}
	return nil
}

var (
	once sync.Once
	cs   *shenMoTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &shenMoTemplateService{}
		err = cs.init()
	})
	return err
}

func GetShenMoTemplateService() ShenMoTemplateService {
	return cs
}
