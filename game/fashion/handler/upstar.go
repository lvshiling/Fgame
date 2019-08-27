package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/fashion/fashion"
	fashionlogic "fgame/fgame/game/fashion/logic"
	"fgame/fgame/game/fashion/pbutil"
	playerfashion "fgame/fgame/game/fashion/player"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_FASHION_UPGRADE_STAR_TYPE), dispatch.HandlerFunc(handleFashionUpstar))
}

//处理时装升星信息
func handleFashionUpstar(s session.Session, msg interface{}) (err error) {
	log.Debug("fashion:处理时装升星信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csFashionUpstar := msg.(*uipb.CSFashionUpstar)
	fashionId := csFashionUpstar.GetFashionId()

	err = fashionUpstar(tpl, fashionId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"fashionId": fashionId,
				"error":     err,
			}).Error("fashion:处理时装升星信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("fashion:处理时装升星完成")
	return nil
}

//时装升星的逻辑
func fashionUpstar(pl player.Player, fashionId int32) (err error) {
	fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(fashionId))
	if fashionTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("fashion:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if fashionTemplate.FashionUpgradeBeginId == 0 {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("fashion:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	fashionManager := pl.GetPlayerDataManager(types.PlayerFashionDataManagerType).(*playerfashion.PlayerFashionDataManager)
	flag := fashionManager.IfFashionExist(fashionId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("fashion:未激活的时装,无法升星")
		playerlogic.SendSystemMessage(pl, lang.FashionNotActiveNotUpstar)
		return
	}

	flag = fashionManager.IfCanUpStar(fashionId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("fashion:时装已满星")
		playerlogic.SendSystemMessage(pl, lang.FashionReacheFullStar)
		return
	}

	//升星需要物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	fashionInfo := fashionManager.GetFashion(fashionId)
	if fashionInfo == nil {
		return
	}
	star := fashionInfo.Star
	to := fashion.GetFashionService().GetFashionTemplate(int(fashionId))
	nextStar := star + 1
	fashionUpstarTemplate := to.GetFashionUpstarByLevel(nextStar)
	if fashionUpstarTemplate == nil {
		return
	}

	needItems := fashionUpstarTemplate.GetNeedItemMap(pl.GetRole(), pl.GetSex())
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":  pl.GetId(),
				"fashionId": fashionId,
			}).Warn("fashion:道具不足，无法升星")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//消耗物品
	if len(needItems) != 0 {
		reasonText := commonlog.InventoryLogReasonFashionUpstar.String()
		flag := inventoryManager.BatchRemove(needItems, commonlog.InventoryLogReasonFashionUpstar, reasonText)
		if !flag {
			panic(fmt.Errorf("fashion: fashionUpstar use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//时装升星判断
	pro, _, sucess := fashionlogic.FashionUpstar(fashionInfo.UpStarNum, fashionInfo.UpStarPro, fashionUpstarTemplate)
	flag = fashionManager.Upstar(fashionId, pro, sucess)
	if !flag {
		panic(fmt.Errorf("fashion: fashionUpstar should be ok"))
	}
	if sucess {
		fashionlogic.FashionPropertyChanged(pl)
	}
	scFashionUpstar := pbutil.BuildSCFashionUpstar(fashionId, fashionInfo.Star, fashionInfo.UpStarPro, sucess)
	pl.SendMsg(scFashionUpstar)
	return
}
