package player

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/constant/constant"
	constantttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/tower/dao"
	towereventtypes "fgame/fgame/game/tower/event/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//玩家打宝塔管理器
type PlayerTowerDataManager struct {
	p                 player.Player
	playerTowerObject *PlayerTowerObject
	dabaoTime         int64 //开始打宝塔时间
	dabaoFlag         bool  //打宝状态
	hbRunner          heartbeat.HeartbeatTaskRunner
}

func (m *PlayerTowerDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerTowerDataManager) Load() (err error) {
	//加载玩家打宝塔信息
	towerEntity, err := dao.GetTowerDao().GetTowerEntity(m.p.GetId())
	if err != nil {
		return
	}
	if towerEntity == nil {
		m.initPlayerTowerObject()
	} else {
		m.playerTowerObject = NewPlayerTowerObject(m.p)
		m.playerTowerObject.FromEntity(towerEntity)
	}
	return nil
}

//第一次初始化
func (m *PlayerTowerDataManager) initPlayerTowerObject() {
	o := NewPlayerTowerObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.useTime = 0
	o.extraTime = 0
	o.lastResetTime = now
	o.createTime = now
	m.playerTowerObject = o
	o.SetModified()
}

//加载后
func (m *PlayerTowerDataManager) AfterLoad() (err error) {
	m.hbRunner.AddTask(CreateRefreshTimeTask(m.p))
	m.hbRunner.AddTask(CreateNoticeTask(m.p))
	return
}

//心跳
func (m *PlayerTowerDataManager) Heartbeat() {
	m.hbRunner.Heartbeat()
}

func (m *PlayerTowerDataManager) StartDaBao() {
	now := global.GetGame().GetTimeService().Now()
	m.dabaoTime = now
	m.dabaoFlag = true

	gameevent.Emit(towereventtypes.EventTypeTowerStartDaBao, m.p, nil)
}

// 结束打宝时间
func (m *PlayerTowerDataManager) EndDaBao() {
	if m.dabaoFlag == false {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	isSame, _ := timeutils.IsSameDay(now, m.dabaoTime)
	if !isSame {
		beginDay, _ := timeutils.BeginOfNow(now)
		lastDayUse := beginDay - m.dabaoTime
		newDayUse := now - beginDay

		m.costTowerTime(lastDayUse)
		m.playerTowerObject.useTime = 0
		m.costTowerTime(newDayUse)

	} else {
		useTime := now - m.dabaoTime
		m.costTowerTime(useTime)
	}

	m.playerTowerObject.updateTime = now
	m.playerTowerObject.SetModified()
	m.dabaoFlag = false

	gameevent.Emit(towereventtypes.EventTypeTowerEndDaBao, m.p, nil)
}

// 结算打宝时间
func (m *PlayerTowerDataManager) checkTowerTime() {
	now := global.GetGame().GetTimeService().Now()
	beginDay, _ := timeutils.BeginOfNow(now)
	isSame, _ := timeutils.IsSameDay(now, m.playerTowerObject.lastResetTime)
	if !isSame {
		if m.dabaoFlag == true {
			lastDayUse := beginDay - m.dabaoTime
			m.costTowerTime(lastDayUse)
			m.dabaoTime = beginDay
		}

		m.playerTowerObject.useTime = 0
		m.playerTowerObject.lastResetTime = now
		m.playerTowerObject.updateTime = now
		m.playerTowerObject.SetModified()

		// 基础时间重置
		gameevent.Emit(towereventtypes.EventTypeTowerCrossDay, m.p, nil)
	}

	if m.dabaoFlag == false {
		return
	}

	if !m.IsOnDaBaoTime() {
		m.EndDaBao()
		gameevent.Emit(towereventtypes.EventTypeTowerEndDaBao, m.p, nil)
		return
	}
}

func (m *PlayerTowerDataManager) costTowerTime(useTime int64) {
	initTime := int64(constant.GetConstantService().GetConstant(constantttypes.ConstantTypeTowerInitTime))
	basic := initTime - m.playerTowerObject.useTime
	extral := m.playerTowerObject.extraTime

	remain := useTime
	if remain-basic <= 0 {
		m.playerTowerObject.useTime += remain
		return
	}
	remain -= basic
	m.playerTowerObject.useTime = initTime

	if remain-extral <= 0 {
		m.playerTowerObject.extraTime -= remain
	} else {
		m.playerTowerObject.extraTime = 0
	}
}

// 是否打宝时间
func (m *PlayerTowerDataManager) IsOnDaBaoTime() bool {
	now := global.GetGame().GetTimeService().Now()
	useTime := now - m.dabaoTime
	remainTime := m.getRemainTime()
	return useTime < remainTime
}

// 打宝剩余时间
func (m *PlayerTowerDataManager) GetRemainTime() int64 {
	return m.getRemainTime()
}

// 增加打宝时间
func (m *PlayerTowerDataManager) AddExtraTime(time int64) {
	if time < 0 {
		panic(fmt.Errorf("time不能小于0"))
	}

	m.playerTowerObject.extraTime += time
	m.playerTowerObject.SetModified()
}

func (m *PlayerTowerDataManager) getRemainTime() int64 {
	initTime := int64(constant.GetConstantService().GetConstant(constantttypes.ConstantTypeTowerInitTime))
	basic := initTime - m.playerTowerObject.useTime
	extra := m.playerTowerObject.extraTime
	return basic + extra
}

// 打宝时间提醒
func (m *PlayerTowerDataManager) noticeDaBaoTime() {
	if m.dabaoFlag == false {
		return
	}

	gameevent.Emit(towereventtypes.EventTypeTowerDaBaoNotice, m.p, m.dabaoTime)
}

func CreatePlayerTowerDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerTowerDataManager{}
	m.p = p
	m.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerTowerDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerTowerDataManager))
}
