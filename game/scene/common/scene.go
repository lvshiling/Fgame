package common

import (
	coretypes "fgame/fgame/core/types"
)

type SceneObject interface {
	GetMapId() int32
	GetSceneId() int64
	GetPos() coretypes.Position
	GetLastMapId() int32
	GetLastSceneId() int64
	GetLastPos() coretypes.Position
}

type sceneObject struct {
	mapId       int32
	sceneId     int64
	pos         coretypes.Position
	lastMapId   int32
	lastSceneId int64
	lastPos     coretypes.Position
}

func (o *sceneObject) GetMapId() int32 {
	return o.mapId
}

func (o *sceneObject) GetSceneId() int64 {
	return o.sceneId
}

func (o *sceneObject) GetPos() coretypes.Position {
	return o.pos
}

func (o *sceneObject) GetLastMapId() int32 {
	return o.lastMapId
}

func (o *sceneObject) GetLastSceneId() int64 {
	return o.lastSceneId
}

func (o *sceneObject) GetLastPos() coretypes.Position {
	return o.lastPos
}

func CreateSceneObject() SceneObject {
	o := &sceneObject{}
	return o
}
