package foe

import (
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

type FoeNoticeHandler interface {
	FoeInfoNotice(pl, foePl scene.Player, sceneType scenetypes.SceneType) (err error)
}

type FoeNoticeHandlerFunc func(pl, foePl scene.Player, sceneType scenetypes.SceneType) (err error)

func (f FoeNoticeHandlerFunc) FoeInfoNotice(pl, foePl scene.Player, sceneType scenetypes.SceneType) (err error) {
	return f(pl, foePl, sceneType)
}

var (
	foeNoticeHandlerMap = make(map[scenetypes.SceneType]FoeNoticeHandler)
	defaultHandler      FoeNoticeHandler
)

func RegisterFoeNoticeHandler(sceneType scenetypes.SceneType, h FoeNoticeHandler) {
	_, ok := foeNoticeHandlerMap[sceneType]
	if ok {
		panic(fmt.Errorf("仇人信息推送处理器重复注册，场景类型：%s", sceneType.String()))
	}

	foeNoticeHandlerMap[sceneType] = h
}

func RegisterDefaultHandler(h FoeNoticeHandler) {
	defaultHandler = h
}

func FoeInfoNotice(pl, foePl scene.Player, sceneType scenetypes.SceneType) (err error) {
	h, ok := foeNoticeHandlerMap[sceneType]
	if !ok {
		return defaultHandler.FoeInfoNotice(pl, foePl, sceneType)
	}

	return h.FoeInfoNotice(pl, foePl, sceneType)
}
