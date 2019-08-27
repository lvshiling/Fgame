package scene

import (
	teamcopyeventtypes "fgame/fgame/cross/teamcopy/event/types"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

type TeamCopySceneData interface {
	scene.SceneDelegate
	//队伍
	GetTeamObj() *TeamObject
	GetMember(p scene.Player) *TeamMemberObject
	GetStartTime() int64
	GetReliveTime(p scene.Player) (reliveTime int32)
	UpdateDamage(playerId int64, damage int64)
	AddReliveTime(p scene.Player)
	IfAllLevel() bool
	GetDamage(p scene.Player) int64
	FinishScene()
}

//组队副本战场数据
type teamCopySceneData struct {
	*scene.SceneDelegateBase
	s       scene.Scene
	teamObj *TeamObject
	isEnd   bool
}

func CreateTeamCopySceneData(teamObj *TeamObject) TeamCopySceneData {
	csd := &teamCopySceneData{
		teamObj: teamObj,
	}
	csd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return csd
}

func (sd *teamCopySceneData) GetScene() (s scene.Scene) {
	return sd.s
}

//场景开始
func (sd *teamCopySceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

//刷怪
func (sd *teamCopySceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//场景心跳
func (sd *teamCopySceneData) OnSceneTick(s scene.Scene) {
	sd.Heartbeat()
}

//怪物死亡
func (sd *teamCopySceneData) OnSceneBiologyAllDead(s scene.Scene) {
	sd.isEnd = true
	sd.s.Finish(true)

}

//生物进入
func (sd *teamCopySceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {

}

func (sd *teamCopySceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *teamCopySceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("teamcopy:组队副本战场应该是同一个场景"))
	}

}

//生物重生
func (sd *teamCopySceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("teamcopy:组队副本战场应该是同一个场景"))
	}
}

//玩家复活
func (sd *teamCopySceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("teamcopy:组队副本战场应该是同一个场景"))
	}
}

//玩家死亡
func (sd *teamCopySceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("teamcopy:组队副本战场应该是同一个场景"))
	}

}

func (sd *teamCopySceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *teamCopySceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("teamcopy:组队副本战场应该是同一个场景"))
	}
	//切换pk模式
	p.SwitchPkState(pktypes.PkStateGroup, pktypes.PkCommonCampDefault)
	if !p.IsRobot() {
		gameevent.Emit(teamcopyeventtypes.EvnetTypeTeamCopySceneEnter, p, sd)
	}
}

//玩家退出
func (sd *teamCopySceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("teamcopy:组队副本战场应该是同一个场景"))
	}
}

//场景完成
func (sd *teamCopySceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("teamcopy:组队副本战场应该是同一个场景"))
	}
	sd.isEnd = true
	gameevent.Emit(teamcopyeventtypes.EventTypeTeamCopySceneFinish, sd, success)
}

//场景退出了
func (sd *teamCopySceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("teamcopy:组队副本战场应该是同一个场景"))
	}
}

//场景获取物品
func (sd *teamCopySceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("teamcopy:组队副本战场应该是同一个场景"))
	}
}

//玩家获得经验
func (sd *teamCopySceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("teamcopy:组队副本战场应该是同一个场景"))
	}
}

//心跳
func (sd *teamCopySceneData) Heartbeat() {

}

//获取玩家
func (sd *teamCopySceneData) GetMember(p scene.Player) *TeamMemberObject {
	for _, mem := range sd.teamObj.GetMemberList() {
		if mem.GetPlayerId() == p.GetId() {
			return mem
		}
	}
	return nil
}

//队伍
func (sd *teamCopySceneData) GetTeamObj() *TeamObject {
	return sd.teamObj
}

//开始时间
func (sd *teamCopySceneData) GetStartTime() int64 {
	return sd.s.GetStartTime()
}

//获取复活次数
func (sd *teamCopySceneData) GetReliveTime(p scene.Player) (reliveTime int32) {
	mem := sd.GetMember(p)
	if mem == nil {
		return
	}
	return mem.GetReliveTime()
}

//获取玩家伤害
func (sd *teamCopySceneData) GetDamage(p scene.Player) (damage int64) {
	mem := sd.GetMember(p)
	if mem == nil {
		return
	}
	return mem.GetDamage()
}

//更新复活次数
func (sd *teamCopySceneData) AddReliveTime(p scene.Player) {
	mem := sd.GetMember(p)
	if mem == nil {
		return
	}
	mem.AddReliveTime()
}

//更新伤害
func (sd *teamCopySceneData) UpdateDamage(playerId int64, damage int64) {
	if damage <= 0 {
		return
	}
	pl := sd.GetScene().GetPlayer(playerId)
	if pl == nil {
		return
	}

	mem := sd.GetMember(pl)
	if mem == nil {
		return
	}

	mem.AddDamage(damage)
	gameevent.Emit(teamcopyeventtypes.EventTypeTeamCopySceneDamageChanged, pl, sd)
}

func (sd *teamCopySceneData) FinishScene() {
	if !sd.isEnd {
		sd.s.Finish(false)
	}
}

func (sd *teamCopySceneData) IfAllLevel() (flag bool) {
	for _, mem := range sd.teamObj.GetMemberList() {
		if mem.IsRobot() {
			continue
		}
		status := mem.GetStatus()
		switch status {
		case MemberStatusOffline,
			MemberStatusOnline:
			return
		}
	}
	flag = true
	return
}
