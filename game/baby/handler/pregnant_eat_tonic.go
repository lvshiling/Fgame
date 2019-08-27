package handler

/*
import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	babytemplate "fgame/fgame/game/baby/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BABY_EAT_TONIC_TYPE), dispatch.HandlerFunc(handleBabyEatTonic))
}

//处理吃补品
func handleBabyEatTonic(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理吃补品消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	// csMsg := msg.(*uipb.CSBabyEatTonic)

	err = babyEatTonic(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("baby:处理吃补品消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("baby:处理吃补品消息完成")
	return nil

}

//吃补品界面逻辑
func babyEatTonic(pl player.Player) (err error) {

	babyConstantTemplate := babytemplate.GetBabyTemplateService().GetBabyConstantTemplate()
	addPro := babyConstantTemplate.GetAddTonicPro()
	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	if babyManager.IsFullTonic(addPro) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("baby:处理吃补品消息,当前补品食用已达上限")
		playerlogic.SendSystemMessage(pl, lang.BabyFullTonic)
		return
	}

	needItemId := babyConstantTemplate.BupinItemId
	needItemNum := babyConstantTemplate.BupinItemCount
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughItem(needItemId, needItemNum) {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"needItemId":  needItemId,
				"needItemNum": needItemNum,
			}).Warn("baby:处理吃补品消息, 物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	itemUseReason := commonlog.InventoryLogReasonBabyEatTonicUse
	flag := inventoryManager.UseItem(needItemId, needItemNum, itemUseReason, itemUseReason.String())
	if !flag {
		panic(fmt.Errorf("baby: 吃补品消耗物品应该成功"))
	}

	//吃补品
	babyManager.EatTonic(addPro)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCBabyEatTonic(addPro)
	pl.SendMsg(scMsg)
	return
}
*/
