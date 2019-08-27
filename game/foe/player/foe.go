package player

import (
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/foe/dao"
	friendtemplate "fgame/fgame/game/friend/template"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
)

//玩家仇人管理器
type PlayerFoeDataManager struct {
	p player.Player
	//玩家仇人列表
	foeMap map[int64]*PlayerFoeObject
	//仇人反馈保护
	foeProtect *PlayerFoeProtectObject
	// 仇人反馈列表
	forFeedbackList []*PlayerFoeFeedbackObject
}

func (m *PlayerFoeDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerFoeDataManager) Load() (err error) {
	err = m.loadFoeObj()
	if err != nil {
		return
	}

	if err = m.loadFoeFeedbackObj(); err != nil {
		return
	}

	if err = m.loadFoeProtectObj(); err != nil {
		return
	}

	return
}

//加载玩家仇人信息
func (m *PlayerFoeDataManager) loadFoeObj() (err error) {
	m.foeMap = make(map[int64]*PlayerFoeObject)
	foeList, err := dao.GetFoeDao().GetFoeList(m.p.GetId())
	if err != nil {
		return
	}
	for _, foeObj := range foeList {
		pfo := newPlayerFoeObject(m.p)
		pfo.FromEntity(foeObj)
		m.foeMap[pfo.AttackId] = pfo
	}
	return
}

//加载玩家仇人反馈
func (m *PlayerFoeDataManager) loadFoeFeedbackObj() (err error) {
	entityList, err := dao.GetFoeDao().GetFoeFeedbackList(m.p.GetId())
	if err != nil {
		return
	}

	for _, entity := range entityList {
		obj := newPlayerFoeFeedbackObject(m.p)
		obj.FromEntity(entity)
		m.addFeedbackObj(obj)
	}
	return
}

//加载玩家仇人反馈保护信息
func (m *PlayerFoeDataManager) loadFoeProtectObj() (err error) {
	entity, err := dao.GetFoeDao().GetFoeProtect(m.p.GetId())
	if err != nil {
		return
	}
	if entity != nil {
		obj := newPlayerFoeProtectObject(m.p)
		obj.FromEntity(entity)
		m.foeProtect = obj
	} else {
		obj := newPlayerFoeProtectObject(m.p)
		id, _ := idutil.GetId()
		now := global.GetGame().GetTimeService().Now()

		obj.id = id
		obj.expireTime = 0
		obj.createTime = now
		obj.SetModified()
		m.foeProtect = obj
	}
	return
}

func (m *PlayerFoeDataManager) initFoeObj(attackId int64) (foeObj *PlayerFoeObject) {
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	foeObj = newPlayerFoeObject(m.p)
	foeObj.Id = id
	foeObj.AttackId = attackId
	foeObj.UpdateTime = now
	foeObj.CreateTime = now
	foeObj.SetModified()
	m.foeMap[attackId] = foeObj
	return
}

//加载后
func (m *PlayerFoeDataManager) AfterLoad() (err error) {

	return nil
}

//心跳
func (m *PlayerFoeDataManager) Heartbeat() {

}

func (m *PlayerFoeDataManager) GetFoeMap() map[int64]*PlayerFoeObject {
	return m.foeMap
}

func (m *PlayerFoeDataManager) AddFoe(attackId int64) (flag bool) {
	now := global.GetGame().GetTimeService().Now()
	foeObj, exist := m.foeMap[attackId]
	if exist {
		foeObj.KillTime = now
		foeObj.UpdateTime = now
		foeObj.SetModified()
		return
	}
	foeLen := int32(len(m.foeMap))
	maxNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFoeLimit)
	if foeLen >= maxNum {
		return
	}
	m.initFoeObj(attackId)
	flag = true
	return
}

func (m *PlayerFoeDataManager) IsFoe(attackId int64) bool {
	_, exist := m.foeMap[attackId]
	if !exist {
		return false
	}
	return true
}

func (m *PlayerFoeDataManager) RemoveFoe(attackId int64) {
	if !m.IsFoe(attackId) {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	foeObj := m.foeMap[attackId]
	foeObj.UpdateTime = now
	foeObj.DeleteTime = now
	foeObj.SetModified()
	delete(m.foeMap, attackId)
}

func (m *PlayerFoeDataManager) IsOnProtected() bool {
	now := global.GetGame().GetTimeService().Now()
	return now < m.foeProtect.expireTime
}

func (m *PlayerFoeDataManager) GetFoeFeedbackProtectExpireTime() int64 {
	return m.foeProtect.expireTime
}

func (m *PlayerFoeDataManager) BuyFoeFeedbackProtect() {
	now := global.GetGame().GetTimeService().Now()
	noticeConstantTemp := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeConstanTemplate()
	expireTime := now + int64(noticeConstantTemp.BaoHuTime)
	m.foeProtect.expireTime = expireTime
	m.foeProtect.updateTime = now
	m.foeProtect.SetModified()
}

func (m *PlayerFoeDataManager) GetFoeFeedbackList() []*PlayerFoeFeedbackObject {
	return m.forFeedbackList
}

func (m *PlayerFoeDataManager) ReadFeedback() {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.forFeedbackList {
		obj.updateTime = now
		obj.deleteTime = now
		obj.SetModified()
	}

	m.forFeedbackList = []*PlayerFoeFeedbackObject{}
}

func (m *PlayerFoeDataManager) AddFoeFeedback(isProtect bool, feedbackName string) {
	obj := newPlayerFoeFeedbackObject(m.p)
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj.id = id
	obj.feedbackName = feedbackName
	if isProtect {
		obj.isProtect = 1
	} else {
		obj.isProtect = 0
	}
	obj.createTime = now
	obj.SetModified()

	m.addFeedbackObj(obj)
}

func (m *PlayerFoeDataManager) addFeedbackObj(obj *PlayerFoeFeedbackObject) {
	m.forFeedbackList = append(m.forFeedbackList, obj)
}

func CreatePlayerFoeDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerFoeDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerFoeDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerFoeDataManager))
}
