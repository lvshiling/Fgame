package logic

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sync"
)

var (
	xianTaoSceneMutex sync.Mutex
)

func getXianTaoScene(activityTemplate *gametemplate.ActivityTemplate) (s scene.Scene) {
	//加锁
	xianTaoSceneMutex.Lock()
	defer xianTaoSceneMutex.Unlock()

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
		sd := CreateXianTaoSceneData()
		xiantao := createXianTaoScene(activityTemplate.Mapid, endTime, sd)
		if xiantao == nil {
			return
		}
		s = xiantao
	}
	return
}

func createXianTaoScene(mapId int32, endTime int64, sh scene.SceneDelegate) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeXianTaoDaHui {
		return nil
	}
	s = scene.CreateActivityScene(mapId, endTime, sh)
	return s
}

type XianTaoSceneData interface {
	scene.SceneDelegate
	//获取玩家采集次数
	GetPlayerCollectCount(playerId int64) int32
	//增加玩家采集次数
	AddPlayerCollectCount(playerId int64)
	//设置玩家采集次数
	SetPlayerCollectCount(playerId int64, count int32)
}

// 仙桃大会场景数据
type xiantaoSceneData struct {
	*scene.SceneDelegateBase
	playerCollectMap map[int64]int32
}

func (sd *xiantaoSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("xiantao:仙桃大会应该是同一个场景"))
	}

	if active {
		onExistXianTao(p, sd)
	}
}

func (sd *xiantaoSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("xiantao:仙桃大会应该是同一个场景"))
	}

	onEnterXianTao(p, sd)
}

func (sd *xiantaoSceneData) GetPlayerCollectCount(playerId int64) int32 {
	count, ok := sd.playerCollectMap[playerId]
	if !ok {
		return 0
	}
	return count
}

func (sd *xiantaoSceneData) AddPlayerCollectCount(playerId int64) {
	count, ok := sd.playerCollectMap[playerId]
	if !ok {
		sd.playerCollectMap[playerId] = 1
		return
	}
	count += 1
	sd.playerCollectMap[playerId] = count
	return
}

//场景完成
func (sd *xiantaoSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("xiantao:仙桃大会应该是同一个场景"))
	}

	for _, spl := range s.GetAllPlayers() {
		onFinishXianTao(spl)
	}

	return
}

func (sd *xiantaoSceneData) SetPlayerCollectCount(playerId int64, count int32) {
	sd.playerCollectMap[playerId] = count
	return
}

func CreateXianTaoSceneData() XianTaoSceneData {
	sd := &xiantaoSceneData{}
	sd.playerCollectMap = make(map[int64]int32)
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return sd
}
