package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droppbutil "fgame/fgame/game/drop/pbutil"
	droptemplate "fgame/fgame/game/drop/template"
	propertytypes "fgame/fgame/game/property/types"
	rankentity "fgame/fgame/game/rank/entity"
	scenepbutil "fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	advancedblessfeedbacktypes "fgame/fgame/game/welfare/advanced/bless_feedback/types"
	advancedfeedbacktypes "fgame/fgame/game/welfare/advanced/feedback/types"
	advancedrewpowertypes "fgame/fgame/game/welfare/advancedrew/power/types"
	advancedrewrewtypes "fgame/fgame/game/welfare/advancedrew/rew/types"
	advancedrewrewextendedtypes "fgame/fgame/game/welfare/advancedrew/rew_extended/types"
	cyclechargetypes "fgame/fgame/game/welfare/cycle/charge/types"
	cyclechargesingletypes "fgame/fgame/game/welfare/cycle/charge_single/types"
	discountbeachtypes "fgame/fgame/game/welfare/discount/beach/types"
	discountdiscounttypes "fgame/fgame/game/welfare/discount/discount/types"
	discountkanjiatypes "fgame/fgame/game/welfare/discount/kanjia/types"
	discountyunyintypes "fgame/fgame/game/welfare/discount/yunyin/types"
	discountzhuanshengtypes "fgame/fgame/game/welfare/discount/zhuansheng/types"
	drewchargedrewtypes "fgame/fgame/game/welfare/drew/charge_drew/types"
	drewcrazyboxtypes "fgame/fgame/game/welfare/drew/crazy_box/types"
	drewsmasheggtypes "fgame/fgame/game/welfare/drew/smash_egg/types"
	feedbackchargetypes "fgame/fgame/game/welfare/feedback/charge/types"
	feedbackchargecycletypes "fgame/fgame/game/welfare/feedback/charge_cycle/types"
	feedbackchargedeveloptypes "fgame/fgame/game/welfare/feedback/charge_develop/types"
	feedbackchargesingletypes "fgame/fgame/game/welfare/feedback/charge_single/types"
	feedbackcosttypes "fgame/fgame/game/welfare/feedback/cost/types"
	feedbackgoldbowltypes "fgame/fgame/game/welfare/feedback/gold_bowl/types"
	feedbackhouseinvesttypes "fgame/fgame/game/welfare/feedback/house_invest/types"
	feedbacklabatypes "fgame/fgame/game/welfare/feedback/laba/types"
	feedbackpigtypes "fgame/fgame/game/welfare/feedback/pig/types"
	grouptimesrewtypes "fgame/fgame/game/welfare/group/times_rew/types"
	halllogintypes "fgame/fgame/game/welfare/hall/login/types"
	hallonlinetypes "fgame/fgame/game/welfare/hall/online/types"
	hallrealmtypes "fgame/fgame/game/welfare/hall/realm/types"
	hallupleveltypes "fgame/fgame/game/welfare/hall/uplevel/types"
	investleveltypes "fgame/fgame/game/welfare/invest/level/types"
	investnewleveltypes "fgame/fgame/game/welfare/invest/new_level/types"
	investservendaytypes "fgame/fgame/game/welfare/invest/sevenday/types"
	madeexptypes "fgame/fgame/game/welfare/made/exp/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	rewardschargelimittypes "fgame/fgame/game/welfare/rewards/charge_limit/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func BuildSCOpenActivityFirstCharge(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityFirstCharge {
	scWelfareFirstCharge := &uipb.SCOpenActivityFirstCharge{}
	for itemId, num := range itemMap {
		scWelfareFirstCharge.DropInfo = append(scWelfareFirstCharge.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scWelfareFirstCharge.RewInfo = buildRewProperty(rd)

	return scWelfareFirstCharge
}

func BuildSCOpenActivityWelfareLoginReceiveRew(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityWelfareLoginReceiveRew {
	scWelfareReceiveRew := &uipb.SCOpenActivityWelfareLoginReceiveRew{}
	for itemId, num := range itemMap {
		scWelfareReceiveRew.DropInfo = append(scWelfareReceiveRew.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scWelfareReceiveRew.RewInfo = buildRewProperty(rd)

	return scWelfareReceiveRew
}

func BuildSCOpenActivityReceiveRew(rewId, groupId int32, rd *propertytypes.RewData, itemMap map[int32]int32, record []int32) *uipb.SCOpenActivityReceiveRew {
	scMsg := &uipb.SCOpenActivityReceiveRew{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.RewInfo = buildRewProperty(rd)
	scMsg.RewId = &rewId
	scMsg.GroupId = &groupId
	scMsg.RecordList = record

	return scMsg
}

func BuildSCOpenActivityReceiveRewMultiple(rewId, groupId int32, rd *propertytypes.RewData, itemDataList []*droptemplate.DropItemData, record []int32) *uipb.SCOpenActivityReceiveRewMultiple {
	scMsg := &uipb.SCOpenActivityReceiveRewMultiple{}
	scMsg.DropInfo = buildDropInfoList(itemDataList)
	scMsg.RewInfo = buildRewProperty(rd)
	scMsg.RewId = &rewId
	scMsg.GroupId = &groupId
	scMsg.RecordList = record

	return scMsg
}

func BuildSCOpenActivityRewPoolsDrew(groupId int32, position int32, itemDataList []*droptemplate.DropItemData, backTimes int32, logList []*welfare.DrewLogObject) *uipb.SCOpenActivityRewPoolsDrew {
	scMsg := &uipb.SCOpenActivityRewPoolsDrew{}
	scMsg.DropInfo = buildDropInfoList(itemDataList)
	scMsg.GroupId = &groupId
	scMsg.Position = &position
	scMsg.BackTimes = &backTimes
	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildDrewLog(log))
	}
	return scMsg
}

func BuildSCOpenActivityReceiveRewCycleSingleMaxRewMultiple(rewId, groupId int32, rd *propertytypes.RewData, itemMap map[int32]int32, record map[int32]int32) *uipb.SCOpenActivityReceiveRew {
	scMsg := &uipb.SCOpenActivityReceiveRew{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.RewInfo = buildRewProperty(rd)
	scMsg.RewId = &rewId
	scMsg.GroupId = &groupId
	scMsg.RecordInfo = buildTimesInfo(record)

	return scMsg
}

func BuildSCOpenActivityReceiveRewCycleSingleAllRewMultiple(rewId, groupId int32, rd *propertytypes.RewData, itemMap map[int32]int32, record map[int32]int32) *uipb.SCOpenActivityReceiveRew {
	scMsg := &uipb.SCOpenActivityReceiveRew{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.RewInfo = buildRewProperty(rd)
	scMsg.RewId = &rewId
	scMsg.GroupId = &groupId
	scMsg.RecordInfo = buildTimesInfo(record)

	return scMsg
}

func BuildSCOpenActivityWelfareUplevelReceiveRew(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityWelfareUplevelReceiveRew {
	scWelfareReceiveRew := &uipb.SCOpenActivityWelfareUplevelReceiveRew{}
	for itemId, num := range itemMap {
		scWelfareReceiveRew.DropInfo = append(scWelfareReceiveRew.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scWelfareReceiveRew.RewInfo = buildRewProperty(rd)

	return scWelfareReceiveRew
}

func BuildSCOpenActivityWelfareOnlineReceiveRew(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityWelfareOnlineReceiveRew {
	scWelfareReceiveRew := &uipb.SCOpenActivityWelfareOnlineReceiveRew{}
	for itemId, num := range itemMap {
		scWelfareReceiveRew.DropInfo = append(scWelfareReceiveRew.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scWelfareReceiveRew.RewInfo = buildRewProperty(rd)

	return scWelfareReceiveRew
}

func BuildSCOpenActivityWelfareOnlineDrew(itemId, num int32) *uipb.SCOpenActivityWelfareOnlineDrew {
	scOpenActivityWelfareOnlineDrew := &uipb.SCOpenActivityWelfareOnlineDrew{}
	scOpenActivityWelfareOnlineDrew.DropInfo = buildDropInfo(itemId, num, 0)

	return scOpenActivityWelfareOnlineDrew
}

func BuildSCOpenActivityInvestLevelBuy() *uipb.SCOpenActivityInvestLevelBuy {
	scOpenActivityInvestLevelBuy := &uipb.SCOpenActivityInvestLevelBuy{}
	return scOpenActivityInvestLevelBuy
}

func BuildSCOpenActivityInvestNewLevelBuy() *uipb.SCOpenActivityInvestNewLevelBuy {
	scOpenActivityInvestNewLevelBuy := &uipb.SCOpenActivityInvestNewLevelBuy{}
	return scOpenActivityInvestNewLevelBuy
}

func BuildSCOpenActivityInvestUpLevel() *uipb.SCOpenActivityInvestUpLevel {
	scOpenActivityInvestUpLevel := &uipb.SCOpenActivityInvestUpLevel{}
	return scOpenActivityInvestUpLevel
}

func BuildSCOpenActivityInvestDayBuy() *uipb.SCOpenActivityInvestDayBuy {
	scOpenActivityInvestBuy := &uipb.SCOpenActivityInvestDayBuy{}
	return scOpenActivityInvestBuy
}

func BuildSCOpenActivityNewInvestDayBuy(groupId, typ int32) *uipb.SCOpenActivityNewInvestDayBuy {
	scOpenActivityNewInvestDayBuy := &uipb.SCOpenActivityNewInvestDayBuy{}
	scOpenActivityNewInvestDayBuy.GroupId = &groupId
	scOpenActivityNewInvestDayBuy.Typ = &typ
	return scOpenActivityNewInvestDayBuy
}

func BuildSCMergeActivityDiscountBuy(groupId, discountId, buyNum, itemId, num int32, globalTimesList []*welfaretypes.TimesLimitInfo) *uipb.SCMergeActivityDiscountBuy {
	scMsg := &uipb.SCMergeActivityDiscountBuy{}
	scMsg.GroupId = &groupId
	scMsg.DiscountId = &discountId
	scMsg.Num = &buyNum
	scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	for _, data := range globalTimesList {
		itemIndex := data.Key
		globalNum := data.Times
		scMsg.GlobalTimesList = append(scMsg.GlobalTimesList, buildDiscountGlobalTimesInfo(itemIndex, globalNum))
	}
	return scMsg
}

func BuildSCOpenActivityKanJiaBuy(groupId, giftType int32, buyRecord []int32) *uipb.SCOpenActivityKanJiaBuy {
	scMsg := &uipb.SCOpenActivityKanJiaBuy{}
	scMsg.GroupId = &groupId
	scMsg.Type = &giftType
	scMsg.BuyRecordList = buyRecord
	return scMsg
}

func BuildSCOpenActivityKanJia(groupId int32, kanjiaRecord map[int32]*discountkanjiatypes.KanJiaInfo) *uipb.SCOpenActivityKanJia {
	scMsg := &uipb.SCOpenActivityKanJia{}
	scMsg.GroupId = &groupId
	scMsg.KanjiaInfoList = buildKanJiaInfoList(kanjiaRecord)

	return scMsg
}

func buildKanJiaInfoList(kanjiaRecord map[int32]*discountkanjiatypes.KanJiaInfo) (infoList []*uipb.KanJiaInfo) {
	for typ, kanJiaData := range kanjiaRecord {
		discount := kanJiaData.Discount
		times := kanJiaData.KanJiaTimes
		giftType := int32(typ)

		kanjiaInfo := &uipb.KanJiaInfo{}
		kanjiaInfo.Type = &giftType
		kanjiaInfo.DiscountNum = &discount
		kanjiaInfo.KanjiaTimes = &times
		infoList = append(infoList, kanjiaInfo)
	}
	return infoList
}

func BuildSCOpenActivityTaoCanBuy(groupId int32) *uipb.SCOpenActivityTaoCanBuy {
	scMsg := &uipb.SCOpenActivityTaoCanBuy{}
	scMsg.GroupId = &groupId
	return scMsg
}

func BuildSCOpenActivityMadeRes(groupId int32, useTimes int32) *uipb.SCOpenActivityMadeRes {
	scMsg := &uipb.SCOpenActivityMadeRes{}
	scMsg.GroupId = &groupId
	scMsg.MadeTimes = &useTimes
	return scMsg
}

func BuildSCOpenActivityDiscountZhuanShengBuy(itemMap map[int32]int32, giftType int32, groupId, num int32, usePoint int32) *uipb.SCOpenActivityDiscountZhuanShengBuy {
	scMsg := &uipb.SCOpenActivityDiscountZhuanShengBuy{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	typeInt := int32(giftType)
	scMsg.Typ = &typeInt
	scMsg.GroupId = &groupId
	scMsg.Num = &num
	scMsg.UsePoint = &usePoint
	return scMsg
}

func BuildSCOpenActivityDiscountZhuanShengBuyAll(groupId int32, itemMap map[int32]int32, buyGiftMap map[int32]int32) *uipb.SCOpenActivityDiscountZhuanShengBuyAll {
	scMsg := &uipb.SCOpenActivityDiscountZhuanShengBuyAll{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.GroupId = &groupId
	scMsg.BuyInfoList = buildTimesInfo(buyGiftMap)
	return scMsg
}

func BuildSCOpenActivityInvestLevelReceiveRew(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityInvestLevelReceiveRew {
	scOpenActivityInvestLevelReceiveRew := &uipb.SCOpenActivityInvestLevelReceiveRew{}
	for itemId, num := range itemMap {
		scOpenActivityInvestLevelReceiveRew.DropInfo = append(scOpenActivityInvestLevelReceiveRew.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scOpenActivityInvestLevelReceiveRew.RewInfo = buildRewProperty(rd)
	return scOpenActivityInvestLevelReceiveRew
}

func BuildSCOpenActivityInvestNewLevelReceiveRew(rd *propertytypes.RewData, itemMap map[int32]int32, rewId int32) *uipb.SCOpenActivityInvestNewLevelReceiveRew {
	scOpenActivityInvestNewLevelReceiveRew := &uipb.SCOpenActivityInvestNewLevelReceiveRew{}
	for itemId, num := range itemMap {
		scOpenActivityInvestNewLevelReceiveRew.DropInfo = append(scOpenActivityInvestNewLevelReceiveRew.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scOpenActivityInvestNewLevelReceiveRew.RewInfo = buildRewProperty(rd)
	scOpenActivityInvestNewLevelReceiveRew.RewId = &rewId
	return scOpenActivityInvestNewLevelReceiveRew
}

func BuildSCOpenActivityInvestDayReceiveRew(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityInvestDayReceiveRew {
	scWelfareReceiveRew := &uipb.SCOpenActivityInvestDayReceiveRew{}
	for itemId, num := range itemMap {
		scWelfareReceiveRew.DropInfo = append(scWelfareReceiveRew.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scWelfareReceiveRew.RewInfo = buildRewProperty(rd)

	return scWelfareReceiveRew
}

func BuildSCOpenActivityFeedbackChargeReceiveRew(rd *propertytypes.RewData, itemMap map[int32]int32, rewId int32) *uipb.SCOpenActivityFeedbackChargeReceiveRew {
	scMsg := &uipb.SCOpenActivityFeedbackChargeReceiveRew{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.RewInfo = buildRewProperty(rd)
	scMsg.RewId = &rewId

	return scMsg
}

func BuildSCOpenActivityDiscountZhuanShengReceiveGift(groupId, giftType int32, itemMap map[int32]int32) *uipb.SCOpenActivityDiscountZhuanShengReceiveGift {
	scMsg := &uipb.SCOpenActivityDiscountZhuanShengReceiveGift{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.GroupId = &groupId
	scMsg.Typ = &giftType

	return scMsg
}

func BuildSCOpenActivityFeedbackDevelopReceive(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityFeedbackDevelopReceive {
	scMsg := &uipb.SCOpenActivityFeedbackDevelopReceive{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.RewInfo = buildRewProperty(rd)

	return scMsg
}

func BuildSCOpenActivityFeedbackDevelopRevive(groupId, feedDay int32, isDead bool) *uipb.SCOpenActivityFeedbackDevelopRevive {
	scMsg := &uipb.SCOpenActivityFeedbackDevelopRevive{}
	scMsg.GroupId = &groupId
	scMsg.FeedDay = &feedDay
	scMsg.IsDead = &isDead

	return scMsg
}

func BuildSCOpenActivityFeedbackCostReceiveRew(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityFeedbackCostReceiveRew {
	scOpenActivityFeedbackReceiveRew := &uipb.SCOpenActivityFeedbackCostReceiveRew{}
	for itemId, num := range itemMap {
		scOpenActivityFeedbackReceiveRew.DropInfo = append(scOpenActivityFeedbackReceiveRew.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scOpenActivityFeedbackReceiveRew.RewInfo = buildRewProperty(rd)

	return scOpenActivityFeedbackReceiveRew
}

func BuildSCOpenActivityTimesRewReceive(rewId int32, itemMap map[int32]int32, record map[int32][]int32) *uipb.SCOpenActivityTimesRewReceive {
	scMsg := &uipb.SCOpenActivityTimesRewReceive{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.RewId = &rewId
	scMsg.RecordList = buildTimesRewRecordList(record)
	return scMsg
}

func BuildSCOpenActivityCycleChargeRewards(rewId int32, rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityCycleChargeRewards {
	scMsg := &uipb.SCOpenActivityCycleChargeRewards{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.RewInfo = buildRewProperty(rd)
	scMsg.RewId = &rewId

	return scMsg
}

func BuildSCOpenActivityCycleSingleChargeRewards(rewId int32, rewRecordList []int32, rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityCycleSingleChargeRewards {
	scMsg := &uipb.SCOpenActivityCycleSingleChargeRewards{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.RewInfo = buildRewProperty(rd)
	scMsg.RewId = &rewId
	scMsg.RewRecordList = rewRecordList

	return scMsg
}

func BuildSCMergeActivityAdvancedRewards(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCMergeActivityAdvancedRewards {
	scMergeActivityAdvancedRewards := &uipb.SCMergeActivityAdvancedRewards{}
	for itemId, num := range itemMap {
		scMergeActivityAdvancedRewards.DropInfo = append(scMergeActivityAdvancedRewards.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMergeActivityAdvancedRewards.RewInfo = buildRewProperty(rd)

	return scMergeActivityAdvancedRewards
}

func BuildSCMergeActivityAdvancedBlessRewards(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCMergeActivityAdvancedBlessRewards {
	scMergeActivityAdvancedBlessRewards := &uipb.SCMergeActivityAdvancedBlessRewards{}
	for itemId, num := range itemMap {
		scMergeActivityAdvancedBlessRewards.DropInfo = append(scMergeActivityAdvancedBlessRewards.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMergeActivityAdvancedBlessRewards.RewInfo = buildRewProperty(rd)

	return scMergeActivityAdvancedBlessRewards
}

func BuildSCOpenActivityAdvancedReceiveRew(rd *propertytypes.RewData, itemMap map[int32]int32, rewId int32, record []int32) *uipb.SCOpenActivityAdvancedReceiveRew {
	scMsg := &uipb.SCOpenActivityAdvancedReceiveRew{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.RewInfo = buildRewProperty(rd)
	scMsg.RewId = &rewId
	scMsg.RewRecordList = record

	return scMsg
}

func BuildSCOpenActivityAdvancedPowerReceiveRew(rd *propertytypes.RewData, itemMap map[int32]int32, rewId int32, record []int64) *uipb.SCOpenActivityAdvancedPowerReceiveRew {
	scMsg := &uipb.SCOpenActivityAdvancedPowerReceiveRew{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.RewInfo = buildRewProperty(rd)
	scMsg.RewId = &rewId
	scMsg.RewRecordList = record

	return scMsg
}

func BuildSCOpenActivityAdvancedRewExtendedReceive(rd *propertytypes.RewData, itemMap map[int32]int32, rewId int32, record []int32) *uipb.SCOpenActivityAdvancedRewExtendedReceive {
	scMsg := &uipb.SCOpenActivityAdvancedRewExtendedReceive{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMsg.RewInfo = buildRewProperty(rd)
	scMsg.RewId = &rewId
	scMsg.RewRecordList = record

	return scMsg
}

func BuildSCOpenActivityRealmRewards(rd *propertytypes.RewData, timesList []*welfaretypes.TimesLimitInfo, itemMap map[int32]int32) *uipb.SCOpenActivityRealmRewards {
	scMsg := &uipb.SCOpenActivityRealmRewards{}
	for itemId, num := range itemMap {
		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, 0))
	}
	for _, data := range timesList {
		level := data.Key
		times := data.Times
		scMsg.TimesInfo = append(scMsg.TimesInfo, buildRewardsTimesInfo(level, times))
	}
	scMsg.RewInfo = buildRewProperty(rd)

	return scMsg
}

func BuildSCMergeActivitySingleChargeRewards(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCMergeActivitySingleChargeRewards {
	scMergeActivitySingleChargeRewards := &uipb.SCMergeActivitySingleChargeRewards{}
	for itemId, num := range itemMap {
		scMergeActivitySingleChargeRewards.DropInfo = append(scMergeActivitySingleChargeRewards.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scMergeActivitySingleChargeRewards.RewInfo = buildRewProperty(rd)

	return scMergeActivitySingleChargeRewards
}

func BuildSCOpenActivityFeedbackCycleChargeRewards(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityFeedbackCycleChargeRewards {
	scOpenActivityFeedbackCycleChargeRewards := &uipb.SCOpenActivityFeedbackCycleChargeRewards{}
	for itemId, num := range itemMap {
		scOpenActivityFeedbackCycleChargeRewards.DropInfo = append(scOpenActivityFeedbackCycleChargeRewards.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scOpenActivityFeedbackCycleChargeRewards.RewInfo = buildRewProperty(rd)

	return scOpenActivityFeedbackCycleChargeRewards
}

func BuildSCOpenActivityFeedbackHouseInvestAdvenced(groupId int32, rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityFeedbackHouseInvestAdvenced {
	scOpenActivityFeedbackHouseInvestAdvenced := &uipb.SCOpenActivityFeedbackHouseInvestAdvenced{}

	scOpenActivityFeedbackHouseInvestAdvenced.GroupId = &groupId
	for itemId, num := range itemMap {
		scOpenActivityFeedbackHouseInvestAdvenced.DropInfo = append(scOpenActivityFeedbackHouseInvestAdvenced.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scOpenActivityFeedbackHouseInvestAdvenced.RewInfo = buildRewProperty(rd)

	return scOpenActivityFeedbackHouseInvestAdvenced
}

func BuildSCOpenActivityFeedbackHouseInvestDecor(groupId int32, rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityFeedbackHouseInvestDecor {
	scOpenActivityFeedbackHouseInvestDecor := &uipb.SCOpenActivityFeedbackHouseInvestDecor{}

	scOpenActivityFeedbackHouseInvestDecor.GroupId = &groupId
	for itemId, num := range itemMap {
		scOpenActivityFeedbackHouseInvestDecor.DropInfo = append(scOpenActivityFeedbackHouseInvestDecor.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scOpenActivityFeedbackHouseInvestDecor.RewInfo = buildRewProperty(rd)

	return scOpenActivityFeedbackHouseInvestDecor
}

func BuildSCOpenActivityFeedbackHouseInvestSell(groupId int32, rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityFeedbackHouseInvestSell {
	scOpenActivityFeedbackHouseInvestSell := &uipb.SCOpenActivityFeedbackHouseInvestSell{}

	scOpenActivityFeedbackHouseInvestSell.GroupId = &groupId
	for itemId, num := range itemMap {
		scOpenActivityFeedbackHouseInvestSell.DropInfo = append(scOpenActivityFeedbackHouseInvestSell.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scOpenActivityFeedbackHouseInvestSell.RewInfo = buildRewProperty(rd)

	return scOpenActivityFeedbackHouseInvestSell
}

func BuildSCOpenActivityFeedbackGoldLaBaAttend(groupId int32, logList []*welfare.GoldLaBaLogObject, rewGold, labaTimes int32) *uipb.SCOpenActivityFeedbackGoldLaBaAttend {
	scMsg := &uipb.SCOpenActivityFeedbackGoldLaBaAttend{}
	scMsg.RewGold = &rewGold
	scMsg.GroupId = &groupId
	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildLaBaLog(log))
	}
	scMsg.LabaTimes = &labaTimes
	return scMsg
}

func BuildSCOpenActivityDevelopFeed(groupId int32, favorableNum, dayFavorableNum int32, feedTimesMap map[int32]int32, itemDataList []*droptemplate.DropItemData) *uipb.SCOpenActivityDevelopFeed {
	scMsg := &uipb.SCOpenActivityDevelopFeed{}
	scMsg.GroupId = &groupId
	scMsg.CurFavorableNum = &favorableNum
	scMsg.TimesInfoList = buildTimesInfo(feedTimesMap)
	scMsg.DropInfo = buildDropInfoList(itemDataList)
	scMsg.DayFavorableNum = &dayFavorableNum
	return scMsg
}

func BuildSCOpenActivityLaBaLogIncr(logList []*welfare.GoldLaBaLogObject, groupId int32) *uipb.SCOpenActivityLaBaLogIncr {
	scMsg := &uipb.SCOpenActivityLaBaLogIncr{}
	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildLaBaLog(log))
	}
	scMsg.GroupId = &groupId
	return scMsg
}

func BuildSCOpenActivityDrewLogIncr(logList []*welfare.DrewLogObject, groupId int32) *uipb.SCOpenActivityDrewLogIncr {
	scMsg := &uipb.SCOpenActivityDrewLogIncr{}
	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildDrewLog(log))
	}
	scMsg.GroupId = &groupId
	return scMsg
}

func BuildSCOpenActivityCrazyBoxLogIncr(logList []*welfare.CrazyBoxLogObject, groupId int32) *uipb.SCOpenActivityCrazyBoxLogIncr {
	scMsg := &uipb.SCOpenActivityCrazyBoxLogIncr{}
	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildCrazyBoxLog(log))
	}
	scMsg.GroupId = &groupId
	return scMsg
}

func BuildSCMergeActivityLuckDrewAttends(itemDataList []*droptemplate.DropItemData, extraItemDataList []*droptemplate.DropItemData, groupId int32) *uipb.SCMergeActivityLuckDrewAttend {
	scMergeActivityLuckDrewAttend := &uipb.SCMergeActivityLuckDrewAttend{}

	scMergeActivityLuckDrewAttend.DropList = buildDropInfoList(itemDataList)
	scMergeActivityLuckDrewAttend.ExtraDropList = buildDropInfoList(extraItemDataList)

	scMergeActivityLuckDrewAttend.GroupId = &groupId

	return scMergeActivityLuckDrewAttend
}

func BuildSCOpenActivityChargeDrewAttend(itemDataList []*droptemplate.DropItemData, extraItemDataList []*droptemplate.DropItemData, logList []*welfare.DrewLogObject, groupId, nextRate, drewType int32, rewIndexList []int32) *uipb.SCOpenActivityChargeDrewAttend {
	scMsg := &uipb.SCOpenActivityChargeDrewAttend{}
	scMsg.Rate = &nextRate
	scMsg.GroupId = &groupId
	scMsg.DrewType = &drewType
	scMsg.RewIndexList = rewIndexList

	scMsg.DropList = buildDropInfoList(itemDataList)
	scMsg.ExtraDropList = buildDropInfoList(extraItemDataList)
	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildDrewLog(log))
	}

	return scMsg
}

func BuildSCOpenActivityChargeDrewAttendWithDrewPokerInfo(scMsg *uipb.SCOpenActivityChargeDrewAttend, pokerList, rewPokerTypeList []int32) {
	info := &uipb.DrewPokerInfo{}
	info.PokerList = pokerList
	info.RewPokerTypeList = rewPokerTypeList
	scMsg.PokerInfo = info
	return
}

func BuildSCOpenActivityCostDrewAttend(itemDataList []*droptemplate.DropItemData, extraItemDataList []*droptemplate.DropItemData, logList []*welfare.DrewLogObject, groupId, nextRate, drewType, dropLevel int32, rewIndexList []int32) *uipb.SCOpenActivityCostDrewAttend {
	scMsg := &uipb.SCOpenActivityCostDrewAttend{}
	scMsg.Rate = &nextRate
	scMsg.GroupId = &groupId
	scMsg.DrewType = &drewType
	scMsg.DropLevel = &dropLevel
	scMsg.RewIndexList = rewIndexList
	scMsg.DropList = buildDropInfoList(itemDataList)
	scMsg.ExtraDropList = buildDropInfoList(extraItemDataList)

	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildDrewLog(log))
	}

	return scMsg
}

func BuildSCOpenActivityCrazyBoxAttend(itemDataList []*droptemplate.DropItemData, extraItemDataList []*droptemplate.DropItemData, logList []*welfare.CrazyBoxLogObject, groupId, drewType int32) *uipb.SCOpenActivityCrazyBoxAttend {
	scMsg := &uipb.SCOpenActivityCrazyBoxAttend{}
	scMsg.GroupId = &groupId
	scMsg.DrewType = &drewType
	scMsg.DropList = buildDropInfoList(itemDataList)
	scMsg.ExtraDropList = buildDropInfoList(extraItemDataList)

	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildCrazyBoxLog(log))
	}

	return scMsg
}

func BuildSCOpenActivitySmashEggAttend(itemDataList []*droptemplate.DropItemData, extraItemDataList []*droptemplate.DropItemData, groupId, attendCount, dropNum int32) *uipb.SCOpenActivitySmashEggAttend {
	scMsg := &uipb.SCOpenActivitySmashEggAttend{}
	scMsg.GroupId = &groupId
	scMsg.AttendCount = &attendCount
	scMsg.DropNum = &dropNum
	scMsg.DropInfo = buildDropInfoList(itemDataList)
	scMsg.ExtraDropList = buildDropInfoList(itemDataList)
	return scMsg
}

func BuildSCMergeActivityLuckBombOreAttend(itemDataList []*droptemplate.DropItemData, extraItemDataList []*droptemplate.DropItemData) *uipb.SCMergeActivityLuckBombOreAttend {
	scMergeActivityLuckBombOreAttend := &uipb.SCMergeActivityLuckBombOreAttend{}
	scMergeActivityLuckBombOreAttend.DropList = buildDropInfoList(itemDataList)
	scMergeActivityLuckBombOreAttend.ExtraDropList = buildDropInfoList(extraItemDataList)
	return scMergeActivityLuckBombOreAttend
}

func BuildSCOpenActivityGiftCode(rd *propertytypes.RewData, itemMap map[int32]int32) *uipb.SCOpenActivityGiftCode {
	scOpenActivityGiftCode := &uipb.SCOpenActivityGiftCode{}
	for itemId, num := range itemMap {
		scOpenActivityGiftCode.DropInfo = append(scOpenActivityGiftCode.DropInfo, buildDropInfo(itemId, num, 0))
	}
	scOpenActivityGiftCode.RewInfo = buildRewProperty(rd)

	return scOpenActivityGiftCode
}

func BuildSCOpenActivityFeedbackCostNotice(groupId int32, curCostGold int32) *uipb.SCOpenActivityFeedbackCostNotice {
	scOpenActivityFeedbackCostNotice := &uipb.SCOpenActivityFeedbackCostNotice{}
	scOpenActivityFeedbackCostNotice.GroupId = &groupId
	scOpenActivityFeedbackCostNotice.CurCostGold = &curCostGold
	return scOpenActivityFeedbackCostNotice
}

func BuildSCOpenActivityFeedbackChargeNotice(groupId int32, curChargeGold int64) *uipb.SCOpenActivityFeedbackChargeNotice {
	scMsg := &uipb.SCOpenActivityFeedbackChargeNotice{}
	scMsg.CurChargeGold = &curChargeGold
	scMsg.GroupId = &groupId
	return scMsg
}

func BuildSCOpenActivityRankChargeNotice(curChargeGold int32, groupId int32) *uipb.SCOpenActivityRankChargeNotice {
	scOpenActivityRankChargeNotice := &uipb.SCOpenActivityRankChargeNotice{}
	scOpenActivityRankChargeNotice.CurChargeGold = &curChargeGold
	scOpenActivityRankChargeNotice.GroupId = &groupId
	return scOpenActivityRankChargeNotice
}

func BuildSCOpenActivityDrewTimesNotice(groupId, drewTimes int32) *uipb.SCOpenActivityChargeDrewTimesNotice {
	scMsg := &uipb.SCOpenActivityChargeDrewTimesNotice{}
	scMsg.DrewTimes = &drewTimes
	scMsg.GroupId = &groupId
	return scMsg
}

func BuildSCOpenActivityCrazyBoxTimesNotice(groupId, drewTimes int32) *uipb.SCOpenActivityCrazyBoxTimesNotice {
	scMsg := &uipb.SCOpenActivityCrazyBoxTimesNotice{}
	scMsg.DrewTimes = &drewTimes
	scMsg.GroupId = &groupId
	return scMsg
}

func BuildSCOpenActivityRankCostNotice(curCostGold int64, groupId int32) *uipb.SCOpenActivityRankCostNotice {
	scOpenActivityRankCostNotice := &uipb.SCOpenActivityRankCostNotice{}
	costGold := int32(curCostGold)
	scOpenActivityRankCostNotice.CurCostGold = &costGold
	scOpenActivityRankCostNotice.GroupId = &groupId
	return scOpenActivityRankCostNotice
}

func BuildSCOpenActivityCostRewardsCostInfo(groupId int32, costGold int64) *uipb.SCOpenActivityCostRewardsCostInfo {
	scMsg := &uipb.SCOpenActivityCostRewardsCostInfo{}
	scMsg.GroupId = &groupId
	scMsg.CurCostGold = &costGold
	return scMsg
}

func BuildSCOpenActivityCycleChargeNotice(curChargeGold int32, groupId int32) *uipb.SCOpenActivityCycleChargeNotice {
	scOpenActivityCycleChargeNotice := &uipb.SCOpenActivityCycleChargeNotice{}
	scOpenActivityCycleChargeNotice.CurChargeGold = &curChargeGold
	scOpenActivityCycleChargeNotice.GroupId = &groupId
	return scOpenActivityCycleChargeNotice
}

func BuildSCOpenActivityEndBroadcast(groupId int32) *uipb.SCOpenActivityEndBroadcast {
	scOpenActivityEndBroadcast := &uipb.SCOpenActivityEndBroadcast{}
	scOpenActivityEndBroadcast.GroupId = &groupId
	return scOpenActivityEndBroadcast
}

func BuildSCMergeActivitySingleChargeNotice(groupId, maxSingleNum int32) *uipb.SCMergeActivitySingleChargeNotice {
	scMsg := &uipb.SCMergeActivitySingleChargeNotice{}
	scMsg.MaxSingleChargeNum = &maxSingleNum
	scMsg.GroupId = &groupId
	return scMsg
}

func BuildSCOpenActivitySystemActivate(groupId int32) *uipb.SCOpenActivitySystemActivate {
	scMsg := &uipb.SCOpenActivitySystemActivate{}
	scMsg.GroupId = &groupId
	return scMsg
}

func BuildSCOpenActivityFeedbackCycleChargeNotice(curChargeNum int32) *uipb.SCOpenActivityFeedbackCycleChargeNotice {
	scOpenActivityFeedbackCycleChargeNotice := &uipb.SCOpenActivityFeedbackCycleChargeNotice{}
	scOpenActivityFeedbackCycleChargeNotice.CurChargeNum = &curChargeNum
	return scOpenActivityFeedbackCycleChargeNotice
}

func BuildSCOpenActivityFeedbackChargeReturnMultipleNotice(periodChargeNum int32) *uipb.SCOpenActivityFeedbackChargeReturnMultipleNotice {
	scOpenActivityFeedbackChargeReturnMultipleNotice := &uipb.SCOpenActivityFeedbackChargeReturnMultipleNotice{}
	scOpenActivityFeedbackChargeReturnMultipleNotice.PeriodChargeNum = &periodChargeNum
	return scOpenActivityFeedbackChargeReturnMultipleNotice
}

func BuildSCOpenActivityPeriodChargeNotice(groupId, periodChargeNum int32) *uipb.SCOpenActivityPeriodChargeNotice {
	scOpenActivityPeriodChargeNotice := &uipb.SCOpenActivityPeriodChargeNotice{}
	scOpenActivityPeriodChargeNotice.GroupId = &groupId
	scOpenActivityPeriodChargeNotice.PeriodChargeNum = &periodChargeNum
	return scOpenActivityPeriodChargeNotice
}

func BuildSCOpenActivityFeedbackHouseInvestNotice(groupId int32, chargeNum int32, curDayChargeNum int32) *uipb.SCOpenActivityFeedbackHouseInvestNotice {
	scOpenActivityFeedbackHouseInvestNotice := &uipb.SCOpenActivityFeedbackHouseInvestNotice{}
	scOpenActivityFeedbackHouseInvestNotice.GroupId = &groupId
	scOpenActivityFeedbackHouseInvestNotice.ChargeNum = &chargeNum
	scOpenActivityFeedbackHouseInvestNotice.CurDayChargeNum = &curDayChargeNum
	return scOpenActivityFeedbackHouseInvestNotice
}

func BuildSCOpenActivityLoginInfo(dayNum int32, loginObj *playerwelfare.PlayerOpenActivityObject) *uipb.SCOpenActivityLoginInfo {
	scOpenActivityLoginInfo := &uipb.SCOpenActivityLoginInfo{}
	scOpenActivityLoginInfo.WelfareLoginInfo = buildWelfareLoginInfo(dayNum, loginObj)

	return scOpenActivityLoginInfo
}

func BuildSCOpenActivityUplevelInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId, level int32, startTime, endTime int64) *uipb.SCOpenActivityUplevelInfo {
	scMsg := &uipb.SCOpenActivityUplevelInfo{}
	scMsg.WelfareUplevelInfo = buildWelfareUplevelInfo(obj, groupId, level, startTime, endTime)

	return scMsg
}

func BuildSCOpenActivityOnlineInfo(onlineTime int64, onlineObj *playerwelfare.PlayerOpenActivityObject) *uipb.SCOpenActivityOnlineInfo {
	scOpenActivityOnlineInfo := &uipb.SCOpenActivityOnlineInfo{}
	scOpenActivityOnlineInfo.WelfareOnlineInfo = buildWelfareOnlineInfo(onlineTime, onlineObj)

	return scOpenActivityOnlineInfo
}

func BuildSCOpenActivityInvestLevelInfo(obj *playerwelfare.PlayerOpenActivityObject) *uipb.SCOpenActivityInvestLevelInfo {
	scOpenActivityInvestLevelInfo := &uipb.SCOpenActivityInvestLevelInfo{}
	if obj != nil {
		info := obj.GetActivityData().(*investleveltypes.InvestLevelInfo)
		for typ, record := range info.InvestBuyInfoMap {
			groupId := obj.GetGroupId()
			scOpenActivityInvestLevelInfo.InvestLevelInfo = append(scOpenActivityInvestLevelInfo.InvestLevelInfo, buildInvestLevelInfo(record, typ, groupId))
			break
		}
	}

	return scOpenActivityInvestLevelInfo
}

func BuildSCOpenActivityInvestNewLevelInfo(obj *playerwelfare.PlayerOpenActivityObject) *uipb.SCOpenActivityInvestNewLevelInfo {
	scOpenActivityInvestNewLevelInfo := &uipb.SCOpenActivityInvestNewLevelInfo{}
	if obj != nil {
		info := obj.GetActivityData().(*investnewleveltypes.InvestNewLevelInfo)
		for typ, records := range info.InvestBuyInfoMap {
			groupId := obj.GetGroupId()
			scOpenActivityInvestNewLevelInfo.InvestLevelInfo = append(scOpenActivityInvestNewLevelInfo.InvestLevelInfo, buildInvestNewLevelInfo(records, typ, groupId))
			break
		}
	}

	return scOpenActivityInvestNewLevelInfo
}

func BuildSCOpenActivityInvestDayInfo(investDayObj *playerwelfare.PlayerOpenActivityObject) *uipb.SCOpenActivityInvestDayInfo {
	scOpenActivityInvestDayInfo := &uipb.SCOpenActivityInvestDayInfo{}
	scOpenActivityInvestDayInfo.InvestDayInfo = buildInvestDayInfo(investDayObj)

	return scOpenActivityInvestDayInfo
}

func BuildSCOpenActivityFeedbackChargeInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityFeedbackChargeInfo {
	scOpenActivityFeedbackChargeInfo := &uipb.SCOpenActivityFeedbackChargeInfo{}

	var record []int32
	chargeNum := int32(0)
	if obj != nil {
		charge := obj.GetActivityData().(*feedbackchargetypes.FeedbackChargeInfo)
		record = charge.RewRecord
		chargeNum = charge.GoldNum

	}

	info := &uipb.FeedbackChargeInfo{}
	info.ChargeNum = &chargeNum
	info.ChargeRecordList = record
	info.GroupId = &groupId
	info.StartTime = &startTime
	info.EndTime = &endTime
	scOpenActivityFeedbackChargeInfo.FeedbackChargeInfo = info
	return scOpenActivityFeedbackChargeInfo
}

func BuildSCOpenActivityFeedbackDevelopInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityFeedbackDevelopInfo {
	scMsg := &uipb.SCOpenActivityFeedbackDevelopInfo{}

	isActivate := false
	isFeed := false
	isDead := false
	isReceiveRew := false
	chargeNum := int32(0)
	feedDay := int32(0)
	todayCostNum := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*feedbackchargedeveloptypes.FeedbackDevelopInfo)
		chargeNum = int32(info.ActivateChargeNum)
		isActivate = info.IsActivate
		isFeed = info.IsFeed
		isDead = info.IsDead
		isReceiveRew = info.IsReceiveRew
		feedDay = info.FeedTimes
		todayCostNum = int32(info.TodayCostNum)
	}

	scMsg.GroupId = &groupId
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	scMsg.IsFeed = &isFeed
	scMsg.IsDead = &isDead
	scMsg.IsActivate = &isActivate
	scMsg.ChargeNum = &chargeNum
	scMsg.FeedDay = &feedDay
	scMsg.TodayCostNum = &todayCostNum
	scMsg.IsReceiveRew = &isReceiveRew

	return scMsg
}

func BuildSCOpenActivityMadeResInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityMadeResInfo {
	scMsg := &uipb.SCOpenActivityMadeResInfo{}

	madeTimes := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*madeexptypes.MadeInfo)
		madeTimes = info.Times
	}

	scMsg.GroupId = &groupId
	scMsg.MadeTimes = &madeTimes
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	return scMsg
}

func BuildSCOpenActivityChargeRewardsLimitInfo(obj *playerwelfare.PlayerOpenActivityObject, timesList []*welfaretypes.TimesLimitInfo, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityChargeRewardsLimitInfo {
	scMsg := &uipb.SCOpenActivityChargeRewardsLimitInfo{}
	chargeNum := int32(0)
	playerTimesMap := make(map[int32]int32)
	if obj != nil {
		info := obj.GetActivityData().(*rewardschargelimittypes.ChargeRewLimitInfo)
		chargeNum = info.GoldNum
		playerTimesMap = info.ReceiveTimesMap
	}

	scMsg.ChargeNum = &chargeNum
	scMsg.GroupId = &groupId
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime

	for _, data := range timesList {
		level := data.Key
		times := data.Times
		scMsg.TimesInfo = append(scMsg.TimesInfo, buildRewardsTimesInfo(level, times))
	}

	for level, times := range playerTimesMap {
		scMsg.PlayerTimesInfo = append(scMsg.PlayerTimesInfo, buildRewardsTimesInfo(level, times))
	}

	return scMsg
}

func BuildSCOpenActivityFeedbackCostInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityFeedbackCostInfo {
	scOpenActivityFeedbackCostInfo := &uipb.SCOpenActivityFeedbackCostInfo{}

	var record []int32
	costNum := int32(0)
	if obj != nil {
		costInfo := obj.GetActivityData().(*feedbackcosttypes.FeedbackCostInfo)
		record = costInfo.RewRecord
		costNum = int32(costInfo.GoldNum)
	}
	info := &uipb.FeedbackCostInfo{}
	info.CostNum = &costNum
	info.CostRecordList = record
	info.GroupId = &groupId
	info.StartTime = &startTime
	info.EndTime = &endTime
	scOpenActivityFeedbackCostInfo.FeedbackCostInfo = info
	return scOpenActivityFeedbackCostInfo
}

func BuildSCOpenActivityCycleChargeInfo(cycleObj *playerwelfare.PlayerOpenActivityObject, groupId int32) *uipb.SCOpenActivityCycleChargeInfo {
	scOpenActivityCycleChargeInfo := &uipb.SCOpenActivityCycleChargeInfo{}

	var record []int32
	goldNum := int32(0)
	if cycleObj != nil {
		cycleInfo := cycleObj.GetActivityData().(*cyclechargetypes.CycleChargeInfo)
		record = cycleInfo.RewRecord
		goldNum = cycleInfo.GoldNum
	}

	scOpenActivityCycleChargeInfo.GoldNum = &goldNum
	scOpenActivityCycleChargeInfo.GroupId = &groupId
	scOpenActivityCycleChargeInfo.RewRecordList = record

	return scOpenActivityCycleChargeInfo
}

func BuildSCOpenActivityCycleSingleChargeInfo(cycleObj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityCycleSingleChargeInfo {
	scMsg := &uipb.SCOpenActivityCycleSingleChargeInfo{}

	var record []int32
	goldNum := int32(0)
	if cycleObj != nil {
		info := cycleObj.GetActivityData().(*cyclechargesingletypes.CycleSingleChargeInfo)
		record = info.RewRecord
		goldNum = info.MaxSingleChargeNum
	}

	scMsg.MaxSingleChargeNum = &goldNum
	scMsg.GroupId = &groupId
	scMsg.RewRecordList = record
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime

	return scMsg
}

func BuildSCMergeActivityAdvancedInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCMergeActivityAdvancedInfo {
	scMergeActivityAdvancedInfo := &uipb.SCMergeActivityAdvancedInfo{}

	var record []int32
	danNum := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*advancedfeedbacktypes.AdvancedInfo)
		record = info.RewRecord
		danNum = info.DanNum
	}

	scMergeActivityAdvancedInfo.DanNum = &danNum
	scMergeActivityAdvancedInfo.GroupId = &groupId
	scMergeActivityAdvancedInfo.RewRecordList = record
	scMergeActivityAdvancedInfo.StartTime = &startTime
	scMergeActivityAdvancedInfo.EndTime = &endTime

	return scMergeActivityAdvancedInfo
}

func BuildSCOpenActivityGetInfo(groupId int32, startTime, endTime int64, hadReceiveRecord []int32) *uipb.SCOpenActivityGetInfo {
	scMsg := &uipb.SCOpenActivityGetInfo{}
	scMsg.GroupId = &groupId
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	scMsg.RecordList = hadReceiveRecord
	return scMsg
}

func BuildSCOpenActivityGetInfoGoal(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, goalId, goalCount int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.GoalInfo = buildWelfareGoalInfo(goalId, goalCount)
	return scMsg
}

func BuildSCOpenActivityGetInfoQiYuDao(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, npcMap map[int64]scene.NPC) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.QiYuInfo = buildWelfareQiYuInfo(npcMap)
	return scMsg
}

func BuildSCOpenActivityGetInfoZhuanSheng(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, zhuanSheng int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.ZhuanShengInfo = buildZhuanShengInfo(zhuanSheng)
	return scMsg
}

func BuildSCOpenActivityGetInfoDevelopFame(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, favorableNum, dayFavorableNum int32, feedTimesMap map[int32]int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.DevelopInfo = buildDevelopInfo(favorableNum, dayFavorableNum, feedTimesMap)
	return scMsg
}

func BuildSCOpenActivityGetInfoAdvancedExpendReturn(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, useNum int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.ExpendReturnInfo = buildAdvancedExpendReturnInfo(useNum)
	return scMsg
}

func BuildSCOpenActivityGetInfoAdvancedRewMax(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, chargeNum, initAdvancedNum int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.AdvancedRewMaxInfo = buildAdvancedRewMaxInfo(chargeNum, initAdvancedNum)
	return scMsg
}

func BuildSCOpenActivityGetInfoHouseExtended(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, activateChargeNum, uplevelChargeNum, uplevelGiftLevel int32, isActivateGift, isUplevelGift bool) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.HouseExtendedInfo = buildFeedbackHouseExtendedInfo(activateChargeNum, uplevelChargeNum, uplevelGiftLevel, isActivateGift, isUplevelGift)
	return scMsg
}

func BuildSCOpenActivityGetInfoAdvancedTimesReturn(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, advancedTimes int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.TimesReturnInfo = buildAdvancedTimesReturnInfo(advancedTimes)
	return scMsg
}

func BuildSCOpenActivityGetInfoChargeReturn(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, isReturn bool) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.ReturnInfo = buildChargeReturnInfo(isReturn)
	return scMsg
}

func BuildSCOpenActivityGetInfoChargeReturnLevel(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, isReturn bool) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.ReturnLevelInfo = buildChargeReturnLevelInfo(isReturn)
	return scMsg
}

func BuildSCOpenActivityGetInfoChargeReturnMultiple(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, periodChargeNum int32, rewardCnt int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.ReturnMultipleInfo = buildChargeReturnMultipleInfo(periodChargeNum, rewardCnt)
	return scMsg
}

func BuildSCOpenActivityGetInfoAllianceCheerInfo(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, poolNum int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.CheerInfo = buildAllianceCheerInfo(poolNum)
	return scMsg
}

func BuildSCOpenActivityGetInfoSystemActivate(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, maxSingleNum int32, isActivate bool) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.SystemInfo = buildSystemActivateInfo(maxSingleNum, isActivate)
	return scMsg
}

func BuildSCOpenActivityGetInfoCycleSingleMaxRew(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, maxSingleNum int32, canRewRecord []int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.CycleSingleMaxRewInfo = buildCycleSingleChargeMaxRewInfo(maxSingleNum, canRewRecord)
	return scMsg

}
func BuildSCOpenActivityGetInfoFeedbackSingleMaxRew(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, maxSingleNum int32, canRewRecord []int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.SingleChargeMax = buildFeedbackSingleChargeMaxRewInfo(maxSingleNum, canRewRecord)
	return scMsg
}

func BuildSCOpenActivityCycleSingleChargeMaxRewInfoNotice(groupId int32, maxSingleNum int32, canRewRecord []int32) *uipb.SCOpenActivityCycleSingleChargeMaxRewInfoNotice {
	scMsg := &uipb.SCOpenActivityCycleSingleChargeMaxRewInfoNotice{}
	scMsg.GroupId = &groupId
	scMsg.Info = buildCycleSingleChargeMaxRewInfo(maxSingleNum, canRewRecord)
	return scMsg
}

func BuildSCOpenActivityGetInfoCycleSingleMaxRewMultiple(groupId int32, startTime, endTime int64, hadReceiveRecord map[int32]int32, maxSingleNum int32, canRewRecord map[int32]int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, nil)
	scMsg.RecordInfo = buildTimesInfo(hadReceiveRecord)
	scMsg.CycleSingleMaxRewMultipleInfo = buildCycleSingleChargeMaxRewMultipleInfo(maxSingleNum, canRewRecord)
	return scMsg
}

func BuildSCOpenActivityGetInfoSmelt(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, canRewRecord int32, useItemNum int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.SmeltInfo = buildSmeltInfo(canRewRecord, useItemNum)
	return scMsg
}

func buildSmeltInfo(canRewRecord int32, useItemNum int32) *uipb.SmeltInfo {
	msg := &uipb.SmeltInfo{}
	msg.Num = &useItemNum
	msg.CanRewRecord = &canRewRecord
	return msg
}

func BuildSCOpenActivityGetInfoRewPools(groupId int32, startTime, endTime int64, position int32, backTimes int32, logList []*welfare.DrewLogObject) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, nil)
	scMsg.RewPoolsInfo = buildRewPoolsInfo(position, backTimes, logList)
	return scMsg
}

func buildRewPoolsInfo(position int32, backTimes int32, logList []*welfare.DrewLogObject) *uipb.RewPoolsInfo {
	msg := &uipb.RewPoolsInfo{}
	msg.Position = &position
	msg.BackTimes = &backTimes
	for _, log := range logList {
		msg.LogList = append(msg.LogList, buildDrewLog(log))
	}
	return msg
}

func BuildSCOpenActivityGetInfoXiuxianBook(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, level int32, maxLevel int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.XiuxianBookInfo = buildXiuxianBookInfo(level, maxLevel)
	return scMsg
}

func buildXiuxianBookInfo(level int32, maxLevel int32) *uipb.XiuxianBookInfo {
	msg := &uipb.XiuxianBookInfo{
		Level:    &level,
		MaxLevel: &maxLevel,
	}
	return msg
}

func BuildSCOpenActivityGetInfoCycleSingleAllRewMultiple(groupId int32, startTime, endTime int64, singleChargeRecord []int32, canRewRecord map[int32]int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, nil)
	scMsg.RecordInfo = buildTimesInfo(canRewRecord)
	scMsg.CycleAllMultipleRewInfo = buildCycleSingleChargeAllMultipleRewInfo(singleChargeRecord, canRewRecord)
	return scMsg
}

func buildCycleSingleChargeAllMultipleRewInfo(singleChargeRecord []int32, canRewRecord map[int32]int32) *uipb.CycleSingleChargeAllMultipleRewInfo {
	scMsg := &uipb.CycleSingleChargeAllMultipleRewInfo{}
	scMsg.SingleChargeRecord = singleChargeRecord
	scMsg.CanReceiveRecordInfo = buildTimesInfo(canRewRecord)
	return scMsg
}

func BuildSCOpenActivityCycleSingleChargeAllRewInfoNotice(groupId int32, singleChargeRecord []int32, canRewRecord map[int32]int32) *uipb.SCOpenActivityCycleSingleChargeAllRewInfoNotice {
	scMsg := &uipb.SCOpenActivityCycleSingleChargeAllRewInfoNotice{}
	scMsg.GroupId = &groupId
	scMsg.SingleChargeRecord = singleChargeRecord
	scMsg.CanReceiveRecordInfo = buildTimesInfo(canRewRecord)
	return scMsg
}

func BuildSCOpenActivityCycleSingleChargeMaxRewMultipleInfoNotice(groupId int32, maxSingleNum int32, canRewRecord map[int32]int32) *uipb.SCOpenActivityCycleSingleChargeMaxRewInfoNotice {
	scMsg := &uipb.SCOpenActivityCycleSingleChargeMaxRewInfoNotice{}
	scMsg.GroupId = &groupId
	scMsg.MultipleInfo = buildCycleSingleChargeMaxRewMultipleInfo(maxSingleNum, canRewRecord)
	return scMsg
}

func BuildSCOpenActivityGetInfoCostDrew(groupId, costNum, leftTimes, attendTimes, rate int32, logList []*welfare.DrewLogObject, startTime, endTime int64) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, nil)
	scMsg.CostDrewInfo = buildCostDrewInfo(leftTimes, costNum, attendTimes, rate, logList)
	return scMsg
}

func BuildSCOpenActivityFeedbackChargeArenapvpAssistInfo(groupId int32, startTime, endTime int64, hadReceiveRecord []int32, chargeNum int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadReceiveRecord)
	scMsg.ArenapvpAssistInfo = buildFeedbackChargeArenapvpAssistInfo(chargeNum)
	return scMsg
}

func BuildSCOpenActivityFeedbackChargeArenapvpAssistReturnInfo(groupId int32, startTime, endTime int64, costNum int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, nil)
	scMsg.ArenapvpAssistReturnInfo = buildFeedbackChargeArenapvpAssistReturnInfo(costNum)
	return scMsg
}

func BuildSCOpenActivityFeedbackChargeNewArenapvpAssistReturnInfo(groupId int32, startTime, endTime int64, costNum int64) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, nil)
	scMsg.NewArenapvpAssistReturnInfo = buildFeedbackChargeNewArenapvpAssistReturnInfo(costNum)
	return scMsg
}

func BuildSCOpenActivityRankingDropDownNotice(rankName string, befRanking, curRanking, groupId int32) *uipb.SCOpenActivityRankingDropDownNotice {
	scMsg := &uipb.SCOpenActivityRankingDropDownNotice{}
	scMsg.BeforeRanking = &befRanking
	scMsg.CurRanking = &curRanking
	scMsg.RankName = &rankName
	scMsg.GroupId = &groupId
	return scMsg
}

func buildAdvancedExpendReturnInfo(useNum int32) *uipb.AdvancedExpendReturnInfo {
	scMsg := &uipb.AdvancedExpendReturnInfo{}
	scMsg.UseItemNum = &useNum
	return scMsg
}

func buildFeedbackChargeArenapvpAssistInfo(chargeNum int32) *uipb.FeedbackChargeArenapvpAssistInfo {
	scMsg := &uipb.FeedbackChargeArenapvpAssistInfo{}
	scMsg.ChargeNum = &chargeNum
	return scMsg
}

func buildFeedbackChargeArenapvpAssistReturnInfo(costNum int32) *uipb.FeedbackChargeArenapvpAssistReturnInfo {
	scMsg := &uipb.FeedbackChargeArenapvpAssistReturnInfo{}
	scMsg.CostNum = &costNum
	return scMsg
}

func buildFeedbackChargeNewArenapvpAssistReturnInfo(costNum int64) *uipb.FeedbackChargeNewArenapvpAssistReturnInfo {
	scMsg := &uipb.FeedbackChargeNewArenapvpAssistReturnInfo{}
	scMsg.CostNum = &costNum
	return scMsg
}
func buildAdvancedRewMaxInfo(chargeNum, initAdvancedNum int32) *uipb.AdvancedRewMaxInfo {
	scMsg := &uipb.AdvancedRewMaxInfo{}
	scMsg.PeriodChargeNum = &chargeNum
	scMsg.InitAdvancedNum = &initAdvancedNum
	return scMsg
}

func buildFeedbackHouseExtendedInfo(activateChargeNum, uplevelChargeNum, uplevelGiftLevel int32, isActivateGift, isUplevelGift bool) *uipb.FeedbackHouseExtendedInfo {
	scMsg := &uipb.FeedbackHouseExtendedInfo{}
	scMsg.ActivateChargeNum = &activateChargeNum
	scMsg.UplevelChargeNum = &uplevelChargeNum
	scMsg.CurUplevelGiftLevel = &uplevelGiftLevel
	scMsg.IsActivateGift = &isActivateGift
	scMsg.IsUplevelGift = &isUplevelGift
	return scMsg
}

func buildAdvancedTimesReturnInfo(advancedTimes int32) *uipb.AdvancedTimesReturnInfo {
	scMsg := &uipb.AdvancedTimesReturnInfo{}
	scMsg.AdvancedTimes = &advancedTimes
	return scMsg
}

func buildDevelopInfo(favorableNum, dayFavorableNum int32, feedTimesMap map[int32]int32) *uipb.DevelopInfo {
	scMsg := &uipb.DevelopInfo{}
	scMsg.FavorableNum = &favorableNum
	scMsg.DayFavorableNum = &dayFavorableNum
	scMsg.TimesInfoList = buildTimesInfo(feedTimesMap)
	return scMsg
}

func buildZhuanShengInfo(zhuanSheng int32) *uipb.ZhuanShengInfo {
	info := &uipb.ZhuanShengInfo{}
	info.ZhuanSheng = &zhuanSheng
	return info
}

func buildWelfareQiYuInfo(npcMap map[int64]scene.NPC) *uipb.WelfareQiYuInfo {
	info := &uipb.WelfareQiYuInfo{}
	info.BiologyInfo = scenepbutil.BuildGeneralCollectInfoList(npcMap)
	return info
}

func buildWelfareGoalInfo(goalId, goalCount int32) *uipb.WelfareGoalInfo {
	info := &uipb.WelfareGoalInfo{}
	info.GoalId = &goalId
	info.GoalCount = &goalCount
	return info
}

func buildTimesInfo(timesMap map[int32]int32) (timesList []*uipb.TimesInfo) {
	for key, value := range timesMap {
		itemId := key
		times := value
		info := &uipb.TimesInfo{}
		info.Key = &itemId
		info.Value = &times
		timesList = append(timesList, info)
	}
	return timesList
}

func buildChargeReturnInfo(isReturn bool) *uipb.ChargeReturnInfo {
	scMsg := &uipb.ChargeReturnInfo{}
	scMsg.IsReturn = &isReturn
	return scMsg
}

func buildChargeReturnLevelInfo(isReturn bool) *uipb.ChargeReturnLevelInfo {
	scMsg := &uipb.ChargeReturnLevelInfo{}
	scMsg.IsReturn = &isReturn
	return scMsg
}

func buildChargeReturnMultipleInfo(periodChargeNum int32, rewardCnt int32) *uipb.ChargeReturnMultipleInfo {
	scMsg := &uipb.ChargeReturnMultipleInfo{}
	scMsg.PeriodChargeNum = &periodChargeNum
	scMsg.RewardCnt = &rewardCnt
	return scMsg
}

func buildAllianceCheerInfo(poolNum int32) *uipb.AllianceCheerInfo {
	scMsg := &uipb.AllianceCheerInfo{}
	scMsg.RewPool = &poolNum
	return scMsg
}

func buildSystemActivateInfo(maxSingleNum int32, isActivate bool) *uipb.SystemActivateInfo {
	scMsg := &uipb.SystemActivateInfo{}
	scMsg.IsActivate = &isActivate
	scMsg.MaxSingleChargeNum = &maxSingleNum
	return scMsg
}

func buildCycleSingleChargeMaxRewInfo(maxSingleNum int32, canRewRecord []int32) *uipb.CycleSingleChargeMaxRewInfo {
	scMsg := &uipb.CycleSingleChargeMaxRewInfo{}
	scMsg.CanReceiveRecordList = canRewRecord
	scMsg.MaxSingleChargeNum = &maxSingleNum
	return scMsg
}

func buildCycleSingleChargeMaxRewMultipleInfo(maxSingleNum int32, canRewRecord map[int32]int32) *uipb.CycleSingleChargeMaxRewMultipleInfo {
	scMsg := &uipb.CycleSingleChargeMaxRewMultipleInfo{}
	scMsg.MaxSingleChargeNum = &maxSingleNum
	scMsg.CanReceiveRecordInfo = buildTimesInfo(canRewRecord)
	return scMsg
}

func buildFeedbackSingleChargeMaxRewInfo(maxSingleNum int32, canRewRecord []int32) *uipb.FeedbackSingleChargeMaxRewInfo {
	scMsg := &uipb.FeedbackSingleChargeMaxRewInfo{}
	scMsg.CanReceiveRecordList = canRewRecord
	scMsg.MaxSingleChargeNum = &maxSingleNum
	return scMsg
}

func buildCostDrewInfo(leftTimes, costNum, attendTimes, rate int32, logList []*welfare.DrewLogObject) *uipb.CostDrewInfo {
	scMsg := &uipb.CostDrewInfo{}

	scMsg.DrewTimes = &leftTimes
	scMsg.CostNum = &costNum
	scMsg.AttendTimes = &attendTimes
	scMsg.Rate = &rate
	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildDrewLog(log))
	}

	return scMsg
}

func BuildSCMergeActivityAdvancedBlessInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCMergeActivityAdvancedBlessInfo {
	scMergeActivityAdvancedBlessInfo := &uipb.SCMergeActivityAdvancedBlessInfo{}

	var record []int32
	if obj != nil {
		info := obj.GetActivityData().(*advancedblessfeedbacktypes.BlessAdvancedInfo)
		record = info.RewRecord
	}

	scMergeActivityAdvancedBlessInfo.GroupId = &groupId
	scMergeActivityAdvancedBlessInfo.RewRecordList = record
	scMergeActivityAdvancedBlessInfo.StartTime = &startTime
	scMergeActivityAdvancedBlessInfo.EndTime = &endTime

	return scMergeActivityAdvancedBlessInfo
}

func BuildSCOpenActivityTimesRewInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityTimesRewInfo {
	scMsg := &uipb.SCOpenActivityTimesRewInfo{}

	times := int32(0)
	record := map[int32][]int32{}
	if obj != nil {
		info := obj.GetActivityData().(*grouptimesrewtypes.TimesRewInfo)
		record = info.RewRecord
		times = info.Times
	}
	scMsg.RecordList = buildTimesRewRecordList(record)
	scMsg.GroupId = &groupId
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	scMsg.Times = &times

	return scMsg
}

func BuildSCOpenActivityAdvancedRewInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityAdvancedRewInfo {
	scMsg := &uipb.SCOpenActivityAdvancedRewInfo{}

	var record []int32
	periodChargeNum := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*advancedrewrewtypes.AdvancedRewInfo)
		record = info.RewRecord
		periodChargeNum = info.PeriodChargeNum
	}

	scMsg.GroupId = &groupId
	scMsg.RewRecordList = record
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	scMsg.PeriodChargeNum = &periodChargeNum

	return scMsg
}

func BuildSCOpenActivityAdvancedPowerRewInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityAdvancedPowerRewInfo {
	scMsg := &uipb.SCOpenActivityAdvancedPowerRewInfo{}

	var record []int64
	expire := int64(0)
	if obj != nil {
		info := obj.GetActivityData().(*advancedrewpowertypes.AdvancedPowerInfo)
		record = info.RewRecord
		expire = info.ExpireTime
	}

	scMsg.GroupId = &groupId
	scMsg.RewRecordList = record
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	scMsg.FreeExpireTime = &expire
	return scMsg
}

func BuildSCOpenActivityAdvancedRewExtendedInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityAdvancedRewExtendedInfo {
	scMsg := &uipb.SCOpenActivityAdvancedRewExtendedInfo{}

	var record []int32
	expire := int64(0)
	if obj != nil {
		info := obj.GetActivityData().(*advancedrewrewextendedtypes.AdvancedRewExtendedInfo)
		record = info.RewRecord
		expire = info.ExpireTime
	}

	scMsg.GroupId = &groupId
	scMsg.RewRecordList = record
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	scMsg.FreeExpireTime = &expire
	return scMsg
}

func BuildSCOpenActivityRealmInfo(realmObj *playerwelfare.PlayerOpenActivityObject, groupId int32, timesList []*welfaretypes.TimesLimitInfo, startTime, endTime int64) *uipb.SCOpenActivityRealmInfo {
	scMsg := &uipb.SCOpenActivityRealmInfo{}

	var record []int32
	level := int32(0)
	if realmObj != nil {
		realmInfo := realmObj.GetActivityData().(*hallrealmtypes.WelfareRealmChallengeInfo)
		record = realmInfo.RewRecord
		level = realmInfo.Level
	}

	scMsg.Level = &level
	scMsg.GroupId = &groupId
	scMsg.RewRecordList = record
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime

	for _, data := range timesList {
		level := data.Key
		times := data.Times
		scMsg.TimesInfo = append(scMsg.TimesInfo, buildRewardsTimesInfo(level, times))
	}

	return scMsg
}

func BuildSCMergeActivitySingleChargeInfo(singleChargeObj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCMergeActivitySingleChargeInfo {
	scMergeActivitySingleChargeInfo := &uipb.SCMergeActivitySingleChargeInfo{}

	var record []int32
	maxSingleChargeNum := int32(0)
	if singleChargeObj != nil {
		singleChargeInfo := singleChargeObj.GetActivityData().(*feedbackchargesingletypes.FeedbackSingleChargeInfo)
		record = singleChargeInfo.RewRecord
		maxSingleChargeNum = singleChargeInfo.MaxSingleChargeNum
	}

	scMergeActivitySingleChargeInfo.MaxSingleChargeNum = &maxSingleChargeNum
	scMergeActivitySingleChargeInfo.GroupId = &groupId
	scMergeActivitySingleChargeInfo.RewRecordList = record
	scMergeActivitySingleChargeInfo.StartTime = &startTime
	scMergeActivitySingleChargeInfo.EndTime = &endTime

	return scMergeActivitySingleChargeInfo
}

func BuildSCMergeActivityDiscountInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId, discountDay int32, globalTimesList []*welfaretypes.TimesLimitInfo, startTime, endTime int64) *uipb.SCMergeActivityDiscountInfo {
	scMsg := &uipb.SCMergeActivityDiscountInfo{}

	record := make(map[int32]int32)
	if obj != nil {
		info := obj.GetActivityData().(*discountdiscounttypes.DiscountInfo)
		record = info.BuyRecord
		groupId = obj.GetGroupId()
	}

	scMsg.DiscountDay = &discountDay
	scMsg.GroupId = &groupId
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	for index, num := range record {
		scMsg.BuyRecordList = append(scMsg.BuyRecordList, buildDiscountInfo(index, num))
	}
	for _, data := range globalTimesList {
		itemIndex := data.Key
		globalNum := data.Times
		scMsg.GlobalTimesList = append(scMsg.GlobalTimesList, buildDiscountGlobalTimesInfo(itemIndex, globalNum))
	}

	return scMsg
}

func BuildSCOpenActivityKanJiaInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityKanJiaInfo {
	scMsg := &uipb.SCOpenActivityKanJiaInfo{}

	kanJiaTimes := int32(0)
	chargeNum := int64(0)
	buyRecord := []int32{}
	kanjiaRecord := make(map[int32]*discountkanjiatypes.KanJiaInfo)
	if obj != nil {
		info := obj.GetActivityData().(*discountkanjiatypes.DiscountKanJiaInfo)
		kanJiaTimes = info.UseTimes
		chargeNum = info.GoldNum
		buyRecord = info.BuyRecord
		kanjiaRecord = info.KanJiaRecord
	}

	scMsg.GroupId = &groupId
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	scMsg.UseTimes = &kanJiaTimes
	scMsg.ChargeNum = &chargeNum
	scMsg.BuyRecordList = buyRecord
	scMsg.KanjiaInfoList = buildKanJiaInfoList(kanjiaRecord)
	return scMsg
}

func BuildSCOpenActivityBeachShopInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime int64, endTime int64, record []int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, record)
	if obj != nil {
		info := obj.GetActivityData().(*discountbeachtypes.BeachShopInfo)
		isActivite := info.IsActivite
		scMsg.IsActivite = &isActivite

		var beachShopInfoList []*uipb.OpenActivityBeachShopInfo
		for typ, num := range info.BuyRecord {
			t := typ
			n := num
			beachShopInfo := &uipb.OpenActivityBeachShopInfo{}
			beachShopInfo.Typ = &t
			beachShopInfo.Num = &n
			beachShopInfoList = append(beachShopInfoList, beachShopInfo)
		}

		scMsg.BeachShopInfo = beachShopInfoList
	}

	return scMsg
}

func BuildSCOpenActivityTongTianTaInfo(minForce int32, maxForce int32, groupId int32, startTime int64, endTime int64, record []int32) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, record)
	scMsg.TongTianTaInfo = buildOpenActivityTongTianTaInfo(minForce, maxForce)
	return scMsg
}

func buildOpenActivityTongTianTaInfo(minForce int32, maxForce int32) *uipb.TongTianTaInfo {
	tongTianTaInfo := &uipb.TongTianTaInfo{}
	tongTianTaInfo.MinForce = &minForce
	tongTianTaInfo.MaxForce = &maxForce
	return tongTianTaInfo
}

func BuildSCOpenActivityNewSevenDayInvestInfo(groupId int32, startTime int64, endTime int64, record []int32, receiveMap map[int32]int32, buyTimeMap map[int32]int64) *uipb.SCOpenActivityGetInfo {
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, record)
	scMsg.NewInvestSevenDayInfo = buildSCOpenActivityNewInvestSevenDayInfo(receiveMap, buyTimeMap)
	return scMsg
}

func buildSCOpenActivityNewInvestSevenDayInfo(receiveMap map[int32]int32, buyTimeMap map[int32]int64) *uipb.NewInvestSevenDayInfo {
	newInvestSevenDayInfo := &uipb.NewInvestSevenDayInfo{}
	for typ, buyTime := range buyTimeMap {
		t := typ
		b := buyTime
		receiveDay, _ := receiveMap[typ]
		newInvestSevenDayInfo.NewInvestDayInfo = append(newInvestSevenDayInfo.NewInvestDayInfo, buildSCOpenActivityNewInvestDayInfo(t, receiveDay, b))
	}
	return newInvestSevenDayInfo
}

func buildSCOpenActivityNewInvestDayInfo(typ int32, investDayRecord int32, buyTime int64) *uipb.NewInvestDayInfo {
	newInvestDayInfo := &uipb.NewInvestDayInfo{}
	newInvestDayInfo.Typ = &typ
	newInvestDayInfo.InvestDayRecord = &investDayRecord
	newInvestDayInfo.BuyTime = &buyTime
	return newInvestDayInfo
}

func BuildSCOpenActivityYunYinShopInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime int64, endTime int64) *uipb.SCOpenActivityGetInfo {
	info := obj.GetActivityData().(*discountyunyintypes.YunYinInfo)
	scMsg := BuildSCOpenActivityGetInfo(groupId, startTime, endTime, info.ReceiveRecord)
	yunYinShopInfo := &uipb.YunYinShopInfo{}
	yunYinShopInfo.GoldNum = &info.GoldNum

	for typ, num := range info.BuyRecord {
		t := typ
		n := num
		yunYinCommodity := &uipb.YunYinCommodity{}
		yunYinCommodity.Typ = &t
		yunYinCommodity.Num = &n
		yunYinShopInfo.YunYinCommodity = append(yunYinShopInfo.YunYinCommodity, yunYinCommodity)
	}
	scMsg.YunYinShopInfo = yunYinShopInfo

	return scMsg
}

func BuildSCOpenActivityBeachShopActivite(groupId int32) *uipb.SCOpenActivityBeachShopActivate {
	scMsg := &uipb.SCOpenActivityBeachShopActivate{}
	scMsg.GroupId = &groupId
	return scMsg
}

func BuildSCOpenActivityBeachShopBuy(groupId int32, typ int32, num int32, itemMap map[int32]int32) *uipb.SCOpenActivityBeachBuy {
	scMsg := &uipb.SCOpenActivityBeachBuy{}
	scMsg.GroupId = &groupId
	scMsg.Type = &typ
	scMsg.Num = &num
	dropList := droppbutil.BuildSimpleDropInfoList(itemMap)
	scMsg.DropInfo = dropList
	return scMsg
}

func BuildSCOpenActivityYunYunShopBuy(groupId int32, typ int32, num int32, itemMap map[int32]int32, gold int32) *uipb.SCOpenActivityYunYinBuy {
	scOpenActivityYunYinBuy := &uipb.SCOpenActivityYunYinBuy{}
	scOpenActivityYunYinBuy.GroupId = &groupId
	scOpenActivityYunYinBuy.Type = &typ
	scOpenActivityYunYinBuy.Num = &num
	scOpenActivityYunYinBuy.GoldNum = &gold
	dropList := droppbutil.BuildSimpleDropInfoList(itemMap)
	scOpenActivityYunYinBuy.DropInfo = dropList
	return scOpenActivityYunYinBuy
}

func BuildSCOpenActivityTaoCanInfo(groupId int32, isBuyHuiYaunPlus, isBuyEquipGift, isBuyInvestLevel bool) *uipb.SCOpenActivityTaoCanInfo {
	scMsg := &uipb.SCOpenActivityTaoCanInfo{}
	scMsg.GroupId = &groupId
	scMsg.IsBuyHuiYaunPlus = &isBuyHuiYaunPlus
	scMsg.IsBuyInvestLevel = &isBuyInvestLevel
	scMsg.IsBuyEquipGift = &isBuyEquipGift
	return scMsg
}

func BuildSCOpenActivityDiscountZhuanShengInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityDiscountZhuanShengInfo {
	scMsg := &uipb.SCOpenActivityDiscountZhuanShengInfo{}

	record := make(map[int32]int32)
	giveRecord := make(map[int32]int32)
	chargeNum := int64(0)
	usePoint := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*discountzhuanshengtypes.DiscountZhuanShengInfo)
		record = info.BuyRecord
		chargeNum = info.ChargeNum
		giveRecord = info.GiftReceiveRecord
		usePoint = info.UsePoint
	}

	scMsg.GroupId = &groupId
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	scMsg.ChargeNum = &chargeNum
	scMsg.UsePoint = &usePoint
	scMsg.BuyRecordList = buildBuyTimesInfoList(record)
	scMsg.ReceiveRecordList = buildBuyTimesInfoList(giveRecord)

	return scMsg
}

func BuildSCMergeActivityGoldBowlInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCMergeActivityGoldBowlInfo {
	scMergeActivityGoldBowlInfo := &uipb.SCMergeActivityGoldBowlInfo{}

	costNum := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*feedbackgoldbowltypes.FeedbackGoldBowlInfo)
		costNum = int32(info.GoldNum)
	}

	scMergeActivityGoldBowlInfo.GroupId = &groupId
	scMergeActivityGoldBowlInfo.CostNum = &costNum
	scMergeActivityGoldBowlInfo.StartTime = &startTime
	scMergeActivityGoldBowlInfo.EndTime = &endTime

	return scMergeActivityGoldBowlInfo
}

func BuildSCOpenActivityFeedbackGoldPigInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityFeedbackGoldPigInfo {
	scMsg := &uipb.SCOpenActivityFeedbackGoldPigInfo{}

	costNum := int32(0)
	chargeNum := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*feedbackpigtypes.FeedbackGoldPigInfo)
		chargeNum = info.ChargeGold
		costNum = info.CostGold
	}

	scMsg.GroupId = &groupId
	scMsg.CostNum = &costNum
	scMsg.ChargeNum = &chargeNum
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime

	return scMsg
}

func BuildSCOpenActivityFeedbackGoldLaBaInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, logList []*welfare.GoldLaBaLogObject, startTime, endTime int64) *uipb.SCOpenActivityFeedbackGoldLaBaInfo {
	scMsg := &uipb.SCOpenActivityFeedbackGoldLaBaInfo{}

	chargeNum := int32(0)
	times := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*feedbacklabatypes.FeedbackGoldLaBaInfo)
		chargeNum = int32(info.ChargeNum)
		times = int32(info.Times)
	}

	scMsg.GroupId = &groupId
	scMsg.CurChargeNum = &chargeNum
	scMsg.LabaTimes = &times
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime

	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildLaBaLog(log))
	}

	return scMsg
}

func BuildSCOpenActivityChargeDrewInfo(obj *playerwelfare.PlayerOpenActivityObject, logList []*welfare.DrewLogObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityChargeDrewInfo {
	scMsg := &uipb.SCOpenActivityChargeDrewInfo{}

	chargeNum := int32(0)
	leftTimes := int32(0)
	attendTimes := int32(0)
	rate := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*drewchargedrewtypes.LuckyChargeDrewInfo)
		chargeNum = int32(info.GoldNum)
		leftTimes = info.LeftTimes
		attendTimes = info.AttendTimes
		rate = info.Ratio
	}

	scMsg.GroupId = &groupId
	scMsg.DrewTimes = &leftTimes
	scMsg.ChargeNum = &chargeNum
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	scMsg.AttendTimes = &attendTimes
	scMsg.Rate = &rate
	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildDrewLog(log))
	}

	return scMsg
}

func BuildSCOpenActivityCrazyBoxInfo(obj *playerwelfare.PlayerOpenActivityObject, logList []*welfare.CrazyBoxLogObject, groupId int32, startTime int64, endTime int64, drewTimes int32, curBoxLevel int32, curBoxTimes int32) *uipb.SCOpenActivityCrazyBoxInfo {
	scMsg := &uipb.SCOpenActivityCrazyBoxInfo{}

	goldNum := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*drewcrazyboxtypes.CrazyBoxInfo)
		goldNum = int32(info.GoldNum)
	}

	scMsg.GroupId = &groupId
	scMsg.DrewTimes = &drewTimes
	scMsg.CostNum = &goldNum
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	scMsg.CurBoxLev = &curBoxLevel
	scMsg.CurBoxTimes = &curBoxTimes
	for _, log := range logList {
		scMsg.LogList = append(scMsg.LogList, buildCrazyBoxLog(log))
	}

	return scMsg
}

func BuildSCOpenActivitySmashEggInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime int64, endTime int64) *uipb.SCOpenActivitySmashEggInfo {
	scMsg := &uipb.SCOpenActivitySmashEggInfo{}

	attendTimes := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*drewsmasheggtypes.SmashEggInfo)
		attendTimes = info.AttendTimes
		for i := 0; i < len(info.AttendTimesList); i++ {
			scMsg.Records = append(scMsg.Records, info.AttendTimesList[i])
		}
	}
	scMsg.GroupId = &groupId
	scMsg.AttendTimes = &attendTimes
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime

	return scMsg
}

func BuildSCOpenActivityFeedbackCycleChargeInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityFeedbackCycleChargeInfo {
	scOpenActivityFeedbackCycleChargeInfo := &uipb.SCOpenActivityFeedbackCycleChargeInfo{}

	curChargeNum := int32(0)
	dayNum := int32(0)
	var record []int32
	isReceive := false
	if obj != nil {
		info := obj.GetActivityData().(*feedbackchargecycletypes.FeedbackCycleChargeInfo)
		curChargeNum = info.CurDayChargeNum
		dayNum = info.DayNum
		record = info.RewRecord
		isReceive = info.IsReceiveDayRew
	}

	scOpenActivityFeedbackCycleChargeInfo.GroupId = &groupId
	scOpenActivityFeedbackCycleChargeInfo.CurChargeNum = &curChargeNum
	scOpenActivityFeedbackCycleChargeInfo.DayNum = &dayNum
	scOpenActivityFeedbackCycleChargeInfo.RewRecordList = record
	scOpenActivityFeedbackCycleChargeInfo.IsReceive = &isReceive
	scOpenActivityFeedbackCycleChargeInfo.StartTime = &startTime
	scOpenActivityFeedbackCycleChargeInfo.EndTime = &endTime

	return scOpenActivityFeedbackCycleChargeInfo
}

func BuildSCOpenActivityFeedbackHouseInvestInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId int32, startTime, endTime int64) *uipb.SCOpenActivityFeedbackHouseInvestInfo {
	scOpenActivityFeedbackHouseInvestInfo := &uipb.SCOpenActivityFeedbackHouseInvestInfo{}

	chargeNum := int32(0)
	curDayChargeNum := int32(0)
	decorDays := int32(0)
	isActivity := false
	isCurDayDecor := false
	isSell := false
	if obj != nil {
		info := obj.GetActivityData().(*feedbackhouseinvesttypes.FeedbackHouseInvestInfo)
		chargeNum = info.ChargeNum
		curDayChargeNum = info.CurDayChargeNum
		decorDays = info.DecorDays
		isActivity = info.IsActivity
		isCurDayDecor = info.IsCurDayDecor
		isSell = info.IsSell
	}

	scOpenActivityFeedbackHouseInvestInfo.GroupId = &groupId
	scOpenActivityFeedbackHouseInvestInfo.ChargeNum = &chargeNum
	scOpenActivityFeedbackHouseInvestInfo.CurDayChargeNum = &curDayChargeNum
	scOpenActivityFeedbackHouseInvestInfo.DecorDays = &decorDays
	scOpenActivityFeedbackHouseInvestInfo.IsActivity = &isActivity
	scOpenActivityFeedbackHouseInvestInfo.IsCurDayDecor = &isCurDayDecor
	scOpenActivityFeedbackHouseInvestInfo.IsSell = &isSell
	scOpenActivityFeedbackHouseInvestInfo.StartTime = &startTime
	scOpenActivityFeedbackHouseInvestInfo.EndTime = &endTime

	return scOpenActivityFeedbackHouseInvestInfo
}

func BuildSCOpenActivityRankMountList(page int32, mountList []*rankentity.PlayerOrderData, rankTime, startTime, endTime int64) *uipb.SCOpenActivityRankMountList {
	scOpenActivityRankMountList := &uipb.SCOpenActivityRankMountList{}
	scOpenActivityRankMountList.Page = &page
	for _, mountObj := range mountList {
		scOpenActivityRankMountList.MountList = append(scOpenActivityRankMountList.MountList, buildOrder(mountObj))
	}
	scOpenActivityRankMountList.RankTime = &rankTime
	scOpenActivityRankMountList.StartTime = &startTime
	scOpenActivityRankMountList.EndTime = &endTime
	return scOpenActivityRankMountList
}

func BuildSCOpenActivityRankWingList(page int32, wingList []*rankentity.PlayerOrderData, rankTime, startTime, endTime int64) *uipb.SCOpenActivityRankWingList {
	scOpenActivityRankWingList := &uipb.SCOpenActivityRankWingList{}
	scOpenActivityRankWingList.Page = &page
	for _, wingObj := range wingList {
		scOpenActivityRankWingList.WingList = append(scOpenActivityRankWingList.WingList, buildOrder(wingObj))
	}
	scOpenActivityRankWingList.RankTime = &rankTime
	scOpenActivityRankWingList.StartTime = &startTime
	scOpenActivityRankWingList.EndTime = &endTime
	return scOpenActivityRankWingList
}

func BuildSCOpenActivityRankBodyShieldList(page int32, bodyShieldList []*rankentity.PlayerOrderData, rankTime, startTime, endTime int64) *uipb.SCOpenActivityRankBodyShieldList {
	scOpenActivityRankBodyShieldList := &uipb.SCOpenActivityRankBodyShieldList{}
	scOpenActivityRankBodyShieldList.Page = &page
	for _, bodyShieldObj := range bodyShieldList {
		scOpenActivityRankBodyShieldList.BodyShieldList = append(scOpenActivityRankBodyShieldList.BodyShieldList, buildOrder(bodyShieldObj))
	}
	scOpenActivityRankBodyShieldList.RankTime = &rankTime
	scOpenActivityRankBodyShieldList.StartTime = &startTime
	scOpenActivityRankBodyShieldList.EndTime = &endTime
	return scOpenActivityRankBodyShieldList
}

func BuildSCOpenActivityRankChargeList(page int32, chargeList []*rankentity.PlayerPropertyData, rankTime, startTime, endTime int64, myCharge int32) *uipb.SCOpenActivityRankChargeList {
	scOpenActivityRankChargeList := &uipb.SCOpenActivityRankChargeList{}
	scOpenActivityRankChargeList.Page = &page
	for _, chargeObj := range chargeList {
		scOpenActivityRankChargeList.ChargeList = append(scOpenActivityRankChargeList.ChargeList, buildProperty(chargeObj))
	}
	scOpenActivityRankChargeList.RankTime = &rankTime
	scOpenActivityRankChargeList.StartTime = &startTime
	scOpenActivityRankChargeList.EndTime = &endTime
	scOpenActivityRankChargeList.ChargeGold = &myCharge
	return scOpenActivityRankChargeList
}

func BuildSCOpenActivityRankCostList(page int32, costList []*rankentity.PlayerPropertyData, rankTime, startTime, endTime int64, myCost int64) *uipb.SCOpenActivityRankCostList {
	scOpenActivityRankCostList := &uipb.SCOpenActivityRankCostList{}
	scOpenActivityRankCostList.Page = &page
	for _, costObj := range costList {
		scOpenActivityRankCostList.CostList = append(scOpenActivityRankCostList.CostList, buildProperty(costObj))
	}
	scOpenActivityRankCostList.RankTime = &rankTime
	scOpenActivityRankCostList.StartTime = &startTime
	scOpenActivityRankCostList.EndTime = &endTime
	costGodl := int32(myCost)
	scOpenActivityRankCostList.CostGold = &costGodl
	return scOpenActivityRankCostList
}

func BuildSCOpenActivityRankCharmList(page int32, costList []*rankentity.PlayerPropertyData, rankTime, startTime, endTime int64, curAddNum int32, groupId int32) *uipb.SCOpenActivityRankCharmList {
	scOpenActivityRankCharmList := &uipb.SCOpenActivityRankCharmList{}
	scOpenActivityRankCharmList.Page = &page
	for _, costObj := range costList {
		scOpenActivityRankCharmList.CharmList = append(scOpenActivityRankCharmList.CharmList, buildProperty(costObj))
	}
	scOpenActivityRankCharmList.RankTime = &rankTime
	scOpenActivityRankCharmList.StartTime = &startTime
	scOpenActivityRankCharmList.EndTime = &endTime
	scOpenActivityRankCharmList.CurAddNum = &curAddNum
	scOpenActivityRankCharmList.GroupId = &groupId
	return scOpenActivityRankCharmList
}

func BuildSCOpenActivityRankMarryDevelopList(page int32, costList []*rankentity.PlayerPropertyData, rankTime, startTime, endTime int64, curAddExp int32, groupId int32) *uipb.SCOpenActivityRankMarryDevelopList {
	scOpenActivityRankMarryDevelopList := &uipb.SCOpenActivityRankMarryDevelopList{}
	scOpenActivityRankMarryDevelopList.Page = &page
	for _, costObj := range costList {
		scOpenActivityRankMarryDevelopList.MarryDevelopList = append(scOpenActivityRankMarryDevelopList.MarryDevelopList, buildProperty(costObj))
	}
	scOpenActivityRankMarryDevelopList.RankTime = &rankTime
	scOpenActivityRankMarryDevelopList.StartTime = &startTime
	scOpenActivityRankMarryDevelopList.EndTime = &endTime
	scOpenActivityRankMarryDevelopList.CurAddExp = &curAddExp
	scOpenActivityRankMarryDevelopList.GroupId = &groupId
	return scOpenActivityRankMarryDevelopList
}

func BuildSCOpenActivityRankShenFaList(page int32, shenFaList []*rankentity.PlayerOrderData, rankTime, startTime, endTime int64) *uipb.SCOpenActivityRankShenFaList {
	scOpenActivityRankShenFaList := &uipb.SCOpenActivityRankShenFaList{}
	scOpenActivityRankShenFaList.Page = &page
	for _, shenFaObj := range shenFaList {
		scOpenActivityRankShenFaList.ShenFaList = append(scOpenActivityRankShenFaList.ShenFaList, buildOrder(shenFaObj))
	}
	scOpenActivityRankShenFaList.RankTime = &rankTime
	scOpenActivityRankShenFaList.StartTime = &startTime
	scOpenActivityRankShenFaList.EndTime = &endTime
	return scOpenActivityRankShenFaList
}

func BuildSCOpenActivityRankAnqiList(page int32, anqiList []*rankentity.PlayerOrderData, rankTime, startTime, endTime int64) *uipb.SCOpenActivityRankAnqiList {
	scOpenActivityRankAnqiList := &uipb.SCOpenActivityRankAnqiList{}
	scOpenActivityRankAnqiList.Page = &page
	for _, obj := range anqiList {
		scOpenActivityRankAnqiList.AnqiList = append(scOpenActivityRankAnqiList.AnqiList, buildOrder(obj))
	}
	scOpenActivityRankAnqiList.RankTime = &rankTime
	scOpenActivityRankAnqiList.StartTime = &startTime
	scOpenActivityRankAnqiList.EndTime = &endTime
	return scOpenActivityRankAnqiList
}

func BuildSCOpenActivityRankFaBaoList(page int32, anqiList []*rankentity.PlayerOrderData, rankTime, startTime, endTime int64) *uipb.SCOpenActivityRankFaBaoList {
	scMsg := &uipb.SCOpenActivityRankFaBaoList{}
	scMsg.Page = &page
	for _, obj := range anqiList {
		scMsg.FabaoList = append(scMsg.FabaoList, buildOrder(obj))
	}
	scMsg.RankTime = &rankTime
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	return scMsg
}

func BuildSCOpenActivityRankAdvancedList(page int32, advancedList []*rankentity.PlayerOrderData, groupId, typ, subType int32, rankTime, startTime, endTime int64) *uipb.SCOpenActivityRankAdvancedList {
	scMsg := &uipb.SCOpenActivityRankAdvancedList{}
	scMsg.Page = &page
	for _, obj := range advancedList {
		scMsg.AdvancedList = append(scMsg.AdvancedList, buildOrder(obj))
	}
	scMsg.RankTime = &rankTime
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	scMsg.GroupId = &groupId
	scMsg.Type = &typ
	scMsg.SubType = &subType
	return scMsg
}

func BuildSCOpenActivityRankPropertyList(page int32, advancedList []*rankentity.PlayerPropertyData, groupId, typ, subType int32, rankTime, startTime, endTime int64) *uipb.SCOpenActivityRankPropertyList {
	scMsg := &uipb.SCOpenActivityRankPropertyList{}
	scMsg.Page = &page
	for _, obj := range advancedList {
		scMsg.PropertyList = append(scMsg.PropertyList, buildProperty(obj))
	}
	scMsg.RankTime = &rankTime
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	scMsg.GroupId = &groupId
	scMsg.Type = &typ
	scMsg.SubType = &subType
	return scMsg
}

func BuildSCOpenActivityRankLingYuList(page int32, lingYuList []*rankentity.PlayerOrderData, rankTime, startTime, endTime int64) *uipb.SCOpenActivityRankLingYuList {
	scOpenActivityRankLingYuList := &uipb.SCOpenActivityRankLingYuList{}
	scOpenActivityRankLingYuList.Page = &page
	for _, lingYuObj := range lingYuList {
		scOpenActivityRankLingYuList.LingYuList = append(scOpenActivityRankLingYuList.LingYuList, buildOrder(lingYuObj))
	}
	scOpenActivityRankLingYuList.RankTime = &rankTime
	scOpenActivityRankLingYuList.StartTime = &startTime
	scOpenActivityRankLingYuList.EndTime = &endTime
	return scOpenActivityRankLingYuList
}

func BuildSCOpenActivityRankFeatherList(page int32, featherList []*rankentity.PlayerOrderData, rankTime, startTime, endTime int64) *uipb.SCOpenActivityRankFeatherList {
	scOpenActivityRankFeatherList := &uipb.SCOpenActivityRankFeatherList{}
	scOpenActivityRankFeatherList.Page = &page
	for _, featherObj := range featherList {
		scOpenActivityRankFeatherList.FeahterList = append(scOpenActivityRankFeatherList.FeahterList, buildOrder(featherObj))
	}
	scOpenActivityRankFeatherList.RankTime = &rankTime
	scOpenActivityRankFeatherList.StartTime = &startTime
	scOpenActivityRankFeatherList.EndTime = &endTime
	return scOpenActivityRankFeatherList
}

func BuildSCOpenActivityRankShieldList(page int32, shieldList []*rankentity.PlayerOrderData, rankTime, startTime, endTime int64) *uipb.SCOpenActivityRankShieldList {
	scOpenActivityRankShieldList := &uipb.SCOpenActivityRankShieldList{}
	scOpenActivityRankShieldList.Page = &page
	for _, shieldObj := range shieldList {
		scOpenActivityRankShieldList.ShieldList = append(scOpenActivityRankShieldList.ShieldList, buildOrder(shieldObj))
	}
	scOpenActivityRankShieldList.RankTime = &rankTime
	scOpenActivityRankShieldList.StartTime = &startTime
	scOpenActivityRankShieldList.EndTime = &endTime
	return scOpenActivityRankShieldList
}

func BuildSCOpenActivityRankCountList(page int32, countList []*rankentity.PlayerPropertyData, groupId, useNum int32, rankTime, startTime, endTime int64) *uipb.SCOpenActivityRankCountList {
	scMsg := &uipb.SCOpenActivityRankCountList{}
	scMsg.Page = &page
	for _, countObj := range countList {
		scMsg.CountList = append(scMsg.CountList, buildProperty(countObj))
	}
	scMsg.RankTime = &rankTime
	scMsg.StartTime = &startTime
	scMsg.EndTime = &endTime
	scMsg.GroupId = &groupId
	scMsg.UseTimes = &useNum

	return scMsg
}

func BuildSCOpenActivityRankMyGet(groupId int32, pos int32) *uipb.SCOpenActivityRankMyGet {
	scOpenActivityRankMyGet := &uipb.SCOpenActivityRankMyGet{}
	scOpenActivityRankMyGet.GroupId = &groupId
	scOpenActivityRankMyGet.Pos = &pos
	return scOpenActivityRankMyGet
}

func BuildSCOpenActivityFirstChargeNotice(isFirst, isReceive bool) *uipb.SCOpenActivityFirstChargeNotice {
	scOpenActivityFirstChargeNotice := &uipb.SCOpenActivityFirstChargeNotice{}
	scOpenActivityFirstChargeNotice.IsFirst = &isFirst
	scOpenActivityFirstChargeNotice.IsReceive = &isReceive
	return scOpenActivityFirstChargeNotice
}

func BuildSCOpenActivityWelfareOnlineDataNotice(groupId, drewTimes int32) *uipb.SCOpenActivityWelfareOnlineDataNotice {
	scMsg := &uipb.SCOpenActivityWelfareOnlineDataNotice{}
	scMsg.GroupId = &groupId
	scMsg.OnlineDrewTimes = &drewTimes

	return scMsg
}

func BuildSCOpenActivityAllianceCheerPoolChanged(groupId int32, poolNum int64) *uipb.SCOpenActivityAllianceCheerPoolChanged {
	scMsg := &uipb.SCOpenActivityAllianceCheerPoolChanged{}
	scMsg.GroupId = &groupId
	scMsg.PoolNum = &poolNum

	return scMsg
}

func BuildSCOpenActivityOpenServerTimeChanged(openTime, mergeTime int64) *uipb.SCOpenActivityOpenServerTimeChanged {
	scMsg := &uipb.SCOpenActivityOpenServerTimeChanged{}
	scMsg.ActvityOPenServerTime = &openTime
	scMsg.ActvityMergeTime = &mergeTime

	return scMsg
}

func BuildSCOpenActivityXunHuanInfo(groupIdList []int32) *uipb.SCOpenActivityXunHuanInfo {
	scMsg := &uipb.SCOpenActivityXunHuanInfo{}
	scMsg.InfoList = buildXunHuanInfo(groupIdList)
	return scMsg
}

func buildXunHuanInfo(groupIdList []int32) (infoList []*uipb.XunHuanInfo) {
	for _, groupId := range groupIdList {
		newGroupId := groupId

		star, end := welfare.GetWelfareService().CountOpenActivityTime(groupId)
		info := &uipb.XunHuanInfo{}
		info.GroupId = &newGroupId
		info.StartTime = &star
		info.EndTime = &end

		infoList = append(infoList, info)
	}

	return
}

func buildWelfareLoginInfo(day int32, obj *playerwelfare.PlayerOpenActivityObject) *uipb.WelfareLoginInfo {
	var record []int32
	groupId := int32(0)
	if obj != nil {
		record = obj.GetActivityData().(*halllogintypes.WelfareLoginInfo).RewRecord
		groupId = obj.GetGroupId()
	}

	info := &uipb.WelfareLoginInfo{}
	info.DayNum = &day
	info.LoginRecordList = record
	info.GroupId = &groupId

	return info
}

func buildWelfareUplevelInfo(obj *playerwelfare.PlayerOpenActivityObject, groupId, level int32, startTime, endTime int64) *uipb.WelfareUplevelInfo {
	var record []int32
	if obj != nil {
		record = obj.GetActivityData().(*hallupleveltypes.WelfareUplevelInfo).RewRecord
	}
	info := &uipb.WelfareUplevelInfo{}
	info.Level = &level
	info.UplevelRecordList = record
	info.GroupId = &groupId
	info.StartTime = &startTime
	info.EndTime = &endTime

	return info
}
func buildWelfareOnlineInfo(onlineTime int64, obj *playerwelfare.PlayerOpenActivityObject) *uipb.WelfareOnlineInfo {
	var record []int32
	drewTimes := int32(0)
	groupId := int32(0)
	if obj != nil {
		online := obj.GetActivityData().(*hallonlinetypes.WelfareOnlineInfo)
		record = online.RewRecord
		drewTimes = online.DrawTimes
		groupId = obj.GetGroupId()
	}

	info := &uipb.WelfareOnlineInfo{}
	info.OnlineTime = &onlineTime
	info.OnlineRecordList = record
	info.DrewTimes = &drewTimes
	info.GroupId = &groupId

	return info
}

func buildInvestLevelInfo(record int32, typ investleveltypes.InvestLevelType, groupId int32) *uipb.InvestLevelInfo {
	info := &uipb.InvestLevelInfo{}
	typInt := int32(typ)
	info.Typ = &typInt
	info.InvestLevelRecord = &record
	info.GroupId = &groupId

	return info
}

func buildInvestNewLevelInfo(records []int32, typ investnewleveltypes.InvestNewLevelType, groupId int32) *uipb.InvestNewLevelInfo {
	info := &uipb.InvestNewLevelInfo{}
	typInt := int32(typ)
	info.Typ = &typInt
	for _, record := range records {
		info.InvestLevelRecord = append(info.InvestLevelRecord, record)
	}
	info.GroupId = &groupId

	return info
}

func buildInvestDayInfo(obj *playerwelfare.PlayerOpenActivityObject) *uipb.InvestDayInfo {
	record := int32(0)
	groupId := int32(0)
	buyTime := int64(0)
	isBuy := false
	if obj != nil {
		info := obj.GetActivityData().(*investservendaytypes.InvestDayInfo)
		record = info.ReceiveDay
		groupId = obj.GetGroupId()
		buyTime = info.BuyTime
	}
	if buyTime > 0 {
		isBuy = true
	}

	info := &uipb.InvestDayInfo{}
	info.IsBuyInvestDay = &isBuy
	info.InvestDayRecord = &record
	info.GroupId = &groupId
	info.BuyTime = &buyTime

	return info
}

func buildRewProperty(rd *propertytypes.RewData) *uipb.RewProperty {
	rewProperty := &uipb.RewProperty{}
	rewExp := rd.GetRewExp()
	rewExpPoint := rd.GetRewExpPoint()
	rewGold := rd.GetRewGold()
	rewBindGold := rd.GetRewBindGold()
	rewSilver := rd.GetRewSilver()

	rewProperty.Exp = &rewExp
	rewProperty.ExpPoint = &rewExpPoint
	rewProperty.Silver = &rewSilver
	rewProperty.Gold = &rewGold
	rewProperty.BindGold = &rewBindGold

	return rewProperty
}

func buildDropInfoList(dropItemDataList []*droptemplate.DropItemData) []*uipb.DropInfo {
	dropInfoList := make([]*uipb.DropInfo, 0, len(dropItemDataList))
	for _, itemData := range dropItemDataList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()
		level := itemData.GetLevel()

		dropInfoList = append(dropInfoList, buildDropInfo(itemId, num, level))
	}

	return dropInfoList
}

func buildDropInfo(itemId, num, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level

	return dropInfo
}

func buildOrder(orderData *rankentity.PlayerOrderData) *uipb.RankOrder {
	rankOder := &uipb.RankOrder{}
	id := orderData.PlayerId
	name := orderData.PlayerName
	order := orderData.Order
	power := orderData.Power
	rankOder.PlayerId = &id
	rankOder.PlayerName = &name
	rankOder.Order = &order
	rankOder.Power = &power
	return rankOder
}

func buildProperty(data *rankentity.PlayerPropertyData) *uipb.RankProperty {
	rankProperty := &uipb.RankProperty{}
	id := data.PlayerId
	name := data.PlayerName
	num := data.Num
	power := data.Power
	rankProperty.PlayerId = &id
	rankProperty.PlayerName = &name
	rankProperty.Num = &num
	rankProperty.Power = &power
	return rankProperty
}

func buildDiscountInfo(index, buyTimes int32) *uipb.DiscountBuyInfo {
	discountBuyInfo := &uipb.DiscountBuyInfo{}
	discountBuyInfo.Index = &index
	discountBuyInfo.BuyTimes = &buyTimes
	return discountBuyInfo
}

func buildDiscountGlobalTimesInfo(index, num int32) *uipb.DiscountGlobalTimesInfo {
	info := &uipb.DiscountGlobalTimesInfo{}
	info.Index = &index
	info.GlobalTimes = &num
	return info
}

func buildRewardsTimesInfo(level, times int32) *uipb.RewardsTimesInfo {
	info := &uipb.RewardsTimesInfo{}
	info.Level = &level
	info.ReceiveTimes = &times
	return info
}

func buildBuyTimesInfoList(record map[int32]int32) (infoList []*uipb.BuyTimesInfo) {
	for key, val := range record {
		typ := key
		times := val

		info := &uipb.BuyTimesInfo{}
		info.Type = &typ
		info.BuyTimes = &times

		infoList = append(infoList, info)
	}
	return infoList
}

func buildLaBaLog(obj *welfare.GoldLaBaLogObject) *uipb.LaBaLog {
	log := &uipb.LaBaLog{}
	playerName := obj.GetPlayerName()
	costGold := obj.GetCostGold()
	rewGold := obj.GetRewGold()
	time := obj.GetUpdateTime()

	log.PlayerName = &playerName
	log.RewGold = &rewGold
	log.CostGold = &costGold
	log.CreateTime = &time

	return log
}

func buildDrewLog(obj *welfare.DrewLogObject) *uipb.DrewLog {
	log := &uipb.DrewLog{}
	playerName := obj.GetPlayerName()
	itemId := obj.GetItemId()
	itemNum := obj.GetItemNum()
	time := obj.GetUpdateTime()

	log.PlayerName = &playerName
	log.ItemId = &itemId
	log.ItemNum = &itemNum
	log.CreateTime = &time

	return log
}

func buildCrazyBoxLog(obj *welfare.CrazyBoxLogObject) *uipb.CrazyBoxLog {
	log := &uipb.CrazyBoxLog{}
	playerName := obj.GetPlayerName()
	itemId := obj.GetItemId()
	itemNum := obj.GetItemNum()
	time := obj.GetUpdateTime()

	log.PlayerName = &playerName
	log.ItemId = &itemId
	log.ItemNum = &itemNum
	log.CreateTime = &time

	return log
}

func buildTimesRewRecordList(record map[int32][]int32) (scRecord []*uipb.TimesRewRecord) {
	for vip, timesList := range record {
		vipInt := vip
		for _, times := range timesList {
			timesInt := times
			info := &uipb.TimesRewRecord{}
			info.Times = &timesInt
			info.Vip = &vipInt
			scRecord = append(scRecord, info)
		}
	}

	return scRecord
}
