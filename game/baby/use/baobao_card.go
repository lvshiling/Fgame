package item

import (
	"fgame/fgame/common/lang"
	playerbaby "fgame/fgame/game/baby/player"
	babytypes "fgame/fgame/game/baby/types"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBaoBaoCard, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handleBaoBaoCard))
}

//宝宝卡
func handleBaoBaoCard(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	if num <= 0 {
		return
	}

	baobaoPropertyData, ok := it.PropertyData.(*babytypes.BabyPropertyData)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"chooseIndexList": chooseIndexList,
			}).Warn("inventory:使用玩具,不是宝宝数据结构")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//宝宝数量
	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	if !babyManager.IsCanAddBaby() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("baby:处理洞房消息,当前宝宝个数已达上限，超生可提高宝宝数量")
		playerlogic.SendSystemMessage(pl, lang.BabyMaxNum)
		return
	}

	babyManager.AddBaby(baobaoPropertyData.Quality, baobaoPropertyData.Danbei, baobaoPropertyData.Sex, baobaoPropertyData.TalentList)

	flag = true
	return
}
