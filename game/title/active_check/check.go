package check

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playertitle "fgame/fgame/game/title/player"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	"fgame/fgame/game/wardrobe/wardrobe"

	log "github.com/Sirupsen/logrus"
)

func init() {
	wardrobe.RegisterCheck(wardrobetypes.WardrobeSysTypeTitle, wardrobe.CheckHandlerFunc(handleTitle))
}

func handleTitle(pl player.Player, seqId int32) (flag bool) {
	log.Debug("title:处理判断是否激活")
	manager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	return manager.IfTitleExist(seqId)
}
