package listener

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/event"
	"fgame/fgame/core/utils"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/alliance/alliance"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	transportationlogic "fgame/fgame/game/transportation/logic"
	biaochenpc "fgame/fgame/game/transportation/npc/biaoche"
	"fgame/fgame/game/transportation/pbutil"
	transportationtemplate "fgame/fgame/game/transportation/template"
	"fgame/fgame/game/transportation/transpotation"
	transportationtypes "fgame/fgame/game/transportation/types"
	"fmt"
)

//劫镖完成后
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {

	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}
	if n.GetBiologyTemplate().GetBiologyScriptType() != scenetypes.BiologyScriptTypeBiaoChe {
		return
	}
	attackId := data.(int64)
	if attackId == 0 {
		return
	}

	biaoChe := n.(*biaochenpc.BiaocheNPC)
	transportationObj := biaoChe.GetTransportationObject()
	typ := transportationObj.GetTransportType()
	beRobPlayerId := transportationObj.GetPlayerId()
	beRobAllianceId := transportationObj.GetAllianceId()

	robName := ""
	robPl := player.GetOnlinePlayerManager().GetPlayerById(attackId)
	if robPl != nil {
		robName = robPl.GetName()
		robSuccess(robPl, beRobPlayerId, typ, beRobAllianceId)
	}
	if robName == "" {
		robName = scenetypes.GetKillName(attackId)
	}
	//劫镖成功
	transpotation.GetTransportService().TransportFail(beRobPlayerId, robName)

	beRobPl := player.GetOnlinePlayerManager().GetPlayerById(beRobPlayerId)
	if beRobPl != nil {
		scTransportBriefInfoNotice := pbutil.BuildSCTransportBriefInfoNotice(transportationObj, n)
		beRobPl.SendMsg(scTransportBriefInfoNotice)
	}
	return
}

func robSuccess(robPl player.Player, beRobPlayerId int64, typ transportationtypes.TransportationType, allianceId int64) {
	activityTemp := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeBiaoChe)
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	timeTemp, _ := activityTemp.GetActivityTimeTemplate(now, openTime, mergeTime)

	if typ == transportationtypes.TransportationTypeAlliance {
		beAllianceName := alliance.GetAllianceService().GetAlliance(allianceId).GetAllianceName()
		transportTemp := transportationtemplate.GetTransportationTemplateService().GetTransportationTemplate(typ)
		robRewMap := transportTemp.GetRobItemMap()
		beRobRewMap := transportTemp.GetLastItemMap()
		// 倍数奖励
		if timeTemp != nil {
			ratio := timeTemp.BeiShu
			robRewMap = utils.MultMap(robRewMap, ratio)
			beRobRewMap = utils.MultMap(beRobRewMap, ratio)
		}

		//劫仙盟镖成功，全体发邮件
		title := lang.GetLangService().ReadLang(lang.EmailAllianceTransportationRobName)
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailAllianceTransportationRobContent), beAllianceName)
		if robPl.GetAllianceId() == 0 {
			emaillogic.AddEmail(robPl, title, econtent, robRewMap)
		} else {
			transportationlogic.AllianceTransportRewBroadcast(robPl.GetAllianceId(), title, econtent, robRewMap)
		}

		//推送系统公告
		format := lang.GetLangService().ReadLang(lang.TransportationAllianceRobSystemBroadcast)
		content := fmt.Sprintf(format, robPl.GetName())
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))

		beRobPl := player.GetOnlinePlayerManager().GetPlayerById(beRobPlayerId)
		//被劫仙盟发邮件
		if beRobPl != nil {
			title := lang.GetLangService().ReadLang(lang.EmailAllianceTransportationBeRobTitle)
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailAllianceTransportationBeRobContent), robPl.GetName())
			emaillogic.AddEmail(beRobPl, title, econtent, beRobRewMap)
		} else {
			emaillogic.AddOfflineEmail(beRobPlayerId, title, econtent, beRobRewMap)
		}

		//移除镖车
		transpotation.GetTransportService().RemoveTransportation(beRobPlayerId)

	} else {
		//劫个人镖成功，直接获取奖励
		propertyManager := robPl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		tem := transportationtemplate.GetTransportationTemplateService().GetTransportationTemplate(typ)
		robNum := int64(0)
		rewItemMap := tem.GetRobItemMap()

		switch typ {
		case transportationtypes.TransportationTypeGold:
			{
				robNum = int64(tem.JiebiaoAwardGold)
				break
			}
		case transportationtypes.TransportationTypeSilver:
			{
				robNum = int64(tem.JiebiaoAwardSilver)
				break
			}
		}

		//倍数奖励
		if timeTemp != nil {
			ratio := timeTemp.BeiShu
			robNum = robNum * int64(ratio)
			rewItemMap = utils.MultMap(rewItemMap, ratio)
		}

		if robNum > 0 {
			switch typ {
			case transportationtypes.TransportationTypeGold:
				{
					goldReason := commonlog.GoldLogReasonTransportationRobRew
					goldReasonText := fmt.Sprintf(goldReason.String(), typ.String())
					propertyManager.AddGold(robNum, true, goldReason, goldReasonText)
					break
				}
			case transportationtypes.TransportationTypeSilver:
				{
					silverReason := commonlog.SilverLogReasonTransportationRobRew
					silverReasonText := fmt.Sprintf(silverReason.String(), typ.String())
					propertyManager.AddSilver(robNum, silverReason, silverReasonText)
					break
				}
			}
		}

		// if typ == transportationtypes.TransportationTypeGold {
		// 	robNum = int64(tem.JiebiaoAwardGold)
		// 	if robNum > 0 {
		// 		goldReason := commonlog.GoldLogReasonTransportationRobRew
		// 		goldReasonText := fmt.Sprintf(goldReason.String(), typ.String())
		// 		propertyManager.AddGold(robNum, true, goldReason, goldReasonText)
		// 	}
		// }

		// if typ == transportationtypes.TransportationTypeSilver {
		// 	robNum = int64(tem.JiebiaoAwardSilver)
		// 	if robNum > 0 {
		// 		silverReason := commonlog.SilverLogReasonTransportationRobRew
		// 		silverReasonText := fmt.Sprintf(silverReason.String(), typ.String())
		// 		propertyManager.AddSilver(int64(robNum), silverReason, silverReasonText)
		// 	}
		// }

		//获得物品
		if len(rewItemMap) > 0 {
			inventoryManager := robPl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
			itemGetReason := commonlog.InventoryLogReasonTransportRew
			reasonText := fmt.Sprintf(itemGetReason.String(), typ, transportationtypes.TransportStateTypeFail)
			flag := inventoryManager.BatchAdd(rewItemMap, itemGetReason, reasonText)
			if !flag {
				panic("transport: add item should be ok")
			}

			inventorylogic.SnapInventoryChanged(robPl)
		}

		propertylogic.SnapChangedProperty(robPl)

		//推送系统公告
		format := lang.GetLangService().ReadLang(lang.TransportationPersonalRobSystemBroadcast)
		content := fmt.Sprintf(format, robPl.GetName(), robNum, typ.String())
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))

		//劫镖成功消息
		scRobSuccessNotice := pbutil.BuildSCRobSuccessNotice(robNum, typ)
		robPl.SendMsg(scRobSuccessNotice)
	}
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
