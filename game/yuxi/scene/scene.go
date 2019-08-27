package scene

import (
	"fgame/fgame/game/activity/activity"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	yuxitypes "fgame/fgame/game/yuxi/types"
	"fmt"
	"sync"
)

var (
	yuXiSceneMutex sync.Mutex
)

func GetYuXiScene(activityTemplate *gametemplate.ActivityTemplate) (s scene.Scene, flag bool) {
	//加锁
	yuXiSceneMutex.Lock()
	defer yuXiSceneMutex.Unlock()

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

	// 是否结束
	if activity.GetActivityService().IsActivityEnd(activityTemplate.GetActivityType(), endTime) {
		flag = true
		return
	}

	s = scene.GetSceneService().GetActivitySceneByMapId(activityTemplate.Mapid)
	if s == nil {
		//创建场景
		sd := CreateYuXiSceneData()
		yuxi := createYuXiScene(activityTemplate.Mapid, endTime, sd)
		if yuxi == nil {
			return
		}
		s = yuxi
	}
	return
}

func createYuXiScene(mapId int32, endTime int64, sh scene.SceneDelegate) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeYuXi {
		return nil
	}
	s = scene.CreateActivityScene(mapId, endTime, sh)
	return s
}

type YuXiSceneData interface {
	scene.SceneDelegate
	CollectYuXiFinish(pl scene.Player) int64 //采集玉玺成功
	GetOwerYuXinInfo() (scene.Player, int64) //玉玺拥有者信息
	YuXiRest(pl scene.Player, rebornType yuxitypes.YuXiReborType)
}

// 玉玺之战场景数据
type yuxiSceneData struct {
	*scene.SceneDelegateBase
	yuxiNpc           scene.NPC               //玉玺
	yuXiOwner         scene.Player            //持有玉玺玩家
	startTime         int64                   //持有玉玺开始时间
	lastBroadcastTime int64                   //上次玉玺广播时间
	attendAlMap       map[int64]*allianceData //参战的仙盟
}

//场景开始
func (sd *yuxiSceneData) OnSceneStart(s scene.Scene) {
	sd.SceneDelegateBase.OnSceneStart(s)

	sd.initYuXiCollect(yuxitypes.YuXiReborTypeInitNone)
	return
}

// 场景定时
func (sd *yuxiSceneData) OnSceneTick(s scene.Scene) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("yuxi:玉玺之战应该是同一个场景"))
	}

	// 超过时间
	if sd.isCanEnd() {
		sd.GetScene().Finish(true)
		return
	}

	// 广播玉玺位置
	now := global.GetGame().GetTimeService().Now()
	if sd.isBroadcastYuXiInfo(now) {
		sd.onBroadcastYuXiInfo()
		sd.lastBroadcastTime = now
	}

	return
}

// 玩家死亡
func (sd *yuxiSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("yuxi:玉玺之战应该是同一个场景"))
	}

	sd.onYuXiReset(p, yuxitypes.YuXiReborTypePlayerDead)
	return
}

// 玩家退出场景
func (sd *yuxiSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("yuxi:玉玺之战应该是同一个场景"))
	}

	sd.onYuXiReset(p, yuxitypes.YuXiReborTypePlayerExitScene)
}

//场景完成
func (sd *yuxiSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("yuxi:玉玺之战应该是同一个场景"))
	}

	sd.onFinishYuXi()
	return
}

// 玩家进入
func (sd *yuxiSceneData) OnScenePlayerEnter(s scene.Scene, spl scene.Player) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("yuxi:玉玺之战应该是同一个场景"))
	}

	sd.onPlayerEnter(spl)
	return
}

//采集玉玺成功
func (sd *yuxiSceneData) CollectYuXiFinish(pl scene.Player) int64 {
	now := global.GetGame().GetTimeService().Now()
	sd.yuXiOwner = pl
	sd.startTime = now

	sd.onBroadcastYuXiCollect(yuxitypes.YuXiReborTypeInitNone)
	return now
}

//采集玉玺成功
func (sd *yuxiSceneData) GetOwerYuXinInfo() (scene.Player, int64) {
	return sd.yuXiOwner, sd.startTime
}

//玉玺重置
func (sd *yuxiSceneData) YuXiRest(spl scene.Player, rebornType yuxitypes.YuXiReborType) {
	sd.onYuXiReset(spl, rebornType)
}

func CreateYuXiSceneData() YuXiSceneData {
	sd := &yuxiSceneData{}
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	sd.attendAlMap = make(map[int64]*allianceData)
	return sd
}
