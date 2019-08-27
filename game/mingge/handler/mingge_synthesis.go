package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	droptemplate "fgame/fgame/game/drop/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/mingge/pbutil"
	minggetemplate "fgame/fgame/game/mingge/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MINGGE_SYNTHESIS_TYPE), dispatch.HandlerFunc(handleMingGeSynthesis))
}

//处理命格合成信息
func handleMingGeSynthesis(s session.Session, msg interface{}) (err error) {
	log.Debug("mingge:处理命格合成信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMingGeSynthesis := msg.(*uipb.CSMingGeSynthesis)
	synthesisId := csMingGeSynthesis.GetSynthesisId()

	err = mingGeSynthesis(tpl, synthesisId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"synthesisId": synthesisId,
				"error":       err,
			}).Error("mingge:处理命格合成信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mingge:处理命格合成信息完成")
	return nil
}

//处理命格合成信息逻辑
func mingGeSynthesis(pl player.Player, synthesisId int32) (err error) {
	mingGeTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeSynthesisTemplate(synthesisId)
	if mingGeTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":    pl.GetId(),
			"synthesisId": synthesisId,
		}).Warn("mingge:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	needSilver := int64(mingGeTemplate.NeedSilver)
	needGold := int64(mingGeTemplate.NeedGold)
	needBindGold := int64(mingGeTemplate.NeedBindGold)

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//是否足够银两
	if needSilver != 0 {
		flag := propertyManager.HasEnoughSilver(needSilver)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":    pl.GetId(),
				"synthesisId": synthesisId,
			}).Warn("mingge:银两不足,无法合成")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	//是否足够元宝
	if needGold != 0 {
		flag := propertyManager.HasEnoughGold(needGold, false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":    pl.GetId(),
				"synthesisId": synthesisId,
			}).Warn("mingge:元宝不足,无法合成")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//是否足够绑元
	costBindGold := needGold + needBindGold
	if costBindGold != 0 {
		flag := propertyManager.HasEnoughGold(costBindGold, true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":    pl.GetId(),
				"synthesisId": synthesisId,
			}).Warn("mingge:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItemMap := mingGeTemplate.GetNeedItemMap()
	if len(needItemMap) != 0 {
		if !inventoryManager.HasEnoughItems(needItemMap) {
			log.WithFields(log.Fields{
				"playerId":    pl.GetId(),
				"synthesisId": synthesisId,
			}).Warn("mingge:命格物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		itemUsereason := commonlog.InventoryLogReasonMingGeSynthesisUse
		if flag := inventoryManager.BatchRemove(needItemMap, itemUsereason, itemUsereason.String()); !flag {
			panic(fmt.Errorf("mingge: mingGeSynthesis use item should be ok"))
		}
	}

	//消耗钱
	reasonGoldText := commonlog.GoldLogReasonMingGeSynthesisCost.String()
	reasonSliverText := commonlog.SilverLogReasonMingGeSynthesisCost.String()
	flag := propertyManager.Cost(needBindGold, needGold, commonlog.GoldLogReasonMingGeSynthesisCost, reasonGoldText, needSilver, commonlog.SilverLogReasonMingGeSynthesisCost, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("mingge: mingGeSynthesis Cost should be ok"))
	}

	synthesisItemId := int32(0)
	systhesisItemNum := int32(0)
	//合成成功率
	sucess := false
	if mathutils.RandomHit(common.MAX_RATE, int(mingGeTemplate.SuccessRate)) {
		sucess = true
	}

	if sucess {
		synthesisItemId = mingGeTemplate.ItemId
		systhesisItemNum = mingGeTemplate.ItemCount
	} else {
		synthesisItemId, systhesisItemNum = droptemplate.GetDropTemplateService().GetDropItem(mingGeTemplate.ShiBaiDrop)
	}

	//入背包
	if synthesisItemId != 0 && systhesisItemNum != 0 {
		inventoryReason := commonlog.InventoryLogReasonMingGeSynthesisAdd
		reasonText := inventoryReason.String()
		flag := inventoryManager.AddItem(synthesisItemId, systhesisItemNum, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("mingge: mingGeSynthesis AddItem should be ok"))
		}
	}

	//推送变化
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scMingGeSynthesis := pbutil.BuildSCMingGeSynthesis(sucess, synthesisItemId, systhesisItemNum)
	pl.SendMsg(scMingGeSynthesis)
	return
}
