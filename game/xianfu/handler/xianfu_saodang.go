package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/constant/constant"
	droplogic "fgame/fgame/game/drop/logic"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	xianfulogic "fgame/fgame/game/xianfu/logic"
	"fgame/fgame/game/xianfu/pbutil"
	xianfutypes "fgame/fgame/game/xianfu/types"

	gamesession "fgame/fgame/game/session"
	xianfuplayer "fgame/fgame/game/xianfu/player"
	xianfutemplate "fgame/fgame/game/xianfu/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIANFU_SAODANG_TYPE), dispatch.HandlerFunc(handlerXianfuSaodang))
}

//秘境仙府扫荡请求
func handlerXianfuSaodang(s session.Session, msg interface{}) (err error) {
	log.Debug("xianfu:处理秘境仙府扫荡请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csXianfuSaoDang := msg.(*uipb.CSXianfuSaoDang)
	typ := csXianfuSaoDang.GetXianfuType()
	num := csXianfuSaoDang.GetNum()

	xianfuType := xianfutypes.XianfuType(typ)
	//验证参数
	if !xianfuType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuType": xianfuType,
			}).Warn("xianfu:秘境仙府扫荡请求，参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = xianfuSaodang(tpl, xianfuType, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   tpl.GetId(),
				"xianfuType": xianfuType,
				"err":        err,
			}).Error("xianfu:处理秘境仙府扫荡请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   tpl.GetId(),
			"xianfuType": xianfuType,
		}).Debug("xianfu：处理秘境仙府扫荡请求完成")

	return
}

//仙府扫荡逻辑
func xianfuSaodang(pl player.Player, xianfuType xianfutypes.XianfuType, saoDangNum int32) (err error) {
	// 等级限制
	limitLevel := constant.GetConstantService().GetConstant(xianfuType.GetSaoDangNeedLevelConstantType())
	if pl.GetLevel() < limitLevel {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuType": xianfuType,
				"needLevel":  limitLevel,
			}).Warn("xianfu:秘境仙府扫荡请求，扫荡所需等级不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	xianfuManager := pl.GetPlayerDataManager(types.PlayerXianfuDtatManagerType).(*xianfuplayer.PlayerXinafuDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	xianfuId := xianfuManager.GetXianfuId(xianfuType)
	xfTemplate := xianfutemplate.GetXianfuTemplateService().GetXianfu(xianfuId, xianfuType)
	if xfTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuId":   xianfuId,
				"xianfuType": xianfuType,
			}).Warn("xianfu:秘境仙府扫荡请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.XianfuArgumentInvalid)
		return
	}

	//波数限制
	if xianfuType == xianfutypes.XianfuTypeExp {
		if xianfuManager.GetGroup(xianfuType) < xfTemplate.GetGroupLimit() {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("xianfu:秘境仙府扫荡请求，经验副本波数0不能扫荡")
			playerlogic.SendSystemMessage(pl, lang.XianfuArgumentInvalid)
			return
		}
	}

	//刷新数据
	now := global.GetGame().GetTimeService().Now()
	err = xianfuManager.RefreshData(now)
	if err != nil {
		return
	}

	//挑战次数是否足够
	leftTimes := xianfuManager.GetChallengeTimes(xianfuType)
	if leftTimes < saoDangNum {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuId":   xianfuId,
				"xianfuType": xianfuType,
			}).Warn("xianfu:秘境仙府扫荡请求，副本次数不足，无法扫荡")
		playerlogic.SendSystemMessage(pl, lang.XinafuNotEnoughChallengeTimes)
		return
	}

	showItemList, rewardsItemList, rewardsResMap := xianfulogic.GetSaoDangDrop(saoDangNum, xianfuId, xianfuType)
	if !inventoryManager.HasEnoughSlotsOfItemLevel(rewardsItemList) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"xianfuId":   xianfuId,
				"xianfuType": xianfuType,
			}).Warn("xianfu:秘境仙府扫荡请求，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	attendNeedItemId := xfTemplate.GetNeedItemId()
	attendNeedItemNum := int32(0)
	saodangNeedGold := int32(0)
	freeTime := xianfulogic.FreeTimesCount(pl, xianfuType)
	if saoDangNum > freeTime {
		attendNeedItemNum = xfTemplate.GetNeedItemCount() * (saoDangNum - freeTime)
		saodangNeedGold = xfTemplate.GetSaodangNeedGold() * (saoDangNum - freeTime)
	}
	//挑战所需物品是否足够
	if attendNeedItemNum > 0 {
		if !inventoryManager.HasEnoughItem(attendNeedItemId, attendNeedItemNum) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("xianfu:秘境仙府扫荡请求，副本挑战令不足，无法扫荡")
			playerlogic.SendSystemMessage(pl, lang.XinafuNotEnoughChallengeItem)
			return
		}
	}

	//预留字段校验：扫荡所需元宝是否足够
	if saodangNeedGold > 0 {
		if !propertyManager.HasEnoughGold(int64(saodangNeedGold), true) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("xianfu:秘境仙府扫荡请求，当前元宝不足，无法扫荡")
			playerlogic.SendSystemMessage(pl, lang.XianfuNotEnoughGold)
			return
		}
	}

	//扫荡所需物品是否足够
	saodangNeedItemMap := xfTemplate.GetSaodangItemMap(saoDangNum)
	if len(saodangNeedItemMap) > 0 {
		if !inventoryManager.HasEnoughItems(saodangNeedItemMap) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"xianfuId":   xianfuId,
					"xianfuType": xianfuType,
				}).Warn("xianfu:秘境仙府扫荡请求，当前扫荡券不足，无法扫荡")
			playerlogic.SendSystemMessage(pl, lang.XinafuNotEnoughSaodangItem)
			return
		}
	}

	//扣除扫荡物品
	useItemReason := commonlog.InventoryLogReasonXianfuSaodang
	useItemReasonText := fmt.Sprintf(useItemReason.String(), xianfuId, xianfuType)
	if len(saodangNeedItemMap) > 0 {
		if flag := inventoryManager.BatchRemove(saodangNeedItemMap, useItemReason, useItemReasonText); !flag {
			panic("xianfu: xianfuSaodang use item should be ok")
		}
	}

	//扣除挑战物品
	if attendNeedItemNum > 0 {
		if flag := inventoryManager.UseItem(attendNeedItemId, attendNeedItemNum, useItemReason, useItemReasonText); !flag {
			panic("xianfu: xianfuSaodang use item should be ok")
		}
	}
	//消耗元宝
	if saodangNeedGold > 0 {
		goldReason := commonlog.GoldLogReasonXianfuSaodang
		goldReasonText := fmt.Sprintf(goldReason.String(), xianfuId, xianfuType)
		flag := propertyManager.CostGold(int64(saodangNeedGold), true, goldReason, goldReasonText)
		if !flag {
			panic(fmt.Errorf("xianfu: xianfuSaodang use gold should be ok"))
		}
	}

	//增加物品
	rewardItemReason := commonlog.InventoryLogReasonXianfuSaodangRewards
	rewardItemReasonText := fmt.Sprintf(rewardItemReason.String(), xianfuId, xianfuType)
	flag := inventoryManager.BatchAddOfItemLevel(rewardsItemList, rewardItemReason, rewardItemReasonText)
	if !flag {
		panic("xianfu:xianfuSaodang add item should be ok")
	}

	//获取扫荡固定资源
	reasonGold := commonlog.GoldLogReasonXianfuSaodangRewards
	reasonSilver := commonlog.SilverLogReasonXianfuSaodangRewards
	reasonLevel := commonlog.LevelLogReasonXianfuSaodangRewards
	saodangGoldReasonText := fmt.Sprintf(reasonGold.String(), xianfuId, xianfuType)
	saodangSilverReasonText := fmt.Sprintf(reasonSilver.String(), xianfuId, xianfuType)
	expReasonText := fmt.Sprintf(reasonLevel.String(), xianfuId, xianfuType)

	rewSilver := int32(xfTemplate.GetRawSilver()) * saoDangNum
	rewBindGold := xfTemplate.GetRawBindGold() * saoDangNum
	rewGold := xfTemplate.GetRawGold() * saoDangNum
	rewExp := int32(xfTemplate.GetRawExp()) * saoDangNum

	rewExpPoint := int32(xfTemplate.GetRawExpPoint()) * saoDangNum
	if xianfuType == xianfutypes.XianfuTypeExp {
		group := xianfuManager.GetGroup(xianfuType)
		rewExp *= group
		rewExpPoint *= group
	}

	totalRewData := propertytypes.CreateRewData(rewExp, rewExpPoint, rewSilver, rewGold, rewBindGold)

	flag = propertyManager.AddRewData(totalRewData, reasonGold, saodangGoldReasonText, reasonSilver, saodangSilverReasonText, reasonLevel, expReasonText)
	if !flag {
		panic("xianfu:xianfuSaodang add RewData should be ok")
	}

	//增加资源
	if len(rewardsResMap) > 0 {
		err = droplogic.AddRes(pl, rewardsResMap, reasonGold, saodangGoldReasonText, reasonSilver, saodangSilverReasonText, reasonLevel, expReasonText)
		if err != nil {
			return
		}
	}

	//完成扫荡
	xianfuManager.UseTimes(xianfuType, saoDangNum, now)
	xianfuManager.EmitFinishEvent(xianfuType, saoDangNum)
	xianfuManager.EmitSweepEvent(xianfuType, saoDangNum)

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	//同步背包
	inventorylogic.SnapInventoryChanged(pl)

	scXianfuSaoDang := pbutil.BuildSCXianfuSaoDang(xianfuId, xianfuType, saoDangNum, totalRewData, showItemList)
	pl.SendMsg(scXianfuSaoDang)
	return
}
