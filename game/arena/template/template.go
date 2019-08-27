package template

import (
	"fgame/fgame/core/template"
	arenatypes "fgame/fgame/game/arena/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sort"
	"sync"
)

//快捷缓存
//配置的整合
type ArenaTemplateService interface {
	//3v3常量配置
	GetArenaConstantTemplate() *gametemplate.ArenaConstantTemplate
	//3v3配置
	GetArenaTemplate(arenaType arenatypes.ArenaType, count int32) *gametemplate.ArenaTemplate
	// 机器人
	GetThreeRobotTemplate(level int32) *gametemplate.ThreeRobatTemplate
	//获取排行榜根据名次
	GetArenaRankTemplateByRankNum(rankNum int32) *gametemplate.ArenaRankTemplate
}

type threeRobotList []*gametemplate.ThreeRobatTemplate

func (adl threeRobotList) Len() int {
	return len(adl)
}

func (adl threeRobotList) Less(i, j int) bool {
	return adl[i].LevelMin < adl[j].LevelMin
}

func (adl threeRobotList) Swap(i, j int) {
	adl[i], adl[j] = adl[j], adl[i]
}

//排序
type ArenaRankTemplateList []*gametemplate.ArenaRankTemplate

func (s ArenaRankTemplateList) Len() int {
	return len(s)
}

func (s ArenaRankTemplateList) Less(i, j int) bool {
	return s[i].RankMin < s[j].RankMin
}

func (s ArenaRankTemplateList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//快捷缓存
//配置的整合
type arenaTemplateService struct {
	arenaContantTemplate *gametemplate.ArenaConstantTemplate
	//
	arenaMap    map[arenatypes.ArenaType]map[int32]*gametemplate.ArenaTemplate
	maxCountMap map[arenatypes.ArenaType]int32
	//
	threeRobotList []*gametemplate.ThreeRobatTemplate
	//排行榜
	arenaRankList []*gametemplate.ArenaRankTemplate
}

func (s *arenaTemplateService) init() (err error) {
	tempArenaConstantTemplate := template.GetTemplateService().Get(1, (*gametemplate.ArenaConstantTemplate)(nil))
	if tempArenaConstantTemplate == nil {
		return fmt.Errorf("竞技场常量表不存在")
	}
	s.arenaContantTemplate = tempArenaConstantTemplate.(*gametemplate.ArenaConstantTemplate)

	s.arenaMap = make(map[arenatypes.ArenaType]map[int32]*gametemplate.ArenaTemplate)
	s.maxCountMap = make(map[arenatypes.ArenaType]int32)
	tempArenaTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ArenaTemplate)(nil))
	for _, tempArenaTemplate := range tempArenaTemplateMap {
		arenaTemplate := tempArenaTemplate.(*gametemplate.ArenaTemplate)
		areaType := arenaTemplate.GetArenaType()
		if areaType != arenatypes.ArenaTypeArena &&
			areaType != arenatypes.ArenaTypeFail {
			continue
		}
		subMap, ok := s.arenaMap[areaType]
		if !ok {
			subMap = make(map[int32]*gametemplate.ArenaTemplate)
			s.arenaMap[areaType] = subMap
		}
		subMap[arenaTemplate.LianXuCount] = arenaTemplate

		maxCount := s.maxCountMap[arenaTemplate.GetArenaType()]
		if maxCount < arenaTemplate.LianXuCount {
			s.maxCountMap[arenaTemplate.GetArenaType()] = arenaTemplate.LianXuCount
		}
	}

	tempThreeRobotTemplateMap := template.GetTemplateService().GetAll((*gametemplate.ThreeRobatTemplate)(nil))
	s.threeRobotList = make([]*gametemplate.ThreeRobatTemplate, 0, 4)
	for _, tempThreeRobatTemplate := range tempThreeRobotTemplateMap {
		threeRobatTemplate := tempThreeRobatTemplate.(*gametemplate.ThreeRobatTemplate)
		s.threeRobotList = append(s.threeRobotList, threeRobatTemplate)
	}

	sort.Sort(threeRobotList(s.threeRobotList))

	//排行榜
	s.arenaRankList = make([]*gametemplate.ArenaRankTemplate, 0, 8)
	toMap := template.GetTemplateService().GetAll((*gametemplate.ArenaRankTemplate)(nil))
	for _, to := range toMap {
		rankTemplate, _ := to.(*gametemplate.ArenaRankTemplate)
		s.arenaRankList = append(s.arenaRankList, rankTemplate)
	}
	sort.Sort(sort.Reverse(ArenaRankTemplateList(s.arenaRankList)))
	return nil
}

func (s *arenaTemplateService) GetArenaConstantTemplate() *gametemplate.ArenaConstantTemplate {
	return s.arenaContantTemplate
}

func (s *arenaTemplateService) GetArenaTemplate(arenaType arenatypes.ArenaType, count int32) *gametemplate.ArenaTemplate {
	maxCount := s.maxCountMap[arenaType]
	if count > maxCount {
		count = maxCount
	}
	subMap, ok := s.arenaMap[arenaType]
	if !ok {
		return nil
	}
	arenaTemplate, ok := subMap[count]
	if !ok {
		return nil
	}
	return arenaTemplate
}

func (s *arenaTemplateService) GetThreeRobotTemplate(level int32) *gametemplate.ThreeRobatTemplate {
	for _, threeRobotTemplate := range s.threeRobotList {
		if threeRobotTemplate.LevelMin <= level && level <= threeRobotTemplate.LevelMax {
			return threeRobotTemplate
		}
	}
	return nil
}

//获取排行榜根据名次
func (s *arenaTemplateService) GetArenaRankTemplateByRankNum(rankNum int32) *gametemplate.ArenaRankTemplate {
	if rankNum <= 0 || rankNum > 100 {
		panic(fmt.Errorf("shenmo: rankNum 应该是大于0小于等于100的"))
	}
	for index, arenaRank := range s.arenaRankList {
		if rankNum >= arenaRank.RankMin && index == 0 {
			return arenaRank
		}
		if rankNum < arenaRank.RankMin {
			continue
		}
		if rankNum >= arenaRank.RankMin && rankNum <= arenaRank.RankMax {
			return arenaRank
		}
	}
	return nil
}

var (
	once sync.Once
	cs   *arenaTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &arenaTemplateService{}
		err = cs.init()
	})
	return err
}

func GetArenaTemplateService() ArenaTemplateService {
	return cs
}
