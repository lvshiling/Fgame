package handler

import (
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeTuMoFinishTar, command.CommandHandlerFunc(handleTuMoFinishTar))
}

//任务栏的任务置完成
func handleTuMoFinishTar(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:处理屠魔任务栏任务值完成")

	err = finishTuMoTar(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理屠魔任务栏任务值完成,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理屠魔任务栏任务值完成完成")
	return
}

//任务栏的任务置完成
func finishTuMoTar(p scene.Player) (err error) {
	pl := p.(player.Player)
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)

	manager.GMFinishTuMoTar()

	questlogic.CheckInitQuestList(pl)
	questlogic.CheckAcceptQuestList(pl)

	questlogic.CheckCommitQuestList(pl)
	questList := make([]*playerquest.PlayerQuestObject, 0, 16)
	activeQuests := manager.GetQuestMap(questtypes.QuestStateActive)
	for _, quest := range activeQuests {
		questList = append(questList, quest)
	}
	acceptQuests := manager.GetQuestMap(questtypes.QuestStateAccept)
	for _, quest := range acceptQuests {
		questList = append(questList, quest)
	}
	finishQuests := manager.GetQuestMap(questtypes.QuestStateFinish)
	for _, quest := range finishQuests {
		questList = append(questList, quest)
	}

	finish := false
	mainQuestId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMainQuestId)
	mainQuest := manager.GetQuestById(mainQuestId)
	if mainQuest != nil && mainQuest.QuestState == questtypes.QuestStateCommit {
		finish = true
	}

	scQuestList := pbutil.BuildSCQuestList(questList, finish)
	pl.SendMsg(scQuestList)
	return nil
}
