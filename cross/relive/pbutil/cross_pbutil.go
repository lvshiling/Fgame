package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/game/scene/scene"
)

func BuildISPlayerRelive(reliveTime int32) *crosspb.ISPlayerRelive {
	isPlayerRelive := &crosspb.ISPlayerRelive{}
	isPlayerRelive.ReliveTime = &reliveTime
	return isPlayerRelive
}

func BuildISPlayerReliveSync(pl scene.Player) *crosspb.ISPlayerReliveSync {
	isPlayerReliveSync := &crosspb.ISPlayerReliveSync{}
	playerReliveData := &crosspb.PlayerReliveData{}
	culTime := pl.GetCulReliveTime()
	playerReliveData.CulTime = &culTime
	lastReliveTime := pl.GetLastReliveTime()
	playerReliveData.LastReliveTime = &lastReliveTime
	isPlayerReliveSync.PlayerReliveData = playerReliveData
	return isPlayerReliveSync
}
