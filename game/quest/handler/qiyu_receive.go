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
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_QIYU_RECEIVE_TYPE), dispatch.HandlerFunc(handleQuestQiYuReceive))
}

//处理获取领取奇遇奖励
func handleQuestQiYuReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理获取领取奇遇奖励")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSQuestQiYuReceive)
	qiyuId := csMsg.GetQiyuId()

	err = questQiYuReceive(tpl, qiyuId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("quest:处理获取领取奇遇奖励,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理获取领取奇遇奖励,完成")
	return nil
}

//获取领取奇遇奖励
func questQiYuReceive(pl player.Player, qiyuId int32) (err error) {

	qiYuTemplate := questtemplate.GetQuestTemplateService().GetQiYuTemplate(qiyuId)
	if qiYuTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"qiyuId":   qiyuId,
			}).Warn("quest:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	if !manager.IsFinishQiYu(qiyuId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"qiyuId":   qiyuId,
			}).Warn("quest:奇遇任务领取失败，奇遇任务未完成")
		playerlogic.SendSystemMessage(pl, lang.QuestQiYuReceiveFail)
		return
	}

	if manager.IsReceiveQiYu(qiyuId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"qiyuId":   qiyuId,
			}).Warn("quest:奇遇任务领取失败，奇遇任务奖励已领取")
		playerlogic.SendSystemMessage(pl, lang.QuestRewHadReceive)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//奖励物品
	rewItemMap := qiYuTemplate.GetRewItemMap()
	if len(rewItemMap) != 0 {
		if !inventoryManager.HasEnoughSlots(rewItemMap) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"qiyuId":   qiyuId,
				}).Warn("quest:背包空间不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
		reason := commonlog.InventoryLogReasonQiYuRew
		reasonText := fmt.Sprintf(reason.String(), qiyuId)
		flag := inventoryManager.BatchAdd(rewItemMap, reason, reasonText)
		if !flag {
			panic("quest: questQiYuReceive BatchAdd  should be ok")
		}
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	reasonGold := commonlog.GoldLogReasonQuestQiYuRew
	reasonSilver := commonlog.SilverLogReasonQuestQiYuRew
	reasonLevel := commonlog.LevelLogReasonQuestQiYuRew
	reasonGoldText := fmt.Sprintf(reasonGold.String(), qiyuId)
	reasonSliverText := fmt.Sprintf(reasonSilver.String(), qiyuId)
	reasonlevelText := fmt.Sprintf(reasonLevel.String(), qiyuId)
	totalRewData := propertytypes.CreateRewData(qiYuTemplate.RewExp, qiYuTemplate.RewExpPoint, qiYuTemplate.RewSilver, qiYuTemplate.RewGold, qiYuTemplate.RewBindGold)
	flag := propertyManager.AddRewData(totalRewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
	if !flag {
		panic(fmt.Errorf("quest: questQiYuReceive AddRewData  should be ok"))
	}
	//同步
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	manager.ReceiveQiYu(qiyuId)
	qiyu := manager.GetQiYu(qiyuId)
	scMsg := pbutil.BuildSCQuestQiYuReceive(qiyuId, qiyu.GetIsReceive())
	pl.SendMsg(scMsg)
	return
}
