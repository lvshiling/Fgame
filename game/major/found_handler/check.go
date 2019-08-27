package found_handler

import (
	"fgame/fgame/game/found/found"
	foundtypes "fgame/fgame/game/found/types"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	found.RegistFoundCheckHandler(foundtypes.FoundResourceTypeMajorShuangXiu, found.FoundCheckHandlerFunc(checkFoundBack))
	found.RegistFoundCheckHandler(foundtypes.FoundResourceTypeMajorFuQi, found.FoundCheckHandlerFunc(checkFoundBack))
}

func checkFoundBack(pl player.Player) bool {
	marryManager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	if marryManager.IsTrueMarry() {
		return true
	} else {
		return false
	}
}
