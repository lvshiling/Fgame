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
	"fgame/fgame/game/title/pbutil"
	playertitle "fgame/fgame/game/title/player"
	"fgame/fgame/game/title/title"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TITLE_ADD_VALID_TIME_TYPE), dispatch.HandlerFunc(handleTitleAddValid))
}

//处理称号增加时效
func handleTitleAddValid(s session.Session, msg interface{}) (err error) {
	log.Debug("title:处理称号增加时效")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSTitleAddValidTime)
	titleId := csMsg.GetTitleId()

	err = titleAddValidTime(tpl, titleId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"titleId":  titleId,
				"error":    err,
			}).Error("title:处理称号增加时效,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Debug("dan:处理称号增加时效完成")
	return nil
}

//处理称号增加时效逻辑
func titleAddValidTime(pl player.Player, titleId int32) (err error) {
	titleManager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
	flag := titleManager.IsValid(titleId)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"titleId":  titleId,
			}).Warn("title:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	titleTemp := title.GetTitleService().GetTitleTemplate(int(titleId))
	title := titleManager.GetTitleInfo(titleTemp.GetTitleType(), titleId)
	if title.ValidTime == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("title:该称号不是限时称号,不能叠加时效")
		playerlogic.SendSystemMessage(pl, lang.TitleNotValid)
		return
	}

	//判断物品是否足够
	needItemMap := titleTemp.GetNeedItemMap()
	if len(needItemMap) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		flag := inventoryManager.HasEnoughItems(needItemMap)
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
		flag = inventoryManager.BatchRemove(needItemMap, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("title: titleActive use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag = titleManager.TitleAddValid(titleId)
	if !flag {
		panic(fmt.Errorf("title: title add valid should be ok"))
	}

	scMsg := pbutil.BuildSCTitleAddValidTime(titleId, title.ValidTime)
	err = pl.SendMsg(scMsg)
	return
}
