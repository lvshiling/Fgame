package scene

import (
	pktype "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
)

type ZhongLiSceneData interface {
	scene.SceneDelegate
	IfLineup() bool
}

//中立城池
type zhongLiSceneData struct {
	*scene.SceneDelegateBase
}

func (d *zhongLiSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	// p.SetCamp(chuangshitypes.ChuangShiCampTypeFuxi)

	p.SwitchPkState(pktype.PkStateZhenYing, pktype.PkCommonCampDefault)

}


func (sd *zhongLiSceneData) IfLineup() bool {
	allPlayersNum := int32(len(sd.GetScene().GetAllPlayers()))
	return allPlayersNum >= 200
}


//城池场景数据
func CreateZhongLiSceneData(mapId int32, endTime int64) scene.Scene {
	sd := &zhongLiSceneData{}
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return createZhongLiScene(mapId, endTime, sd)
}

//城池场景
func createZhongLiScene(mapId int32, endTime int64, sd ZhongLiSceneData) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeChuangShiZhiZhanZhongLi {
		return nil
	}

	s = scene.CreateScene(mapTemplate, endTime, sd)
	return s
}
