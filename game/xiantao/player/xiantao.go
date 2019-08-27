package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/game/xiantao/dao"
	xiantaoeventtypes "fgame/fgame/game/xiantao/event/types"
	xiantaotemplate "fgame/fgame/game/xiantao/template"
	"fgame/fgame/pkg/idutil"
)

//玩家仙桃大会管理器
type PlayerXianTaoDataManager struct {
	p player.Player
	//玩家仙桃大会对象
	xianTaoObject *PlayerXianTaoObject
}

func (pddm *PlayerXianTaoDataManager) Player() player.Player {
	return pddm.p
}

func (pddm *PlayerXianTaoDataManager) GetXianTaoObject() *PlayerXianTaoObject {
	return pddm.xianTaoObject
}

//加载
func (pddm *PlayerXianTaoDataManager) Load() (err error) {
	//加载玩家仙桃大会
	xianTaoEntity, err := dao.GetXianTaoDao().GetXianTaoEntity(pddm.p.GetId())
	if err != nil {
		return
	}
	if xianTaoEntity == nil {
		pddm.initPlayerXianTaoObject()
	} else {
		pddm.xianTaoObject = NewPlayerXianTaoObject(pddm.p)
		pddm.xianTaoObject.FromEntity(xianTaoEntity)
	}

	return nil
}

//第一次初始化
func (pddm *PlayerXianTaoDataManager) initPlayerXianTaoObject() {
	pdo := NewPlayerXianTaoObject(pddm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pdo.Id = id
	//生成id
	pdo.PlayerId = pddm.p.GetId()
	pdo.JuniorPeachCount = int32(0)
	pdo.HighPeachCount = int32(0)
	pdo.RobCount = int32(0)
	pdo.BeRobCount = int32(0)
	pdo.EndTime = int64(0)
	pdo.CreateTime = now
	pddm.xianTaoObject = pdo
	pdo.SetModified()
}

//加载后
func (pddm *PlayerXianTaoDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (pddm *PlayerXianTaoDataManager) Heartbeat() {

}

//获取下一个劫取数据
func (pddm *PlayerXianTaoDataManager) GetNextRobTimesTemplate() *gametemplate.XianTaoTimesTemplate {
	var nextTemplate *gametemplate.XianTaoTimesTemplate
	if pddm.xianTaoObject.RobCount == 0 {
		nextTemplate = xiantaotemplate.GetXianTaoTemplateService().GetXianTaoTimesTempByTimes(1)
	} else {
		curTemplate := xiantaotemplate.GetXianTaoTemplateService().GetXianTaoTimesTempByTimes(pddm.xianTaoObject.RobCount)
		nextTemplate = curTemplate.GetNextTemplate()
	}
	return nextTemplate
}

//获取下一个被劫取数据
func (pddm *PlayerXianTaoDataManager) GetNextBeRobTimesTemplate() *gametemplate.XianTaoTimesTemplate {
	var nextTemplate *gametemplate.XianTaoTimesTemplate
	if pddm.xianTaoObject.BeRobCount == 0 {
		nextTemplate = xiantaotemplate.GetXianTaoTemplateService().GetXianTaoTimesTempByTimes(1)
	} else {
		curTemplate := xiantaotemplate.GetXianTaoTemplateService().GetXianTaoTimesTempByTimes(pddm.xianTaoObject.BeRobCount)
		nextTemplate = curTemplate.GetNextTemplate()
		if nextTemplate == nil {
			nextTemplate = curTemplate
		}
	}
	return nextTemplate
}

//玩家增加劫取次数
func (pddm *PlayerXianTaoDataManager) AddRobCount(count int32) {
	now := global.GetGame().GetTimeService().Now()
	pddm.xianTaoObject.RobCount += count
	pddm.xianTaoObject.UpdateTime = now
	pddm.xianTaoObject.SetModified()
	return
}

//玩家设置被劫取次数
func (pddm *PlayerXianTaoDataManager) SetBeRobCount(count int32) {
	now := global.GetGame().GetTimeService().Now()
	pddm.xianTaoObject.BeRobCount = count
	pddm.xianTaoObject.UpdateTime = now
	pddm.xianTaoObject.SetModified()
	return
}

//玩家增加百年仙桃数量
func (pddm *PlayerXianTaoDataManager) AddJuniorPeachCount(count int32) (curCount, addCount int32) {
	constTemp := xiantaotemplate.GetXianTaoTemplateService().GetXianTaoConstTemplate()
	if constTemp.XianTaoMax > pddm.xianTaoObject.JuniorPeachCount {
		if constTemp.XianTaoMax < pddm.xianTaoObject.JuniorPeachCount+count {
			count = constTemp.XianTaoMax - pddm.xianTaoObject.JuniorPeachCount
		}
		now := global.GetGame().GetTimeService().Now()
		pddm.xianTaoObject.JuniorPeachCount += count
		pddm.xianTaoObject.UpdateTime = now
		pddm.xianTaoObject.SetModified()
	} else {
		count = 0
	}
	addCount = count
	curCount = pddm.xianTaoObject.JuniorPeachCount
	gameevent.Emit(xiantaoeventtypes.EventTypeBaiNianXianTaoChange, pddm.p, nil)
	return
}

//玩家减少百年仙桃数量
func (pddm *PlayerXianTaoDataManager) SubJuniorPeachCount(count int32) (curCount, subCount int32) {
	now := global.GetGame().GetTimeService().Now()
	if pddm.xianTaoObject.JuniorPeachCount > count {
		pddm.xianTaoObject.JuniorPeachCount -= count
	} else {
		count = pddm.xianTaoObject.JuniorPeachCount
		pddm.xianTaoObject.JuniorPeachCount = 0
	}
	subCount = count
	curCount = pddm.xianTaoObject.JuniorPeachCount

	pddm.xianTaoObject.UpdateTime = now
	pddm.xianTaoObject.SetModified()
	gameevent.Emit(xiantaoeventtypes.EventTypeBaiNianXianTaoChange, pddm.p, nil)
	return
}

//玩家增加千年仙桃数量
func (pddm *PlayerXianTaoDataManager) AddHighPeachCount(count int32) (curCount, addCount int32) {
	constTemp := xiantaotemplate.GetXianTaoTemplateService().GetXianTaoConstTemplate()
	if constTemp.XianTaoMax > pddm.xianTaoObject.HighPeachCount {
		if constTemp.XianTaoMax < pddm.xianTaoObject.HighPeachCount+count {
			count = constTemp.XianTaoMax - pddm.xianTaoObject.HighPeachCount
		}
		now := global.GetGame().GetTimeService().Now()
		pddm.xianTaoObject.HighPeachCount += count
		pddm.xianTaoObject.UpdateTime = now
		pddm.xianTaoObject.SetModified()
	} else {
		count = 0
	}
	addCount = count
	curCount = pddm.xianTaoObject.HighPeachCount
	gameevent.Emit(xiantaoeventtypes.EventTypeQianNianXianTaoChange, pddm.p, nil)
	return
}

//玩家减少千年仙桃数量
func (pddm *PlayerXianTaoDataManager) SubHighPeachCount(count int32) (curCount, subCount int32) {
	now := global.GetGame().GetTimeService().Now()
	if pddm.xianTaoObject.HighPeachCount > count {
		pddm.xianTaoObject.HighPeachCount -= count
	} else {
		count = pddm.xianTaoObject.HighPeachCount
		pddm.xianTaoObject.HighPeachCount = 0
	}
	subCount = count
	curCount = pddm.xianTaoObject.HighPeachCount

	pddm.xianTaoObject.UpdateTime = now
	pddm.xianTaoObject.SetModified()
	gameevent.Emit(xiantaoeventtypes.EventTypeQianNianXianTaoChange, pddm.p, nil)
	return
}

//玩家清理所有仙桃数量
func (pddm *PlayerXianTaoDataManager) ClearAllPeachCount() {
	now := global.GetGame().GetTimeService().Now()
	pddm.xianTaoObject.JuniorPeachCount = int32(0)
	gameevent.Emit(xiantaoeventtypes.EventTypeBaiNianXianTaoChange, pddm.p, nil)
	pddm.xianTaoObject.HighPeachCount = int32(0)
	gameevent.Emit(xiantaoeventtypes.EventTypeQianNianXianTaoChange, pddm.p, nil)
	pddm.xianTaoObject.UpdateTime = now
	pddm.xianTaoObject.SetModified()
	return
}

//进入仙桃大会
func (pddm *PlayerXianTaoDataManager) EnterXianTao(endTime int64) {
	if pddm.xianTaoObject.EndTime == endTime {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pddm.xianTaoObject.JuniorPeachCount = int32(0)
	pddm.xianTaoObject.HighPeachCount = int32(0)
	pddm.xianTaoObject.RobCount = int32(0)
	pddm.xianTaoObject.BeRobCount = int32(0)
	pddm.xianTaoObject.EndTime = endTime
	pddm.xianTaoObject.UpdateTime = now
	pddm.xianTaoObject.SetModified()
	return
}

//玩家退出
func (pddm *PlayerXianTaoDataManager) ExitXianTao() {
	return
}

func CreatePlayerXianTaoDataManager(p player.Player) player.PlayerDataManager {
	pddm := &PlayerXianTaoDataManager{}
	pddm.p = p
	return pddm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerXianTaoDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerXianTaoDataManager))
}
