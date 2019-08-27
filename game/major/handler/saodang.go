package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	droplogic "fgame/fgame/game/drop/logic"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	majoreventtypes "fgame/fgame/game/major/event/types"
	majorlogic "fgame/fgame/game/major/logic"
	"fgame/fgame/game/major/pbutil"
	playermajor "fgame/fgame/game/major/player"
	majortemplate "fgame/fgame/game/major/template"
	majortypes "fgame/fgame/game/major/types"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	weeklogic "fgame/fgame/game/week/logic"

	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MAJOR_SAODANG_TYPE), dispatch.HandlerFunc(handlerMajorSaodang))
}

//副本扫荡请求
func handlerMajorSaodang(s session.Session, msg interface{}) (err error) {
	log.Debug("major:处理副本扫荡请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSMajorSaoDang)
	typ := csMsg.GetMajorType()
	fuBenId := csMsg.GetFubenId()
	num := csMsg.GetNum()

	majorType := majortypes.MajorType(typ)
	//验证参数
	if !majorType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"majorType": majorType,
			}).Warn("major:参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	if num <= 0 {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"majorType": majorType,
			}).Warn("major:参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	if fuBenId <= 0 {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"majorType": majorType,
			}).Warn("major:参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = majorSaodang(tpl, majorType, fuBenId, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("major:处理双修扫荡消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("major:处理双修扫荡消息完成")
	return
}

//副本扫荡逻辑
func majorSaodang(pl player.Player, majorType majortypes.MajorType, fubenId int32, saoDangNum int32) (err error) {
	majorTemp := majortemplate.GetMajorTemplateService().GetMajorTemplate(majorType, fubenId)
	if majorTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"majorType": majorType,
			}).Warn("major:夫妻副本扫荡请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMajorDataManagerType).(*playermajor.PlayerMajorDataManager)
	flag := manager.HasMajorNumByNum(majorType, saoDangNum)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("major:今日双修次数已用完")
		playerlogic.SendSystemMessage(pl, lang.MajorInviteNoTimes)
		return
	}

	marryManager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := marryManager.GetMarryInfo()
	if marryInfo.SpouseId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("major:前没有配偶,结婚后可邀请配偶双修")
		playerlogic.SendSystemMessage(pl, lang.MajorInviteNoSpouse)
		return
	}

	if marryInfo.Status != marrytypes.MarryStatusTypeMarried {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("major:举办婚礼成为正式夫妻才可进入")
		playerlogic.SendSystemMessage(pl, lang.MajorNoHoldWed)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	saodangNeedGold := int64(saoDangNum * majorTemp.GetSaodangNeedGold())
	saodangNeedItemMap := majorTemp.GetSaodangItemMap(saoDangNum)
	if weeklogic.IsSeniorWeek(pl) {
		saodangNeedGold = int64(0)
		saodangNeedItemMap = nil
	}

	//扫荡所需物品是否足够
	if len(saodangNeedItemMap) > 0 {
		if !inventoryManager.HasEnoughItems(saodangNeedItemMap) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"majorType":  majorType,
					"saoDangNum": saoDangNum,
				}).Warn("major:夫妻副本扫荡请求，当前扫荡券不足，无法扫荡")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//预留字段校验：扫荡所需元宝是否足够
	if saodangNeedGold > 0 {
		if !propertyManager.HasEnoughGold(int64(saodangNeedGold), true) {
			log.WithFields(
				log.Fields{
					"playerId":        pl.GetId(),
					"majorType":       majorType,
					"saoDangNum":      saoDangNum,
					"saodangNeedGold": saodangNeedGold,
				}).Warn("major:夫妻副本扫荡请求，当前元宝不足，无法扫荡")
			playerlogic.SendSystemMessage(pl, lang.XianfuNotEnoughGold)
			return
		}
	}

	showItemList, rewardsItemList, rewardsResMap, totalRewData, flagSaoDang := majorlogic.GetSaoDangDrop(pl, saoDangNum, majorType, fubenId)
	if !flagSaoDang {
		return
	}

	//判断是否空间足够
	if !inventoryManager.HasEnoughSlotsOfItemLevel(rewardsItemList) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"majorType":  majorType,
				"saoDangNum": saoDangNum,
			}).Warn("major:夫妻副本扫荡请求，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//扣除扫荡物品
	useItemReason := commonlog.InventoryLogReasonMajorSaodangUse
	useItemReasonText := fmt.Sprintf(useItemReason.String(), saoDangNum, majorType.String(), fubenId)
	if len(saodangNeedItemMap) > 0 {
		if flag := inventoryManager.BatchRemove(saodangNeedItemMap, useItemReason, useItemReasonText); !flag {
			panic("major: majorSaodang use item should be ok")
		}
	}

	//消耗元宝
	if saodangNeedGold > 0 {
		goldReason := commonlog.GoldLogReasonMajorSaoDangUse
		goldReasonText := fmt.Sprintf(goldReason.String(), saoDangNum, majorType.String(), fubenId)
		flag := propertyManager.CostGold(saodangNeedGold, true, goldReason, goldReasonText)
		if !flag {
			panic(fmt.Errorf("major: majorSaodang use gold should be ok"))
		}
	}

	//增加物品
	getItemReason := commonlog.InventoryLogReasonMajorSaodangGet
	getItemReasonText := fmt.Sprintf(getItemReason.String(), saoDangNum, majorType.String(), fubenId)
	flag = inventoryManager.BatchAddOfItemLevel(rewardsItemList, getItemReason, getItemReasonText)
	if !flag {
		panic("major: majorSaodang add item should be ok")
	}

	//获取扫荡固定资源
	reasonGold := commonlog.GoldLogReasonMajorSaoDangGet
	reasonSilver := commonlog.SilverLogReasonMajorSaoDangGet
	reasonLevel := commonlog.LevelLogReasonMajorSaoDangGet

	saodangGoldReasonText := fmt.Sprintf(reasonGold.String(), saoDangNum, majorType.String(), fubenId)
	saodangSilverReasonText := fmt.Sprintf(reasonSilver.String(), saoDangNum, majorType.String(), fubenId)
	expReasonText := fmt.Sprintf(reasonLevel.String(), saoDangNum, majorType.String(), fubenId)
	flag = propertyManager.AddRewData(totalRewData, reasonGold, saodangGoldReasonText, reasonSilver, saodangSilverReasonText, reasonLevel, expReasonText)
	if !flag {
		panic("major: majorSaodang add RewData should be ok")
	}

	//增加资源
	if len(rewardsResMap) > 0 {
		err = droplogic.AddRes(pl, rewardsResMap, reasonGold, saodangGoldReasonText, reasonSilver, saodangSilverReasonText, reasonLevel, expReasonText)
		if err != nil {
			return
		}
	}

	//完成扫荡
	manager.UseMajorNumByNum(majorType, saoDangNum)
	//通关事件
	data := majoreventtypes.CreateMajorSweepEventData(majorTemp, saoDangNum)

	gameevent.Emit(majoreventtypes.EventTypePlayerMajorSweep, pl, data)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCMajorSaoDang(majorType, fubenId, saoDangNum, showItemList)
	pl.SendMsg(scMsg)
	return
}
