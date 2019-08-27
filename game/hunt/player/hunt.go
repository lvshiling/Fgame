package player

import (
	"fgame/fgame/core/heartbeat"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/hunt/dao"
	hunteventtypes "fgame/fgame/game/hunt/event/types"
	hunttemplate "fgame/fgame/game/hunt/template"
	hunttypes "fgame/fgame/game/hunt/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//玩家寻宝管理器
type PlayerHuntDataManager struct {
	p player.Player
	//寻宝数据
	huntObjMap map[hunttypes.HuntType]*PlayerHuntObject
	//
	hbRunner heartbeat.HeartbeatTaskRunner
}

func (m *PlayerHuntDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerHuntDataManager) Load() (err error) {
	//加载寻宝数据
	entityList, err := dao.GetHuntDao().GetHuntEntityList(m.p.GetId())
	if err != nil {
		return
	}

	m.huntObjMap = make(map[hunttypes.HuntType]*PlayerHuntObject)
	for _, entity := range entityList {
		obj := NewPlayerHuntObject(m.p)
		obj.FromEntity(entity)
		m.huntObjMap[obj.huntType] = obj
	}

	m.initHuntObj()

	return
}

func (m *PlayerHuntDataManager) initHuntObj() {
	now := global.GetGame().GetTimeService().Now()
	for initType := hunttypes.MinHuntType; initType <= hunttypes.MaxHuntType; initType++ {
		if m.getHuntObj(initType) != nil {
			continue
		}

		id, _ := idutil.GetId()
		newObj := NewPlayerHuntObject(m.p)
		newObj.id = id
		newObj.huntType = initType
		newObj.freeHuntCount = 0
		newObj.totalHuntCount = 0
		newObj.lastHuntTime = now
		newObj.createTime = now
		newObj.SetModified()

		m.huntObjMap[initType] = newObj
	}
}

//加载后
func (m *PlayerHuntDataManager) AfterLoad() (err error) {
	m.hbRunner.AddTask(CreateFreeHuntRefresh(m.p))
	m.refreshFreeHuntTimes()
	return nil
}

//心跳
func (m *PlayerHuntDataManager) Heartbeat() {
	m.hbRunner.Heartbeat()
}

//更新免费次数信息
func (m *PlayerHuntDataManager) UpdateFreeHunt(huntType hunttypes.HuntType) {

	obj := m.getHuntObj(huntType)
	if obj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj.freeHuntCount += 1
	obj.lastHuntTime = now
	obj.updateTime = now
	obj.SetModified()
}

//更新免费次数信息
func (m *PlayerHuntDataManager) UpdateHuntCount(huntType hunttypes.HuntType, attendTimes int32) {
	if attendTimes < 1 {
		panic(fmt.Errorf("寻宝次数不能小于1，attendTimes:%d", attendTimes))
	}

	obj := m.getHuntObj(huntType)
	if obj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj.totalHuntCount += attendTimes
	obj.updateTime = now
	obj.SetModified()
}

func (m *PlayerHuntDataManager) RefreshFreeHuntTimes() {
	isCrossDay := m.refreshFreeHuntTimes()
	if isCrossDay {
		gameevent.Emit(hunteventtypes.EventTypeHuntCrossDay, m.p, nil)
	}
}

func (m *PlayerHuntDataManager) refreshFreeHuntTimes() bool {
	isCrossDay := false
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.huntObjMap {
		isSame, _ := timeutils.IsSameFive(now, obj.lastHuntTime)
		if isSame {
			continue
		}

		obj.freeHuntCount = 0
		obj.lastHuntTime = now
		obj.SetModified()

		isCrossDay = true
	}

	return isCrossDay
}

//是否免费次数
func (m *PlayerHuntDataManager) IsFreeTimes(huntType hunttypes.HuntType) bool {
	obj := m.getHuntObj(huntType)
	if obj == nil {
		return false
	}

	nextTemp := hunttemplate.GetHuntTemplateService().GetHuntTimesTemplat(obj.freeHuntCount + 1)
	if nextTemp == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	diff := now - obj.lastHuntTime
	if diff < nextTemp.JianGeTime {
		return false
	}

	return true
}

func (m *PlayerHuntDataManager) GetAllHuntInfo() map[hunttypes.HuntType]*PlayerHuntObject {
	return m.huntObjMap
}

func (m *PlayerHuntDataManager) GetHuntInfo(huntType hunttypes.HuntType) *PlayerHuntObject {
	return m.getHuntObj(huntType)
}

func (m *PlayerHuntDataManager) getHuntObj(huntType hunttypes.HuntType) *PlayerHuntObject {
	obj, ok := m.huntObjMap[huntType]
	if !ok {
		return nil
	}

	return obj
}

func CreatePlayerHuntDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerHuntDataManager{}
	m.p = p
	m.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerHuntDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerHuntDataManager))
}
