package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	gamesession "fgame/fgame/game/session"
	investnewleveltypes "fgame/fgame/game/welfare/invest/new_level/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_INVEST_NEW_LEVEL_RECEIVE_REW_TYPE), dispatch.HandlerFunc(handlerReceiveInvestNewLevel))
}

//处理领取新等级投资计划
func handlerReceiveInvestNewLevel(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理新等级投资计划领取奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityInvestNewLevelReceiveRew := msg.(*uipb.CSOpenActivityInvestNewLevelReceiveRew)
	rewId := csOpenActivityInvestNewLevelReceiveRew.GetRewId()

	err = receiveInvestNewLevel(tpl, rewId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理新等级投资计划领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理新等级投资计划领取奖励请求完成")

	return
}

//新等级投资计划领取奖励请求逻辑
func receiveInvestNewLevel(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeInvest
	subType := welfaretypes.OpenActivityInvestSubTypeNewLevel

	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:新等级投资计划领取奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, openTemp.Group)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(openTemp.Group)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  openTemp.Group,
			}).Warn("welfare:新等级投资计划领取奖励请求，活动不存在")
		return
	}
	investLevelType := investnewleveltypes.InvestNewLevelType(openTemp.Value1)
	info := obj.GetActivityData().(*investnewleveltypes.InvestNewLevelInfo)
	receiveList, ok := info.InvestBuyInfoMap[investLevelType]
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"investLevelType": investLevelType,
			}).Warn("welfare:新等级投资计划领取奖励请求，未购买投资计划")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotBuyInvest)
		return
	}

	rewardLevel := openTemp.Value2
	if rewardLevel > pl.GetLevel() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"subType":     subType,
				"rewardLevel": rewardLevel,
			}).Warn("welfare:新等级投资计划领取奖励请求，等级不足")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	if !info.IsCanReceiveRew(investLevelType, rewardLevel) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"subType":  subType,
			}).Warn("welfare:新等级投资计划领取奖励请求，已领取奖励")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	showRewItemMap := make(map[int32]int32)
	rewItemMap := make(map[int32]int32)
	totalRewData := propertytypes.CreateRewData(0, 0, 0, 0, 0)
	//资源
	totalRewData.RewBindGold += openTemp.RewGoldBind
	totalRewData.RewGold += openTemp.RewGold
	totalRewData.RewSilver += openTemp.RewSilver
	//物品
	for itemId, num := range openTemp.GetRewItemMap() {
		_, ok := rewItemMap[itemId]
		if ok {
			rewItemMap[itemId] += num
		} else {
			rewItemMap[itemId] = num
		}
	}
	//前端展示
	for itemId, num := range openTemp.GetEmailRewItemMap() {
		_, ok := showRewItemMap[itemId]
		if ok {
			showRewItemMap[itemId] += num
		} else {
			showRewItemMap[itemId] = num
		}
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	newItemDataList := welfarelogic.ConvertToItemData(rewItemMap, inventorytypes.NewItemLimitTimeTypeNone, 0)
	//背包空间
	if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemDataList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:新等级投资计划领取奖励请求，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//增加物品
	itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
	itemGetReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)
	flag := inventoryManager.BatchAddOfItemLevel(newItemDataList, itemGetReason, itemGetReasonText)
	if !flag {
		panic("welfare:invest rewards add item should be ok")
	}

	reasonGold := commonlog.GoldLogReasonOpenActivityRew
	reasonSilver := commonlog.SilverLogReasonOpenActivityRew
	reasonLevel := commonlog.LevelLogReasonOpenActivityRew
	reasonGoldText := fmt.Sprintf(reasonGold.String(), typ, subType)
	reasonSilverText := fmt.Sprintf(reasonSilver.String(), typ, subType)
	reasonLevelText := fmt.Sprintf(reasonLevel.String(), typ, subType)
	flag = propertyManager.AddRewData(totalRewData, reasonGold, reasonGoldText, reasonSilver, reasonSilverText, reasonLevel, reasonLevelText)
	if !flag {
		panic("welfare:invest rewards add RewData should be ok")
	}

	//更新信息
	receiveList = append(receiveList, rewardLevel)
	info.InvestBuyInfoMap[investLevelType] = receiveList
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scOpenActivityInvestNewLevelReceiveRew := pbutil.BuildSCOpenActivityInvestNewLevelReceiveRew(totalRewData, showRewItemMap, rewId)
	pl.SendMsg(scOpenActivityInvestNewLevelReceiveRew)
	return
}
