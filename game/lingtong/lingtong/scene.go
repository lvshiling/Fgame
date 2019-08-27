package lingtong

import (
	"fgame/fgame/core/aoi"
	coretypes "fgame/fgame/core/types"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//场景管理器
type LingTongSceneManager struct {
	*scene.SceneObjectBase
	p scene.LingTong
	//加载过的玩家
	loadedPlayers map[int64]scene.Player
}

const (
	defaultAngle = 0
)

func (m *LingTongSceneManager) EnterScene(s scene.Scene) {
	//设置加载过的玩家
	m.loadedPlayers = make(map[int64]scene.Player)
	m.SceneObjectBase.EnterScene(s)
	//发送事件
	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongEnterScene, m.p, nil)

}

func (m *LingTongSceneManager) ExitScene(active bool) {
	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongExitScene, m.p, active)
	//TODO 清除技能管理器
	m.p.ClearAllSkillAction()
	//暂停移动
	m.p.PauseMove()
	m.SceneObjectBase.ExitScene(active)
}

//重写进入aoi
func (m *LingTongSceneManager) OnEnterAOI(other aoi.AOI) {
	m.SceneObjectBase.OnEnterAOI(other)

	otherPlayer, ok := other.(scene.Player)
	if ok {
		if !otherPlayer.IsRobot() {
			m.loadedPlayers[otherPlayer.GetId()] = otherPlayer
		}
	}

}

//重写进入aoi
func (m *LingTongSceneManager) OnLeaveAOI(other aoi.AOI, complete bool) {

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

}

func (m *LingTongSceneManager) GetLoadedPlayers() map[int64]scene.Player {
	return m.loadedPlayers
}

func (m *LingTongSceneManager) RemoveLoadedPlayer(id int64) {
	delete(m.loadedPlayers, id)
}

func CreateLingTongSceneManager(p scene.LingTong, pos coretypes.Position, angle float64) *LingTongSceneManager {
	m := &LingTongSceneManager{
		p: p,
	}
	m.loadedPlayers = make(map[int64]scene.Player)
	m.SceneObjectBase = scene.NewSceneObjectBase(p, pos, angle, scenetypes.BiologyTypePlayer)
	return m
}
