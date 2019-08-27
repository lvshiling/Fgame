package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	danlogic "fgame/fgame/game/dan/logic"
	"fgame/fgame/game/dan/pbutil"
	playerdan "fgame/fgame/game/dan/player"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DAN_USE_TYPE), dispatch.HandlerFunc(handleDanUse))
}

//处理全部食用
func handleDanUse(s session.Session, msg interface{}) (err error) {
	log.Debug("dan:处理全部食用消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = danUse(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("dan:处理全部食用消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("dan:处理全部食用消息完成")
	return nil
}

//处理全部食用逻辑
func danUse(pl player.Player) (err error) {
	danManager := pl.GetPlayerDataManager(types.PlayerDanDataManagerType).(*playerdan.PlayerDanDataManager)
	flag := danManager.CheckFullLevel()
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("dan:当前食丹等级已达最高级,无法再食用")
		playerlogic.SendSystemMessage(pl, lang.DanLevelReachedLimitNotEat)
		return
	}

	dans, flag := danManager.WhatEat()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("dan:当前食用已达上限，请升级后再食用")
		playerlogic.SendSystemMessage(pl, lang.DanUseReachedLimit)
		return
	}

	//从背包获取丹药数量
	nums := int32(0)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	for danId, danNum := range dans {
		num := inventoryManager.NumOfItems(int32(danId))
		if num <= 0 {
			delete(dans, danId)
			continue
		}
		nums = num
		if num >= danNum {
			nums = danNum
		}
		dans[danId] = nums
	}
	if len(dans) == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("dan:当前等级无可食用丹药，请先炼丹")
		playerlogic.SendSystemMessage(pl, lang.DanNotCanUse)
		return
	}

	//食丹扣除物品
	reasonText := commonlog.InventoryLogReasonDanEat.String()
	flag = inventoryManager.BatchRemove(dans, commonlog.InventoryLogReasonDanEat, reasonText)
	if !flag {
		panic(fmt.Errorf("dan: danUse use item should be ok"))
	}
	//同步物品
	inventorylogic.SnapInventoryChanged(pl)
	//食丹
	danManager.EatDan(dans)
	danlogic.DanPropertyChanged(pl)
	danInfo := danManager.GetDanInfo()
	scDanGet := pbuitl.BuildSCDanUse(danInfo.DanInfoMap)
	pl.SendMsg(scDanGet)
	return
}
