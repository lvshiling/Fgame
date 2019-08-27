package marry

import (
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/marry/dao"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	hunche "fgame/fgame/game/marry/npc/hunche"
	marryscene "fgame/fgame/game/marry/scene"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenescene "fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"sort"
	"sync"
)

//婚期接口处理
type MarryService interface {
	Heartbeat()
	//创建结婚数据
	CreateMarrySceneData()
	//获取结婚场景
	GetScene() scenescene.Scene
	//婚车巡游结束
	WeddingCarEnd()
	//设置婚车对象
	WeddingCarNPC(*hunche.HunCheNPC)
	//获取婚车对象
	GetHunCheNpc() *hunche.HunCheNPC
	//获取婚宴场景
	GetMarrySceneData() *marryscene.MarrySceneStatusData
	//获取喜帖列表
	GetWeddingCardList() []*MarryWedCardObject
	//获取婚期预定列表
	GetMarryWeddingList() []*MarryWedObject
	//获取婚礼举办场次
	GetWeddingPeriod(playerId int64) (period int32)
	//获取婚烟信息
	GetMarry(playerId int64) *MarryObject
	//获取求婚婚戒
	GetMarryProposalRing(playerId int64) *MarryRingObject
	//获取玩家的婚礼戒指信息
	GetMarryRing(playerId int64) marrytypes.MarryRingType
	//添加亲密度
	AddPoint(playerId int64, addPoint int32)
	//删除求婚婚戒
	RemoveMarryRing(playerId int64)
	//同步婚戒等级
	MarryRingLevel(playerId int64, ringLevel int32)
	//婚戒替换
	MarryReplace(playerId int64, ringType marrytypes.MarryRingType)
	//求婚
	MarryProposalRing(playerId int64, peerId int64, peerName string, ringType marrytypes.MarryRingType)
	//喜帖是否存在
	WeddingCardIsExist(cardId int64) bool
	//赠送贺礼
	IfCanGiveWedGift(period int32) bool
	//离婚
	Divorce(pl player.Player, spouseId int64, typ marrytypes.MarryDivorceType) error
	//离婚决策
	DivorceDeal(pl player.Player, spouseId int64, result marrytypes.MarryResultType)
	//被求婚者决策
	ProposalDeal(pl player.Player, dealId int64, ringLevel int32, ringType marrytypes.MarryRingType, point int32, result marrytypes.MarryResultType) error
	//预定婚期
	MarryPreWedding(pl player.Player, period int32, marryGrade *marrytypes.MarryGrade, spouseId int64, reserveTime int64) error
	//婚期预定配偶决策
	MarryPreWedDeal(pl player.Player, result bool, isFirst bool) (err error)
	//玩家修改名字
	PlayerNameChanged(pl player.Player)
	//获取配偶名字
	GetSpouseName(playerId int64) (spouseName string)
	//判断是否请求过婚期预定
	MarryPreWedIsExist(playerId int64) (flag bool)
	//gm 使用 清空婚期
	GmClearMarryWed()
	//同步表白等级
	SyncMarryDevelopLevel(playerId int64, developLevel int32)
	GetSpouseDevelopLevel(playerId int64) int32

	GetMarrySpouseId(playerId int64) (spouseId int64)

	//同步定情信物
	SyncMarryDingQing(playerId int64, dqMap map[int32]map[int32]int32)
	GetMarryDingQing(playerId int64) map[int32]map[int32]int32
	ExistsSpouseDingQing(playerId int64, suiId int32, posId int32) bool
	//获取伴侣id
	GetSpouseId(playerId int64) int64
}

type marryService struct {
	//读写锁
	rwm sync.RWMutex
	//结婚场景状态
	sceneStatusData *marryscene.MarrySceneStatusData
	//婚期预定列表
	marryWedList []*MarryWedObject
	//喜帖map
	marryWedCardMap map[int64]*MarryWedCardObject
	//婚烟map
	marryMap map[int64]*MarryObject

	//结婚场景
	marryScene scenescene.Scene
	//婚车npc
	huncheNpc *hunche.HunCheNPC
	//求婚戒指
	marryRingMap map[int64]*MarryRingObject
	//预定婚期map
	marryPreWedMap map[int64]*MarryPreWedObject
}

//初始化
func (ms *marryService) init() (err error) {

	//婚期map
	if err = ms.initMarryWedList(); err != nil {
		return
	}

	//喜帖map
	if err = ms.initMarryCardList(); err != nil {
		return
	}

	//婚烟map
	if err = ms.initMarryList(); err != nil {
		return
	}

	//婚戒map
	if err = ms.initMarryRingList(); err != nil {
		return
	}

	//预定婚期
	if err = ms.initMarryPreWedList(); err != nil {
		return
	}

	//合服
	isMerge := merge.GetMergeService().IsMerge()
	if isMerge {
		ms.mergeServerMarryWed()
		ms.mergeServerMarryWedCard()
		ms.mergeServerMarry()
	} else { //关服补偿
		ms.closeServerCompensation()
	}

	return
}

func (ms *marryService) initMarryWedList() (err error) {
	ms.marryWedList = make([]*MarryWedObject, 0, 12)
	marryWedList, err := dao.GetMarryDao().GetMarryWedList()
	if err != nil {
		return
	}
	for _, marryWed := range marryWedList {
		mwo := NewMarryWedObject()
		mwo.FromEntity(marryWed)
		ms.marryWedList = append(ms.marryWedList, mwo)
	}
	return nil
}

func (ms *marryService) initMarryCardList() (err error) {
	ms.marryWedCardMap = make(map[int64]*MarryWedCardObject)
	now := global.GetGame().GetTimeService().Now()
	marryWedCardList, err := dao.GetMarryDao().GetMarryWedCardList(now)
	if err != nil {
		return
	}
	for _, marryWedCard := range marryWedCardList {
		mwco := NewMarryWedCardObject()
		mwco.FromEntity(marryWedCard)
		ms.marryWedCardMap[mwco.Id] = mwco
	}
	return
}

func (ms *marryService) initMarryList() (err error) {
	ms.marryMap = make(map[int64]*MarryObject)
	//婚烟map
	marryList, err := dao.GetMarryDao().GetMarryList()
	if err != nil {
		return
	}
	for _, marryObj := range marryList {
		mo := NewMarryObject()
		mo.FromEntity(marryObj)
		if mo.Status != marrytypes.MarryStatusTypeUnmarried && mo.Status != marrytypes.MarryStatusTypeDivorce { //未结婚
			ms.marryMap[mo.PlayerId] = mo
			ms.marryMap[mo.SpouseId] = mo
		}
	}
	return
}

func (ms *marryService) initMarryRingList() (err error) {
	ms.marryRingMap = make(map[int64]*MarryRingObject)
	//婚戒map
	marryRingList, err := dao.GetMarryDao().GetMarryRingList()
	if err != nil {
		return
	}
	for _, marryRingObj := range marryRingList {
		mro := NewMarryRingObject()
		mro.FromEntity(marryRingObj)
		ms.marryRingMap[mro.PlayerId] = mro
	}
	return
}

func (ms *marryService) initMarryPreWedList() (err error) {
	ms.marryPreWedMap = make(map[int64]*MarryPreWedObject)

	//预定婚期
	marryPreWedList, err := dao.GetMarryDao().GetMarryPreWedList()
	if err != nil {
		return
	}
	for _, marryPreWedObj := range marryPreWedList {
		mpwo := NewMarryPreWedObject()
		mpwo.FromEntity(marryPreWedObj)
		ms.marryPreWedMap[mpwo.PlayerId] = mpwo
	}
	return
}

func (ms *marryService) mergeServerMarryWed() {
	now := global.GetGame().GetTimeService().Now()
	for _, marryWed := range ms.marryWedList {
		period := marryWed.Period
		status := marryWed.Status
		if status != marrytypes.MarryWedStatusTypeNoStart && period != -1 {
			continue
		}
		marryWed.DeleteTime = now
		marryWed.SetModified()
		gameevent.Emit(marryeventtypes.EventTypeMarryMergeServer, marryWed, nil)
	}
	ms.marryWedList = make([]*MarryWedObject, 0, 12)
}

//TODO:zrc 修改半小时设定
func (ms *marryService) closeServerCompensation() {
	now := global.GetGame().GetTimeService().Now()
	for _, marryWed := range ms.marryWedList {
		if marryWed.Status != marrytypes.MarryWedStatusTypeNoStart {
			continue
		}
		//婚礼半个小时后的不补偿
		leftTime := marryWed.HTime - now
		if leftTime >= int64(30*common.MINUTE) {
			continue
		}
		playerId := marryWed.PlayerId
		marryWed.Status = marrytypes.MarryWedStatusTypeCancle
		marryWed.UpdateTime = now

		marryObj, exist := ms.marryMap[playerId]
		if exist {
			marryObj.Status = marrytypes.MarryStatusTypeProposal
			marryObj.UpdateTime = now
		}
		gameevent.Emit(marryeventtypes.EventTypeMarryCloseServer, marryWed, nil)
		//半小时还能预定的保持预定
		if leftTime > 0 && leftTime < int64(30*common.MINUTE) {
			marryWed.Period = -1
			marryWed.Grade = 0
			marryWed.HunCheGrade = 0
			marryWed.SugarGrade = 0
		}
		marryWed.SetModified()
	}
}

func (ms *marryService) mergeServerMarryWedCard() {
	now := global.GetGame().GetTimeService().Now()
	for _, marryWedCard := range ms.marryWedCardMap {
		marryWedCard.DeleteTime = now
		marryWedCard.SetModified()
		delete(ms.marryWedCardMap, marryWedCard.Id)
	}
}

func (ms *marryService) mergeServerMarry() {
	now := global.GetGame().GetTimeService().Now()
	for _, marryObj := range ms.marryMap {
		if marryObj.Status != marrytypes.MarryStatusTypeEngagement {
			continue
		}
		marryObj.Status = marrytypes.MarryStatusTypeProposal
		marryObj.UpdateTime = now
		marryObj.SetModified()
	}
}

func (ms *marryService) initMarryPreWedObj(pl player.Player, period int32, marryGrade *marrytypes.MarryGrade, spouseId int64, reserveTime int64) {
	now := global.GetGame().GetTimeService().Now()
	mpwo := NewMarryPreWedObject()
	id, _ := idutil.GetId()
	mpwo.Id = id
	mpwo.ServerId = global.GetGame().GetServerIndex()
	mpwo.Period = period
	mpwo.PlayerId = pl.GetId()
	mpwo.PlayerName = pl.GetName()
	mpwo.PeerId = spouseId
	mpwo.Grade = marryGrade.Grade
	mpwo.HunCheGrade = marryGrade.HunCheGrade
	mpwo.SugarGrade = marryGrade.SugarGrade
	mpwo.Status = marrytypes.MarryPreWedStatusTypeOngoing
	mpwo.PreWedTime = now
	mpwo.HoldTime = reserveTime
	mpwo.SetModified()
	ms.marryPreWedMap[pl.GetId()] = mpwo
	return
}

//同步表白等级
func (ms *marryService) SyncMarryDevelopLevel(playerId int64, developLevel int32) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	marryObj, exist := ms.marryMap[playerId]
	if !exist {
		return
	}

	if marryObj.PlayerId == playerId {
		marryObj.DevelopLevel = developLevel
	} else {
		marryObj.SpouseDevelopLevel = developLevel
	}
	now := global.GetGame().GetTimeService().Now()
	marryObj.UpdateTime = now
	marryObj.SetModified()

	eventData := marryeventtypes.CreateMarryDevelopLevelChangedEventData(playerId, developLevel)
	gameevent.Emit(marryeventtypes.EventTypeMarryDevelopLevelChanged, marryObj, eventData)
}

//同步婚戒等级
func (ms *marryService) MarryRingLevel(playerId int64, ringLevel int32) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	marryObj, exist := ms.marryMap[playerId]
	if !exist {
		return
	}

	if marryObj.PlayerId == playerId {
		marryObj.PlayerRingLevel = ringLevel
	} else {
		marryObj.SpouseRingLevel = ringLevel
	}
	now := global.GetGame().GetTimeService().Now()
	marryObj.UpdateTime = now
	marryObj.SetModified()
}

//同步配偶名字
func (ms *marryService) GetSpouseName(playerId int64) (spouseName string) {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()

	marryObj, exist := ms.marryMap[playerId]
	if !exist {
		return
	}
	if marryObj.PlayerId == playerId {
		spouseName = marryObj.SpouseName
	} else {
		spouseName = marryObj.PlayerName
	}
	return
}

//同步配偶表白等级
func (ms *marryService) GetSpouseDevelopLevel(playerId int64) int32 {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()

	marryObj, exist := ms.marryMap[playerId]
	if !exist {
		return 0
	}
	if marryObj.PlayerId == playerId {
		return marryObj.SpouseDevelopLevel
	} else {
		return marryObj.DevelopLevel
	}
}

//获取配偶id
func (ms *marryService) GetMarrySpouseId(playerId int64) (spouseId int64) {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()

	marryObj, exist := ms.marryMap[playerId]
	if !exist {
		return
	}
	if marryObj.PlayerId == playerId {
		spouseId = marryObj.SpouseId
	}
	if marryObj.SpouseId == playerId {
		spouseId = marryObj.PlayerId
	}
	return
}

func (ms *marryService) SyncMarryDingQing(playerId int64, dqMap map[int32]map[int32]int32) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	marryObj, exist := ms.marryMap[playerId]
	if !exist {
		return
	}
	newMap := ms.copySuitMap(dqMap)
	if marryObj.PlayerId == playerId {
		marryObj.PlayerSuit = newMap
	} else {
		marryObj.SpouseSuit = newMap
	}

}

func (ms *marryService) GetMarryDingQing(playerId int64) map[int32]map[int32]int32 {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()
	return ms.getMarryDingQing(playerId)
}

func (ms *marryService) getMarryDingQing(playerId int64) map[int32]map[int32]int32 {
	marryObj, exist := ms.marryMap[playerId]
	if !exist {
		return nil
	}
	var suitMap map[int32]map[int32]int32
	if marryObj.PlayerId == playerId {
		suitMap = marryObj.PlayerSuit
	} else {
		suitMap = marryObj.SpouseSuit
	}
	newMap := ms.copySuitMap(suitMap)
	return newMap
}

func (ms *marryService) copySuitMap(p_suitMap map[int32]map[int32]int32) map[int32]map[int32]int32 {
	newMap := make(map[int32]map[int32]int32)
	for key, value := range p_suitMap {
		_, exists := newMap[key]
		if !exists {
			newMap[key] = make(map[int32]int32)
		}
		for vkey, vvalue := range value {
			newMap[key][vkey] = vvalue
		}
	}
	return newMap
}

func (ms *marryService) ExistsSpouseDingQing(playerId int64, suiId int32, posId int32) bool {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()
	marryObj, exist := ms.marryMap[playerId]
	if !exist {
		return false
	}
	var suitMap map[int32]map[int32]int32
	if marryObj.PlayerId == playerId {
		suitMap = marryObj.SpouseSuit
	} else {
		suitMap = marryObj.PlayerSuit
	}
	if suitMap == nil {
		return false
	}
	suit, exists := suitMap[suiId]
	if !exists {
		return false
	}
	_, exists = suit[posId]

	return exists
}

func (ms *marryService) GetSpouseId(playerId int64) int64 {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()
	return ms.getSpouseId(playerId)
}

func (ms *marryService) getSpouseId(playerId int64) int64 {
	marryObj, exist := ms.marryMap[playerId]
	if !exist {
		return 0
	}
	if marryObj.PlayerId == playerId {
		return marryObj.SpouseId
	} else {
		return marryObj.PlayerId
	}
}

//获取求婚婚戒
func (ms *marryService) GetMarryProposalRing(playerId int64) *MarryRingObject {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()

	marryRingObj, exist := ms.marryRingMap[playerId]
	if !exist {
		return nil
	}
	return marryRingObj
}

//求婚婚戒
func (ms *marryService) MarryProposalRing(playerId int64, peerId int64, peerName string, ringType marrytypes.MarryRingType) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	mro := NewMarryRingObject()
	id, _ := idutil.GetId()
	mro.Id = id
	mro.ServerId = global.GetGame().GetServerIndex()
	mro.PlayerId = playerId
	mro.PeerId = peerId
	mro.PeerName = peerName
	mro.Ring = ringType
	mro.ProposalTime = now
	mro.SetModified()
	ms.marryRingMap[playerId] = mro
}

//添加亲密度
func (ms *marryService) AddPoint(playerId int64, addPoint int32) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	marryObj, exist := ms.marryMap[playerId]
	if exist {
		marryObj.Point += addPoint
		marryObj.UpdateTime = now
		marryObj.SetModified()
	}

	return
}

//婚戒替换
func (ms *marryService) MarryReplace(playerId int64, ringType marrytypes.MarryRingType) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	marryObj, exist := ms.marryMap[playerId]
	if !exist {
		return
	}
	if marryObj.Ring >= ringType {
		return
	}
	marryObj.Ring = ringType
	marryObj.UpdateTime = now
	marryObj.SetModified()
	return
}

//获取婚宴场景
func (ms *marryService) GetMarrySceneData() *marryscene.MarrySceneStatusData {
	return ms.sceneStatusData
}

//获取玩家的婚礼戒指信息
func (ms *marryService) GetMarryRing(playerId int64) (ringType marrytypes.MarryRingType) {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()

	marryObj, exist := ms.marryMap[playerId]
	if !exist {
		return -1
	}
	return marryObj.Ring
}

//获取婚烟信息
func (ms *marryService) GetMarry(playerId int64) *MarryObject {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()

	marryObj, flag := ms.marryMap[playerId]
	if !flag {
		return nil
	}
	return marryObj
}

func (ms *marryService) getMarryWedByTime(now int64) (marryWed *MarryWedObject, pl player.Player, spl player.Player) {
	marryDurationTime := marrytemplate.GetMarryTemplateService().GetMarryDurationTime()
	for _, marryWedObj := range ms.marryWedList {
		if marryWedObj.Period == -1 || marryWedObj.Status != marrytypes.MarryWedStatusTypeNoStart {
			continue
		}
		hTime := marryWedObj.HTime

		if now >= hTime && now < marryDurationTime+hTime {
			//判断玩家双方是否在线
			playerId := marryWedObj.PlayerId
			spouseId := marryWedObj.SpouseId
			marryWed = marryWedObj

			pl = player.GetOnlinePlayerManager().GetPlayerById(playerId)
			spl = player.GetOnlinePlayerManager().GetPlayerById(spouseId)
			if pl == nil || spl == nil {
				break
			}
			//跨服默认不在线,取消婚礼
			if pl != nil && pl.IsCross() {
				pl = nil
			}
			if spl != nil && spl.IsCross() {
				spl = nil
			}
			return marryWed, pl, spl
		}
	}
	return marryWed, pl, spl
}

//设置结婚场景的婚礼状态
func (ms *marryService) WeddingCarEnd() {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()
	//避免gm刷婚车
	if ms.sceneStatusData.Status != marryscene.MarrySceneStatusCruise {
		return
	}
	ms.sceneStatusData.Status = marryscene.MarrySceneStatusBanquet
	ms.huncheNpc = nil
}

//获取婚车对象
func (ms *marryService) GetHunCheNpc() *hunche.HunCheNPC {
	return ms.huncheNpc
}

//设置婚车对象
func (ms *marryService) WeddingCarNPC(huncheNpc *hunche.HunCheNPC) {
	ms.huncheNpc = huncheNpc
}

//系统举办婚礼取消
func (ms *marryService) cancleMarryWedding(playerId int64, now int64) {

	marryObj, exist := ms.marryMap[playerId]
	if !exist {
		return
	}
	if marryObj.Status != marrytypes.MarryStatusTypeEngagement {
		return
	}
	marryObj.Status = marrytypes.MarryStatusTypeProposal
	marryObj.UpdateTime = now
	marryObj.SetModified()
}

func (ms *marryService) refreshMarryRingMap(now int64) {
	for playerId, marryRingObj := range ms.marryRingMap {
		proposalTime := marryRingObj.ProposalTime
		diffTime := now - proposalTime
		if diffTime > marrytypes.RingTime {
			ms.marryRingStatusFail(playerId, marryRingObj.Ring, now)
		}
	}
}

func (ms *marryService) refreshMarryPreWedMap(now int64) {
	for playerId, marryPreWedObj := range ms.marryPreWedMap {
		preWedTime := marryPreWedObj.PreWedTime
		diffTime := now - preWedTime
		if diffTime > marrytypes.PreWedTime {
			ms.marryPreWedStatusFail(playerId, now, true)
		}
	}
}

func (ms *marryService) refreshMarryMap(now int64) {
	ms.refreshMarryRingMap(now)
	ms.refreshMarryPreWedMap(now)
}

//心跳
func (ms *marryService) Heartbeat() {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()
	now := global.GetGame().GetTimeService().Now()
	ms.refreshMarryMap(now)
	if ms.sceneStatusData.Period == -1 {
		marryWedObj, pl, spl := ms.getMarryWedByTime(now)
		if marryWedObj == nil {
			return
		}
		period := int32(marryWedObj.Period)
		grade := int32(marryWedObj.Grade)
		hunCheGrade := int32(marryWedObj.HunCheGrade)
		sugarGrade := int32(marryWedObj.SugarGrade)
		marryObj := ms.marryMap[marryWedObj.PlayerId]
		if pl == nil || spl == nil {
			//婚宴取消
			ms.cancleMarryWedding(marryWedObj.PlayerId, now)
			marryWedObj.Status = marrytypes.MarryWedStatusTypeCancle
			gameevent.Emit(marryeventtypes.EventTypeMarryWedCancle, marryWedObj, marryObj)
		} else {
			//修改结婚状态
			marryObj.Status = marrytypes.MarryStatusTypeMarried
			marryObj.UpdateTime = now
			marryObj.SetModified()
			//婚宴开始
			ms.setSceneStatusData(pl, spl, marryWedObj)
			marryWedObj.Status = marrytypes.MarryWedStatusTypeOngoing
			eventData := marryeventtypes.CreateMarryWedStartEventData(marryObj.Id, period, grade, hunCheGrade, sugarGrade, pl, spl, marryWedObj.IsFirst)
			gameevent.Emit(marryeventtypes.EventTypeMarryWedStart, nil, eventData)
		}
		marryWedObj.LastTime = now
		marryWedObj.UpdateTime = now
		marryWedObj.SetModified()
	} else {
		marryWedObj := ms.getWedOngoing(ms.sceneStatusData.Period)
		if marryWedObj == nil {
			return
		}

		marryDurationTime := marrytemplate.GetMarryTemplateService().GetMarryDurationTime()
		endTime := marryWedObj.HTime + marryDurationTime
		if now >= endTime {
			ms.initSceneStatusData()
			marryWedObj.Status = marrytypes.MarryWedStatusTypeHeld
			marryWedObj.LastTime = now
			marryWedObj.UpdateTime = now
			marryWedObj.SetModified()
			//发送事件
			gameevent.Emit(marryeventtypes.EventTypeMarryWedEnd, nil, nil)
		}
	}

}

func (ms *marryService) setSceneStatusData(pl player.Player, spl player.Player, marryWed *MarryWedObject) {
	ms.sceneStatusData.Id = marryWed.Id
	ms.sceneStatusData.Status = marryscene.MarrySceneStatusCruise
	ms.sceneStatusData.Period = marryWed.Period
	ms.sceneStatusData.Grade = marryWed.Grade
	ms.sceneStatusData.PlayerId = pl.GetId()
	ms.sceneStatusData.PlayerName = pl.GetName()
	ms.sceneStatusData.Role = int32(pl.GetRole())
	ms.sceneStatusData.Sex = int32(pl.GetSex())
	ms.sceneStatusData.SpouseId = spl.GetId()
	ms.sceneStatusData.SpouseName = spl.GetName()
	ms.sceneStatusData.SpouseRole = int32(spl.GetRole())
	ms.sceneStatusData.SpouseSex = int32(spl.GetSex())
}

func (ms *marryService) initSceneStatusData() {
	ms.sceneStatusData.Id = 0
	ms.sceneStatusData.Status = marryscene.MarrySceneStatusTypeInit
	ms.sceneStatusData.Period = -1
	ms.sceneStatusData.Grade = 0
	ms.sceneStatusData.PlayerId = 0
	ms.sceneStatusData.PlayerName = ""
	ms.sceneStatusData.Role = 1
	ms.sceneStatusData.Sex = 1
	ms.sceneStatusData.SpouseId = 0
	ms.sceneStatusData.SpouseName = ""
	ms.sceneStatusData.SpouseRole = 1
	ms.sceneStatusData.SpouseSex = 1
}

//结婚场景创建
func (ms *marryService) CreateMarrySceneData() {
	marryTemplate := marrytemplate.GetMarryTemplateService().GetMarryConstTempalte()
	marryData := marryscene.CreateMarrySceneData()
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(marryTemplate.MarryMapId)
	if mapTemplate == nil {
		return
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeMarry {
		return
	}
	ms.marryScene = scene.CreateScene(mapTemplate, 0, marryData)
	ms.sceneStatusData = marryscene.CreateMarrySceneStatusData(0, -1, 0, marryscene.MarrySceneStatusTypeInit, 0, 0, "", "", 1, 1, 1, 1)
}

//获取结婚场景
func (ms *marryService) GetScene() scenescene.Scene {
	return ms.marryScene
}

//赠送贺礼
func (ms *marryService) IfCanGiveWedGift(period int32) (flag bool) {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()
	if ms.sceneStatusData.Status != marryscene.MarrySceneStatusTypeInit &&
		ms.sceneStatusData.Period == period {
		flag = true
	}
	return
}

//获取当前进行的婚期
func (ms *marryService) getWedOngoing(period int32) (marryWed *MarryWedObject) {
	for _, marryWedObj := range ms.marryWedList {
		if marryWedObj.Period == period {
			marryWed = marryWedObj
			break
		}
	}
	return
}

//离婚决策
func (ms *marryService) DivorceDeal(pl player.Player, spouseId int64, result marrytypes.MarryResultType) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	agree := true
	marryObj, exist := ms.marryMap[pl.GetId()]
	if !exist {
		return
	}

	spouseName := marryObj.SpouseName
	if marryObj.PlayerId != pl.GetId() {
		spouseName = marryObj.PlayerName
	}
	marryReturnList := make([]*marryeventtypes.MarryDivorceDealEventDataItem, 0)
	//同意
	if result == marrytypes.MarryResultTypeOk {

		now := global.GetGame().GetTimeService().Now()
		marryObj.Status = marrytypes.MarryStatusTypeDivorce
		marryObj.DeleteTime = now
		marryObj.SetModified()
		delete(ms.marryMap, pl.GetId())
		delete(ms.marryMap, spouseId)
		preMarryList := ms.marryPreWedRemove(pl)
		for _, value := range preMarryList {
			item := marryeventtypes.CreateMarryDivorceDealEventDataItem(value.PlayerId, value.Grade, value.HunCheGrade, value.SugarGrade)
			marryReturnList = append(marryReturnList, item)
		}
		ms.cancelWedding(pl.GetId())
	} else {
		agree = false
	}

	//发送事件
	eventData := marryeventtypes.CreateMarryDivorceDealEventData(agree, spouseId, spouseName, marryReturnList)
	gameevent.Emit(marryeventtypes.EventTypeMarryDivorceDeal, pl, eventData)
	return
}

//婚期取消
func (ms *marryService) cancelWedding(playerId int64) {
	now := global.GetGame().GetTimeService().Now()
	for _, marryWedObj := range ms.marryWedList {
		if (marryWedObj.PlayerId == playerId || marryWedObj.SpouseId == playerId) &&
			marryWedObj.Status == marrytypes.MarryWedStatusTypeNoStart {
			marryWedObj.Status = marrytypes.MarryWedStatusTypeCancle
			marryWedObj.LastTime = now
			marryWedObj.UpdateTime = now
			marryWedObj.Period = -1
			marryWedObj.PlayerId = 0
			marryWedObj.Name = ""
			marryWedObj.SpouseId = 0
			marryWedObj.SpouseName = ""
			marryWedObj.Grade = 0
			marryWedObj.HunCheGrade = 0
			marryWedObj.SugarGrade = 0
			marryWedObj.HTime = 0
			marryWedObj.SetModified()
		}
	}
	return
}

//离婚
func (ms *marryService) Divorce(pl player.Player, spouseId int64, typ marrytypes.MarryDivorceType) (err error) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	marryReturnList := make([]*marryeventtypes.MarryDivorceDealEventDataItem, 0)
	//强制离婚
	if typ == marrytypes.MarryDivorceTypeForce {
		// ms.cancelWedding(pl.GetId())
		marryObj, eixst := ms.marryMap[pl.GetId()]
		if eixst {
			marryObj.Status = marrytypes.MarryStatusTypeDivorce
			marryObj.DeleteTime = now
			marryObj.SetModified()
		}

		delete(ms.marryMap, pl.GetId())
		delete(ms.marryMap, spouseId)
		preMarryList := ms.marryPreWedRemove(pl)
		for _, value := range preMarryList {
			item := marryeventtypes.CreateMarryDivorceDealEventDataItem(value.PlayerId, value.Grade, value.HunCheGrade, value.SugarGrade)
			marryReturnList = append(marryReturnList, item)
		}
		ms.cancelWedding(pl.GetId())
	} else {
		spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
		if spl == nil {
			err = ErrorMarryDivorceNoOnline
			return
		}
	}

	//发送事件
	eventData := marryeventtypes.CreateMarryDivorceEventData(spouseId, typ, marryReturnList)
	gameevent.Emit(marryeventtypes.EventTypeMarryDivorce, pl, eventData)
	return
}

func (ms *marryService) newMarryObject(pl player.Player, spl player.Player, ringLevel int32, ringType marrytypes.MarryRingType, point int32) {
	now := global.GetGame().GetTimeService().Now()
	mo := NewMarryObject()
	playerId := pl.GetId()
	playerName := pl.GetName()
	spouseId := spl.GetId()
	spouseName := spl.GetName()
	id, _ := idutil.GetId()
	mo.Id = id
	mo.ServerId = global.GetGame().GetServerIndex()
	mo.PlayerId = playerId
	mo.PlayerName = playerName
	mo.SpouseId = spouseId
	mo.SpouseName = spouseName
	mo.PlayerRingLevel = 0
	if ringLevel == 0 {
		ringLevel = 1
	}
	mo.Role = int32(pl.GetRole())
	mo.Sex = int32(pl.GetSex())
	mo.SpouseRole = int32(spl.GetRole())
	mo.SpouseSex = int32(spl.GetSex())
	mo.SpouseRingLevel = ringLevel
	mo.PlayerRingLevel = ringLevel
	mo.DevelopLevel = pl.GetMarryDevelopLevel()
	mo.SpouseDevelopLevel = spl.GetMarryDevelopLevel()
	mo.Point = point
	mo.Ring = ringType
	mo.Status = marrytypes.MarryStatusTypeProposal
	mo.CreateTime = now
	ms.marryMap[playerId] = mo
	ms.marryMap[spouseId] = mo
	mo.SetModified()
	return
}

func (ms *marryService) RemoveMarryRing(playerId int64) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	marryRingObj, exist := ms.marryRingMap[playerId]
	if !exist {
		return
	}
	delete(ms.marryRingMap, playerId)

	now := global.GetGame().GetTimeService().Now()
	marryRingObj.DeleteTime = now
	marryRingObj.UpdateTime = now
	marryRingObj.SetModified()
}

func (ms *marryService) marryRingStatusFail(playerId int64, ringType marrytypes.MarryRingType, now int64) {
	marryRingObj, exist := ms.marryRingMap[playerId]
	if !exist {
		return
	}
	marryRingObj.Status = marrytypes.MarryRingStatusTypeFail
	marryRingObj.UpdateTime = now
	marryRingObj.SetModified()
	eventTypeData := marryeventtypes.CreateMarryGiveBackRingEventData(ringType, marryRingObj.PeerName)

	gameevent.Emit(marryeventtypes.EventTypeMarryRingGiveBack, playerId, eventTypeData)

}

func (ms *marryService) marryRingStatusSucess(playerId int64, now int64) {
	marryRingObj, exist := ms.marryRingMap[playerId]
	if !exist {
		return
	}
	marryRingObj.Status = marrytypes.MarryRingStatusTypeSucess
	marryRingObj.UpdateTime = now
	marryRingObj.DeleteTime = now
	marryRingObj.SetModified()

	delete(ms.marryRingMap, playerId)
}

//被求婚者决策
func (ms *marryService) ProposalDeal(pl player.Player, dealId int64, ringLevel int32, ringType marrytypes.MarryRingType, point int32, result marrytypes.MarryResultType) (err error) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	agree := true
	ppl := player.GetOnlinePlayerManager().GetPlayerById(dealId)
	now := global.GetGame().GetTimeService().Now()
	//同意
	if result == marrytypes.MarryResultTypeOk {
		if ppl == nil {
			//求婚者下线
			ms.marryRingStatusFail(dealId, ringType, now)
			err = ErrorMarryProposalIsNoOnline
			return
		}

		//求婚者性别修改了
		if ppl.GetSex() == pl.GetSex() {
			ms.marryRingStatusFail(dealId, ringType, now)
			err = ErrorMarryDealIsSexChanged
			return
		}

		if _, exist := ms.marryMap[pl.GetId()]; exist {
			//决策者已婚
			ms.marryRingStatusFail(dealId, ringType, now)
			err = ErrorMarryDealIsMarried
			return
		}

		ms.marryRingStatusSucess(dealId, now)
		ms.newMarryObject(ppl, pl, ringLevel, ringType, point)
	} else {
		ms.marryRingStatusFail(dealId, ringType, now)
		agree = false
	}

	//发送事件
	eventData := marryeventtypes.CreateMarryProposalDealEventData(agree, dealId)
	gameevent.Emit(marryeventtypes.EventTypeMarryProposalDeal, pl, eventData)
	return
}

//获取喜帖列表
func (ms *marryService) GetWeddingCardList() (wedCardList []*MarryWedCardObject) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	ms.refreshWeddingCard(now)
	for _, wedCardObj := range ms.marryWedCardMap {
		wedCardList = append(wedCardList, wedCardObj)
	}
	if len(wedCardList) > 0 {
		sort.Sort(sort.Reverse(MarryWedCardList(wedCardList)))
	}

	return
}

func (ms *marryService) refreshWeddingCard(now int64) {
	for wedCardId, wedCardObj := range ms.marryWedCardMap {
		if wedCardObj.OutOfTime <= now {
			delete(ms.marryWedCardMap, wedCardId)
			continue
		}
	}
}

//喜帖是否存在
func (ms *marryService) WeddingCardIsExist(cardId int64) (existFlag bool) {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()
	_, existFlag = ms.marryWedCardMap[cardId]
	return
}

//获取婚礼举办场次
func (ms *marryService) GetWeddingPeriod(playerId int64) (period int32) {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()
	period = -1
	ms.refreshMarryWed()
	for _, marryWed := range ms.marryWedList {
		if (marryWed.PlayerId == playerId || marryWed.SpouseId == playerId) &&
			marryWed.Status == marrytypes.MarryWedStatusTypeNoStart {
			return marryWed.Period
		}
	}
	return
}

//获取婚期预定列表
func (ms *marryService) GetMarryWeddingList() []*MarryWedObject {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()
	ms.refreshMarryWed()
	return ms.marryWedList
}

func (ms *marryService) newMarryWedCardObject(playerId int64, playerName string, spouseId int64, spouseName string, reserveTime int64, hTime string) (mwco *MarryWedCardObject) {
	now := global.GetGame().GetTimeService().Now()
	ms.refreshWeddingCard(now)
	mwco = NewMarryWedCardObject()
	id, _ := idutil.GetId()
	mwco.Id = id

	mwco.ServerId = global.GetGame().GetServerIndex()
	mwco.PlayerId = playerId
	mwco.PlayerName = playerName
	mwco.SpouseId = spouseId
	mwco.SpouseName = spouseName
	mwco.HoldTime = hTime
	mwco.OutOfTime = reserveTime
	mwco.CreateTime = now
	ms.marryWedCardMap[id] = mwco
	mwco.SetModified()
	return mwco
}

//跨天刷新婚期预定列表
func (ms *marryService) refreshMarryWed() {
	now := global.GetGame().GetTimeService().Now()
	refreshFlag := false
	for _, marryWedObj := range ms.marryWedList {
		flag, _ := timeutils.IsSameDay(marryWedObj.LastTime, now)
		// flag, _ := timeutils.IsSameFive(marryWedObj.LastTime, now)
		if !flag {
			refreshFlag = true
		}
		break
	}
	if refreshFlag {
		for _, marryWedObj := range ms.marryWedList {
			marryWedObj.DeleteTime = now
			marryWedObj.SetModified()
		}
		ms.marryWedList = make([]*MarryWedObject, 0, 12)
	}
}

func (ms *marryService) ifExistPeriod(period int32) bool {
	for _, marryWed := range ms.marryWedList {
		if period == marryWed.Period {
			return true
		}
	}
	return false
}

func (ms *marryService) newMarryWedObject(period int32, marryGrade *marrytypes.MarryGrade, playerId int64, playerName string, spouseId int64, spouseName string, reserveTime int64, now int64, isFirst bool) {
	mwo := NewMarryWedObject()
	id, _ := idutil.GetId()
	mwo.Id = id
	mwo.ServerId = global.GetGame().GetServerIndex()
	mwo.Period = period
	mwo.Grade = marryGrade.Grade
	mwo.HunCheGrade = marryGrade.HunCheGrade
	mwo.SugarGrade = marryGrade.SugarGrade
	mwo.Status = marrytypes.MarryWedStatusTypeNoStart
	mwo.PlayerId = playerId
	mwo.Name = playerName
	mwo.SpouseId = spouseId
	mwo.SpouseName = spouseName
	mwo.HTime = reserveTime
	mwo.CreateTime = now
	mwo.LastTime = now
	mwo.IsFirst = isFirst
	ms.marryWedList = append(ms.marryWedList, mwo)
	mwo.SetModified()
	return
}

func (ms *marryService) marryWedding(period int32, marryGrade *marrytypes.MarryGrade, playerId int64, playerName string, spouseId int64, spouseName string, reserveTime int64, isFirst bool) {
	now := global.GetGame().GetTimeService().Now()
	updateFlag := false
	for _, marryWedObj := range ms.marryWedList {
		if marryWedObj.Period == -1 && !updateFlag {
			marryWedObj.Period = period
			marryWedObj.Grade = marryGrade.Grade
			marryWedObj.HunCheGrade = marryGrade.HunCheGrade
			marryWedObj.SugarGrade = marryGrade.SugarGrade
			marryWedObj.PlayerId = playerId
			marryWedObj.Name = playerName
			marryWedObj.SpouseId = spouseId
			marryWedObj.Status = marrytypes.MarryWedStatusTypeNoStart
			marryWedObj.SpouseName = spouseName
			marryWedObj.HTime = reserveTime
			marryWedObj.IsFirst = isFirst
			updateFlag = true

			marryWedObj.UpdateTime = now
			marryWedObj.LastTime = now
			marryWedObj.SetModified()
			break
		}

	}
	if !updateFlag {
		ms.newMarryWedObject(period, marryGrade, playerId, playerName, spouseId, spouseName, reserveTime, now, isFirst)
	}

	marryObj, exist := ms.marryMap[playerId]
	if exist {
		marryObj.Status = marrytypes.MarryStatusTypeEngagement
		marryObj.UpdateTime = now
		marryObj.SetModified()
	}
	return
}

func (ms *marryService) marryPreWedStatusFail(playerId int64, now int64, isRefuse bool) {
	marryPreWedObj, exist := ms.marryPreWedMap[playerId]
	if !exist {
		return
	}
	defer delete(ms.marryPreWedMap, playerId)

	marryPreWedObj.Status = marrytypes.MarryPreWedStatusTypeFail
	marryPreWedObj.UpdateTime = now
	marryPreWedObj.DeleteTime = now
	marryPreWedObj.SetModified()

	grade := marryPreWedObj.Grade
	hunCheGrade := marryPreWedObj.HunCheGrade
	sugarGrade := marryPreWedObj.SugarGrade
	eventData := marryeventtypes.CreateMarryPreWedGiveBackEventData(isRefuse, grade, hunCheGrade, sugarGrade)
	gameevent.Emit(marryeventtypes.EventTypeMarryPreWedGiveBack, playerId, eventData)
}

func (ms *marryService) marryPreWedStatusSucess(playerId int64, now int64) {
	marryPreWedObj, exist := ms.marryPreWedMap[playerId]
	if !exist {
		return
	}

	defer delete(ms.marryPreWedMap, playerId)

	marryPreWedObj.Status = marrytypes.MarryPreWedStatusTypeSucess
	marryPreWedObj.UpdateTime = now
	marryPreWedObj.DeleteTime = now
	marryPreWedObj.SetModified()

}

func (ms *marryService) MarryPreWedIsExist(playerId int64) (flag bool) {
	ms.rwm.RLock()
	defer ms.rwm.RUnlock()

	_, flag = ms.marryPreWedMap[playerId]
	return
}

//婚期预定判断
func (ms *marryService) MarryPreWedding(pl player.Player, period int32, marryGrade *marrytypes.MarryGrade, spouseId int64, reserveTime int64) (err error) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	//该婚期已存在
	exist := ms.ifExistPeriod(period)
	if exist {
		err = ErrorMarryWedIsBeScheduled
		return
	}
	//是否已结婚
	_, exist = ms.marryMap[pl.GetId()]
	if !exist {
		err = ErrorMarryReserveNoMarried
		return
	}
	//配偶不在线
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	if spl == nil {
		err = ErrorMarryPreWedSpouseNoOnline
		return
	}

	ms.initMarryPreWedObj(pl, period, marryGrade, spouseId, reserveTime)
	eventData := marryeventtypes.CreateMarryPreWedEventData(period, pl.GetName(), spouseId, marryGrade)
	gameevent.Emit(marryeventtypes.EventTypeMarryPreWed, pl, eventData)
	return
}

func (ms *marryService) isPreWedExist(peerId int64) (preWedObj *MarryPreWedObject, flag bool) {
	for _, marryPreWed := range ms.marryPreWedMap {
		if marryPreWed.Status != marrytypes.MarryPreWedStatusTypeOngoing {
			continue
		}
		if marryPreWed.PeerId != peerId {
			continue
		}
		return marryPreWed, true
	}
	return
}

//婚期预定配偶决策
func (ms *marryService) MarryPreWedDeal(spl player.Player, result bool, isFirst bool) (err error) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	preWedObj, flag := ms.isPreWedExist(spl.GetId())
	if !flag {
		err = ErrorMarryPreWedIsOverdue
		return
	}

	playerId := preWedObj.PlayerId
	playerName := preWedObj.PlayerName
	period := preWedObj.Period
	marryGrade := marrytypes.CreateMarryGrade(preWedObj.Grade, preWedObj.HunCheGrade, preWedObj.SugarGrade)
	reserveTime := preWedObj.HoldTime
	spouseId := spl.GetId()
	spouseName := spl.GetName()
	if !result {
		ms.marryPreWedStatusFail(playerId, now, true)
		return
	}
	//婚期被抢先预定了
	if ms.ifExistPeriod(period) {
		ms.marryPreWedStatusFail(playerId, now, false)
		err = ErrorMarryWedIsBeScheduled
		return
	}
	ms.marryPreWedStatusSucess(playerId, now)
	ms.marryWedding(period, marryGrade, playerId, playerName, spouseId, spouseName, reserveTime, isFirst)
	year, month, day, hour, min := timeutils.GetYearMonthDay(reserveTime)
	marryDurationTime := marrytemplate.GetMarryTemplateService().GetMarryDurationTime()
	_, _, _, nextHour, nextMin := timeutils.GetYearMonthDay(reserveTime + marryDurationTime)
	hTime := fmt.Sprintf("%d年%d月%d日%d:%d~%d:%d", year, month, day, hour, min, nextHour, nextMin)
	mwco := ms.newMarryWedCardObject(playerId, playerName, spouseId, spouseName, reserveTime, hTime)
	//发送事件
	eventData := marryeventtypes.CreateMarryWedEventData(period, spl, playerId, playerName, hTime)
	gameevent.Emit(marryeventtypes.EventTypeMarryWed, mwco, eventData)
	return
}

//预定婚期取消,返回取消的婚礼
func (ms *marryService) marryPreWedRemove(spl player.Player) []*MarryWedObject {
	result := make([]*MarryWedObject, 0)
	playerId := spl.GetId()
	for _, value := range ms.marryWedList {
		if value.Status != marrytypes.MarryWedStatusTypeNoStart {
			continue
		}
		if value.PlayerId == playerId || value.SpouseId == playerId {
			result = append(result, value)
			value.Period = -1 //目前直接这样子删除，释放出去的在求婚那边有判断
			value.SetModified()
		}
	}
	return result
}

//婚期预定
// func (ms *marryService) MarryWedding(pl player.Player, period int32, marryGrade *marrytypes.MarryGrade, spouseId int64, spouseName string, reserveTime int64) (wedCode marrytypes.MarryWedCodeType) {
// 	ms.rwm.Lock()
// 	defer ms.rwm.Unlock()

// 	wedCode = marrytypes.MarryCodeTypeWeddingSucess
// 	//该婚期已存在
// 	exist := ms.ifExistPeriod(period)
// 	if exist {
// 		wedCode = marrytypes.MarryCodeTypeWeddingExist
// 		return
// 	}

// 	ms.marryWedding(period, marryGrade, pl, spouseId, spouseName, reserveTime)
// 	now := global.GetGame().GetTimeService().Now()
// 	year, month, day := timeutils.GetYearMonthDay(now)
// 	marryDurationTime := marrytemplate.GetMarryTemplateService().GetMarryDurationTime() / int64(common.MINUTE)
// 	hTime := fmt.Sprintf("%d年%d月%d日%d:00~%d:%d", year, month, day, 10+period-1, 10+period-1, marryDurationTime)
// 	mwco := ms.newMarryWedCardObject(pl.GetId(), pl.GetName(), spouseId, spouseName, reserveTime, hTime)

// 	//发送事件
// 	eventData := marryeventtypes.CreateMarryWedEventData(period, pl, spouseId, spouseName)
// 	gameevent.Emit(marryeventtypes.EventTypeMarryWed, mwco, eventData)
// 	return
// }

func (ms *marryService) playerChangeNameMarry(pl player.Player) {
	marryObj, exist := ms.marryMap[pl.GetId()]
	if !exist {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if marryObj.PlayerId == pl.GetId() {
		marryObj.PlayerName = pl.GetName()
	} else {
		marryObj.SpouseName = pl.GetName()
	}
	marryObj.UpdateTime = now
	marryObj.SetModified()
	gameevent.Emit(marryeventtypes.EventTypeMarryPlayerNameChanged, marryObj, pl.GetId())
}

func (ms *marryService) playerChangeNameWeddingCard(pl player.Player) {
	now := global.GetGame().GetTimeService().Now()
	for _, weddingCard := range ms.marryWedCardMap {
		if weddingCard.PlayerId == pl.GetId() {
			weddingCard.PlayerName = pl.GetName()
			weddingCard.UpdateTime = now
			weddingCard.SetModified()
		}
		if weddingCard.SpouseId == pl.GetId() {
			weddingCard.SpouseName = pl.GetName()
			weddingCard.UpdateTime = now
			weddingCard.SetModified()
		}
	}
}

func (ms *marryService) playerChangeNameWedding(pl player.Player) {
	now := global.GetGame().GetTimeService().Now()
	for _, marryWed := range ms.marryWedList {
		if marryWed.PlayerId == pl.GetId() {
			marryWed.Name = pl.GetName()
			marryWed.UpdateTime = now
			marryWed.SetModified()
		}
		if marryWed.SpouseId == pl.GetId() {
			marryWed.SpouseName = pl.GetName()
			marryWed.UpdateTime = now
			marryWed.SetModified()
		}
	}
}

func (ms *marryService) playerChangeNameSceneData(pl player.Player) {
	if ms.sceneStatusData.PlayerId == pl.GetId() {
		ms.sceneStatusData.PlayerName = pl.GetName()
	}
	if ms.sceneStatusData.SpouseId == pl.GetId() {
		ms.sceneStatusData.SpouseName = pl.GetName()
	}
}

//玩家修改名字
func (ms *marryService) PlayerNameChanged(pl player.Player) {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()
	ms.playerChangeNameMarry(pl)
	ms.playerChangeNameWeddingCard(pl)
	ms.playerChangeNameWedding(pl)
	ms.playerChangeNameSceneData(pl)
}

//gm 使用 清空婚期
func (ms *marryService) GmClearMarryWed() {
	ms.rwm.Lock()
	defer ms.rwm.Unlock()

	for _, marryWedObj := range ms.marryWedList {
		marryWedObj.Period = -1
		marryWedObj.Grade = 0
		marryWedObj.Status = marrytypes.MarryWedStatusTypeNoStart
		marryWedObj.PlayerId = 0
		marryWedObj.SpouseId = 0
		marryWedObj.Name = ""
		marryWedObj.SpouseName = ""
		marryWedObj.HTime = 0
		marryWedObj.LastTime = 0
		marryWedObj.SetModified()
	}

	ms.initSceneStatusData()
	ms.huncheNpc = nil
	return
}

var (
	once sync.Once
	cs   *marryService
	// mset *marrySetService
)

func Init() (err error) {
	once.Do(func() {
		cs = &marryService{}
		err = cs.init()
		// mset = &marrySetService{}
		// mset.init()
	})
	return err
}

func GetMarryService() MarryService {
	return cs
}

// func GetMarrySetService() MarrySetService {
// 	return mset
// }
