package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	wardrobeeventtypes "fgame/fgame/game/wardrobe/event/types"
	wardrobelogic "fgame/fgame/game/wardrobe/logic"
	"fgame/fgame/game/wardrobe/pbutil"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
	wardrobetemplate "fgame/fgame/game/wardrobe/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WARDROBE_PEIYANG_TYPE), dispatch.HandlerFunc(handleWardrobePeiYang))
}

//处理衣橱培养信息
func handleWardrobePeiYang(s session.Session, msg interface{}) (err error) {
	log.Debug("wardrobe:处理衣橱培养信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csWardrobePeiYang := msg.(*uipb.CSWardrobePeiYang)
	typ := csWardrobePeiYang.GetType()
	num := csWardrobePeiYang.GetNum()
	err = wardrobePeiYang(tpl, int32(typ), num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"type":     typ,
				"num":      num,
				"error":    err,
			}).Error("wardrobe:处理衣橱培养信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"type":     typ,
			"num":      num,
		}).Debug("wardrobe:处理衣橱培养信息完成")
	return nil
}

//处理衣橱界面信息逻辑
func wardrobePeiYang(pl player.Player, typ int32, num int32) (err error) {
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"type":     typ,
			"num":      num,
		}).Warn("wardrobe:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	flag := manager.IsWardrobeActive(typ)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"type":     typ,
			"num":      num,
		}).Warn("wardrobe:未激活的衣橱套装,无法食培养丹")
		suitTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuSuitTemplate(typ)
		playerlogic.SendSystemMessage(pl, lang.WardrobeNotActiveNotEat, suitTemplate.Name)
		return
	}

	eatNum, flag := manager.IfCanPeiYang(typ)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"type":     typ,
			"num":      num,
		}).Warn("wardrobe:食丹等级已达最大")
		playerlogic.SendSystemMessage(pl, lang.WardrobeEatDanReachedLimit)
		return
	}
	suitTemplate := wardrobetemplate.GetWardrobeTemplateService().GetYiChuSuitTemplate(typ)
	if suitTemplate == nil {
		return
	}

	culLevel := manager.GetWardrobePeiYangNum(typ)
	culTemplate := suitTemplate.GetPeiYangByLevel(culLevel + 1)
	if culTemplate == nil {
		return
	}

	if num > eatNum {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"type":     typ,
			"num":      num,
			"eatNum":   eatNum,
		}).Warn("wardrobe:当前套装食用资质丹数量超过上限,无法食用")
		playerlogic.SendSystemMessage(pl, lang.WardrobeEatDanNumReachedLimit)
		return
	}

	reachPeiYangTemplate, flag := suitTemplate.GetEatPeiYangTemplate(culLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"type":     typ,
			"num":      num,
		}).Warn("wardrobe:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := culTemplate.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"type":     typ,
				"num":      num,
			}).Warn("wardrobe:物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonWardrobeEatDan := commonlog.InventoryLogReasonWardrobeEatDan
		reasonText := fmt.Sprintf(reasonWardrobeEatDan.String(), typ)
		flag = inventoryManager.UseItem(useItem, num, reasonWardrobeEatDan, reasonText)
		if !flag {
			panic(fmt.Errorf("wardrobe:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	manager.EatCulDan(typ, reachPeiYangTemplate.Level)
	//同步属性
	wardrobelogic.WardrobePropertyChanged(pl)

	scWardrobePeiYang := pbutil.BuildSCWardrobePeiYang(int32(typ), reachPeiYangTemplate.Level)
	pl.SendMsg(scWardrobePeiYang)

	//日志
	wardrobeReason := commonlog.WardrobeLogReasonUpgrade
	reasonText := fmt.Sprintf(wardrobeReason.String(), typ, suitTemplate.Name)
	eventData := wardrobeeventtypes.CreateWardrobePeiYangLogEventData(typ, reachPeiYangTemplate.Level, culLevel, wardrobeReason, reasonText)
	gameevent.Emit(wardrobeeventtypes.EventTypeWardrobePeiYangLog, pl, eventData)
	return
}
