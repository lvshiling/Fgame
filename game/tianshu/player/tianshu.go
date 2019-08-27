package player

import (
	commonlog "fgame/fgame/common/log"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/tianshu/dao"
	tianshueventtypes "fgame/fgame/game/tianshu/event/types"
	tianshutypes "fgame/fgame/game/tianshu/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//玩家天书管理器
type PlayerTianShuDataManager struct {
	p          player.Player
	tianshuMap map[tianshutypes.TianShuType]*PlayerTianShuObject
}

func (m *PlayerTianShuDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerTianShuDataManager) Load() (err error) {
	//加载玩家打宝塔信息
	entityList, err := dao.GetTianShuDao().GetTianShuEntityList(m.p.GetId())
	if err != nil {
		return
	}

	for _, entity := range entityList {
		obj := NewPlayerTianShuObject(m.p)
		obj.FromEntity(entity)
		m.addTianshu(obj)
	}
	return nil
}

//加载后
func (m *PlayerTianShuDataManager) AfterLoad() (err error) {
	m.RefreshReceive()
	return
}

//心跳
func (m *PlayerTianShuDataManager) Heartbeat() {
}

//第一次初始化
func (m *PlayerTianShuDataManager) initPlayerTianShuObject(typ tianshutypes.TianShuType) *PlayerTianShuObject {
	o := NewPlayerTianShuObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.typ = typ
	o.level = 1
	o.isReceive = 0
	o.createTime = now
	o.SetModified()

	return o
}

// 激活天书
func (m *PlayerTianShuDataManager) ActivateTianShu(typ tianshutypes.TianShuType) {
	obj := m.getTianShu(typ)
	if obj != nil {
		return
	}
	obj = m.initPlayerTianShuObject(typ)
	m.addTianshu(obj)

	gameevent.Emit(tianshueventtypes.EventTypePlayerTianShuActivate, m.p, obj)
	tsReason := commonlog.TianShuLogReasonActivate
	tsReasonText := fmt.Sprintf(tsReason.String(), typ)
	data := tianshueventtypes.CreatePlayerTianShuLogEventData(0, obj.level, typ, tsReason, tsReasonText)
	gameevent.Emit(tianshueventtypes.EventTypePlayerTianShuLog, m.p, data)
}

// 是否激活天书
func (m *PlayerTianShuDataManager) IsActivateTianShu(typ tianshutypes.TianShuType) bool {
	obj := m.getTianShu(typ)
	if obj == nil {
		return false
	}

	return true
}

// 领取每日奖励
func (m *PlayerTianShuDataManager) ReceiveTianShuGift(typ tianshutypes.TianShuType) {
	obj := m.getTianShu(typ)
	if obj == nil {
		return
	}

	if obj.isReceive == 1 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj.isReceive = 1
	obj.updateTime = now
	obj.SetModified()
}

func (m *PlayerTianShuDataManager) IsReceive(typ tianshutypes.TianShuType) bool {
	obj := m.getTianShu(typ)
	if obj == nil {
		return true
	}

	if obj.isReceive == 1 {
		return true
	}

	return false
}

func (m *PlayerTianShuDataManager) RefreshReceive() {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.tianshuMap {

		isSame, _ := timeutils.IsSameFive(now, obj.updateTime)
		if !isSame {
			obj.isReceive = 0
			obj.updateTime = now
			obj.SetModified()
		}
	}
}

func (m *PlayerTianShuDataManager) UplevelTianShu(typ tianshutypes.TianShuType) {
	obj := m.getTianShu(typ)
	if obj == nil {
		return
	}

	beforLevel := obj.level
	now := global.GetGame().GetTimeService().Now()
	obj.level += 1
	obj.updateTime = now
	obj.SetModified()

	gameevent.Emit(tianshueventtypes.EventTypePlayerTianShuUplevel, m.p, obj)
	tsReason := commonlog.TianShuLogReasonUplevel
	tsReasonText := fmt.Sprintf(tsReason.String(), typ)
	data := tianshueventtypes.CreatePlayerTianShuLogEventData(beforLevel, 1, typ, tsReason, tsReasonText)
	gameevent.Emit(tianshueventtypes.EventTypePlayerTianShuLog, m.p, data)
}

func (m *PlayerTianShuDataManager) GetTianShuLevel(typ tianshutypes.TianShuType) int32 {
	obj := m.getTianShu(typ)
	if obj == nil {
		return 0
	}

	return obj.level
}

func (m *PlayerTianShuDataManager) GetTianShuAll() map[tianshutypes.TianShuType]*PlayerTianShuObject {
	return m.tianshuMap
}

func (m *PlayerTianShuDataManager) getTianShu(typ tianshutypes.TianShuType) *PlayerTianShuObject {
	obj, ok := m.tianshuMap[typ]
	if !ok {
		return nil
	}
	return obj
}

func (m *PlayerTianShuDataManager) addTianshu(obj *PlayerTianShuObject) {
	_, ok := m.tianshuMap[obj.typ]
	if !ok {
		m.tianshuMap[obj.typ] = obj
	}
}

func CreatePlayerTianShuDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerTianShuDataManager{}
	m.p = p
	m.tianshuMap = make(map[tianshutypes.TianShuType]*PlayerTianShuObject)
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerTianShuDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerTianShuDataManager))
}
