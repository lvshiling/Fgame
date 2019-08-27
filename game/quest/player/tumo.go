package player

import (
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"

	"fgame/fgame/game/quest/dao"
	questentity "fgame/fgame/game/quest/entity"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"

	"github.com/pkg/errors"
)

//屠魔次数对象
type PlayerTuMoObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Num        int32
	ExtraNum   int32
	UsedNum    int32
	UsedBuyNum int32
	BuyNum     int32
	LastTime   int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerTuMoObject(pl player.Player) *PlayerTuMoObject {
	ptmo := &PlayerTuMoObject{
		player: pl,
	}
	return ptmo
}

func (ptmo *PlayerTuMoObject) GetPlayerId() int64 {
	return ptmo.player.GetId()
}

func (ptmo *PlayerTuMoObject) GetDBId() int64 {
	return ptmo.Id
}

func (ptmo *PlayerTuMoObject) ToEntity() (e storage.Entity, err error) {
	e = &questentity.PlayerTuMoEntity{
		Id:         ptmo.Id,
		PlayerId:   ptmo.PlayerId,
		Num:        ptmo.Num,
		ExtraNum:   ptmo.ExtraNum,
		UsedNum:    ptmo.UsedNum,
		UsedBuyNum: ptmo.UsedBuyNum,
		BuyNum:     ptmo.BuyNum,
		LastTime:   ptmo.LastTime,
		UpdateTime: ptmo.UpdateTime,
		CreateTime: ptmo.CreateTime,
		DeleteTime: ptmo.DeleteTime,
	}
	return e, nil
}

func (ptmo *PlayerTuMoObject) FromEntity(e storage.Entity) error {
	ptme, _ := e.(*questentity.PlayerTuMoEntity)
	ptmo.Id = ptme.Id
	ptmo.PlayerId = ptme.PlayerId
	ptmo.Num = ptme.Num
	ptmo.ExtraNum = ptme.ExtraNum
	ptmo.UsedNum = ptme.UsedNum
	ptmo.UsedBuyNum = ptme.UsedBuyNum
	ptmo.BuyNum = ptme.BuyNum
	ptmo.LastTime = ptme.LastTime
	ptmo.UpdateTime = ptme.UpdateTime
	ptmo.CreateTime = ptme.CreateTime
	ptmo.DeleteTime = ptme.DeleteTime
	return nil
}

func (ptmo *PlayerTuMoObject) SetModified() {
	e, err := ptmo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Quest"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	ptmo.player.AddChangedObject(obj)
	return
}

func (pqdm *PlayerQuestDataManager) loadTuMo() (err error) {
	//加载屠魔次数信息
	tuMoEntity, err := dao.GetQuestDao().GetTuMoQuestNum(pqdm.p.GetId())
	if err != nil {
		return
	}
	if tuMoEntity == nil {
		pqdm.initPlayerTuMoObject()
	} else {
		pqdm.playerTuMoObject = NewPlayerTuMoObject(pqdm.p)
		pqdm.playerTuMoObject.FromEntity(tuMoEntity)
	}
	return
}

//第一次初始化
func (pqdm *PlayerQuestDataManager) initPlayerTuMoObject() {
	ptmo := NewPlayerTuMoObject(pqdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	ptmo.Id = id
	//生成id
	ptmo.PlayerId = pqdm.p.GetId()
	ptmo.Num = int32(0)
	ptmo.ExtraNum = int32(0)
	ptmo.UsedNum = int32(0)
	ptmo.UsedBuyNum = int32(0)
	ptmo.BuyNum = int32(0)
	ptmo.LastTime = int64(0)
	ptmo.CreateTime = now
	pqdm.playerTuMoObject = ptmo
	ptmo.SetModified()
}

func (pqdm *PlayerQuestDataManager) afterLoadTuMoLine() (err error) {
	//刷新屠魔次数
	err = pqdm.refreshTuMoNum()
	if err != nil {
		return
	}
	return
}

//获取屠魔任务quest
func (pqdm *PlayerQuestDataManager) GetTuMoQuestById(questId int32) (quest *PlayerQuestObject) {
	for _, questMap := range pqdm.tuMoQuestMap {
		quest, exist := questMap[questId]
		if exist {
			return quest
		}
	}
	return
}

//获取屠魔任务剩余次数
func (pqdm *PlayerQuestDataManager) GetTuMoLeftDefaultNum() (leftNum int32) {
	leftNum = 0
	maxNum := questtemplate.GetQuestTemplateService().GetQuestTuMoInitialNum()
	curNum, _, _ := pqdm.GetTuMoNum()
	leftNum = maxNum - curNum
	return
}

//获取屠魔次数和购买次数
func (pqdm *PlayerQuestDataManager) GetTuMoNum() (num int32, buyNum int32, extraNum int32) {
	//刷新屠魔次数
	pqdm.refreshTuMoNum()
	num = pqdm.playerTuMoObject.Num
	buyNum = pqdm.playerTuMoObject.BuyNum
	extraNum = pqdm.playerTuMoObject.ExtraNum
	return
}

//获取玩家屠魔购买次数
func (pqdm *PlayerQuestDataManager) GetBuyNum() int32 {
	return pqdm.playerTuMoObject.BuyNum
}

//获取额外屠魔总次数
func (pqdm *PlayerQuestDataManager) GetExtraNum() int32 {
	return pqdm.playerTuMoObject.ExtraNum
}

//获取玩家剩余的额外屠魔次数
func (pqdm *PlayerQuestDataManager) getLeftExtraNum() int32 {
	leftNum := int32(0)
	extraNum := pqdm.playerTuMoObject.ExtraNum
	usedBuyNum := pqdm.playerTuMoObject.UsedBuyNum
	if extraNum > usedBuyNum {
		leftNum = extraNum - usedBuyNum
	}
	return leftNum
}

//获取屠魔任务列表个数
func (pqdm *PlayerQuestDataManager) NumOfQuestTuMoList() int32 {
	num := int32(0)
	for _, questMap := range pqdm.tuMoQuestMap {
		num += int32(len(questMap))
	}
	return num
}

//获取屠魔任务id
func (pqdm *PlayerQuestDataManager) GetTuMoQuestIdByToken(token questtypes.QuestLevelType) (int32, bool) {
	if !token.Valid() {
		return 0, false
	}
	level := pqdm.p.GetLevel()
	usedList, questTb := pqdm.getUsedTuMoListByToken(token, level)
	return questtemplate.GetQuestTemplateService().GetQuestIdForTuMo(token, usedList, questTb, level)
}

//增加购买次数
func (pqdm *PlayerQuestDataManager) AddNumByBuy() error {
	flag, err := pqdm.IfReachBuyLimit()
	if err != nil {
		return err
	}
	if flag {
		return nil
	}
	now := global.GetGame().GetTimeService().Now()
	pqdm.playerTuMoObject.BuyNum += 1
	pqdm.playerTuMoObject.ExtraNum += 1
	pqdm.playerTuMoObject.LastTime = now
	pqdm.playerTuMoObject.UpdateTime = now
	pqdm.playerTuMoObject.SetModified()
	return nil
}

//任务列表是已满
func (pqdm *PlayerQuestDataManager) IfTuMoListReachLimit() bool {
	curNum := pqdm.NumOfQuestTuMoList()
	maxNum := questtemplate.GetQuestTemplateService().GetTuMoTaskBarDefaultNum()
	//vip 限制先改成等级限制
	fourthLimit := questtemplate.GetQuestTemplateService().GetTuMoTaskBarOpenFourthVipLevel()
	thirdLimit := questtemplate.GetQuestTemplateService().GetTuMoTaskBarOpenThirdPlayerLevel()
	level := pqdm.p.GetLevel()
	if level >= thirdLimit {
		maxNum += 1
	}
	//vip 限制先改成等级限制
	vip := pqdm.p.GetLevel()
	if vip >= fourthLimit {
		maxNum += 1
	}
	if curNum >= maxNum {
		return true
	}
	return false
}

//屠魔次数是已达上限
func (pqdm *PlayerQuestDataManager) IfTuMoNumReachLimit() bool {
	maxNum := questtemplate.GetQuestTemplateService().GetQuestTuMoInitialNum()
	curNum, _, extraNum := pqdm.GetTuMoNum()
	totalLeft := maxNum + extraNum - curNum
	if totalLeft > 0 {
		return false
	}
	return true
}

func (pqdm *PlayerQuestDataManager) maxBuyNum() (maxBuyNum int32) {
	maxBuyLimit := questtemplate.GetQuestTemplateService().GetQuestTuMoVipAddBuyNum()
	maxBuyNum = maxBuyLimit
	//TODO 获取VIP
	// curVip := int32(0)
	// minVipLimit := questtemplate.GetQuestTemplateService().GetQuestTuMoBuyNumVipLimit()
	// if curVip < minVipLimit {
	// 	maxBuyNum = 0
	// }
	return maxBuyNum
}

//是否到达最大购买次数
func (pqdm *PlayerQuestDataManager) IfReachBuyLimit() (bool, error) {
	maxBuyNum := pqdm.maxBuyNum()
	//刷新屠魔次数
	err := pqdm.refreshTuMoNum()
	if err != nil {
		return false, err
	}

	curBuyNum := pqdm.GetBuyNum()
	if curBuyNum >= maxBuyNum {
		return true, nil
	}
	return false, nil

}

//刷新屠魔次数
func (pqdm *PlayerQuestDataManager) refreshTuMoNum() error {
	//是否跨天
	now := global.GetGame().GetTimeService().Now()
	lastTime := pqdm.playerTuMoObject.LastTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameFive(lastTime, now)
		if err != nil {
			return err
		}
		if !flag {
			pqdm.playerTuMoObject.Num = 0
			pqdm.playerTuMoObject.BuyNum = 0
			pqdm.playerTuMoObject.UsedNum = 0
			pqdm.playerTuMoObject.ExtraNum -= pqdm.playerTuMoObject.UsedBuyNum
			pqdm.playerTuMoObject.UsedBuyNum = 0
			pqdm.playerTuMoObject.LastTime = 0
			pqdm.playerTuMoObject.UpdateTime = now
			pqdm.playerTuMoObject.SetModified()
		}
	}
	return nil
}

//获取正在进行的屠魔任务 (questTb仅橙色有效)
func (pqdm *PlayerQuestDataManager) getUsedTuMoListByToken(token questtypes.QuestLevelType, level int32) (usedList []int32, questTb int32) {
	questTb = 0
	questMap, exist := pqdm.tuMoQuestMap[token]
	if !exist {
		return nil, 0
	}
	if len(questMap) <= 0 {
		return nil, 0
	}
	for _, obj := range questMap {
		usedList = append(usedList, obj.QuestId)
	}

	//获取橙色questTb 可能跨级 需要等级判断
	for _, questId := range usedList {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
		if questTemplate.GetQuestLevel() != questtypes.QuestLevelTypeTuMoOrange {
			continue
		}
		minLevel := questTemplate.MinLevel
		maxLevel := questTemplate.MaxLevel
		if level >= minLevel && level <= maxLevel {
			questTb = questTemplate.QuestTb
			break
		}
	}
	return
}

//添加屠魔任务
func (pqdm *PlayerQuestDataManager) addTuMoQuest(questLevel questtypes.QuestLevelType, quest *PlayerQuestObject) {
	questMap, exist := pqdm.tuMoQuestMap[questLevel]
	if !exist {
		questMap = make(map[int32]*PlayerQuestObject)
		pqdm.tuMoQuestMap[questLevel] = questMap
	}
	questMap[quest.QuestId] = quest
}

//添加屠魔次数
func (pqdm *PlayerQuestDataManager) addTuMoNum() {
	//屠魔次数+1
	now := global.GetGame().GetTimeService().Now()
	pqdm.playerTuMoObject.Num += 1
	//优先使用购买次数
	leftExtraNum := pqdm.getLeftExtraNum()
	if leftExtraNum > 0 {
		pqdm.playerTuMoObject.UsedBuyNum += 1
	} else {
		pqdm.playerTuMoObject.UsedNum += 1
	}
	pqdm.playerTuMoObject.LastTime = now
	pqdm.playerTuMoObject.UpdateTime = now
	pqdm.playerTuMoObject.SetModified()
}

//接受屠魔任务
func (pqdm *PlayerQuestDataManager) AcceptTumoQuest(questId int32) (flag bool) {
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	typ := questTemplate.GetQuestType()
	if typ != questtypes.QuestTypeTuMo {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	questLevel := questTemplate.GetQuestLevel()
	questObj := pqdm.GetQuestById(questId)
	if questObj == nil {
		questObj = createQuest(pqdm.p, questId)
	} else {
		if questObj.QuestState != questtypes.QuestStateCommit &&
			questObj.QuestState != questtypes.QuestStateDiscard {
			return
		}
		questObj.QuestDataMap = make(map[int32]int32)
		questObj.CollectItemDataMap = make(map[int32]int32)
		//移除
		pqdm.removeQuestByIdAndState(questObj.QuestState, questId)
	}
	questObj.QuestState = questtypes.QuestStateAccept
	questObj.UpdateTime = now
	questObj.SetModified()

	pqdm.addTuMoQuest(questLevel, questObj)
	pqdm.addQuest(questObj)
	pqdm.addTuMoNum()
	flag = true
	//TODO 发送事件
	gameevent.Emit(questeventtypes.EventTypeQuestAccept, pqdm.p, questId)
	return
}

//放弃屠魔任务
func (pqdm *PlayerQuestDataManager) DiscardTuMoQuest(questId int32) {
	questObj := pqdm.GetQuestByIdAndState(questtypes.QuestStateAccept, questId)
	tuMoObj := pqdm.GetTuMoQuestById(questId)
	if questObj == nil || tuMoObj == nil {
		return
	}

	//移除
	pqdm.removeTuMoQuestByIdAndState(questtypes.QuestStateAccept, questId)
	questObj.QuestState = questtypes.QuestStateDiscard
	now := global.GetGame().GetTimeService().Now()
	questObj.UpdateTime = now
	questObj.SetModified()
	pqdm.addQuest(questObj)
	return
}

//移除屠魔任务
func (pqdm *PlayerQuestDataManager) removeTuMoQuestByIdAndState(questState questtypes.QuestState, questId int32) {
	//移除任务对象
	pqdm.removeQuestByIdAndState(questState, questId)

	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	questLevel := questTemplate.GetQuestLevel()
	//移除屠魔任务
	questMap, exist := pqdm.tuMoQuestMap[questLevel]
	if !exist {
		return
	}
	delete(questMap, questId)
}

//获取剩余的屠魔次数
func (pqdm *PlayerQuestDataManager) GetTuMoAcceptQuestAndNeedNum() (questIdList []int32, needNum int32, accepectNum int32) {
	maxNum := questtemplate.GetQuestTemplateService().GetQuestTuMoInitialNum()
	curNum, _, extraNum := pqdm.GetTuMoNum()

	accepectNum = 0
	for _, questMap := range pqdm.tuMoQuestMap {
		for _, questObj := range questMap {
			if questObj.QuestState == questtypes.QuestStateAccept {
				accepectNum++
				questIdList = append(questIdList, questObj.QuestId)
			}
		}
	}

	needNum = maxNum + extraNum - curNum
	return
}

//一键完成使用屠魔次数
func (pqdm *PlayerQuestDataManager) UseTuMoNum(addNum int32) (num int32) {
	if addNum <= 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pqdm.playerTuMoObject.Num += addNum
	for i := int32(0); i < num; i++ {
		leftExtraNum := pqdm.getLeftExtraNum()
		if leftExtraNum > 0 {
			pqdm.playerTuMoObject.UsedBuyNum += 1
		} else {
			pqdm.playerTuMoObject.UsedNum += 1
		}
	}

	pqdm.playerTuMoObject.LastTime = now
	pqdm.playerTuMoObject.UpdateTime = now
	num = pqdm.playerTuMoObject.Num
	pqdm.playerTuMoObject.SetModified()

	data := questeventtypes.CreateQuestFinishAllEventData(questtypes.QuestTypeTuMo, addNum)
	gameevent.Emit(questeventtypes.EventTypeQuestFinishAll, pqdm.p, data)
	return
}
