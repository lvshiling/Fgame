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
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	gamesession "fgame/fgame/game/session"
	investsevendaytemplate "fgame/fgame/game/welfare/invest/sevenday/template"
	investsevendaytypes "fgame/fgame/game/welfare/invest/sevenday/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"

	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_INVEST_DAY_RECEIVE_REW_TYPE), dispatch.HandlerFunc(handlerReceiveInvestDay))
}

//处理领取七日投资
func handlerReceiveInvestDay(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理七日投资领取奖励请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csOpenActivityInvestDayReceiveRew := msg.(*uipb.CSOpenActivityInvestDayReceiveRew)
	groupId := csOpenActivityInvestDayReceiveRew.GetGroupId()

	err = receiveInvestDay(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理七日投资领取奖励请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理七日投资领取奖励请求完成")

	return
}

//七日投资领取奖励请求逻辑
func receiveInvestDay(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeInvest
	subType := welfaretypes.OpenActivityInvestSubTypeServenDay

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:七日投资领取奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:七日投资领取奖励请求，活动不存在")
		return
	}
	info := obj.GetActivityData().(*investsevendaytypes.InvestDayInfo)
	maxExclude := info.ReceiveDay
	curDay := info.GetCurDay()

	if info.BuyTime <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:七日投资领取奖励请求，未购买七日投资")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotBuyInvest)
		return
	}
	groupTemp := groupInterface.(*investsevendaytemplate.GroupTemplateInvestDay)
	maxRewards := groupTemp.GetInvestDayMaxRewardsLevel()
	if maxExclude >= maxRewards {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:七日投资领取奖励请求，没有可领取奖励")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
		return
	}

	showRewItemMap := make(map[int32]int32)
	rewItemMap := make(map[int32]int32)
	totalRewData := propertytypes.CreateRewData(0, 0, 0, 0, 0)
	tempList := groupTemp.GetInvestDayRewTempList(maxExclude, curDay)
	for _, temp := range tempList {
		//资源
		totalRewData.RewBindGold += temp.RewGoldBind
		totalRewData.RewGold += temp.RewGold
		totalRewData.RewSilver += temp.RewSilver

		//物品
		for itemId, num := range temp.GetRewItemMap() {
			_, ok := rewItemMap[itemId]
			if ok {
				rewItemMap[itemId] += num
			} else {
				rewItemMap[itemId] = num
			}
		}

		//前端展示
		for itemId, num := range temp.GetEmailRewItemMap() {
			_, ok := showRewItemMap[itemId]
			if ok {
				showRewItemMap[itemId] += num
			} else {
				showRewItemMap[itemId] = num
			}
		}
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//背包空间
	firsOpenTemp := groupTemp.GetFirstOpenTemp()
	newItemDataList := welfarelogic.ConvertToItemData(rewItemMap, firsOpenTemp.GetExpireType(), firsOpenTemp.GetExpireTime())
	if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemDataList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:七日投资领取奖励请求，背包空间不足")
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
	info.ReceiveDay = curDay
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scOpenActivityInvestDayReceiveRew := pbutil.BuildSCOpenActivityInvestDayReceiveRew(totalRewData, showRewItemMap)
	pl.SendMsg(scOpenActivityInvestDayReceiveRew)
	return
}
