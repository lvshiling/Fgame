package player

import (
	gameevent "fgame/fgame/game/event"
	fushieventtypes "fgame/fgame/game/fushi/event/types"
	fushitypes "fgame/fgame/game/fushi/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	"fgame/fgame/game/fushi/dao"
)

// 玩家八卦符石管理器
type PlayerFuShiDataManager struct {
	p player.Player
	// 玩家符石对象
	playerFushiObjectMap map[fushitypes.FuShiType]*PlayerFuShiObject
}

func (m *PlayerFuShiDataManager) Player() player.Player {
	return m.p
}

func (m *PlayerFuShiDataManager) Load() (err error) {
	fushiEntityList, err := dao.GetFushiDao().GetFuShiEntity(m.p.GetId())
	if err != nil {
		return
	}

	m.playerFushiObjectMap = make(map[fushitypes.FuShiType]*PlayerFuShiObject)
	for _, entity := range fushiEntityList {
		obj := NewPlayerFuShiObject(m.p)
		obj.FromEntity(entity)

		m.playerFushiObjectMap[obj.typ] = obj
	}

	return nil
}

func (m *PlayerFuShiDataManager) AfterLoad() (err error) {
	return
}

func (m *PlayerFuShiDataManager) Heartbeat() {
}

func (m *PlayerFuShiDataManager) GetFushiInfo() map[fushitypes.FuShiType]*PlayerFuShiObject {
	return m.playerFushiObjectMap
}

// 通过符石类型获取符石对象
func (m *PlayerFuShiDataManager) GetFushiByTyp(typ fushitypes.FuShiType) *PlayerFuShiObject {
	fushiObj, ok := m.playerFushiObjectMap[typ]
	if !ok {
		fushiObj = m.initFuShiObject(typ)
		m.playerFushiObjectMap[typ] = fushiObj
	}

	return fushiObj
}

func (m *PlayerFuShiDataManager) initFuShiObject(typ fushitypes.FuShiType) *PlayerFuShiObject {
	now := global.GetGame().GetTimeService().Now()
	fushiObj := NewPlayerFuShiObject(m.p)
	id, _ := idutil.GetId()
	fushiObj.id = id
	fushiObj.typ = typ
	fushiObj.fushiLevel = 0
	fushiObj.createTime = now
	return fushiObj
}

// 符石是否激活
func (m *PlayerFuShiDataManager) IsFuShiActivite(typ fushitypes.FuShiType) bool {
	obj := m.GetFushiByTyp(typ)
	return obj.fushiLevel > 0
}

// 符石激活成功
func (m *PlayerFuShiDataManager) FushiActitiveSucess(typ fushitypes.FuShiType) bool {
	now := global.GetGame().GetTimeService().Now()
	obj := m.GetFushiByTyp(typ)
	if obj.fushiLevel > 0 {
		return false
	}
	obj.fushiLevel = 1
	obj.updateTime = now

	obj.SetModified()

	m.emitFuShiEvent(obj.typ, obj.fushiLevel)

	return true
}

// 符石升级成功
func (m *PlayerFuShiDataManager) FushiUpLevelSucess(typ fushitypes.FuShiType) bool {
	now := global.GetGame().GetTimeService().Now()
	obj, ok := m.playerFushiObjectMap[typ]
	if !ok {
		return false
	}

	obj.fushiLevel++
	obj.updateTime = now

	obj.SetModified()

	m.emitFuShiEvent(obj.typ, obj.fushiLevel)

	return true
}

// 符石等级改变，发送事件
func (m *PlayerFuShiDataManager) emitFuShiEvent(typ fushitypes.FuShiType, level int32) {
	data := fushieventtypes.CreateFuShiLevelChangedData(typ, level)
	gameevent.Emit(fushieventtypes.FushiEventTypeLevelChanged, m.p, data)
}

func CreatePlayerFuShiDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerFuShiDataManager{
		p: p,
	}

	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerFuShiDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerFuShiDataManager))
}
