package check

import (
	playeradditionsys "fgame/fgame/game/additionsys/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	additionsystypes "fgame/fgame/game/additionsys/types"
	clicktypes "fgame/fgame/game/click/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeUpgradeSysOperation, quest.CheckHandlerFunc(handleUpgradeSysOperation))
}

//check 处理升级系统点击事件
func handleUpgradeSysOperation(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理升级系统点击事件")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for demandId, num := range questDemandMap {
		clickType := clicktypes.ClickSubTypeUpgradeSys(demandId)
		if !clickType.Valid() {
			return
		}
		sysType := additionsystypes.AdditionSysTypeMountEquip
		manager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
		switch clickType {
		case clicktypes.ClickSubTypeUpgradeSysMount:
			sysType = additionsystypes.AdditionSysTypeMountEquip
			break
		case clicktypes.ClickSubTypeUpgradeSysWing:
			sysType = additionsystypes.AdditionSysTypeWingStone
			break
		case clicktypes.ClickSubTypeUpgradeSysAnQi:
			sysType = additionsystypes.AdditionSysTypeAnqiJiguan
			break
		}

		levelInfo := manager.GetAdditionSysLevelInfoByType(sysType)
		shengJiTemplate := levelInfo.GetNextShengJiTemplate()
		if shengJiTemplate != nil {
			return
		}
		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), demandId, num)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理升级系统点击事件,完成")
	return nil
}
