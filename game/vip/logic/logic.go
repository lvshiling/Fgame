package logic

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/vip/pbutil"
	playervip "fgame/fgame/game/vip/player"
	viptemplate "fgame/fgame/game/vip/template"
	viptypes "fgame/fgame/game/vip/types"
	"math"
	"sort"

	log "github.com/Sirupsen/logrus"
)

// 推送vip信息
func VipInfoNotice(pl player.Player) {
	vipManager := pl.GetPlayerDataManager(playertypes.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)
	level, star := vipManager.GetVipLevel()
	giftRecord, freeGiftRecord := vipManager.GetGiftRecord()
	chargeNum := vipManager.GetChargeNum()
	scMsg := pbutil.BuildSCVipInfoNotice(level, star, chargeNum, giftRecord, freeGiftRecord)
	pl.SendMsg(scMsg)
}

// 付费等级计算系统升阶规则
func CountWithCostLevel(pl player.Player, moduleType viptypes.CostLevelRuleType, min, max int32) (minTimes int32, maxTimes int32) {
	vipManager := pl.GetPlayerDataManager(playertypes.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)
	costLevel := vipManager.GetCostLevel()
	costTemplate := viptemplate.GetVipTemplateService().GetCostLevelTemplate(costLevel)
	if costTemplate == nil {
		log.WithFields(
			log.Fields{
				"consumerLevel": costLevel,
			}).Warn("付费模板不存在")
		minTimes = min
		maxTimes = max
		return
	}

	ruleData, ok := costTemplate.GetAdvancedRuleMap()[moduleType]
	if !ok {
		log.WithFields(
			log.Fields{
				"moduleType": moduleType,
			}).Warn("该系统没有付费规则")
		minTimes = min
		maxTimes = max
		return
	}

	minRate := ruleData.GetMinRate()
	maxRate := ruleData.GetMaxRate()
	minTimes = int32(math.Ceil(float64(min) * float64(minRate) / float64(common.MAX_RATE)))
	maxTimes = int32(math.Ceil(float64(max) * float64(maxRate) / float64(common.MAX_RATE)))

	log.WithFields(
		log.Fields{
			"min":        minTimes,
			"max":        maxTimes,
			"costLevel":  costLevel,
			"playerId":   pl.GetId(),
			"moduleType": moduleType,
		}).Debug("付费规则计算进阶次数")
	return
}

// 付费等级计算固定次数掉落包
func CountDropTimesWithCostLevel(pl player.Player, moduleType viptypes.CostLevelRuleType, timesDescList []int) (ruleTimesMap map[int]int32) {
	vipManager := pl.GetPlayerDataManager(playertypes.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)
	costLevel := vipManager.GetCostLevel()
	costTemplate := viptemplate.GetVipTemplateService().GetCostLevelTemplate(costLevel)
	if costTemplate == nil {
		log.WithFields(
			log.Fields{
				"consumerLevel": costLevel,
			}).Warn("付费模板不存在")
		return
	}

	ruleTimesList, ok := costTemplate.GetDropRuleMap()[moduleType]
	if !ok {
		log.WithFields(
			log.Fields{
				"moduleType": moduleType,
			}).Warn("该系统没有付费规则")
		return
	}

	if len(ruleTimesList) != len(timesDescList) {
		log.WithFields(
			log.Fields{
				"moduleType": moduleType,
			}).Warn("付费掉落规则和基础掉落规则数不一致")
		return
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ruleTimesList)))

	// 计算固定掉落次数
	ruleTimesMap = make(map[int]int32)
	for index, times := range timesDescList {
		newTimes := int32(math.Ceil(float64(times) * float64(ruleTimesList[index]) / float64(common.MAX_RATE)))
		ruleTimesMap[times] = newTimes
	}

	log.WithFields(
		log.Fields{
			"costLevel":     costLevel,
			"playerId":      pl.GetId(),
			"moduleType":    moduleType,
			"timesDescList": timesDescList,
			"ruleTimesMap":  ruleTimesMap,
		}).Debug("付费规则计算掉落次数")

	return
}

// 幻境boss疲劳值最大购买数
func GetUnrealBossMaxBuyTimes(pl player.Player) int32 {
	vipManager := pl.GetPlayerDataManager(playertypes.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)
	leve, star := vipManager.GetVipLevel()
	vipTemp := viptemplate.GetVipTemplateService().GetVipTemplate(leve, star)
	if vipTemp == nil {
		return 0
	}
	return vipTemp.BuyPilaoCount
}

// 红包额外领取数量
func GetHongBaoSnatchCount(pl player.Player) int32 {
	vipManager := pl.GetPlayerDataManager(playertypes.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)
	leve, star := vipManager.GetVipLevel()
	vipTemp := viptemplate.GetVipTemplateService().GetVipTemplate(leve, star)
	if vipTemp == nil {
		return 0
	}
	return vipTemp.HbSnatchCount
}
