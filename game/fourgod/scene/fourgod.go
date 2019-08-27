package scene

import (
	fourgodtemplate "fgame/fgame/game/fourgod/template"
	"fgame/fgame/game/scene/scene"

	"sync"
)

type FourGodSceneData interface {
	//获取结束时间
	GetEndTime() int64
	GetScene() scene.Scene
}

//四神遗迹
type fourGodSceneData struct {
	rwm sync.Mutex
	//主战场
	war FourGodWarSceneData
	//结束时间
	endTime int64
}

//四神遗迹场景数据
func CreateFourGodSceneData(warId int32, endTime int64) FourGodSceneData {
	fsd := &fourGodSceneData{}
	fsd.endTime = endTime
	defaultPos, _ := fourgodtemplate.GetFourGodTemplateService().GetFourGodSpecialDefaultPos()

	warData := createFourGodWarSceneData(defaultPos)
	warScene := createFourGodWarScene(warId, endTime, warData)
	if warScene == nil {
		return nil
	}
	fsd.war = warData
	return fsd
}

//获取活动结束时间
func (sd *fourGodSceneData) GetEndTime() int64 {
	return sd.endTime
}

//获取活动结束时间
func (sd *fourGodSceneData) GetScene() scene.Scene {
	return sd.war.GetScene()
}

type FourGodSubSceneData interface {
	scene.SceneDelegate
}
