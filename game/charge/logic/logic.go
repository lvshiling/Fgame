package logic

import (
	"context"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/common/message"
	"fgame/fgame/game/charge/charge"
	"fgame/fgame/game/charge/pbutil"
	playercharge "fgame/fgame/game/charge/player"
	chargetemplate "fgame/fgame/game/charge/template"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

func OnPlayerCharge(pl player.Player, orderId string, chargeId int32) {
	chargeTemplate := chargetemplate.GetChargeTemplateService().GetChargeTemplate(chargeId)
	if chargeTemplate == nil {
		return
	}

	chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

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

	// 生成充值订单
	flag := chargeManager.AddCharge(orderId, chargeId)
	if !flag {
		panic(fmt.Errorf("charge:玩家充值应该成功"))
	}

	chargeGold := int64(chargeTemplate.Gold)
	money := int64(chargeTemplate.Rmb)
	reason := commonlog.GoldLogReasonCharge
	propertyManager.AddGold(chargeGold, false, reason, reason.String())

	//同步
	propertylogic.SnapChangedProperty(pl)

	pl.AddChargeInfo(chargeGold, money)
	charge.GetChargeService().FinishOrder(pl, orderId)

	scCharge := pbutil.BuildSCCharge(chargeId)
	pl.SendMsg(scCharge)

	scMsg := pbutil.BuildSCChargeGoldNotice(chargeGold)
	pl.SendMsg(scMsg)

	reocrd := chargeManager.GetNewFirstChargeRecordInfo(startTime).GetRecord()
	scNewFirstChargeRecordNotice := pbutil.BuildSCNewFirstChargeRecordNotice(reocrd)
	pl.SendMsg(scNewFirstChargeRecordNotice)
}

// 后台充值
func OnPlayerPrivilegeCharge(pl player.Player) {
	chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	unfinishPrivilegeChargeList := charge.GetChargeService().GetUnfinishPrivilegeChargeList(pl)
	if len(unfinishPrivilegeChargeList) < 1 {
		return
	}

	for _, privilegeCharge := range unfinishPrivilegeChargeList {
		// 添加元宝
		gold := privilegeCharge.GetGoldNum()
		reason := commonlog.GoldLogReasonPrivilegeCharge
		propertyManager.AddGold(gold, false, reason, reason.String())
		// 更新充值状态
		charge.GetChargeService().FinishPriviCharge(pl, privilegeCharge.GetPrivilegeCharge())
		pl.AddPrivilegeChargeInfo(gold)

		chargeManager.AddPrivilegeCharge(int32(gold))

		scMsg := pbutil.BuildSCChargeGoldNotice(gold)
		pl.SendMsg(scMsg)
	}

	// 同步
	propertylogic.SnapChangedProperty(pl)

}

// func UsePrivilegeItem(pl player.Player, ) {
// 	chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
// 	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
// 	// 添加元宝
// 	gold := privilegeCharge.GetGoldNum()
// 	reason := commonlog.GoldLogReasonPrivilegeItem
// 	propertyManager.AddGold(gold, false, reason, reason.String())
// 	// 更新充值状态
// 	charge.GetChargeService().FinishPriviCharge(pl, privilegeCharge.GetPrivilegeCharge())
// 	pl.AddPrivilegeChargeInfo(gold)

// 	chargeManager.AddPrivilegeCharge(int32(gold))

// 	scMsg := pbutil.BuildSCChargeGoldNotice(gold)
// 	pl.SendMsg(scMsg)
// }

func BroadcastFirstCharge() {
	for _, p := range player.GetOnlinePlayerManager().GetAllPlayers() {
		ctx := scene.WithPlayer(context.Background(), p)
		msg := message.NewScheduleMessage(onFirstChargeNotice, ctx, nil, nil)
		p.Post(msg)
	}
}

func onFirstChargeNotice(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	spl, ok := pl.(player.Player)
	if !ok {
		return nil
	}

	// 推送新首充活动信息
	chargeManager := spl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
	start, duration := charge.GetChargeService().GetNewFirstChargeTime()
	info := chargeManager.GetNewFirstChargeRecordInfo(start)
	reocrd := info.GetRecord()
	scNewFirstChargeRecord := pbutil.BuildSCNewFirstChargeRecord(start, duration, reocrd)
	spl.SendMsg(scNewFirstChargeRecord)

	return nil
}
