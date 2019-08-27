package player

import (
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/huiyuan/dao"
	huiyuaneventtypes "fgame/fgame/game/huiyuan/event/types"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家会员管理器
type PlayerHuiYuanManager struct {
	p player.Player
	//会员数据
	huiyuanObject *PlayerHuiYuanObject
}

func (m *PlayerHuiYuanManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerHuiYuanManager) Load() (err error) {

	//加载会员数据
	huiyuanEntity, err := dao.GetHuiYuanDao().GetHuiYuanEntity(m.p.GetId())
	if err != nil {
		return
	}

	if huiyuanEntity != nil {
		obj := newPlayerHuiYuanObject(m.p)
		obj.FromEntity(huiyuanEntity)
		m.huiyuanObject = obj
	}

	return nil
}

//加载后
func (m *PlayerHuiYuanManager) AfterLoad() (err error) {
	err = m.RefresHuiYuanRewards()
	if err != nil {
		return
	}
	return
}

//心跳
func (m *PlayerHuiYuanManager) Heartbeat() {
}

//是否临时会员
func (m *PlayerHuiYuanManager) IsHuiYuan(huiyuanType huiyuantypes.HuiYuanType) bool {
	if m.huiyuanObject == nil {
		return false
	}

	if huiyuanType == huiyuantypes.HuiYuanTypePlus {
		if m.huiyuanObject.huiyuanType == huiyuanType {
			return true
		}
	} else {
		now := global.GetGame().GetTimeService().Now()
		if now < m.huiyuanObject.expireTime {
			return true
		}
	}

	return false
}

//购买会员
func (m *PlayerHuiYuanManager) BuyHuiYuan(huiyuanType huiyuantypes.HuiYuanType, expireTime int64) (huiyuan *PlayerHuiYuanObject) {
	now := global.GetGame().GetTimeService().Now()
	if m.huiyuanObject == nil {
		obj := newPlayerHuiYuanObject(m.p)
		id, _ := idutil.GetId()
		obj.id = id
		obj.createTime = now
		obj.level = 1
		m.huiyuanObject = obj
	}

	if huiyuanType == huiyuantypes.HuiYuanTypeInterim {
		expireTime += now
		m.huiyuanObject.interimBuyTime = now
		m.huiyuanObject.expireTime = expireTime
	} else {
		m.huiyuanObject.plusBuyTime = now
	}

	m.huiyuanObject.huiyuanType = huiyuanType
	m.huiyuanObject.updateTime = now
	m.huiyuanObject.SetModified()

	gameevent.Emit(huiyuaneventtypes.EventTypeHuiYuanBuy, m.p, huiyuanType)

	return m.huiyuanObject
}

//是否领取奖励
func (m *PlayerHuiYuanManager) IsReceiveRewards(huiyuanType huiyuantypes.HuiYuanType) bool {
	if m.huiyuanObject == nil {
		return false
	}

	lastReceiveTime := int64(0)
	if huiyuanType == huiyuantypes.HuiYuanTypeInterim {
		lastReceiveTime = m.huiyuanObject.lastInterimReceiveTime
	} else {
		lastReceiveTime = m.huiyuanObject.lastReceiveTime
	}

	now := global.GetGame().GetTimeService().Now()
	diffDay, _ := timeutils.DiffDay(now, lastReceiveTime)
	if diffDay > 0 {
		return false
	}

	return true
}

//领取奖励
func (m *PlayerHuiYuanManager) ReceiveHuiYuanRewards(huiyuanType huiyuantypes.HuiYuanType) {

	now := global.GetGame().GetTimeService().Now()
	if huiyuanType == huiyuantypes.HuiYuanTypeInterim {
		m.huiyuanObject.lastInterimReceiveTime = now
	} else {
		m.huiyuanObject.lastReceiveTime = now
	}
	m.huiyuanObject.updateTime = now
	m.huiyuanObject.SetModified()
	return
}

//刷新会员领奖信息
func (m *PlayerHuiYuanManager) RefresHuiYuanRewards() (err error) {
	if m.huiyuanObject == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	canRewardsTime := now - int64(common.DAY)
	//至尊会员
	if m.huiyuanObject.huiyuanType == huiyuantypes.HuiYuanTypePlus {
		diffDay := int32(0)
		lastTime := m.huiyuanObject.lastReceiveTime
		if lastTime == 0 {
			lastTime = m.huiyuanObject.plusBuyTime
			diffDay, _ = timeutils.DiffDay(canRewardsTime, lastTime)
			diffDay += 1
			if diffDay <= 0 {
				goto Interim
			}
		} else {
			diffDay, _ = timeutils.DiffDay(canRewardsTime, lastTime)
			if diffDay <= 0 {
				goto Interim
			}
		}

		//发事件
		data := huiyuaneventtypes.CreateHuiYuanRewardsEventData(huiyuantypes.HuiYuanTypePlus, diffDay, m.huiyuanObject.plusBuyTime)
		gameevent.Emit(huiyuaneventtypes.EventTypeHuiYuanRewards, m.p, data)

		m.huiyuanObject.lastReceiveTime = canRewardsTime
		m.huiyuanObject.updateTime = now
		m.huiyuanObject.SetModified()
	}
Interim:
	//临时会员
	interimCanRewardsTime := now - int64(common.DAY)
	if m.huiyuanObject.expireTime > 0 {
		if now > m.huiyuanObject.expireTime {
			interimCanRewardsTime = m.huiyuanObject.expireTime
		}

		diffDay := int32(0)
		lastTime := m.huiyuanObject.lastInterimReceiveTime
		if lastTime == 0 {
			lastTime = m.huiyuanObject.interimBuyTime
			diffDay, _ = timeutils.DiffDay(interimCanRewardsTime, lastTime)
			diffDay += 1
			if diffDay <= 0 {
				return
			}
		} else {
			diffDay, _ = timeutils.DiffDay(interimCanRewardsTime, lastTime)
			if diffDay <= 0 {
				return
			}
		}

		//发事件
		data := huiyuaneventtypes.CreateHuiYuanRewardsEventData(huiyuantypes.HuiYuanTypeInterim, diffDay, m.huiyuanObject.interimBuyTime)
		gameevent.Emit(huiyuaneventtypes.EventTypeHuiYuanRewards, m.p, data)

		m.huiyuanObject.lastInterimReceiveTime = interimCanRewardsTime
		m.huiyuanObject.updateTime = now
		m.huiyuanObject.SetModified()
	}

	return
}

// 是否今日购买
func (m *PlayerHuiYuanManager) IsFirstRew(huiyuanType huiyuantypes.HuiYuanType) bool {
	if m.huiyuanObject == nil {
		return true
	}

	now := global.GetGame().GetTimeService().Now()
	buyTime := int64(0)
	if huiyuanType == huiyuantypes.HuiYuanTypePlus {
		buyTime = m.huiyuanObject.plusBuyTime
	}
	if huiyuanType == huiyuantypes.HuiYuanTypeInterim {
		buyTime = m.huiyuanObject.interimBuyTime
	}

	isSame, err := timeutils.IsSameDay(now, buyTime)
	if err != nil {
		return false
	}

	return isSame
}

func (m *PlayerHuiYuanManager) GetHuiYuanExpireTiem() int64 {
	if m.huiyuanObject == nil {
		return 0
	}

	return m.huiyuanObject.expireTime
}

func (m *PlayerHuiYuanManager) GetHuiYuanLevel() int32 {
	if m.huiyuanObject == nil {
		return 0
	}

	return m.huiyuanObject.level
}

func (m *PlayerHuiYuanManager) GetHuiYuanType() huiyuantypes.HuiYuanType {
	if m.huiyuanObject == nil {
		return huiyuantypes.HuiYuanTypeCommon
	}

	if m.huiyuanObject.huiyuanType == huiyuantypes.HuiYuanTypeInterim {
		now := global.GetGame().GetTimeService().Now()
		if now > m.huiyuanObject.expireTime {
			return huiyuantypes.HuiYuanTypeCommon
		}
	}

	return m.huiyuanObject.huiyuanType
}

func createPlayerHuiYuanDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerHuiYuanManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerHuiYuanDataManagerType, player.PlayerDataManagerFactoryFunc(createPlayerHuiYuanDataManager))
}
