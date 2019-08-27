package player

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/qixue/dao"
	qixueeventtypes "fgame/fgame/game/qixue/event/types"
	qixuetemplate "fgame/fgame/game/qixue/template"
	qixuetypes "fgame/fgame/game/qixue/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/mathutils"
	"fmt"
	"math"
)

const (
	maxQiXueNum = 9999999999 //杀戮之心最大数量
)

//玩家泣血枪管理器
type PlayerQiXueDataManager struct {
	p player.Player
	//玩家泣血枪对象
	playerQiXueObject *PlayerQiXueObject
}

func (m *PlayerQiXueDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerQiXueDataManager) Load() (err error) {
	//加载玩家泣血枪信息
	qixueEntity, err := dao.GetQiXueDao().GetQiXueEntity(m.p.GetId())
	if err != nil {
		return
	}
	if qixueEntity == nil {
		m.initPlayerQiXueObject()
	} else {
		m.playerQiXueObject = NewPlayerQiXueObject(m.p)
		m.playerQiXueObject.FromEntity(qixueEntity)
	}

	return nil
}

//第一次初始化
func (m *PlayerQiXueDataManager) initPlayerQiXueObject() {
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(m.p.GetRole(), m.p.GetSex())

	obj := NewPlayerQiXueObject(m.p)
	obj.id = id
	qixueTemplate := qixuetemplate.GetQiXueTemplateService().GetQiXueTemplate(playerCreateTemplate.QiXue)
	if qixueTemplate == nil {
		obj.currLevel = int32(0)
		obj.currStar = int32(0)
	} else {
		obj.currLevel = qixueTemplate.Level
		obj.currStar = qixueTemplate.Star
	}
	obj.timesNum = int32(0)
	obj.shaLuNum = int64(0)
	obj.lastTime = int64(0)
	obj.power = int64(0)
	obj.createTime = now
	obj.SetModified()

	m.playerQiXueObject = obj
}

//加载后
func (m *PlayerQiXueDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (m *PlayerQiXueDataManager) Heartbeat() {

}

//泣血枪信息对象
func (m *PlayerQiXueDataManager) GetQiXueInfo() *PlayerQiXueObject {
	return m.playerQiXueObject
}

func (m *PlayerQiXueDataManager) IfDropCD() bool {
	now := global.GetGame().GetTimeService().Now()
	dropCd := int64(qixuetemplate.GetQiXueTemplateService().GetQiXueConstantTemplate().DropCd)
	return now > m.playerQiXueObject.lastTime+dropCd
}

//TODO:xzk 放在模版处理
//杀戮心掉落百分比
func (m *PlayerQiXueDataManager) getDropPercent() (percent float64) {
	minDropPercent := int(qixuetemplate.GetQiXueTemplateService().GetQiXueConstantTemplate().DropPercentMin)
	maxDropPercent := int(qixuetemplate.GetQiXueTemplateService().GetQiXueConstantTemplate().DropPercentMax) + 1
	return float64(mathutils.RandomRange(minDropPercent, maxDropPercent)) / float64(common.MAX_RATE)
}

//TODO:xzk 放在模版处理
//杀戮心掉落几率
func (m *PlayerQiXueDataManager) ifHitDropShaLu() bool {
	dropSqRate := int(qixuetemplate.GetQiXueTemplateService().GetQiXueConstantTemplate().DropRate)
	return mathutils.RandomHit(common.MAX_RATE, dropSqRate)
}

//掉落
func (m *PlayerQiXueDataManager) QiXueDrop() (flag bool, dropNum int64, costStar int32) {
	//是否掉落冷却中
	now := global.GetGame().GetTimeService().Now()
	if !m.IfDropCD() {
		return
	}

	//杀戮心掉落
	slNum := m.playerQiXueObject.shaLuNum
	if slNum > 0 {
		if m.ifHitDropShaLu() {
			dropNum = int64(math.Ceil(float64(slNum) * m.getDropPercent()))
			m.playerQiXueObject.shaLuNum -= dropNum
			m.playerQiXueObject.lastTime = now
		}
	}

	//泣血枪掉星
	curLev := m.playerQiXueObject.currLevel
	curStar := m.playerQiXueObject.currStar
	curTemplate := qixuetemplate.GetQiXueTemplateService().GetQiXueTemplateByLevel(curLev, curStar)
	if curTemplate != nil {
		if curTemplate.IfHitReduceStar() {
			//掉落几颗星数
			newLev := curLev
			newStar := curStar
			temp := curTemplate
			subStar := temp.GetRandomReduceStar()
			for subStar > 0 && newStar > 0 {
				subStar -= 1
				costStar += 1
				dropNum += int64(temp.StarCount)
				newStar -= 1
				if newStar == 0 {
					newLev -= 1
					newStar = qixuetemplate.GetQiXueTemplateService().GetQiXueStarCount(newLev)
				}
				if newLev == 0 && newStar == 0 {
					break
				}

				temp = qixuetemplate.GetQiXueTemplateService().GetQiXueTemplateByLevel(newLev, newStar)
			}

			m.playerQiXueObject.currLevel = newLev
			m.playerQiXueObject.currStar = newStar
			m.playerQiXueObject.lastTime = now

			//发送降阶事件
			if curLev != newLev {
				eventData := qixueeventtypes.CreatePlayerQiXueChangedWeaponEventData(curLev, curStar, newLev, newStar)
				gameevent.Emit(qixueeventtypes.EventTypeQiXueDegrade, m.p, eventData)
			}

			// //日志
			// qixueReason := commonlog.QiXueLogReasonDegrade
			// reasonText := qixueReason.String()
			// data := qixueeventtypes.CreatePlayerQiXueChangedLogEventData(int32(oldAdvanceId), costStar, slNum, qixueReason, reasonText)
			// gameevent.Emit(qixueeventtypes.EventTypeQiXueChangedLog, m.p, data)
		}
	}

	m.playerQiXueObject.updateTime = now
	m.playerQiXueObject.SetModified()
	flag = true
	return
}

func (m *PlayerQiXueDataManager) AddShaLu(num int32) bool {
	if num <= 0 {
		panic(fmt.Errorf("掉落的数量不能小于1，num:%d", num))
	}

	now := global.GetGame().GetTimeService().Now()
	newNum := m.playerQiXueObject.shaLuNum + int64(num)
	if newNum > maxQiXueNum {
		newNum = maxQiXueNum
	}
	m.playerQiXueObject.shaLuNum = newNum
	m.playerQiXueObject.updateTime = now
	m.playerQiXueObject.SetModified()

	return true
}

//泣血枪战斗力
func (m *PlayerQiXueDataManager) QiXuePower(power int64) {
	if power <= 0 {
		return
	}

	//相同的话不做存储,影响存储io
	if m.playerQiXueObject.power == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerQiXueObject.power = power
	m.playerQiXueObject.updateTime = now
	m.playerQiXueObject.SetModified()
	return
}

func (m *PlayerQiXueDataManager) ToQiXueInfo() *qixuetypes.QiXueInfo {
	bodyShieldInfo := &qixuetypes.QiXueInfo{
		CurrLevel: int32(m.playerQiXueObject.currLevel),
		CurrStar:  int32(m.playerQiXueObject.currStar),
		ShaLuNum:  int64(m.playerQiXueObject.shaLuNum),
	}
	return bodyShieldInfo
}

//升级泣血枪
func (m *PlayerQiXueDataManager) UseShaLuNum() (flag bool) {
	if !m.playerQiXueObject.IfEnoughShaLuNum() {
		return
	}

	oldLev := m.playerQiXueObject.currLevel
	oldStar := m.playerQiXueObject.currStar
	oldTemplate := qixuetemplate.GetQiXueTemplateService().GetQiXueTemplateByLevel(oldLev, oldStar)
	if oldTemplate == nil {
		return
	}

	nextLevelTemp := oldTemplate.GetNextTemp()
	if nextLevelTemp == nil {
		return
	}

	newLev := oldLev
	newStar := oldStar

	shaLuTimes := m.playerQiXueObject.timesNum
	needShaLuNum := int64(nextLevelTemp.UseResources - shaLuTimes)
	remainShaLuNum := m.playerQiXueObject.shaLuNum
	if remainShaLuNum > needShaLuNum {
		remainShaLuNum -= needShaLuNum
		shaLuTimes = 0
		newLev = nextLevelTemp.Level
		newStar = nextLevelTemp.Star

	} else {
		shaLuTimes += int32(remainShaLuNum)
		remainShaLuNum = 0
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerQiXueObject.timesNum = shaLuTimes
	m.playerQiXueObject.shaLuNum = remainShaLuNum
	m.playerQiXueObject.currLevel = newLev
	m.playerQiXueObject.currStar = newStar
	m.playerQiXueObject.updateTime = now
	m.playerQiXueObject.SetModified()

	if oldLev != newLev || oldStar != newStar {
		//发送升阶事件
		eventData := qixueeventtypes.CreatePlayerQiXueChangedWeaponEventData(oldLev, oldStar, newLev, newStar)
		gameevent.Emit(qixueeventtypes.EventTypeQiXueAdvanced, m.p, eventData)
		// TODO ：xzk25 进阶日志
		// qixueReason := commonlog.QiXueLogReasonAdvanced
		// reasonText := fmt.Sprintf(qixueReason.String(), commontypes.AdvancedTypeShaQi.String())
		// data := qixueeventtypes.CreatePlayerQiXueChangedLogEventData(beforeNum, 1, beforeShaLuNum, qixueReason, reasonText)
		// gameevent.Emit(qixueeventtypes.EventTypeQiXueChangedLog, pl, data)
	}

	flag = true
	return
}

//仅gm使用 泣血枪进阶
func (m *PlayerQiXueDataManager) GmSetQiXueAdvanced(lev int32, star int32) {
	qiXueTemplate := qixuetemplate.GetQiXueTemplateService().GetQiXueTemplateByLevel(lev, star)
	if qiXueTemplate == nil {
		return
	}
	oldLev := m.playerQiXueObject.currLevel
	oldStar := m.playerQiXueObject.currStar

	now := global.GetGame().GetTimeService().Now()
	m.playerQiXueObject.currLevel = lev
	m.playerQiXueObject.currStar = star
	m.playerQiXueObject.timesNum = int32(0)
	m.playerQiXueObject.updateTime = now
	m.playerQiXueObject.SetModified()

	//发送事件
	if oldLev > lev {
		eventData := qixueeventtypes.CreatePlayerQiXueChangedWeaponEventData(oldLev, oldStar, lev, star)
		gameevent.Emit(qixueeventtypes.EventTypeQiXueDegrade, m.p, eventData)
	} else {
		eventData := qixueeventtypes.CreatePlayerQiXueChangedWeaponEventData(oldLev, oldStar, lev, star)
		gameevent.Emit(qixueeventtypes.EventTypeQiXueAdvanced, m.p, eventData)
	}

	return
}

//仅gm使用 泣血枪设置杀气值
func (m *PlayerQiXueDataManager) GmSetQiXueShaLuNum(num int64) {
	if num < 0 {
		panic(fmt.Errorf("GM设置杀戮数量不能小于0，num:%d", num))
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerQiXueObject.shaLuNum = num
	m.playerQiXueObject.updateTime = now
	m.playerQiXueObject.SetModified()
	return
}

func CreatePlayerQiXueDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerQiXueDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerQiXueDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerQiXueDataManager))
}
