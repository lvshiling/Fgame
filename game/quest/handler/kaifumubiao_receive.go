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
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_KAIFUMUBIAO_RECEIVE_TYPE), dispatch.HandlerFunc(handleQuestKaiFuMuBiaoReceive))
}

//处理获取领取开服目标组奖励
func handleQuestKaiFuMuBiaoReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理获取领取开服目标组奖励")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csQuestkaiFuMuBiaoReceive := msg.(*uipb.CSQuestkaiFuMuBiaoReceive)
	kaiFuTime := csQuestkaiFuMuBiaoReceive.GetKaiFuTime()

	err = questKaiFuMuBiaoReceive(tpl, kaiFuTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("quest:处理获取领取开服目标组奖励,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理获取领取开服目标组奖励,完成")
	return nil
}

//获取领取开服目标组奖励
func questKaiFuMuBiaoReceive(pl player.Player, kaiFuTime int32) (err error) {
	if kaiFuTime < 1 || kaiFuTime > questtypes.KaiFuMuBiaoDayMax {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"kaiFuTime": kaiFuTime,
			}).Warn("quest:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	kaiFuMuBiaoTemplate := questtemplate.GetQuestTemplateService().GetKaiFuMuBiaoTemplate(kaiFuTime)
	if kaiFuMuBiaoTemplate == nil {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	flag := manager.IfCanReceive(kaiFuTime)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"kaiFuTime": kaiFuTime,
			}).Warn("quest:领取组奖励失败")
		playerlogic.SendSystemMessage(pl, lang.QuestKaiFuMuBiaoGroupReceive)
		return
	}

	//奖励物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	rewItemMap := kaiFuMuBiaoTemplate.GetRewItemMap()
	if len(rewItemMap) != 0 {
		//添加物品
		if !inventoryManager.HasEnoughSlots(rewItemMap) {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"kaiFuTime": kaiFuTime,
				}).Warn("quest:背包空间不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonKaiFuMuBiaoGroupRew.String(), kaiFuTime)
		flag := inventoryManager.BatchAdd(rewItemMap, commonlog.InventoryLogReasonKaiFuMuBiaoGroupRew, reasonText)
		if !flag {
			panic("quest: questKaiFuMuBiaoReceive BatchAdd  should be ok")
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	silver := int64(kaiFuMuBiaoTemplate.RewardSilver)
	//奖励银两
	if silver != 0 {
		reason := commonlog.SilverLogReasonKaiFuMuBiaoReward
		reasonText := fmt.Sprintf(reason.String(), kaiFuTime)
		propertyManager.AddSilver(silver, reason, reasonText)
		propertylogic.SnapChangedProperty(pl)
	}

	flag = manager.GroupReward(kaiFuTime)
	if !flag {
		panic("quest: GroupReward  should be ok")
	}
	scQuestKaiFuMuBiaoReceive := pbutil.BuildSCQuestKaiFuMuBiaoReceive(kaiFuTime)
	pl.SendMsg(scQuestKaiFuMuBiaoReceive)
	return
}
