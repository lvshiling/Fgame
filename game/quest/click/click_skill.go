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
	quest.RegisterClick(clicktypes.ClickTypeSkill, quest.HandlerFunc(handleClickSkill))
}

//处理技能点击事件
func handleClickSkill(pl player.Player, clickSubType clicktypes.ClickSubType) (err error) {
	log.Debug("quest:处理技能点击事件")

	switch clickSubType {
	case clicktypes.ClickSubTypeSkillUpgrade:
		{
			err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeClickSkillUpgradeButton, 0, 1)
			break
		}
	}

	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"clickSubType": clickSubType,
				"error":        err,
			}).Error("quest:处理技能点击事件,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":     pl.GetId(),
			"clickSubType": clickSubType,
		}).Debug("quest:处理技能点击事件,完成")
	return nil
}
