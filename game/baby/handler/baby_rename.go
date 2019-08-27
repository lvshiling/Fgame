package handler

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
	processor.Register(codec.MessageType(uipb.MessageType_CS_BABY_CHANGE_NAME_TYPE), dispatch.HandlerFunc(handleBabyChangeName))
}

//处理宝宝改名
func handleBabyChangeName(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理宝宝改名消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSBabyChangeName)
	babyId := csMsg.GetBabyId()
	newName := csMsg.GetNewName()

	err = handlerRename(tpl, babyId, newName)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("baby:处理宝宝改名消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("baby:处理宝宝改名消息完成")
	return nil

}

// 改名卡
func handlerRename(pl player.Player, babyId int64, newName string) (err error) {

	//物品
	babyConstantTemplate := babytemplate.GetBabyTemplateService().GetBabyConstantTemplate()
	useItemMap := make(map[int32]int32)
	useItemMap[babyConstantTemplate.GaiMingKaId] = 1
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughItems(useItemMap) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"useItemMap": useItemMap,
			}).Warn("baby:处理宝宝改名消息, 物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	if len(newName) <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"newName":  newName,
			}).Warn("baby:处理宝宝改名消息,名字无效")
		playerlogic.SendSystemMessage(pl, lang.NameInvalid)
		return
	}

	itemUseReason := commonlog.InventoryLogReasonBabyChangeName
	itemUseReasonText := fmt.Sprintf(itemUseReason.String(), babyId)
	flag := inventoryManager.BatchRemove(useItemMap, itemUseReason, itemUseReasonText)
	if !flag {
		panic(fmt.Errorf("baby: 宝宝改名消耗物品应该成功"))
	}
	inventorylogic.SnapInventoryChanged(pl)

	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	babyManager.UpdateBabyName(babyId, newName)

	scMsg := pbutil.BuildSCBabyChangeName(babyId, newName)
	pl.SendMsg(scMsg)
	return
}
