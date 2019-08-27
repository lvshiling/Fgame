package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"sync"
)

type QuestTemplateService interface {
	//获取任务模板通过questId
	GetQuestTemplate(questId int32) *gametemplate.QuestTemplate
	//获取屠魔任务类型
	GetQuestTypeById(questId int32) questtypes.QuestType
	//获取屠魔任务id
	GetQuestIdForTuMo(questLevel questtypes.QuestLevelType, usedQuestList []int32, questTb int32, level int32) (int32, bool)
	//获取一键完成橙色屠魔任务id
	GetQuestIdListForTuMoFinishAll(needNum int32, level int32) []int32
	//屠魔任务次数额外购买vip限制
	GetQuestTuMoBuyNumVipLimit() int32
	//屠魔任务VIP额外购买次数
	GetQuestTuMoVipAddBuyNum() int32
	//屠魔任务初始次数
	GetQuestTuMoInitialNum() int32
	//屠魔任务栏默认开启个数
	GetTuMoTaskBarDefaultNum() int32
	//屠魔任务栏第三栏开启所需人物等级
	GetTuMoTaskBarOpenThirdPlayerLevel() int32
	//屠魔任务栏第四栏所需VIP等级
	GetTuMoTaskBarOpenFourthVipLevel() int32
	//屠魔任务每次购买消耗元宝
	GetQuestTuMoBuyNumCostGold() int64
	//屠魔任务一键完成单次消耗
	GetQuestTuMoFinishCostGold() int32
	//屠魔任务直接完成消耗的绑元
	GetQuestTuMoImmediateFinishCostGold() int32
	//日环任务直接完成消耗的绑元
	GetQuestDailyImmediateFinish(dailyTag questtypes.QuestDailyTag) int32
	//日环任务提交双倍奖励
	GetQuestDailyCommitDouble(dailyTag questtypes.QuestDailyTag) int32
	//获取活跃度任务
	GetQuestHuoYueMap() map[int32]*gametemplate.QuestTemplate
	//日环任务
	GetQuestDailyTemplateBySeq(dailyTag questtypes.QuestDailyTag, seqId int32) gametemplate.DailyTagTemplate
	//日环任务一键完成
	GetQuestDailyFinishAll(pl player.Player, dailyTag questtypes.QuestDailyTag, times questtypes.QuestDailyType) (seqIdList []int32, seqId int32)

	//获取日环任务
	GetQuestDailyTemplate(pl player.Player, dailyTag questtypes.QuestDailyTag, times questtypes.QuestDailyType) (dailyTemplate gametemplate.DailyTagTemplate, flag bool)
	//获取日环任务数量上限
	GetQuestDailyMaxNum(dailyTag questtypes.QuestDailyTag) int32
	//获取开服目标任务
	GetKaiFuMuBiaoMap() map[int32]*gametemplate.KaiFuMuBiaoTemplate
	//获取开服目标模板
	GetKaiFuMuBiaoTemplate(kaiFuDay int32) *gametemplate.KaiFuMuBiaoTemplate
	//获取开服解锁天数根据任务id
	GetKaiFuMuBiaoKaiFuDay(questId int32) (kaiFuDayList []int32, flag bool)
	//奇遇任务
	GetQiYuTemplate(qiyuId int32) *gametemplate.QiYuTemplate
	GetQiYuTemplateAll() map[int32]*gametemplate.QiYuTemplate
	GetQiYuIdByQuestId(questId int32) int32
	//运营活动目标任务
	GetYunYingGoalTemplate(goalId int32) *gametemplate.YunYingGoalTemplate
}
type questTemplateService struct {
	//任务map
	questTemplateMap map[int32]*gametemplate.QuestTemplate
	//屠魔任务map
	questTuMoTemplateMap map[questtypes.QuestLevelType]map[int32][]*gametemplate.QuestTemplate
	//活跃度任务map
	huoYueTemplateMap map[int32]*gametemplate.QuestTemplate
	//日环任务map
	dailyTemplateMap map[questtypes.QuestDailyTag]map[questtypes.QuestDailyType]map[int32][]gametemplate.DailyTagTemplate
	dailyMap         map[questtypes.QuestDailyTag]map[int32]gametemplate.DailyTagTemplate
	// dailyTimesRangeListMap map[questtypes.QuestDailyTag][]gametemplate.DailyTagTemplate
	//日环任务概率map
	dailyRateMap map[questtypes.QuestDailyTag]map[questtypes.QuestDailyType]map[int32][]int64

	//开服目标任务map
	kaiFuMuBiaoMap map[int32]*gametemplate.KaiFuMuBiaoTemplate
	//奇遇任务
	qiyuMap map[int32]*gametemplate.QiYuTemplate
	//运营活动目标任务
	yunYingGoalMap map[int32]*gametemplate.YunYingGoalTemplate
}

func (qs *questTemplateService) init() (err error) {
	qs.questTemplateMap = make(map[int32]*gametemplate.QuestTemplate)
	qs.questTuMoTemplateMap = make(map[questtypes.QuestLevelType]map[int32][]*gametemplate.QuestTemplate)
	qs.huoYueTemplateMap = make(map[int32]*gametemplate.QuestTemplate)
	qs.dailyTemplateMap = make(map[questtypes.QuestDailyTag]map[questtypes.QuestDailyType]map[int32][]gametemplate.DailyTagTemplate)
	qs.dailyRateMap = make(map[questtypes.QuestDailyTag]map[questtypes.QuestDailyType]map[int32][]int64)
	qs.dailyMap = make(map[questtypes.QuestDailyTag]map[int32]gametemplate.DailyTagTemplate)
	//qs.dailyTimesRangeList = make([]*gametemplate.DailyTemplate, 0, questtypes.QuestDailyTypeMax)
	// qs.dailyTimesRangeListMap = make(map[questtypes.QuestDailyTag][]gametemplate.DailyTagTemplate)
	qs.kaiFuMuBiaoMap = make(map[int32]*gametemplate.KaiFuMuBiaoTemplate)
	qs.qiyuMap = make(map[int32]*gametemplate.QiYuTemplate)
	qs.yunYingGoalMap = make(map[int32]*gametemplate.YunYingGoalTemplate)

	err = qs.initQuest()
	if err != nil {
		return
	}

	err = qs.initDaily()
	if err != nil {
		return
	}

	err = qs.initDailyAlliance()
	if err != nil {
		return
	}

	err = qs.initKaiFuMuBiao()
	if err != nil {
		return
	}

	err = qs.checkData()
	if err != nil {
		return
	}

	err = qs.initQiYu()
	if err != nil {
		return
	}

	err = qs.initYunYingGoal()
	if err != nil {
		return
	}
	return
}

func (qs *questTemplateService) initQuest() (err error) {
	for _, tempQuestTemplate := range template.GetTemplateService().GetAll((*gametemplate.QuestTemplate)(nil)) {
		questTemplate := tempQuestTemplate.(*gametemplate.QuestTemplate)
		//任务map
		_, exist := qs.questTemplateMap[int32(questTemplate.TemplateId())]
		if exist {
			return fmt.Errorf("questtemplateservice:重复任务questId:%d", questTemplate.TemplateId())
		}
		qs.questTemplateMap[int32(questTemplate.TemplateId())] = questTemplate

		//赋值屠魔任务map
		typ := questTemplate.GetQuestType()
		switch typ {
		case questtypes.QuestTypeTuMo:
			{
				//任务品质
				questLevel := questTemplate.GetQuestLevel()
				tuMoTbTemplateMap, ok := qs.questTuMoTemplateMap[questLevel]
				if !ok {
					tuMoTbTemplateMap = make(map[int32][]*gametemplate.QuestTemplate)
					qs.questTuMoTemplateMap[questLevel] = tuMoTbTemplateMap
				}
				//任务标记
				tuMoTbTemplateList := tuMoTbTemplateMap[questTemplate.QuestTb]
				tuMoTbTemplateList = append(tuMoTbTemplateList, questTemplate)
				tuMoTbTemplateMap[questTemplate.QuestTb] = tuMoTbTemplateList
				break
			}
		case questtypes.QuestTypeLiveness:
			{
				qs.huoYueTemplateMap[int32(questTemplate.TemplateId())] = questTemplate
				break
			}
		}
	}
	return
}

func (qs *questTemplateService) initDaily() (err error) {
	for _, tempDailyTemplate := range template.GetTemplateService().GetAll((*gametemplate.DailyTemplate)(nil)) {
		dailyTemplate := tempDailyTemplate.(*gametemplate.DailyTemplate)
		dailyType := dailyTemplate.GetDailyTimesType()
		//赋值 dailyTemplateMap
		dailyTemplateMap, ok := qs.dailyTemplateMap[questtypes.QuestDailyTagPerson]
		if !ok {
			dailyTemplateMap = make(map[questtypes.QuestDailyType]map[int32][]gametemplate.DailyTagTemplate)
			qs.dailyTemplateMap[questtypes.QuestDailyTagPerson] = dailyTemplateMap
		}
		dailyTimeLevelMap, ok := dailyTemplateMap[dailyType]
		if !ok {
			dailyTimeLevelMap = make(map[int32][]gametemplate.DailyTagTemplate)
			dailyTemplateMap[dailyType] = dailyTimeLevelMap
		}
		dailyTimesList := dailyTimeLevelMap[dailyTemplate.LevelMin]
		dailyTimesList = append(dailyTimesList, dailyTemplate)
		dailyTimeLevelMap[dailyTemplate.LevelMin] = dailyTimesList

		//赋值 dailyRateMap
		percent := int64(dailyTemplate.Percent)
		dailyRateMap, ok := qs.dailyRateMap[questtypes.QuestDailyTagPerson]
		if !ok {
			dailyRateMap = make(map[questtypes.QuestDailyType]map[int32][]int64)
			qs.dailyRateMap[questtypes.QuestDailyTagPerson] = dailyRateMap
		}
		dailyRateLevelMap, ok := dailyRateMap[dailyType]
		if !ok {
			dailyRateLevelMap = make(map[int32][]int64)
			dailyRateMap[dailyType] = dailyRateLevelMap
		}
		dailyRateList := dailyRateLevelMap[dailyTemplate.LevelMin]
		dailyRateList = append(dailyRateList, percent)
		dailyRateLevelMap[dailyTemplate.LevelMin] = dailyRateList

		//赋值 dailyMap
		dailyMap, ok := qs.dailyMap[questtypes.QuestDailyTagPerson]
		if !ok {
			dailyMap = make(map[int32]gametemplate.DailyTagTemplate)
			qs.dailyMap[questtypes.QuestDailyTagPerson] = dailyMap
		}
		dailyMap[int32(dailyTemplate.TemplateId())] = dailyTemplate

		qs.initDailyTimes(questtypes.QuestDailyTagPerson, dailyTemplate)
	}

	return
}

func (qs *questTemplateService) initDailyAlliance() (err error) {
	for _, tempDailyTemplate := range template.GetTemplateService().GetAll((*gametemplate.UnionRiChangTemplate)(nil)) {
		dailyTemplate := tempDailyTemplate.(*gametemplate.UnionRiChangTemplate)
		dailyType := dailyTemplate.GetDailyTimesType()
		//赋值 dailyTemplateMap
		dailyTemplateMap, ok := qs.dailyTemplateMap[questtypes.QuestDailyTagAlliance]
		if !ok {
			dailyTemplateMap = make(map[questtypes.QuestDailyType]map[int32][]gametemplate.DailyTagTemplate)
			qs.dailyTemplateMap[questtypes.QuestDailyTagAlliance] = dailyTemplateMap
		}
		dailyTimeLevelMap, ok := dailyTemplateMap[dailyType]
		if !ok {
			dailyTimeLevelMap = make(map[int32][]gametemplate.DailyTagTemplate)
			dailyTemplateMap[dailyType] = dailyTimeLevelMap
		}
		dailyTimesList := dailyTimeLevelMap[dailyTemplate.LevelMin]
		dailyTimesList = append(dailyTimesList, dailyTemplate)
		dailyTimeLevelMap[dailyTemplate.LevelMin] = dailyTimesList

		//赋值 dailyRateMap
		percent := int64(dailyTemplate.Percent)
		dailyRateMap, ok := qs.dailyRateMap[questtypes.QuestDailyTagAlliance]
		if !ok {
			dailyRateMap = make(map[questtypes.QuestDailyType]map[int32][]int64)
			qs.dailyRateMap[questtypes.QuestDailyTagAlliance] = dailyRateMap
		}
		dailyRateLevelMap, ok := dailyRateMap[dailyType]
		if !ok {
			dailyRateLevelMap = make(map[int32][]int64)
			dailyRateMap[dailyType] = dailyRateLevelMap
		}
		dailyRateList := dailyRateLevelMap[dailyTemplate.LevelMin]
		dailyRateList = append(dailyRateList, percent)
		dailyRateLevelMap[dailyTemplate.LevelMin] = dailyRateList

		//赋值 dailyMap
		dailyMap, ok := qs.dailyMap[questtypes.QuestDailyTagAlliance]
		if !ok {
			dailyMap = make(map[int32]gametemplate.DailyTagTemplate)
			qs.dailyMap[questtypes.QuestDailyTagAlliance] = dailyMap
		}
		dailyMap[int32(dailyTemplate.TemplateId())] = dailyTemplate

		qs.initDailyTimes(questtypes.QuestDailyTagAlliance, dailyTemplate)
	}

	return
}

func (qs *questTemplateService) initKaiFuMuBiao() (err error) {
	for _, tempKaiFuMuBiaoTemplate := range template.GetTemplateService().GetAll((*gametemplate.KaiFuMuBiaoTemplate)(nil)) {
		kaiFuMuBiaoTemplate := tempKaiFuMuBiaoTemplate.(*gametemplate.KaiFuMuBiaoTemplate)
		kaiFuTime := kaiFuMuBiaoTemplate.KaiFuTime
		qs.kaiFuMuBiaoMap[kaiFuTime] = kaiFuMuBiaoTemplate
	}
	return
}

func (qs *questTemplateService) initQiYu() (err error) {
	for _, to := range template.GetTemplateService().GetAll((*gametemplate.QiYuTemplate)(nil)) {
		qiYuTemplate := to.(*gametemplate.QiYuTemplate)
		qs.qiyuMap[int32(qiYuTemplate.Id)] = qiYuTemplate
	}
	return
}

func (qs *questTemplateService) initYunYingGoal() (err error) {
	for _, to := range template.GetTemplateService().GetAll((*gametemplate.YunYingGoalTemplate)(nil)) {
		temp := to.(*gametemplate.YunYingGoalTemplate)
		qs.yunYingGoalMap[int32(temp.Id)] = temp
	}
	return
}

func (qs *questTemplateService) initDailyTimes(dailyTag questtypes.QuestDailyTag, dailyTempalte gametemplate.DailyTagTemplate) {
	// dailyTimeRangeList := qs.dailyTimesRangeListMap[dailyTag]
	// if len(dailyTimeRangeList) == 0 {
	// 	dailyTimeRangeList = append(dailyTimeRangeList, dailyTempalte)
	// 	qs.dailyTimesRangeListMap[dailyTag] = dailyTimeRangeList
	// }

	// for _, tempTempalte := range dailyTimeRangeList {
	// 	if tempTempalte.GetTimesMin() == dailyTempalte.GetTimesMin() {
	// 		return
	// 	}
	// }
	// dailyTimeRangeList = append(dailyTimeRangeList, dailyTempalte)
	// qs.dailyTimesRangeListMap[dailyTag] = dailyTimeRangeList
}

func (qs *questTemplateService) checkData() (err error) {
	tuMoLingMap := item.GetItemService().GetItemClassMap(itemtypes.ItemTypeTuMoLing)
	if tuMoLingMap == nil {
		return fmt.Errorf("questtemplateservice:屠魔令配置应该是存在的")
	}
	for questLevel, _ := range questtypes.QuestTuMoMap {
		//校验屠魔令
		itemSubType := questLevel.ItemTumoSubType()
		_, exist := tuMoLingMap[itemSubType]
		if !exist {
			return fmt.Errorf("questtemplateservice:屠魔令配置应该是存在的")
		}
		//验证任务品质
		_, exist = qs.questTuMoTemplateMap[questLevel]
		if !exist {
			return fmt.Errorf("questtemplateservice:屠魔任务品质配置不全")
		}
	}
	err = qs.checkKaiFuMuBiao()
	if err != nil {
		return
	}

	return nil
}

func (qs *questTemplateService) checkKaiFuMuBiao() (err error) {
	for _, kaiFuMuBiaoTemplate := range qs.kaiFuMuBiaoMap {
		for questId, _ := range kaiFuMuBiaoTemplate.GetQuestMap() {
			for _, otherKaiFuMuBiaoTemplate := range qs.kaiFuMuBiaoMap {
				for otherQuestId, _ := range otherKaiFuMuBiaoTemplate.GetQuestMap() {
					if questId == otherQuestId &&
						kaiFuMuBiaoTemplate != otherKaiFuMuBiaoTemplate {
						return fmt.Errorf("questtemplateservice:开服目标任务id重复:%d", questId)
					}
				}
			}
		}
	}
	return
}

//获取任务模板通过questId
func (qs *questTemplateService) GetQuestTemplate(questId int32) *gametemplate.QuestTemplate {
	q, exist := qs.questTemplateMap[questId]
	if !exist {
		return nil
	}
	return q
}

//获取活跃度任务map
func (qs *questTemplateService) GetQuestHuoYueMap() map[int32]*gametemplate.QuestTemplate {
	return qs.huoYueTemplateMap
}

//获取开服目标任务
func (qs *questTemplateService) GetKaiFuMuBiaoMap() map[int32]*gametemplate.KaiFuMuBiaoTemplate {
	return qs.kaiFuMuBiaoMap
}

//获取屠魔任务类型
func (qs *questTemplateService) GetQuestTypeById(questId int32) questtypes.QuestType {
	to := qs.GetQuestTemplate(questId)
	if to == nil {
		return 0
	}
	return to.GetQuestType()
}

//获取屠魔任务id
func (qs *questTemplateService) GetQuestIdForTuMo(questLevel questtypes.QuestLevelType, usedQuestIdList []int32, questTb int32, level int32) (int32, bool) {
	if len(usedQuestIdList) == 0 {
		if questTb != 0 {
			panic(fmt.Errorf("quest: when len(usedQuestIdList)==0,questTb should be 0 "))
		}
	}
	//获取任务品质map
	tuMoTbTemplateMap, exist := qs.questTuMoTemplateMap[questLevel]
	if !exist {
		return 0, false
	}
	//(questTb =0跨级或未接取questlevel品质) || 非橙色的纯随机
	if questTb == 0 || questLevel != questtypes.QuestLevelTypeTuMoOrange {
		questId, flag := qs.getTuMoQuestIdByLevel(tuMoTbTemplateMap, usedQuestIdList, level)
		if !flag {
			return 0, false
		}
		return questId, true
	}
	//橙色的接取同questTb
	tuMoLevelTemplateMap, exist := tuMoTbTemplateMap[questTb]
	if !exist {
		return 0, false
	}
	questId, flag := qs.getTuMoQuestIdByLevelAndTb(tuMoLevelTemplateMap, usedQuestIdList, level)
	if !flag {
		return 0, false
	}
	return questId, true
}

//获取屠魔任务id根据玩家等级
func (qs *questTemplateService) getTuMoQuestIdByLevel(questMap map[int32][]*gametemplate.QuestTemplate, usedQuestIdList []int32, level int32) (int32, bool) {
	questIdList := make([]int32, 0, 16)
	weightList := make([]int64, 0, 16)
	for _, questList := range questMap {
		for _, questTemplate := range questList {
			minLevel := questTemplate.MinLevel
			maxLevel := questTemplate.MaxLevel
			if level < minLevel || level > maxLevel {
				continue
			}
			questId := int32(questTemplate.TemplateId())
			if len(usedQuestIdList) != 0 {
				flag := utils.ContainInt32(usedQuestIdList, questId)
				if flag {
					continue
				}
			}
			questIdList = append(questIdList, questId)
			weightList = append(weightList, 1)
		}
	}
	if len(questIdList) == 0 {
		return 0, false
	}
	index := mathutils.RandomWeights(weightList)
	return questIdList[index], true
}

//获取屠魔任务id根据玩家等级和tb
func (qs *questTemplateService) getTuMoQuestIdByLevelAndTb(questLevelList []*gametemplate.QuestTemplate, usedQuestIdList []int32, level int32) (int32, bool) {
	questIdList := make([]int32, 0, 16)
	weightList := make([]int64, 0, 16)
	for _, questObj := range questLevelList {
		minLevel := questObj.MinLevel
		maxLevel := questObj.MaxLevel
		if level < minLevel || level > maxLevel {
			continue
		}
		questId := int32(questObj.TemplateId())
		if len(usedQuestIdList) != 0 {
			flag := utils.ContainInt32(usedQuestIdList, questId)
			if flag {
				continue
			}
		}
		questIdList = append(questIdList, questId)
		weightList = append(weightList, 1)
	}
	if len(questIdList) == 0 {
		return 0, false
	}
	index := mathutils.RandomWeights(weightList)
	return questIdList[index], true
}

//获取一键完成橙色屠魔任务id
func (qs *questTemplateService) GetQuestIdListForTuMoFinishAll(needNum int32, level int32) (questIdList []int32) {
	questTbMap, _ := qs.questTuMoTemplateMap[questtypes.QuestLevelTypeTuMoOrange]
	pollIdList := make([]int32, 0, 8)
	for _, questList := range questTbMap {
		for _, questTemplate := range questList {
			minLevel := questTemplate.MinLevel
			maxLevel := questTemplate.MaxLevel
			if level < minLevel || level > maxLevel {
				continue
			}
			pollIdList = append(pollIdList, int32(questTemplate.TemplateId()))
		}
	}
	if len(pollIdList) == 0 {
		panic(fmt.Errorf("quest: 玩家等级:%d橙色任务配置不存在", level))
	}

	for curSize := int32(len(questIdList)); curSize < needNum; {
		listPoll := mathutils.RandomList(pollIdList, needNum-curSize)
		for i := 0; i < len(listPoll); i++ {
			questIdList = append(questIdList, listPoll[i])
		}
		curSize = int32(len(questIdList))
	}
	return questIdList
}

//屠魔任务次数额外购买vip限制
func (qs *questTemplateService) GetQuestTuMoBuyNumVipLimit() int32 {
	minVipLimit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuMoTaskBuyNumVipLimit)
	return minVipLimit
}

//屠魔任务VIP额外购买次数
func (qs *questTemplateService) GetQuestTuMoVipAddBuyNum() int32 {
	maxBuyLimit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuMoTaskVipAddBuyNum)
	return maxBuyLimit
}

//屠魔任务初始次数
func (qs *questTemplateService) GetQuestTuMoInitialNum() int32 {
	initNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuMoTaskInitialNum)
	return initNum
}

//屠魔任务每次购买消耗元宝
func (qs *questTemplateService) GetQuestTuMoBuyNumCostGold() int64 {
	costGold := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuMoTaskBuyNumCostGold)
	return int64(costGold)
}

//屠魔任务栏默认开启个数
func (qs *questTemplateService) GetTuMoTaskBarDefaultNum() int32 {
	taskBarDefaultNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuMoTaskBarDefaultNum)
	return taskBarDefaultNum
}

//屠魔任务栏第三栏开启所需人物等级
func (qs *questTemplateService) GetTuMoTaskBarOpenThirdPlayerLevel() int32 {
	openThirdBarLevel := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuMoTaskBarOpenThirdPlayerLevel)
	return openThirdBarLevel
}

//屠魔任务栏第四栏所需VIP等级
func (qs *questTemplateService) GetTuMoTaskBarOpenFourthVipLevel() int32 {
	openFourthBarVip := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuMoTaskBarOpenFourthVipLevel)
	return openFourthBarVip
}

//屠魔任务一键完成单次消耗
func (qs *questTemplateService) GetQuestTuMoFinishCostGold() int32 {
	finishCostGold := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuMoTaskFinishCostGold)
	return finishCostGold
}

//屠魔任务直接完成消耗的绑元
func (qs *questTemplateService) GetQuestTuMoImmediateFinishCostGold() int32 {
	finishCostGold := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuMoImmediateFinishCostGold)
	return finishCostGold
}

//日环任务直接完成消耗的绑元
func (qs *questTemplateService) GetQuestDailyImmediateFinish(dailyTag questtypes.QuestDailyTag) int32 {
	constantType := constanttypes.ConstantTypeDailyFinishBindGold
	switch dailyTag {
	case questtypes.QuestDailyTagPerson:
		constantType = constanttypes.ConstantTypeDailyFinishBindGold
	case questtypes.QuestDailyTagAlliance:
		constantType = constanttypes.ConstantTypeAllianceDailyFinishBindGold
	}
	finishCostGold := constant.GetConstantService().GetConstant(constanttypes.ConstantType(constantType))
	return finishCostGold
}

//日环任务数量限制
func (qs *questTemplateService) GetQuestDailyMaxNum(dailyTag questtypes.QuestDailyTag) int32 {
	maxDailyNum := int32(0)
	switch dailyTag {
	case questtypes.QuestDailyTagAlliance:
		maxDailyNum = constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceDailyCount)
	case questtypes.QuestDailyTagPerson:
		maxDailyNum = constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDailyQuestNumLimit)
	}

	return maxDailyNum
}

//日环任务提交双倍奖励
func (qs *questTemplateService) GetQuestDailyCommitDouble(dailyTag questtypes.QuestDailyTag) int32 {
	constantType := constanttypes.ConstantTypeDailyCommitBindGold
	switch dailyTag {
	case questtypes.QuestDailyTagPerson:
		constantType = constanttypes.ConstantTypeDailyCommitBindGold
	case questtypes.QuestDailyTagAlliance:
		constantType = constanttypes.ConstantTypeAllianceDailyCommitBindGold
	}

	finishCostGold := constant.GetConstantService().GetConstant(constanttypes.ConstantType(constantType))
	return finishCostGold
}

func (qs *questTemplateService) GetQuestDailyTemplateBySeq(dailyTag questtypes.QuestDailyTag, seqId int32) gametemplate.DailyTagTemplate {
	dailyMap, ok := qs.dailyMap[dailyTag]
	if !ok {
		return nil
	}
	return dailyMap[seqId]
}

//日环任务一键完成
func (qs *questTemplateService) GetQuestDailyFinishAll(pl player.Player, dailyTag questtypes.QuestDailyTag, times questtypes.QuestDailyType) (seqIdList []int32, seqId int32) {
	maxDailyType := questtypes.QuestDailyType(qs.GetQuestDailyMaxNum(dailyTag))
	for nextTime := times + 1; nextTime <= maxDailyType; nextTime++ {
		dailyTemplate, flag := qs.GetQuestDailyTemplate(pl, dailyTag, nextTime)
		if !flag {
			continue
		}
		seqId = int32(dailyTemplate.TemplateId())
		seqIdList = append(seqIdList, seqId)
	}
	return
}

// 任务区间起始次数
func (qs *questTemplateService) getDailyTimesRange(dailyTag questtypes.QuestDailyTag, lev int32, needTimes questtypes.QuestDailyType) (times questtypes.QuestDailyType, flag bool) {
	// dailyTimesRangeList, ok := qs.dailyTimesRangeListMap[dailyTag]
	// if !ok {
	// 	return
	// }
	// for _, dailyTemplate := range dailyTimesRangeList {
	// 	if int32(needTimes) >= dailyTemplate.GetTimesMin() && int32(needTimes) <= dailyTemplate.GetTimesMax() {
	// 		flag = true
	// 		times = questtypes.QuestDailyType(dailyTemplate.GetTimesMin())
	// 		return
	// 	}
	// }
	dailyTemplateMap, ok := qs.dailyTemplateMap[dailyTag]
	if !ok {
		return
	}
	for timesMin, dailyLevelTemplateMap := range dailyTemplateMap {
		if timesMin > needTimes {
			continue
		}

		for levMin, dailyTemplateList := range dailyLevelTemplateMap {
			if levMin > lev {
				continue
			}
			for _, dailyTemplate := range dailyTemplateList {
				if int32(needTimes) > dailyTemplate.GetTimesMax() {
					continue
				}
				if lev > dailyTemplate.GetLevelMax() {
					continue
				}
				return timesMin, true
			}
		}
	}
	return
}

//获取日环任务
func (qs *questTemplateService) GetQuestDailyTemplate(pl player.Player, dailyTag questtypes.QuestDailyTag, needTimes questtypes.QuestDailyType) (dailyTemplate gametemplate.DailyTagTemplate, flag bool) {
	level := pl.GetLevel()

	// 任务区间起始次数
	times, flag := qs.getDailyTimesRange(dailyTag, level, needTimes)
	if !flag {
		return
	}

	// 次数区间所有日环任务
	dailyTemplateMap, ok := qs.dailyTemplateMap[dailyTag]
	if !ok {
		return
	}
	dailyLevelTemplateMap, ok := dailyTemplateMap[times]
	if !ok {
		return
	}

	// 概率
	dailyRateMap, ok := qs.dailyRateMap[dailyTag]
	if !ok {
		return
	}
	dailyLevelRateMap, ok := dailyRateMap[times]
	if !ok {
		return
	}

	levelIndex := int32(0)

Loop:
	for _, dailyLevelList := range dailyLevelTemplateMap {
		for _, dailyLevelTemplate := range dailyLevelList {
			if level >= dailyLevelTemplate.GetLevelMin() && level <= dailyLevelTemplate.GetLevelMax() {
				levelIndex = dailyLevelTemplate.GetLevelMin()
				break Loop
			}
			break
		}
	}

	// 等级区间所有日环任务
	dailyLevelTemplateList := dailyLevelTemplateMap[levelIndex]
	if len(dailyLevelTemplateList) == 0 {
		return
	}

	dailyLevelList := dailyLevelRateMap[levelIndex]
	index := mathutils.RandomWeights(dailyLevelList)
	if index == -1 {
		return
	}
	flag = true
	dailyTemplate = dailyLevelTemplateList[index]
	return
}

//获取开服目标模板
func (qs *questTemplateService) GetKaiFuMuBiaoTemplate(kaiFuDay int32) *gametemplate.KaiFuMuBiaoTemplate {
	kaiFuMuBiaoTemplate, ok := qs.kaiFuMuBiaoMap[kaiFuDay]
	if !ok {
		return nil
	}
	return kaiFuMuBiaoTemplate
}

func (qs *questTemplateService) GetKaiFuMuBiaoKaiFuDay(questId int32) (kaiFuDayList []int32, flag bool) {
	for i := 1; i <= questtypes.KaiFuMuBiaoDayMax; i++ {
		kaiFuMuBiaoTempalte, ok := qs.kaiFuMuBiaoMap[int32(i)]
		if !ok {
			continue
		}
		questMap := kaiFuMuBiaoTempalte.GetQuestMap()
		if len(questMap) == 0 {
			continue
		}
		_, ok = questMap[questId]
		if ok {
			flag = true
			kaiFuDayList = append(kaiFuDayList, kaiFuMuBiaoTempalte.KaiFuTime)
		}
	}
	return
}

//
func (qs *questTemplateService) GetQiYuIdByQuestId(questId int32) (qiyuId int32) {
	for _, qiyuTemp := range qs.qiyuMap {
		questMap := qiyuTemp.GetQuestMap()
		if len(questMap) == 0 {
			continue
		}
		_, ok := questMap[questId]
		if ok {
			return int32(qiyuTemp.Id)
		}
	}
	return
}

//获取运营活动目标任务
func (qs *questTemplateService) GetYunYingGoalTemplate(goalId int32) *gametemplate.YunYingGoalTemplate {
	temp, ok := qs.yunYingGoalMap[goalId]
	if !ok {
		return nil
	}
	return temp
}

//获取奇遇任务
func (qs *questTemplateService) GetQiYuTemplate(qiyuId int32) *gametemplate.QiYuTemplate {
	temp, ok := qs.qiyuMap[qiyuId]
	if !ok {
		return nil
	}
	return temp
}

//获取奇遇任务
func (qs *questTemplateService) GetQiYuTemplateAll() map[int32]*gametemplate.QiYuTemplate {
	return qs.qiyuMap
}

var (
	once sync.Once
	cs   *questTemplateService
)

func Init() (err error) {
	once.Do(func() {
		cs = &questTemplateService{}
		err = cs.init()
	})
	return err
}

func GetQuestTemplateService() QuestTemplateService {
	return cs
}
