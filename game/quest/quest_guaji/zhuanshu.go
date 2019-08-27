package quest_guaji

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

//转生
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeZhuanSheng, guaji.QuestGuaJiFunc(zhuanSheng))
}

//转生
func zhuanSheng(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		return true
	}
	//获取当前任务数据
	questManager := p.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	propertyManager := p.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	q := questManager.GetQuestById(int32(questTemplate.TemplateId()))
	if q == nil {
		log.WithFields(
			log.Fields{
				"playerId":  p.GetId(),
				"questType": questTemplate.GetQuestSubType().String(),
			}).Warn("quest_guaji:任务不存在")
		return false
	}
	if q.QuestState != questtypes.QuestStateAccept {
		log.WithFields(
			log.Fields{
				"playerId":  p.GetId(),
				"questType": questTemplate.GetQuestSubType().String(),
			}).Warn("quest_guaji:任务没在进行中")
		return false
	}
	//TODO 验证是否可以升级
	num := demandMap[0]
	if num <= propertyManager.GetZhuanSheng() {
		return true
	}
	propertyManager.SetZhuanSheng(num, commonlog.ZhuanShengLogReasonGuaJi, commonlog.ZhuanShengLogReasonGuaJi.String())

	return true
}
