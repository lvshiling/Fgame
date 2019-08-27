package relive_handler

import (
	"fgame/fgame/common/lang"
	relivelogic "fgame/fgame/cross/relive/logic"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	worldbosspbutil "fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterReliveHandler(scenetypes.SceneTypeArenaShengShou, scene.ReliveHandlerFunc(relive))

}

//一般的正常复活
func relive(pl scene.Player, autoBuy bool) {

	s := pl.GetScene()
	if s == nil {
		return
	}
	bossType := worldbosstypes.BossTypeArena
	reliveTime := pl.GetBossReliveTime(bossType)

	maxReliveTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeShengShouReliveTimes)
	if reliveTime >= maxReliveTime {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"reliveTime":    reliveTime,
				"maxReliveTime": maxReliveTime,
			}).Warn("zhenxi:复活超过最大次数")
		playerlogic.SendSystemMessage(pl, lang.ReliveBaseExceedMaxTimes)
		return
	}

	relivelogic.Relive(pl)
	pl.PlayerBossRelive(bossType)
	reliveTime = pl.GetBossReliveTime(bossType)

	scWorldBossReliveTimeNotice := worldbosspbutil.BuildSCWorldBossReliveTimeNotice(int32(bossType), reliveTime)
	pl.SendMsg(scWorldBossReliveTimeNotice)

}
