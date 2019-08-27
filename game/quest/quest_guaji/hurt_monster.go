package quest_guaji

import (
	coreutils "fgame/fgame/core/utils"
	guidereplicalogic "fgame/fgame/game/guidereplica/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/quest/guaji/guaji"
	questtypes "fgame/fgame/game/quest/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeHurtMonster, guaji.QuestGuaJiFunc(hurtMonster))
}

var (
	questDistance = 5
)

func hurtMonster(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {
	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		log.Warn("quest_guaji:任务需求是空")
		return false
	}
	//进入引导副本
	if questTemplate.GetGuideReplicaTemplate() != nil {
		//进入引导副本
		return guidereplicalogic.PlayerEnterQuestGuideReplica(p, int32(questTemplate.TemplateId()))
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

		s := p.GetScene()
		if s != nil && s.MapId() == sceneTemplate.SceneID {
			if coreutils.Distance(p.GetPosition(), sceneTemplate.GetPos()) > float64(s.MapTemplate().GetMapType().GetEnterDistance()) {
				return p.SetDestPosition(sceneTemplate.GetPos())
			}

			npcList := s.GetNPCListByBiology(npcId)
			if len(npcList) <= 0 {
				return false
			}
			var foundNpc scene.NPC
			for _, npc := range npcList {
				if !npc.IsDead() {
					foundNpc = npc
					break
				}
			}
			if foundNpc == nil {
				//等待复活
				foundNpc = npcList[0]
			}

			p.SetAttackTarget(foundNpc)
			p.GuaJiTrace()
			return true

		}

		if !scenelogic.PlayerMoveThroughPortal(p, questTemplate.GetPortalTemplate(), sceneTemplate.SceneID, sceneTemplate.GetPos()) {
			log.WithFields(
				log.Fields{
					"npcId": npcId,
				}).Warn("quest_guaji:找不到路")
			return false
		}

		return true
	}
	return false

}
