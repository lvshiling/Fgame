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
	processor.Register(codec.MessageType(uipb.MessageType_CS_FASHION_ACTIVE_TYPE), dispatch.HandlerFunc(handleFashionActive))
}

//处理时装激活信息
func handleFashionActive(s session.Session, msg interface{}) (err error) {
	log.Debug("fashion:处理时装激活信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csFashionActive := msg.(*uipb.CSFashionActive)
	fashionId := csFashionActive.GetFashionId()

	err = fashionActive(tpl, fashionId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"fashionId": fashionId,
				"error":     err,
			}).Error("fashion:处理时装激活信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Debug("fashion:处理时装激活信息完成")
	return nil
}

//处理时装激活信息逻辑
func fashionActive(pl player.Player, fashionId int32) (err error) {
	fashionManager := pl.GetPlayerDataManager(types.PlayerFashionDataManagerType).(*playerfashion.PlayerFashionDataManager)
	flag := fashionManager.IsValid(fashionId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("fashion:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag = fashionManager.IfFashionExist(fashionId)
	if flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("fashion:该时装已激活,无需激活")
		playerlogic.SendSystemMessage(pl, lang.FashionRepeatActive)
		return
	}

	fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(fashionId))
	items := fashionTemplate.GetNeedItemMap(pl.GetRole(), pl.GetSex())
	if len(items) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//判断物品是否足够
		flag := inventoryManager.HasEnoughItems(items)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":  pl.GetId(),
				"fashionId": fashionId,
			}).Warn("fashion:当前道具不足，无法激活该时装")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		inventoryReason := commonlog.InventoryLogReasonFashionActive
		reasonText := fmt.Sprintf(inventoryReason.String(), fashionId)
		flag = inventoryManager.BatchRemove(items, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("fashion: fashionActive use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	activeTime, flag := fashionManager.FashionActive(fashionId, true)
	if !flag {
		panic(fmt.Errorf("fashion: fashionActive should be ok"))
	}

	//同步属性
	fashionlogic.FashionPropertyChanged(pl)
	scFashionActive := pbutil.BuildSCFashionActive(fashionId, activeTime)
	pl.SendMsg(scFashionActive)
	return
}
