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
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	xiantilogic "fgame/fgame/game/xianti/logic"
	"fgame/fgame/game/xianti/pbutil"
	playerxianti "fgame/fgame/game/xianti/player"
	"fgame/fgame/game/xianti/xianti"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIANTI_UNREALDAN_TYPE), dispatch.HandlerFunc(handleXianTiUnrealDan))
}

//处理仙体食幻化丹信息
func handleXianTiUnrealDan(s session.Session, msg interface{}) (err error) {
	log.Debug("xianti:处理仙体食幻化丹信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csXianTiUnrealDan := msg.(*uipb.CSXiantiUnrealDan)
	num := csXianTiUnrealDan.GetNum()

	err = xianTiUnrealDan(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("xianti:处理仙体食幻化丹信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("xianti:处理仙体食幻化丹信息完成")
	return nil

}

//仙体食幻化丹的逻辑
func xianTiUnrealDan(pl player.Player, num int32) (err error) {
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("xianti:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	xianTiManager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	xianTiInfo := xianTiManager.GetXianTiInfo()
	advancedId := xianTiInfo.AdvanceId
	unrealLevel := xianTiInfo.UnrealLevel
	xianTiTemplate := xianti.GetXianTiService().GetXianTiNumber(int32(advancedId))
	if xianTiTemplate == nil {
		return
	}
	hunaHuaTemplate := xianti.GetXianTiService().GetXianTiHuanHuaTemplate(unrealLevel + 1)
	if hunaHuaTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("xianti:幻化丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.XianTiEatUnDanReachedFull)
		return
	}

	if unrealLevel >= xianTiTemplate.ShidanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("xianti:幻化丹食丹等级已达最大,请进阶后再试")
		playerlogic.SendSystemMessage(pl, lang.XianTiEatUnDanReachedLimit)
		return
	}

	reachHuanHuaTemplate, flag := xianti.GetXianTiService().GetXianTiEatHuanHuanTemplate(unrealLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("xianti:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachHuanHuaTemplate.Level > xianTiTemplate.ShidanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("xianti:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := hunaHuaTemplate.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("xianti:当前幻化丹药数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonXianTiEatUn.String()
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonXianTiEatUn, reasonText)
		if !flag {
			panic(fmt.Errorf("xianti:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	xianTiManager.EatUnrealDan(reachHuanHuaTemplate.Level)
	//同步属性
	xiantilogic.XianTiPropertyChanged(pl)

	scXianTiUnrealDan := pbutil.BuildSCXianTiUnrealDan(xianTiInfo.UnrealLevel, xianTiInfo.UnrealPro)
	pl.SendMsg(scXianTiUnrealDan)
	return
}
