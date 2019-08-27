package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/liveness/pbutil"
	playerliveness "fgame/fgame/game/liveness/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questpbutil "fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/scene/scene"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeLivenessClear, command.CommandHandlerFunc(handleLivenessClear))
}

func handleLivenessClear(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)

	manager := pl.GetPlayerDataManager(types.PlayerLivenessDataManagerType).(*playerliveness.PlayerLivenessDataManager)
	manager.GmSetLivenessClearNum()

	livenessObj := manager.GetLiveness()
	livenessMap := manager.GetLivenessQuestMap()
	scLivenessGet := pbutil.BuildSCLivenessGet(livenessObj, livenessMap)
	p.SendMsg(scLivenessGet)

	questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	for questId, _ := range livenessMap {
		quest := questManager.CommitLivenessResetInit(questId)
		if quest != nil {
			questlogic.CheckInitQuest(pl, quest)
			scQuestUpdate := questpbutil.BuildSCQuestUpdate(quest)
			pl.SendMsg(scQuestUpdate)
		}
	}

	return
}
