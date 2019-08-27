package scene

import (
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	pktype "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
)

type FuShuSceneData interface {
	scene.SceneDelegate
	IfLineup() bool
}

//附属城池
type fuShuSceneData struct {
	*scene.SceneDelegateBase
	jieMeng    bool
	defendCamp chuangshitypes.ChuangShiCampType
}

func (d *fuShuSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	// p.SetCamp(chuangshitypes.ChuangShiCampTypeFuxi)
	if d.jieMeng {
		if p.GetCamp() == d.defendCamp {
			p.SwitchPkState(pktype.PkStateCamp, ChuangShiPkCampDefend)
		} else {
			p.SwitchPkState(pktype.PkStateCamp, ChuangShiPkCampAttack)

		}
	} else {
		if p.GetCamp() == d.defendCamp {
			p.SwitchPkState(pktype.PkStateZhenYing, ChuangShiPkCampDefend)
		} else {
			p.SwitchPkState(pktype.PkStateZhenYing, ChuangShiPkCampAttack)
		}
	}
}

func (sd *fuShuSceneData) IfLineup() bool {
	allPlayersNum := int32(len(sd.GetScene().GetAllPlayers()))
	return allPlayersNum >= 200
}

//城池场景数据
func CreateFuShuSceneData(mapId int32, campType chuangshitypes.ChuangShiCampType, endTime int64) scene.Scene {
	sd := &fuShuSceneData{}
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	sd.jieMeng = true
	sd.defendCamp = campType
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
