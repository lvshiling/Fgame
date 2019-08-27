package player

import (
	"fgame/fgame/core/heartbeat"
	daliwandao "fgame/fgame/game/daliwan/dao"
	daliwaneventtypes "fgame/fgame/game/daliwan/event/types"
	daliwantemplate "fgame/fgame/game/daliwan/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
)

type PlayerDaLiWanManager struct {
	p          player.Player
	daliWanMap map[int32]*DaLiWanObject //大力丸
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

//***********接口开始****************

//玩家
func (m *PlayerDaLiWanManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerDaLiWanManager) Load() error {
	playerId := m.p.GetId()
	daliWanEntityList, err := daliwandao.GetDaLiWanDao().GetDailiWanList(playerId)
	if err != nil {
		return err
	}

	m.daliWanMap = make(map[int32]*DaLiWanObject)

	for _, daliWanEntity := range daliWanEntityList {
		dailiWanObj := NewDaLiWanObject(m.p)
		err = dailiWanObj.FromEntity(daliWanEntity)
		if err != nil {
			return err
		}
		m.daliWanMap[dailiWanObj.typ] = dailiWanObj
	}
	m.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	return nil
}

//加载后
func (m *PlayerDaLiWanManager) AfterLoad() error {
	m.checkExpire()
	m.heartbeatRunner.AddTask(CreateDaLiWanTask(m.p))
	return nil
}

func (m *PlayerDaLiWanManager) checkExpire() {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.daliWanMap {
		if obj.expired != 0 {
			continue
		}
		if obj.IsExpire(now) {
			obj.expired = 1
			obj.updateTime = now
			obj.SetModified()

		}
	}
}

//心跳
func (m *PlayerDaLiWanManager) Heartbeat() {
	m.heartbeatRunner.Heartbeat()
}

func (m *PlayerDaLiWanManager) Use(typ int32) bool {
	linshiTemplate := daliwantemplate.GetDaLiWanTemplateService().GetLinShiTemplate(typ)
	if linshiTemplate == nil {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	daLiWanObj, ok := m.daliWanMap[typ]
	if !ok {
		daLiWanObj = NewDaLiWanObject(m.p)
		daLiWanObj.id, _ = idutil.GetId()
		daLiWanObj.startTime = now
		daLiWanObj.typ = typ
		daLiWanObj.duration = int64(linshiTemplate.Time)
		daLiWanObj.expired = 0
		daLiWanObj.createTime = now
		daLiWanObj.SetModified()
		m.daliWanMap[typ] = daLiWanObj
		gameevent.Emit(daliwaneventtypes.DaLiWanEventTypeUpdate, m.p, daLiWanObj)
		return true
	}
	if daLiWanObj.IsExpire(now) {
		daLiWanObj.startTime = now
		daLiWanObj.duration = int64(linshiTemplate.Time)
		daLiWanObj.expired = 0
		daLiWanObj.updateTime = now
		daLiWanObj.SetModified()
		gameevent.Emit(daliwaneventtypes.DaLiWanEventTypeUpdate, m.p, daLiWanObj)
	} else {
		daLiWanObj.duration += int64(linshiTemplate.Time)
		daLiWanObj.updateTime = now
		daLiWanObj.SetModified()
		gameevent.Emit(daliwaneventtypes.DaLiWanEventTypeUpdate, m.p, daLiWanObj)
	}
	return true
}

func (m *PlayerDaLiWanManager) IsExpire(typ int32) bool {

	now := global.GetGame().GetTimeService().Now()
	daLiWanObj, ok := m.daliWanMap[typ]
	if !ok {
		return true
	}
	if daLiWanObj.IsExpire(now) {
		return true
	}
	return false
}

func (m *PlayerDaLiWanManager) CheckExpire() {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.daliWanMap {
		if obj.expired != 0 {
			continue
		}
		if obj.IsExpire(now) {
			obj.expired = 1
			obj.updateTime = now
			obj.SetModified()
			//移除
			gameevent.Emit(daliwaneventtypes.DaLiWanEventTypeExpire, m.p, obj)
		}
	}
}

func (m *PlayerDaLiWanManager) GetDaliWanMap() map[int32]*DaLiWanObject {
	return m.daliWanMap
}

func NewPlayerDaLiWanManager(pl player.Player) player.PlayerDataManager {
	rst := &PlayerDaLiWanManager{
		p: pl,
	}
	return rst
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerDaLiWanDataManagerType, player.PlayerDataManagerFactoryFunc(NewPlayerDaLiWanManager))
}
