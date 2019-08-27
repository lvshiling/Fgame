package scene

import (
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

type ReliveHandler interface {
	Relive(Player, bool)
}

type ReliveHandlerFunc func(pl Player, autoBuy bool)

func (f ReliveHandlerFunc) Relive(pl Player, autoBuy bool) {
	f(pl, autoBuy)
}

var (
	reliveHandlerMap = make(map[scenetypes.SceneType]ReliveHandler)
)

func RegisterReliveHandler(sceneType scenetypes.SceneType, h ReliveHandler) {
	_, ok := reliveHandlerMap[sceneType]
	if ok {
		panic(fmt.Errorf("relive:repeat register %s relive", sceneType.String()))
	}
	reliveHandlerMap[sceneType] = h
}

func GetReliveHandler(sceneType scenetypes.SceneType) ReliveHandler {
	h, ok := reliveHandlerMap[sceneType]
	if !ok {
		return nil
	}
	return h
}

//回进入点复活
type ReliveEntryPointHandler interface {
	ReliveEntryPoint(Player) bool
}

type ReliveEntryPointHandlerFunc func(pl Player) bool

func (f ReliveEntryPointHandlerFunc) ReliveEntryPoint(pl Player) bool {
	return f(pl)
}

var (
	reliveEntryPointHandlerMap = make(map[scenetypes.SceneType]ReliveEntryPointHandler)
)

func RegisterReliveEntryPointHandler(sceneType scenetypes.SceneType, h ReliveEntryPointHandler) {
	_, ok := reliveEntryPointHandlerMap[sceneType]
	if ok {
		panic(fmt.Errorf("relive:repeat register %s relive", sceneType.String()))
	}
	reliveEntryPointHandlerMap[sceneType] = h
}

func GetReliveEntryPointHandler(sceneType scenetypes.SceneType) ReliveEntryPointHandler {
	h, ok := reliveEntryPointHandlerMap[sceneType]
	if !ok {
		return nil
	}
	return h
}

//回复活点复活
type RelivePointHandler interface {
	RelivePoint(Player) bool
}

type RelivePointHandlerFunc func(pl Player) bool

func (f RelivePointHandlerFunc) RelivePoint(pl Player) bool {
	return f(pl)
}

var (
	relivePointHandlerMap = make(map[scenetypes.SceneType]RelivePointHandler)
)

func RegisterRelivePointHandler(sceneType scenetypes.SceneType, h RelivePointHandler) {
	_, ok := relivePointHandlerMap[sceneType]
	if ok {
		panic(fmt.Errorf("relive:repeat register %s relive", sceneType.String()))
	}
	relivePointHandlerMap[sceneType] = h
}

func GetRelivePointHandler(sceneType scenetypes.SceneType) RelivePointHandler {
	h, ok := relivePointHandlerMap[sceneType]
	if !ok {
		return nil
	}
	return h
}

//自动复活

type AutoReliveHandler interface {
	AutoRelive(Player) bool
}

type AutoReliveHandlerFunc func(pl Player) bool

func (f AutoReliveHandlerFunc) AutoRelive(pl Player) bool {
	return f(pl)
}

var (
	autoReliveHandlerMap = make(map[scenetypes.SceneType]AutoReliveHandler)
)

func RegisterAutoReliveHandler(sceneType scenetypes.SceneType, h AutoReliveHandler) {
	_, ok := reliveHandlerMap[sceneType]
	if ok {
		panic(fmt.Errorf("relive:repeat register %s relive", sceneType.String()))
	}
	autoReliveHandlerMap[sceneType] = h
}

func GetAutoReliveHandler(sceneType scenetypes.SceneType) AutoReliveHandler {
	h, ok := autoReliveHandlerMap[sceneType]
	if !ok {
		return nil
	}
	return h
}
