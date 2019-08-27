package handler

// import (
// 	uipb "fgame/fgame/common/codec/pb/ui"
// 	"fgame/fgame/common/lang"
// 	commonlog "fgame/fgame/common/log"
// 	"fgame/fgame/core/session"
// 	"fgame/fgame/game/constant/constant"
// 	droplogic "fgame/fgame/game/drop/logic"
// 	"fgame/fgame/game/global"
// 	inventorylogic "fgame/fgame/game/inventory/logic"
// 	playerinventory "fgame/fgame/game/inventory/player"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	"fgame/fgame/game/player/types"
// 	propertylogic "fgame/fgame/game/property/logic"
// 	playerproperty "fgame/fgame/game/property/player"
// 	propertytypes "fgame/fgame/game/property/types"
// 	xianfulogic "fgame/fgame/game/xianfu/logic"
// 	"fgame/fgame/game/xianfu/pbutil"
// 	xianfutemplate "fgame/fgame/game/xianfu/template"

// 	gamesession "fgame/fgame/game/session"
// 	xianfuplayer "fgame/fgame/game/xianfu/player"
// 	xianfutypes "fgame/fgame/game/xianfu/types"
// 	"fmt"

// 	log "github.com/Sirupsen/logrus"
// )

// func init() {
// 	// processor.Register(codec.MessageType(uipb.MessageType_CS_XIANFU_FINISH_ALL_TYPE), dispatch.HandlerFunc(handlerMaterialFinishAll))
// }

// //材料副本一键完成请求
// func handlerMaterialFinishAll(s session.Session, msg interface{}) (err error) {
// 	log.Debug("xianfu:处理材料副本一键完成请求")

// 	pl := gamesession.SessionInContext(s.Context()).Player()
// 	tpl := pl.(player.Player)
// 	csMaterialFinishAll := msg.(*uipb.CSMaterialFinishAll)
// 	typ := csMaterialFinishAll.GetMaterialType()
// 	//验证参数
// 	xianfuType := xianfutypes.MaterialType(typ)
// 	if !xianfuType.Valid() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":   pl.GetId(),
// 				"xianfuType": xianfuType,
// 			}).Warn("xianfu:材料副本一键完成请求，参数错误")
// 		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
// 		return
// 	}

// 	err = xianfuFinishAll(tpl, xianfuType)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":   tpl.GetId(),
// 				"xianfuType": xianfuType,
// 				"err":        err,
// 			}).Error("xianfu:处理材料副本一键完成请求，错误")
// 		return
// 	}

// 	log.WithFields(
// 		log.Fields{
// 			"playerId":   tpl.GetId(),
// 			"xianfuType": xianfuType,
// 		}).Debug("xianfu：处理材料副本一键完成请求完成")

// 	return
// }

// //仙府一键完成逻辑
// func xianfuFinishAll(pl player.Player, xianfuType xianfutypes.MaterialType) (err error) {
// 	xianfuManager := pl.GetPlayerDataManager(types.PlayerMaterialDtatManagerType).(*xianfuplayer.PlayerXinafuDataManager)
// 	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
// 	xianfuId := xianfuManager.GetMaterialId(xianfuType)

// 	xfTemplate := xianfutemplate.GetMaterialTemplateService().GetMaterial(xianfuId, xianfuType)
// 	if xfTemplate == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":   pl.GetId(),
// 				"xianfuId":   xianfuId,
// 				"xianfuType": xianfuType,
// 			}).Warn("xianfu:材料副本一键完成请求，模板不存在")

// 		playerlogic.SendSystemMessage(pl, lang.MaterialArgumentInvalid)
// 		return
// 	}

// 	//刷新数据
// 	now := global.GetGame().GetTimeService().Now()
// 	err = xianfuManager.RefreshData(now)
// 	if err != nil {
// 		return
// 	}

// 	leftNum := xianfuManager.GetChallengeTimes(xianfuType)
// 	costGold := constant.GetConstantService().GetConstant(xianfuType.GetFinishAllConstantType())
// 	finishAllNeedGold := costGold * leftNum

// 	//判断元宝是否足够
// 	if !propertyManager.HasEnoughGold(int64(finishAllNeedGold), false) {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":   pl.GetId(),
// 				"xianfuType": xianfuType,
// 			}).Warn("xianfu:材料副本一键完成请求，当前元宝不足，无法升级")
// 		playerlogic.SendSystemMessage(pl, lang.MaterialNotEnoughGold)
// 		return
// 	}

// 	showItemList, rewardsItemList, rewardsResMap := xianfulogic.GetSaoDangDrop(leftNum, xianfuId, xianfuType)
// 	//背包不足
// 	if len(rewardsItemList) > 0 {
// 		if !inventoryManager.HasEnoughSlotsOfItemLevel(rewardsItemList) {
// 			log.WithFields(
// 				log.Fields{
// 					"playerId":   pl.GetId(),
// 					"xianfuType": xianfuType,
// 				}).Warn("xianfu:材料副本一键完成请求，背包空间不足")
// 			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
// 			return
// 		}
// 	}

// 	//消耗元宝
// 	goldReason := commonlog.GoldLogReasonMaterialFinishAllUse
// 	goldReasonText := fmt.Sprintf(goldReason.String(), xianfuId, xianfuType)
// 	flag := propertyManager.CostGold(int64(finishAllNeedGold), true, goldReason, goldReasonText)
// 	if !flag {
// 		panic(fmt.Errorf("xianfu: finish all use gold should be ok"))
// 	}

// 	//增加物品
// 	itemReason := commonlog.InventoryLogReasonMaterialSaodangRewards
// 	rewardItemReasonText := fmt.Sprintf(itemReason.String(), xianfuId, xianfuType)
// 	flag = inventoryManager.BatchAddOfItemLevel(rewardsItemList, itemReason, rewardItemReasonText)
// 	if !flag {
// 		panic("xianfu:finish all add item should be ok")
// 	}

// 	//获取一键完成固定资源
// 	reasonGold := commonlog.GoldLogReasonMaterialFinishAllRew
// 	reasonSilver := commonlog.SilverLogReasonMaterialFinishAllRew
// 	reasonLevel := commonlog.LevelLogReasonMaterialFinishAllRew
// 	finishExpReasonText := fmt.Sprintf(reasonLevel.String(), xianfuId, xianfuType)
// 	finishGoldReasonText := fmt.Sprintf(reasonGold.String(), xianfuId, xianfuType)
// 	finishSilverReasonText := fmt.Sprintf(reasonSilver.String(), xianfuId, xianfuType)

// 	rewSilver := int32(xfTemplate.GetRawSilver()) * leftNum
// 	rewBindGold := xfTemplate.GetRawBindGold() * leftNum
// 	rewGold := xfTemplate.GetRawGold() * leftNum
// 	rewExp := int32(xfTemplate.GetRawExp()) * leftNum
// 	rewExpPoint := int32(xfTemplate.GetRawExpPoint()) * leftNum
// 	totalRewData := propertytypes.CreateRewData(rewExp, rewExpPoint, rewSilver, rewGold, rewBindGold)

// 	flag = propertyManager.AddRewData(totalRewData, reasonGold, finishGoldReasonText, reasonSilver, finishSilverReasonText, reasonLevel, finishExpReasonText)
// 	if !flag {
// 		panic("xianfu:finish all add RewData should be ok")
// 	}

// 	//增加资源
// 	if len(rewardsResMap) > 0 {
// 		err = droplogic.AddRes(pl, rewardsResMap, reasonGold, finishGoldReasonText, reasonSilver, finishSilverReasonText, reasonLevel, finishExpReasonText)
// 		if err != nil {
// 			return
// 		}
// 	}

// 	//完成一键完成
// 	xianfuManager.UseTimes(xianfuType, leftNum, now)
// 	xianfuManager.EmitFinishEvent(xianfuType, leftNum)

// 	//同步资源
// 	propertylogic.SnapChangedProperty(pl)

// 	//同步背包
// 	inventorylogic.SnapInventoryChanged(pl)

// 	scMaterialSaoDang := pbutil.BuildSCMaterialFinishAll(xianfuId, xianfuType, totalRewData, showItemList)
// 	pl.SendMsg(scMaterialSaoDang)
// 	return
// }
