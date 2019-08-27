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
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONG_FASHION_UPSTAR_TYPE), dispatch.HandlerFunc(handleLingTongFashionUpstar))
}

//处理灵童时装升星信息
func handleLingTongFashionUpstar(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtong:处理灵童时装升星信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongFashionUpstar := msg.(*uipb.CSLingTongFashionUpstar)
	fashionId := csLingTongFashionUpstar.GetFashionId()

	err = lingTongFashionUpstar(tpl, fashionId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"fashionId": fashionId,
				"error":     err,
			}).Error("lingtong:处理灵童时装升星信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lingtong:处理灵童时装升星完成")
	return nil
}

//灵童时装升星的逻辑
func lingTongFashionUpstar(pl player.Player, fashionId int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongObj := manager.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("lingtong:请先激活灵童时装激活系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongActiveSystem)
		return
	}

	flag := lingtongtemplate.GetLingTongTemplateService().IsBornFashion(fashionId)
	if flag {
		return
	}

	lingTongFashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
	if lingTongFashionTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("lingtong:模板为空")
		return
	}

	if lingTongFashionTemplate.LingTongUpstarId == 0 {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("lingtong:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	fashionInfo := manager.GetFashionInfoById(fashionId)
	if fashionInfo == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("lingtong:未激活的时装无法升星")
		playerlogic.SendSystemMessage(pl, lang.LingTongFashionUpstarNoActivate)
		return
	}

	flag = manager.IfCanUpStar(fashionId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("lingtong:灵童时装已满星")
		playerlogic.SendSystemMessage(pl, lang.LingTongFashionReacheFullStar)
		return
	}

	//升星需要物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	star := fashionInfo.GetUpgradeLevel()
	nextStar := star + 1
	fashionUpstarTemplate := lingTongFashionTemplate.GetLingTongFashionUpstarByLevel(nextStar)
	if fashionUpstarTemplate == nil {
		return
	}

	needItems := fashionUpstarTemplate.GetNeedItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":  pl.GetId(),
				"fashionId": fashionId,
			}).Warn("lingtong:道具不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//消耗物品
	if len(needItems) != 0 {
		reasonText := commonlog.InventoryLogReasonLingTongFashionUpstar.String()
		flag := inventoryManager.BatchRemove(needItems, commonlog.InventoryLogReasonLingTongFashionUpstar, reasonText)
		if !flag {
			panic(fmt.Errorf("lingtong: lingTongFashionUpstar BatchRemove item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//灵童时装升星判断
	pro, _, sucess := lingtonglogic.LingTongFashionUpstar(fashionInfo.GetUpradeNum(), fashionInfo.GetUpgradePro(), fashionUpstarTemplate)
	flag = manager.Upstar(fashionId, pro, sucess)
	if !flag {
		panic(fmt.Errorf("lingtong: lingTongFashionUpstar should be ok"))
	}
	if sucess {
		lingtonglogic.LingTongFashionPropertyChanged(pl)
	}
	scLingTongFashionUpstar := pbutil.BuildSCLingTongFashionUpstar(fashionId, fashionInfo.GetUpgradeLevel(), fashionInfo.GetUpgradePro(), sucess)
	pl.SendMsg(scLingTongFashionUpstar)
	return
}
