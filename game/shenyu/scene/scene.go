package scene

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	shenyueventtypes "fgame/fgame/game/shenyu/event/types"
	shenyutypes "fgame/fgame/game/shenyu/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

type ShenYuSceneData interface {
	scene.SceneDelegate
	//获取神域模板
	GetShenYuTemplate() *gametemplate.ShenYuTemplate
	//获取活动结束时间
	GetActivityEndTime() int64
}

// 神域场景数据
type shenyuSceneData struct {
	*scene.SceneDelegateBase
	shenyuTemplate  *gametemplate.ShenYuTemplate //
	activityEndTime int64                        //活动结束时间
	lastLuckyTime   int64                        //上次幸运奖时间
}

func (sd *shenyuSceneData) OnSceneTick(s scene.Scene) {
	now := global.GetGame().GetTimeService().Now()
	if now-sd.lastLuckyTime < int64(sd.shenyuTemplate.LuckyRewTime) {
		return
	}

	gameevent.Emit(shenyueventtypes.EventTypeShenYuLuckRew, sd, nil)
	sd.lastLuckyTime = now
	return
}

func (sd *shenyuSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("shenyu:神域应该是同一个场景"))
	}

	//去除玩家排行榜数据
	s.RemovePlayer(shenyutypes.ShenYuSceneRankTypeKey, p.GetId())
	return
}

func (sd *shenyuSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("shenyu:神域应该是同一个场景"))
	}

	// 神域结束
	gameevent.Emit(shenyueventtypes.EventTypeShenYuFinish, sd, nil)
	return
}

func (sd *shenyuSceneData) OnSceneStop(s scene.Scene) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("shenyu:神域应该是同一个场景"))
	}
	//
	gameevent.Emit(shenyueventtypes.EventTypeShenYuStop, sd, nil)

	return
}

func (sd *shenyuSceneData) GetShenYuTemplate() *gametemplate.ShenYuTemplate {
	return sd.shenyuTemplate
}

func (sd *shenyuSceneData) GetActivityEndTime() int64 {
	return sd.activityEndTime
}

func CreateShenYuSceneData(temp *gametemplate.ShenYuTemplate, endTime int64) ShenYuSceneData {
	sd := &shenyuSceneData{}
	sd.shenyuTemplate = temp
	sd.activityEndTime = endTime
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return sd
}
