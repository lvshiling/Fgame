package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/game/scene/scene"
)

func BuildISPlayerBossReliveSync(pl scene.Player, reliveData *scene.PlayerBossReliveData) *crosspb.ISPlayerBossReliveSync {
	isPlayerBossReliveSync := &crosspb.ISPlayerBossReliveSync{}
	playerBossReliveData := &crosspb.PlayerBossReliveData{}
	bossType := int32(reliveData.GetBossType())
	reliveTime := reliveData.GetReliveTime()
	playerBossReliveData.BossType = &bossType
	playerBossReliveData.ReliveTime = &reliveTime
	isPlayerBossReliveSync.PlayerBossReliveData = playerBossReliveData
	return isPlayerBossReliveSync
}
