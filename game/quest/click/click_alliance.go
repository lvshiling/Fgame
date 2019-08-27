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
	quest.RegisterClick(clicktypes.ClickTypeAlliance, quest.HandlerFunc(handleClickAlliance))
}

//处理帝魂点击事件
func handleClickAlliance(pl player.Player, clickSubType clicktypes.ClickSubType) (err error) {
	log.Debug("quest:处理仙盟点击事件")

	switch clickSubType {
	case clicktypes.ClickSubTypeAllianceApplyJoin:
		{
			err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeApplyJoinAlliance, 0, 1)
			break
		}
	}
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"clickSubType": clickSubType,
				"error":        err,
			}).Error("quest:处理仙盟点击事件,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":     pl.GetId(),
			"clickSubType": clickSubType,
		}).Debug("quest:处理仙盟点击事件,完成")
	return nil
}
