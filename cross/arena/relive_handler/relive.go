package relive_handler

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/lang"
	arenatemplate "fgame/fgame/game/arena/template"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterReliveHandler(scenetypes.SceneTypeArena, scene.ReliveHandlerFunc(Relive))

}

//一般的正常复活
func Relive(pl scene.Player, autoBuy bool) {
	reliveTime := pl.GetArenaReliveTime()
	maxReliveTime := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RebornAmountMax
	if reliveTime >= maxReliveTime {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"reliveTime":    reliveTime,
				"maxReliveTime": maxReliveTime,
			}).Warn("arena:复活超过最大次数")
		playerlogic.SendSystemMessage(pl, lang.ReliveBaseExceedMaxTimes)
		return
	}
	isArenaRelive := &crosspb.ISArenaRelive{}
	pl.SendMsg(isArenaRelive)
	//判断复活单

	// pl.ArenaRelive()
	// mi := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeItemNumAddEveryRelive)) / float64(common.MAX_RATE)
	// first := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFirstReliveItemNum))
	// reliveItemId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeReliveItemId)
	//TODO: 防止越界
	//判断消耗
	// itemNum := int32(math.Ceil(first * math.Pow(float64(reliveTime), mi)))
	//扣除复活丹

	// pl.Reborn(pl.GetPosition())

}
