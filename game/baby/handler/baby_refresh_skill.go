package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	babylogic "fgame/fgame/game/baby/logic"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	babytemplate "fgame/fgame/game/baby/template"
	gameevent "fgame/fgame/game/event"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_BABY_REFRESH_SKILL_TYPE), dispatch.HandlerFunc(handleBabyRefreshSkill))
}

//处理宝宝洗练技能
func handleBabyRefreshSkill(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理宝宝洗练技能消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSBabyRefreshSkill)
	babyId := csMsg.GetBabyId()

	err = handlerRefreshSkill(tpl, babyId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("baby:处理宝宝洗练技能消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("baby:处理宝宝洗练技能消息完成")
	return nil
}

// 洗练技能
func handlerRefreshSkill(pl player.Player, babyId int64) (err error) {
	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	baby := babyManager.GetBabyInfo(babyId)
	if baby == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"babyId":   babyId,
			}).Warn("baby:处理宝宝洗练技能, 宝宝不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	pregnantTemp := babytemplate.GetBabyTemplateService().GetBabyPregnantTemplateByQuality(baby.GetQuality())
	if pregnantTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"babyId":   babyId,
			}).Warn("baby:处理宝宝洗练技能, 宝宝模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//物品
	refreshTime := baby.GetRefreshTimes()
	babyConstantTemplate := babytemplate.GetBabyTemplateService().GetBabyConstantTemplate()
	useItemMap := babyConstantTemplate.GetRefreshTalentUseItemMap(refreshTime)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(useItemMap) > 0 {
		if !inventoryManager.HasEnoughItems(useItemMap) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"useItemMap": useItemMap,
				}).Warn("baby:处理宝宝洗练技能消息, 物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}

		itemUseReason := commonlog.InventoryLogReasonBabyRefreshUse
		itemUseReasonText := fmt.Sprintf(itemUseReason.String(), babyId, refreshTime)
		flag := inventoryManager.BatchRemove(useItemMap, itemUseReason, itemUseReasonText)
		if !flag {
			panic(fmt.Errorf("baby: 宝宝洗练技能消耗物品应该成功"))
		}
		inventorylogic.SnapInventoryChanged(pl)

		eventdata := babyeventtypes.CreatePlayerBabyLearnUseItemEventData(useItemMap)
		gameevent.Emit(babyeventtypes.EventTypeBabyLearnUseItem, pl, eventdata)
	}

	flag := babyManager.RefeshBabySkill(babyId)
	if !flag {
		panic(fmt.Errorf("baby：宝宝洗练应该成功,babyId:%d", babyId))
	}

	babylogic.BabyPropertyChanged(pl)

	scMsg := pbutil.BuildSCBabyRefreshSkill(babyId, baby.GetSkillList())
	pl.SendMsg(scMsg)
	return
}
