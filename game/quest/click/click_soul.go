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
	quest.RegisterClick(clicktypes.ClickTypeSoul, quest.HandlerFunc(handleClickSoul))
}

//处理帝魂点击事件
func handleClickSoul(pl player.Player, clickSubType clicktypes.ClickSubType) (err error) {
	log.Debug("quest:处理帝魂点击事件")

	err = clickSoulStrengthenButton(pl, clickSubType)
	if err != nil {
		return
	}

	err = clickSoulUpgradeButton(pl, clickSubType)
	if err != nil {
		return
	}
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"clickSubType": clickSubType,
				"error":        err,
			}).Error("quest:处理帝魂点击事件,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":     pl.GetId(),
			"clickSubType": clickSubType,
		}).Debug("quest:处理帝魂点击事件,完成")
	return nil
}

//强化帝魂X次
func clickSoulStrengthenButton(pl player.Player, clickSubType clicktypes.ClickSubType) (err error) {
	switch clickSubType {
	case clicktypes.ClickSubTypeSoulStrengthen:
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeClickSoulStrengthenButton, 0, 1)
	}
	return
}

//魂技升级X次
func clickSoulUpgradeButton(pl player.Player, clickSubType clicktypes.ClickSubType) (err error) {
	switch clickSubType {
	case clicktypes.ClickSubTypeSoulUpgrade:
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeClickSoulUpgradeButton, 0, 1)
	}
	return
}
