package quest_guaji

import (
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/player"
	"fgame/fgame/game/quest/guaji/guaji"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	scenelogic "fgame/fgame/game/scene/logic"
	scenetemplate "fgame/fgame/game/scene/template"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeCollect, guaji.QuestGuaJiFunc(collectQuestGuaJi))
}

var (
	collectTime     = 2
	collectDistance = 1
)

func collectQuestGuaJi(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {
	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		return true
	}

	if p.IsMove() {
		return true
	}

	for npcId, _ := range demandMap {
		//获取npc位置
		sceneTemplate := scenetemplate.GetSceneTemplateService().GetNPC(npcId)
		//npc不存在
		if sceneTemplate == nil {
			log.Warn("quest_guaji:npc不存在")
			return false
		}
		currentScene := p.GetScene()
		if currentScene != nil && currentScene.MapId() == sceneTemplate.SceneID {
			if coreutils.Distance(p.GetPosition(), sceneTemplate.GetPos()) <= float64(collectDistance) {
				questlogic.IncreaseQuestData(p, questtypes.QuestSubTypeCollect, npcId, 1)
				return true
			}
		}

		if !scenelogic.PlayerMoveThroughPortal(p, questTemplate.GetPortalTemplate(), sceneTemplate.SceneID, sceneTemplate.GetPos()) {
			log.WithFields(
				log.Fields{
					"npcId":     npcId,
					"questType": questTemplate.GetQuestSubType().String(),
				}).Warn("quest_guaji:找不到路")
			return false
		}

		return true
	}
	return false
}
