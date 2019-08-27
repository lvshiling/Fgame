package player

import (
	"fgame/fgame/game/constant/constant"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/tianmo/dao"
	tianmoeventtypes "fgame/fgame/game/tianmo/event/types"
	tianmotemplate "fgame/fgame/game/tianmo/template"
	tianmotypes "fgame/fgame/game/tianmo/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家天魔管理器
type PlayerTianMoDataManager struct {
	p player.Player
	//玩家天魔对象
	playerTianMoObject *PlayerTianMoObject
}

func (m *PlayerTianMoDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerTianMoDataManager) Load() (err error) {
	//加载玩家天魔信息
	tianmoEntity, err := dao.GetTianMoDao().GetTianMoEntity(m.p.GetId())
	if err != nil {
		return
	}
	if tianmoEntity == nil {
		m.initPlayerTianMoObject()
	} else {
		m.playerTianMoObject = NewPlayerTianMoObject(m.p)
		m.playerTianMoObject.FromEntity(tianmoEntity)
	}

	return nil
}

//第一次初始化
func (m *PlayerTianMoDataManager) initPlayerTianMoObject() {
	obj := NewPlayerTianMoObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.Id = id
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(m.p.GetRole(), m.p.GetSex())
	advanceId := playerCreateTemplate.TianMoTi
	obj.AdvanceId = advanceId
	obj.TianMoDanLevel = 0
	obj.TianMoDanNum = int32(0)
	obj.TianMoDanPro = 0
	obj.TimesNum = int32(0)
	obj.Bless = int32(0)
	obj.BlessTime = int64(0)
	obj.Power = int64(0)
	obj.CreateTime = now
	obj.SetModified()

	m.playerTianMoObject = obj
}

func (m *PlayerTianMoDataManager) refreshBless() {
	now := global.GetGame().GetTimeService().Now()
	number := int32(m.playerTianMoObject.AdvanceId)
	nextNumber := number + 1
	tianMoTemplate := tianmotemplate.GetTianMoTemplateService().GetTianMoNumber(nextNumber)
	if tianMoTemplate == nil {
		return
	}
	if !tianMoTemplate.GetIsClear() {
		return
	}
	lastTime := m.playerTianMoObject.BlessTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return
		}
		if !flag {
			m.playerTianMoObject.Bless = 0
			m.playerTianMoObject.BlessTime = 0
			m.playerTianMoObject.TimesNum = 0
			m.playerTianMoObject.SetModified()
		}
	}
	return
}

//加载后
func (m *PlayerTianMoDataManager) AfterLoad() (err error) {
	m.refreshBless()
	return nil
}

//天魔信息对象
func (m *PlayerTianMoDataManager) GetTianMoInfo() *PlayerTianMoObject {
	m.refreshBless()
	return m.playerTianMoObject
}

func (m *PlayerTianMoDataManager) GetTianMoAdvanced() int32 {
	return int32(m.playerTianMoObject.AdvanceId)
}

func (m *PlayerTianMoDataManager) EatTianMoDan(level int32) {
	if m.playerTianMoObject.TianMoDanLevel == level || level <= 0 {
		return
	}
	tianmoDanTemplate := tianmotemplate.GetTianMoTemplateService().GetTianMoDan(level)
	if tianmoDanTemplate == nil {
		return
	}
	m.playerTianMoObject.TianMoDanLevel = level
	now := global.GetGame().GetTimeService().Now()
	m.playerTianMoObject.UpdateTime = now
	m.playerTianMoObject.SetModified()
	return
}

//心跳
func (m *PlayerTianMoDataManager) Heartbeat() {

}

//进阶
func (m *PlayerTianMoDataManager) TianMoAdvanced(pro, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		tianmoTemplate := tianmotemplate.GetTianMoTemplateService().GetTianMoNumber(m.playerTianMoObject.AdvanceId + 1)
		if tianmoTemplate == nil {
			return
		}
		m.playerTianMoObject.AdvanceId += 1
		m.playerTianMoObject.TimesNum = 0
		m.playerTianMoObject.Bless = 0
		m.playerTianMoObject.BlessTime = 0
		gameevent.Emit(tianmoeventtypes.EventTypeTianMoAdvanced, m.p, int32(m.playerTianMoObject.AdvanceId))
	} else {
		m.playerTianMoObject.TimesNum += addTimes
		if m.playerTianMoObject.Bless == 0 {
			m.playerTianMoObject.BlessTime = now
		}
		m.playerTianMoObject.Bless += pro
	}
	m.playerTianMoObject.UpdateTime = now
	m.playerTianMoObject.SetModified()
	return
}

//直升券进阶
func (m *PlayerTianMoDataManager) TianMoAdvancedTicket(addAdvancedNum int32) {
	if addAdvancedNum <= 0 {
		return
	}

	canAddNum := int32(0)
	nextAdvancedId := m.playerTianMoObject.AdvanceId + 1
	for addAdvancedNum > 0 {
		tianmoTemplate := tianmotemplate.GetTianMoTemplateService().GetTianMoNumber(nextAdvancedId)
		if tianmoTemplate == nil {
			return
		}
		canAddNum += 1
		nextAdvancedId += 1
		addAdvancedNum -= 1
	}

	if canAddNum == 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerTianMoObject.AdvanceId += canAddNum
	m.playerTianMoObject.TimesNum = 0
	m.playerTianMoObject.Bless = 0
	m.playerTianMoObject.BlessTime = 0
	m.playerTianMoObject.UpdateTime = now
	m.playerTianMoObject.SetModified()
	gameevent.Emit(tianmoeventtypes.EventTypeTianMoAdvanced, m.p, m.playerTianMoObject.AdvanceId)
	return
}

//天魔战斗力
func (m *PlayerTianMoDataManager) TianMoPower(power int64) {
	if power <= 0 {
		return
	}
	curPower := m.playerTianMoObject.Power
	if curPower == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerTianMoObject.Power = power
	m.playerTianMoObject.UpdateTime = now
	m.playerTianMoObject.SetModified()
	gameevent.Emit(tianmoeventtypes.EventTypeTianMoPowerChanged, m.p, power)
	return
}

func (m *PlayerTianMoDataManager) ToTianMoInfo() *tianmotypes.TianMoInfo {
	bodyShieldInfo := &tianmotypes.TianMoInfo{
		AdvancedId:     int32(m.playerTianMoObject.AdvanceId),
		TianMoDanLevel: m.playerTianMoObject.TianMoDanLevel,
		TianMoDanPro:   m.playerTianMoObject.TianMoDanPro,
	}
	return bodyShieldInfo
}

func (m *PlayerTianMoDataManager) IfFullAdvanced() (flag bool) {
	if m.playerTianMoObject.AdvanceId == 0 {
		return
	}
	tianMoTemplate := tianmotemplate.GetTianMoTemplateService().GetTianMoNumber(m.playerTianMoObject.AdvanceId)
	if tianMoTemplate == nil {
		return
	}
	if tianMoTemplate.NextId == 0 {
		return true
	}
	return
}

//天魔体充值数
func (m *PlayerTianMoDataManager) AddChargeNum(chargeNum int32) {
	if chargeNum <= 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerTianMoObject.ChargeVal += int64(chargeNum)
	m.playerTianMoObject.UpdateTime = now
	m.playerTianMoObject.SetModified()
	return
}

//仅gm使用 天魔进阶
func (m *PlayerTianMoDataManager) GmSetTianMoAdvanced(advancedId int32) {
	m.playerTianMoObject.AdvanceId = advancedId
	m.playerTianMoObject.TimesNum = int32(0)
	m.playerTianMoObject.Bless = int32(0)
	m.playerTianMoObject.BlessTime = int64(0)
	now := global.GetGame().GetTimeService().Now()
	m.playerTianMoObject.UpdateTime = now
	m.playerTianMoObject.SetModified()

	gameevent.Emit(tianmoeventtypes.EventTypeTianMoAdvanced, m.p, m.playerTianMoObject.AdvanceId)
	return
}

//仅gm使用 天魔食丹等级
func (m *PlayerTianMoDataManager) GmSetTianMoTianMoDanLevel(level int32) {
	m.playerTianMoObject.TianMoDanLevel = level
	m.playerTianMoObject.TianMoDanNum = 0
	m.playerTianMoObject.TianMoDanPro = 0

	now := global.GetGame().GetTimeService().Now()
	m.playerTianMoObject.UpdateTime = now
	m.playerTianMoObject.SetModified()
}

func CreatePlayerTianMoDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerTianMoDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerTianMoDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerTianMoDataManager))
}
