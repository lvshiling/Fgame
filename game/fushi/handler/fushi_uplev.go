package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_FUSHI_UP_LEVEL_TYPE), dispatch.HandlerFunc(handleFushiUpLevel))
}

func handleFushiUpLevel(s session.Session, msg interface{}) (err error) {
	log.Debug("处理升级八卦符石请求消息,开始")
	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSFuShiUplevel)
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

	err = fushiUpLevel(tpl, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"error":    err,
			}).Error("fushi:处理升级八卦符石请求消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"index":    index,
		}).Error("fushi:处理升级八卦符石请求消息,成功")

	return
}

func fushiUpLevel(pl player.Player, typ fushitypes.FuShiType) (err error) {

	fushiManager := pl.GetPlayerDataManager(playertypes.PlayerFuShiDataManagerType).(*playerfushi.PlayerFuShiDataManager)
	if !fushiManager.IsFuShiActivite(typ) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("fushi:符石未激活")
		playerlogic.SendSystemMessage(pl, lang.FuShiNotActivite)
		return
	}

	fushiObj := fushiManager.GetFushiByTyp(typ)
	level := fushiObj.GetFushiLevel() + 1
	fushiTemp := fushitemplate.GetFuShiTemplateService().GetFuShiLevelByFuShiTypeAndLevel(typ, int32(level))
	if fushiTemp == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("fushi:符石已达到最高等级")
		playerlogic.SendSystemMessage(pl, lang.FuShiAlreadyTopLevel)
		return
	}

	itemMap := fushiTemp.GetItemMap()
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
		temp := fushitemplate.GetFuShiTemplateService().GetFuShiByFuShiType(typ)
		if temp == nil {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("fushi:模板不存在")
			playerlogic.SendSystemMessage(pl, lang.FuShiTemplateNotExist)
			return
		}
		reason := commonlog.InventoryLogReasonFuShiUpLevel
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonFuShiUpLevel.String(), temp.Name)
		flag := inventoryManager.BatchRemove(itemMap, reason, reasonText)
		if !flag {
			panic(fmt.Errorf("fushi: 消耗物品应该成功"))
		}

		// 同步背包信息
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag := fushiManager.FushiUpLevelSucess(typ)
	if !flag {
		panic(fmt.Errorf("fushi: 符石升级应该成功"))
	}

	scMsg := pbutil.BuildSCFuShiUpLevel(int32(typ), level)
	pl.SendMsg(scMsg)

	return
}
