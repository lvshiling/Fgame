package scene

import (
	chuangshieventtypes "fgame/fgame/cross/chuangshi/event/types"
	chuangshitemplate "fgame/fgame/game/chuangshi/template"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	pktype "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type FuShuSceneData interface {
	scene.SceneDelegate
	//是否排队
	IfLineup() bool
	//获取初始守方阵营
	GetInitDefendCampType() chuangshitypes.ChuangShiCampType
	//获取当前守方阵营
	GetCurrentDefendCampType() chuangshitypes.ChuangShiCampType
	// //获取结束时间
	// GetEndTime() int64
	//守方玉玺被采集
	YuXiNpcCollectFinish(n scene.NPC, collectCampType chuangshitypes.ChuangShiCampType) (flag bool)

	//获取当前门
	GetCurrentDoor() int32
	//获取当前玉玺
	GetCollectYuXi() scene.NPC
	//防护罩是否打破
	IsProtectBroken() bool

	// //获取占领复活点的联盟
	// GetCurrentReliveAllianceId() int64
	// //获取正在采集复活点开始时间
	// GetCollectReliveStartTime() int64
	// //获取正在采集复活点的用户
	// GetCollectRelivePlayerId() int64
	// //获取正在采集复活点的联盟
	// GetCollectReliveAllianceId() int64
	// //复活点占领
	// ReliveOccupy(allianceId int64, playerId int64) bool
	// //获取复活旗子
	// GetReliveFlag() scene.NPC
	// //清除复活占领
	// ClearReliveOccupy()
}

//附属城池
type fuShuSceneData struct {
	*scene.SceneDelegateBase

	jieMeng            bool
	initDefendCampType chuangshitypes.ChuangShiCampType
	curDefenCampType   chuangshitypes.ChuangShiCampType

	//当前复活旗的阵营
	currentReliveCampType chuangshitypes.ChuangShiCampType
	//正在采集的阵营
	collectReliveCampType chuangshitypes.ChuangShiCampType
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
}

// 场景开始
func (sd *fuShuSceneData) OnSceneStart(s scene.Scene) {
	sd.SceneDelegateBase.OnSceneStart(s)

	sd.initProtectNpc()
	return
}

//场景心跳
func (sd *fuShuSceneData) OnSceneTick(s scene.Scene) {
	if s.State() == scene.SceneStateFinish || s.State() == scene.SceneStateStopped {
		return
	}

	//占领复活点检查
	if sd.ifReliveOccuoyFinish() {
		sd.onReliveOccupyFinish()
	}

	// //定时奖励
	// gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneTickRew, sd, nil)
}

// 玩家进入
func (sd *fuShuSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	sd.initCamp(p)

	// //场景信息推送
	// gameevent.Emit(allianceeventtypes.EventTypePlayerEnterAllianceScene, sd, p)
}

//生物进入
func (sd *fuShuSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
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

//玩家退出
func (sd *fuShuSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("chuangshji:城战应该是同一个场景"))
	}

	//设置阵营
	p.SetFactionType(scenetypes.FactionTypePlayer)

	// //清空占领复活点读条
	// if sd.collectRelivePlayerId == p.GetId() {
	// 	sd.ClearReliveOccupy()
	// }
}

//场景完成
func (sd *fuShuSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("chuangshji:城战应该是同一个场景"))
	}

	// gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneFinish, sd, nil)
}

//怪物死亡
func (sd *fuShuSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	sd.protectNpcDead(npc)
	sd.doorNpcDead(npc)
}

// 初始化防护罩
func (sd *fuShuSceneData) initProtectNpc() {
	warTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiWarTemp(sd.initDefendCampType)
	boTemp := warTemp.GetProtectBiologyTemp()
	pos := warTemp.GetProtectPos()
	so := scene.CreateNPC(scenetypes.OwnerTypeAlliance, 0, 0, 0, 0, boTemp, pos, 0, 0)
	sd.GetScene().AddSceneObject(so)
	sd.protectNpc = so

	gameevent.Emit(chuangshieventtypes.EventTypeChuangShiSceneInitProtect, sd, so)
}

// 初始化玉玺
func (sd *fuShuSceneData) initYuXiNpc() {

	warTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiWarTemp(sd.initDefendCampType)
	boTemp := warTemp.GetYuXiBiologyTemp()
	pos := warTemp.GetYuXiPos()
	// TODO xzk27 玉玺所属问题
	so := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, 0, 0, boTemp, pos, 0, 0)
	sd.GetScene().AddSceneObject(so)
	sd.yuxiNpc = so

	gameevent.Emit(chuangshieventtypes.EventTypeChuangShiSceneInitYuXi, sd, so)
}

//保护罩死亡
func (sd *fuShuSceneData) protectNpcDead(npc scene.NPC) {
	// biology := npc.GetBiologyTemplate()
	// if biology.GetBiologyScriptType() != scenetypes.BiologyScriptTypeMonster {
	// 	return
	// }
	// warTemp := alliancetemplate.GetChuangShiTemplateService().GetChuangShiWarTemp()
	// if biology.Id != warTemp.GetProtectBiologyTemp().Id {
	// 	return
	// }

	//生成玉玺
	sd.initYuXiNpc()
}

//城门死亡
func (sd *fuShuSceneData) doorNpcDead(npc scene.NPC) {
	biology := npc.GetBiologyTemplate()
	if biology.GetBiologyScriptType() != scenetypes.BiologyScriptTypeBuildingMonster {
		return
	}
	sd.currentDoor += 1
	// gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneDoorBroke, sd, nil)
}

// 初始化阵营
func (sd *fuShuSceneData) initCamp(p scene.Player) {
	// p.SetCamp(chuangshitypes.ChuangShiCampTypeFuxi)
	if sd.jieMeng {
		if p.GetCamp() == sd.curDefenCampType {
			p.SwitchPkState(pktype.PkStateCamp, ChuangShiPkCampDefend)
			//设置阵营（针对NPC）
			p.SetFactionType(scenetypes.FactionTypeChengZhanDefendPlayer)
		} else {
			p.SwitchPkState(pktype.PkStateCamp, ChuangShiPkCampAttack)
			p.SetFactionType(scenetypes.FactionTypeChengZhanAttackPlayer)

		}
	} else {
		if p.GetCamp() == sd.curDefenCampType {
			p.SwitchPkState(pktype.PkStateZhenYing, ChuangShiPkCampDefend)
			p.SetFactionType(scenetypes.FactionTypeChengZhanDefendPlayer)
		} else {
			p.SwitchPkState(pktype.PkStateZhenYing, ChuangShiPkCampAttack)
			p.SetFactionType(scenetypes.FactionTypeChengZhanAttackPlayer)
		}
	}
}

//是否占领完成
func (sd *fuShuSceneData) ifReliveOccuoyFinish() bool {
	if sd.collectReliveCampType == chuangshitypes.ChuangShiCampTypeNone {
		return false
	}

	//已经采集过了
	collectFlagTime := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiWarTemp(sd.initDefendCampType).OccupyFlagTime
	now := global.GetGame().GetTimeService().Now()
	elapse := now - sd.collectReliveFlagStartTime
	if elapse > int64(collectFlagTime) {
		return true
	}

	return false
}

//玩家占领复活点旗子了
func (sd *fuShuSceneData) onReliveOccupyFinish() {
	// collectRelivePlayerId := sd.collectRelivePlayerId
	collectReliveCampType := sd.collectReliveCampType
	previousCampType := sd.currentReliveCampType

	sd.collectReliveCampType = 0
	sd.collectRelivePlayerId = 0
	sd.currentReliveCampType = collectReliveCampType
	log.WithFields(
		log.Fields{
			"currentCampType":  collectReliveCampType,
			"previousCampType": previousCampType,
		}).Info("chuangshji:复活点读条完成")

	// TODO : xzk27 采集广播
	// eventData := CreateAllianceReliveCollectEventData(collectReliveCampType, collectRelivePlayerId)
	gameevent.Emit(chuangshieventtypes.EventTypeChuangShiSceneReliveOccupyFinish, sd, nil)
}

//
//------接口方法------------
// 是否排队
func (sd *fuShuSceneData) IfLineup() bool {
	allPlayersNum := int32(len(sd.GetScene().GetAllPlayers()))
	return allPlayersNum >= 200
}

//玉玺被采集
func (sd *fuShuSceneData) YuXiNpcCollectFinish(collecNpc scene.NPC, collectCampType chuangshitypes.ChuangShiCampType) (flag bool) {
	biology := collecNpc.GetBiologyTemplate()
	if biology.GetBiologyScriptType() != scenetypes.BiologyScriptTypeGeneralCollect {
		return
	}

	warTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiWarTemp(sd.initDefendCampType)
	if biology.Id != warTemp.GetYuXiBiologyTemp().Id {
		return
	}
	if sd.curDefenCampType == collectCampType {
		return
	}

	sd.curDefenCampType = collectCampType

	// 阵营转换
	for _, p := range sd.GetScene().GetAllPlayers() {
		sd.initCamp(p)
	}

	//清空所有守卫
	// TODO xzk27
	// for _, n := range sd.GetScene().GetAllNPCS() {
	// 	if n.GetBiologyTemplate().ScriptType != scenetypes.BiologyScriptTypeXianMengNPC {
	// 		continue
	// 	}
	// 	if n.GetOwnerId() != 0 {
	// 		n.Recycle(0)
	// 	}
	// }

	// 生成防护罩
	sd.initProtectNpc()

	flag = true
	return
}

//清除占领复活点读条
func (sd *fuShuSceneData) ClearReliveOccupy() {
	if sd.GetScene().State() == scene.SceneStateFinish || sd.GetScene().State() == scene.SceneStateStopped {
		return
	}
	if sd.collectRelivePlayerId == 0 {
		return
	}
	log.WithFields(
		log.Fields{
			"currentReliveCampType": sd.currentReliveCampType,
			"collectReliveCampType": sd.collectReliveCampType,
		}).Info("chuangshji:清除读条")

	// gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneReliveOccupyStop, sd, sd.collectReliveCampType)
	sd.collectReliveCampType = 0
	sd.collectRelivePlayerId = 0
}

func (sd *fuShuSceneData) GetCurrentDefendCampType() chuangshitypes.ChuangShiCampType {
	return sd.curDefenCampType
}

func (sd *fuShuSceneData) GetInitDefendCampType() chuangshitypes.ChuangShiCampType {
	return sd.initDefendCampType
}

func (sd *fuShuSceneData) GetCurrentDoor() int32 {
	return sd.currentDoor
}

func (sd *fuShuSceneData) GetCollectYuXi() scene.NPC {
	return sd.yuxiNpc
}

func (sd *fuShuSceneData) IsProtectBroken() bool {
	if sd.protectNpc == nil {
		return true
	}

	if sd.protectNpc.IsDead() {
		return true
	}

	return false
}

//城池场景数据
func CreateFuShuSceneData(mapId int32, campType chuangshitypes.ChuangShiCampType, endTime int64) scene.Scene {
	sd := &fuShuSceneData{}
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	sd.jieMeng = true
	sd.initDefendCampType = campType
	sd.curDefenCampType = campType
	return createFuShuScene(mapId, endTime, sd)
}

//城池场景
func createFuShuScene(mapId int32, endTime int64, sd FuShuSceneData) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeChuangShiZhiZhanFuShu {
		return nil
	}

	s = scene.CreateScene(mapTemplate, endTime, sd)
	return s
}
