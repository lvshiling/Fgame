package quest_guaji

import (
	chesslogic "fgame/fgame/game/chess/logic"
	playerchess "fgame/fgame/game/chess/player"
	chesstemplate "fgame/fgame/game/chess/template"
	chesstypes "fgame/fgame/game/chess/types"
	"fgame/fgame/game/global"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

//苍龙棋局
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeDragonChess, guaji.QuestGuaJiFunc(dragonChess))
}

//苍龙棋局
func dragonChess(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {
	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		return true
	}
	//获取当前任务数据
	questManager := p.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	q := questManager.GetQuestById(int32(questTemplate.TemplateId()))
	if q == nil {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("quest_guaji:任务不存在")
		return false
	}
	if q.QuestState != questtypes.QuestStateAccept {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("quest_guaji:任务没在进行中")
		return false
	}
	for k, v := range demandMap {
		chessType := chesstypes.ChessType(k)
		if !chessType.Valid() {
			log.WithFields(
				log.Fields{
					"playerId":  p.GetId(),
					"chessType": chessType,
				}).Warn("quest_guaji:棋局类型无效")
			return false
		}

		num := q.QuestDataMap[k]
		if num >= v {
			continue
		}
		needNum := v - num
		for i := 0; i < int(needNum); i++ {
			//棋局参加
			flag := chessAttend(p, chessType)
			if !flag {
				return false
			}
		}
		break
	}
	return true
}

func chessAttend(pl player.Player, typ chesstypes.ChessType) bool {
	chessManager := pl.GetPlayerDataManager(playertypes.PlayerChessDataManagerType).(*playerchess.PlayerChessDataManager)
	curChessId := chessManager.GetChessId(typ)
	chessTemplate := chesstemplate.GetChessTemplateService().GetChessByTypAndChessId(typ, curChessId)
	if chessTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"curChessId": curChessId,
				"typ":        typ.String(),
			}).Warn("quest_guaji:破解棋局,模板不存在")
		return false
	}

	//次数限制
	flag := chessManager.IsEnoughTimes(typ, 1)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("quest_guaji:破解棋局")
		return false
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needGold := int64(chessTemplate.GoldUse)
	needBindGold := int64(chessTemplate.BindGoldUse)
	needSilver := int64(chessTemplate.SilverUse)
	needItemId := chessTemplate.UseItemId
	needItemCount := chessTemplate.UseItemCount

	//物品是否足够
	totalNum := inventoryManager.NumOfItems(int32(needItemId))
	if totalNum < needItemCount {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"typ":           typ,
				"needItemId":    needItemId,
				"needItemCount": needItemCount,
			}).Warn("quest_guaji:破解棋局错误，道具不足")
		return false
	}

	//是否足够银两
	flag = propertyManager.HasEnoughSilver(needSilver)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("quest_guaji:破解棋局错误，银两不足")
		return false
	}

	//是否足够元宝
	flag = propertyManager.HasEnoughGold(needGold, false)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("quest_guaji:破解棋局错误，元宝不足")
		return false
	}
	//是否足够绑元
	needCostBindGold := needBindGold + needGold
	flag = propertyManager.HasEnoughGold(needCostBindGold, true)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("quest_guaji:破解棋局错误，绑元不足")
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	chesslogic.ChessAttend(pl, typ, now, false)
	return true
}
