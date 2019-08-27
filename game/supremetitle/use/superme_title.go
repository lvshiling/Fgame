package use

import (
	"fgame/fgame/common/lang"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"

	supremetitlelogic "fgame/fgame/game/supremetitle/logic"
	"fgame/fgame/game/supremetitle/pbutil"
	playersupremetitle "fgame/fgame/game/supremetitle/player"
	supremetitletemplate "fgame/fgame/game/supremetitle/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeTitle, itemtypes.ItemTitleSubTypeDingZhiCard, playerinventory.ItemUseHandleFunc(handleDingZhi))
}

func handleDingZhi(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	//参数不对
	itemId := it.ItemId
	if num != 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("supremetitle:使用定制称号卡")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	titleId := itemTemplate.TypeFlag1

	manager := pl.GetPlayerDataManager(types.PlayerSupremeTitleDataManagerType).(*playersupremetitle.PlayerSupremeTitleDataManager)
	titleTemplate := supremetitletemplate.GetTitleDingZhiTemplateService().GetTitleDingZhiTempalte(titleId)
	if titleTemplate == nil {
		return
	}
	existFlag := manager.IfTitleExist(titleId)
	if existFlag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("supremetitle:该至尊称号已激活,无需激活")
		playerlogic.SendSystemMessage(pl, lang.SupremeTitleRepeatActive)
		return
	}

	flag = manager.TitleActive(titleId)
	if !flag {
		panic(fmt.Errorf("supremetitle: supremeTitleActive should be ok"))
	}

	//同步属性
	supremetitlelogic.SupremeTitlePropertyChanged(pl)
	scTitleActive := pbutil.BuildSCSupremeTitleActive(titleId)
	pl.SendMsg(scTitleActive)

	return
}
