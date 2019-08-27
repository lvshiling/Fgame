package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	baguaplayer "fgame/fgame/game/bagua/player"
	"fgame/fgame/game/fushi/pbutil"
	playerfushi "fgame/fgame/game/fushi/player"
	fushitemplate "fgame/fgame/game/fushi/template"
	fushitypes "fgame/fgame/game/fushi/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FUSHI_ACTIVITE_TYPE), dispatch.HandlerFunc(handleFushiActivite))
}

func handleFushiActivite(s session.Session, msg interface{}) (err error) {
	log.Debug("处理激活八卦符石请求消息,开始")
	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSFuShiActivite)
	index := csMsg.GetTyp()
	typ := fushitypes.FuShiType(index)
	if !typ.Vaild() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"error":    err,
			}).Warn("fushi:参数错误")
		return
	}

	err = fushiActivite(tpl, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"error":    err,
			}).Error("fushi:处理激活八卦符石请求消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"index":    index,
		}).Error("fushi:处理激活八卦符石请求消息,成功")

	return
}

func fushiActivite(pl player.Player, typ fushitypes.FuShiType) (err error) {
	fushiTemp := fushitemplate.GetFuShiTemplateService().GetFuShiByFuShiType(typ)
	if fushiTemp == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("fushi:模板不存在")
		playerlogic.SendSystemMessage(pl, lang.FuShiTemplateNotExist)
		return
	}

	minMiJingLevel := fushiTemp.NeedBaGuaMiJing
	maxTongGuanLevel := getBaGuaMiJingMaxLevel(pl)
	if maxTongGuanLevel < minMiJingLevel {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("fushi:八卦秘境未达到指定通关等级，无法激活")
		playerlogic.SendSystemMessage(pl, lang.FuShiBaGuaMiJingLevelNotEnough)
		return
	}

	fushiManager := pl.GetPlayerDataManager(playertypes.PlayerFuShiDataManagerType).(*playerfushi.PlayerFuShiDataManager)
	if fushiManager.IsFuShiActivite(typ) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("fushi:符石已经激活")
		playerlogic.SendSystemMessage(pl, lang.FuShiAlreadyActivite)
		return
	}

	level := 1
	fushiLevelTemp := fushitemplate.GetFuShiTemplateService().GetFuShiLevelByFuShiTypeAndLevel(typ, int32(level))
	if fushiLevelTemp == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("fushi:模板不存在")
		playerlogic.SendSystemMessage(pl, lang.FuShiTemplateNotExist)
		return
	}

	itemMap := fushiLevelTemp.GetItemMap()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	if len(itemMap) != 0 {
		if !inventoryManager.HasEnoughItems(itemMap) {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("fushi:没有足够的物品")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		// 消耗物品
		reason := commonlog.InventoryLogReasonFuShiActivite
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonFuShiActivite.String(), fushiTemp.Name)
		flag := inventoryManager.BatchRemove(itemMap, reason, reasonText)
		if !flag {
			panic(fmt.Errorf("fushi: 消耗物品应该成功"))
		}

		// 同步背包信息
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag := fushiManager.FushiActitiveSucess(typ)
	if !flag {
		panic(fmt.Errorf("fushi: 符石激活应该成功"))
	}

	scMsg := pbutil.BuildSCFuShiActivite(int32(typ))
	pl.SendMsg(scMsg)

	return
}

// 获得八卦秘境通关等级
func getBaGuaMiJingMaxLevel(pl player.Player) int32 {
	baguaManager := pl.GetPlayerDataManager(playertypes.PlayerBaGuaDataManagerType).(*baguaplayer.PlayerBaGuaDataManager)
	return baguaManager.GetLevel()
}
