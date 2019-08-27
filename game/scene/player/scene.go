package player

import (
	"fgame/fgame/game/constant/constant"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"

	"fgame/fgame/pkg/idutil"

	"fgame/fgame/game/scene/dao"
)

//玩家场景数据管理器
type PlayerSceneDataManager struct {
	p                 player.Player
	playerSceneObject *PlayerSceneObject
}

func (m *PlayerSceneDataManager) Player() player.Player {
	return m.p
}

func (m *PlayerSceneDataManager) GetPlayerScene() *PlayerSceneObject {
	return m.playerSceneObject
}

//加载
func (m *PlayerSceneDataManager) Load() (err error) {
	pse, err := dao.GetSceneDao().GetSceneEntity(m.p.GetId())
	if err != nil {
		return
	}
	if pse == nil {
		m.initPlayerSceneObject()
	} else {
		m.playerSceneObject = NewPlayerSceneObject(m.p)
		m.playerSceneObject.FromEntity(pse)
	}

	return nil
}

func (m *PlayerSceneDataManager) AfterLoad() (err error) {
	return nil
}

func (m *PlayerSceneDataManager) Heartbeat() {

}

//第一次初始化
func (m *PlayerSceneDataManager) initPlayerSceneObject() {
	pso := NewPlayerSceneObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	role := m.Player().GetRole()
	sex := m.Player().GetSex()
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(role, sex)
	bornPos := playerCreateTemplate.GetBornPos()
	pso.Id = id
	//生成id
	pso.PlayerId = m.p.GetId()
	pso.MapId = playerCreateTemplate.MapId
	pso.PosX = float64(bornPos.X)
	pso.PosY = float64(bornPos.Y)
	pso.PosZ = float64(bornPos.Z)
	pso.LastMapId = 0
	pso.LastPosX = 0.0
	pso.LastPosY = 0.0
	pso.LastPosZ = 0.0
	pso.CreateTime = now
	m.playerSceneObject = pso
	pso.SetModified()
}

//保存数据
func (m *PlayerSceneDataManager) Save() {
	m.playerSceneObject.MapId = m.p.GetMapId()
	m.playerSceneObject.SceneId = m.p.GetSceneId()
	pos := m.p.GetPosition()
	m.playerSceneObject.PosX = pos.X
	m.playerSceneObject.PosY = pos.Y
	m.playerSceneObject.PosZ = pos.Z
	m.playerSceneObject.LastMapId = m.p.GetLastMapId()
	m.playerSceneObject.LastSceneId = m.p.GetLastSceneId()
	lastPos := m.p.GetLastPos()
	m.playerSceneObject.LastPosX = lastPos.X
	m.playerSceneObject.LastPosZ = lastPos.Z
	m.playerSceneObject.LastPosY = lastPos.Y
	m.playerSceneObject.SetModified()
}

func CreatePlayerSceneDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerSceneDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerSceneDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerSceneDataManager))
}
