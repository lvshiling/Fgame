package scene

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	welfaresceneeventtypes "fgame/fgame/game/welfarescene/event/types"
	"fgame/fgame/game/welfarescene/pbutil"
	welfarescenetypes "fgame/fgame/game/welfarescene/types"
	"fmt"
)

func init() {
	RegisterWelfareSceneSdHandler(welfarescenetypes.WelfareSceneTypeQiYuDao, WelfareSceneSdHandlerFunc(CreateWelfareQiYuBossSceneData))
}

// d)当BOSS刷新时，前端弹出提示框

type WelfareQiYuBossSceneData interface {
	GetWelfareSceneTemp() *gametemplate.WelfareSceneTemplate
	GetBossMap() map[int64]scene.NPC
	GetScene() scene.Scene
	GetGroupId() int32
	GetCollectNum() int32
}

//奇遇岛
type qiYuBossSceneData struct {
	*scene.SceneDelegateBase
	wsTemp          *gametemplate.WelfareSceneTemplate
	groupId         int32
	bossMap         map[int64]scene.NPC
	collectNum      int32
	lastRefreshTime int64
}

func (sd *qiYuBossSceneData) OnSceneTick(s scene.Scene) {
	now := global.GetGame().GetTimeService().Now()
	qiYuTemp := sd.wsTemp.GetQiYuTemp()
	timeList := qiYuTemp.GetRefreshTimeList(now)
	for _, freshTime := range timeList {
		if now < freshTime || sd.lastRefreshTime >= freshTime {
			continue
		}

		sd.refreshBiology(s)
		sd.lastRefreshTime = now
	}

	return
}

func (sd *qiYuBossSceneData) OnSceneStart(s scene.Scene) {
	sd.SceneDelegateBase.OnSceneStart(s)

	//刷新怪物
	now := global.GetGame().GetTimeService().Now()
	qiYuTemp := sd.wsTemp.GetQiYuTemp()
	start := qiYuTemp.GetStartTime(now)
	end := qiYuTemp.GetEndTime(now)
	if now < start || now > end {
		return
	}

	sd.refreshBiology(s)
	sd.lastRefreshTime = now

	return
}

//刷新怪物
func (sd *qiYuBossSceneData) refreshBiology(s scene.Scene) {
	var addNpcList []scene.NPC
	if len(sd.bossMap) > 0 {
		for _, npc := range sd.bossMap {
			if !npc.IsDead() {
				continue
			}

			sn := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, 0, 0, npc.GetBiologyTemplate(), npc.GetPosition(), 0, 0)
			s.AddSceneObject(sn)
			addNpcList = append(addNpcList, sn)

			delete(sd.bossMap, npc.GetId())
		}

	} else {
		for _, boPosTemp := range sd.wsTemp.GetQiYuTemp().GetBiologyPosList() {
			biologyTemp := boPosTemp.GetBiologyTemp()
			sn := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, 0, 0, biologyTemp, boPosTemp.GetPos(), 0, 0)
			s.AddSceneObject(sn)

			addNpcList = append(addNpcList, sn)
		}
	}
	sd.addBossNpc(addNpcList)

	gameevent.Emit(welfaresceneeventtypes.EventTypeWelfareSceneRefresh, sd, nil)
	return
}

func (sd *qiYuBossSceneData) addBossNpc(newNpcList []scene.NPC) {
	for _, newN := range newNpcList {
		if newN.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeQiYuDao {
			continue
		}

		sd.bossMap[newN.GetId()] = newN
	}
}

//玩家进入
func (sd *qiYuBossSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("welfarescene:运营活动副本应该是同一个场景"))
	}

	// 场景信息
	sd.onPlayerEnter(p, s)
}

func (sd *qiYuBossSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("welfarescene:运营活动副本应该是同一个场景"))
	}

	gameevent.Emit(welfaresceneeventtypes.EventTypeWelfareSceneFinish, sd, sd.wsTemp.GroupId)
	return
}

func (sd *qiYuBossSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("welfarescene:运营活动副本应该是同一个场景"))
	}

	if npc.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeGeneralCollect {
		return
	}

	sd.collectNum += 1
	return
}

func (sd *qiYuBossSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("welfarescene:运营活动副本应该是同一个场景"))
	}

	if npc.GetBiologyTemplate().GetBiologyScriptType() == scenetypes.BiologyScriptTypeGeneralCollect {
		sd.collectNum -= 1
	}

	npcMap := make(map[int64]scene.NPC)
	_, ok := sd.bossMap[npc.GetId()]
	if ok {
		npcMap[npc.GetId()] = npc
	}

	scMsg := pbutil.BuildSCWelfareSceneDataChangedNotice(int32(sd.wsTemp.Id), npcMap, sd.collectNum)
	s.BroadcastMsg(scMsg)
	return
}

func (sd *qiYuBossSceneData) GetWelfareSceneTemp() *gametemplate.WelfareSceneTemplate {
	return sd.wsTemp
}

func (sd *qiYuBossSceneData) GetBossMap() map[int64]scene.NPC {
	return sd.bossMap
}

func (sd *qiYuBossSceneData) GetGroupId() int32 {
	return sd.groupId
}

func (sd *qiYuBossSceneData) GetCollectNum() int32 {
	return sd.collectNum
}

func (sd *qiYuBossSceneData) onPlayerEnter(spl scene.Player, s scene.Scene) {
	startTime := s.GetStartTime()
	scMsg := pbutil.BuildSCWelfareSceneInfo(startTime, int32(sd.wsTemp.Id), sd.bossMap, sd.collectNum)
	spl.SendMsg(scMsg)
}

func CreateWelfareQiYuBossSceneData(groupId int32, temp *gametemplate.WelfareSceneTemplate) scene.SceneDelegate {
	csd := &qiYuBossSceneData{
		wsTemp:  temp,
		groupId: groupId,
		bossMap: make(map[int64]scene.NPC),
	}
	csd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return csd
}
