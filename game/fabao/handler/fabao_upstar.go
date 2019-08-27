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
	processor.Register(codec.MessageType(uipb.MessageType_CS_FABAO_UPSTAR_TYPE), dispatch.HandlerFunc(handleFaBaoUpstar))
}

//处理法宝皮肤升星信息
func handleFaBaoUpstar(s session.Session, msg interface{}) (err error) {
	log.Debug("fabao:处理法宝皮肤升星信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csFaBaoUpstar := msg.(*uipb.CSFaBaoUpstar)
	faBaoId := csFaBaoUpstar.GetFaBaoId()

	err = faBaoUpstar(tpl, faBaoId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"faBaoId":  faBaoId,
				"error":    err,
			}).Error("fabao:处理法宝皮肤升星信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("fabao:处理法宝皮肤升星完成")
	return nil
}

//法宝皮肤升星的逻辑
func faBaoUpstar(pl player.Player, faBaoId int32) (err error) {
	faBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBao(int(faBaoId))
	if faBaoTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"faBaoId":  faBaoId,
		}).Warn("fabao:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if faBaoTemplate.FaBaoUpstarBeginId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"faBaoId":  faBaoId,
		}).Warn("fabao:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	faBaoOtherInfo, flag := manager.IfFaBaoSkinExist(faBaoId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"faBaoId":  faBaoId,
		}).Warn("fabao:未激活的法宝皮肤,无法升星")
		playerlogic.SendSystemMessage(pl, lang.FaBaoSkinUpstarNoActive)
		return
	}

	_, flag = manager.IfCanUpStar(faBaoId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"faBaoId":  faBaoId,
		}).Warn("fabao:法宝皮肤已满星")
		playerlogic.SendSystemMessage(pl, lang.FaBaoSkinReacheFullStar)
		return
	}

	//升星需要物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	level := faBaoOtherInfo.GetLevel()
	nextLevel := level + 1
	to := fabaotemplate.GetFaBaoTemplateService().GetFaBao(int(faBaoId))
	if to == nil {
		return
	}
	faBaoUpstarTemplate := to.GetFaBaoUpstarByLevel(nextLevel)
	if faBaoUpstarTemplate == nil {
		return
	}

	needItems := faBaoUpstarTemplate.GetNeedItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
				"faBaoId":  faBaoId,
			}).Warn("fabao:道具不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//消耗物品
	if len(needItems) != 0 {
		reasonText := commonlog.InventoryLogReasonFaBaoUpstar.String()
		flag := inventoryManager.BatchRemove(needItems, commonlog.InventoryLogReasonFaBaoUpstar, reasonText)
		if !flag {
			panic(fmt.Errorf("fabao: faBaoUpstar use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//皮肤升星判断
	pro, _, sucess := fabaologic.FaBaoSkinUpstar(faBaoOtherInfo.GetUpNum(), faBaoOtherInfo.GetUpPro(), faBaoUpstarTemplate)
	flag = manager.Upstar(faBaoId, pro, sucess)
	if !flag {
		panic(fmt.Errorf("fabao: faBaoUpstar should be ok"))
	}
	if sucess {
		fabaologic.FaBaoPropertyChanged(pl)
	}
	scFaBaoUpstar := pbutil.BuildSCFaBaoUpstar(faBaoId, faBaoOtherInfo.GetLevel(), faBaoOtherInfo.GetUpPro())
	pl.SendMsg(scFaBaoUpstar)
	return
}
