package player

import (
	dianxingcommon "fgame/fgame/game/dianxing/common"
	"fgame/fgame/game/dianxing/dao"
	dianxingeventtypes "fgame/fgame/game/dianxing/event/types"
	dianxingtemplate "fgame/fgame/game/dianxing/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	xiuxianbookeventtypes "fgame/fgame/game/welfare/xiuxianbook/event/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

const XINGCHEN_MAX = 9999999999

//玩家点星系统管理器
type PlayerDianXingDataManager struct {
	p player.Player
	//玩家点星系统对象
	playerDianXingObject *PlayerDianXingObject
}

func (m *PlayerDianXingDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerDianXingDataManager) Load() (err error) {
	//加载玩家点星系统信息
	dianXingEntity, err := dao.GetDianXingDao().GetDianXingEntity(m.p.GetId())
	if err != nil {
		return
	}
	if dianXingEntity == nil {
		m.initPlayerDianXingObject()
	} else {
		m.playerDianXingObject = NewPlayerDianXingObject(m.p)
		m.playerDianXingObject.FromEntity(dianXingEntity)
	}

	return nil
}

//第一次初始化
func (m *PlayerDianXingDataManager) initPlayerDianXingObject() {
	obj := NewPlayerDianXingObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.Id = id
	obj.CreateTime = now
	obj.SetModified()

	m.playerDianXingObject = obj
}

//加载后
func (m *PlayerDianXingDataManager) AfterLoad() (err error) {
	m.refreshDianXingBless()
	return nil
}

//点星系统对象
func (m *PlayerDianXingDataManager) GetDianXingObject() *PlayerDianXingObject {
	m.refreshDianXingBless()
	return m.playerDianXingObject
}

//心跳
func (m *PlayerDianXingDataManager) Heartbeat() {

}

//进阶
func (m *PlayerDianXingDataManager) DianXingAdvanced(pro int32, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		nextTemplate := m.GetNextDianXingTemplate()
		if nextTemplate == nil {
			return
		}
		oldXingPu := m.playerDianXingObject.CurrType
		m.playerDianXingObject.CurrType = nextTemplate.XingPuType
		m.playerDianXingObject.CurrLevel = nextTemplate.Level
		m.playerDianXingObject.DianXingBless = 0
		m.playerDianXingObject.DianXingBlessTime = 0
		m.playerDianXingObject.DianXingTimes = 0
		if oldXingPu != 0 && oldXingPu < m.playerDianXingObject.CurrType {
			//发送升阶事件
			eventData := oldXingPu
			gameevent.Emit(dianxingeventtypes.EventTypeDianXingAdvanced, m.p, eventData)
		}
	} else {
		m.playerDianXingObject.DianXingTimes += addTimes
		if m.playerDianXingObject.DianXingBless == 0 {
			m.playerDianXingObject.DianXingBlessTime = now
		}
		m.playerDianXingObject.DianXingBless += pro
	}
	m.playerDianXingObject.UpdateTime = now
	m.playerDianXingObject.SetModified()
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, m.p, nil)
	return
}

//解封进阶
func (m *PlayerDianXingDataManager) DianXingJieFengAdvanced(pro int32, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		nextTemplate := m.GetNextDianXingJieFengTemplate()
		if nextTemplate == nil {
			return
		}
		m.playerDianXingObject.JieFengLev = nextTemplate.Level
		m.playerDianXingObject.JieFengBless = 0
		m.playerDianXingObject.JieFengTimes = 0
	} else {
		m.playerDianXingObject.JieFengTimes += addTimes
		m.playerDianXingObject.JieFengBless += pro
	}
	m.playerDianXingObject.UpdateTime = now
	m.playerDianXingObject.SetModified()
	gameevent.Emit(xiuxianbookeventtypes.XiuxianBookEventTypeUpdateObj, m.p, nil)
	return
}

func (m *PlayerDianXingDataManager) refreshDianXingBless() (err error) {
	now := global.GetGame().GetTimeService().Now()
	nextTemplate := m.GetNextDianXingTemplate()
	if nextTemplate == nil {
		return
	}
	if !nextTemplate.GetIsClear() {
		return
	}
	lastTime := m.playerDianXingObject.DianXingBlessTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return err
		}
		if !flag {
			m.playerDianXingObject.DianXingBless = 0
			m.playerDianXingObject.DianXingBlessTime = 0
			m.playerDianXingObject.DianXingTimes = 0
			m.playerDianXingObject.SetModified()
		}
	}
	return
}

//获取下一个点星升级数据
func (m *PlayerDianXingDataManager) GetNextDianXingTemplate() *gametemplate.DianXingTemplate {
	var nextDianXingTemplate *gametemplate.DianXingTemplate
	if m.playerDianXingObject.CurrType == 0 {
		nextDianXingTemplate = dianxingtemplate.GetDianXingTemplateService().GetDianXingTemplateByArg(1, 1)
	} else {
		//判断系统是否可以升级
		currTemplate := dianxingtemplate.GetDianXingTemplateService().GetDianXingTemplateByArg(m.playerDianXingObject.CurrType, m.playerDianXingObject.CurrLevel)
		nextDianXingTemplate = currTemplate.GetNextTemplate()
	}
	return nextDianXingTemplate
}

//获取下一个点星解封升级数据
func (m *PlayerDianXingDataManager) GetNextDianXingJieFengTemplate() *gametemplate.DianXingJieFengTemplate {
	var nextDianXingJieFengTemplate *gametemplate.DianXingJieFengTemplate
	if m.playerDianXingObject.JieFengLev == 0 {
		nextDianXingJieFengTemplate = dianxingtemplate.GetDianXingTemplateService().GetDianXingJieFengTemplateByLev(1)
	} else {
		//判断系统是否可以升级
		currTemplate := dianxingtemplate.GetDianXingTemplateService().GetDianXingJieFengTemplateByLev(m.playerDianXingObject.JieFengLev)
		nextDianXingJieFengTemplate = currTemplate.GetNextTemplate()
	}
	return nextDianXingJieFengTemplate
}

//捡起掉落星尘
func (m *PlayerDianXingDataManager) DropXingChen(num int32) (mine *PlayerDianXingObject) {
	now := global.GetGame().GetTimeService().Now()
	oidNum := m.playerDianXingObject.XingChenNum
	newNum := oidNum + int64(num)
	if newNum > XINGCHEN_MAX {
		newNum = XINGCHEN_MAX
	}
	m.playerDianXingObject.XingChenNum = int64(newNum)
	m.playerDianXingObject.UpdateTime = now
	m.playerDianXingObject.SetModified()
	mine = m.playerDianXingObject
	return
}

//点星系统战斗力
func (m *PlayerDianXingDataManager) DianXingPower(power int64) {
	if power <= 0 {
		return
	}
	if m.playerDianXingObject.Power == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerDianXingObject.Power = power
	m.playerDianXingObject.UpdateTime = now
	m.playerDianXingObject.SetModified()
	return
}

func (m *PlayerDianXingDataManager) ToDianXingInfo() *dianxingcommon.DianXingInfo {
	dianXingInfo := &dianxingcommon.DianXingInfo{
		CurrType:     int32(m.playerDianXingObject.CurrType),
		CurrLevel:    int32(m.playerDianXingObject.CurrLevel),
		JieFengLev:   int32(m.playerDianXingObject.JieFengLev),
		JieFengBless: int32(m.playerDianXingObject.JieFengBless),
	}
	return dianXingInfo
}

func (m *PlayerDianXingDataManager) SubXingChenNum(num int64) bool {
	flag := false
	if num <= 0 {
		return flag
	}
	if num > m.playerDianXingObject.XingChenNum {
		return flag
	}
	m.playerDianXingObject.XingChenNum -= num
	now := global.GetGame().GetTimeService().Now()
	m.playerDianXingObject.UpdateTime = now
	m.playerDianXingObject.SetModified()
	flag = true

	return flag
}

func (m *PlayerDianXingDataManager) GmSetDianXingAdvanced(xingPu int32, lev int32) {
	now := global.GetGame().GetTimeService().Now()
	oldXingPu := m.playerDianXingObject.CurrType
	m.playerDianXingObject.CurrType = xingPu
	m.playerDianXingObject.CurrLevel = lev
	m.playerDianXingObject.DianXingBless = 0
	m.playerDianXingObject.DianXingBlessTime = 0
	m.playerDianXingObject.DianXingTimes = 0
	m.playerDianXingObject.UpdateTime = now
	if oldXingPu != 0 && oldXingPu < m.playerDianXingObject.CurrType {
		//发送升阶事件
		eventData := oldXingPu
		gameevent.Emit(dianxingeventtypes.EventTypeDianXingAdvanced, m.p, eventData)
	}
	m.playerDianXingObject.SetModified()
	return
}

func (m *PlayerDianXingDataManager) GmSetDianXingJieFengAdvanced(lev int32) {
	now := global.GetGame().GetTimeService().Now()
	m.playerDianXingObject.JieFengLev = lev
	m.playerDianXingObject.JieFengBless = 0
	m.playerDianXingObject.JieFengTimes = 0
	m.playerDianXingObject.UpdateTime = now
	m.playerDianXingObject.SetModified()
	return
}

//仅gm使用 设置星尘值
func (m *PlayerDianXingDataManager) GmSetXingChenNum(num int64) {
	m.playerDianXingObject.XingChenNum = num
	now := global.GetGame().GetTimeService().Now()
	m.playerDianXingObject.UpdateTime = now
	m.playerDianXingObject.SetModified()
	return
}

func CreatePlayerDianXingDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerDianXingDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerDianXingDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerDianXingDataManager))
}
