package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/fourgod/pbutil"
	"fgame/fgame/game/fourgod/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FOURGOD_USE_MASKED_TYPE), dispatch.HandlerFunc(handleFourGodUseMasked))
}

//处理使用蒙面衣
func handleFourGodUseMasked(s session.Session, msg interface{}) (err error) {
	log.Debug("fourgod:处理使用蒙面衣信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csFourGodUseMasked := msg.(*uipb.CSFourGodUseMasked)
	agree := csFourGodUseMasked.GetResult()
	err = fourGodUseMasked(tpl, agree)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"agree":    agree,
				"error":    err,
			}).Error("fourgod:处理使用蒙面衣信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"agree":    agree,
		}).Debug("fourgod:处理使用蒙面衣信息完成")
	return nil
}

//处理使用蒙面衣信息逻辑
func fourGodUseMasked(pl player.Player, agree bool) (err error) {
	//判断场景
	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeFourGodWar {
		return
	}

	fourGodConstTemplate := template.GetFourGodTemplateService().GetFourGodConstTemplate()
	itemId := fourGodConstTemplate.ItemId
	blackerBuffId := fourGodConstTemplate.BlackerBuffId
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	num := inventoryManager.NumOfItems(itemId)
	if num == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"agree":    agree,
		}).Warn("fourgod:当前未拥有蒙面衣")
		playerlogic.SendSystemMessage(pl, lang.FourGodBlackItemNoExist)
		return
	}

	reasonText := commonlog.InventoryLogReasonFourGodBlack.String()
	flag := inventoryManager.UseItem(itemId, num, commonlog.InventoryLogReasonFourGodBlack, reasonText)
	if !flag {
		panic(fmt.Errorf("fourgod: fourGodUseMasked use item should be ok"))
	}
	inventorylogic.SnapInventoryChanged(pl)

	if agree {
		scenelogic.AddBuff(pl, blackerBuffId, pl.GetId(), common.MAX_RATE)
	}
	scFourGodUseMarsked := pbuitl.BuildSCFourGodUseMarsked(agree)
	pl.SendMsg(scFourGodUseMarsked)
	return
}
