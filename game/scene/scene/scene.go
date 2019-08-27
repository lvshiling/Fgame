package scene

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/aoi"
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/nav"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/common/common"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"
	"fgame/fgame/game/global"
	robottypes "fgame/fgame/game/robot/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

type SceneState int32

const (
	SceneStateInit SceneState = iota
	SceneStateStart
	SceneStateFinish
	SceneStateStopped
)

var (
	sceneStateMap = map[SceneState]string{
		SceneStateInit:    "初始化",
		SceneStateStart:   "开始",
		SceneStateFinish:  "完成",
		SceneStateStopped: "停止",
	}
)

func (s SceneState) String() string {
	return sceneStateMap[s]
}

//TODO 限制场景人数
//场景接口
type Scene interface {
	SceneRankManager
	//场景唯一id
	Id() int64
	SceneDelegate() SceneDelegate
	State() SceneState
	Success() bool
	//场景地图id
	MapId() int32
	//获取场景模板数据
	MapTemplate() *gametemplate.MapTemplate
	//获取地图
	GetWorld() *nav.World
	Stop(active bool, complete bool)
	AsyncStop()
	Post(msg message.Message) bool
	AddSceneObject(so SceneObject)
	OnDead(bo BattleObject)
	OnReborn(bo BattleObject, pos coretypes.Position)

	OnPlayerGetItem(pl Player, itemData *droptemplate.DropItemData)
	OnPlayerGetExp(pl Player, exp int64)
	Move(so BattleObject, pos coretypes.Position)
	RemoveSceneObject(so SceneObject, active bool)
	Finish(success bool)
	GetStartTime() int64
	GetEndTime() int64
	GetFinishTime() int64
	IsFinish() bool
	GetCurrentGroup() int32
	//刷新怪物
	RefreshBiology(groupId int32)
	//获取所有物品
	GetAllItems() map[int64]DropItem
	//获取物品
	GetItem(itemId int64) DropItem
	//获取场景对象
	GetSceneObject(id int64) SceneObject
	//获取玩家
	GetPlayer(id int64) Player
	//获取灵童
	GetLingTong(id int64) LingTong
	//获取所有npc
	GetAllNPCS() map[int64]NPC
	//获取npc列表
	GetNPCListByBiology(biologyId int32) []NPC
	//获取NPC
	GetNPCS(biologyTypes scenetypes.BiologyScriptType) map[int64]NPC
	//获取npc根据部怪表
	GetNPCByIdx(idx int64) NPC
	//获取剩余npc数量
	GetNumOfNPC() int32
	//获取所有玩家
	GetAllPlayers() map[int64]Player
	//广播消息
	BroadcastMsg(msg proto.Message)
	//获取经验加成
	GetExpAddPercent(p Player) int32

	GetNumOfRobot() int32                //接口只取任务类型机器人
	GetAllQuestRobots() map[int64]Player //接口只取任务类型机器人
	GetAllRobots() map[int64]Player
	GetNumOfAllRobot() int32 //记录所有机器人
	GetLastRobotTime() int64
	SetLastRobotTime(t int64)
	String() string
}

type SceneDelegate interface {
	GetScene() Scene
	//场景开始
	OnSceneStart(s Scene)
	//场景心跳
	OnSceneTick(s Scene)
	//场景完成
	OnSceneFinish(s Scene, success bool, useTime int64)
	//场景停止
	OnSceneStop(s Scene)
	//生物进入
	OnSceneBiologyEnter(s Scene, npc NPC)
	OnSceneBiologyExit(s Scene, npc NPC)
	//场景生物死亡
	OnSceneBiologyDead(s Scene, npc NPC)
	//生物重生
	OnSceneBiologyReborn(s Scene, npc NPC)
	//场景生物全部死亡
	OnSceneBiologyAllDead(s Scene)
	//玩家复活
	OnScenePlayerReborn(s Scene, p Player)
	//玩家死亡
	OnScenePlayerDead(s Scene, p Player)
	//玩家退出
	OnScenePlayerExit(s Scene, p Player, active bool)

	OnScenePlayerBeforeEnter(s Scene, p Player)
	//玩家进入
	OnScenePlayerEnter(s Scene, p Player)
	//玩家获取物品
	OnScenePlayerGetItem(s Scene, p Player, itemData *droptemplate.DropItemData)
	//玩家获得经验
	OnScenePlayerGetExp(s Scene, p Player, num int64)
	//场景刷怪
	OnSceneRefreshGroup(s Scene, group int32)
	//获取经验加成
	GetExpAddPercent(p Player) int32
}

type SceneDelegateBase struct {
	s Scene
}

func (sd *SceneDelegateBase) GetScene() Scene {
	return sd.s
}

func (sd *SceneDelegateBase) GetExpAddPercent(p Player) int32 {
	return 0
}
func (sd *SceneDelegateBase) OnSceneStart(s Scene) {
	sd.s = s
	return
}

func (sd *SceneDelegateBase) OnSceneTick(s Scene) {

	// //TODO xzk27 封装到scene里面
	// func propertyTickRew(s scene.Scene) {
	// 	warTemplate := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiWarTemp()
	// 	for _, spl := range s.GetAllPlayers() {
	// 		tickRewData := spl.GetActivityTickRewData()
	// 		if tickRewData == nil {
	// 			continue
	// 		}

	// 		now := global.GetGame().GetTimeService().Now()
	// 		enterTime := tickRewData.GetEnterTime()
	// 		lastRewTime := tickRewData.GetLastRewTime()
	// 		lastRewSpecialTime := tickRewData.GetLastRewSpecialTime()

	// 		//定时奖励
	// 		resMap := make(map[int32]int32)
	// 		if warTemplate.IfTickRew(now, enterTime, lastRewTime) {
	// 			silver := int64(0)
	// 			exp := int64(0)
	// 			expPoint := int32(0)
	// 			if s.MapTemplate().IsSafe(spl.GetPos()) {
	// 				silver = warTemplate.RewSilverSafeArea
	// 				exp = warTemplate.RewExpSafeArea
	// 				expPoint = warTemplate.RewExpPointSafeArea
	// 			} else {
	// 				silver = warTemplate.RewSilver
	// 				exp = warTemplate.RewExp
	// 				expPoint = warTemplate.RewExpPoint
	// 			}
	// 			exp += propertylogic.ExpPointConvertExp(expPoint, spl.GetLevel())

	// 			resMap[constanttypes.SilverItem] = int32(silver)
	// 			resMap[constanttypes.ExpItem] = int32(exp)
	// 		}

	// 		//
	// 		specialResMap := make(map[int32]int32)
	// 		if warTemplate.IfTickRew(now, enterTime, lastRewSpecialTime) {
	// 			if !s.MapTemplate().IsSafe(spl.GetPos()) {
	// 				specialResMap[constanttypes.ChuangShiJifen] = warTemplate.RewJiFen
	// 			}
	// 		}

	// 		if len(resMap) > 0 || len(specialResMap) > 0 {
	// 			spl.AddActivityTickRew(resMap, specialResMap)
	// 		}
	// 	}
	// }
	return
}
func (sd *SceneDelegateBase) OnSceneFinish(s Scene, success bool, useTime int64) {
	return
}
func (sd *SceneDelegateBase) OnSceneStop(s Scene) {
	return
}
func (sd *SceneDelegateBase) OnSceneBiologyEnter(s Scene, npc NPC) {
	return
}
func (sd *SceneDelegateBase) OnSceneBiologyExit(s Scene, npc NPC) {
	return
}
func (sd *SceneDelegateBase) OnSceneBiologyDead(s Scene, npc NPC) {
	return
}

func (sd *SceneDelegateBase) OnSceneBiologyReborn(s Scene, npc NPC) {
	return
}

func (sd *SceneDelegateBase) OnSceneBiologyAllDead(s Scene) {
	return
}

func (sd *SceneDelegateBase) OnScenePlayerReborn(s Scene, p Player) {
	return
}

func (sd *SceneDelegateBase) OnScenePlayerDead(s Scene, p Player) {
	return
}
func (sd *SceneDelegateBase) OnScenePlayerExit(s Scene, p Player, active bool) {
	return
}

func (sd *SceneDelegateBase) OnScenePlayerBeforeEnter(s Scene, p Player) {
	return
}

func (sd *SceneDelegateBase) OnScenePlayerEnter(s Scene, p Player) {
	return
}
func (sd *SceneDelegateBase) OnScenePlayerGetItem(s Scene, p Player, itemData *droptemplate.DropItemData) {
	return
}
func (sd *SceneDelegateBase) OnScenePlayerGetExp(s Scene, p Player, num int64) {
	return
}
func (sd *SceneDelegateBase) OnSceneRefreshGroup(s Scene, group int32) {
	return
}

func NewSceneDelegateBase() *SceneDelegateBase {
	return &SceneDelegateBase{}
}

//场景基础结构和功能
type scene struct {
	m sync.Mutex
	SceneRankManager
	//场景状态
	state SceneState
	//结果
	success bool
	//唯一id
	id int64
	//aoi
	aoiManager aoi.AOIManager
	//寻路地图
	world *nav.World
	//地图模板
	mapTemplate *gametemplate.MapTemplate
	//对象管理
	sceneObjects map[int64]SceneObject
	//玩家管理
	players map[int64]Player
	//灵童管理
	lingTongs map[int64]LingTong
	//npc管理
	npcs    map[int64]NPC
	idxNpcs map[int64]NPC
	//npc分类
	npcMapOfMap map[scenetypes.BiologyScriptType]map[int64]NPC
	//npc分类
	npcListOfMap map[int32][]NPC
	//物品管理
	dropItems map[int64]DropItem
	//开始时间
	startTime int64
	//完成时间
	finishTime int64
	//结束时间 活动使用
	endTime int64
	//场景handler
	sceneDelegate SceneDelegate
	// 定时器
	heartbeatTimer *time.Timer
	//消息
	msgs chan message.Message
	//当前波数
	currentGroup int32

	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
	//机器人列表
	robotMap      map[robottypes.RobotType]map[int64]Player
	allRobotMap   map[int64]Player
	lastRobotTime int64
	done          chan struct{}
	lastLogTime   int64
}

func (sb *scene) Id() int64 {
	return sb.id
}

func (sb *scene) State() SceneState {
	return sb.state
}

func (sb *scene) SceneDelegate() SceneDelegate {
	return sb.sceneDelegate
}

func (sb *scene) Success() bool {
	return sb.success
}

func (sb *scene) MapTemplate() *gametemplate.MapTemplate {
	return sb.mapTemplate
}

func (sb *scene) GetWorld() *nav.World {
	return sb.world
}

func (sb *scene) MapId() int32 {
	return int32(sb.mapTemplate.Id)
}

func (sb *scene) GetStartTime() int64 {
	return sb.startTime
}

func (sb *scene) GetEndTime() int64 {
	return sb.endTime
}

func (sb *scene) GetFinishTime() int64 {
	return sb.finishTime
}

func (sb *scene) IsFinish() bool {
	return sb.state == SceneStateFinish
}

func (sb *scene) GetCurrentGroup() int32 {
	return sb.currentGroup
}

func (sb *scene) GetAllNPCS() map[int64]NPC {
	return sb.npcs
}
func (sb *scene) GetAllItems() map[int64]DropItem {
	return sb.dropItems
}
func (sb *scene) GetNPCS(npcType scenetypes.BiologyScriptType) map[int64]NPC {
	return sb.npcMapOfMap[npcType]
}
func (sb *scene) GetNPCListByBiology(biologyId int32) []NPC {
	return sb.npcListOfMap[biologyId]
}

func (sb *scene) GetNPCByIdx(idx int64) NPC {
	return sb.idxNpcs[idx]
}

func (sb *scene) GetAllPlayers() map[int64]Player {
	return sb.players
}

func (sb *scene) GetAllQuestRobots() map[int64]Player {
	return sb.robotMap[robottypes.RobotTypeQuest]
}

func (sb *scene) GetAllRobots() map[int64]Player {
	return sb.allRobotMap
}

func (sb *scene) GetLastRobotTime() int64 {
	return sb.lastRobotTime
}

func (sb *scene) SetLastRobotTime(robotTime int64) {
	sb.lastRobotTime = robotTime
}

func (sb *scene) GetNumOfNPC() int32 {
	return int32(len(sb.npcs))
}

func (sb *scene) GetNumOfRobot() int32 {
	return int32(len(sb.robotMap[robottypes.RobotTypeQuest]))
}

func (sb *scene) GetNumOfAllRobot() int32 {
	return int32(len(sb.allRobotMap))
}

func (sb *scene) RefreshBiology(groupId int32) {
	sb.currentGroup = groupId
	//回收npc
	biologyMap := sb.MapTemplate().GetSceneBiologyMapByGroup(groupId)
	if biologyMap == nil {
		log.WithFields(
			log.Fields{
				"mapId":   sb.MapId(),
				"groupId": groupId,
			}).Warn("scene:波数不存在")
		return
	}

	for _, biology := range biologyMap {

		if biology.GetBiology().GetRebornType() == scenetypes.BiologyRebornTypeTime {
			continue
		}
		deadTime := int64(0)
		n := CreateNPC(scenetypes.OwnerTypeNone, 0, 0, 0, int32(biology.IndexID), biology.GetBiology(), biology.GetPos(), biology.Angle, deadTime)
		if n == nil {
			log.WithFields(
				log.Fields{
					"mapId":       sb.MapId(),
					"biologyType": biology.GetBiology().GetBiologyScriptType(),
				}).Warnln("创建npc,不存在")
			continue
		}
		//映射部怪表
		sb.idxNpcs[int64(biology.Id)] = n
		// log.WithFields(
		// 	log.Fields{
		// 		"mapId":       sb.MapId(),
		// 		"biologyType": biology.GetBiology().GetBiologyScriptType().String(),
		// 		"pos":         n.GetPosition(),
		// 		"id":          n.GetId(),
		// 	}).Infoln("创建npc,成功")

		//设置场景
		sb.AddSceneObject(n)
	}
	if sb.sceneDelegate != nil {
		sb.sceneDelegate.OnSceneRefreshGroup(sb, sb.currentGroup)
	}
}

func (sb *scene) Post(msg message.Message) bool {
	sb.m.Lock()
	defer sb.m.Unlock()
	if sb.state == SceneStateStopped {
		return false
	}
	sb.msgs <- msg
	return true
}

//添加场景对象
func (sb *scene) AddSceneObject(so SceneObject) {

	switch obj := so.(type) {
	case Player:
		if sb.sceneDelegate != nil {
			sb.sceneDelegate.OnScenePlayerBeforeEnter(sb, obj)
		}
		break
	case NPC:
		break
	case DropItem:
		break
	}
	so.EnterScene(sb)
	// log.WithFields(
	// 	log.Fields{
	// 		"id":   so.GetId(),
	// 		"pos":  so.GetPosition(),
	// 		"goId": runtimeutils.Goid(),
	// 	}).Infoln("场景对象添加")

	//进入AOI
	bo, ok := so.(NPC)
	if ok {
		if !bo.IsDead() {
			sb.aoiManager.Enter(so, so.GetPosition())
		}
	} else {
		sb.aoiManager.Enter(so, so.GetPosition())
	}
	sb.sceneObjects[so.GetId()] = so
	switch obj := so.(type) {
	case Player:
		sb.addPlayer(obj)
		if sb.sceneDelegate != nil {
			sb.sceneDelegate.OnScenePlayerEnter(sb, obj)
		}
		if obj.GetLingTong() != nil && !obj.IsLingTongHidden() {
			obj.GetLingTong().SetPosition(so.GetPosition())
			sb.AddSceneObject(obj.GetLingTong())
		}
		break
	case NPC:
		sb.addNPC(obj)
		if sb.sceneDelegate != nil {
			sb.sceneDelegate.OnSceneBiologyEnter(sb, obj)
		}
		break
	case DropItem:
		sb.addDropItem(obj)
		break
	case LingTong:
		sb.addLingTong(obj)
		break
	default:
		panic("never reach here")
	}
}

//bo死亡
func (sb *scene) OnDead(bo BattleObject) {

	for _, obj := range bo.GetNeighbors() {
		switch nei := obj.(type) {
		case BattleObject:
			nei.OnDead(bo)
		}
	}

	switch tbo := bo.(type) {
	case NPC:
		sb.onNPCDead(tbo)
		break
	case Player:
		sb.onPlayerDead(tbo)
		break
	}
}

//玩家死亡
func (sb *scene) onPlayerDead(p Player) {

	if sb.sceneDelegate != nil {
		sb.sceneDelegate.OnScenePlayerDead(sb, p)
	}
}

//npc死亡
func (sb *scene) onNPCDead(npc NPC) {

	//不复活 移除

	sb.aoiManager.Leave(npc)

	if sb.sceneDelegate != nil {
		sb.sceneDelegate.OnSceneBiologyDead(sb, npc)
	}
	//所有怪物都死亡了
	if len(sb.npcs) == 1 {
		if sb.sceneDelegate != nil {
			sb.sceneDelegate.OnSceneBiologyAllDead(sb)
		}
	}
}

func (sb *scene) OnPlayerGetItem(pl Player, itemData *droptemplate.DropItemData) {
	if sb.sceneDelegate != nil {
		sb.sceneDelegate.OnScenePlayerGetItem(sb, pl, itemData)
	}

}

func (sb *scene) OnPlayerGetExp(pl Player, num int64) {
	if sb.sceneDelegate != nil {
		sb.sceneDelegate.OnScenePlayerGetExp(sb, pl, num)
	}
}

//npc生
func (sb *scene) OnReborn(bo BattleObject, pos coretypes.Position) {
	for _, obj := range bo.GetNeighbors() {
		switch neiBo := obj.(type) {
		case BattleObject:
			neiBo.OnReborn(bo)
			break
		}
	}
	switch tbo := bo.(type) {
	case NPC:
		sb.onNPCReborn(tbo, pos)
		break
	case Player:
		sb.onPlayerReborn(tbo, pos)
		break
	}
}

//npc重生
func (sb *scene) onNPCReborn(npc NPC, pos coretypes.Position) {
	sb.aoiManager.Enter(npc, pos)
	if sb.sceneDelegate != nil {
		sb.sceneDelegate.OnSceneBiologyReborn(sb, npc)
	}
}

//玩家重生
func (sb *scene) onPlayerReborn(pl Player, pos coretypes.Position) {
	sb.aoiManager.Move(pl, pos)
	if sb.sceneDelegate != nil {
		sb.sceneDelegate.OnScenePlayerReborn(sb, pl)
	}
}

func (sb *scene) Move(bo BattleObject, pos coretypes.Position) {
	//保存aoi数据
	sb.aoiManager.Move(bo, pos)
}

//移除场景对象
func (sb *scene) RemoveSceneObject(so SceneObject, active bool) {

	so, exist := sb.sceneObjects[so.GetId()]
	if !exist {
		// panic(fmt.Errorf("scene:场景对象不存在"))
		return
	}

	switch obj := so.(type) {
	case Player:
		if obj.GetLingTong() != nil && !obj.IsLingTongHidden() {
			sb.RemoveSceneObject(obj.GetLingTong(), active)
		}
		sb.removePlayer(obj)
		if sb.sceneDelegate != nil {
			sb.sceneDelegate.OnScenePlayerExit(sb, obj, active)
		}

		break
	case NPC:
		sb.removeNPC(obj)
		break
	case DropItem:
		sb.removeDropItem(obj)
		break
	case LingTong:
		sb.removeLingTong(obj)
		break
	default:
		panic("never reach here")
	}
	delete(sb.sceneObjects, so.GetId())

	so.ExitScene(active)

	//退出aoi
	sb.aoiManager.Leave(so)
	//TODO 广播给玩家退出场景 以防内存泄露
	switch so.(type) {
	case Player:
		for _, tempSceneObj := range sb.sceneObjects {
			switch sceneObj := tempSceneObj.(type) {
			case Player:
				sceneObj.RemoveLoadedPlayer(so.GetId())
				break
			case LingTong:
				sceneObj.RemoveLoadedPlayer(so.GetId())
				break
			}
		}
		break
	}

	// log.WithFields(
	// 	log.Fields{
	// 		"id":              so.GetId(),
	// 		"sceneObjectType": so.GetSceneObjectType(),
	// 		"pos":             so.GetPosition(),
	// 		"goId":            runtimeutils.Goid(),
	// 	}).Infoln("场景对象移除")

}

func (sb *scene) addLingTong(p LingTong) {
	sb.lingTongs[p.GetId()] = p
}

func (sb *scene) removeLingTong(p LingTong) {

	delete(sb.lingTongs, p.GetId())
}

func (sb *scene) addPlayer(p Player) {
	sb.players[p.GetId()] = p
	if p.IsRobot() {
		robot := p.(RobotPlayer)
		subMap, ok := sb.robotMap[robot.GetRobotType()]
		if !ok {
			subMap = make(map[int64]Player)
			sb.robotMap[robot.GetRobotType()] = subMap
		}
		subMap[robot.GetId()] = robot
		sb.allRobotMap[robot.GetId()] = robot
	}
}

func (sb *scene) removePlayer(p Player) {
	if p.IsRobot() {
		robot := p.(RobotPlayer)
		subMap, ok := sb.robotMap[robot.GetRobotType()]
		if ok {
			delete(subMap, robot.GetId())
		}
		delete(sb.allRobotMap, robot.GetId())
	}
	delete(sb.players, p.GetId())

}

func (sb *scene) addNPC(npc NPC) {
	sb.npcs[npc.GetId()] = npc
	npcMap, ok := sb.npcMapOfMap[npc.GetBiologyTemplate().GetBiologyScriptType()]
	if !ok {
		npcMap = make(map[int64]NPC)
		sb.npcMapOfMap[npc.GetBiologyTemplate().GetBiologyScriptType()] = npcMap
	}
	npcMap[npc.GetId()] = npc

	npcList := sb.npcListOfMap[int32(npc.GetBiologyTemplate().TemplateId())]
	sb.npcListOfMap[int32(npc.GetBiologyTemplate().TemplateId())] = append(npcList, npc)
}

func (sb *scene) removeNPC(npc NPC) {
	if sb.sceneDelegate != nil {
		sb.sceneDelegate.OnSceneBiologyExit(sb, npc)
	}
	delete(sb.npcs, npc.GetId())
	npcMap, ok := sb.npcMapOfMap[npc.GetBiologyTemplate().GetBiologyScriptType()]
	if ok {
		delete(npcMap, npc.GetId())
	}
	npcList := sb.npcListOfMap[int32(npc.GetBiologyTemplate().TemplateId())]

	foundIndex := -1
	for index, tempNpc := range npcList {
		if tempNpc == npc {
			foundIndex = index
		}
	}
	if foundIndex < 0 {
		return
	}
	sb.npcListOfMap[int32(npc.GetBiologyTemplate().TemplateId())] = append(npcList[:foundIndex], npcList[foundIndex+1:]...)
}

func (sb *scene) addDropItem(dropItem DropItem) {
	sb.dropItems[dropItem.GetId()] = dropItem
}

func (sb *scene) removeDropItem(dropItem DropItem) {
	delete(sb.dropItems, dropItem.GetId())
}

func (sb *scene) init() {
	log.WithFields(
		log.Fields{
			"id":      sb.Id(),
			"sceneId": sb.mapTemplate.Id,
			"name":    sb.mapTemplate.Name,
		}).Info("scene:开启")
	sb.state = SceneStateStart
	sb.startTime = global.GetGame().GetTimeService().Now()
	sb.heartbeatTimer.Reset(heartbeatTime)
	if sb.sceneDelegate != nil {
		sb.sceneDelegate.OnSceneStart(sb)
	}

	//刷怪
	sb.currentGroup = 0
	sb.RefreshBiology(sb.currentGroup)
}

const (
	warnElapse = int64(100)
)

//goroutine不安全 都会在scene runner 调用
//开始
func (sb *scene) start() {
Loop:
	for {
		select {
		case <-sb.heartbeatTimer.C:
			//计算时间
			before := global.GetGame().GetTimeService().Now()
			sb.tick()
			sb.heartbeatTimer.Reset(heartbeatTime)
			now := global.GetGame().GetTimeService().Now()
			elapse := now - before
			if elapse > warnElapse {
				log.WithFields(
					log.Fields{
						"id":      sb.Id(),
						"sceneId": sb.mapTemplate.Id,
						"name":    sb.mapTemplate.Name,
						"elapse":  elapse,
						"players": len(sb.GetAllPlayers()),
						"items":   len(sb.GetAllItems()),
						"npcs":    len(sb.GetAllNPCS()),
					}).Warn("scene:心跳太久")
			}
			break
		case m, ok := <-sb.msgs:
			{
				if !ok {
					break Loop
				}
				err := global.GetGame().GetMessageHandler().HandleMessage(m)
				if err != nil {
					log.WithFields(
						log.Fields{
							"error": err,
						}).Error("message:处理消息,错误")
				}
			}
		case <-sb.done:
			{
				sb.Stop(true, true)
			}
		}
	}
	log.WithFields(
		log.Fields{
			"id":      sb.Id(),
			"sceneId": sb.mapTemplate.Id,
			"name":    sb.mapTemplate.Name,
		}).Info("scene:场景心跳结束")
}

const (
	afterFinishTime = 15 * common.SECOND
	logTime         = common.MINUTE
)

//消息处理
func (sb *scene) tick() {
	defer func() {
		if terr := recover(); terr != nil {
			debug.PrintStack()
			exceptionContent := string(debug.Stack())
			log.WithFields(
				log.Fields{
					"mapName": sb.mapTemplate.Name,
					"error":   terr,
					"stack":   string(debug.Stack()),
				}).Error("scene:tick,错误")
			gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
		}
	}()
	//记录日志
	// sb.log()
	now := global.GetGame().GetTimeService().Now()
	elapseTime := now - sb.startTime
	//副本
	if sb.mapTemplate.IsFuBen() {
		if sb.state == SceneStateStart {
			//失败
			if int64(sb.mapTemplate.PointsTime) <= elapseTime {
				sb.Finish(false)
			}
			//结束
			if int64(sb.mapTemplate.LastTime) <= elapseTime {
				//停止
				sb.Finish(false)
			}
		} else if sb.state == SceneStateFinish {
			elapseFinishTime := now - sb.finishTime
			//完成后 几秒退出
			if elapseFinishTime >= int64(afterFinishTime) {
				sb.Stop(true, false)
				return
			}
		}

	} else if sb.mapTemplate.IsArena() {
		if sb.state == SceneStateStart {
			// 失败
			if int64(sb.mapTemplate.PointsTime) <= elapseTime {
				sb.Finish(false)
			}
			if sb.endTime <= now {
				sb.Finish(false)
			}
		} else if sb.state == SceneStateFinish {

			//完成后 几秒退出
			if sb.endTime <= now {
				sb.Stop(true, false)
				return
			}

		}

	} else if sb.mapTemplate.IsActivity() || sb.mapTemplate.IsActivityFuBen() {
		//活动副本
		now := global.GetGame().GetTimeService().Now()
		if sb.state == SceneStateStart {
			//完成
			if int64(sb.endTime) <= now {
				sb.Finish(true)
			}
		} else if sb.state == SceneStateFinish {
			elapseFinishTime := now - sb.finishTime
			//完成后 几秒退出
			if elapseFinishTime >= int64(afterFinishTime) {
				sb.Stop(true, false)
				return
			}
		}
	} else if sb.mapTemplate.IsActivitySub() {
		now := global.GetGame().GetTimeService().Now()
		if sb.state == SceneStateFinish {
			elapseFinishTime := now - sb.finishTime
			//完成后 几秒退出
			if elapseFinishTime >= int64(afterFinishTime) {
				sb.Stop(true, false)
				return
			}
		}
	}

	if sb.sceneDelegate != nil {
		sb.sceneDelegate.OnSceneTick(sb)
	}

	for _, p := range sb.players {
		p.Tick()
	}

	sb.heartbeat()

	sb.tickNPC()

	sb.tickDropItem()

	sb.heartbeatRunner.Heartbeat()
}

//消息处理
func (sb *scene) log() {
	defer func() {
		if terr := recover(); terr != nil {
			debug.PrintStack()
			exceptionContent := string(debug.Stack())
			log.WithFields(
				log.Fields{
					"mapName": sb.mapTemplate.Name,
					"error":   terr,
					"stack":   string(debug.Stack()),
				}).Error("scene:log,错误")
			gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
		}
	}()
	now := global.GetGame().GetTimeService().Now()
	elapseTime := now - sb.lastLogTime
	if elapseTime > int64(logTime) {
		sb.lastLogTime = now
		gameevent.Emit(sceneeventtypes.EventTypeSceneInfoLog, sb, nil)
	}

}

//数据更新
func (sb *scene) heartbeat() error {
	for _, p := range sb.players {
		p.Heartbeat()
	}
	for _, p := range sb.lingTongs {
		p.Heartbeat()
	}
	return nil
}

func (sb *scene) tickNPC() {

	//npc移动
	for _, npc := range sb.npcs {
		if npc.ShouldRemove() {
			sb.RemoveSceneObject(npc, false)
			continue
		}
		npc.Heartbeat()
	}
	return
}

func (sb *scene) tickDropItem() {
	//掉落物品
	for _, dropItem := range sb.dropItems {
		dropItem.Tick()
	}
	return
}

func (sb *scene) shouldFinish() bool {
	if sb.state == SceneStateStart {
		return true
	}
	return false
}

//goroutine不安全 都会在scene runner 调用
func (sb *scene) Finish(success bool) {
	//世界服 没有结算
	if sb.mapTemplate.IsWorld() || sb.mapTemplate.IsBoss() {
		return
	}
	if !sb.shouldFinish() {
		return
	}

	log.WithFields(
		log.Fields{
			"mapId":   sb.MapId(),
			"sceneId": sb.Id(),
			"success": success,
		}).Info("scene:场景结束")
	sb.success = success
	sb.state = SceneStateFinish
	now := global.GetGame().GetTimeService().Now()
	sb.finishTime = now
	if sb.sceneDelegate != nil {
		sb.sceneDelegate.OnSceneFinish(sb, sb.success, sb.getUseTime())
		gameevent.Emit(sceneeventtypes.EventTypeSceneFinish, sb, nil)
	}

	return
}

//获取使用时间
func (sb *scene) getUseTime() int64 {
	return sb.finishTime - sb.startTime
}

func (sb *scene) GetItem(itemId int64) DropItem {
	it, ok := sb.dropItems[itemId]
	if !ok {
		return nil
	}
	return it
}

func (sb *scene) GetSceneObject(itemId int64) SceneObject {
	it, ok := sb.sceneObjects[itemId]
	if !ok {
		return nil
	}
	return it
}

func (sb *scene) GetPlayer(id int64) Player {
	p, ok := sb.players[id]
	if !ok {
		return nil
	}
	return p
}

func (sb *scene) GetLingTong(id int64) LingTong {
	p, ok := sb.lingTongs[id]
	if !ok {
		return nil
	}
	return p
}

func (sb *scene) BroadcastMsg(msg proto.Message) {
	for _, pl := range sb.players {
		pl.SendMsg(msg)
	}
}

func (sb *scene) AsyncStop() {
	close(sb.done)
}

//goroutine不安全 都会在scene runner 调用
func (sb *scene) Stop(active, stopComplete bool) {
	if sb.state == SceneStateStopped {
		return
	}

	sb.state = SceneStateStopped

	defer func() {
		if terr := recover(); terr != nil {
			debug.PrintStack()
			exceptionContent := string(debug.Stack())
			log.WithFields(
				log.Fields{
					"mapName": sb.mapTemplate.Name,
					"error":   terr,
					"stack":   string(debug.Stack()),
				}).Error("scene:结束,错误")
			gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
		}
		//主动关闭
		close(sb.msgs)
		log.WithFields(
			log.Fields{
				"id":      sb.Id(),
				"sceneId": sb.mapTemplate.Id,
				"name":    sb.mapTemplate.Name,
			}).Info("scene:关闭")
	}()

	if sb.sceneDelegate != nil {
		sb.sceneDelegate.OnSceneStop(sb)
	}
	sb.beforeStop(stopComplete)

	return
}

//关闭前清理
func (sb *scene) beforeStop(stopComplete bool) {
	sb.clearPlayers(stopComplete)
}

//清理玩家
func (sb *scene) clearPlayers(stopComplete bool) {
	for _, p := range sb.players {
		if stopComplete {
			p.ExitScene(false)
		} else {
			p.BackLastScene()
		}
	}
}

func (sb *scene) GetExpAddPercent(p Player) int32 {
	if sb.sceneDelegate == nil {
		return 0
	}
	return sb.sceneDelegate.GetExpAddPercent(p)
}

func (sb *scene) String() string {

	return fmt.Sprintf("场景id[%d],地图id[%d],地图名[%s],玩家数[%d],机器人[%d],物品数[%d],npc数量[%d],状态[%s]", sb.Id(), sb.mapTemplate.Id, sb.mapTemplate.Name, len(sb.GetAllPlayers()), sb.GetNumOfRobot(), len(sb.GetAllItems()), len(sb.GetAllNPCS()), sb.state.String())
}

const (
	defaultSceneCapacity = 10000
)

const (
	queueCapacity = 20000
	maxTime       = time.Millisecond
	heartbeatTime = 40 * time.Millisecond
)

func newScene(mapTemplate *gametemplate.MapTemplate, sh SceneDelegate) *scene {
	sb := &scene{}
	sb.id, _ = idutil.GetId()

	sb.mapTemplate = mapTemplate
	enterDistance := mapTemplate.GetMapType().GetEnterDistance()

	exitDistance := mapTemplate.GetMapType().GetExitDistance()

	sb.aoiManager = aoi.NewXZListAOIManager(enterDistance, exitDistance)
	startX, startZ := sb.mapTemplate.GetMap().GetStartXZ()
	sb.world = nav.NewWorld(startX, startZ, sb.mapTemplate.GetMap().Accuracy(), sb.mapTemplate.GetMap().GetMapMask())
	sb.players = make(map[int64]Player)
	sb.robotMap = make(map[robottypes.RobotType]map[int64]Player)
	sb.allRobotMap = make(map[int64]Player)
	sb.lingTongs = make(map[int64]LingTong)
	sb.npcs = make(map[int64]NPC)
	sb.idxNpcs = make(map[int64]NPC)
	sb.npcMapOfMap = make(map[scenetypes.BiologyScriptType]map[int64]NPC)
	sb.npcListOfMap = make(map[int32][]NPC)

	sb.dropItems = make(map[int64]DropItem)
	sb.sceneObjects = make(map[int64]SceneObject)
	sb.sceneDelegate = sh
	sb.msgs = make(chan message.Message, queueCapacity)
	sb.heartbeatTimer = time.NewTimer(heartbeatTime)
	sb.heartbeatTimer.Stop()
	sb.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	sb.heartbeatRunner.AddTask(CreateSceneRobotTask(sb))
	sb.heartbeatRunner.AddTask(createRankTask(sb))
	sb.init()
	return sb
}

//创建活动副本
func CreateActivityScene(mapId int32, endTime int64, sh SceneDelegate) Scene {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if !mapTemplate.IsActivity() {
		return nil
	}
	s := CreateScene(mapTemplate, endTime, sh)
	return s
}

//创建活动子副本
func CreateActivitySubScene(mapId int32, endTime int64, sh SceneDelegate) Scene {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if !mapTemplate.IsActivitySub() {
		return nil
	}
	s := CreateScene(mapTemplate, endTime, sh)
	return s
}

//验证是否是副本
func CreateFuBenScene(mapId int32, sh SceneDelegate) Scene {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if !mapTemplate.IsFuBen() {
		return nil
	}
	s := CreateScene(mapTemplate, 0, sh)
	return s
}

func CreateArenaScene(mapId int32, endTime int64, sh SceneDelegate) Scene {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if !mapTemplate.IsArena() {
		return nil
	}
	s := CreateScene(mapTemplate, endTime, sh)
	return s
}

// 创建打宝塔场景
func CreateTowerScene(mapId int32, sh SceneDelegate) Scene {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if !mapTemplate.IsTower() {
		return nil
	}

	s := CreateScene(mapTemplate, 0, sh)
	return s
}

func CreateScene(mapTemplate *gametemplate.MapTemplate, endTime int64, sh SceneDelegate) Scene {
	s := newScene(mapTemplate, sh)
	s.endTime = endTime
	s.SceneRankManager = newSceneRankManager(s)
	s.done = make(chan struct{})
	cs.OnSceneStart(s)
	go func() {
		s.start()
		cs.OnSceneStop(s)
	}()
	return s
}

type contextKey string

const (
	sceneContextKey contextKey = "scene"
)

func WithScene(parent context.Context, s Scene) context.Context {
	return context.WithValue(parent, sceneContextKey, s)
}

func SceneInContext(ctx context.Context) Scene {
	s := ctx.Value(sceneContextKey)
	if s == nil {
		return nil
	}
	ts, ok := s.(Scene)
	if !ok {
		return nil
	}
	return ts
}
