package handler

import (
	"fgame/fgame/common/lang"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	houseextendedtypes "fgame/fgame/game/welfare/feedback/house_extended/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeHouseExtended, welfare.ReceiveHandlerFunc(receiveHouseExtended))
}

//房产活动领取奖励请求逻辑
func receiveHouseExtended(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeHouseExtended
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取房产活动奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	groupId := openTemp.Group

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"subType":  subType,
				"groupId":  groupId,
			}).Warn("welfare:运营活动,不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取房产活动奖励请求，活动不存在")
		return
	}
	info := obj.GetActivityData().(*houseextendedtypes.FeedbackHouseExtendedInfo)

	rewType := houseextendedtypes.HouseRewType(openTemp.Value1)
	needCharge := openTemp.Value2
	switch rewType {
	case houseextendedtypes.HouseRewTypeActivate:
		{
			if !info.IsCanReceiveActivateGift(needCharge) {
				log.WithFields(
					log.Fields{
						"playerId":          pl.GetId(),
						"groupId":           groupId,
						"needCharge":        needCharge,
						"activateChargeNum": info.ActivateChargeNum,
						"isActivate":        info.IsActivateGift,
					}).Warn("welfare:领取房产活动奖励请求，不满足领取条件")
				playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
				return
			}
		}
	case houseextendedtypes.HouseRewTypeUplevel:
		{
			rewLevel := openTemp.Value3
			if info.CurUplevelGiftLevel != rewLevel {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
						"rewId":    rewId,
						"rewLevel": rewLevel,
						"curLevel": info.CurUplevelGiftLevel,
					}).Warn("welfare:领取房产活动奖励请求，礼包等级错误")
				playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
				return
			}

			if !info.IsCanReceiveUplevelGift(needCharge) {
				log.WithFields(
					log.Fields{
						"playerId":         pl.GetId(),
						"groupId":          groupId,
						"needCharge":       needCharge,
						"UplevelChargeNum": info.UplevelChargeNum,
						"IsUplevelGift":    info.IsUplevelGift,
					}).Warn("welfare:领取房产活动奖励请求，不满足领取条件")
				playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
				return
			}
		}
	default:
		break
	}

	// //领取条件
	// groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	// if groupInterface == nil {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"groupId":  groupId,
	// 		}).Warn("welfare:领取房产活动奖励请求，模板不存在")
	// 	playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
	// 	return
	// }
	// groupTemp := groupInterface.(*feedbackhouseextendedtemplate.GroupTemplateHouseExtended)

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新信息
	switch rewType {
	case houseextendedtypes.HouseRewTypeActivate:
		{
			info.ReceiveActivateGift()
		}
	case houseextendedtypes.HouseRewTypeUplevel:
		{
			info.ReceiveUplevelGift()
		}
	default:
		break
	}
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityReceiveRew(rewId, groupId, totalRewData, rewItemMap, nil)
	pl.SendMsg(scMsg)
	return
}
