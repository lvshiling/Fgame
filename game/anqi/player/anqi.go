package player

import (
	"fgame/fgame/game/anqi/dao"
	anqieventtypes "fgame/fgame/game/anqi/event/types"
	anqitemplate "fgame/fgame/game/anqi/template"
	anqitypes "fgame/fgame/game/anqi/types"
	"fgame/fgame/game/constant/constant"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家暗器管理器
type PlayerAnqiDataManager struct {
	p player.Player
	//玩家暗器对象
	playerAnqiObject *PlayerAnqiObject
}

func (m *PlayerAnqiDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerAnqiDataManager) Load() (err error) {
	//加载玩家暗器信息
	anqiEntity, err := dao.GetAnQiDao().GetAnQiEntity(m.p.GetId())
	if err != nil {
		return
	}
	if anqiEntity == nil {
		m.initPlayerAnqiObject()
	} else {
		m.playerAnqiObject = NewPlayerAnqiObject(m.p)
		m.playerAnqiObject.FromEntity(anqiEntity)
	}

	return nil
}

//第一次初始化
func (m *PlayerAnqiDataManager) initPlayerAnqiObject() {
	obj := NewPlayerAnqiObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.Id = id
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(m.p.GetRole(), m.p.GetSex())
	advanceId := playerCreateTemplate.Anqi
	obj.AdvanceId = int(advanceId)
	obj.AnqiDanLevel = 0
	obj.AnqiDanNum = int32(0)
	obj.AnqiDanPro = 0
	obj.TimesNum = int32(0)
	obj.Bless = int32(0)
	obj.BlessTime = int64(0)
	obj.Power = int64(0)
	obj.CreateTime = now
	obj.SetModified()

	m.playerAnqiObject = obj
}

func (m *PlayerAnqiDataManager) refreshBless() {
	now := global.GetGame().GetTimeService().Now()
	number := int32(m.playerAnqiObject.AdvanceId)
	nextNumber := number + 1
	anQiTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(nextNumber)
	if anQiTemplate == nil {
		return
	}
	if !anQiTemplate.GetIsClear() {
		return
	}
	lastTime := m.playerAnqiObject.BlessTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return
		}
		if !flag {
			m.playerAnqiObject.Bless = 0
			m.playerAnqiObject.BlessTime = 0
			m.playerAnqiObject.TimesNum = 0
			m.playerAnqiObject.SetModified()
		}
	}
	return
}

//加载后
func (m *PlayerAnqiDataManager) AfterLoad() (err error) {
	m.refreshBless()
	return nil
}

//暗器信息对象
func (m *PlayerAnqiDataManager) GetAnqiInfo() *PlayerAnqiObject {
	m.refreshBless()
	return m.playerAnqiObject
}

func (m *PlayerAnqiDataManager) GetAnqiAdvanced() int32 {
	return int32(m.playerAnqiObject.AdvanceId)
}

// func (m *PlayerAnqiDataManager) EatAnqiDan(pro int32, sucess bool) {
// 	if pro < 0 {
// 		return
// 	}
// 	if sucess {
// 		m.playerAnqiObject.AnqiDanLevel += 1
// 		m.playerAnqiObject.AnqiDanNum = 0
// 		m.playerAnqiObject.AnqiDanPro = pro
// 	} else {
// 		m.playerAnqiObject.AnqiDanNum += 1
// 		m.playerAnqiObject.AnqiDanPro += pro
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerAnqiObject.UpdateTime = now
// 	m.playerAnqiObject.SetModified()
// 	return
// }

func (m *PlayerAnqiDataManager) EatAnqiDan(level int32) {
	if m.playerAnqiObject.AnqiDanLevel == level || level <= 0 {
		return
	}
	anqiDanTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiDan(level)
	if anqiDanTemplate == nil {
		return
	}
	m.playerAnqiObject.AnqiDanLevel = level
	now := global.GetGame().GetTimeService().Now()
	m.playerAnqiObject.UpdateTime = now
	m.playerAnqiObject.SetModified()
	return
}

//心跳
func (m *PlayerAnqiDataManager) Heartbeat() {

}

//进阶
func (m *PlayerAnqiDataManager) AnqiAdvanced(pro, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		anqiTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(int32(m.playerAnqiObject.AdvanceId + 1))
		if anqiTemplate == nil {
			return
		}
		m.playerAnqiObject.AdvanceId += 1
		m.playerAnqiObject.TimesNum = 0
		m.playerAnqiObject.Bless = 0
		m.playerAnqiObject.BlessTime = 0
		gameevent.Emit(anqieventtypes.EventTypeAnqiAdvanced, m.p, int32(m.playerAnqiObject.AdvanceId))
	} else {
		m.playerAnqiObject.TimesNum += addTimes
		if m.playerAnqiObject.Bless == 0 {
			m.playerAnqiObject.BlessTime = now
		}
		m.playerAnqiObject.Bless += pro
	}
	m.playerAnqiObject.UpdateTime = now
	m.playerAnqiObject.SetModified()
	return
}

//直升券进阶
func (m *PlayerAnqiDataManager) AnqiAdvancedTicket(addAdvancedNum int32) {
	if addAdvancedNum <= 0 {
		return
	}

	canAddNum := 0
	nextAdvancedId := m.playerAnqiObject.AdvanceId + 1
	for addAdvancedNum > 0 {
		anqiTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(int32(nextAdvancedId))
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
	m.playerAnqiObject.AdvanceId += canAddNum
	m.playerAnqiObject.TimesNum = 0
	m.playerAnqiObject.Bless = 0
	m.playerAnqiObject.BlessTime = 0
	m.playerAnqiObject.UpdateTime = now
	m.playerAnqiObject.SetModified()
	gameevent.Emit(anqieventtypes.EventTypeAnqiAdvanced, m.p, int32(m.playerAnqiObject.AdvanceId))
	return
}

//暗器战斗力
func (m *PlayerAnqiDataManager) AnqiPower(power int64) {
	if power <= 0 {
		return
	}
	curPower := m.playerAnqiObject.Power
	if curPower == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerAnqiObject.Power = power
	m.playerAnqiObject.UpdateTime = now
	m.playerAnqiObject.SetModified()
	gameevent.Emit(anqieventtypes.EventTypeAnqiPowerChanged, m.p, power)
	return
}

func (m *PlayerAnqiDataManager) ToAnqiInfo() *anqitypes.AnqiInfo {
	bodyShieldInfo := &anqitypes.AnqiInfo{
		AdvancedId:   int32(m.playerAnqiObject.AdvanceId),
		AnqiDanLevel: m.playerAnqiObject.AnqiDanLevel,
		AnqiDanPro:   m.playerAnqiObject.AnqiDanPro,
	}
	return bodyShieldInfo
}

func (m *PlayerAnqiDataManager) IfFullAdvanced() (flag bool) {
	if m.playerAnqiObject.AdvanceId == 0 {
		return
	}
	anQiTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(int32(m.playerAnqiObject.AdvanceId))
	if anQiTemplate == nil {
		return
	}
	if anQiTemplate.NextId == 0 {
		return true
	}
	return
}

//仅gm使用 暗器进阶
func (m *PlayerAnqiDataManager) GmSetAnqiAdvanced(advancedId int) {
	m.playerAnqiObject.AdvanceId = advancedId
	m.playerAnqiObject.TimesNum = int32(0)
	m.playerAnqiObject.Bless = int32(0)
	m.playerAnqiObject.BlessTime = int64(0)
	now := global.GetGame().GetTimeService().Now()
	m.playerAnqiObject.UpdateTime = now
	m.playerAnqiObject.SetModified()

	gameevent.Emit(anqieventtypes.EventTypeAnqiAdvanced, m.p, int32(m.playerAnqiObject.AdvanceId))
	return
}

//仅gm使用 暗器食丹等级
func (m *PlayerAnqiDataManager) GmSetAnqiAnqiDanLevel(level int32) {
	m.playerAnqiObject.AnqiDanLevel = level
	m.playerAnqiObject.AnqiDanNum = 0
	m.playerAnqiObject.AnqiDanPro = 0

	now := global.GetGame().GetTimeService().Now()
	m.playerAnqiObject.UpdateTime = now
	m.playerAnqiObject.SetModified()
}

func CreatePlayerAnqiDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerAnqiDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerAnqiDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerAnqiDataManager))
}
