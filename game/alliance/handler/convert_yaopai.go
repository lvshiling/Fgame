package handler

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/processor"
	"fmt"

	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_YAO_PAI_CONVERT_TYPE), dispatch.HandlerFunc(handleAllianceConvert))
}

//处理腰牌兑换
func handleAllianceConvert(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理腰牌兑换")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = allianceConvert(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:处理腰牌兑换,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理腰牌兑换,完成")
	return nil

}

//腰牌兑换
func allianceConvert(pl player.Player) (err error) {
	allianceManager := pl.GetPlayerDataManager(types.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//物品是否足够
	flag := allianceManager.HasEnoughYaoPai()
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:处理腰牌兑换,腰牌数量不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//是否有兑换次数
	flag = allianceManager.HasEnoughConvetTiems()
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:处理腰牌兑换,兑换次数不足")
		playerlogic.SendSystemMessage(pl, lang.AllianceConvertTimesNotEnough)
		return
	}

	// 背包是否足够
	itemId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeConvertToQianKunBaDaiId)
	num := allianceManager.GetCanConvertTimes()
	if !inventoryManager.HasEnoughSlot(itemId, num) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:处理腰牌兑换,背包不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//扣除腰牌
	convertOneNeedNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeConvertNeedNum)
	needYaoPai := num * convertOneNeedNum
	yaopaiReason := commonlog.YaoPaiLogReasonExchange
	reasonText := fmt.Sprintf(yaopaiReason.String(), needYaoPai)
	flag = allianceManager.CostYaoPai(needYaoPai, yaopaiReason, reasonText)
	if !flag {
		panic("alliance:腰牌兑换应该成功")
	}

	//添加兑换物品
	inventoryReason := commonlog.InventoryLogReasonAllianceConvert
	flag = inventoryManager.AddItem(itemId, num, inventoryReason, inventoryReason.String())
	if !flag {
		panic("alliance:兑换物品，物品添加应该成功")
	}

	//更新兑换次数
	allianceManager.UpdateConvertTimes(num)

	//同步物品
	inventorylogic.SnapInventoryChanged(pl)

	scYaoPaiConvert := pbutil.BuildSCYaoPaiConvert(itemId, num)
	pl.SendMsg(scYaoPaiConvert)

	return
}
