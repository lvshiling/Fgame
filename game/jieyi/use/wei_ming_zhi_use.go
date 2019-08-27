package use

import (
	"fgame/fgame/common/lang"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/jieyi/jieyi"
	jieyilogic "fgame/fgame/game/jieyi/logic"
	"fgame/fgame/game/jieyi/pbutil"
	playerjieyi "fgame/fgame/game/jieyi/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeShengWei, playerinventory.ItemUseHandleFunc(handleResource))
}

func handleResource(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	if num <= 0 {
		return
	}

	plObj := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
	if plObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi: 玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}

	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	jieYiManager.AddShengWeiZhi(num)

	// 推送给结义兄弟
	jieYi := plObj.GetJieYi()
	scJieBrotherInfoOnChange := pbutil.BuildSCJieBrotherInfoOnChange(plObj)
	jieyilogic.BroadcastJieYi(jieYi, scJieBrotherInfoOnChange)

	flag = true
	return
}
