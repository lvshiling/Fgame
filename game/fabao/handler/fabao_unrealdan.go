package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	fabaologic "fgame/fgame/game/fabao/logic"
	"fgame/fgame/game/fabao/pbutil"
	playerfabao "fgame/fgame/game/fabao/player"
	fabaotemplate "fgame/fgame/game/fabao/template"
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

	processor.Register(codec.MessageType(uipb.MessageType_CS_FABAO_UNREALDAN_TYPE), dispatch.HandlerFunc(handleFaBaoUnrealDan))

}

//处理法宝食幻化丹信息
func handleFaBaoUnrealDan(s session.Session, msg interface{}) (err error) {
	log.Debug("fabao:处理法宝食幻化丹信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csFaBaoUnrealDan := msg.(*uipb.CSFaBaoUnrealDan)
	num := csFaBaoUnrealDan.GetNum()

	err = faBaoUnrealDan(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("fabao:处理法宝食幻化丹信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("fabao:处理法宝食幻化丹信息完成")
	return nil

}

//法宝食幻化丹的逻辑
func faBaoUnrealDan(pl player.Player, num int32) (err error) {
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("fabao:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	faBaoInfo := manager.GetFaBaoInfo()
	advancedId := faBaoInfo.GetAdvancedId()
	unrealLevel := faBaoInfo.GetUnrealLevel()
	fabaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(int32(advancedId))
	if fabaoTemplate == nil {
		return
	}

	huanHuaTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoHuanHuaTemplate(unrealLevel + 1)
	if huanHuaTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("fabao:幻化丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.FaBaoUnrealDanReachedFull)
		return
	}

	if unrealLevel >= fabaoTemplate.ShidanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("fabao:幻化丹食丹等级已达最大,请进阶后再试")
		playerlogic.SendSystemMessage(pl, lang.FaBaoUnrealDanReachedLimit)
		return
	}

	reachHuanHuaTemplate, flag := fabaotemplate.GetFaBaoTemplateService().GetFaBaoEatHuanHuanTemplate(unrealLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("fabao:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachHuanHuaTemplate.Level > fabaoTemplate.ShidanLimit {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("fabao:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := huanHuaTemplate.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("fabao:当前幻化丹药数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonFaBaoEatUn.String()
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonFaBaoEatUn, reasonText)
		if !flag {
			panic(fmt.Errorf("fabao:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	manager.EatUnrealDan(reachHuanHuaTemplate.Level)
	//同步属性
	fabaologic.FaBaoPropertyChanged(pl)

	scFaBaoShiDan := pbutil.BuildSCFaBaoUnrealDan(faBaoInfo.GetUnrealLevel(), faBaoInfo.GetUnrealPro())
	pl.SendMsg(scFaBaoShiDan)
	return
}
