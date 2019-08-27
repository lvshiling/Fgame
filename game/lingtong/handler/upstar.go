package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	lingtonglogic "fgame/fgame/game/lingtong/logic"
	"fgame/fgame/game/lingtong/pbutil"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONG_UPSTAR_TYPE), dispatch.HandlerFunc(handleLingTongUpstar))
}

//处理灵童升星信息
func handleLingTongUpstar(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtong:处理灵童升星信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongUpstar := msg.(*uipb.CSLingTongUpstar)
	lingTongId := csLingTongUpstar.GetLingTongId()

	err = lingTongUpstar(tpl, lingTongId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"lingTongId": lingTongId,
				"error":      err,
			}).Error("lingtong:处理灵童升星信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Debug("lingtong:处理灵童升星完成")
	return nil
}

//灵童升星的逻辑
func lingTongUpstar(pl player.Player, lingTongId int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeLingTongUpstar) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("lingtong: 灵童升星失败，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Warn("lingtong:模板为空")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongInfo, flag := manager.GetLingTongInfo(lingTongId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Warn("lingtong:未激活该灵童")
		playerlogic.SendSystemMessage(pl, lang.LingTongNoActive)
		return
	}

	if lingTongTemplate.LingTongUpstarBeginId == 0 {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Warn("lingtong:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	curLevel := lingTongInfo.GetStarLevel()
	nextLevel := curLevel + 1
	nextLingTongUpstarTemplate := lingTongTemplate.GetLingTongUpstarByLevel(nextLevel)
	if nextLingTongUpstarTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Warn("lingtong:已达满级")
		playerlogic.SendSystemMessage(pl, lang.LingTongUpstarReachFull)
		return
	}

	//升星需要物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItems := nextLingTongUpstarTemplate.GetItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":   pl.GetId(),
				"lingTongId": lingTongId,
			}).Warn("lingtong:道具不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//消耗物品
	if len(needItems) != 0 {
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonLingTongUpstar.String(), lingTongId, curLevel)
		flag := inventoryManager.BatchRemove(needItems, commonlog.InventoryLogReasonLingTongUpstar, reasonText)
		if !flag {
			panic(fmt.Errorf("lingtong: lingTongUpstar BatchRemove item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//灵童升星判断
	pro, _, sucess := lingtonglogic.LingTongUpstarJudge(lingTongInfo.GetStarNum(), lingTongInfo.GetStarPro(), nextLingTongUpstarTemplate)
	flag = manager.LingTongUpstar(lingTongId, pro, sucess)
	if !flag {
		panic(fmt.Errorf("lingtong: lingTongUpstar should be ok"))
	}
	if sucess {
		//manager.AddLevel()
		lingtonglogic.LingTongPropertyChanged(pl)
	}
	scMsg := pbutil.BuildSCLingTongUpstar(lingTongId, lingTongInfo.GetStarLevel(), lingTongInfo.GetStarPro(), sucess)
	pl.SendMsg(scMsg)
	return
}
