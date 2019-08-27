package click

import (
	clicktypes "fgame/fgame/game/click/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterClick(clicktypes.ClickTypeAdvanced, quest.HandlerFunc(handleClickAdvanced))
}

//处理进阶点击事件
func handleClickAdvanced(pl player.Player, clickSubType clicktypes.ClickSubType) (err error) {
	log.Debug("quest:处理进阶点击事件")

	switch clickSubType {
	case clicktypes.ClickSubTypeAdvancedMount,
		clicktypes.ClickSubTypeAdvancedWing,
		clicktypes.ClickSubTypeAdvancedShenFa,
		clicktypes.ClickSubTypeAdvancedFaBao,
		clicktypes.ClickSubTypeAdvancedXianTi,
		clicktypes.ClickSubTypeAdvancedAnQi:
		{
			err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeAdvancedOperation, int32(clickSubType.SubType()), 1)
			break
		}
	}

	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"clickSubType": clickSubType,
				"error":        err,
			}).Error("quest:处理进阶点击事件,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":     pl.GetId(),
			"clickSubType": clickSubType,
		}).Debug("quest:处理进阶点击事件,完成")
	return nil
}
