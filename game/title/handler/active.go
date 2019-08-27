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
	titlelogic "fgame/fgame/game/title/logic"
	"fgame/fgame/game/title/pbutil"
	playertitle "fgame/fgame/game/title/player"
	"fgame/fgame/game/title/title"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TITLE_ACTIVE_TYPE), dispatch.HandlerFunc(handleTitleActive))
}

//处理称号激活信息
func handleTitleActive(s session.Session, msg interface{}) (err error) {
	log.Debug("title:处理称号激活信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTitleActive := msg.(*uipb.CSTitleActive)
	titleId := csTitleActive.GetTitleId()

	err = titleActive(tpl, titleId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"titleId":  titleId,
				"error":    err,
			}).Error("title:处理称号激活信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Debug("dan:处理称号激活信息完成")
	return nil
}

//处理称号激活信息逻辑
func titleActive(pl player.Player, titleId int32) (err error) {
	titleManager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	flag := titleManager.IsValid(titleId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("title:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag = titleManager.IfTitleExist(titleId)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("title:该称号已激活,无需激活")
		playerlogic.SendSystemMessage(pl, lang.TitleRepeatActive)
		return
	}

	titleTemplate := title.GetTitleService().GetTitleTemplate(int(titleId))
	items := titleTemplate.GetNeedItemMap()
	if len(items) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//判断物品是否足够
		flag := inventoryManager.HasEnoughItems(items)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"titleId":  titleId,
			}).Warn("title:当前道具不足，无法激活该称号")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		inventoryReason := commonlog.InventoryLogReasonTitleActive
		reasonText := fmt.Sprintf(inventoryReason.String(), titleId)
		flag = inventoryManager.BatchRemove(items, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("title: titleActive use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)

	}

	titleObj, flag := titleManager.TitleActive(titleId)
	if !flag {
		panic(fmt.Errorf("title: titleActive should be ok"))
	}

	//同步属性
	titlelogic.TitlePropertyChanged(pl)

	scTitleActive := pbutil.BuildSCTitleActive(titleId, titleObj.ActiveTime, titleObj.ValidTime)
	err = pl.SendMsg(scTitleActive)
	return
}
