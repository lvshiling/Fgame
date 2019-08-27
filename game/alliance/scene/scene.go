package scene

import (
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancetemplate "fgame/fgame/game/alliance/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type AllianceSceneData interface {
	GetScene() scene.Scene
	//获取初始守方联盟id
	GetFirstDefendAllianceId() int64
	//获取当前守方id
	GetCurrentDefendAllianceId() int64
	//获取当前守方名字
	GetCurrentDefendAllianceName() string
	//获取结束时间
	GetEndTime() int64
	//守方玉玺被采集
	YuXiNpcCollectFinish(n scene.NPC, collectAlId int64, alName string, huFu int64) (flag bool)
	//虎符
	OnHuFuChanged(allianceId int64, huFu int64)
	//获取当前联盟虎符
	GetCurrentDefendAllianceHuFu() int64

	//召唤守卫
	CallGuard(index int32) (flag bool)
	//获取召唤列表
	GetCallGuardList() []int32
	//守卫重生
	GuardReborn(npc scene.NPC)

	//获取当前门
	GetCurrentDoor() int32
	//获取当前玉玺
	GetCollectYuXi() scene.NPC
	//防护罩是否打破
	IsProtectBroken() bool

	//获取占领复活点的联盟
	GetCurrentReliveAllianceId() int64
	//获取正在采集复活点开始时间
	GetCollectReliveStartTime() int64
	//获取正在采集复活点的用户
	GetCollectRelivePlayerId() int64
	//获取正在采集复活点的联盟
	GetCollectReliveAllianceId() int64
	//复活点占领
	ReliveOccupy(allianceId int64, playerId int64) bool
	//获取复活旗子
	GetReliveFlag() scene.NPC
	//清除复活占领
	ClearReliveOccupy()
}

//九霄城战
type allianceSceneData struct {
	*scene.SceneDelegateBase
	//防守方
	firstDefendAllianceId int64
	currentAllianceId     int64
	currentAllianceName   string
	currentAllianceHuFu   int64

	//当前复活旗的联盟
	currentReliveAllianceId int64
	//正在采集的阵营
	collectReliveAllianceId int64
	//正在采集复活旗子的玩家
	collectRelivePlayerId int64
	//采集复活旗子开始时间
	collectReliveFlagStartTime int64
	//复活旗子
	reliveFlag scene.NPC

	//保护罩
	protectNpc scene.NPC
	yuxiNpc    scene.NPC

	//攻破的城门
	currentDoor int32
	//结束时间
	endTime int64
	//守卫
	guardMap map[int32]scene.NPC
	//召唤守卫
	callGuardMap map[int32]scene.NPC
}

//城战场景数据
func CreateAllianceSceneData(mapId int32, defAllianceId int64, allianceName string, allianceHuFu int64, endTime int64) AllianceSceneData {
	sd := &allianceSceneData{
		firstDefendAllianceId:   defAllianceId,
		currentAllianceId:       defAllianceId,
		currentReliveAllianceId: defAllianceId,
		currentAllianceName:     allianceName,
		currentAllianceHuFu:     allianceHuFu,
		currentDoor:             0,
		endTime:                 endTime,
		guardMap:                make(map[int32]scene.NPC),
		callGuardMap:            make(map[int32]scene.NPC),
		SceneDelegateBase:       scene.NewSceneDelegateBase(),
	}

	chengWaiScene := createChengWaiScene(mapId, sd, endTime)
	if chengWaiScene == nil {
		return nil
	}

	return sd
}

//城战场景
func createChengWaiScene(mapId int32, sd *allianceSceneData, endTime int64) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeChengZhan {
		return nil
	}
	s = scene.CreateActivityScene(mapId, endTime, sd)
	return s
}

// 场景开始
func (sd *allianceSceneData) OnSceneStart(s scene.Scene) {
	sd.SceneDelegateBase.OnSceneStart(s)

	sd.initProtectNpc()
	return
}

//生物进入
func (sd *allianceSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
	// 复活点
	if npc.GetBiologyTemplate().GetBiologyScriptType() == scenetypes.BiologyScriptTypeRelivePoint {
		sd.reliveFlag = npc
	}

	// 守卫
	if npc.GetBiologyTemplate().ScriptType == scenetypes.BiologyScriptTypeXianMengNPC {
		_, exist := sd.guardMap[npc.GetIdInScene()]
		if !exist {
			sd.guardMap[npc.GetIdInScene()] = npc
		}
	}
}

//场景心跳
func (sd *allianceSceneData) OnSceneTick(s scene.Scene) {
	if s.State() == scene.SceneStateFinish || s.State() == scene.SceneStateStopped {
		return
	}

	//占领复活点检查
	if sd.ifReliveOccuoyFinish() {
		sd.onReliveOccupyFinish()
	}

	//定时奖励
	gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneTickRew, sd, nil)
}

//怪物死亡
func (sd *allianceSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("alliance:城战应该是同一个场景"))
	}

	sd.doorNpcDead(npc)
	sd.protectNpcDead(npc)
}

//玩家进入
func (sd *allianceSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("alliance:城战应该是同一个场景"))
	}

	//初始化阵营
	sd.initCamp(p)

	//场景信息推送
	gameevent.Emit(allianceeventtypes.EventTypePlayerEnterAllianceScene, sd, p)
}

//玩家退出
func (sd *allianceSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("alliance:城战应该是同一个场景"))
	}

	//设置阵营
	p.SetFactionType(scenetypes.FactionTypePlayer)

	//清空占领复活点读条
	if sd.collectRelivePlayerId == p.GetId() {
		sd.ClearReliveOccupy()
	}
}

//场景完成
func (sd *allianceSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("alliance:城战应该是同一个场景"))
	}

	gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneFinish, sd, nil)
}

//召唤守卫
func (sd *allianceSceneData) CallGuard(index int32) (flag bool) {
	guardNpc, exist := sd.guardMap[index]
	if !exist {
		return
	}

	_, exist = sd.callGuardMap[index]
	if exist {
		return
	}
	sd.callGuardMap[index] = guardNpc

	guardNpc.SetFactionType(scenetypes.FactionTypeChengZhanDefendNPC)
	guardNpc.AllianceCalled(sd.currentAllianceId)
	flag = true
	return
}

func (sd *allianceSceneData) GuardReborn(npc scene.NPC) {
	npcId := npc.GetIdInScene()
	delete(sd.callGuardMap, npcId)
}

func (sd *allianceSceneData) GetCallGuardList() (guardList []int32) {
	for index, _ := range sd.callGuardMap {
		guardList = append(guardList, index)
	}
	return
}

//清除占领复活点读条
func (sd *allianceSceneData) ClearReliveOccupy() {
	if sd.GetScene().State() == scene.SceneStateFinish || sd.GetScene().State() == scene.SceneStateStopped {
		return
	}
	if sd.collectRelivePlayerId == 0 {
		return
	}
	log.WithFields(
		log.Fields{
			"currentReliveAllianceId": sd.currentReliveAllianceId,
			"collectReliveAllianceId": sd.collectReliveAllianceId,
		}).Info("alliance:清除读条")

	gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneReliveOccupyStop, sd, sd.collectReliveAllianceId)
	sd.collectReliveAllianceId = 0
	sd.collectRelivePlayerId = 0
}

//虎符变化
func (sd *allianceSceneData) OnHuFuChanged(allianceId int64, huFu int64) {
	if sd.currentAllianceId == allianceId {
		sd.currentAllianceHuFu = huFu
		//推送虎符改变
		gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneDefendHuFuChanged, sd, sd.currentAllianceHuFu)
	} else {
		eventData := allianceeventtypes.CreateAllianceSceneHuFuChangedEvent(allianceId, huFu)
		gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneHuFuChanged, sd, eventData)
	}
}

//占领复活点
func (sd *allianceSceneData) ReliveOccupy(allianceId int64, playerId int64) bool {
	if sd.GetScene().State() == scene.SceneStateFinish || sd.GetScene().State() == scene.SceneStateStopped {
		return false
	}
	if sd.currentReliveAllianceId == allianceId {
		return false
	}
	if sd.collectReliveAllianceId != 0 {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	//读条
	sd.collectReliveAllianceId = allianceId
	sd.collectRelivePlayerId = playerId
	sd.collectReliveFlagStartTime = now

	log.WithFields(
		log.Fields{
			"currentReliveAllianceId": sd.currentReliveAllianceId,
			"collectReliveAllianceId": sd.collectReliveAllianceId,
			"collectRelivePlayerId":   sd.collectRelivePlayerId,
		}).Info("alliance:复活点占领读条")

	// 发事件
	eventData := allianceeventtypes.CreateReliveOccupyEvent(sd.collectReliveAllianceId, sd.collectRelivePlayerId)
	gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneReliveOccupying, sd, eventData)
	return true
}

//玉玺被采集
func (sd *allianceSceneData) YuXiNpcCollectFinish(collecNpc scene.NPC, collectAlId int64, alName string, huFu int64) (flag bool) {
	biology := collecNpc.GetBiologyTemplate()
	if biology.GetBiologyScriptType() != scenetypes.BiologyScriptTypeGeneralCollect {
		return
	}

	warTemp := alliancetemplate.GetAllianceTemplateService().GetWarTemplate()
	if biology.Id != warTemp.GetYuXiBiologyTemp().Id {
		return
	}
	if sd.currentAllianceId == collectAlId {
		return
	}

	sd.currentAllianceId = collectAlId
	sd.currentAllianceName = alName
	sd.currentAllianceHuFu = huFu

	// 阵营转换
	for _, p := range sd.GetScene().GetAllPlayers() {
		sd.initCamp(p)
	}

	//清空所有守卫
	for _, n := range sd.GetScene().GetAllNPCS() {
		if n.GetBiologyTemplate().ScriptType != scenetypes.BiologyScriptTypeXianMengNPC {
			continue
		}
		if n.GetOwnerId() != 0 {
			n.Recycle(0)
		}
	}

	// 生成防护罩
	sd.initProtectNpc()

	flag = true
	return
}

func (sd *allianceSceneData) IsProtectBroken() bool {
	if sd.protectNpc == nil {
		return true
	}

	if sd.protectNpc.IsDead() {
		return true
	}

	return false
}

func (sd *allianceSceneData) GetEndTime() int64 {
	return sd.endTime
}

func (sd *allianceSceneData) GetCurrentDefendAllianceId() int64 {
	return sd.currentAllianceId
}

func (sd *allianceSceneData) GetCurrentDefendAllianceName() string {
	return sd.currentAllianceName
}

func (sd *allianceSceneData) GetCollectRelivePlayerId() int64 {
	return sd.collectRelivePlayerId
}

func (sd *allianceSceneData) GetCurrentDoor() int32 {
	return sd.currentDoor
}

func (sd *allianceSceneData) GetCollectYuXi() scene.NPC {
	return sd.yuxiNpc
}

func (sd *allianceSceneData) GetFirstDefendAllianceId() int64 {
	return sd.firstDefendAllianceId
}

func (sd *allianceSceneData) GetCollectReliveAllianceId() int64 {
	return sd.collectReliveAllianceId
}

func (sd *allianceSceneData) GetCollectReliveStartTime() int64 {
	return sd.collectReliveFlagStartTime
}

func (sd *allianceSceneData) GetCurrentDefendAllianceHuFu() int64 {
	return sd.currentAllianceHuFu
}
func (sd *allianceSceneData) GetCurrentReliveAllianceId() int64 {
	return sd.currentReliveAllianceId
}

func (sd *allianceSceneData) GetReliveFlag() scene.NPC {
	return sd.reliveFlag
}
