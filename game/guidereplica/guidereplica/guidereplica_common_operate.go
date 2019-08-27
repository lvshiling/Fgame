package guidereplica

import (
	guidereplicatypes "fgame/fgame/game/guidereplica/types"
	"fgame/fgame/game/player"
	"fmt"
)

type GuideCommonOperateHandler interface {
	GuideCommonOperate(pl player.Player) error
}

type GuideCommonOperateHandlerFunc func(pl player.Player) error

func (f GuideCommonOperateHandlerFunc) GuideCommonOperate(pl player.Player) error {
	return f(pl)
}

var (
	guideCommonOperateHandlerMap = make(map[guidereplicatypes.GuideReplicaType]GuideCommonOperateHandler)
)

func RegisterGuideCommonOperateHandler(tag guidereplicatypes.GuideReplicaType, h GuideCommonOperateHandler) {
	_, ok := guideCommonOperateHandlerMap[tag]
	if ok {
		panic(fmt.Errorf("guidereplica:repeat register %s", tag.String()))
	}
	guideCommonOperateHandlerMap[tag] = h
}

func GetGuideCommonOperate(pl player.Player, tag guidereplicatypes.GuideReplicaType) GuideCommonOperateHandler {
	h, ok := guideCommonOperateHandlerMap[tag]
	if !ok {
		return nil
	}

	return h
}
