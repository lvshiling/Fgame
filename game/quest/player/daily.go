package player

import (
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/dao"
	questentity "fgame/fgame/game/quest/entity"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"

	"github.com/pkg/errors"
)

//日环对象
type PlayerDailyObject struct {
	player       player.Player
	id           int64
	dailyTag     questtypes.QuestDailyTag
	seqId        int32
	times        int32
	lastTime     int64
	crossDayTime int64
	updateTime   int64
	createTime   int64
	deleteTime   int64
}

func NewPlayerDailyObject(pl player.Player) *PlayerDailyObject {
	pmo := &PlayerDailyObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerDailyObjectToEntity(pqo *PlayerDailyObject) (e *questentity.PlayerDailyEntity, err error) {

	e = &questentity.PlayerDailyEntity{
		Id:            pqo.id,
		DailyTag:      int32(pqo.dailyTag),
		SeqId:         pqo.seqId,
		PlayerId:      pqo.player.GetId(),
		Times:         pqo.times,
		LastTime:      pqo.lastTime,
		CrossFiveTime: pqo.crossDayTime,
		UpdateTime:    pqo.updateTime,
		CreateTime:    pqo.createTime,
		DeleteTime:    pqo.deleteTime,
	}
	return
}

func (pqo *PlayerDailyObject) GetPlayerId() int64 {
	return pqo.player.GetId()
}

func (pqo *PlayerDailyObject) GetDBId() int64 {
	return pqo.id
}

func (pqo *PlayerDailyObject) GetDailyTag() questtypes.QuestDailyTag {
	return pqo.dailyTag
}

func (pqo *PlayerDailyObject) GetTimes() int32 {
	return pqo.times
}

func (pqo *PlayerDailyObject) GetLeftTimes() int32 {
	maxTimes := questtemplate.GetQuestTemplateService().GetQuestDailyMaxNum(pqo.dailyTag)
	return maxTimes - pqo.times
}

func (pqo *PlayerDailyObject) GetSeqId() int32 {
	return pqo.seqId
}

func (pqo *PlayerDailyObject) GetCrossFiveTime() int64 {
	return pqo.crossDayTime
}

func (pqo *PlayerDailyObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerDailyObjectToEntity(pqo)
	return
}

func (pqo *PlayerDailyObject) FromEntity(e storage.Entity) error {
	pqe, _ := e.(*questentity.PlayerDailyEntity)

	pqo.id = pqe.Id
	pqo.dailyTag = questtypes.QuestDailyTag(pqe.DailyTag)
	pqo.seqId = pqe.SeqId
	pqo.times = pqe.Times
	pqo.updateTime = pqe.UpdateTime
	pqo.lastTime = pqe.LastTime
	pqo.crossDayTime = pqe.CrossFiveTime
	pqo.createTime = pqe.CreateTime
	pqo.deleteTime = pqe.DeleteTime
	return nil
}

func (pqo *PlayerDailyObject) SetModified() {
	e, err := pqo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "daily"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pqo.player.AddChangedObject(obj)
	return
}

func (pqo *PlayerDailyObject) IsCrossFive() (isCrossFive bool, err error) {
	if !pqo.player.IsFuncOpen(funcopentypes.FuncOpenTypeDailyQuest) {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	lastTime := pqo.crossDayTime

	flag, err := timeutils.IsSameFive(lastTime, now)
	if err != nil {
		return false, err
	}
	if !flag {
		pqo.crossDayTime = now
		pqo.updateTime = now
		pqo.SetModified()
		isCrossFive = true
	}
	return
}

//数据库加载
func (pqdm *PlayerQuestDataManager) loadDaily() (err error) {
	pqdm.playerDailyObjectMap = make(map[questtypes.QuestDailyTag]*PlayerDailyObject)
	//加载日环次数信息
	dailyEntityList, err := dao.GetQuestDao().GetDailyQuestNum(pqdm.p.GetId())
	if err != nil {
		return
	}

	for _, dailyEntity := range dailyEntityList {
		obj := NewPlayerDailyObject(pqdm.p)
		err = obj.FromEntity(dailyEntity)
		if err != nil {
			return
		}
		pqdm.playerDailyObjectMap[obj.GetDailyTag()] = obj
	}

	for dailyTag := questtypes.QuestDailyTagMin; dailyTag <= questtypes.QuestDailyTagMax; dailyTag++ {
		_, ok := pqdm.playerDailyObjectMap[dailyTag]
		if ok {
			continue
		}
		pqdm.initPlayerDailyObject(dailyTag)
	}

	return
}

func (pqdm *PlayerQuestDataManager) getQuestTag() (personDailys []*PlayerQuestObject, allianceDailys []*PlayerQuestObject) {
	personMaxTimes := questtemplate.GetQuestTemplateService().GetQuestDailyMaxNum(questtypes.QuestDailyTagPerson)
	allianceMaxTimes := questtemplate.GetQuestTemplateService().GetQuestDailyMaxNum(questtypes.QuestDailyTagAlliance)
	personDailys = make([]*PlayerQuestObject, 0, personMaxTimes)
	allianceDailys = make([]*PlayerQuestObject, 0, allianceMaxTimes)

	for _, stateQuestMap := range pqdm.stateQuestMap {
		for questId, quest := range stateQuestMap {
			questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
			if questTemplate == nil {
				continue
			}
			questType := questTemplate.GetQuestType()
			switch questType {
			case questtypes.QuestTypeDaily:
				personDailys = append(personDailys, quest)
			case questtypes.QuestTypeDailyAlliance:
				allianceDailys = append(allianceDailys, quest)
			}
		}
	}
	return
}

func (pqdm *PlayerQuestDataManager) resetDaily(dailyTag questtypes.QuestDailyTag, questList []*PlayerQuestObject) (updateQuestList []*PlayerQuestObject) {
	if len(questList) == 0 {
		return
	}
	for _, quest := range questList {
		flag := pqdm.resetDailyInit(quest)
		if flag {
			updateQuestList = append(updateQuestList, quest)
		}
	}
	pqdm.clearDaily(dailyTag)
	return
}

//日环任务
func (pqdm *PlayerQuestDataManager) AfterLoadDailyQuest() (err error) {
	personDailyList, allianceDailyList := pqdm.getQuestTag()
	var questList []*PlayerQuestObject
	for questDailyTag, playerDailyObj := range pqdm.playerDailyObjectMap {
		flag, err := playerDailyObj.IsCrossFive()
		if err != nil {
			return err
		}
		if !flag {
			continue
		}
		switch questDailyTag {
		case questtypes.QuestDailyTagPerson:
			questList = personDailyList
		case questtypes.QuestDailyTagAlliance:
			questList = allianceDailyList
		}
		pqdm.resetDaily(questDailyTag, questList)
	}
	return
}

//第一次初始化
func (pqdm *PlayerQuestDataManager) initPlayerDailyObject(dailyType questtypes.QuestDailyTag) (ptmo *PlayerDailyObject) {
	ptmo, ok := pqdm.playerDailyObjectMap[dailyType]
	if ok {
		return
	}
	if !dailyType.Valid() {
		return
	}
	ptmo = NewPlayerDailyObject(pqdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	ptmo.id = id
	//生成id
	ptmo.times = 0
	ptmo.dailyTag = dailyType
	ptmo.lastTime = 0
	ptmo.crossDayTime = now
	ptmo.createTime = now
	pqdm.playerDailyObjectMap[dailyType] = ptmo
	ptmo.SetModified()
	return
}

func (pqdm *PlayerQuestDataManager) clearDaily(dailyTag questtypes.QuestDailyTag) {
	now := global.GetGame().GetTimeService().Now()
	playerDailyObj, ok := pqdm.playerDailyObjectMap[dailyTag]
	if !ok {
		return
	}
	playerDailyObj.seqId = 0
	playerDailyObj.times = 0
	playerDailyObj.lastTime = now
	playerDailyObj.updateTime = now
	playerDailyObj.SetModified()
}

func (pqdm *PlayerQuestDataManager) resetDailyInit(quest *PlayerQuestObject) bool {
	questId := quest.QuestId
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return false
	}
	questType := questTemplate.GetQuestType()
	if questType != questtypes.QuestTypeDaily && questType != questtypes.QuestTypeDailyAlliance {
		return false
	}
	if quest.QuestState == questtypes.QuestStateInit ||
		quest.QuestState == questtypes.QuestStateCommit ||
		quest.QuestState == questtypes.QuestStateDiscard {
		return false
	}
	//系统领取奖励
	if quest.QuestState == questtypes.QuestStateFinish {
		dailyTag, flag := questType.GetDailyTag()
		if flag {
			gameevent.Emit(questeventtypes.EventTypeQuestDailyReward, pqdm.p, pqdm.playerDailyObjectMap[dailyTag])
		}

	}
	pqdm.restQuestInit(quest)
	return true
}

//日环任务过天跨5点
func (pqdm *PlayerQuestDataManager) resetDailyCrossFive() (err error) {
	for questDailyTag, playerDailyObj := range pqdm.playerDailyObjectMap {
		flag, err := playerDailyObj.IsCrossFive()
		if err != nil {
			return err
		}
		if !flag {
			continue
		}

		seqId := playerDailyObj.GetSeqId()
		dailyTempalte := questtemplate.GetQuestTemplateService().GetQuestDailyTemplateBySeq(questDailyTag, seqId)
		if dailyTempalte == nil {
			pqdm.clearDaily(questDailyTag)
			gameevent.Emit(questeventtypes.EventTypeQuestDailyCrossFive, pqdm.p, questDailyTag)
			continue
		}
		questId := dailyTempalte.GetQuestId()
		quest := pqdm.GetQuestById(questId)
		//一键完成了
		if quest == nil {
			pqdm.clearDaily(questDailyTag)
			gameevent.Emit(questeventtypes.EventTypeQuestDailyCrossFive, pqdm.p, questDailyTag)
			continue
		}
		now := global.GetGame().GetTimeService().Now()
		switch quest.QuestState {
		case questtypes.QuestStateFinish: //系统领取奖励
			gameevent.Emit(questeventtypes.EventTypeQuestDailyReward, pqdm.p, playerDailyObj)
			pqdm.removeQuestByIdAndState(quest.QuestState, questId)
			quest.QuestState = questtypes.QuestStateCommit
			quest.UpdateTime = now
			quest.SetModified()
			break
		case questtypes.QuestStateAccept: //放弃任务
			pqdm.removeQuestByIdAndState(quest.QuestState, questId)
			quest.QuestState = questtypes.QuestStateDiscard
			quest.UpdateTime = now
			quest.SetModified()
			break
		case questtypes.QuestStateCommit, //放弃和提交
			questtypes.QuestStateDiscard,
			questtypes.QuestStateInit:
			goto DailyCrossFive
		}
		pqdm.addQuest(quest)
		gameevent.Emit(questeventtypes.EventTypeQuestDailyUpdate, pqdm.p, quest)

	DailyCrossFive:
		pqdm.clearDaily(questDailyTag)
		gameevent.Emit(questeventtypes.EventTypeQuestDailyCrossFive, pqdm.p, questDailyTag)
	}
	return
}

//获取日环次数
func (pqdm *PlayerQuestDataManager) GetDailyObj(dailyTag questtypes.QuestDailyTag) *PlayerDailyObject {
	playerDailyObj, ok := pqdm.playerDailyObjectMap[dailyTag]
	if !ok {
		return nil
	}
	return playerDailyObj
}

//随机第一个日环任务
func (pqdm *PlayerQuestDataManager) AcceptDailyQuest(dailyTag questtypes.QuestDailyTag) (questObj *PlayerQuestObject, flag bool) {
	questType := dailyTag.GetQuestType()
	if !questType.Valid() {
		return
	}
	if dailyTag == questtypes.QuestDailyTagAlliance && pqdm.p.GetAllianceId() == 0 {
		return
	}
	funcOpen := dailyTag.GetFuncOpen()
	playerDailyObj := pqdm.playerDailyObjectMap[dailyTag]
	if playerDailyObj == nil {
		return
	}
	if !pqdm.p.IsFuncOpen(funcOpen) {
		return
	}
	for _, stateQuestMap := range pqdm.stateQuestMap {
		for questId, quest := range stateQuestMap {
			switch quest.QuestState {
			case questtypes.QuestStateInit,
				questtypes.QuestStateActive,
				questtypes.QuestStateDiscard,
				questtypes.QuestStateCommit:
				continue
			}
			questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
			if questTemplate == nil {
				continue
			}
			if questTemplate.GetQuestType() != questType {
				continue
			}
			return
		}
	}

	maxTimes := questtemplate.GetQuestTemplateService().GetQuestDailyMaxNum(dailyTag)
	if playerDailyObj.times >= maxTimes {
		return
	}

	dailyTemplate, flag := questtemplate.GetQuestTemplateService().GetQuestDailyTemplate(pqdm.p, dailyTag, questtypes.QuestDailyTypeMin)
	if !flag {
		return
	}
	questObj = pqdm.AddDailyQuest(dailyTag, dailyTemplate)
	if questObj != nil {
		flag = true
	}
	return
}

//设置日环任务信息
func (pqdm *PlayerQuestDataManager) setDailyTimes(dailyTag questtypes.QuestDailyTag, dailyTemplate gametemplate.DailyTagTemplate) {
	if dailyTemplate == nil {
		return
	}
	if !dailyTag.Valid() {
		return
	}
	seqId := int32(dailyTemplate.TemplateId())
	now := global.GetGame().GetTimeService().Now()

	playerDailyObj, ok := pqdm.playerDailyObjectMap[dailyTag]
	if !ok {
		playerDailyObj = pqdm.initPlayerDailyObject(dailyTag)
	}
	playerDailyObj.seqId = int32(seqId)
	playerDailyObj.times += 1
	playerDailyObj.lastTime = now
	playerDailyObj.updateTime = now
	playerDailyObj.SetModified()
}

//添加日环任务
func (pqdm *PlayerQuestDataManager) AddDailyQuest(dailyTag questtypes.QuestDailyTag, dailyTemplate gametemplate.DailyTagTemplate) (quest *PlayerQuestObject) {
	if dailyTemplate == nil {
		return
	}
	questId := dailyTemplate.GetQuestId()
	now := global.GetGame().GetTimeService().Now()
	questObj := pqdm.GetQuestById(questId)
	if questObj != nil {
		if questObj.QuestState != questtypes.QuestStateCommit &&
			questObj.QuestState != questtypes.QuestStateDiscard &&
			questObj.QuestState != questtypes.QuestStateInit {
			return
		}
		pqdm.removeQuestByIdAndState(questObj.QuestState, questId)
		questObj.QuestState = questtypes.QuestStateAccept
		questObj.QuestDataMap = make(map[int32]int32)
		questObj.CollectItemDataMap = make(map[int32]int32)
		questObj.UpdateTime = now
		questObj.SetModified()
	} else {
		questObj = createQuest(pqdm.p, questId)
		questObj.QuestState = questtypes.QuestStateAccept
		questObj.SetModified()
	}
	pqdm.addQuest(questObj)
	pqdm.setDailyTimes(dailyTag, dailyTemplate)
	return questObj
}

//获取下一个日环任务
func (pqdm *PlayerQuestDataManager) GetNextDailyQuest(dailyTag questtypes.QuestDailyTag) (quest *PlayerQuestObject) {
	if !dailyTag.Valid() {
		return
	}
	obj, ok := pqdm.playerDailyObjectMap[dailyTag]
	if !ok {
		return
	}

	maxTimes := questtemplate.GetQuestTemplateService().GetQuestDailyMaxNum(dailyTag)
	if obj.times >= maxTimes {
		return
	}

	nextDailyTimes := questtypes.QuestDailyType(obj.times + 1)
	dailyTempalte, flag := questtemplate.GetQuestTemplateService().GetQuestDailyTemplate(pqdm.p, dailyTag, nextDailyTimes)
	if !flag {
		return
	}
	quest = pqdm.AddDailyQuest(dailyTag, dailyTempalte)
	return
}

//日环任务直接完成
func (pqdm *PlayerQuestDataManager) QuestDailyImmediateFinish(questId int32) (qu *PlayerQuestObject, flag bool) {
	now := global.GetGame().GetTimeService().Now()
	quest := pqdm.GetQuestById(questId)
	if quest == nil {
		return
	}
	if quest.QuestState != questtypes.QuestStateAccept &&
		quest.QuestState != questtypes.QuestStateFinish {
		return
	}

	q := pqdm.GetQuestByIdAndState(quest.QuestState, questId)
	//移除
	questType := questtemplate.GetQuestTemplateService().GetQuestTypeById(questId)
	if questType == questtypes.QuestTypeTuMo {
		pqdm.removeTuMoQuestByIdAndState(quest.QuestState, questId)
	} else {
		pqdm.removeQuestByIdAndState(quest.QuestState, questId)
	}

	q.QuestState = questtypes.QuestStateCommit
	q.UpdateTime = now
	q.SetModified()

	pqdm.addQuest(q)
	qu = q
	flag = true
	return
}

func (pqdm *PlayerQuestDataManager) QuestDailyFinishAll(dailyTag questtypes.QuestDailyTag, seqId int32) {
	now := global.GetGame().GetTimeService().Now()
	playerDailyObj, ok := pqdm.playerDailyObjectMap[dailyTag]
	if !ok {
		return
	}

	playerDailyObj.seqId = seqId
	playerDailyObj.times = questtemplate.GetQuestTemplateService().GetQuestDailyMaxNum(dailyTag)
	playerDailyObj.updateTime = now
	playerDailyObj.lastTime = now
	playerDailyObj.SetModified()
}

//功能开启日环任务
func (pqdm *PlayerQuestDataManager) CheckDailyQuest(funcOpenType funcopentypes.FuncOpenType) (quest *PlayerQuestObject) {

	if funcOpenType != funcopentypes.FuncOpenTypeDailyQuest &&
		funcOpenType != funcopentypes.FuncOpenTypeAllianceDaily {
		return
	}
	if !pqdm.p.IsFuncOpen(funcOpenType) {
		return
	}

	dailyTag := questtypes.QuestDailyTagAlliance
	switch funcOpenType {
	case funcopentypes.FuncOpenTypeDailyQuest:
		dailyTag = questtypes.QuestDailyTagPerson
	case funcopentypes.FuncOpenTypeAllianceBoss:
		dailyTag = questtypes.QuestDailyTagAlliance
	}
	quest, flag := pqdm.AcceptDailyQuest(dailyTag)
	if !flag {
		return nil
	}
	return quest
}

//判断是否日环任务完成
func (pqdm *PlayerQuestDataManager) IfDailyQuestFinish(dailyTag questtypes.QuestDailyTag) bool {
	playerDailyObj, ok := pqdm.playerDailyObjectMap[dailyTag]
	if !ok {
		return false
	}

	maxTimes := questtemplate.GetQuestTemplateService().GetQuestDailyMaxNum(dailyTag)
	if playerDailyObj.times >= maxTimes {
		return true
	}
	return false
}

func (pqdm *PlayerQuestDataManager) GetCurrentDailyQuest() *PlayerQuestObject {
	acceptQuestMap := pqdm.GetQuestMap(questtypes.QuestStateAccept)
	for _, q := range acceptQuestMap {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(q.QuestId)
		if questTemplate == nil {
			continue
		}
		if questTemplate.GetQuestType() == questtypes.QuestTypeDaily {
			return q
		}
	}
	finishQuestMap := pqdm.GetQuestMap(questtypes.QuestStateFinish)
	for _, q := range finishQuestMap {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(q.QuestId)
		if questTemplate == nil {
			continue
		}
		if questTemplate.GetQuestType() == questtypes.QuestTypeDaily {
			return q
		}
	}
	return nil
}
