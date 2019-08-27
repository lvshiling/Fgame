package player

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/massacre/dao"
	massacreeventtypes "fgame/fgame/game/massacre/event/types"
	massacretemplate "fgame/fgame/game/massacre/template"
	massacretypes "fgame/fgame/game/massacre/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/mathutils"
	"math"
)

//玩家戮仙刃管理器
type PlayerMassacreDataManager struct {
	p player.Player
	//玩家戮仙刃对象
	playerMassacreObject *PlayerMassacreObject
}

func (m *PlayerMassacreDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerMassacreDataManager) Load() (err error) {
	//加载玩家戮仙刃信息
	massacreEntity, err := dao.GetMassacreDao().GetMassacreEntity(m.p.GetId())
	if err != nil {
		return
	}
	if massacreEntity == nil {
		m.initPlayerMassacreObject()
	} else {
		m.playerMassacreObject = NewPlayerMassacreObject(m.p)
		m.playerMassacreObject.FromEntity(massacreEntity)
	}

	return nil
}

//第一次初始化
func (m *PlayerMassacreDataManager) initPlayerMassacreObject() {
	obj := NewPlayerMassacreObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.Id = id
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(m.p.GetRole(), m.p.GetSex())
	advanceId := playerCreateTemplate.Massacre
	obj.AdvanceId = int(advanceId)

	massacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(obj.AdvanceId)
	if massacreTemplate == nil {
		obj.CurrLevel = int32(0)
		obj.CurrStar = int32(0)
	} else {
		obj.CurrLevel = massacreTemplate.Type
		obj.CurrStar = massacreTemplate.Star
	}
	obj.TimesNum = int32(0)
	obj.ShaQiNum = int64(0)
	obj.LastTime = int64(0)
	obj.Power = int64(0)
	obj.CreateTime = now
	obj.SetModified()

	m.playerMassacreObject = obj
}

//加载后
func (m *PlayerMassacreDataManager) AfterLoad() (err error) {
	return nil
}

//戮仙刃信息对象
func (m *PlayerMassacreDataManager) GetMassacreInfo() *PlayerMassacreObject {
	return m.playerMassacreObject
}

//心跳
func (m *PlayerMassacreDataManager) Heartbeat() {

}

//进阶
func (m *PlayerMassacreDataManager) MassacreAdvanced(addTimes int32, sucess bool) {
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		massacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(m.playerMassacreObject.AdvanceId + 1)
		if massacreTemplate == nil {
			return
		}
		oldAdvanceId := m.playerMassacreObject.AdvanceId
		oldLev := m.playerMassacreObject.CurrLevel
		m.playerMassacreObject.AdvanceId += 1
		m.playerMassacreObject.CurrLevel = massacreTemplate.Type
		m.playerMassacreObject.CurrStar = massacreTemplate.Star
		m.playerMassacreObject.TimesNum = 0
		if oldLev != m.playerMassacreObject.CurrLevel {
			eventData := massacreeventtypes.CreatePlayerMassacreAdvanceEventData(int32(oldAdvanceId), int32(m.playerMassacreObject.AdvanceId))
			//发送升阶事件
			gameevent.Emit(massacreeventtypes.EventTypeMassacreAdvanced, m.p, eventData)

		}
	} else {
		m.playerMassacreObject.TimesNum += addTimes
	}
	m.playerMassacreObject.UpdateTime = now
	m.playerMassacreObject.SetModified()
	return
}

func (m *PlayerMassacreDataManager) IfCanDrop() bool {
	now := global.GetGame().GetTimeService().Now()
	dropCd := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypePlayerDropSqCd))
	if now < m.playerMassacreObject.LastTime+dropCd {
		return false
	}
	return true
}

//掉落
func (m *PlayerMassacreDataManager) MassacreDrop(attackName string) (flag bool, bagDropNum int64, costStar int32) {
	//是否掉落冷却中
	now := global.GetGame().GetTimeService().Now()
	if !m.IfCanDrop() {
		return
	}

	itemNum := m.playerMassacreObject.ShaQiNum
	minDropPercent := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBagDropSqMinPer))
	maxDropPercent := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBagDropSqMaxPer))
	dropSqRate := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDropSqRate))
	if m.playerMassacreObject.ShaQiNum > 0 && mathutils.RandomHit(common.MAX_RATE, dropSqRate) {
		bagDropPercent := float64(mathutils.RandomRange(minDropPercent, maxDropPercent)) / float64(common.MAX_RATE)
		bagDropNum = int64(math.Ceil(float64(itemNum) * bagDropPercent))
		m.playerMassacreObject.ShaQiNum -= bagDropNum
		m.playerMassacreObject.LastTime = now
	}

	oldAdvanceId := m.playerMassacreObject.AdvanceId
	oldMassacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(oldAdvanceId)

	//是否掉戮仙刃等级,0阶0星的不掉落杀气
	if oldMassacreTemplate != nil && mathutils.RandomHit(common.MAX_RATE, int(oldMassacreTemplate.GasPercent)) {
		//掉落几颗星数
		newAdvanceId := oldAdvanceId
		befLev := oldMassacreTemplate.Type
		befStar := oldMassacreTemplate.Star
		newLev := befLev
		newStar := befStar
		subSrar := int32(mathutils.RandomRange(int(oldMassacreTemplate.GasMin), int(oldMassacreTemplate.GasMax)))
		if subSrar > 0 {
			tempTemplate := oldMassacreTemplate
			for subSrar > 0 {
				newAdvanceId--
				subSrar--
				costStar++
				bagDropNum += int64(tempTemplate.StarCount)
				if newAdvanceId == 0 {
					newLev = 0
					newStar = 0
					break
				}
				tempTemplate = massacretemplate.GetMassacreTemplateService().GetMassacre(newAdvanceId)
				newLev = tempTemplate.Type
				newStar = tempTemplate.Star
			}
			m.playerMassacreObject.AdvanceId = newAdvanceId
			m.playerMassacreObject.CurrLevel = newLev
			m.playerMassacreObject.CurrStar = newStar
			m.playerMassacreObject.LastTime = now
			if befLev != newLev {
				eventData := massacreeventtypes.CreatePlayerMassacreDegradeEventData(int32(oldAdvanceId), int32(newAdvanceId), attackName)
				//发送降阶事件
				gameevent.Emit(massacreeventtypes.EventTypeMassacreDegrade, m.p, eventData)
			}
		}
		//日志
		massacreReason := commonlog.MassacreLogReasonDegrade
		reasonText := massacreReason.String()
		data := massacreeventtypes.CreatePlayerMassacreChangedLogEventData(int32(oldAdvanceId), costStar, itemNum, massacreReason, reasonText)
		gameevent.Emit(massacreeventtypes.EventTypeMassacreChangedLog, m.p, data)
	}
	m.playerMassacreObject.UpdateTime = now
	m.playerMassacreObject.SetModified()
	flag = true
	return
}

//设置玩家戮仙刃阶星
// func (m *PlayerMassacreDataManager) SetMassacreObjInfo(advanceId int) {
// 	if advanceId == m.playerMassacreObject.AdvanceId {
// 		return
// 	}
// 	now := global.GetGame().GetTimeService().Now()
// 	massacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(advanceId)
// 	if massacreTemplate == nil {
// 		m.playerMassacreObject.AdvanceId = 0
// 		m.playerMassacreObject.CurrLevel = 0
// 		m.playerMassacreObject.CurrStar = 0
// 		m.playerMassacreObject.TimesNum = 0
// 	} else {
// 		m.playerMassacreObject.AdvanceId = massacreTemplate.TemplateId()
// 		m.playerMassacreObject.CurrLevel = massacreTemplate.Type
// 		m.playerMassacreObject.CurrStar = massacreTemplate.Star
// 		m.playerMassacreObject.TimesNum = 0
// 	}

// 	m.playerMassacreObject.UpdateTime = now
// 	m.playerMassacreObject.SetModified()
// 	return
// }

//设置玩家戮仙刃杀气掉落cd
// func (m *PlayerMassacreDataManager) SetMassacreDrop(num int64) {
// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerMassacreObject.ShaQiNum = num
// 	m.playerMassacreObject.LastTime = now
// 	m.playerMassacreObject.UpdateTime = now
// 	m.playerMassacreObject.SetModified()
// 	return
// }

const SHAQI_MAX = 9999999999

func (m *PlayerMassacreDataManager) DropShaQi(num int32) (mine *PlayerMassacreObject) {
	now := global.GetGame().GetTimeService().Now()
	oidNum := m.playerMassacreObject.ShaQiNum
	newNum := oidNum + int64(num)
	if newNum > SHAQI_MAX {
		newNum = SHAQI_MAX
	}
	m.playerMassacreObject.ShaQiNum = int64(newNum)
	m.playerMassacreObject.UpdateTime = now
	m.playerMassacreObject.SetModified()
	mine = m.playerMassacreObject
	return
}

//戮仙刃战斗力
func (m *PlayerMassacreDataManager) MassacrePower(power int64) {
	if power <= 0 {
		return
	}
	if m.playerMassacreObject.Power == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerMassacreObject.Power = power
	m.playerMassacreObject.UpdateTime = now
	m.playerMassacreObject.SetModified()
	return
}

func (m *PlayerMassacreDataManager) ToMassacreInfo() *massacretypes.MassacreInfo {
	bodyShieldInfo := &massacretypes.MassacreInfo{
		AdvancedId: int32(m.playerMassacreObject.AdvanceId),
		CurrLevel:  int32(m.playerMassacreObject.CurrLevel),
		CurrStar:   int32(m.playerMassacreObject.CurrStar),
		ShaQiNum:   int64(m.playerMassacreObject.ShaQiNum),
	}
	return bodyShieldInfo
}

func (m *PlayerMassacreDataManager) SubShaQiNum(num int64) bool {
	flag := false
	if num <= 0 {
		return flag
	}
	if num > m.playerMassacreObject.ShaQiNum {
		return flag
	}
	m.playerMassacreObject.ShaQiNum -= num
	now := global.GetGame().GetTimeService().Now()
	m.playerMassacreObject.UpdateTime = now
	m.playerMassacreObject.SetModified()
	flag = true

	return flag
}

//仅gm使用 戮仙刃进阶
func (m *PlayerMassacreDataManager) GmSetMassacreAdvanced(lev int32, star int32) {
	massacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacreNumber(lev, star)
	if massacreTemplate == nil {
		return
	}
	oldAdvanceId := m.playerMassacreObject.AdvanceId
	oldMassacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(oldAdvanceId)
	befLev := int32(0)
	if oldMassacreTemplate != nil {
		befLev = oldMassacreTemplate.Type
	}
	newLev := int32(0)
	newAdvanceId := massacreTemplate.Id
	newMassacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(newAdvanceId)
	if newMassacreTemplate != nil {
		newLev = newMassacreTemplate.Type
	}
	m.playerMassacreObject.AdvanceId = massacreTemplate.TemplateId()
	m.playerMassacreObject.CurrLevel = massacreTemplate.Type
	m.playerMassacreObject.CurrStar = massacreTemplate.Star
	m.playerMassacreObject.TimesNum = int32(0)
	now := global.GetGame().GetTimeService().Now()
	m.playerMassacreObject.UpdateTime = now
	m.playerMassacreObject.SetModified()
	if befLev > newLev {
		eventData := massacreeventtypes.CreatePlayerMassacreDegradeEventData(int32(oldAdvanceId), int32(newAdvanceId), "")
		//发送降阶事件
		gameevent.Emit(massacreeventtypes.EventTypeMassacreDegrade, m.p, eventData)
	} else if befLev < newLev {
		eventData := massacreeventtypes.CreatePlayerMassacreAdvanceEventData(int32(oldAdvanceId), int32(newAdvanceId))
		//发送升阶事件
		gameevent.Emit(massacreeventtypes.EventTypeMassacreAdvanced, m.p, eventData)
	}

	return
}

//仅gm使用 戮仙刃设置杀气值
func (m *PlayerMassacreDataManager) GmSetMassacreShaQiNum(num int64) {
	m.playerMassacreObject.ShaQiNum = num
	now := global.GetGame().GetTimeService().Now()
	m.playerMassacreObject.UpdateTime = now
	m.playerMassacreObject.SetModified()
	return
}

func CreatePlayerMassacreDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerMassacreDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerMassacreDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerMassacreDataManager))
}
