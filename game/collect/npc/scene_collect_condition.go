package npc

import (
	scene "fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

const (
	defaultCollectCondition = true
)

type SceneCollectConditionHandler interface {
	SceneCollectCondition(scene.Player, scene.CollectNPC) bool
}

type SceneCollectConditionHandlerFunc func(scene.Player, scene.CollectNPC) bool

func (f SceneCollectConditionHandlerFunc) SceneCollectCondition(pl scene.Player, cn scene.CollectNPC) bool {
	return f(pl, cn)
}

var (
	sceneCollectConditionHandlerMap = make(map[scenetypes.SceneType]SceneCollectConditionHandler)
)

func RegisterSceneCollectConditionHandler(tag scenetypes.SceneType, h SceneCollectConditionHandler) {
	_, ok := sceneCollectConditionHandlerMap[tag]
	if ok {
		panic(fmt.Errorf("scene_collect_condition:repeat register %s", tag.String()))
	}
	sceneCollectConditionHandlerMap[tag] = h
}

func IsSceneCollectCondition(pl scene.Player, cn scene.CollectNPC, tag scenetypes.SceneType) bool {
	h, ok := sceneCollectConditionHandlerMap[tag]
	if !ok {
		return defaultCollectCondition
	}

	return h.SceneCollectCondition(pl, cn)
}
