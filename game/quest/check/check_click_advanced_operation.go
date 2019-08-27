package check

import (
	playeranqi "fgame/fgame/game/anqi/player"
	playerfabao "fgame/fgame/game/fabao/player"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	playershenfa "fgame/fgame/game/shenfa/player"
	gametemplate "fgame/fgame/game/template"
	playerwing "fgame/fgame/game/wing/player"
	playerxianti "fgame/fgame/game/xianti/player"

	clicktypes "fgame/fgame/game/click/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeAdvancedOperation, quest.CheckHandlerFunc(handleAdvancedOperation))
}

//check 处理进阶点击事件
func handleAdvancedOperation(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理进阶点击事件")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for demandId, num := range questDemandMap {
		clickType := clicktypes.ClickSubTypeAdvanced(demandId)
		if !clickType.Valid() {
			return
		}
		isFull := false
		switch clickType {
		case clicktypes.ClickSubTypeAdvancedMount:
			manager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
			isFull = manager.IfFullAdvanced()
			break
		case clicktypes.ClickSubTypeAdvancedWing:
			manager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
			isFull = manager.IfFullAdvanced()
			break
		case clicktypes.ClickSubTypeAdvancedShenFa:
			manager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
			isFull = manager.IfFullAdvanced()
			break
		case clicktypes.ClickSubTypeAdvancedFaBao:
			manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
			isFull = manager.IfFullAdvanced()
			break
		case clicktypes.ClickSubTypeAdvancedXianTi:
			manager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
			isFull = manager.IfFullAdvanced()
			break
		case clicktypes.ClickSubTypeAdvancedAnQi:
			manager := pl.GetPlayerDataManager(types.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
			isFull = manager.IfFullAdvanced()
			break
		}
		if !isFull {
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
		}).Debug("quest:处理进阶点击事件,完成")
	return nil
}
