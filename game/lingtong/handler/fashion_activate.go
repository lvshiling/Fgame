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
	playerlingtong "fgame/fgame/game/lingtong/player"
	"fmt"

	lingtonglogic "fgame/fgame/game/lingtong/logic"
	"fgame/fgame/game/lingtong/pbutil"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONG_FASHION_ACTIVE_TYPE), dispatch.HandlerFunc(handleLingTongFashionActivate))

}

//处理灵童时装激活信息
func handleLingTongFashionActivate(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtong:处理获取灵童时装激活消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongFashionActive := msg.(*uipb.CSLingTongFashionActive)
	fashionId := csLingTongFashionActive.GetFashionId()

	err = lingTongFashionActivate(tpl, fashionId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"fashionId": fashionId,
				"error":     err,
			}).Error("lingtong:处理获取灵童时装激活消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lingtong:处理获取灵童时装激活消息完成")
	return nil
}

//获取灵童时装激活界面信息逻辑
func lingTongFashionActivate(pl player.Player, fashionId int32) (err error) {
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
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
	if lingTongTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("lingtong:模板为空")
		return
	}
	fashionInfo := manager.GetFashionInfoById(fashionId)
	if fashionInfo != nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("lingtong:灵童时装重复激活")
		playerlogic.SendSystemMessage(pl, lang.LingTongFashionRepeateActivate)
		return
	}

	//激活需要物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItems := lingTongTemplate.GetNeedItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":  pl.GetId(),
				"fashionId": fashionId,
			}).Warn("lingtong:道具不足，无法激活灵童时装")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonLingTongFashionActivate.String(), fashionId)
		flag = inventoryManager.BatchRemove(needItems, commonlog.InventoryLogReasonLingTongFashionActivate, reasonText)
		if !flag {
			panic(fmt.Errorf("lingtong: lingTongFashionActivate BatchRemove item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	obj, flag := manager.FashionActive(fashionId)
	if !flag {
		panic(fmt.Errorf("lingtong: lingTongFashionActivate FashionActive should be ok"))
	}

	lingtonglogic.LingTongFashionPropertyChanged(pl)

	scLingTongFashionActivate := pbutil.BuildSCLingTongFashionActivate(obj)
	pl.SendMsg(scLingTongFashionActivate)
	return
}
