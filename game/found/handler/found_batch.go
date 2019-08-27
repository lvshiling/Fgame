package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/found/pbutil"
	playerfound "fgame/fgame/game/found/player"
	foundtemplate "fgame/fgame/game/found/template"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FOUND_BATCH_TYPE), dispatch.HandlerFunc(handlerFoundBatch))
}

//一键找回
func handlerFoundBatch(s session.Session, msg interface{}) (err error) {
	log.Debug("found:处理资源一键找回请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csFoundBatch := msg.(*uipb.CSFoundBatch)
	typ := csFoundBatch.GetTyp()
	foundType := foundtypes.FoundType(typ)

	if !foundType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("found:参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = foundBatch(tpl, foundType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"typ":      typ,
				"err":      err,
			}).Error("found:处理资源一键找回请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
			"typ":      typ,
		}).Debug("found：处理资源一键找回请求完成")

	return
}

func foundBatch(pl player.Player, typ foundtypes.FoundType) (err error) {
	foundManager := pl.GetPlayerDataManager(types.PlayerFoundDataManagerType).(*playerfound.PlayerFoundDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	totalNeedGold := int64(0)
	totalNeedSilver := int64(0)
	foundTemplateMap := make(map[foundtypes.FoundResourceType]foundtypes.FoundData)
	goldUseReasonMap := make(map[foundtypes.FoundResourceType]int32)
	for _, obj := range foundManager.GetPreDayFoundList() {
		if obj.GetFoundStatus() == foundtypes.FoundBackStatusWaitReceive {
			temp := foundtemplate.GetFoundTemplateService().GetFoundTemplateByType(obj.GetResType(), obj.GetResLevel())
			if temp == nil {
				continue
			}
			foundTimes := foundManager.GetFoundTimes(obj.GetResType())

			totalNeedGold += int64(temp.FoundUsing * foundTimes)
			totalNeedSilver += int64(temp.FoundUsingSilver * foundTimes)
			foundTemplateMap[temp.GetResType()] = temp.GetFoundData(typ)
			goldUseReasonMap[temp.GetResType()] = foundTimes
		}
	}

	switch typ {
	case foundtypes.FoundTypeGold:
		{
			//判断元宝是否足够
			if !propertyManager.HasEnoughGold(int64(totalNeedGold), true) {
				log.WithFields(log.Fields{
					"playerId":      pl.GetId(),
					"typ":           typ,
					"totalNeedGold": totalNeedGold,
				}).Warn("found:元宝不足，无法完美找回")
				playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
				return
			}

			//消耗元宝
			if totalNeedGold > 0 {
				goldReason := commonlog.GoldLogReasonFoundResourceBatchCost
				goldReasonText := fmt.Sprintf(goldReason.String(), goldUseReasonMap)
				flag := propertyManager.CostGold(totalNeedGold, true, goldReason, goldReasonText)
				if !flag {
					panic("found:使用元宝应该成功")
				}
			}
		}
	case foundtypes.FoundTypeFree:
		{
			//判断银两是否足够
			if !propertyManager.HasEnoughSilver(totalNeedSilver) {
				log.WithFields(
					log.Fields{
						"playerId":        pl.GetId(),
						"typ":             typ,
						"totalNeedSilver": totalNeedSilver,
					}).Warn("found:银两不足，无法普通找回")
				playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
				return
			}

			//消耗银两
			if totalNeedSilver > 0 {
				silverUseReason := commonlog.SilverLogReasonFoundResourceBatchUse
				silverUseReasonText := fmt.Sprintf(silverUseReason.String(), goldUseReasonMap)
				flag := propertyManager.CostSilver(totalNeedSilver, silverUseReason, silverUseReasonText)
				if !flag {
					panic("found:消耗银两应该成功")
				}
			}
		}

	default:
		break
	}
	totalItemMap := make(map[int32]int32)

	for resType, foundData := range foundTemplateMap {
		tempTotalItemMap := addRes(pl, resType, foundData)
		for itemId, itemNum := range tempTotalItemMap {
			totalItemMap[itemId] = totalItemMap[itemId] + itemNum
		}
	}

	scFoundBatch := pbutil.BuildSCFoundBatch(totalItemMap)
	pl.SendMsg(scFoundBatch)
	return
}
