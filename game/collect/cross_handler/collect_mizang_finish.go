package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/collect/pbutil"
	collecttemplate "fgame/fgame/game/collect/template"
	collecttypes "fgame/fgame/game/collect/types"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_COLLECT_MIZANG_FINISH_TYPE), dispatch.HandlerFunc(handleCollectMiZangFinish))
}

//处理采集完成
func handleCollectMiZangFinish(s session.Session, msg interface{}) (err error) {
	log.Debug("collect:处理跨服采集完成")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isCollectMiZangFinish := msg.(*crosspb.ISCollectMiZangFinish)
	biologyId := isCollectMiZangFinish.GetBiologyId()
	npcId := isCollectMiZangFinish.GetNpcId()
	openTypeInt := isCollectMiZangFinish.GetOpenType()
	miZangId := isCollectMiZangFinish.GetMiZangId()

	openType := collecttypes.MiZangOpenType(openTypeInt)

	if !openType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
				"openType": openType,
			}).Warn("collect:参数错误")

		return
	}
	err = collectMiZangFinish(tpl, biologyId, npcId, miZangId, openType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": biologyId,
			}).Error("collect:处理跨服采集完成,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"biologyId": biologyId,
		}).Debug("tulong:处理跨服采集完成,完成")
	return nil

}

//跨服采集完成
func collectMiZangFinish(pl player.Player, biologyId int32, npcId int64, miZangId int32, openType collecttypes.MiZangOpenType) (err error) {
	log.Debug("collect:密藏采集完成回跨服")

	miZangTemplate := collecttemplate.GetCollectTemplateService().GetMiZang(miZangId)
	if miZangTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"npcId":     npcId,
				"biologyId": biologyId,
				"miZangId":  miZangId,
			}).Warn("collect:处理:密藏采集消息,密藏不存在")
		return
	}
	miZangOpen := miZangTemplate.GetMiZang(openType)
	if miZangOpen == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("collect:处理:密藏采集消息,没有采集密藏")
		return
	}
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItemMap := miZangOpen.GetItemMap()
	if !inventoryManager.HasEnoughItems(needItemMap) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"npcId":    npcId,
			}).Warn("collect:处理:密藏采集消息,物品不足")

		return
	}
	reason := commonlog.InventoryLogReasonMiZangCost
	reasonText := fmt.Sprintf(reason.String(), openType.String())
	flag := inventoryManager.BatchRemove(needItemMap, reason, reasonText)
	if !flag {
		panic(fmt.Errorf("collect:物品移除应该成功"))
	}

	dropList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(miZangOpen.GetDropList())
	if len(dropList) != 0 {
		if !inventoryManager.HasEnoughSlotsOfItemLevel(dropList) {
			title := lang.GetLangService().ReadLang(lang.CollectMiZangTitle)
			content := lang.GetLangService().ReadLang(lang.CollectMiZangContent)
			now := global.GetGame().GetTimeService().Now()
			emaillogic.AddEmailItemLevel(pl, title, content, now, dropList)
		} else {
			reason := commonlog.InventoryLogReasonMiZangGet
			reasonText := fmt.Sprintf(reason.String(), openType.String())
			inventoryManager.BatchAddOfItemLevel(dropList, reason, reasonText)
		}
	}

	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCSceneCollectMiZangOpen(int32(openType), npcId, dropList)
	pl.SendMsg(scMsg)

	siCollectMiZangFinish := pbutil.BuildSICollectMiZangFinish(npcId)
	pl.SendCrossMsg(siCollectMiZangFinish)
	return
}
