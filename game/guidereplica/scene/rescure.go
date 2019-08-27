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
	guidereplica.RegisterGuideReplicaSdHandler(guidereplicatypes.GuideReplicaTypeRescure, guidereplica.GuideReplicaSdHandlerFunc(CreateGuideRescureSceneData))
}

type GuideRescureSceneData interface {
	GetGuideTemp() *gametemplate.GuideReplicaTemplate
}

// 救援小医仙
type guideRescureSceneData struct {
	*scene.SceneDelegateBase
	ownerId          int64
	questId          int32
	guidereplicaTemp *gametemplate.GuideReplicaTemplate
}

//玩家进入
func (sd *guideRescureSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}

	// 场景信息
	sd.onPlayerEnter(p, s)
}

//玩家退出
func (sd *guideRescureSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}

	//主动退出 结束副本
	sd.GetScene().Stop(true, false)
}

//场景完成
func (sd *guideRescureSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.GetScene() != s {
		panic(fmt.Errorf("guidereplica:引导副本应该是同一个场景"))
	}

	spl := sd.GetScene().GetPlayer(sd.ownerId)
	if spl == nil {
		return
	}

	gameevent.Emit(guidereplicaeventtypes.EventTypeGuideFinish, spl, success)
}

func (sd *guideRescureSceneData) GetGuideTemp() *gametemplate.GuideReplicaTemplate {
	return sd.guidereplicaTemp
}

func (sd *guideRescureSceneData) onPlayerEnter(spl scene.Player, s scene.Scene) {
	startTime := s.GetStartTime()
	scMsg := pbutil.BuildSCGuideReplicaSceneInfo(startTime, sd.guidereplicaTemp.MapId, int32(sd.guidereplicaTemp.GetGuideType()), sd.questId)
	spl.SendMsg(scMsg)
}

func CreateGuideRescureSceneData(pl player.Player, temp *gametemplate.GuideReplicaTemplate, questId int32) scene.SceneDelegate {
	csd := &guideRescureSceneData{
		ownerId:          pl.GetId(),
		guidereplicaTemp: temp,
		questId:          questId,
	}
	csd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return csd
}
