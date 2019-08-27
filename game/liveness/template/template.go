package template

import (
	"fgame/fgame/core/template"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sort"
	"sync"
)

//排序
type HuoYueLevelTemplateList []*gametemplate.HuoYueLevelTemplate

func (l HuoYueLevelTemplateList) Len() int {
	return len(l)
}

func (l HuoYueLevelTemplateList) Less(i, j int) bool {
	return l[i].LevelMax < l[j].LevelMax
}

func (l HuoYueLevelTemplateList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

//排序
type HuoYueBoxTemplateList []*gametemplate.HuoYueBoxTemplate

func (l HuoYueBoxTemplateList) Len() int {
	return len(l)
}

func (l HuoYueBoxTemplateList) Less(i, j int) bool {
	return l[i].NeedStar < l[j].NeedStar
}

func (l HuoYueBoxTemplateList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

//活跃度模板接口处理
type HuoYueTemplateService interface {
	//获取活跃度模板配置
	GetHuoYueTemplate(questId int32) *gametemplate.HuoYueTemplate
	//获取星数奖励
	GetHuoYueBoxTemplate(boxId int32) *gametemplate.HuoYueBoxTemplate
	//根据等级获取奖励
	GetHuoYueLevelTemplate(questId int32, level int32) (*gametemplate.HuoYueLevelTemplate, bool)
	//活跃活跃度宝箱列表
	GetHuoYueBoxList() []*gametemplate.HuoYueBoxTemplate
}

type huoYueTemplateService struct {
	huoYueMap      map[int32]*gametemplate.HuoYueTemplate
	huoYueBoxList  []*gametemplate.HuoYueBoxTemplate
	huoYueLevelMap map[int32][]*gametemplate.HuoYueLevelTemplate
}

//初始化
func (s *huoYueTemplateService) init() error {
	s.huoYueMap = make(map[int32]*gametemplate.HuoYueTemplate)
	s.huoYueBoxList = make([]*gametemplate.HuoYueBoxTemplate, 0, 3)
	s.huoYueLevelMap = make(map[int32][]*gametemplate.HuoYueLevelTemplate)
	for _, tempCardTemplate := range template.GetTemplateService().GetAll((*gametemplate.HuoYueTemplate)(nil)) {
		huoYueTemplate := tempCardTemplate.(*gametemplate.HuoYueTemplate)
		s.huoYueMap[int32(huoYueTemplate.TemplateId())] = huoYueTemplate
	}

	for _, tempBoxTemplate := range template.GetTemplateService().GetAll((*gametemplate.HuoYueBoxTemplate)(nil)) {
		huoYueBoxTemplate := tempBoxTemplate.(*gametemplate.HuoYueBoxTemplate)
		s.huoYueBoxList = append(s.huoYueBoxList, huoYueBoxTemplate)
	}
	sort.Sort(sort.Reverse(HuoYueBoxTemplateList(s.huoYueBoxList)))

	for _, tempLevelTemplate := range template.GetTemplateService().GetAll((*gametemplate.HuoYueLevelTemplate)(nil)) {
		huoYueLevelTemplate := tempLevelTemplate.(*gametemplate.HuoYueLevelTemplate)
		levelList := s.huoYueLevelMap[huoYueLevelTemplate.QuestId]
		levelList = append(levelList, huoYueLevelTemplate)
		s.huoYueLevelMap[huoYueLevelTemplate.QuestId] = levelList
	}

	//排序
	for questId, levelList := range s.huoYueLevelMap {
		if len(levelList) == 0 {
			return fmt.Errorf("liveness: 活跃度任务列表应该是存在的")
		}
		sort.Sort(sort.Reverse(HuoYueLevelTemplateList(levelList)))
		s.huoYueLevelMap[questId] = levelList
	}

	for questId, _ := range s.huoYueMap {
		_, exist := s.huoYueLevelMap[questId]
		if !exist {
			return fmt.Errorf("liveness: 活跃度任务等级奖励应该是存在的 questId:%d", questId)
		}
	}

	return nil
}

//获取活跃度模板配置
func (s *huoYueTemplateService) GetHuoYueTemplate(questId int32) *gametemplate.HuoYueTemplate {
	to, exist := s.huoYueMap[questId]
	if !exist {
		return nil
	}
	return to
}

//活跃活跃度宝箱列表
func (s *huoYueTemplateService) GetHuoYueBoxList() []*gametemplate.HuoYueBoxTemplate {
	return s.huoYueBoxList
}

//获取宝箱奖励
func (s *huoYueTemplateService) GetHuoYueBoxTemplate(boxId int32) *gametemplate.HuoYueBoxTemplate {
	for _, huoYueBox := range s.huoYueBoxList {
		if int32(huoYueBox.Id) == boxId {
			return huoYueBox
		}
	}
	return nil
}

//根据等级获取奖励
func (s *huoYueTemplateService) GetHuoYueLevelTemplate(questId int32, level int32) (levelTemplate *gametemplate.HuoYueLevelTemplate, flag bool) {
	huoYueLevelList, exist := s.huoYueLevelMap[questId]
	if !exist {
		return
	}
	if len(huoYueLevelList) == 0 {
		return
	}

	for _, huoYueLevel := range huoYueLevelList {
		if level >= huoYueLevel.LevelMin {
			flag = true
			levelTemplate = huoYueLevel
		}
	}

	//以防策划配错
	if levelTemplate == nil {
		flag = true
		levelTemplate = huoYueLevelList[0]
	}
	return
}

var (
	once sync.Once
	cs   *huoYueTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &huoYueTemplateService{}
		err = cs.init()
	})
	return err
}

func GetHuoYueTempalteService() HuoYueTemplateService {
	return cs
}
