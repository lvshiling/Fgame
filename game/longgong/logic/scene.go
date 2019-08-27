package logic

import (
	collectlogic "fgame/fgame/game/collect/logic"
	collectnpc "fgame/fgame/game/collect/npc"
	"fgame/fgame/game/global"
	"fgame/fgame/game/longgong/pbutil"
	longgongtemplate "fgame/fgame/game/longgong/template"
	longgongtypes "fgame/fgame/game/longgong/types"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sync"
)

var (
	longGongSceneMutex sync.Mutex
)

func getLongGongScene(activityTemplate *gametemplate.ActivityTemplate) (s scene.Scene) {
	//加锁
	longGongSceneMutex.Lock()
	defer longGongSceneMutex.Unlock()

	s = scene.GetSceneService().GetActivitySceneByMapId(activityTemplate.Mapid)
	if s == nil {
		now := global.GetGame().GetTimeService().Now()
		openTime := global.GetGame().GetServerTime()
		mergeTime := merge.GetMergeService().GetMergeTime()
		activityTimeTemplate, err := activityTemplate.GetActivityTimeTemplate(now, openTime, mergeTime)
		if err != nil {
			return
		}
		if activityTimeTemplate == nil {
			return
		}
		endTime, _ := activityTimeTemplate.GetEndTime(now)

		//创建场景
		sd := CreateLongGongSceneData()
		longgong := createLongGongScene(activityTemplate.Mapid, endTime, sd)
		if longgong == nil {
			return
		}
		s = longgong
	}
	return
}

func createLongGongScene(mapId int32, endTime int64, sh scene.SceneDelegate) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeLongGong {
		return nil
	}
	s = scene.CreateActivityScene(mapId, endTime, sh)
	return s
}

//黑龙boss信息
type heiLongInfo struct {
	status longgongtypes.HeiLongStatusType //黑龙状态
	npc    scene.NPC                       //场景npc
}

func newHeiLongInfo() *heiLongInfo {
	heiLong := &heiLongInfo{
		status: longgongtypes.HeiLongStatusTypeInit,
	}
	return heiLong
}

func (be *heiLongInfo) GetStatus() longgongtypes.HeiLongStatusType {
	return be.status
}

func (be *heiLongInfo) GetNpc() scene.NPC {
	return be.npc
}

type LongGongSceneData interface {
	scene.SceneDelegate
	//获取玩家黑龙财宝采集次数
	GetPlayerTreasureCollectCount(playerId int64) int32
	//增加玩家黑龙财宝采集次数
	AddPlayerTreasureCollectCount(playerId int64)
	//Gm使用设置玩家黑龙财宝采集次数
	GmSetPlayerTreasureCollectCount(playerId int64, count int32)
	//获取珍珠采集数
	GetPearlCollectCount() int32
	//增加珍珠采集数
	AddPearlCollectCount(val int32)
	//Gm使用设置珍珠采集数
	GmSetPearlCollectCount(count int32) (isSucceed bool)
	//获取黑龙boss信息
	GetHeiLongBossInfo() *heiLongInfo
	//处理黑龙boss死亡
	DealHeiLongBossDead()
	//龙宫珍珠采集数和Boss出生广播
	PearlCountWithBossBornBroadcast(isBorn bool)
	//获取财宝采集点
	GetTreasureCollectPointNpc() scene.NPC
}

// 龙宫探宝场景数据
type longgongSceneData struct {
	*scene.SceneDelegateBase
	playerCollectMap  map[int64]int32
	pearlCollectCount int32
	heiLongBoss       *heiLongInfo
	cpNpc             scene.NPC
	//上次系统采集珍珠时间
	lastSysAddPearlCountTime int64
}

func (sd *longgongSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {

}

func (sd *longgongSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("longgong:龙宫探宝应该是同一个场景"))
	}

	onEnterLongGong(p, sd)
}

//场景完成
func (sd *longgongSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("longgong:龙宫探宝应该是同一个场景"))
	}

	onFinishLongGong(s)
	return
}

//场景心跳
func (sd *longgongSceneData) OnSceneTick(s scene.Scene) {
	if s.State() == scene.SceneStateFinish || s.State() == scene.SceneStateStopped {
		return
	}
	if sd.heiLongBoss.status != longgongtypes.HeiLongStatusTypeInit {
		return
	}

	longGongService := longgongtemplate.GetLongGongTemplateService()
	constTemp := longGongService.GetLongGongConstTemplate()
	now := global.GetGame().GetTimeService().Now()
	elapse := now - sd.lastSysAddPearlCountTime
	//已经采集过了
	if elapse < int64(constTemp.XuJiaCaiJiAddTime) {
		return
	}
	sd.lastSysAddPearlCountTime = now
	//增加珍珠采集数
	sd.AddPearlCollectCount(constTemp.XuJiaCaiJiAddCount)
}

//龙宫珍珠采集数和Boss出生广播
func (sd *longgongSceneData) PearlCountWithBossBornBroadcast(isBorn bool) {
	splM := sd.GetScene().GetAllPlayers()
	for _, spl := range splM {
		if isBorn {
			curHp := sd.heiLongBoss.npc.GetHP()
			scMsg := pbutil.BuildSCLonggongSceneValBroadcast(sd.pearlCollectCount, sd.heiLongBoss.status, curHp)
			spl.SendMsg(scMsg)
		} else {
			scMsg := pbutil.BuildSCLonggongScenePearlCountBroadcast(sd.pearlCollectCount)
			spl.SendMsg(scMsg)
		}
	}
}

func (sd *longgongSceneData) GetHeiLongBossInfo() *heiLongInfo {
	return sd.heiLongBoss
}

func (sd *longgongSceneData) DealHeiLongBossDead() {
	sd.heiLongBoss.npc = nil
	sd.heiLongBoss.status = longgongtypes.HeiLongStatusTypeDead
	//刷出宝藏采集物
	longGongService := longgongtemplate.GetLongGongTemplateService()
	constTemp := longGongService.GetLongGongConstTemplate()
	collectTemp := constTemp.GetCollectBiology()
	cn := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, int64(0), int32(0), collectTemp, constTemp.GetPos(), 0, 0)
	if cn == nil {
		return
	}
	sd.cpNpc = cn
	sd.GetScene().AddSceneObject(cn)

	//广播一下
	splM := sd.GetScene().GetAllPlayers()
	for _, spl := range splM {
		scMsg := pbutil.BuildSCLonggongSceneBossDieBroadcast(sd.heiLongBoss.status, cn.GetId())
		spl.SendMsg(scMsg)
	}
	return
}

//获取财宝采集点
func (sd *longgongSceneData) GetTreasureCollectPointNpc() scene.NPC {
	return sd.cpNpc
}

func (sd *longgongSceneData) GetPlayerTreasureCollectCount(playerId int64) int32 {
	count, ok := sd.playerCollectMap[playerId]
	if !ok {
		return 0
	}
	return count
}

func (sd *longgongSceneData) AddPlayerTreasureCollectCount(playerId int64) {
	count, ok := sd.playerCollectMap[playerId]
	if !ok {
		sd.playerCollectMap[playerId] = 1
		return
	}
	count += 1
	sd.playerCollectMap[playerId] = count
	return
}

func (sd *longgongSceneData) GmSetPlayerTreasureCollectCount(playerId int64, count int32) {
	sd.playerCollectMap[playerId] = count
	return
}

func (sd *longgongSceneData) GetPearlCollectCount() int32 {
	return sd.pearlCollectCount
}

func (sd *longgongSceneData) AddPearlCollectCount(val int32) {
	sd.pearlCollectCount += val
	//采集一定数量的珍珠刷出黑龙boss
	isBorn := false
	longGongService := longgongtemplate.GetLongGongTemplateService()
	constTemp := longGongService.GetLongGongConstTemplate()
	if sd.pearlCollectCount >= constTemp.BossNeedCaiJiCount && sd.heiLongBoss.status == longgongtypes.HeiLongStatusTypeInit {
		bossTemp := constTemp.GetBossBiology()
		bn := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, int64(0), int32(0), bossTemp, constTemp.GetPos(), 0, 0)
		if bn != nil {
			sd.GetScene().AddSceneObject(bn)
			sd.heiLongBoss.status = longgongtypes.HeiLongStatusTypeLive
			sd.heiLongBoss.npc = bn
			isBorn = true

			//boss出生移除龙宫所有珍珠采集物
			pnM := sd.GetScene().GetNPCS(scenetypes.BiologyScriptTypePearl)
			for _, pn := range pnM {
				ccn, ok := pn.(*collectnpc.CollectChooseNPC)
				if !ok {
					continue
				}
				ccnPl := ccn.GetCollect().GetPlayer()
				if ccnPl != nil {
					collectlogic.CollectInterrupt(ccnPl, ccn)
				}
				sd.GetScene().RemoveSceneObject(pn, false)
			}
		}
	}
	//广播一下
	sd.PearlCountWithBossBornBroadcast(isBorn)
	return
}

func (sd *longgongSceneData) GmSetPearlCollectCount(count int32) (isSucceed bool) {
	if count <= sd.pearlCollectCount {
		return
	}
	s := sd.GetScene()
	//处理黑龙boss和财宝采集点
	isBorn := false
	longGongService := longgongtemplate.GetLongGongTemplateService()
	constTemp := longGongService.GetLongGongConstTemplate()
	if count >= constTemp.BossNeedCaiJiCount && sd.heiLongBoss.status == longgongtypes.HeiLongStatusTypeInit {
		bossTemp := constTemp.GetBossBiology()
		bn := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, int64(0), int32(0), bossTemp, constTemp.GetPos(), 0, 0)
		if bn != nil {
			s.AddSceneObject(bn)
			sd.heiLongBoss.status = longgongtypes.HeiLongStatusTypeLive
			sd.heiLongBoss.npc = bn
			isBorn = true
			//boss出生移除龙宫所有珍珠采集物
			pnM := sd.GetScene().GetNPCS(scenetypes.BiologyScriptTypePearl)
			for _, pn := range pnM {
				ccn, ok := pn.(*collectnpc.CollectChooseNPC)
				if !ok {
					continue
				}
				ccnPl := ccn.GetCollect().GetPlayer()
				if ccnPl != nil {
					collectlogic.CollectInterrupt(ccnPl, ccn)
				}
				sd.GetScene().RemoveSceneObject(pn, false)
			}
		}
	}
	sd.pearlCollectCount = count
	isSucceed = true
	//广播一下
	sd.PearlCountWithBossBornBroadcast(isBorn)
	return
}

func CreateLongGongSceneData() LongGongSceneData {
	sd := &longgongSceneData{}
	sd.playerCollectMap = make(map[int64]int32)
	sd.pearlCollectCount = int32(0)
	sd.lastSysAddPearlCountTime = int64(0)
	sd.heiLongBoss = newHeiLongInfo()
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return sd
}
