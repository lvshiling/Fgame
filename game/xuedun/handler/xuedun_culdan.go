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
	xuedunlogic "fgame/fgame/game/xuedun/logic"
	"fgame/fgame/game/xuedun/pbutil"
	playerxuedun "fgame/fgame/game/xuedun/player"
	xueduntemplate "fgame/fgame/game/xuedun/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_XUEDUN_PEIYANGE_TYPE), dispatch.HandlerFunc(handleXueDunPeiYang))

}

//处理血盾培养信息
func handleXueDunPeiYang(s session.Session, msg interface{}) (err error) {
	log.Debug("xuedun:处理血盾培养信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csXueDunPeiYang := msg.(*uipb.CSXueDunPeiYang)
	num := csXueDunPeiYang.GetNum()

	err = xueDunCulDan(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("xuedun:处理血盾培养信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("xuedun:处理血盾培养信息完成")
	return nil

}

//血盾培养逻辑
func xueDunCulDan(pl player.Player, num int32) (err error) {
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("xuedun:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerXueDunDataManagerType).(*playerxuedun.PlayerXueDunDataManager)
	xueDunInfo := manager.GetXueDunInfo()
	culLevel := xueDunInfo.GetCulLevel()
	limitLevel := manager.GetCulLevelLimit()
	peiYangTempalte := xueduntemplate.GetXueDunTemplateService().GetXueDunPeiYangTemplate(culLevel + 1)
	if peiYangTempalte == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("xuedun:培养丹食丹等级满级")
		playerlogic.SendSystemMessage(pl, lang.XueDunPeiYangReachFull)
		return
	}

	if culLevel >= limitLevel {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("xuedun:血盾培养等级已达最大,请升阶后再试")
		playerlogic.SendSystemMessage(pl, lang.XueDunEatCulDanReachedLimit)
		return
	}

	reachPeiYangTemplate, flag := xueduntemplate.GetXueDunTemplateService().GetXueDunEatPeiYangTemplate(culLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("xuedun:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if reachPeiYangTemplate.Level > limitLevel {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("xuedun:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := peiYangTempalte.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("xuedun:当前培养丹数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonXueDunEatClu.String()
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonXueDunEatClu, reasonText)
		if !flag {
			panic(fmt.Errorf("xuedun:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	manager.EatCulDan(reachPeiYangTemplate.Level)
	xuedunlogic.XueDunPropertyChanged(pl)

	scXueDunPeiYang := pbutil.BuildSCXueDunPeiYang(xueDunInfo)
	pl.SendMsg(scXueDunPeiYang)
	return
}
