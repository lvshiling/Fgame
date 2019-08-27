package scene

import (
	gameevent "fgame/fgame/game/event"
	guidereplicaeventtypes "fgame/fgame/game/guidereplica/event/types"
	guidereplica "fgame/fgame/game/guidereplica/guidereplica"
	"fgame/fgame/game/guidereplica/pbutil"
	guidereplicatypes "fgame/fgame/game/guidereplica/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

func init() {
	guidereplica.RegisterGuideReplicaSdHandler(guidereplicatypes.GuideReplicaTypeCatDog, guidereplica.GuideReplicaSdHandlerFunc(CreateGuideCatDogSceneData))
}

type GuideCatDogSceneData interface {
	GetKillMap() map[guidereplicatypes.CatDogKillType]int32
	GetGuideTemp() *gametemplate.GuideReplicaTemplate
}

// 猫狗大战
type guideCatDogSceneData struct {
	*scene.SceneDelegateBase
	ownerId          int64
	questId          int32
	guidereplicaTemp *gametemplate.GuideReplicaTemplate
	killMap          map[guidereplicatypes.CatDogKillType]int32 //击杀数量记录
}

//怪物死亡
func (sd *guideCatDogSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}
	boId := int32(npc.GetBiologyTemplate().Id)

	catDogTemp := sd.guidereplicaTemp.GetCatDogGuideTemp()
	catId := catDogTemp.CatBiologyId
	dogId := catDogTemp.DogBiologyId
	switch boId {
	case catId:
		sd.killMap[guidereplicatypes.CatDogKillTypeCat] += 1
	case dogId:
		sd.killMap[guidereplicatypes.CatDogKillTypeDog] += 1
	}

	sd.onKillChangedNotice()
}

func (sd *guideCatDogSceneData) OnSceneBiologyAllDead(s scene.Scene) {
	sd.GetScene().Finish(true)
	return
}

//玩家进入
func (sd *guideCatDogSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}

	// 场景信息
	sd.onPlayerEnter(p, s)
}

//玩家退出
func (sd *guideCatDogSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}

	//主动退出 结束副本
	sd.GetScene().Stop(true, false)
}

//场景完成
func (sd *guideCatDogSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}

	// 自然结束副本成功
	if !success && useTime >= int64(s.MapTemplate().PointsTime) {
		success = true
	}

	spl := sd.GetScene().GetPlayer(sd.ownerId)
	if spl == nil {
		return
	}

	gameevent.Emit(guidereplicaeventtypes.EventTypeGuideFinish, spl, success)
}

func (sd *guideCatDogSceneData) GetKillMap() map[guidereplicatypes.CatDogKillType]int32 {
	return sd.killMap
}

func (sd *guideCatDogSceneData) GetGuideTemp() *gametemplate.GuideReplicaTemplate {
	return sd.guidereplicaTemp
}

func (sd *guideCatDogSceneData) onPlayerEnter(spl scene.Player, s scene.Scene) {
	startTime := s.GetStartTime()
	scMsg := pbutil.BuildSCGuideReplicaSceneInfoWithCatDog(startTime, sd.guidereplicaTemp.MapId, int32(sd.guidereplicaTemp.GetGuideType()), sd.questId, sd.killMap)
	spl.SendMsg(scMsg)
}

func (sd *guideCatDogSceneData) onKillChangedNotice() {
	spl := sd.GetScene().GetPlayer(sd.ownerId)
	if spl == nil {
		return
	}
	scMsg := pbutil.BuildSCGuideReplicaSceneDataChangedNoticeWithCatDog(int32(sd.guidereplicaTemp.GetGuideType()), sd.killMap)
	spl.SendMsg(scMsg)
}

func CreateGuideCatDogSceneData(pl player.Player, temp *gametemplate.GuideReplicaTemplate, questId int32) scene.SceneDelegate {
	csd := &guideCatDogSceneData{
		ownerId:          pl.GetId(),
		guidereplicaTemp: temp,
		questId:          questId,
		killMap:          make(map[guidereplicatypes.CatDogKillType]int32),
	}
	csd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return csd
}
