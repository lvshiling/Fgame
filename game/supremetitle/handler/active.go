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
	supremetitlelogic "fgame/fgame/game/supremetitle/logic"
	"fgame/fgame/game/supremetitle/pbutil"
	playersupremetitle "fgame/fgame/game/supremetitle/player"
	supremetitletemplate "fgame/fgame/game/supremetitle/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SUPREME_TITLE_ACTIVE_TYPE), dispatch.HandlerFunc(handleSupremeTitleActive))
}

//处理至尊称号激活信息
func handleSupremeTitleActive(s session.Session, msg interface{}) (err error) {
	log.Debug("supremetitle:处理至尊称号激活信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSupremeTitleActive := msg.(*uipb.CSSupremeTitleActive)
	titleId := csSupremeTitleActive.GetTitleId()

	err = supremeTitleActive(tpl, titleId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"titleId":  titleId,
				"error":    err,
			}).Error("supremetitle:处理至尊称号激活信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Debug("dan:处理至尊称号激活信息完成")
	return nil
}

//处理至尊称号激活信息逻辑
func supremeTitleActive(pl player.Player, titleId int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSupremeTitleDataManagerType).(*playersupremetitle.PlayerSupremeTitleDataManager)
	titleTemplate := supremetitletemplate.GetTitleDingZhiTemplateService().GetTitleDingZhiTempalte(titleId)
	if titleTemplate == nil {
		return
	}
	flag := manager.IfTitleExist(titleId)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("supremetitle:该至尊称号已激活,无需激活")
		playerlogic.SendSystemMessage(pl, lang.SupremeTitleRepeatActive)
		return
	}

	items := titleTemplate.GetNeedItemMap()
	if len(items) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//判断物品是否足够
		flag := inventoryManager.HasEnoughItems(items)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"titleId":  titleId,
			}).Warn("supremetitle:当前道具不足，无法激活该至尊称号")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		inventoryReason := commonlog.InventoryLogReasonTitleActive
		reasonText := fmt.Sprintf(inventoryReason.String(), titleId)
		flag = inventoryManager.BatchRemove(items, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("supremetitle: supremeTitleActive use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)

	}

	flag = manager.TitleActive(titleId)
	if !flag {
		panic(fmt.Errorf("supremetitle: supremeTitleActive should be ok"))
	}

	//同步属性
	supremetitlelogic.SupremeTitlePropertyChanged(pl)
	scTitleActive := pbutil.BuildSCSupremeTitleActive(titleId)
	pl.SendMsg(scTitleActive)
	return
}
