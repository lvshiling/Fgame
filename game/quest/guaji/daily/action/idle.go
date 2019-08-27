package action

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/guaji/guaji"
	questlogic "fgame/fgame/game/quest/logic"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeDailyQuest, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
	questId int32
}

func (a *idleAction) OnEnter() {
	a.questId = 0
}

//任务挂机中
func (a *idleAction) GuaJi(p scene.Player) {
	pl, ok := p.(player.Player)
	if !ok {
		log.Info("quest:挂机者不是真实玩家")
		p.ExitGuaJi()
		return
	}
	m := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	if a.questId == 0 {
		q := m.GetCurrentDailyQuest()
		if q == nil {
			log.Info("quest:没有日环任务了")
			p.ExitGuaJi()
			return
		}
		a.questId = q.QuestId
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"questId":  a.questId,
				"times":    m.GetDailyObj(questtypes.QuestDailyTagPerson).GetTimes(),
			}).Info("quest:正在做日环任务")
	}
	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(a.questId)
	if questTemplate == nil {
		log.WithFields(
			log.Fields{
				"questId": a.questId,
			}).Info("quest:日环任务不存在")
		p.ExitGuaJi()
		return
	}
	q := m.GetQuestById(a.questId)
	if q == nil {
		log.WithFields(
			log.Fields{
				"questId": a.questId,
			}).Info("quest:日环任务不存在")
		p.ExitGuaJi()
		return
	}

	//	double := false
	//	guaJiManager := pl.GetPlayerDataManager(playertypes.PlayerGuaJiManagerType).(*playerguaji.PlayerGuaJiManager)
	//	guaJiData, _ := guaJiManager.GetCurrentGuaJiType()

	//	if guaJiData.GetType() == guajitypes.GuaJiTypeDailyQuest {
	//		val := guaJiData.GetOptionValue(guajitypes.GuaJiTypeDailyQuestOptionDouble)
	//		if val != 0 {
	//			double = true
	//		}
	//	}

	//已经接受
	switch q.QuestState {
	case questtypes.QuestStateAccept:
		{
			guaJi := guaji.GetQuestGuaJi(questTemplate.GetQuestSubType())
			if guaJi == nil {
				log.WithFields(
					log.Fields{
						"questId": a.questId,
						"name":    questTemplate.Name,
					}).Info("quest:日环挂机不存在,无法完成")
				return
			}

			//判断是否是已经
			flag := guaJi.DoQuest(pl, questTemplate)
			if !flag {
				log.WithFields(
					log.Fields{
						"questId": a.questId,
					}).Info("quest:处理任务挂机失败")
				p.ExitGuaJi()
				return
			}
			break
		}
	case questtypes.QuestStateFinish:
		{
			//判断是否是手动完成
			if !questTemplate.AutoCommit() {
				//判断选项
				//	if double {
				//		if questlogic.IfQuestCanCommitDouble(pl, a.questId) {
				//			questlogic.CommitQuest(pl, a.questId, true)
				//		}
				//	}
				questlogic.CommitQuest(pl, a.questId, false)
			}
			break
		}
	case questtypes.QuestStateCommit:
		{
			a.questId = 0
			break
		}
	}
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
