package scene

import (
	coretypes "fgame/fgame/core/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

// 场景进入检查
type CheckEnterSceneHandler interface {
	CheckEnter(pl Player, s Scene, pos coretypes.Position, enterType scenetypes.SceneEnterType) (flag bool, enterPos coretypes.Position)
}

type CheckEnterSceneHandlerFunc func(pl Player, s Scene, pos coretypes.Position, enterType scenetypes.SceneEnterType) (flag bool, enterPos coretypes.Position)

func (f CheckEnterSceneHandlerFunc) CheckEnter(pl Player, s Scene, pos coretypes.Position, enterType scenetypes.SceneEnterType) (flag bool, enterPos coretypes.Position) {
	return f(pl, s, pos, enterType)
}

func defaultCheck(enterType scenetypes.SceneEnterType) bool {
	switch enterType {
	case scenetypes.SceneEnterTypeTrac:
		return false
	default:
		return true
	}
}

var (
	checkEnterMap = make(map[scenetypes.SceneType]CheckEnterSceneHandler)
)

func RegisterCheckEnterHandler(mapType scenetypes.SceneType, h CheckEnterSceneHandler) {
	_, ok := checkEnterMap[mapType]
	if ok {
		panic(fmt.Errorf("场景进入检查器重复注册，地图类型：%s", mapType.String()))
	}

	checkEnterMap[mapType] = h
}

func CheckEnterScene(pl Player, s Scene, pos coretypes.Position, enterType scenetypes.SceneEnterType) (flag bool, enterPos coretypes.Position) {
	mapType := s.MapTemplate().GetMapType()
	h, ok := checkEnterMap[mapType]
	if !ok {
		return defaultCheck(enterType), pos
	}

	return h.CheckEnter(pl, s, pos, enterType)
}

// 场景位置修正检查
type CheckFixPosSceneHandler interface {
	CheckFixPos(pl Player, s Scene) bool
}

type CheckFixPosSceneHandlerFunc func(pl Player, s Scene) bool

func (f CheckFixPosSceneHandlerFunc) CheckFixPos(pl Player, s Scene) bool {
	return f(pl, s)
}

func defaultFixCheck() bool {
	return true
}

var (
	checkFixPosMap = make(map[scenetypes.SceneType]CheckFixPosSceneHandler)
)

func RegisterCheckFixPosHandler(mapType scenetypes.SceneType, h CheckFixPosSceneHandler) {
	_, ok := checkFixPosMap[mapType]
	if ok {
		panic(fmt.Errorf("场景位置修正检查器重复注册，地图类型：%s", mapType.String()))
	}

	checkFixPosMap[mapType] = h
}

func CheckFixPosScene(pl Player, s Scene) bool {
	mapType := s.MapTemplate().GetMapType()
	h, ok := checkFixPosMap[mapType]
	if !ok {
		return defaultFixCheck()
	}

	return h.CheckFixPos(pl, s)
}
