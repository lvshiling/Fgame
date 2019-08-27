package scene

import (
	pktype "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
)

type MainSceneData interface {
	scene.SceneDelegate
	IfLineup() bool
}

//主城
type mainSceneData struct {
	*scene.SceneDelegateBase
}

func (d *mainSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	// p.SetCamp(chuangshitypes.ChuangShiCampTypeFuxi)

	p.SwitchPkState(pktype.PkStateZhenYing, pktype.PkCommonCampDefault)

}


func (sd *mainSceneData) IfLineup() bool {
	allPlayersNum := int32(len(sd.GetScene().GetAllPlayers()))
	return allPlayersNum >= 200
}

//城池场景数据
func CreateMainSceneData(mapId int32, endTime int64) scene.Scene {
	sd := &mainSceneData{}
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return createMainScene(mapId, endTime, sd)
}

//城池场景
func createMainScene(mapId int32, endTime int64, sd MainSceneData) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeChuangShiZhiZhanMain {
		return nil
	}

	s = scene.CreateScene(mapTemplate, endTime, sd)
	return s
}
