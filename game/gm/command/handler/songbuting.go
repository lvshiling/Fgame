package handler

import (
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/songbuting/pbutil"
	playersongbuting "fgame/fgame/game/songbuting/player"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeSongBuTing, command.CommandHandlerFunc(handleSongBuTingClear))

}

func handleSongBuTingClear(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)

	manager := pl.GetPlayerDataManager(types.PlayerSongBuTingDataManagerType).(*playersongbuting.PlayerSongBuTingDataManager)
	manager.GmClearNum()
	songBuTingObj := manager.GetSongBuTingObj()

	scSongBuTingChanged := pbutil.BuildSCSongBuTingChanged(songBuTingObj)
	err = pl.SendMsg(scSongBuTingChanged)
	return
}
