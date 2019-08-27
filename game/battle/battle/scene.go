package battle

import (
	"fgame/fgame/core/aoi"
	coretypes "fgame/fgame/core/types"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenecommon "fgame/fgame/game/scene/common"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
)

//场景管理器
type PlayerSceneManager struct {
	*scene.SceneObjectBase
	p           scene.Player
	scene       scene.Scene
	mapId       int32
	sceneId     int64
	pos         coretypes.Position
	lastMapId   int32
	lastSceneId int64
	lastPos     coretypes.Position
	enterPos    coretypes.Position
	//加载过的玩家
	loadedPlayers map[int64]scene.Player
	//离开的
	leaveNeighbors map[int64]aoi.AOI
	//进入的
	enterNeighbors map[int64]aoi.AOI
}

const (
	defaultAngle = 0
)

func (m *PlayerSceneManager) EnterScene(s scene.Scene) {

	//设置加载过的玩家
	m.loadedPlayers = make(map[int64]scene.Player)
	sb := scene.NewSceneObjectBase(m.p, m.pos, defaultAngle, scenetypes.BiologyTypePlayer)
	m.SceneObjectBase = sb
	m.scene = s
	m.SceneObjectBase.EnterScene(s)
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(m.mapId)
	if mapTemplate != nil {
		if mapTemplate.IsWorld() {
			m.lastMapId = m.mapId
			m.lastSceneId = m.sceneId
			m.lastPos = m.pos
		}
	}
	m.mapId = m.scene.MapId()
	m.sceneId = m.scene.Id()
	m.pos = m.enterPos
	m.SetPosition(m.pos)
	angle := float64(s.MapTemplate().Rotate)
	m.SetAngle(angle)

	//重置敌人
	m.p.ResetEnemy()
	//发送事件
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerEnterScene, m.p, nil)

}

func (m *PlayerSceneManager) ExitScene(active bool) {
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerExitScene, m.p, active)
	m.SceneObjectBase.ExitScene(active)
	//TODO 清除技能管理器
	m.p.ClearAllSkillAction()
	m.p.ClearPvpBattle()
	//暂停移动
	m.p.PauseMove()
	//TODO:zrc 退出场景死亡active是干嘛用的
	//用户死亡了 需要重生
	if m.p.IsDead() {
		m.p.Reborn(m.p.GetPosition())
	}
	//发送事件
	m.scene = nil
	m.loadedPlayers = nil
}

func (m *PlayerSceneManager) GetEnterNeighborsAndClear() map[int64]aoi.AOI {
	r := m.enterNeighbors
	//TODO 清空
	m.enterNeighbors = make(map[int64]aoi.AOI)
	return r
}

func (m *PlayerSceneManager) GetLeaveNeighborsAndClear() map[int64]aoi.AOI {
	r := m.leaveNeighbors
	//TODO 清空
	m.leaveNeighbors = make(map[int64]aoi.AOI)
	return r
}

//重写进入aoi
func (m *PlayerSceneManager) OnEnterAOI(other aoi.AOI) {
	m.SceneObjectBase.OnEnterAOI(other)
	_, exist := m.leaveNeighbors[other.GetId()]
	if exist {
		delete(m.leaveNeighbors, other.GetId())
	} else {
		m.enterNeighbors[other.GetId()] = other
	}
	otherPlayer, ok := other.(scene.Player)
	if ok {
		//记录真实玩家
		if !otherPlayer.IsRobot() {
			m.loadedPlayers[otherPlayer.GetId()] = otherPlayer
		}
	}

	bo, ok := other.(scene.BattleObject)
	if !ok {
		return
	}
	m.p.EnterBattle(bo)
}

//重写进入aoi
func (m *PlayerSceneManager) OnLeaveAOI(other aoi.AOI, complete bool) {

	defer m.SceneObjectBase.OnLeaveAOI(other, complete)

	switch otherObj := other.(type) {
	case scene.Player:
		if complete {
			if !otherObj.IsRobot() {
				delete(m.loadedPlayers, otherObj.GetId())
			}
		}
		break
	}
	if complete {
		switch other.(type) {
		case scene.Player,
			scene.LingTong:
			//不需要同步了
			delete(m.enterNeighbors, other.GetId())
			delete(m.leaveNeighbors, other.GetId())
			break
		default:
			_, exist := m.enterNeighbors[other.GetId()]
			if exist {
				delete(m.enterNeighbors, other.GetId())
			} else {
				m.leaveNeighbors[other.GetId()] = other
			}
			break
		}
	} else {
		_, exist := m.enterNeighbors[other.GetId()]
		if exist {
			delete(m.enterNeighbors, other.GetId())
		} else {
			m.leaveNeighbors[other.GetId()] = other
		}
	}

	// _, exist := m.enterNeighbors[other.GetId()]
	// if exist {
	// 	delete(m.enterNeighbors, other.GetId())
	// } else {
	// 	m.leaveNeighbors[other.GetId()] = other
	// }
	bo, ok := other.(scene.BattleObject)
	if !ok {
		return
	}
	m.p.ExitBattle(bo)
}

//别的对象移动
func (m *PlayerSceneManager) OnMove(bo scene.BattleObject, pos coretypes.Position, angle float64) {
	m.p.EnterBattle(bo)
}

//重生
func (m *PlayerSceneManager) OnReborn(bo scene.BattleObject) {
	m.p.EnterBattle(bo)
}

//死亡
func (m *PlayerSceneManager) OnDead(bo scene.BattleObject) {
	m.p.ExitBattle(bo)
}

func (m *PlayerSceneManager) Move(pos coretypes.Position, angle float64) {
	m.pos = pos
	m.SetAngle(angle)
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerMove, m.p, nil)
	//发送事件
}

//退回上一个场景
func (m *PlayerSceneManager) BackLastScene() {
	//发送事件
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerBackLastScene, m.p, nil)

}

//被攻击位移
func (m *PlayerSceneManager) AttackedMove(pos coretypes.Position, angle float64, moveSpeed float64, stopTime float64) {
	m.Move(pos, angle)
}

func (m *PlayerSceneManager) GetLoadedPlayers() map[int64]scene.Player {
	return m.loadedPlayers
}

func (m *PlayerSceneManager) RemoveLoadedPlayer(id int64) {
	delete(m.loadedPlayers, id)
}

func (m *PlayerSceneManager) GetLastMapId() int32 {
	return m.lastMapId
}

func (m *PlayerSceneManager) GetLastSceneId() int64 {
	return m.lastSceneId
}

func (m *PlayerSceneManager) GetLastPos() coretypes.Position {
	return m.lastPos
}

func (m *PlayerSceneManager) GetMapId() int32 {
	return m.mapId
}

func (m *PlayerSceneManager) GetSceneId() int64 {
	return m.sceneId
}

func (m *PlayerSceneManager) GetPos() coretypes.Position {
	return m.pos
}

func (p *PlayerSceneManager) GetScene() scene.Scene {
	return p.scene
}

func (p *PlayerSceneManager) SetEnterPos(pos coretypes.Position) {
	p.enterPos = pos
}

func (p *PlayerSceneManager) GetEnterPosition() coretypes.Position {
	return p.enterPos
}

func CreatePlayerSceneManagerWithObject(p scene.Player, so scenecommon.SceneObject) *PlayerSceneManager {
	m := &PlayerSceneManager{
		p:           p,
		mapId:       so.GetMapId(),
		sceneId:     so.GetSceneId(),
		pos:         so.GetPos(),
		lastMapId:   so.GetLastMapId(),
		lastSceneId: so.GetLastSceneId(),
		lastPos:     so.GetLastPos(),
	}
	m.loadedPlayers = make(map[int64]scene.Player)
	m.leaveNeighbors = make(map[int64]aoi.AOI)
	m.enterNeighbors = make(map[int64]aoi.AOI)
	return m
}

func CreatePlayerSceneManager(p scene.Player) *PlayerSceneManager {
	m := &PlayerSceneManager{
		p: p,
	}
	m.loadedPlayers = make(map[int64]scene.Player)
	m.leaveNeighbors = make(map[int64]aoi.AOI)
	m.enterNeighbors = make(map[int64]aoi.AOI)
	return m
}
