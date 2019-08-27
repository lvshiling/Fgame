package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	secretcardlogic "fgame/fgame/game/secretcard/logic"
	"fgame/fgame/game/secretcard/pbutil"
	playersecretcard "fgame/fgame/game/secretcard/player"
	"fgame/fgame/game/secretcard/secretcard"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_STAR_REW_TYPE), dispatch.HandlerFunc(handleSecretStarRew))
}

//处理天机牌星数奖励信息
func handleSecretStarRew(s session.Session, msg interface{}) (err error) {
	log.Debug("secretcard:处理天机牌星数奖励消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csQuestSecretStarRew := msg.(*uipb.CSQuestSecretStarRew)
	openBox := csQuestSecretStarRew.GetOpenBox()

	err = secretStarRew(tpl, openBox)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"openBox":  openBox,
				"error":    err,
			}).Error("secretcard:处理天机牌星数奖励消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("secretcard:处理天机牌星数奖励消息完成")
	return nil
}

//天机牌星数奖励信息的逻辑
func secretStarRew(pl player.Player, openBox int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSecretCardDataManagerType).(*playersecretcard.PlayerSecretCardDataManager)
	flag := manager.IfSecretStarRew(openBox)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"openBox":  openBox,
			}).Warn("secretcard:条件不足或已领取玩")
		playerlogic.SendSystemMessage(pl, lang.SecretCardStarNoEnough)
		return
	}

	starTemplate := secretcard.GetSecretCardService().GetStarTemplate(openBox)
	if starTemplate == nil {
		return
	}

	//判断背包是否足够
	rewItemMap := starTemplate.GetRewItemMap()
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(rewItemMap) != 0 {
		flag = inventoryManager.HasEnoughSlots(rewItemMap)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("secretcard:背包空间不足,清理后再来")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
	}

	//宝箱奖励
	isReturn := secretcardlogic.GiveSecretCardBoxReward(pl, starTemplate)
	if isReturn {
		return
	}
	flag = manager.SecretStarRew(openBox)
	if !flag {
		panic(fmt.Errorf("secretcard: secretStarRew SecretStarRew should be ok"))
	}
	scQuestSecretStarRew := pbutil.BuildSCQuestSecretStarRew(openBox)
	pl.SendMsg(scQuestSecretStarRew)
	return
}
