package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	coreutils "fgame/fgame/core/utils"
	playercharge "fgame/fgame/game/charge/player"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeYuanBaoKa, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handleUseYuanbaoKa))
}

func handleUseYuanbaoKa(pl player.Player, it *playerinventory.PlayerItemObject, itemNum int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	goldNum := int64(itemTemplate.TypeFlag1)
	needChargeGoldNum := int64(itemTemplate.TypeFlag2)

	// 玩家本日充值元宝数量
	chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
	chargeGoldNum := chargeManager.GetTodayChargeNum()
	if needChargeGoldNum > chargeGoldNum {
		log.WithFields(
			log.Fields{
				"playerId":          pl.GetId(),
				"chargeGoldNum":     chargeGoldNum,
				"needChargeGoldNum": needChargeGoldNum,
			}).Warn("welfare:未满足使用元宝卡条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityYuanBaoKaPlayerChargeGoldNotEnough)
		return
	}

	// 使用元宝卡获得元宝
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	goldReason := commonlog.GoldLogReasonUseYuanBaoKa
	goldReasonText := fmt.Sprintf(goldReason.String(), itemId, itemNum)
	propertyManager.AddGold(goldNum, false, goldReason, goldReasonText)

	// 公告
	arg := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(funcopentypes.FuncOpenTypeDuanWuQiYuDao)}
	link := coreutils.FormatLink(chattypes.ButtonTypeToGet, arg)
	plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	acName := chatlogic.FormatMailKeyWordNoticeStr(itemTemplate.Name)
	goldText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", goldNum))
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityYuanBaoKaUseNotice), plName, acName, goldText, link)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	propertylogic.SnapChangedProperty(pl)

	flag = true
	return
}
