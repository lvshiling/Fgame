package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	wingeventtypes "fgame/fgame/game/wing/event/types"
	winglogic "fgame/fgame/game/wing/logic"
	"fgame/fgame/game/wing/pbutil"
	playerwing "fgame/fgame/game/wing/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeWingActivateCard, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handleWingActivateCard))
}

// 战翼激活卡
func handleWingActivateCard(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	//参数不对
	if num != 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("wing:使用战翼激活卡,使用物品数量错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	// 是否已经激活
	wingManager := pl.GetPlayerDataManager(playertypes.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	info := wingManager.GetWingInfo()
	if info.AdvanceId > 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("wing:使用战翼激活卡错误,已激活")
		playerlogic.SendSystemMessage(pl, lang.WingHadActivate)
		return
	}
	beforeNum := int32(info.AdvanceId)
	wingManager.WingActivate()

	wingReason := commonlog.WingLogReasonAdvanced
	reasonText := fmt.Sprintf(wingReason.String(), commontypes.AdvancedTypeActivateCard.String())
	data := wingeventtypes.CreatePlayerWingAdvancedLogEventData(beforeNum, 1, wingReason, reasonText)
	gameevent.Emit(wingeventtypes.EventTypeWingAdvancedLog, pl, data)

	//同步属性
	winglogic.WingPropertyChanged(pl)

	scWingAdvanced := pbutil.BuildSCWingAdavanced(int32(info.AdvanceId), info.WingId, info.Bless, info.Bless, info.BlessTime, commontypes.AdvancedTypeActivateCard, false)
	pl.SendMsg(scWingAdvanced)

	flag = true
	return
}
