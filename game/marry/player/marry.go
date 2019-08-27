package player

import (
	commonlog "fgame/fgame/common/log"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/marry/dao"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//玩家结婚管理器
type PlayerMarryDataManager struct {
	p player.Player
	//玩家结婚
	playerMarryObject *PlayerMarryObject
	//玩家求婚map
	proposalMap map[int64]int64
	//玩家求婚婚戒类型
	ringTypeMap map[int64]marrytypes.MarryRingType
	//收到离婚消息
	receiveDivorce bool
	//玩家豪气值
	playerHeroismObject *PlayerMarryHeroismObject
	//推送玩家婚礼按钮记录
	playerPushWedRecordObj *PlayerPushWedRecordObject
	//玩家点击婚车时间
	clickTime int64
	//玩家收到预定婚期
	preWedFlag bool
	//查看过的喜帖id
	viewCardList []int64

	//玩家结婚纪念
	jinianMap map[marrytypes.MarryBanquetSubTypeWed]*PlayerMarryJiNianObject
	//玩家结婚纪念时装获取信息
	jinianSjObj *PlayerMarryJiNianSjObject

	//定情信物
	dingQing *PlayerMarryDingQingObject
	//结婚伴侣定情信物
	spouseSuitMap map[int32]map[int32]int32
	//信物索取的时间
	xinwuTime int64
}

func (m *PlayerMarryDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerMarryDataManager) Load() (err error) {
	m.proposalMap = make(map[int64]int64)
	m.ringTypeMap = make(map[int64]marrytypes.MarryRingType)
	m.viewCardList = make([]int64, 0, 8)
	m.jinianMap = make(map[marrytypes.MarryBanquetSubTypeWed]*PlayerMarryJiNianObject)
	m.spouseSuitMap = make(map[int32]map[int32]int32)
	m.receiveDivorce = false
	m.clickTime = 0
	m.xinwuTime = 0

	//加载玩家结婚信息
	marryEntity, err := dao.GetMarryDao().GetMarryPlayerEntity(m.p.GetId())
	if err != nil {
		return
	}
	if marryEntity == nil {
		m.initPlayerMarryObject()
	} else {
		m.playerMarryObject = NewPlayerMarryObject(m.p)
		m.playerMarryObject.FromEntity(marryEntity)
	}

	//豪气值
	heroismEntity, err := dao.GetMarryDao().GetHeroismEntity(m.p.GetId())
	if err != nil {
		return
	}
	if heroismEntity == nil {
		m.initPlayerMarryHeroismObject()
	} else {
		m.playerHeroismObject = NewPlayerMarryHeroismObject(m.p)
		m.playerHeroismObject.FromEntity(heroismEntity)
	}
	//婚帖
	pushWedRecordEntity, err := dao.GetMarryDao().GetPlayerPushWedRecord(m.p.GetId())
	if err != nil {
		return
	}
	if pushWedRecordEntity == nil {
		m.initPlayerPushWedRecordObject()
	} else {
		m.playerPushWedRecordObj = NewPlayerPushWedRecordObject(m.p)
		m.playerPushWedRecordObj.FromEntity(pushWedRecordEntity)
	}

	//获取喜帖
	now := global.GetGame().GetTimeService().Now()
	beginDay, err := timeutils.BeginOfNow(now)
	if err != nil {
		return err
	}
	viewWedCardList, err := dao.GetMarryDao().GetViewWedCardList(m.p.GetId(), beginDay)
	if err != nil {
		return
	}
	for _, viewWedCard := range viewWedCardList {
		m.viewCardList = append(m.viewCardList, viewWedCard.CardId)
	}

	//结婚纪念
	jinianList, err := dao.GetMarryDao().GetPlayerMarryJiNianList(m.p.GetId())
	if err != nil {
		return
	}

	for _, jinianObject := range jinianList {
		itemType := marrytypes.MarryBanquetSubTypeWed(jinianObject.JiNianType)
		jinianItem := NewPlayerMarryJiNianObject(m.Player())
		err = jinianItem.FromEntity(jinianObject)
		if err != nil {
			return err
		}
		m.jinianMap[itemType] = jinianItem
	}

	//定情信物
	dingQing, err := dao.GetMarryDao().GetPlayerMarryDingQingList(m.Player().GetId())
	if err != nil {
		return
	}

	m.dingQing = NewPlayerMarryDingQingObject(m.Player())
	m.dingQing.FromEntity(dingQing)
	if m.dingQing.Id == 0 {
		m.dingQing.Id, _ = idutil.GetId()
		m.dingQing.PlayerId = m.Player().GetId()
		m.dingQing.CreateTime = global.GetGame().GetTimeService().Now()
		m.dingQing.SetModified()
	}

	jinianSj, err := dao.GetMarryDao().GetPlayerJiNianSjInfo(m.Player().GetId())
	if err != nil {
		return
	}
	m.jinianSjObj = NewPlayerMarryJiNianSjObject(m.Player())
	if jinianSj == nil || jinianSj.Id == 0 {
		m.jinianSjObj.Id, _ = idutil.GetId()
		m.jinianSjObj.PlayerId = m.Player().GetId()
		m.jinianSjObj.CreateTime = global.GetGame().GetTimeService().Now()
		m.jinianSjObj.SjGetFlag = false
		m.jinianSjObj.SetModified()
	} else {
		m.jinianSjObj.FromEntity(jinianSj)
	}

	return nil
}

func (m *PlayerMarryDataManager) initPlayerMarryObject() {
	m.playerMarryObject = m.getInitPlayerMarryObject()
	m.playerMarryObject.SetModified()
}

//第一次初始化
// func (m *PlayerMarryDataManager) initPlayerMarryObject() {

// 	m.playerMarryObject = m.getInitPlayerMarryObject()
// 	// pmo.SetModified() 不生成
// }

func (m *PlayerMarryDataManager) getInitPlayerMarryObject() *PlayerMarryObject {
	pmo := NewPlayerMarryObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	pmo.Id, _ = idutil.GetId()
	pmo.PlayerId = m.p.GetId()
	pmo.SpouseId = int64(0)
	pmo.SpouseName = ""
	pmo.Status = marrytypes.MarryStatusTypeUnmarried
	pmo.Ring = -1
	pmo.RingLevel = 0
	pmo.RingNum = 0
	pmo.RingExp = 0
	pmo.TreeLevel = 0
	pmo.TreeNum = 0
	pmo.TreeExp = 0
	pmo.IsProposal = 0
	pmo.WedStatus = marrytypes.MarryWedStatusSelfTypeNo
	pmo.developExp = 0
	pmo.developLevel = 0
	pmo.coupleDevelopLevel = 0
	pmo.CreateTime = now
	return pmo
}

//第一次初始化
func (m *PlayerMarryDataManager) initPlayerMarryHeroismObject() {
	pmho := NewPlayerMarryHeroismObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pmho.Id = id
	//生成id
	pmho.PlayerId = m.p.GetId()
	pmho.Heroism = 0
	pmho.OutOfTime = 0
	pmho.CreateTime = now
	m.playerHeroismObject = pmho
	pmho.SetModified()
}

//第一次初始化
func (m *PlayerMarryDataManager) initPlayerPushWedRecordObject() {
	pmho := NewPlayerPushWedRecordObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pmho.Id = id
	//生成id
	pmho.PlayerId = m.p.GetId()
	pmho.WedId = 0
	pmho.HunCheTime = 0
	pmho.BanquetTime = 0
	pmho.CreateTime = now
	m.playerPushWedRecordObj = pmho
	pmho.SetModified()
}

//清空豪气值
func (m *PlayerMarryDataManager) clearPlayerHeroism(now int64) {
	m.playerHeroismObject.Heroism = 0
	m.playerHeroismObject.OutOfTime = 0
	m.playerHeroismObject.UpdateTime = now
	m.playerHeroismObject.SetModified()
	return
}

func (m *PlayerMarryDataManager) AddHeroism(heroism int32, outOfTime int64) {
	now := global.GetGame().GetTimeService().Now()

	if m.playerHeroismObject.OutOfTime > now {
		m.playerHeroismObject.OutOfTime = outOfTime
		m.playerHeroismObject.Heroism = heroism
	} else {
		m.playerHeroismObject.Heroism += heroism
	}
	m.playerHeroismObject.UpdateTime = now
	m.playerHeroismObject.SetModified()

	return
}

//加载后
func (m *PlayerMarryDataManager) AfterLoad() (err error) {
	m.RefreshHeriosm()
	return nil
}

//心跳
func (m *PlayerMarryDataManager) Heartbeat() {

}

func (m *PlayerMarryDataManager) GetViewCardList() []int64 {
	return m.viewCardList
}

func (m *PlayerMarryDataManager) GetPushWedRecord() *PlayerPushWedRecordObject {
	return m.playerPushWedRecordObj
}

func (m *PlayerMarryDataManager) GetMarryInfo() *PlayerMarryObject {
	return m.playerMarryObject
}

func (m *PlayerMarryDataManager) IsMarry() bool {
	if m.playerMarryObject.SpouseId != 0 {
		return true
	}

	return false
}

func (m *PlayerMarryDataManager) IsTrueMarry() bool {
	if m.playerMarryObject.SpouseId != 0 && m.playerMarryObject.Status == marrytypes.MarryStatusTypeMarried {
		return true
	}

	return false
}

func (m *PlayerMarryDataManager) ToMarryInfo() *marrytypes.MarryInfo {
	info := &marrytypes.MarryInfo{
		SpouseId:   m.playerMarryObject.SpouseId,
		SpouseName: m.playerMarryObject.SpouseName,
		Ring:       int32(m.playerMarryObject.Ring),
		RLevel:     m.playerMarryObject.RingLevel,
		RNum:       m.playerMarryObject.RingNum,
		RProgress:  m.playerMarryObject.RingExp,
		TLevel:     m.playerMarryObject.TreeLevel,
		TNum:       m.playerMarryObject.TreeNum,
		TProgress:  m.playerMarryObject.TreeExp,
		IsProposal: m.playerMarryObject.IsProposal,
		Status:     int32(m.playerMarryObject.Status),
		DLevel:     m.playerMarryObject.developLevel,
		DExp:       m.playerMarryObject.developExp,
		MarryCount: m.playerMarryObject.MarryCount,
	}
	return info
}

func (m *PlayerMarryDataManager) GetClickTime() int64 {
	return m.clickTime
}

func (m *PlayerMarryDataManager) SetClickTime(now int64) {
	m.clickTime = now
}

func (m *PlayerMarryDataManager) GetXinWuTime() int64 {
	return m.xinwuTime
}

func (m *PlayerMarryDataManager) SetXinWuTime(now int64) {
	m.xinwuTime = now
}

func (m *PlayerMarryDataManager) RefreshHeriosm() {
	now := global.GetGame().GetTimeService().Now()
	if m.playerHeroismObject.OutOfTime != 0 &&
		now > m.playerHeroismObject.OutOfTime {
		m.clearPlayerHeroism(now)
	}
	return
}

//TODO:zrc 复用查看的喜帖id
func (m *PlayerMarryDataManager) newPlayerMarryWedCardObject(wedCardId int64) {
	pvwco := NewPlayerViewWedCardObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pvwco.Id = id
	//生成id
	pvwco.PlayerId = m.p.GetId()
	pvwco.CardId = wedCardId
	pvwco.ViewTime = now
	pvwco.CreateTime = now

	pvwco.SetModified()
	return
}

//TODO:zrc 复用查看的喜帖id
//查看喜帖
func (m *PlayerMarryDataManager) ViewWedCard(wedCardId int64) {
	m.newPlayerMarryWedCardObject(wedCardId)
}

//求婚成功
func (m *PlayerMarryDataManager) ProposalMarry(spouseId int64, spouseName string, ringType marrytypes.MarryRingType, isProposal bool) {
	// now := global.GetGame().GetTimeService().Now()
	// if int32(ringType) > int32(m.playerMarryObject.Ring) {
	// 	m.playerMarryObject.Ring = ringType
	// }
	// if m.playerMarryObject.RingLevel == 0 {
	// 	m.playerMarryObject.RingLevel = 1
	// }
	// if m.playerMarryObject.TreeLevel == 0 {
	// 	m.playerMarryObject.TreeLevel = 1
	// }
	// if isProposal {
	// 	m.playerMarryObject.IsProposal = 1
	// } else {
	// 	m.playerMarryObject.IsProposal = 0
	// }
	// m.playerMarryObject.SpouseId = spouseId
	// m.playerMarryObject.SpouseName = spouseName
	// m.playerMarryObject.UpdateTime = now
	// m.playerMarryObject.Status = marrytypes.MarryStatusTypeProposal
	// m.playerMarryObject.SetModified()
	m.proposalMarry(spouseId, spouseName, ringType, isProposal)
	gameevent.Emit(marryeventtypes.EventTypePlayerMarrySpouseChange, m.p, nil)
	gameevent.Emit(marryeventtypes.EventTypePlayerMarryRingChange, m.p, nil)
	return
}

//结婚成功以后，记录包含还原的
func (m *PlayerMarryDataManager) proposalMarry(spouseId int64, spouseName string, ringType marrytypes.MarryRingType, isProposal bool) {
	now := global.GetGame().GetTimeService().Now()
	// _, exists := m.hisMarryMap[spouseId]
	// if exists { //还原历史结婚信息
	// 	hisMarry := m.hisMarryMap[spouseId]
	// 	m.playerMarryObject.SpouseId = hisMarry.SpouseId
	// 	// m.playerMarryObject.Ring = hisMarry.Ring
	// 	// m.playerMarryObject.RingLevel = hisMarry.RingLevel
	// 	m.playerMarryObject.RingNum = hisMarry.RingNum
	// 	m.playerMarryObject.RingExp = hisMarry.RingExp
	// 	m.playerMarryObject.TreeLevel = hisMarry.TreeLevel
	// 	m.playerMarryObject.TreeNum = hisMarry.TreeNum
	// 	m.playerMarryObject.TreeExp = hisMarry.TreeExp
	// 	m.playerMarryObject.IsProposal = hisMarry.IsProposal
	// 	m.playerMarryObject.developLevel = hisMarry.developLevel
	// 	m.playerMarryObject.developExp = hisMarry.developExp
	// 	m.playerMarryObject.coupleDevelopLevel = hisMarry.coupleDevelopLevel
	// }

	//戒指变大则重新开始
	if ringType.BetterThan(m.playerMarryObject.Ring) {
		m.playerMarryObject.Ring = ringType
		m.playerMarryObject.RingLevel = 1
	}
	// if int32(ringType) > int32(m.playerMarryObject.Ring) { //戒指变大则重新开始
	// 	m.playerMarryObject.Ring = ringType
	// 	m.playerMarryObject.RingLevel = 1
	// }
	if isProposal {
		m.playerMarryObject.IsProposal = 1
	} else {
		m.playerMarryObject.IsProposal = 0
	}
	//等级修正
	// if m.playerMarryObject.RingLevel == 0 {
	// 	m.playerMarryObject.RingLevel = 1
	// }
	if m.playerMarryObject.TreeLevel == 0 {
		m.playerMarryObject.TreeLevel = 1
	}
	m.playerMarryObject.SpouseId = spouseId
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.Status = marrytypes.MarryStatusTypeProposal
	m.playerMarryObject.SpouseName = spouseName
	// m.playerMarryObject.MarryCount += 1
	m.playerMarryObject.SetModified()
}

//获取求婚的婚戒类型
func (m *PlayerMarryDataManager) GetProposalRingType(playerId int64) (ringType marrytypes.MarryRingType) {
	ringType = m.ringTypeMap[playerId]
	return
}

//被求婚
func (m *PlayerMarryDataManager) Proposaled(playerId int64, ringType marrytypes.MarryRingType) {
	now := global.GetGame().GetTimeService().Now()
	m.proposalMap[playerId] = now
	m.ringTypeMap[playerId] = ringType
}

//决策后移除求婚
func (m *PlayerMarryDataManager) RemoveProposaled(playerId int64) {
	delete(m.proposalMap, playerId)
	delete(m.ringTypeMap, playerId)
}

//求婚者id是否存在
func (m *PlayerMarryDataManager) IfProposalDealIdExist(dealId int64) bool {
	_, exist := m.proposalMap[dealId]
	if !exist {
		return false
	}
	return true
}

//离婚
func (m *PlayerMarryDataManager) Divorce() {
	// now := global.GetGame().GetTimeService().Now()
	// m.playerMarryObject.SpouseId = 0
	// m.playerMarryObject.Status = marrytypes.MarryStatusTypeDivorce
	// m.playerMarryObject.UpdateTime = now
	// m.receiveDivorce = false
	// m.playerMarryObject.SetModified()
	// //表白系统
	// m.playerMarryObject.coupleDevelopLevel = 0
	// m.playerMarryObject.SetModified()
	m.divorce()
	gameevent.Emit(marryeventtypes.EventTypePlayerMarrySpouseChange, m.p, nil)
	gameevent.Emit(marryeventtypes.EventTypePlayerMarryRingChange, m.p, nil)
	return
}

//离婚
func (m *PlayerMarryDataManager) divorce() {
	now := global.GetGame().GetTimeService().Now()
	//历史中是否已经存在
	// var hisMarry *PlayerMarryLogObject
	// _, exists := m.hisMarryMap[m.playerMarryObject.SpouseId]
	// if exists {
	// 	hisMarry = m.hisMarryMap[m.playerMarryObject.SpouseId]
	// } else {
	// 	hisMarry = NewPlayerMarryLogObject(m.Player())
	// 	hisMarry.Id, _ = idutil.GetId()
	// 	hisMarry.CreateTime = now
	// 	m.hisMarryMap[m.playerMarryObject.SpouseId] = hisMarry
	// }

	// hisMarry.PlayerId = m.playerMarryObject.PlayerId
	// hisMarry.SpouseId = m.playerMarryObject.SpouseId
	// hisMarry.SpouseName = m.playerMarryObject.SpouseName
	// hisMarry.Status = m.playerMarryObject.Status
	// hisMarry.Ring = m.playerMarryObject.Ring
	// hisMarry.RingLevel = m.playerMarryObject.RingLevel
	// hisMarry.RingNum = m.playerMarryObject.RingNum
	// hisMarry.RingExp = m.playerMarryObject.RingExp
	// hisMarry.TreeLevel = m.playerMarryObject.TreeLevel
	// hisMarry.TreeNum = m.playerMarryObject.TreeNum
	// hisMarry.TreeExp = m.playerMarryObject.TreeExp
	// hisMarry.IsProposal = m.playerMarryObject.IsProposal
	// hisMarry.WedStatus = m.playerMarryObject.WedStatus
	// hisMarry.developLevel = m.playerMarryObject.developLevel
	// hisMarry.developExp = m.playerMarryObject.developExp
	// hisMarry.coupleDevelopLevel = m.playerMarryObject.coupleDevelopLevel
	// hisMarry.UpdateTime = now
	// hisMarry.SetModified()

	m.playerMarryObject.SpouseId = 0
	m.playerMarryObject.Status = marrytypes.MarryStatusTypeDivorce
	m.playerMarryObject.UpdateTime = now
	m.receiveDivorce = false
	//表白系统
	// m.playerMarryObject.coupleDevelopLevel = 0
	//清除掉其他的属性信息
	// m.playerMarryObject.Ring = -1
	// m.playerMarryObject.RingLevel = 0
	// m.playerMarryObject.TreeLevel = 0
	// m.playerMarryObject.TreeExp = 0
	// m.playerMarryObject.developLevel = 0

	m.playerMarryObject.SetModified()
}

func (m *PlayerMarryDataManager) IsReceiveDivorce() bool {
	return m.receiveDivorce
}

//收到离婚消息
func (m *PlayerMarryDataManager) ReceiveDivorce() {
	m.receiveDivorce = true
}

//婚戒培养
func (m *PlayerMarryDataManager) RingFeed(pro int32, sucess bool) {
	if pro < 0 {
		return
	}
	if sucess {
		ringType := m.playerMarryObject.Ring
		nextRingLevel := m.playerMarryObject.RingLevel + 1
		ringTemplate := marrytemplate.GetMarryTemplateService().GetMarryRingTemplate(ringType, nextRingLevel)
		if ringTemplate == nil {
			return
		}
		m.playerMarryObject.RingLevel += 1
		m.playerMarryObject.RingNum = 0
		m.playerMarryObject.RingExp = pro
		gameevent.Emit(marryeventtypes.EventTypeMarryRingFeedUpgrade, m.p, m.playerMarryObject.RingLevel)
	} else {
		m.playerMarryObject.RingNum += 1
		m.playerMarryObject.RingExp += pro
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
	return
}

//爱情树培养
func (m *PlayerMarryDataManager) TreeFeed(pro int32, sucess bool) {
	if pro < 0 {
		return
	}
	if sucess {
		nextTreeLevel := m.playerMarryObject.TreeLevel + 1
		treeTemplate := marrytemplate.GetMarryTemplateService().GetMarryLoveTreeTemplate(nextTreeLevel)
		if treeTemplate == nil {
			return
		}
		m.playerMarryObject.TreeLevel += 1
		m.playerMarryObject.TreeNum = 0
		m.playerMarryObject.TreeExp = pro
	} else {
		m.playerMarryObject.TreeNum += 1
		m.playerMarryObject.TreeExp += pro
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
	return
}

//被动替换
func (m *PlayerMarryDataManager) RingReplacedBySpouse(ringType marrytypes.MarryRingType) {
	// if ringType <= m.playerMarryObject.Ring {
	// 	return
	// }
	if !ringType.BetterThan(m.playerMarryObject.Ring) {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerMarryObject.Ring = ringType
	if m.playerMarryObject.RingLevel <= 0 {
		m.playerMarryObject.RingLevel = 1
	}
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
	gameevent.Emit(marryeventtypes.EventTypePlayerMarryRingChange, m.p, nil)
	return
}

//婚戒替换
func (m *PlayerMarryDataManager) RingReplace(ringType marrytypes.MarryRingType) {
	m.RingReplacedBySpouse(ringType)
	if m.playerMarryObject.Status != marrytypes.MarryStatusTypeDivorce {
		//发送事件
		eventData := marryeventtypes.CreateMarryRingReplaceEventData(ringType, m.playerMarryObject.SpouseId)
		gameevent.Emit(marryeventtypes.EventTypeMarryRingReplace, m.p, eventData)
	}
	return
}

//预定婚礼
func (m *PlayerMarryDataManager) DueToWedding() {
	if m.playerMarryObject.Status != marrytypes.MarryStatusTypeProposal {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerMarryObject.Status = marrytypes.MarryStatusTypeEngagement
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
	gameevent.Emit(marryeventtypes.EventTypePlayerMarrySpouseChange, m.p, nil)
	return
}

//举办婚礼
func (m *PlayerMarryDataManager) Holdwedding(isCruise bool) {
	if m.playerMarryObject.Status != marrytypes.MarryStatusTypeEngagement && m.playerMarryObject.Status != marrytypes.MarryStatusTypeMarried {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerMarryObject.Status = marrytypes.MarryStatusTypeMarried
	if isCruise {
		m.playerMarryObject.WedStatus = marrytypes.MarryWedStatusSelfTypeCruise
	} else {
		m.playerMarryObject.WedStatus = marrytypes.MarryWedStatusSelfTypeBanquet
	}
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
	gameevent.Emit(marryeventtypes.EventTypePlayerMarrySpouseChange, m.p, nil)
	gameevent.Emit(marryeventtypes.EventTypePlayerMarryWedStatusChange, m.p, nil)
	return
}

//举办婚礼
func (m *PlayerMarryDataManager) Marry() {
	if m.playerMarryObject.Status != marrytypes.MarryStatusTypeEngagement {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerMarryObject.Status = marrytypes.MarryStatusTypeMarried
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
	gameevent.Emit(marryeventtypes.EventTypePlayerMarrySpouseChange, m.p, nil)
	return
}

//同步玩家个人婚宴状态
func (m *PlayerMarryDataManager) SynchronousWedStatus(wedStatus marrytypes.MarryWedStatusSelfType) {
	now := global.GetGame().GetTimeService().Now()
	if m.playerMarryObject.WedStatus == wedStatus {
		return
	}
	m.playerMarryObject.WedStatus = wedStatus
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
	gameevent.Emit(marryeventtypes.EventTypePlayerMarryWedStatusChange, m.p, nil)
}

//巡游结束
func (m *PlayerMarryDataManager) EndCruise() {
	now := global.GetGame().GetTimeService().Now()
	if m.playerMarryObject.WedStatus != marrytypes.MarryWedStatusSelfTypeCruise {
		return
	}
	m.playerMarryObject.WedStatus = marrytypes.MarryWedStatusSelfTypeBanquet
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
	gameevent.Emit(marryeventtypes.EventTypePlayerMarryWedStatusChange, m.p, nil)
}

//婚宴结束
func (m *PlayerMarryDataManager) EndWedding() {
	now := global.GetGame().GetTimeService().Now()
	if m.playerMarryObject.WedStatus == marrytypes.MarryWedStatusSelfTypeNo {
		return
	}
	m.playerMarryObject.WedStatus = marrytypes.MarryWedStatusSelfTypeNo
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
	gameevent.Emit(marryeventtypes.EventTypePlayerMarryWedStatusChange, m.p, nil)
}

//婚期取消
func (m *PlayerMarryDataManager) CancleWedding() {
	if m.playerMarryObject.Status != marrytypes.MarryStatusTypeEngagement {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerMarryObject.Status = marrytypes.MarryStatusTypeProposal
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
	gameevent.Emit(marryeventtypes.EventTypePlayerMarrySpouseChange, m.p, nil)
	return
}

//婚车帖推送记录
func (m *PlayerMarryDataManager) PushWedRecordHunChe(wedId int64) {
	now := global.GetGame().GetTimeService().Now()
	m.playerPushWedRecordObj.WedId = wedId
	m.playerPushWedRecordObj.HunCheTime = now
	m.playerPushWedRecordObj.BanquetTime = 0
	m.playerPushWedRecordObj.UpdateTime = now
	m.playerPushWedRecordObj.SetModified()
}

//酒席帖推送记录
func (m *PlayerMarryDataManager) PushWedRecordBanquet(wedId int64) {
	now := global.GetGame().GetTimeService().Now()
	m.playerPushWedRecordObj.WedId = wedId
	m.playerPushWedRecordObj.HunCheTime = 0
	m.playerPushWedRecordObj.BanquetTime = now
	m.playerPushWedRecordObj.UpdateTime = now
	m.playerPushWedRecordObj.SetModified()
}

//展示用 特殊处理
func (m *PlayerMarryDataManager) GetRingType() int32 {
	switch m.playerMarryObject.Status {
	case marrytypes.MarryStatusTypeMarried,
		marrytypes.MarryStatusTypeProposal,
		marrytypes.MarryStatusTypeEngagement:
		return int32(m.playerMarryObject.Ring) + 1
	}

	return 0
}

func (m *PlayerMarryDataManager) GetRingLevel() int32 {
	return m.playerMarryObject.RingLevel
}

func (m *PlayerMarryDataManager) GetWedStatus() marrytypes.MarryWedStatusSelfType {
	return m.playerMarryObject.WedStatus
}

func (m *PlayerMarryDataManager) GetSpouseName() string {
	switch m.playerMarryObject.Status {
	case marrytypes.MarryStatusTypeMarried:
		{
			spouseName := m.playerMarryObject.SpouseName
			if m.p.GetSex() == types.SexTypeMan {
				return spouseName + "的丈夫"
			} else {
				return spouseName + "的妻子"
			}
		}
	case marrytypes.MarryStatusTypeProposal,
		marrytypes.MarryStatusTypeEngagement:
		{
			spouseName := m.playerMarryObject.SpouseName
			if m.p.GetSex() == types.SexTypeMan {
				return spouseName + "的未婚夫"
			} else {
				return spouseName + "的未婚妻"
			}
		}
	}
	return ""
}

func (m *PlayerMarryDataManager) GetSpouseId() int64 {
	return m.playerMarryObject.SpouseId
}

func (m *PlayerMarryDataManager) SpouseNameChanged(name string) {
	now := global.GetGame().GetTimeService().Now()
	m.playerMarryObject.SpouseName = name
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
	gameevent.Emit(marryeventtypes.EventTypePlayerMarrySpouseChange, m.p, nil)
}

func (m *PlayerMarryDataManager) MarryPreWedFlag(preWedFlag bool) {
	m.preWedFlag = preWedFlag
}

func (m *PlayerMarryDataManager) GetMarryPreWedFlag() bool {
	return m.preWedFlag
}

func (m *PlayerMarryDataManager) GetMarryDevelopLevel() int32 {
	return m.playerMarryObject.developLevel
}

func (m *PlayerMarryDataManager) GetCoupleMarryDevelopLevel() int32 {
	return m.playerMarryObject.coupleDevelopLevel
}

func (m *PlayerMarryDataManager) UpdateCoupleMarryDevelopLevel(developLevel int32) {
	if developLevel == m.playerMarryObject.coupleDevelopLevel {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerMarryObject.coupleDevelopLevel = developLevel
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
}

func (m *PlayerMarryDataManager) AddDevelopExp(exp int32) {
	if exp < 0 {
		panic(fmt.Errorf("经验值不能小于0，exp:%d", exp))
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerMarryObject.developExp += exp
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
	gameevent.Emit(marryeventtypes.EventTypePlayerMarryDevelopExpAdd, m.p, exp)
}

func (m *PlayerMarryDataManager) GetMarryDevelopExp() int32 {
	return m.playerMarryObject.developExp
}

func (m *PlayerMarryDataManager) MarryDevelopUpLevel() (flag bool) {
	nextLevel := m.playerMarryObject.developLevel + 1
	nextDevelopTemp := marrytemplate.GetMarryTemplateService().GetMarryDeveopTemplate(nextLevel)
	if nextDevelopTemp == nil {
		return
	}

	if m.playerMarryObject.developExp < nextDevelopTemp.Experience {
		return
	}

	beforeDevelopLevel := m.playerMarryObject.developLevel
	beforeDevelopExp := m.playerMarryObject.developExp
	now := global.GetGame().GetTimeService().Now()
	m.playerMarryObject.developLevel = nextLevel
	m.playerMarryObject.developExp -= nextDevelopTemp.Experience
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()

	gameevent.Emit(marryeventtypes.EventTypePlayerMarryDevelopLevelChanged, m.p, m.playerMarryObject.developLevel)

	// reason := commonlog.MarryLogReasonDevelop
	// logEventData := marryeventtypes.CreatePlayerDevelopLevelLogEventData(beforeDevelopLevel, nextLevel, reason, reason.String())
	// gameevent.Emit(marryeventtypes.EventTypePlayerMarryDevelopLog, m.p, logEventData)
	reason := commonlog.MarryLogReasonDevelopExpToLevel
	reasonText := fmt.Sprintf(reason.String(), nextLevel, beforeDevelopLevel)
	logEventData := marryeventtypes.CreatePlayerDevelopExpLogEventData(beforeDevelopExp, m.playerMarryObject.developExp, reason, reasonText)
	gameevent.Emit(marryeventtypes.EventTypePlayerMarryDevelopExpLog, m.p, logEventData)

	flag = true
	return
}

func (m *PlayerMarryDataManager) GetPlayerMarryJinNianInfo(marryType marrytypes.MarryBanquetSubTypeWed) *PlayerMarryJiNianObject {
	return m.jinianMap[marryType]
}

func (m *PlayerMarryDataManager) IsCanAddJiNian(marryType marrytypes.MarryBanquetSubTypeWed, prizeCount int32) bool {
	_, exists := m.jinianMap[marryType]
	if !exists {
		if prizeCount == 1 {
			return true
		}
		return false
	}
	if m.jinianMap[marryType].SendFlag == 1 {
		return false
	}
	if m.jinianMap[marryType].JiNianCount+1 == prizeCount {
		return true
	}
	return false

}

func (m *PlayerMarryDataManager) UpdateMarryJiNianCount(marryType marrytypes.MarryBanquetSubTypeWed, prizeCount int32) {
	mcInfo, exists := m.jinianMap[marryType]

	now := global.GetGame().GetTimeService().Now()
	if !exists {
		marryCountInfo := NewPlayerMarryJiNianObject(m.Player())
		marryCountInfo.Id, _ = idutil.GetId()
		marryCountInfo.PlayerId = m.Player().GetId()
		marryCountInfo.JiNianCount = 0
		marryCountInfo.SendFlag = 0
		marryCountInfo.JiNianType = marryType
		marryCountInfo.CreateTime = now
		m.jinianMap[marryType] = marryCountInfo
		mcInfo = marryCountInfo
	}
	if mcInfo.SendFlag == 0 && mcInfo.JiNianCount+1 == prizeCount {
		mcInfo.SendFlag = 1
	}
	if mcInfo.JiNianCount < prizeCount {
		mcInfo.JiNianCount += 1
	}
	mcInfo.SetModified()
}

func (m *PlayerMarryDataManager) GetJiNianMap() map[marrytypes.MarryBanquetSubTypeWed]*PlayerMarryJiNianObject {
	return m.jinianMap
}

//获取所有定情信息
func (m *PlayerMarryDataManager) GetAllDingQing() *PlayerMarryDingQingObject {
	return m.dingQing
}

func (m *PlayerMarryDataManager) GetAllDingQingMap() map[int32]map[int32]int32 {
	newMap := make(map[int32]map[int32]int32)
	for key, value := range m.dingQing.SuitMap {
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

func (m *PlayerMarryDataManager) ExistsDingQing(suitId int32, posId int32) bool {
	_, exists := m.dingQing.SuitMap[suitId]
	if exists {
		_, exists := m.dingQing.SuitMap[suitId][posId]
		return exists
	}
	return false
}

func (m *PlayerMarryDataManager) AddDingQing(suitId int32, posId int32) {
	exists := m.ExistsDingQing(suitId, posId)
	if exists {
		return
	}
	_, exists = m.dingQing.SuitMap[suitId]
	if !exists {
		m.dingQing.SuitMap[suitId] = make(map[int32]int32)
	}
	m.dingQing.SuitMap[suitId][posId] = 0
	m.dingQing.SetModified()

	gameevent.Emit(marryeventtypes.EventTypeDingQingTokenActivite, m.p, nil)
}

//设置伴侣定情信物
func (m *PlayerMarryDataManager) UpdateSpouseSuit(spouseMap map[int32]map[int32]int32) {
	m.spouseSuitMap = spouseMap
	gameevent.Emit(marryeventtypes.EventTypeDingQingTokenActivite, m.p, nil)
}

//获取伴侣定情信物
func (m *PlayerMarryDataManager) GetSpouseSuit() map[int32]map[int32]int32 {
	return m.spouseSuitMap
}

//增加伴侣的定情信物
func (m *PlayerMarryDataManager) AddSpouseSuit(suitId int32, posId int32) {
	_, exists := m.spouseSuitMap[suitId]
	if !exists {
		m.spouseSuitMap[suitId] = make(map[int32]int32)
	}
	_, exists = m.spouseSuitMap[suitId][posId]
	if exists {
		return
	}
	m.spouseSuitMap[suitId][posId] = 0

	gameevent.Emit(marryeventtypes.EventTypeDingQingTokenActivite, m.p, nil)
}

//是否已经领取过时装
func (m *PlayerMarryDataManager) IfCanGetJiNianShiZhuang() bool {
	if m.jinianSjObj.SjGetFlag {
		return false
	}
	allJiNianTemplate := marrytemplate.GetMarryTemplateService().GetAllMarryJiNianTemplate()
	for key, value := range allJiNianTemplate {
		_, exists := m.jinianMap[key]
		if !exists {
			return false
		}
		if value.NeedNum > m.jinianMap[key].JiNianCount { //未达到赠送标准
			return false
		}
	}
	return true
}

func (m *PlayerMarryDataManager) AddMarryCount() {
	m.playerMarryObject.MarryCount += 1
	m.playerMarryObject.UpdateTime = global.GetGame().GetTimeService().Now()
	m.playerMarryObject.SetModified()
}

func (m *PlayerMarryDataManager) GetMarryCount() int32 {
	return m.playerMarryObject.MarryCount
}

func (m *PlayerMarryDataManager) GetJiNianShiZhuang() {
	if m.jinianSjObj.SjGetFlag {
		return
	}
	m.jinianSjObj.SjGetFlag = true
	m.jinianSjObj.UpdateTime = global.GetGame().GetTimeService().Now()
	m.jinianSjObj.SetModified()
}

//仅gm 使用
func (m *PlayerMarryDataManager) GmMarryRingLevel(level int32) {
	ringType := m.playerMarryObject.Ring
	ringTemplate := marrytemplate.GetMarryTemplateService().GetMarryRingTemplate(ringType, level)
	if ringTemplate == nil {
		return
	}
	m.playerMarryObject.RingLevel = level
	m.playerMarryObject.RingNum = 0
	m.playerMarryObject.RingExp = 0
	gameevent.Emit(marryeventtypes.EventTypeMarryRingFeedUpgrade, m.p, m.playerMarryObject.RingLevel)

	now := global.GetGame().GetTimeService().Now()
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
}

//仅gm 使用
func (m *PlayerMarryDataManager) GmMarryTreeLevel(level int32) {

	treeTemplate := marrytemplate.GetMarryTemplateService().GetMarryLoveTreeTemplate(level)
	if treeTemplate == nil {
		return
	}
	m.playerMarryObject.TreeLevel += 1
	m.playerMarryObject.TreeNum = 0
	m.playerMarryObject.TreeExp = 0

	now := global.GetGame().GetTimeService().Now()
	m.playerMarryObject.UpdateTime = now
	m.playerMarryObject.SetModified()
}

func CreatePlayerMarryDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerMarryDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerMarryDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerMarryDataManager))
}
