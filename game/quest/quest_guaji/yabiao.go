package quest_guaji

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	transportationlogic "fgame/fgame/game/transportation/logic"
	playertransportation "fgame/fgame/game/transportation/player"
	transportationtemplate "fgame/fgame/game/transportation/template"
	transportationtypes "fgame/fgame/game/transportation/types"

	log "github.com/Sirupsen/logrus"
)

//进行X次押镖
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeDart, guaji.QuestGuaJiFunc(yaBiao))
}

//进行X次押镖
func yaBiao(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		return true
	}
	//获取当前任务数据
	questManager := p.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	q := questManager.GetQuestById(int32(questTemplate.TemplateId()))
	needNum := demandMap[0]
	num := q.QuestDataMap[0]
	if num >= needNum {
		return true
	}
	//押镖
	flag := doPersonalTransportation(p, transportationtypes.TransportationTypeSilver)
	if !flag {
		return false
	}
	return true
}

//押镖
func doPersonalTransportation(pl player.Player, typ transportationtypes.TransportationType) bool {
	tem := transportationtemplate.GetTransportationTemplateService().GetTransportationTemplate(typ)
	if tem == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("quest_guaji:参数错误,模板不存在")

		return false
	}

	transManager := pl.GetPlayerDataManager(playertypes.PlayerTransportationType).(*playertransportation.PlayerTransportationDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//是否足够次数
	if !transManager.IsEnoughTransportTimes() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("quest_guaji:领取押镖任务错误，次数不足")
		return false
	}

	if typ == transportationtypes.TransportationTypeSilver {
		//是否足够银两
		flag := propertyManager.HasEnoughSilver(int64(tem.BiaocheSilver))
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ,
				}).Warn("quest_guaji:领取押镖任务错误，银两不足")

			return false
		}
	}

	if typ == transportationtypes.TransportationTypeGold {
		//是否足够元宝
		flag := propertyManager.HasEnoughGold(int64(tem.BiaocheGold), true)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ,
				}).Warn("quest_guaji:领取押镖任务错误，元宝不足")

			return false
		}

	}
	transportationlogic.HandlePersonalTransportation(pl, typ)
	return true
}
