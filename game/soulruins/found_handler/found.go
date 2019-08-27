package found_handler

import (
	"fgame/fgame/game/found/found"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playersoulruins "fgame/fgame/game/soulruins/player"
	"fgame/fgame/game/soulruins/soulruins"
)

func init() {
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeDiHunYiJi, found.FoundObjDataHandlerFunc(getDiHunYiJiFoundParam))
}

func getDiHunYiJiFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	group = int32(1)

	diHunManager := pl.GetPlayerDataManager(playertypes.PlayerSoulRuinsDataManagerType).(*playersoulruins.PlayerSoulRuinsDataManager)
	chapter, soulTyp, level := diHunManager.GetCurMaxLevel()
	resLevel = int32(soulruins.GetSoulRuinsService().GetSoulRuinsTemplate(chapter, soulTyp, level).TemplateId())
	maxTimes = soulruins.GetSoulRuinsService().GetSoulRuinsChallengeNum()
	return
}
