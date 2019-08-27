package secretcard

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/secretcard/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"sync"
)

//天机牌接口处理
type SecretCardService interface {
	//获取任务池
	GetQuestPool() map[types.SecretCardPoolType][]*gametemplate.TianJiPaiTemplate
	//获取任务对应任务池
	GetQuestPoolType(cardId int32) (typ types.SecretCardPoolType, flag bool)
	//获取任务id
	GetQuestIdByCardId(cardId int32) (questId int32, questName string, flag bool)
	//获取天机牌配置
	GetSecretCardTemplate(cardId int32) *gametemplate.TianJiPaiTemplate
	//获取掉落id
	GetDropIdByNum(cardId int32, num int32) (dropId int32)
	//获取星数奖励
	GetStarTemplate(boxId int32) *gametemplate.TianJiPaiStarTemplate
	//星数随机
	GetRandomStar(idList []int32) map[int32]int32
	//获取剩余未开启的宝箱模板
	GetLeftBoxTemplate(openedBox []int32) (closeStarTemplates []*gametemplate.TianJiPaiStarTemplate)
	//获取天机牌每日翻牌次数
	GetConstSecretCardNum() int32
	//天机牌一键完成消耗的元宝数量
	GetConstSecretCardCostGold() int32
	//获取天机牌玩家最小等级的
	GetSecretCardTemplateByQuestId(questId int32) (*gametemplate.TianJiPaiTemplate, bool)
}

type secretCardService struct {
	tianJiPaiMap      map[int32]*gametemplate.TianJiPaiTemplate
	tianJiPaiPoolMap  map[types.SecretCardPoolType][]*gametemplate.TianJiPaiTemplate
	starMap           map[int32]*gametemplate.TianJiPaiStarTemplate
	tianJiPaiLevepMap map[int32]map[int32]*gametemplate.TianJiPaiTemplate
}

//初始化
func (s *secretCardService) init() error {
	s.starMap = make(map[int32]*gametemplate.TianJiPaiStarTemplate)
	s.tianJiPaiMap = make(map[int32]*gametemplate.TianJiPaiTemplate)
	s.tianJiPaiPoolMap = make(map[types.SecretCardPoolType][]*gametemplate.TianJiPaiTemplate)
	s.tianJiPaiLevepMap = make(map[int32]map[int32]*gametemplate.TianJiPaiTemplate)
	for _, tempCardTemplate := range template.GetTemplateService().GetAll((*gametemplate.TianJiPaiTemplate)(nil)) {
		cardTemplate := tempCardTemplate.(*gametemplate.TianJiPaiTemplate)

		s.tianJiPaiMap[int32(cardTemplate.TemplateId())] = cardTemplate

		poolTyp := cardTemplate.GetPoolTyp()
		poolList := s.tianJiPaiPoolMap[poolTyp]
		poolList = append(poolList, cardTemplate)
		s.tianJiPaiPoolMap[poolTyp] = poolList

		tianJiPaiLevelMap, exist := s.tianJiPaiLevepMap[cardTemplate.QuestId]
		if !exist {
			tianJiPaiLevelMap = make(map[int32]*gametemplate.TianJiPaiTemplate)
			s.tianJiPaiLevepMap[cardTemplate.QuestId] = tianJiPaiLevelMap
		}
		tianJiPaiLevelMap[cardTemplate.LevelMin] = cardTemplate
	}

	for _, tempStarTemplate := range template.GetTemplateService().GetAll((*gametemplate.TianJiPaiStarTemplate)(nil)) {
		starTemplate := tempStarTemplate.(*gametemplate.TianJiPaiStarTemplate)
		s.starMap[int32(starTemplate.TemplateId())] = starTemplate
	}

	//校验轮询池
	pollList, exist := s.tianJiPaiPoolMap[types.SecretCardPoolTypePoll]
	if !exist {
		return fmt.Errorf("secretcard: 轮询池任务应该存在的")
	}
	shouldPollLen := int32(0)
	for _, pollObj := range pollList {
		if pollObj.ModuleOpenedId == 0 {
			shouldPollLen++
		}
	}
	if shouldPollLen < 3 {
		return fmt.Errorf("secretcard: 轮询池至少存在三个跟功能开启无关的任务")
	}
	return nil
}

//获取天机牌玩家最小等级的
func (s *secretCardService) GetSecretCardTemplateByQuestId(questId int32) (tianJiPaiTemplate *gametemplate.TianJiPaiTemplate, flag bool) {
	tianJiPaiLevelMap, exist := s.tianJiPaiLevepMap[questId]
	if !exist {
		return
	}
	minLevel := int32(500)
	for level, _ := range tianJiPaiLevelMap {
		if level < minLevel {
			minLevel = level
		}
	}
	tianJiPaiTemplate, exist = tianJiPaiLevelMap[minLevel]
	if exist {
		flag = true
	}
	return
}

//获取
func (s *secretCardService) GetQuestPool() map[types.SecretCardPoolType][]*gametemplate.TianJiPaiTemplate {
	return s.tianJiPaiPoolMap
}

//获取任务对应任务池
func (s *secretCardService) GetQuestPoolType(id int32) (typ types.SecretCardPoolType, flag bool) {
	tempTemplate, exist := s.tianJiPaiMap[id]
	if !exist {
		return
	}
	typ = tempTemplate.GetPoolTyp()
	flag = true
	return
}

//获取掉落id
func (s *secretCardService) GetDropIdByNum(cardId int32, num int32) (dropId int32) {
	if num < 1 {
		return
	}

	to := s.GetSecretCardTemplate(cardId)
	if num <= 5 {
		speDropList := to.GetSpeDropList()
		return speDropList[num-1]
	}

	dropIdMap := to.GetIntervalDropMap()
	for _, value := range to.GetDropIntervalList() {
		ret := num % int32(value)
		if ret == 0 {
			return dropIdMap[int32(value)]
		}
	}

	return

}

//获取天机牌配置
func (s *secretCardService) GetSecretCardTemplate(cardId int32) *gametemplate.TianJiPaiTemplate {
	to, exist := s.tianJiPaiMap[cardId]
	if !exist {
		return nil
	}
	return to
}

//获取星数奖励
func (s *secretCardService) GetStarTemplate(boxId int32) *gametemplate.TianJiPaiStarTemplate {
	to, exist := s.starMap[boxId]
	if !exist {
		return nil
	}
	return to
}

//星数随机
func (s *secretCardService) GetRandomStar(idList []int32) map[int32]int32 {
	cardMap := make(map[int32]int32)
	for _, id := range idList {
		to := s.GetSecretCardTemplate(id)
		if to == nil {
			continue
		}
		min := to.StarMin
		max := to.StarMax
		star := int32(mathutils.RandomRange(int(min), int(max)))
		cardMap[id] = star
	}
	return cardMap
}

//获取剩余未开启的宝箱数
func (s *secretCardService) GetLeftBoxTemplate(openedBox []int32) (starTemplateList []*gametemplate.TianJiPaiStarTemplate) {
	for boxId, starTemplate := range s.starMap {
		if len(openedBox) > 0 {
			flag := utils.ContainInt32(openedBox, boxId)
			if flag {
				continue
			}
		}
		starTemplateList = append(starTemplateList, starTemplate)
	}
	return
}

//获取任务id
func (s *secretCardService) GetQuestIdByCardId(cardId int32) (questId int32, questName string, flag bool) {
	to, exist := s.tianJiPaiMap[cardId]
	if !exist {
		return
	}
	questId = to.QuestId
	questName = to.Remarks
	flag = true
	return
}

//天机牌每日翻牌次数
func (s *secretCardService) GetConstSecretCardNum() int32 {
	return constant.GetConstantService().GetConstant(constanttypes.ConstantTypeSecretCardNum)
}

//天机牌一键完成消耗的元宝数量
func (s *secretCardService) GetConstSecretCardCostGold() int32 {
	return constant.GetConstantService().GetConstant(constanttypes.ConstantTypeSecretCardCostGold)
}

var (
	once sync.Once
	cs   *secretCardService
)

func Init() (err error) {
	once.Do(func() {
		cs = &secretCardService{}
		err = cs.init()
	})
	return err
}

func GetSecretCardService() SecretCardService {
	return cs
}
