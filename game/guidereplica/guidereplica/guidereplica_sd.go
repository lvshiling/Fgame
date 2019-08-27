package guidereplica

import (
	guidereplicatypes "fgame/fgame/game/guidereplica/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

type GuideReplicaSdHandler interface {
	GuideReplicaSd(pl player.Player, temp *gametemplate.GuideReplicaTemplate, questId int32) (sh scene.SceneDelegate)
}

type GuideReplicaSdHandlerFunc func(pl player.Player, temp *gametemplate.GuideReplicaTemplate, questId int32) (sh scene.SceneDelegate)

func (f GuideReplicaSdHandlerFunc) GuideReplicaSd(pl player.Player, temp *gametemplate.GuideReplicaTemplate, questId int32) (sh scene.SceneDelegate) {
	return f(pl, temp, questId)
}

var (
	guideReplicaSdHandlerMap = make(map[guidereplicatypes.GuideReplicaType]GuideReplicaSdHandler)
)

func RegisterGuideReplicaSdHandler(tag guidereplicatypes.GuideReplicaType, h GuideReplicaSdHandler) {
	_, ok := guideReplicaSdHandlerMap[tag]
	if ok {
		panic(fmt.Errorf("guidereplica:repeat register %s", tag.String()))
	}
	guideReplicaSdHandlerMap[tag] = h
}

func GetGuideReplicaSd(pl player.Player, temp *gametemplate.GuideReplicaTemplate, questId int32) (sh scene.SceneDelegate) {
	if temp == nil {
		return
	}

	h, ok := guideReplicaSdHandlerMap[temp.GetGuideType()]
	if !ok {
		return
	}

	return h.GuideReplicaSd(pl, temp, questId)
}
