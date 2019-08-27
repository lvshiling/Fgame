package player

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fmt"

	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/quest/dao"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家任务管理器
type PlayerQuestDataManager struct {
	p player.Player
	//任务对象
	stateQuestMap map[questtypes.QuestState]map[int32]*PlayerQuestObject
	//屠魔任务对象
	tuMoQuestMap map[questtypes.QuestLevelType]map[int32]*PlayerQuestObject
	//屠魔次数对象
	playerTuMoObject *PlayerTuMoObject
	//日环次数对象
	playerDailyObjectMap map[questtypes.QuestDailyTag]*PlayerDailyObject
	//活跃度对象
	playerLivenessObject *PlayerLivenessCrossFiveObject
	//开服目标任务
	playerMuBiaoMap map[int32]*PlayerKaiFuMuBiaoObject
	//任务模块跨天对象
	playerCrossDayObject *PlayerQuestCrossDayObject
	//奇遇任务对象
	playerQiYuMap map[int32]*PlayerQiYuObject
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

func (pqdm *PlayerQuestDataManager) Player() player.Player {
	return pqdm.p
}

//加载
func (pqdm *PlayerQuestDataManager) Load() (err error) {
	//加载玩家任务信息
	pqdm.stateQuestMap = make(map[questtypes.QuestState]map[int32]*PlayerQuestObject)
	pqdm.tuMoQuestMap = make(map[questtypes.QuestLevelType]map[int32]*PlayerQuestObject)

	err = pqdm.loadQuest()
	if err != nil {
		return
	}

	//加载屠魔次数信息
	err = pqdm.loadTuMo()
	if err != nil {
		return
	}

	//活跃度
	err = pqdm.loadLiveness()
	if err != nil {
		return
	}

	//加载日环
	err = pqdm.loadDaily()
	if err != nil {
		return
	}

	//跨12点检查
	err = pqdm.loadQuestCrossDay()
	if err != nil {
		return
	}
	//开服目标
	err = pqdm.loadKaiFuMuBiao()
	if err != nil {
		return
	}
	//奇遇任务
	err = pqdm.loadQiYu()
	if err != nil {
		return
	}
	return nil
}

//加载所有任务
func (pqdm *PlayerQuestDataManager) loadQuest() (err error) {
	questEntityList, err := dao.GetQuestDao().GetQuestList(pqdm.p.GetId())
	if err != nil {
		return
	}

	for _, questEntity := range questEntityList {
		questObj := NewPlayerQuestObject(pqdm.p)
		err = questObj.FromEntity(questEntity)
		if err != nil {
			return
		}
		pqdm.addQuest(questObj)
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questObj.QuestId)
		questType := questTemplate.GetQuestType()
		//屠魔任务列表
		if questType == questtypes.QuestTypeTuMo {
			state := questObj.QuestState
			if state == questtypes.QuestStateAccept ||
				state == questtypes.QuestStateFinish {
				questLevel := questTemplate.GetQuestLevel()
				pqdm.addTuMoQuest(questLevel, questObj)
			}
		}
	}
	return
}

//加载后
func (pqdm *PlayerQuestDataManager) AfterLoad() (err error) {
	err = pqdm.afterLoadMainLine()
	if err != nil {
		return
	}
	err = pqdm.afterLoadTuMoLine()
	if err != nil {
		return
	}
	err = pqdm.afterLoadLivenessQuest()
	if err != nil {
		return
	}

	err = pqdm.afterLoadKaiFuMuBiaoQuest()
	if err != nil {
		return
	}

	pqdm.heartbeatRunner.AddTask(CreateQuestTask(pqdm.p))
	pqdm.heartbeatRunner.AddTask(CreateQiYuTask(pqdm.p))
	return nil
}

func (pqdm *PlayerQuestDataManager) restQuestInit(quest *PlayerQuestObject) {
	now := global.GetGame().GetTimeService().Now()
	//移除
	pqdm.removeQuestByIdAndState(quest.QuestState, quest.QuestId)
	quest.QuestState = questtypes.QuestStateInit
	quest.QuestDataMap = make(map[int32]int32)
	quest.CollectItemDataMap = make(map[int32]int32)
	quest.UpdateTime = now
	quest.SetModified()
	pqdm.addQuest(quest)
}

//添加任务
func (pqdm *PlayerQuestDataManager) AddQuest(questId int32) (questObj *PlayerQuestObject, flag bool) {
	questObj = pqdm.GetQuestById(questId)
	if questObj != nil {
		return
	}
	questObj = createQuest(pqdm.p, questId)
	questObj.SetModified()

	pqdm.addQuest(questObj)
	flag = true
	//TODO 发送事件
	return
}

//添加任务
func (pqdm *PlayerQuestDataManager) addQuest(quest *PlayerQuestObject) {
	m, ok := pqdm.stateQuestMap[quest.QuestState]
	if !ok {
		m = make(map[int32]*PlayerQuestObject)
		pqdm.stateQuestMap[quest.QuestState] = m
	}
	m[quest.QuestId] = quest
}

//获取任务
func (pqdm *PlayerQuestDataManager) GetQuestMap(state questtypes.QuestState) (questMap map[int32]*PlayerQuestObject) {
	questMap, exist := pqdm.stateQuestMap[state]
	if !exist {
		return nil
	}
	return
}

//根据任务获取状态
func (pqdm *PlayerQuestDataManager) GetQuestByIdAndState(state questtypes.QuestState, questId int32) (quest *PlayerQuestObject) {
	questMap, exist := pqdm.stateQuestMap[state]
	if !exist {
		return nil
	}
	quest, exist = questMap[questId]
	if !exist {
		return
	}
	return
}

//根据任务获取状态
func (pqdm *PlayerQuestDataManager) GetTuMoQuestByLevelAndId(questLevel questtypes.QuestLevelType, questId int32) (quest *PlayerQuestObject) {
	questMap, exist := pqdm.tuMoQuestMap[questLevel]
	if !exist {
		return nil
	}
	quest, exist = questMap[questId]
	if !exist {
		return nil
	}
	return
}

//根据任务获取状态
func (pqdm *PlayerQuestDataManager) IsCommit(questId int32) bool {
	questMap, exist := pqdm.stateQuestMap[questtypes.QuestStateCommit]
	if !exist {
		return false
	}
	_, exist = questMap[questId]
	if !exist {
		return false
	}
	return true
}

//移除
func (pqdm *PlayerQuestDataManager) removeQuestByIdAndState(state questtypes.QuestState, questId int32) {
	questMap, exist := pqdm.stateQuestMap[state]
	if !exist {
		return
	}
	delete(questMap, questId)
}

func (pqdm *PlayerQuestDataManager) GetQuestById(questId int32) (quest *PlayerQuestObject) {
	for _, questMap := range pqdm.stateQuestMap {
		quest, exist := questMap[questId]
		if exist {
			return quest
		}
	}
	return
}

func (pqdm *PlayerQuestDataManager) NumOfQuest() int32 {
	num := int32(0)
	for _, questMap := range pqdm.stateQuestMap {
		num += int32(len(questMap))
	}
	return num
}

//是否可以激活任务
func (pqdm *PlayerQuestDataManager) ShouldActiveQuest(questId int32) (flag bool) {
	q := pqdm.GetQuestByIdAndState(questtypes.QuestStateInit, questId)
	if q == nil {
		return
	}
	return true
}

//激活任务
func (pqdm *PlayerQuestDataManager) ActiveQuest(questId int32) (flag bool) {
	if !pqdm.ShouldActiveQuest(questId) {
		return
	}
	q := pqdm.GetQuestByIdAndState(questtypes.QuestStateInit, questId)
	//移除
	pqdm.removeQuestByIdAndState(questtypes.QuestStateInit, questId)

	q.QuestState = questtypes.QuestStateActive
	now := global.GetGame().GetTimeService().Now()
	q.UpdateTime = now
	q.SetModified()

	pqdm.addQuest(q)
	flag = true
	return
}

//是否可以接受任务
func (pqdm *PlayerQuestDataManager) ShouldAcceptQuest(questId int32) (flag bool) {
	q := pqdm.GetQuestByIdAndState(questtypes.QuestStateActive, questId)
	if q == nil {
		return
	}
	return true
}

//接受任务
func (pqdm *PlayerQuestDataManager) AcceptQuest(questId int32) (flag bool) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	if !pqdm.ShouldAcceptQuest(questId) {
		return
	}
	q := pqdm.GetQuestByIdAndState(questtypes.QuestStateActive, questId)
	//移除
	pqdm.removeQuestByIdAndState(questtypes.QuestStateActive, questId)

	q.QuestState = questtypes.QuestStateAccept
	now := global.GetGame().GetTimeService().Now()
	q.UpdateTime = now
	q.SetModified()
	pqdm.addQuest(q)
	flag = true

	gameevent.Emit(questeventtypes.EventTypeQuestAccept, pqdm.p, questId)
	return
}

//完成任务
func (pqdm *PlayerQuestDataManager) ShouldFinishQuest(questId int32) (flag bool) {
	q := pqdm.GetQuestByIdAndState(questtypes.QuestStateAccept, questId)
	if q == nil {
		return
	}
	return true
}

//完成任务
func (pqdm *PlayerQuestDataManager) FinishQuest(questId int32) (flag bool) {
	if !pqdm.ShouldFinishQuest(questId) {
		return
	}
	q := pqdm.GetQuestByIdAndState(questtypes.QuestStateAccept, questId)
	//移除
	pqdm.removeQuestByIdAndState(questtypes.QuestStateAccept, questId)

	q.QuestState = questtypes.QuestStateFinish
	now := global.GetGame().GetTimeService().Now()
	q.UpdateTime = now
	q.SetModified()

	pqdm.addQuest(q)
	flag = true
	gameevent.Emit(questeventtypes.EventTypeQuestFinish, pqdm.p, questId)
	return
}

//是否可以交付任务
func (pqdm *PlayerQuestDataManager) ShouldCommitQuest(questId int32) (flag bool) {
	q := pqdm.GetQuestByIdAndState(questtypes.QuestStateFinish, questId)
	if q == nil {
		return
	}
	return true
}

//交付任务
func (pqdm *PlayerQuestDataManager) CommitQuest(questId int32, isDouble bool) (flag bool) {
	if !pqdm.ShouldCommitQuest(questId) {
		return
	}
	q := pqdm.GetQuestByIdAndState(questtypes.QuestStateFinish, questId)

	//移除
	questType := questtemplate.GetQuestTemplateService().GetQuestTypeById(questId)
	if questType == questtypes.QuestTypeTuMo {
		pqdm.removeTuMoQuestByIdAndState(questtypes.QuestStateFinish, questId)
	} else {
		pqdm.removeQuestByIdAndState(questtypes.QuestStateFinish, questId)
	}

	q.QuestState = questtypes.QuestStateCommit
	now := global.GetGame().GetTimeService().Now()
	q.UpdateTime = now
	q.SetModified()

	pqdm.addQuest(q)
	flag = true
	gameevent.Emit(questeventtypes.EventTypeQuestCommit, pqdm.p, questId)
	return
}

//添加完成数据
func (pqdm *PlayerQuestDataManager) IncreaseQuestData(questId int32, dataId int32, num int32) (flag bool) {
	if num <= 0 {
		panic(fmt.Errorf("quest:添加进度数据[%d]应该大于0", num))
	}
	q := pqdm.GetQuestByIdAndState(questtypes.QuestStateAccept, questId)
	if q == nil {
		return
	}

	oldNum, _ := q.QuestDataMap[dataId]
	q.QuestDataMap[dataId] = oldNum + num
	now := global.GetGame().GetTimeService().Now()
	q.UpdateTime = now
	q.SetModified()

	flag = true
	return
}

//设置完成数据
func (pqdm *PlayerQuestDataManager) SetQuestData(questId int32, dataId int32, num int32) (flag bool) {
	q := pqdm.GetQuestByIdAndState(questtypes.QuestStateAccept, questId)
	if q == nil {
		return
	}

	oldNum, _ := q.QuestDataMap[dataId]
	if oldNum == num {
		flag = true
		return
	}
	q.QuestDataMap[dataId] = num
	now := global.GetGame().GetTimeService().Now()
	q.UpdateTime = now
	q.SetModified()
	flag = true
	return
}

//设置收集数据
func (pqdm *PlayerQuestDataManager) SetCollectItemData(questId int32, dataId int32, num int32) (flag bool) {
	if num <= 0 {
		panic(fmt.Errorf("quest:添加进度数据[%d]应该大于0", num))
	}
	q := pqdm.GetQuestByIdAndState(questtypes.QuestStateAccept, questId)
	if q == nil {
		return
	}

	oldNum, _ := q.CollectItemDataMap[dataId]
	if oldNum >= num {
		return
	}
	q.CollectItemDataMap[dataId] = num
	now := global.GetGame().GetTimeService().Now()
	q.UpdateTime = now
	q.SetModified()
	flag = true
	return
}

//放弃任务
func (pqdm *PlayerQuestDataManager) DiscardQuest(questId int32) (quest *PlayerQuestObject) {
	quest = pqdm.GetQuestByIdAndState(questtypes.QuestStateAccept, questId)
	if quest == nil {
		return
	}
	//移除
	pqdm.removeQuestByIdAndState(questtypes.QuestStateAccept, questId)
	quest.QuestState = questtypes.QuestStateDiscard
	now := global.GetGame().GetTimeService().Now()
	quest.UpdateTime = now
	quest.SetModified()
	pqdm.addQuest(quest)
	return
}

//任务直接完成
func (pqdm *PlayerQuestDataManager) QuestImmediateFinish(questId int32) (qu *PlayerQuestObject, flag bool) {
	now := global.GetGame().GetTimeService().Now()
	quest := pqdm.GetQuestById(questId)
	if quest == nil {
		return
	}
	if quest.QuestState != questtypes.QuestStateAccept {
		return
	}

	q := pqdm.GetQuestByIdAndState(questtypes.QuestStateAccept, questId)
	//移除
	questType := questtemplate.GetQuestTemplateService().GetQuestTypeById(questId)
	if questType == questtypes.QuestTypeTuMo {
		pqdm.removeTuMoQuestByIdAndState(questtypes.QuestStateAccept, questId)
	} else {
		pqdm.removeQuestByIdAndState(questtypes.QuestStateAccept, questId)
	}

	q.QuestState = questtypes.QuestStateCommit
	q.UpdateTime = now
	q.SetModified()

	pqdm.addQuest(q)
	qu = q
	flag = true
	return
}

//任务一键完成
func (pqdm *PlayerQuestDataManager) QuestToKeyComplete(questIdList []int32) (questList []*PlayerQuestObject) {
	for _, questId := range questIdList {
		q, flag := pqdm.QuestImmediateFinish(questId)
		if !flag {
			continue
		}
		questList = append(questList, q)
	}
	return
}

//定时器校验
func (pqdm *PlayerQuestDataManager) QuestReset() (err error) {
	err = pqdm.resetDailyCrossFive()
	if err != nil {
		return
	}
	err = pqdm.resetLivenessCrossFive()
	if err != nil {
		return
	}
	return
}

//心跳
func (pqdm *PlayerQuestDataManager) Heartbeat() {
	pqdm.heartbeatRunner.Heartbeat()
}

func (pqdm *PlayerQuestDataManager) GetCurrentMainQuest() *PlayerQuestObject {
	acceptQuestMap := pqdm.GetQuestMap(questtypes.QuestStateAccept)
	for _, q := range acceptQuestMap {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(q.QuestId)
		if questTemplate == nil {
			continue
		}
		if questTemplate.GetQuestType() == questtypes.QuestTypeOnce {
			return q
		}
	}
	finishQuestMap := pqdm.GetQuestMap(questtypes.QuestStateFinish)
	for _, q := range finishQuestMap {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(q.QuestId)
		if questTemplate == nil {
			continue
		}
		if questTemplate.GetQuestType() == questtypes.QuestTypeOnce {
			return q
		}
	}
	return nil
}

func CreatePlayerQuestDataManager(p player.Player) player.PlayerDataManager {
	pmdm := &PlayerQuestDataManager{}
	pmdm.p = p
	pmdm.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	return pmdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerQuestDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerQuestDataManager))
}

func createQuest(p player.Player, questId int32) *PlayerQuestObject {
	pqo := NewPlayerQuestObject(p)
	id, _ := idutil.GetId()
	pqo.Id = id
	pqo.QuestState = questtypes.QuestStateInit
	now := global.GetGame().GetTimeService().Now()
	pqo.CreateTime = now
	pqo.UpdateTime = now
	pqo.QuestDataMap = make(map[int32]int32)
	pqo.CollectItemDataMap = make(map[int32]int32)
	pqo.QuestId = questId
	return pqo
}
