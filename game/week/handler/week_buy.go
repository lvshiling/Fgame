package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/global"
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
	"fgame/fgame/game/week/pbutil"
	playerweek "fgame/fgame/game/week/player"
	weektemplate "fgame/fgame/game/week/template"
	weektypes "fgame/fgame/game/week/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WEEK_BUY_TYPE), dispatch.HandlerFunc(handlerBuyWeek))
}

const (
	initDay = 1
)

//处理购买周卡
func handlerBuyWeek(s session.Session, msg interface{}) (err error) {
	log.Debug("week:处理购买周卡请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSWeekBuy)
	typ := csMsg.GetWeekType()

	weekType := weektypes.WeekType(typ)
	if !weekType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weekType": weekType,
			}).Warn("week:购买周卡请求，类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = buyWeek(tpl, weekType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("week:处理购买周卡请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("week:处理购买Week请求完成")

	return
}

//购买周卡请求逻辑
func buyWeek(pl player.Player, weekType weektypes.WeekType) (err error) {
	weekTemp := weektemplate.GetWeekTemplateService().GetWeekTemplate(weekType)
	if weekTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weekType": weekType,
			}).Warn("week:购买周卡请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	weekManager := pl.GetPlayerDataManager(playertypes.PlayerWeekDataManagerType).(*playerweek.PlayerWeekManager)
	weekInfo := weekManager.GetWeekInfo(weekType)
	if weekInfo == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weekType": weekType,
			}).Warn("week:领取周卡奖励请求，周卡未初始化")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//是否已经是周卡
	now := global.GetGame().GetTimeService().Now()
	if weekInfo.IsWeek(now) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weekType": weekType,
			}).Warn("week:购买周卡请求，已经是周卡")
		playerlogic.SendSystemMessage(pl, lang.WeekHadBuyWeek)
		return
	}

	//元宝是否足够
	needGold := weekTemp.NeedGold
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(int64(needGold), false) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": needGold,
			}).Warn("week:购买周卡请求，当前元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//每日奖励模板
	rewDayInt := weekInfo.GetNextCycleDay()
	weekDayTemp := weekTemp.GetCycDayRew(rewDayInt)
	if weekDayTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"weekType":  weekType,
				"rewDayInt": rewDayInt,
			}).Warn("week:购买周卡请求，每日奖励模板")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//背包空间
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	totalItemMap := make(map[int32]int32)
	if !weekInfo.IsReceiveRewards(now) {
		if weekTemp.IsExtralRew(rewDayInt) {
			totalItemMap = coreutils.MergeMap(totalItemMap, weekTemp.GetExtralRewItemMap())
		}

		dayRewItemMap := weekDayTemp.GetRewItemMap()
		totalItemMap = coreutils.MergeMap(totalItemMap, dayRewItemMap)
	}
	firstRewItemMap := weekTemp.GetRewFirstItemMap()
	totalItemMap = coreutils.MergeMap(totalItemMap, firstRewItemMap)
	if !inventoryManager.HasEnoughSlots(totalItemMap) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"totalItemMap": totalItemMap,
			}).Warn("week:购买周卡请求，背包不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗元宝
	goldReason := commonlog.GoldLogReasonBuyWeekCost
	goldReasonText := fmt.Sprintf(goldReason.String(), weekType.String())
	flag := propertyManager.CostGold(int64(needGold), false, goldReason, goldReasonText)
	if !flag {
		panic(fmt.Errorf("week: buy week use gold should be ok"))
	}

	//添加物品
	itemReason := commonlog.InventoryLogReasonWeedBuyRew
	itemReasonText := fmt.Sprintf(itemReason.String(), weekType.String())
	flag = inventoryManager.BatchAdd(totalItemMap, itemReason, itemReasonText)
	if !flag {
		panic(fmt.Errorf("week: buy week add item should be ok"))
	}

	addBindGod := int64(weekTemp.RewBindGold + weekDayTemp.RewBindGold)
	addGold := int64(weekTemp.RewGold + weekDayTemp.RewGold)
	addSilver := int64(weekTemp.RewSilver + weekDayTemp.RewSilver)
	addGoldReason := commonlog.GoldLogReasonBuyWeekRew
	addGoldReasonText := fmt.Sprintf(addGoldReason.String(), weekType.String())
	silverReason := commonlog.SilverLogReasonBuyWeekRew
	silverReasonText := fmt.Sprintf(silverReason.String(), weekType.String())
	flag = propertyManager.AddMoney(addBindGod, addGold, addGoldReason, addGoldReasonText, addSilver, silverReason, silverReasonText)
	if !flag {
		panic(fmt.Errorf("week:添加银两元宝应该成功"))
	}

	//更新
	flag = weekManager.BuyWeek(weekType)
	if !flag {
		panic(fmt.Errorf("week:购买周卡失败"))
	}

	if !weekInfo.IsReceiveRewards(now) {
		flag = weekManager.ReceiveWeekRewards(weekType)
		if !flag {
			panic(fmt.Errorf("week:领取周卡奖励失败"))
		}
	}

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	rd := propertytypes.CreateRewData(0, 0, int32(addSilver), int32(addGold), int32(addBindGod))
	expire := weekInfo.GetExpireTime()
	scBuyWeek := pbutil.BuildSCWeekBuy(weekType, expire, totalItemMap, rd, weekInfo)
	pl.SendMsg(scBuyWeek)
	return
}
