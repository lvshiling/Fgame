package player

import (
	"fgame/fgame/game/arenapvp/dao"
	arenapvpeventtypes "fgame/fgame/game/arenapvp/event/types"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func init() {
	player.RegisterPlayerDataManager(types.PlayerArenapvpDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerArenapvpDataManager))
}

const (
	logSize = 50
)

//玩家竞技场数据管理器
type PlayerArenapvpDataManager struct {
	p               player.Player
	arenapvpObject  *PlayerArenapvpObject
	guessLogObjList []*PlayerArenapvpGuessLogObject
}

//玩家
func (m *PlayerArenapvpDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerArenapvpDataManager) Load() (err error) {
	//加载玩家活动数据列表
	arenapvpEntity, err := dao.GetArenapvpDao().GetPlayerArenapvp(m.Player().GetId())
	if err != nil {
		return
	}

	if arenapvpEntity != nil {
		arenapvpObject := NewPlayerArenapvpObject(m.p)
		err = arenapvpObject.FromEntity(arenapvpEntity)
		if err != nil {
			return
		}
		m.arenapvpObject = arenapvpObject
	} else {
		m.initPlayerArenapvpObject()
	}

	//加载玩家竞猜日志列表
	logEntityList, err := dao.GetArenapvpDao().GetPlayerArenapvpGuessLogList(m.Player().GetId())
	if err != nil {
		return
	}

	for _, entity := range logEntityList {
		obj := NewPlayerArenapvpGuessLogObject(m.p)
		obj.FromEntity(entity)
		m.guessLogObjList = append(m.guessLogObjList, obj)
	}

	return nil
}

//第一次初始化
func (m *PlayerArenapvpDataManager) initPlayerArenapvpObject() {
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o := NewPlayerArenapvpObject(m.p)
	o.id = id
	o.reliveTimes = 0
	o.outStatus = 0
	o.pvpRecord = arenapvptypes.ArenapvpTypeInit
	o.guessNotice = 1
	o.createTime = now
	o.SetModified()
	m.arenapvpObject = o
}

//加载后
func (m *PlayerArenapvpDataManager) AfterLoad() error {
	m.refresh()
	return nil
}

//心跳
func (m *PlayerArenapvpDataManager) Heartbeat() {

}

//个人竞猜日志列表
func (m *PlayerArenapvpDataManager) GetGuessLogList() []*PlayerArenapvpGuessLogObject {
	return m.guessLogObjList
}

//竞猜结算
func (m *PlayerArenapvpDataManager) GuessResult(raceNum int32, guessType arenapvptypes.ArenapvpType, winnerId int64) (logObj *PlayerArenapvpGuessLogObject) {
	logObj = m.getGuessLog(raceNum, guessType)
	if logObj == nil {
		return nil
	}

	if logObj.winnerId != 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	logObj.winnerId = winnerId
	logObj.updateTime = now
	logObj.SetModified()

	return
}

//添加竞猜日志
func (m *PlayerArenapvpDataManager) AddGuessLog(raceNum int32, guessType arenapvptypes.ArenapvpType, guessId int64) (flag bool) {
	lastObj := m.GetLastGuessLog()

	if lastObj != nil && lastObj.IfAttendGuess(raceNum, guessType) {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	var logObj *PlayerArenapvpGuessLogObject
	if len(m.guessLogObjList) >= logSize {
		logObj = m.guessLogObjList[0]
		m.guessLogObjList = m.guessLogObjList[1:]
	} else {
		logObj = NewPlayerArenapvpGuessLogObject(m.p)
		id, _ := idutil.GetId()
		logObj.id = id
		logObj.createTime = now
	}

	logObj.guessId = guessId
	logObj.guessType = guessType
	logObj.raceNum = raceNum
	logObj.updateTime = now
	logObj.SetModified()
	m.guessLogObjList = append(m.guessLogObjList, logObj)

	flag = true
	return
}

//最近的竞猜日志
func (m *PlayerArenapvpDataManager) GetLastGuessLog() *PlayerArenapvpGuessLogObject {
	len := len(m.guessLogObjList)
	if len == 0 {
		return nil
	}

	return m.guessLogObjList[len-1]
}

func (m *PlayerArenapvpDataManager) getGuessLog(raceNum int32, guessType arenapvptypes.ArenapvpType) *PlayerArenapvpGuessLogObject {
	for _, logObj := range m.guessLogObjList {
		if logObj.raceNum != raceNum {
			continue
		}

		if logObj.guessType != guessType {
			continue
		}

		return logObj
	}

	return nil
}

func (m *PlayerArenapvpDataManager) refresh() {
	m.refreshArenapvp()
}

//登录检查玩家pvp信息
func (m *PlayerArenapvpDataManager) refreshArenapvp() {
	now := global.GetGame().GetTimeService().Now()
	isSame, _ := timeutils.IsSameDay(now, m.arenapvpObject.updateTime)
	if !isSame {
		m.arenapvpObject.outStatus = 0
		m.arenapvpObject.pvpRecord = arenapvptypes.ArenapvpTypeInit
		m.arenapvpObject.ticketFlag = 0
		m.arenapvpObject.updateTime = now
		m.arenapvpObject.SetModified()
	}
}

//进入pvp
func (m *PlayerArenapvpDataManager) EnterArenapvp() {
	if m.arenapvpObject.pvpRecord == arenapvptypes.ArenapvpTypeInit {
		m.arenapvpObject.pvpRecord = arenapvptypes.ArenapvpTypeElection
	}
	now := global.GetGame().GetTimeService().Now()
	m.arenapvpObject.reliveTimes = 0
	m.arenapvpObject.updateTime = now
	m.arenapvpObject.SetModified()
	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpReliveChanged, m.p, nil)
}

// 玩家复活
func (m *PlayerArenapvpDataManager) Relive() {
	now := global.GetGame().GetTimeService().Now()
	m.arenapvpObject.reliveTimes += 1
	m.arenapvpObject.updateTime = now
	m.arenapvpObject.SetModified()

	//发送事件
	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpReliveChanged, m.p, nil)
}

// 重置复活次数
func (m *PlayerArenapvpDataManager) ResetRelive() {
	now := global.GetGame().GetTimeService().Now()
	m.arenapvpObject.reliveTimes = 0
	m.arenapvpObject.updateTime = now
	m.arenapvpObject.SetModified()
	//发送事件
	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpReliveChanged, m.p, nil)
}

// pvp结果
func (m *PlayerArenapvpDataManager) ArenapvpResult(win bool, pvpType arenapvptypes.ArenapvpType, addJiFen int32) {
	if addJiFen < 0 {
		panic(fmt.Errorf("addJiFen 小于0，%d", addJiFen))
	}

	now := global.GetGame().GetTimeService().Now()
	if !win {
		m.arenapvpObject.outStatus = 1
	}
	m.arenapvpObject.setPvpRecord(win, pvpType)
	m.arenapvpObject.jiFen += addJiFen
	m.arenapvpObject.updateTime = now
	m.arenapvpObject.SetModified()

	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpJiFenChanged, m.p, m.arenapvpObject)
	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpResult, m.p, m.arenapvpObject)
}

//消耗积分
func (m *PlayerArenapvpDataManager) UseJiFen(num int32) bool {
	if !m.arenapvpObject.IfEnoughJiFen(num) {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	m.arenapvpObject.jiFen -= num
	m.arenapvpObject.updateTime = now
	m.arenapvpObject.SetModified()

	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpJiFenChanged, m.p, m.arenapvpObject)
	return true
}

//添加积分
func (m *PlayerArenapvpDataManager) AddJiFen(num int32) {
	if num < 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.arenapvpObject.jiFen += num
	m.arenapvpObject.updateTime = now
	m.arenapvpObject.SetModified()

	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpJiFenChanged, m.p, m.arenapvpObject)
	return
}

//设置积分
func (m *PlayerArenapvpDataManager) GMSetJiFen(num int32) {
	if num < 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.arenapvpObject.jiFen = num
	m.arenapvpObject.updateTime = now
	m.arenapvpObject.SetModified()

	gameevent.Emit(arenapvpeventtypes.EventTypeArenapvpJiFenChanged, m.p, m.arenapvpObject)
	return
}

//玩家pvp信息
func (m *PlayerArenapvpDataManager) GetPlayerArenapvpObj() *PlayerArenapvpObject {
	m.refreshArenapvp()
	return m.arenapvpObject
}

//玩家pvp信息
func (m *PlayerArenapvpDataManager) GuessNoticeSetting(notice int32) {
	if m.arenapvpObject.guessNotice == notice {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.arenapvpObject.guessNotice = notice
	m.arenapvpObject.updateTime = now
	m.arenapvpObject.SetModified()
}

//玩家pvp信息
func (m *PlayerArenapvpDataManager) BuyArenapvpTicket() (flag bool) {
	if m.arenapvpObject.IfBuyTicket() {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.arenapvpObject.ticketFlag = 1
	m.arenapvpObject.updateTime = now
	m.arenapvpObject.SetModified()

	flag = true
	return
}

func CreatePlayerArenapvpDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerArenapvpDataManager{}
	m.p = p
	return m
}
