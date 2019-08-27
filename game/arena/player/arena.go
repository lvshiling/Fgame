package player

import (
	"fgame/fgame/game/arena/dao"
	arenaeventtypes "fgame/fgame/game/arena/event/types"
	arenatemplate "fgame/fgame/game/arena/template"
	arenatypes "fgame/fgame/game/arena/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"math"
)

func init() {
	player.RegisterPlayerDataManager(types.PlayerArenaDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerArenaDataManager))
}

//玩家竞技场数据管理器
type PlayerArenaDataManager struct {
	p                 player.Player
	playerArenaObject *PlayerArenaObject
	inviteTime        int64
}

//玩家
func (m *PlayerArenaDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerArenaDataManager) Load() (err error) {
	//加载玩家活动数据列表
	arenaEntity, err := dao.GetArenaDao().GetArena(m.Player().GetId())
	if err != nil {
		return
	}

	if arenaEntity != nil {
		playerArenaObject := CreatePlayerArenaObject(m.Player())
		err = playerArenaObject.FromEntity(arenaEntity)
		if err != nil {
			return
		}
		m.playerArenaObject = playerArenaObject
	} else {
		m.initPlayerArenaObject()
	}

	return nil
}

//第一次初始化
func (m *PlayerArenaDataManager) initPlayerArenaObject() {
	o := CreatePlayerArenaObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.culRewardTime = 0
	o.totalRewardTime = 0
	o.endTime = 0
	o.createTime = now
	o.SetModified()
	m.playerArenaObject = o
}

//加载后
func (m *PlayerArenaDataManager) AfterLoad() error {
	return nil
}

//心跳
func (m *PlayerArenaDataManager) Heartbeat() {

}

//获胜
func (m *PlayerArenaDataManager) EnterArena(endTime int64) {
	now := global.GetGame().GetTimeService().Now()
	m.playerArenaObject.reliveTime = 0
	m.playerArenaObject.updateTime = now
	m.playerArenaObject.SetModified()
	gameevent.Emit(arenaeventtypes.EventTypeArenaReliveChanged, m.p, nil)
}

//3v3结束奖励
func (m *PlayerArenaDataManager) ArenaFinish(win bool, ratio, maxJifenPercent, addJifenPercent int32) {
	m.refresh()

	now := global.GetGame().GetTimeService().Now()

	// 积分相关
	if win {
		m.playerArenaObject.winCount += 1
		m.playerArenaObject.dayWinCount += 1
		m.playerArenaObject.failCount = 0
	} else {
		m.playerArenaObject.winCount = 0
		m.playerArenaObject.dayWinCount = 0
		m.playerArenaObject.failCount += 1
	}

	curDayWinCount := m.playerArenaObject.dayWinCount
	if curDayWinCount > m.playerArenaObject.dayMaxWinCount {
		m.playerArenaObject.dayMaxWinCount = curDayWinCount
	}

	m.arenaWinJiFen(win, ratio, maxJifenPercent, addJifenPercent)
	m.playerArenaObject.culRewardTime += 1
	m.playerArenaObject.totalRewardTime += 1
	m.playerArenaObject.updateTime = now
	m.playerArenaObject.SetModified()

	//发送事件
	gameevent.Emit(arenaeventtypes.EventTypeArenaWinChanged, m.p, win)
	return
}

func (m *PlayerArenaDataManager) Relive() {
	now := global.GetGame().GetTimeService().Now()
	m.playerArenaObject.reliveTime += 1
	m.playerArenaObject.updateTime = now
	m.playerArenaObject.SetModified()
	//发送事件
	gameevent.Emit(arenaeventtypes.EventTypeArenaReliveChanged, m.p, nil)
}

// 重置复活次数
func (m *PlayerArenaDataManager) ResetRelive() {
	now := global.GetGame().GetTimeService().Now()
	m.playerArenaObject.reliveTime = 0
	m.playerArenaObject.updateTime = now
	m.playerArenaObject.SetModified()
	//发送事件
	gameevent.Emit(arenaeventtypes.EventTypeArenaReliveChanged, m.p, nil)
}

func (m *PlayerArenaDataManager) refresh() {
	m.refreshJiFenDay()
	m.refreshRewardTime()
}

func (m *PlayerArenaDataManager) refreshRewardTime() {
	now := global.GetGame().GetTimeService().Now()
	lastTime := m.playerArenaObject.updateTime
	flag, err := timeutils.IsSameFive(lastTime, now)
	if err != nil {
		return
	}
	if !flag {
		m.playerArenaObject.updateTime = now
		m.playerArenaObject.culRewardTime = 0
		m.playerArenaObject.SetModified()
	}
}

func (m *PlayerArenaDataManager) refreshJiFenDay() {
	now := global.GetGame().GetTimeService().Now()
	lastTime := m.playerArenaObject.arenaTime
	flag, err := timeutils.IsSameFive(lastTime, now)
	if err != nil {
		return
	}
	if !flag {
		m.playerArenaObject.arenaTime = now
		m.playerArenaObject.dayWinCount = 0
		m.playerArenaObject.dayMaxWinCount = 0
		m.playerArenaObject.jiFenDay = 0
		m.playerArenaObject.SetModified()
	}
}

func (m *PlayerArenaDataManager) GetPlayerArenaObject() *PlayerArenaObject {
	return m.playerArenaObject
}

func (m *PlayerArenaDataManager) GetPlayerArenaObjectByRefresh() *PlayerArenaObject {
	m.refresh()
	return m.playerArenaObject
}

const (
	inviteCD = 30 * common.SECOND
)

func (m *PlayerArenaDataManager) Invite() bool {
	if !m.IfInviteCD() {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	m.inviteTime = now
	return true
}

func (m *PlayerArenaDataManager) IfInviteCD() bool {
	now := global.GetGame().GetTimeService().Now()
	if now-m.inviteTime < int64(inviteCD) {
		return false
	}
	return true
}

// 是否领取周榜奖励
func (m *PlayerArenaDataManager) IsHasedReward(rankTime int64) (flag bool) {
	if m.playerArenaObject.rankRewTime != rankTime {
		return
	}

	flag = true
	return
}

//领取周榜奖励
func (m *PlayerArenaDataManager) RankRewardGet(rankTime int64) (flag bool) {
	if m.IsHasedReward(rankTime) {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerArenaObject.rankRewTime = rankTime
	m.playerArenaObject.updateTime = now
	m.playerArenaObject.SetModified()
	flag = true
	return
}

//消耗积分
func (m *PlayerArenaDataManager) UsePoint(num int32) bool {
	if !m.playerArenaObject.IfEnoughPoint(num) {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerArenaObject.jiFenCount -= num
	m.playerArenaObject.updateTime = now
	m.playerArenaObject.SetModified()

	gameevent.Emit(arenaeventtypes.EventTypeArenaJiFenChanged, m.p, nil)
	return true
}

//放弃/退出
func (m *PlayerArenaDataManager) ArenaGiveUp(ratio, maxJifenPercent, addJifenPercent int32) {
	now := global.GetGame().GetTimeService().Now()
	m.playerArenaObject.winCount = 0
	m.playerArenaObject.dayWinCount = 0
	m.playerArenaObject.failCount += 1
	m.playerArenaObject.updateTime = now
	m.arenaWinJiFen(false, ratio, maxJifenPercent, addJifenPercent)
	m.playerArenaObject.SetModified()

	//发送事件
	gameevent.Emit(arenaeventtypes.EventTypeArenaWinChanged, m.p, false)
	return
}

//添加积分
func (m *PlayerArenaDataManager) arenaWinJiFen(win bool, ratio, maxJifenPercent, addJifenPercent int32) {
	if ratio <= 0 {
		return
	}

	var areaType arenatypes.ArenaType
	arenaCount := int32(0)
	if win {
		areaType = arenatypes.ArenaTypeArena
		arenaCount = m.playerArenaObject.winCount
	} else {
		areaType = arenatypes.ArenaTypeFail
		arenaCount = m.playerArenaObject.failCount
	}

	arenaTemplate := arenatemplate.GetArenaTemplateService().GetArenaTemplate(areaType, arenaCount)
	if arenaTemplate == nil {
		return
	}

	//小于最大积分上限
	addJiFen := arenaTemplate.LianXuGetJifen
	addJiFen = int32(math.Ceil(float64(addJiFen*ratio) * (float64(common.MAX_RATE+addJifenPercent) / float64(common.MAX_RATE))))
	// addJiFen += int32(math.Ceil(float64(addJiFen*addJifenPercent) / float64(common.MAX_RATE)))
	arenaConstantTemplate := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate()
	dayMax := arenaConstantTemplate.JiFenMaxDay
	dayMax += int32(math.Ceil(float64(dayMax*maxJifenPercent) / float64(common.MAX_RATE)))
	remainJiFen := dayMax - m.playerArenaObject.jiFenDay
	if addJiFen > remainJiFen {
		addJiFen = remainJiFen
	}

	if addJiFen <= 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerArenaObject.jiFenDay += addJiFen
	m.playerArenaObject.jiFenCount += addJiFen
	m.playerArenaObject.arenaTime = now
}

//添加积分
func (m *PlayerArenaDataManager) AddJiFen(num int32) {
	if num <= 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerArenaObject.jiFenCount += num
	m.playerArenaObject.updateTime = now
	m.playerArenaObject.SetModified()

	gameevent.Emit(arenaeventtypes.EventTypeArenaJiFenChanged, m.p, nil)
	return
}

//消耗积分
func (m *PlayerArenaDataManager) GMSetPoint(num int32) {
	if num < 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerArenaObject.jiFenCount = num
	m.playerArenaObject.updateTime = now
	m.playerArenaObject.SetModified()

	gameevent.Emit(arenaeventtypes.EventTypeArenaJiFenChanged, m.p, nil)
	return
}

//消耗积分
func (m *PlayerArenaDataManager) GMSetTodayPoint(num int32) {
	if num < 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerArenaObject.jiFenDay = num
	m.playerArenaObject.updateTime = now
	m.playerArenaObject.SetModified()

	gameevent.Emit(arenaeventtypes.EventTypeArenaJiFenChanged, m.p, nil)
	return
}

func CreatePlayerArenaDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerArenaDataManager{}
	m.p = p
	return m
}
