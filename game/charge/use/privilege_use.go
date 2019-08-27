package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/charge/charge"
	"fgame/fgame/game/charge/pbutil"
	playercharge "fgame/fgame/game/charge/player"
	chargetemplate "fgame/fgame/game/charge/template"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeFuChi, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handleFuChi))
}

// 扶持卡
func handleFuChi(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	//参数不对
	itemId := it.ItemId
	if num != 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("wing:使用扶持卡,使用物品数量错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	itemTemp := item.GetItemService().GetItem(int(itemId))
	typ := pl.GetSDKType()
	subType := itemTemp.TypeFlag1
	chargeTemplate := chargetemplate.GetChargeTemplateService().GetChargeTemplateByType(typ, subType)
	if chargeTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("wing:使用扶持卡,充值模板不存在")
		playerlogic.SendSystemMessage(pl, lang.ChargePrivilegeFailed)
		return
	}

	chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	chargeId := int32(chargeTemplate.TemplateId())
	// // 档次首次充值赠送
	// isHadRecord := chargeManager.IsHadFirstChargeRecord(chargeId)
	// if !isHadRecord {
	// 	giftGodlReason := commonlog.GoldLogReasonFirstLevelCharge
	// 	reasonText := fmt.Sprintf(giftGodlReason.String(), chargeId)
	// 	giftGold := int64(chargeTemplate.FanhuanGold)
	// 	giftBindGold := int64(chargeTemplate.FanhuanBindGold)
	// 	if giftGold > 0 {
	// 		propertyManager.AddGold(giftGold, false, giftGodlReason, reasonText)
	// 	}
	// 	if giftBindGold > 0 {
	// 		propertyManager.AddGold(giftBindGold, true, giftGodlReason, reasonText)
	// 	}
	// }

	// // 特权充值
	// goldBackRate := pl.GetTianShuRate(tianshutypes.TianShuTypeGold)
	// if goldBackRate > 0 {
	// 	extraGold := int32(math.Ceil(float64(chargeGold) * float64(goldBackRate) / float64(common.MAX_RATE)))
	// 	title := lang.GetLangService().ReadLang(lang.EmailTianShuGoldFeedbackTitle)
	// 	content := lang.GetLangService().ReadLang(lang.EmailTianShuGoldFeedbackContent)
	// 	attachment := make(map[int32]int32)
	// 	attachment[int32(constanttypes.GoldItem)] = extraGold
	// 	emaillogic.AddEmail(pl, title, content, attachment)
	// }
	// bindGoldBackRate := pl.GetTianShuRate(tianshutypes.TianShuTypeBindGold)
	// if bindGoldBackRate > 0 {
	// 	extraBindGold := int32(math.Ceil(float64(chargeGold) * float64(bindGoldBackRate) / float64(common.MAX_RATE)))
	// 	title := lang.GetLangService().ReadLang(lang.EmailTianShuBindGoldFeedbackTitle)
	// 	content := lang.GetLangService().ReadLang(lang.EmailTianShuBindGoldFeedbackContent)
	// 	attachment := make(map[int32]int32)
	// 	attachment[int32(constanttypes.BindGoldItem)] = extraBindGold
	// 	emaillogic.AddEmail(pl, title, content, attachment)
	// }

	// 判断活动时间
	startTime, _ := charge.GetChargeService().GetNewFirstChargeTime()

	if charge.GetChargeService().IsNewFirstChargeDuration() {
		// 判断是否有没有记录信息
		if !chargeManager.IsHasNewFirstChargeRecord(chargeId, startTime) {
			flag := chargeManager.UpdateNewFirstChargeRecord(chargeId, startTime)
			if !flag {
				panic("charge: 更新新首充记录信息应该成功")
			}

			propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
			addGodlReason := commonlog.GoldLogReasonNewFirstChargeReturn
			addGodlReasonText := fmt.Sprintf(addGodlReason.String(), chargeTemplate.Gold)
			propertyManager.AddGold(int64(chargeTemplate.Gold), true, addGodlReason, addGodlReasonText)
		}
	}
	chargeGold := int64(chargeTemplate.Gold)
	chargeManager.AddPrivilegeChargeId(int32(chargeGold), chargeId)
	reason := commonlog.GoldLogReasonPrivilegeItem
	propertyManager.AddGold(chargeGold, false, reason, reason.String())

	propertylogic.SnapChangedProperty(pl)
	pl.AddPrivilegeChargeInfo(chargeGold)

	scCharge := pbutil.BuildSCCharge(chargeId)
	pl.SendMsg(scCharge)

	scMsg := pbutil.BuildSCChargeGoldNotice(chargeGold)
	pl.SendMsg(scMsg)
	flag = true
	return
}
