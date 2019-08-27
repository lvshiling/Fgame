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
	gamesession "fgame/fgame/game/session"
	shihunfanlogic "fgame/fgame/game/shihunfan/logic"
	"fgame/fgame/game/shihunfan/pbutil"
	playershihunfan "fgame/fgame/game/shihunfan/player"
	shihunfantemplate "fgame/fgame/game/shihunfan/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHIHUNFAN_DAN_ADVANCED_TYPE), dispatch.HandlerFunc(handleShiHunFanEatDan))
}

//处理使用噬魂幡丹
func handleShiHunFanEatDan(s session.Session, msg interface{}) (err error) {
	log.Debug("shihunfan:处理使用噬魂幡丹")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csShiHunFanDan := msg.(*uipb.CSShihunfanDanAdvanced)
	num := csShiHunFanDan.GetNum()
	if num <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("shihunfan:参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = shiHunFanEatDan(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("shihunfan:处理使用噬魂幡丹,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shihunfan:处理使用噬魂幡丹完成")
	return nil

}

//食用噬魂幡丹
func shiHunFanEatDan(pl player.Player, num int32) (err error) {
	shihunfanManager := pl.GetPlayerDataManager(playertypes.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	shihunfanInfo := shihunfanManager.GetShiHunFanInfo()
	advancedId := int32(shihunfanInfo.AdvanceId)
	danLevel := shihunfanInfo.DanLevel
	shihunfanTemplate := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanNumber(advancedId)
	if shihunfanTemplate == nil {
		return
	}
	nextTemplate := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanDan(danLevel + 1)
	if nextTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shihunfan:噬魂幡丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.ShiHunFanEatDanReachedFull)
		return
	}

	if danLevel >= shihunfanTemplate.ShidanLimit {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shihunfan:噬魂幡丹食丹等级已达最大,请进阶后再试")
		playerlogic.SendSystemMessage(pl, lang.ShiHunFanEatDanReachedLimit)
		return
	}

	reachDanTemplate, flag := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanEatDanTemplate(danLevel, num)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shihunfan:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachDanTemplate.Level > shihunfanTemplate.ShidanLimit {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("shihunfan:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := nextTemplate.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("shihunfan:当前丹药数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonShiHunFanEatDan.String()
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonShiHunFanEatDan, reasonText)
		if !flag {
			panic(fmt.Errorf("shihunfan:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	shihunfanManager.EatShiHunFanDan(reachDanTemplate.Level)
	//同步属性
	shihunfanlogic.ShiHunFanPropertyChanged(pl)

	scShiHunFanEatDan := pbutil.BuildSCShiHunFanEatDan(shihunfanInfo)
	pl.SendMsg(scShiHunFanEatDan)
	return
}
