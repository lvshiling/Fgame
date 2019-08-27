package teamcopy

import (
	teamcopyeventtypes "fgame/fgame/cross/teamcopy/event/types"
	scene "fgame/fgame/cross/teamcopy/scene"
	gameevent "fgame/fgame/game/event"
	gamescene "fgame/fgame/game/scene/scene"
	teamcopytemplate "fgame/fgame/game/teamcopy/template"
)

const (
	maxServer = 10
)

type teamCopyData struct {
	sceneData scene.TeamCopySceneData
}

func (s *teamCopyData) init() {
	s.sceneData = nil

}

func createSceneData(teamObj *scene.TeamObject) (data *teamCopyData) {
	data = &teamCopyData{}
	data.init()
	data.sceneData = scene.CreateTeamCopySceneData(teamObj)

	purpose := teamObj.GetTeamPurpose()
	teamCopyTemplate := teamcopytemplate.GetTeamCopyTemplateService().GetTeamCopyTempalte(purpose)
	if teamCopyTemplate == nil {
		return
	}
	mapTemplate := teamCopyTemplate.GetMapTemplate()
	mapId := int32(mapTemplate.Id)
	gamescene.CreateFuBenScene(mapId, data.sceneData)
	gameevent.Emit(teamcopyeventtypes.EventTypeTeamCopySceneCreateFinish, data.sceneData, nil)
	return data
}

func (s *teamCopyData) GetTeamCopySceneData() scene.TeamCopySceneData {
	return s.sceneData
}

func (s *teamCopyData) teamCopySceneFinish() {
	s.sceneData = nil
}
