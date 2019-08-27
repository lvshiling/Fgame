package player

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	xianzuncardcommon "fgame/fgame/game/xianzuncard/common"
	"fgame/fgame/game/xianzuncard/dao"
	xianzuncardeventtypes "fgame/fgame/game/xianzuncard/event/types"
	xianzuncardtemplate "fgame/fgame/game/xianzuncard/template"
	xianzuncardtypes "fgame/fgame/game/xianzuncard/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

type PlayerXianZunCardDataManager struct {
	pl                player.Player
	xianZunCardObjMap map[xianzuncardtypes.XianZunCardType]*PlayerXianZunCardObject
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

// 加载
func (m *PlayerXianZunCardDataManager) Load() (err error) {
	xianZunEntityList, err := dao.GetXianZunCard().GetPlayerXianZunCardEntityList(m.pl.GetId())
	if err != nil {
		return
	}

	m.xianZunCardObjMap = make(map[xianzuncardtypes.XianZunCardType]*PlayerXianZunCardObject)
	for _, entity := range xianZunEntityList {
		obj := NewXianZunCardObject(m.pl)
		obj.FromEntity(entity)
		m.xianZunCardObjMap[obj.cardType] = obj
	}

	return nil
}

func (m *PlayerXianZunCardDataManager) Player() player.Player {
	return m.pl
}

//加载后
func (m *PlayerXianZunCardDataManager) AfterLoad() (err error) {
	m.heartbeatRunner.AddTask(CreateXianZunCardTask(m.pl))
	return
}

//心跳
func (m *PlayerXianZunCardDataManager) Heartbeat() {
	m.heartbeatRunner.Heartbeat()
}

func (m *PlayerXianZunCardDataManager) GetXianZunCardObjectMap() map[xianzuncardtypes.XianZunCardType]*PlayerXianZunCardObject {
	return m.xianZunCardObjMap
}

func (m *PlayerXianZunCardDataManager) GetXianZunCardTypeList() (infoList []*xianzuncardcommon.XianZunCardInfo) {
	for typ, obj := range m.xianZunCardObjMap {
		if !obj.IsActivite() {
			continue
		}
		info := &xianzuncardcommon.XianZunCardInfo{
			Typ: int32(typ),
		}
		infoList = append(infoList, info)
	}
	return
}

// 是否已经激活
func (m *PlayerXianZunCardDataManager) IsActivite(typ xianzuncardtypes.XianZunCardType) bool {
	obj, ok := m.xianZunCardObjMap[typ]
	if ok {
		return obj.IsActivite()
	}
	return false
}

// 是否已经领取
func (m *PlayerXianZunCardDataManager) IsReceive(typ xianzuncardtypes.XianZunCardType) bool {
	obj, ok := m.xianZunCardObjMap[typ]
	if ok {
		return obj.IsReceive()
	}
	return false
}

// 购买仙尊特权卡成功
func (m *PlayerXianZunCardDataManager) BuySuccess(typ xianzuncardtypes.XianZunCardType) bool {
	obj, ok := m.xianZunCardObjMap[typ]
	if !ok {
		obj = m.initXianZunCardObject(typ)
		m.xianZunCardObjMap[typ] = obj
	}

	now := global.GetGame().GetTimeService().Now()
	obj.isActivite = 1
	obj.activiteTime = now
	obj.receiveTime = now
	obj.updateTime = now
	obj.SetModified()

	event.Emit(xianzuncardeventtypes.EventTypeXianZunCardDataChange, m.pl, nil)
	return true
}

// 领取每日仙尊特权卡奖励成功
func (m *PlayerXianZunCardDataManager) ReceiveSuccess(typ xianzuncardtypes.XianZunCardType) bool {
	obj, ok := m.xianZunCardObjMap[typ]
	if !ok {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	obj.isReceive = 1
	obj.receiveTime = now
	obj.updateTime = now
	obj.SetModified()
	return true
}

func (m *PlayerXianZunCardDataManager) refreshData() {
	now := global.GetGame().GetTimeService().Now()
	for typ, obj := range m.xianZunCardObjMap {
		// 是否激活
		if !obj.IsActivite() {
			continue
		}

		// 判断是否跨天
		diffDay, err := timeutils.DiffDay(now, obj.receiveTime)
		if err != nil {
			continue
		}
		if diffDay != 0 {

			// 前面有一天已经领取过了
			if obj.IsReceive() {
				diffDay -= 1
			}

			// 未领取发邮件补奖励
			if diffDay >= 1 {
				data := xianzuncardeventtypes.CreatePlayerXianZunCardCrossDayEventData(typ, diffDay)
				event.Emit(xianzuncardeventtypes.EventTypeXianZunCardCrossDay, m.pl, data)
			}

			obj.isReceive = 0
			obj.receiveTime = now
			obj.updateTime = now
			obj.SetModified()

		}

		temp := xianzuncardtemplate.GetXianZunCardTemplateService().GetXianZunCardTemplate(typ)
		if temp == nil {
			continue
		}
		duration := temp.Duration
		// 判断是否到期
		if now > obj.activiteTime+duration {
			obj.isActivite = 0
			obj.isReceive = 0
			obj.activiteTime = 0
			obj.receiveTime = 0
			obj.SetModified()

			// 至尊特权卡过期
			event.Emit(xianzuncardeventtypes.EventTypeXianZunCardExpire, m.pl, typ)
			event.Emit(xianzuncardeventtypes.EventTypeXianZunCardDataChange, m.pl, nil)
		}

	}
}

func (m *PlayerXianZunCardDataManager) initXianZunCardObject(typ xianzuncardtypes.XianZunCardType) (obj *PlayerXianZunCardObject) {
	obj = NewXianZunCardObject(m.pl)
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj.id = id
	obj.cardType = typ
	obj.createTime = now
	return
}

func CreatePlayerXianZunCardDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerXianZunCardDataManager{}
	m.pl = p
	m.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerXianZunCardManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerXianZunCardDataManager))
}
