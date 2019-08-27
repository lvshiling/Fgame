package scene

import (
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	welfarescenetypes "fgame/fgame/game/welfarescene/types"
	"fmt"
)

type WelfareSceneSdHandler interface {
	WelfareSceneSd(groupId int32, temp *gametemplate.WelfareSceneTemplate) (sh scene.SceneDelegate)
}

type WelfareSceneSdHandlerFunc func(groupId int32, temp *gametemplate.WelfareSceneTemplate) (sh scene.SceneDelegate)

func (f WelfareSceneSdHandlerFunc) WelfareSceneSd(groupId int32, temp *gametemplate.WelfareSceneTemplate) (sh scene.SceneDelegate) {
	return f(groupId, temp)
}

var (
	welfareSceneSdHandlerMap = make(map[welfarescenetypes.WelfareSceneType]WelfareSceneSdHandler)
)

func RegisterWelfareSceneSdHandler(typ welfarescenetypes.WelfareSceneType, h WelfareSceneSdHandler) {
	_, ok := welfareSceneSdHandlerMap[typ]
	if ok {
		panic(fmt.Errorf("welfarescene:repeat register sceneData handler; type:%d ", typ))
	}
	welfareSceneSdHandlerMap[typ] = h
}

func GetWelfareSceneSd(typ welfarescenetypes.WelfareSceneType, groupId int32, temp *gametemplate.WelfareSceneTemplate) (sh scene.SceneDelegate) {
	h, ok := welfareSceneSdHandlerMap[typ]
	if !ok {
		return nil
	}

	return h.WelfareSceneSd(groupId, temp)
}
