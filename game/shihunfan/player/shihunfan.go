package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	shihunfancommon "fgame/fgame/game/shihunfan/common"
	"fgame/fgame/game/shihunfan/dao"
	shihunfaneventtypes "fgame/fgame/game/shihunfan/event/types"
	shihunfantemplate "fgame/fgame/game/shihunfan/template"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家噬魂幡管理器
type PlayerShiHunFanDataManager struct {
	p player.Player
	//玩家噬魂幡对象
	playerShiHunFanObject *PlayerShiHunFanObject
}

func (m *PlayerShiHunFanDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerShiHunFanDataManager) Load() (err error) {
	//加载玩家噬魂幡信息
	shihunfanEntity, err := dao.GetShiHunFanDao().GetShiHunFanEntity(m.p.GetId())
	if err != nil {
		return
	}
	if shihunfanEntity == nil {
		m.initPlayerShiHunFanObject()
	} else {
		m.playerShiHunFanObject = NewPlayerShiHunFanObject(m.p)
		m.playerShiHunFanObject.FromEntity(shihunfanEntity)
	}

	return nil
}

//第一次初始化
func (m *PlayerShiHunFanDataManager) initPlayerShiHunFanObject() {
	obj := NewPlayerShiHunFanObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.Id = id
	obj.CreateTime = now
	obj.SetModified()

	m.playerShiHunFanObject = obj
}

func (m *PlayerShiHunFanDataManager) refreshBless() {
	now := global.GetGame().GetTimeService().Now()
	number := int32(m.playerShiHunFanObject.AdvanceId)
	nextNumber := number + 1
	nextTemplate := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanNumber(nextNumber)
	if nextTemplate == nil {
		return
	}
	if !nextTemplate.GetIsClear() {
		return
	}
	lastTime := m.playerShiHunFanObject.BlessTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return
		}
		if !flag {
			m.playerShiHunFanObject.Bless = 0
			m.playerShiHunFanObject.BlessTime = 0
			m.playerShiHunFanObject.TimesNum = 0
			m.playerShiHunFanObject.SetModified()
		}
	}
	return
}

//加载后
func (m *PlayerShiHunFanDataManager) AfterLoad() (err error) {
	m.refreshBless()
	return nil
}

//噬魂幡信息对象
func (m *PlayerShiHunFanDataManager) GetShiHunFanInfo() *PlayerShiHunFanObject {
	m.refreshBless()
	return m.playerShiHunFanObject
}

func (m *PlayerShiHunFanDataManager) GetShiHunFanAdvanced() int32 {
	return int32(m.playerShiHunFanObject.AdvanceId)
}

func (m *PlayerShiHunFanDataManager) EatShiHunFanDan(level int32) {
	if m.playerShiHunFanObject.DanLevel == level || level <= 0 {
		return
	}
	danTemplate := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanDan(level)
	if danTemplate == nil {
		return
	}
	m.playerShiHunFanObject.DanLevel = level
	now := global.GetGame().GetTimeService().Now()
	m.playerShiHunFanObject.UpdateTime = now
	m.playerShiHunFanObject.SetModified()
	return
}

//心跳
func (m *PlayerShiHunFanDataManager) Heartbeat() {

}

//进阶
func (m *PlayerShiHunFanDataManager) ShiHunFanAdvanced(pro, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		shihunfanTemplate := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanNumber(int32(m.playerShiHunFanObject.AdvanceId + 1))
		if shihunfanTemplate == nil {
			return
		}
		m.playerShiHunFanObject.AdvanceId += 1
		m.playerShiHunFanObject.TimesNum = 0
		m.playerShiHunFanObject.Bless = 0
		m.playerShiHunFanObject.BlessTime = 0
		gameevent.Emit(shihunfaneventtypes.EventTypeShiHunFanAdvanced, m.p, int32(m.playerShiHunFanObject.AdvanceId))
	} else {
		m.playerShiHunFanObject.TimesNum += addTimes
		if m.playerShiHunFanObject.Bless == 0 {
			m.playerShiHunFanObject.BlessTime = now
		}
		m.playerShiHunFanObject.Bless += pro
	}
	m.playerShiHunFanObject.UpdateTime = now
	m.playerShiHunFanObject.SetModified()
	return
}

//直升券进阶
func (m *PlayerShiHunFanDataManager) ShiHunFanAdvancedTicket(addAdvancedNum int32) {
	if addAdvancedNum <= 0 {
		return
	}

	canAddNum := 0
	nextAdvancedId := m.playerShiHunFanObject.AdvanceId + 1
	for addAdvancedNum > 0 {
		anqiTemplate := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanNumber(int32(nextAdvancedId))
		if anqiTemplate == nil {
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
	m.playerShiHunFanObject.AdvanceId += canAddNum
	m.playerShiHunFanObject.TimesNum = 0
	m.playerShiHunFanObject.Bless = 0
	m.playerShiHunFanObject.BlessTime = 0
	m.playerShiHunFanObject.UpdateTime = now
	m.playerShiHunFanObject.SetModified()
	gameevent.Emit(shihunfaneventtypes.EventTypeShiHunFanAdvanced, m.p, int32(m.playerShiHunFanObject.AdvanceId))
	return
}

//噬魂幡充值数
func (m *PlayerShiHunFanDataManager) ShiHunFanCharge(num int32) {
	if num <= 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerShiHunFanObject.ChargeVal += num
	m.playerShiHunFanObject.UpdateTime = now
	m.playerShiHunFanObject.SetModified()
	return
}

//噬魂幡战斗力
func (m *PlayerShiHunFanDataManager) ShiHunFanPower(power int64) {
	if power <= 0 {
		return
	}
	curPower := m.playerShiHunFanObject.Power
	if curPower == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerShiHunFanObject.Power = power
	m.playerShiHunFanObject.UpdateTime = now
	m.playerShiHunFanObject.SetModified()
	gameevent.Emit(shihunfaneventtypes.EventTypeShiHunFanPowerChanged, m.p, power)
	return
}

func (m *PlayerShiHunFanDataManager) ToShiHunFanInfo() *shihunfancommon.ShiHunFanInfo {
	shiHunFanInfo := &shihunfancommon.ShiHunFanInfo{
		AdvanceId: int32(m.playerShiHunFanObject.AdvanceId),
		DanLevel:  m.playerShiHunFanObject.DanLevel,
		DanPro:    m.playerShiHunFanObject.DanPro,
		ChargeVal: m.playerShiHunFanObject.ChargeVal,
	}
	return shiHunFanInfo
}

func (m *PlayerShiHunFanDataManager) IfFullAdvanced() (flag bool) {
	if m.playerShiHunFanObject.AdvanceId == 0 {
		return
	}
	temp := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanNumber(int32(m.playerShiHunFanObject.AdvanceId))
	if temp == nil {
		return
	}
	if temp.NextId == 0 {
		return true
	}
	return
}

//仅gm使用 噬魂幡进阶
func (m *PlayerShiHunFanDataManager) GmSetShiHunFanAdvanced(advancedId int) {
	m.playerShiHunFanObject.AdvanceId = advancedId
	m.playerShiHunFanObject.TimesNum = int32(0)
	m.playerShiHunFanObject.Bless = int32(0)
	m.playerShiHunFanObject.BlessTime = int64(0)
	now := global.GetGame().GetTimeService().Now()
	m.playerShiHunFanObject.UpdateTime = now
	m.playerShiHunFanObject.SetModified()

	gameevent.Emit(shihunfaneventtypes.EventTypeShiHunFanAdvanced, m.p, int32(m.playerShiHunFanObject.AdvanceId))
	return
}

//仅gm使用 噬魂幡食丹等级
func (m *PlayerShiHunFanDataManager) GmSetShiHunFanShiHunFanDanLevel(level int32) {
	m.playerShiHunFanObject.DanLevel = level
	m.playerShiHunFanObject.DanNum = 0
	m.playerShiHunFanObject.DanPro = 0

	now := global.GetGame().GetTimeService().Now()
	m.playerShiHunFanObject.UpdateTime = now
	m.playerShiHunFanObject.SetModified()
}

//仅gm使用 噬魂幡充值数
func (m *PlayerShiHunFanDataManager) GmSetShiHunFanShiHunFanChargeVal(num int32) {
	m.playerShiHunFanObject.ChargeVal = num
	now := global.GetGame().GetTimeService().Now()
	m.playerShiHunFanObject.UpdateTime = now
	m.playerShiHunFanObject.SetModified()
}

func CreatePlayerShiHunFanDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerShiHunFanDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerShiHunFanDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerShiHunFanDataManager))
}
