package scene

import (
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	arenatemplate "fgame/fgame/game/arena/template"
	arenatypes "fgame/fgame/game/arena/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"

	"fmt"
)

//四神兽场景
func CreateFourGodScene(mapId int32, fourGodType arenatypes.FourGodType, endTime int64) (s scene.Scene) {
	asd := &fourGodSceneData{}
	asd.SceneDelegateBase = scene.NewSceneDelegateBase()
	asd.fourGodType = fourGodType
	asd.teamMap = make(map[int64]*ArenaTeam)
	asd.currentTeamList = make([]*ArenaTeam, 0, 16)
	asd.currentTeamQueue = make([]*ArenaTeam, 0, 16)
	asd.collectMap = make(map[int64]*CollectItem)
	asd.maxTeam = arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().TeamCount
	s = createFourGodScene(mapId, endTime, asd)
	return s
}

//创建神兽
func createFourGodScene(mapId int32, endTime int64, sh scene.SceneDelegate) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeArenaShengShou {
		return nil
	}
	s = scene.CreateActivityScene(mapId, endTime, sh)
	return s
}

type FourGodSceneData interface {
	scene.SceneDelegate
	//四神类型
	GetFourGodType() arenatypes.FourGodType
	//队伍加入
	TeamJoin(to *ArenaTeam) bool
	//退出排队
	TeamLeaveQueue(teamId int64)
	//获取当前队伍信息
	GetCurrentTeamList() []*ArenaTeam
	//获取队伍排队
	GetCurrentTeamQueue() []*ArenaTeam
	//获取队伍
	GetTeam(teamId int64) *ArenaTeam
	//获取经验树
	GetExpTree() scene.NPC
	//清除采集
	ClearCollect(pl scene.Player)
	// //获取正在采集的玩家
	// GetCollectPlayerId() int64
	//获取boss信息
	GetArenaBoss() scene.NPC

	//是否在采集范围内
	IfCollectDistance(collectId int64, pos coretypes.Position) bool
	//是否可以采集
	IfCollect(collectId int64) bool
	//采集经验树
	Collect(p scene.Player, collectId int64) bool

	//获取当前排队顺序
	GetQueueTeam(teamId int64) (index int32, t *ArenaTeam)
}

type CollectItem struct {
	n                scene.NPC
	collectPlayerId  int64
	collectStartTime int64
	collectTime      int64
}

func (c *CollectItem) GetNPC() scene.NPC {
	return c.n
}

func (c *CollectItem) GetCollectPlayerId() int64 {
	return c.collectPlayerId
}

func (c *CollectItem) GetCollectStartTime() int64 {
	return c.collectStartTime
}

func (c *CollectItem) GetCollectTime() int64 {
	return c.collectTime
}

func (c *CollectItem) Collect(playerId int64) bool {
	if c.collectPlayerId != 0 {
		return false
	}
	c.collectPlayerId = playerId
	now := global.GetGame().GetTimeService().Now()
	c.collectStartTime = now
	return true
}

func (c *CollectItem) ClearCollect() {
	c.collectPlayerId = 0
}

func CreateCollectItem(n scene.NPC, collectTime int64) *CollectItem {
	c := &CollectItem{
		n:           n,
		collectTime: collectTime,
	}
	return c
}

type fourGodSceneData struct {
	*scene.SceneDelegateBase
	s           scene.Scene
	fourGodType arenatypes.FourGodType
	//所有队伍
	teamMap map[int64]*ArenaTeam
	//当前队伍
	currentTeamList []*ArenaTeam
	//当前排队队伍
	currentTeamQueue []*ArenaTeam
	//经验树
	expTree scene.NPC
	//boss信息
	boss scene.NPC
	//宝箱
	collectMap map[int64]*CollectItem
	maxTeam    int32
}

func (sd *fourGodSceneData) GetFourGodType() arenatypes.FourGodType {
	return sd.fourGodType
}

func (sd *fourGodSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

func (sd *fourGodSceneData) GetScene() scene.Scene {
	return sd.s
}

func (sd *fourGodSceneData) OnSceneTick(s scene.Scene) {
	for _, collectItem := range sd.collectMap {
		if collectItem.GetCollectPlayerId() != 0 {
			now := global.GetGame().GetTimeService().Now()
			elapse := now - collectItem.GetCollectStartTime()
			if elapse >= collectItem.GetCollectTime() {
				sd.collectDone(collectItem)
			}
		}
	}

}

func (sd *fourGodSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	gameevent.Emit(arenaeventtypes.EventTypeArenaFourGodSceneFinish, sd, nil)
}

func (sd *fourGodSceneData) OnSceneStop(s scene.Scene) {

}

func (sd *fourGodSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
	if npc.GetBiologyTemplate().GetBiologyScriptType() == scenetypes.BiologyScriptTypeArenaExpTree {
		sd.expTree = npc
		collectTime := int64(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().TreeGetTime)
		sd.collectMap[npc.GetId()] = CreateCollectItem(npc, collectTime)
	}
	if npc.GetBiologyTemplate().GetBiologyScriptType() == scenetypes.BiologyScriptTypeArenaShengShou {
		sd.boss = npc
	}
	if npc.GetBiologyTemplate().GetBiologyScriptType() == scenetypes.BiologyScriptTypeArenaTreasure {
		collectTime := int64(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().EquipBoxTime)
		sd.collectMap[npc.GetId()] = CreateCollectItem(npc, collectTime)
	}
}

func (sd *fourGodSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {
	if npc.GetBiologyTemplate().GetBiologyScriptType() == scenetypes.BiologyScriptTypeArenaExpTree {
		delete(sd.collectMap, npc.GetId())
	}

	if npc.GetBiologyTemplate().GetBiologyScriptType() == scenetypes.BiologyScriptTypeArenaTreasure {
		delete(sd.collectMap, npc.GetId())
	}
}

func (sd *fourGodSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {

}

func (sd *fourGodSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

func (sd *fourGodSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}

func (sd *fourGodSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

func (sd *fourGodSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {

}

func (sd *fourGodSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	gameevent.Emit(arenaeventtypes.EventTypeArenaFourGodScenePlayerExit, p, active)

	sd.ClearCollect(p)

	teamId := p.GetTeamId()
	if teamId == 0 {
		return
	}
	found := false
	//判断是否所有成员退出了
	for _, tempP := range sd.s.GetAllPlayers() {
		if tempP.GetTeamId() == teamId {
			found = true
			break
		}
	}
	//所有成员都退出了
	if !found {
		sd.removeTeamFromScene(teamId)
	}
}
func (sd *fourGodSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}
func (sd *fourGodSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	gameevent.Emit(arenaeventtypes.EventTypeArenaFourGodScenePlayerEnter, sd, p)
}

func (sd *fourGodSceneData) OnScenePlayerGetItem(s scene.Scene, p scene.Player, itemData *droptemplate.DropItemData) {

}
func (sd *fourGodSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
}

func (sd *fourGodSceneData) OnSceneRefreshGroup(s scene.Scene, group int32) {
}

func (sd *fourGodSceneData) TeamJoin(to *ArenaTeam) bool {
	//判断是否是已经在里面了
	teamId := to.GetTeam().GetTeamId()
	originTeam := sd.getTeam(teamId)
	if originTeam != nil {
		return false
	}

	sd.addTeam(to)
	return true
}

func (sd *fourGodSceneData) addTeam(to *ArenaTeam) {
	teamId := to.GetTeam().GetTeamId()
	sd.teamMap[teamId] = to
	//判断是否要排队
	if sd.isFull() {
		sd.addTeamInQueue(to)
		return
	}
	sd.addTeamInScene(to)
}

//加入队列
func (sd *fourGodSceneData) addTeamInQueue(to *ArenaTeam) {
	sd.currentTeamQueue = append(sd.currentTeamQueue, to)
	//发送事件,加入队列
	gameevent.Emit(arenaeventtypes.EventTypeArenaFourGodSceneTeamQueue, sd, to)
}

func (sd *fourGodSceneData) addTeamInScene(to *ArenaTeam) {
	sd.currentTeamList = append(sd.currentTeamList, to)
	//发送事件,加入战场
	gameevent.Emit(arenaeventtypes.EventTypeArenaFourGodSceneTeamJoin, sd, to)
}

//队伍离开排队
func (sd *fourGodSceneData) TeamLeaveQueue(teamId int64) {
	to, _ := sd.getTeamFromQueue(teamId)
	if to == nil {
		return
	}

	sd.removeTeamFromQueue(teamId)
}

//从游戏队伍中移除
func (sd *fourGodSceneData) removeTeamFromScene(teamId int64) {
	t, pos := sd.getTeamFromScene(teamId)
	if t == nil {
		return
	}
	delete(sd.teamMap, teamId)
	sd.currentTeamList = append(sd.currentTeamList[:pos], sd.currentTeamList[pos+1:]...)
	//队伍退出
	gameevent.Emit(arenaeventtypes.EventTypeArenaFourGodSceneTeamLeave, sd, teamId)

	if len(sd.currentTeamQueue) > 0 {
		addTeam := sd.currentTeamQueue[0]
		sd.currentTeamQueue = sd.currentTeamQueue[1:]
		sd.addTeamInScene(addTeam)
		gameevent.Emit(arenaeventtypes.EventTypeArenaFourGodSceneTeamQueueChanged, sd, nil)
	}

}

//从排队队伍中移除
func (sd *fourGodSceneData) removeTeamFromQueue(teamId int64) {
	t, pos := sd.getTeamFromQueue(teamId)
	if t == nil {
		return
	}

	delete(sd.teamMap, teamId)
	sd.currentTeamQueue = append(sd.currentTeamQueue[:pos], sd.currentTeamQueue[pos+1:]...)
	gameevent.Emit(arenaeventtypes.EventTypeArenaFourGodSceneTeamQueueCancel, sd, t)
}

//获取队伍
func (sd *fourGodSceneData) getTeam(teamId int64) *ArenaTeam {
	teamObj, ok := sd.teamMap[teamId]
	if !ok {
		return nil
	}
	return teamObj
}

//获取游戏队伍
func (sd *fourGodSceneData) getTeamFromScene(teamId int64) (to *ArenaTeam, pos int32) {
	for index, t := range sd.currentTeamList {
		if t.GetTeam().GetTeamId() == teamId {
			to = t
			pos = int32(index)
			return
		}
	}
	return nil, -1
}

//获取队列队伍
func (sd *fourGodSceneData) getTeamFromQueue(teamId int64) (to *ArenaTeam, pos int32) {
	for index, t := range sd.currentTeamQueue {
		if t.GetTeam().GetTeamId() == teamId {
			to = t
			pos = int32(index)
			return
		}
	}
	return nil, -1
}

//是否满了
func (sd *fourGodSceneData) isFull() bool {
	return len(sd.currentTeamList) >= int(sd.maxTeam)
}

func (sd *fourGodSceneData) GetCurrentTeamList() []*ArenaTeam {
	return sd.currentTeamList
}

func (sd *fourGodSceneData) GetCurrentTeamQueue() []*ArenaTeam {
	return sd.currentTeamQueue
}

func (sd *fourGodSceneData) GetTeam(teamId int64) *ArenaTeam {
	for _, t := range sd.currentTeamList {
		if t.GetTeam().GetTeamId() == teamId {
			return t
		}
	}
	return nil
}

func (sd *fourGodSceneData) GetQueueTeam(teamId int64) (index int32, t *ArenaTeam) {
	for i, t := range sd.currentTeamQueue {
		if t.GetTeam().GetTeamId() == teamId {
			return int32(i), t
		}
	}
	return -1, nil
}

func (sd *fourGodSceneData) GetExpTree() scene.NPC {
	return sd.expTree
}

func (sd *fourGodSceneData) GetArenaBoss() scene.NPC {
	return sd.boss
}

func (sd *fourGodSceneData) IfCollectDistance(collectId int64, pos coretypes.Position) (flag bool) {
	collectItem, ok := sd.collectMap[collectId]
	if !ok {
		return
	}
	npcPos := collectItem.GetNPC().GetPosition()
	distance := coreutils.Distance(npcPos, pos)

	collectDistance := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCollectDistance)) / float64(1000)
	if distance > float64(collectDistance) {
		return
	}
	flag = true
	return
}

func (sd *fourGodSceneData) IfCollect(collectId int64) bool {
	collectItem, ok := sd.collectMap[collectId]
	if !ok {
		fmt.Println("采集物不存在")
		return false
	}
	if collectItem.GetCollectPlayerId() != 0 {
		return false
	}
	return true
}

func (sd *fourGodSceneData) Collect(pl scene.Player, collectId int64) bool {
	collectItem, ok := sd.collectMap[collectId]
	if !ok {
		return false
	}
	if !ok {
		return false
	}
	if collectItem.GetCollectPlayerId() != 0 {
		return false
	}
	flag := collectItem.Collect(pl.GetId())
	if !flag {
		return false
	}
	gameevent.Emit(arenaeventtypes.EventTypeArenaFourGodSceneCollecting, sd, collectItem)
	return true
}

//清除读条
func (sd *fourGodSceneData) ClearCollect(pl scene.Player) {
	if sd.s.State() == scene.SceneStateFinish || sd.s.State() == scene.SceneStateStopped {
		return
	}

	for _, box := range sd.collectMap {
		if box.GetCollectPlayerId() == pl.GetId() {
			gameevent.Emit(arenaeventtypes.EventTypeArenaFourGodSceneCollectStop, sd, box)
			box.ClearCollect()
		}
	}
}

func (sd *fourGodSceneData) collectDone(collectItem *CollectItem) {
	if collectItem.GetCollectPlayerId() == 0 {
		return
	}
	gameevent.Emit(arenaeventtypes.EventTypeArenaFourGodSceneCollect, sd, collectItem)
	n := collectItem.GetNPC()
	// n.CostHP(n.GetHP(), collectItem.GetCollectPlayerId())
	n.Recycle(collectItem.GetCollectPlayerId())
	collectItem.ClearCollect()
}
