package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/teamcopy/pbutil"
	playerteamcopy "fgame/fgame/game/teamcopy/player"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTeamClearNum, command.CommandHandlerFunc(handleTeamClearNum))

}

func handleTeamClearNum(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)

	manager := pl.GetPlayerDataManager(types.PlayerTeamCopyDataManagerType).(*playerteamcopy.PlayerTeamCopyDataManager)
	manager.GmClearNum()
	teamCopyMap := manager.GetTeamCopyMap()
	scTeamCopyAllGet := pbutil.BuildSCTeamCopyAllGet(teamCopyMap)
	pl.SendMsg(scTeamCopyAllGet)
	return
}
