package ai

import (
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/global"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	robotlogic "fgame/fgame/game/robot/logic"
	"fgame/fgame/game/robot/robot"
	robottypes "fgame/fgame/game/robot/types"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"

	log "github.com/Sirupsen/logrus"
)

func init() {
	robot.RegisterActionFactory(robottypes.RobotTypeQuest, robot.RobotPlayerStateIdle, robot.RobotActionFactoryFunc(newIdleAction))
}

const (
	minQuestTime = 3 * common.SECOND
	maxQuestTime = 5 * common.SECOND
)

type idleAction struct {
	*robot.DummyAction
	currentQuestId int32
	//正在做任务
	doQuest       bool
	questTime     int64
	waitStartTime int64
	questDone     bool
	//任务全做完
	questAllDone bool
	//退出
	shouldExit bool
}

func (a *idleAction) Action(p scene.RobotPlayer) {

	//正在移动
	if p.IsMove() {
		return
	}

	if a.shouldExit {
		//退出
		robotlogic.RemoveRobot(p)
		return
	}

	//TODO:zrc 场景可能是空的
	if p.GetScene() == nil {
		return
	}

	//TODO 随机睡眠
	if a.doQuest {
		now := global.GetGame().GetTimeService().Now()
		if a.waitStartTime != 0 {
			elapse := now - a.waitStartTime
			if elapse > a.questTime {
				a.doQuest = false
				a.waitStartTime = 0
				a.questTime = 0
				a.questDone = true
			}
		} else {
			a.waitStartTime = now
		}
		return
	}

	//任务做完了
	if a.questAllDone {
		//移动
		sdkList := center.GetCenterService().GetSdkList()
		robotQuestTemplate := scenetemplate.GetSceneTemplateService().GetRobotQuestTemplate(sdkList, p.GetScene().MapId())
		if robotQuestTemplate == nil {
			a.shouldExit = true
			return
		}
		if robotQuestTemplate.GetPortalTemplate() == nil {
			a.shouldExit = true
			return
		}
		//获取传送阵
		portalSceneTemplate := scenetemplate.GetSceneTemplateService().GetPortalSceneTemplate(int32(robotQuestTemplate.GetPortalTemplate().TemplateId()))
		if portalSceneTemplate == nil {
			a.shouldExit = true
			return
		}

		p.SetDestPosition(portalSceneTemplate.GetPos())

		a.shouldExit = true
		return
	}

	//第一个任务
	if a.currentQuestId == 0 {
		a.currentQuestId = p.GetQuestBeginId()
	} else {
		if a.questDone {
			//任务结束
			if a.currentQuestId == p.GetQuestEndId() {
				a.questAllDone = true
				return
			}
			questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(a.currentQuestId)
			nextQuestIds := questTemplate.GetNextQuestIds()
			if len(nextQuestIds) == 0 {
				a.questAllDone = true
				return
			}
			//获取下一个任务
			a.currentQuestId = nextQuestIds[0]
			a.questDone = false
		}
	}

	//做任务
	finishQuest := true
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(a.currentQuestId)
	switch questTemplate.GetQuestSubType() {
	case questtypes.QuestSubTypeDialog:
		finishQuest = dialogQuest(p, questTemplate)
		if !finishQuest {
			a.questTime = int64(mathutils.RandomRange(int(minQuestTime), int(maxQuestTime)))
			a.doQuest = true
		} else {
			a.questDone = true
		}
		break
	case questtypes.QuestSubTypeCollect:
		finishQuest = collectQuest(p, questTemplate)
		if !finishQuest {
			a.questTime = int64(mathutils.RandomRange(int(minQuestTime), int(maxQuestTime)))
			a.doQuest = true
		} else {
			a.questDone = true
		}
		break
	case questtypes.QuestSubTypeHurtMonster:
		finishQuest = hurtMonster(p, questTemplate)
		if finishQuest {
			a.questDone = true
		}
		break
	case questtypes.QuestSubTypeKillMonster:
		finishQuest = hurtMonster(p, questTemplate)
		if finishQuest {
			a.questDone = true
		}
		break
	default:
		a.questDone = true
		break
	}

	return
}

func newIdleAction() scene.RobotAction {
	a := &idleAction{}
	a.DummyAction = robot.NewDummyAction()
	return a
}

func collectQuest(p scene.Player, questTemplate *gametemplate.QuestTemplate) bool {
	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		return true
	}

	for npcId, _ := range demandMap {
		//获取npc位置
		sceneTemplate := scenetemplate.GetSceneTemplateService().GetNPC(npcId)
		//npc不存在
		if sceneTemplate == nil {
			log.Warn("quest_guaji:npc不存在")
			return true
		}
		if !p.SetDestPosition(sceneTemplate.GetPos()) {
			continue
		}
		return false
	}
	return true
}

func dialogQuest(p scene.Player, questTemplate *gametemplate.QuestTemplate) bool {
	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		return true
	}

	for npcId, _ := range demandMap {
		//获取npc位置
		sceneTemplate := scenetemplate.GetSceneTemplateService().GetQuestNPC(npcId)
		//npc不存在
		if sceneTemplate == nil {
			log.Warn("quest_robot:npc不存在")
			return true
		}
		if !p.SetDestPosition(sceneTemplate.GetPos()) {
			continue
		}

		return false
	}
	return true
}

func hurtMonster(p scene.RobotPlayer, questTemplate *gametemplate.QuestTemplate) bool {
	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		return true
	}

	for npcId, _ := range demandMap {
		//获取npc位置
		sceneTemplate := scenetemplate.GetSceneTemplateService().GetNPC(npcId)
		//npc不存在
		if sceneTemplate == nil {
			log.Warn("quest_robot:npc不存在")
			return true
		}

		s := p.GetScene()
		if s != nil && s.MapId() == sceneTemplate.SceneID {

			if coreutils.Distance(p.GetPosition(), sceneTemplate.GetPos()) > float64(s.MapTemplate().GetMapType().GetEnterDistance()) {
				flag := p.SetDestPosition(sceneTemplate.GetPos())
				if !flag {
					log.WithFields(
						log.Fields{
							"id": sceneTemplate.Idx,
						}).Warn("quest_robot:找不到模板的路")
					return true
				}
				return false
			}

			npcList := s.GetNPCListByBiology(npcId)
			if len(npcList) <= 0 {
				return true
			}
			weights := make([]int64, 0, len(npcList))
			for _, _ = range npcList {
				weights = append(weights, 1)
			}
			index := mathutils.RandomWeights(weights)
			//随机
			foundNpc := npcList[index]
			if foundNpc.IsDead() {
				for _, npc := range npcList {
					if npc.IsDead() {
						continue
					}
					foundNpc = npc
					break
				}
			}
			if foundNpc == nil {
				return true
			}

			p.SetAttackTarget(foundNpc)
			p.Trace()
			return true
		}
		return true
	}
	return true

}
