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
	quest.RegisterClick(clicktypes.ClickTypeEquip, quest.HandlerFunc(handleClickEquip))
}

//处理装备点击事件
func handleClickEquip(pl player.Player, clickSubType clicktypes.ClickSubType) (err error) {
	log.Debug("quest:处理装备点击事件")

	err = clickEquipmentStrengthenButton(pl, clickSubType)
	if err != nil {
		return
	}

	err = clickEquipmentUpgradeButton(pl, clickSubType)
	if err != nil {
		return
	}

	err = clickEquipmentUpgradeStarButton(pl, clickSubType)
	if err != nil {
		return
	}

	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"clickSubType": clickSubType,
				"error":        err,
			}).Error("quest:处理装备点击事件,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":     pl.GetId(),
			"clickSubType": clickSubType,
		}).Debug("quest:处理装备点击事件,完成")
	return nil
}

//一键强化或者强化装备X次
func clickEquipmentStrengthenButton(pl player.Player, clickSubType clicktypes.ClickSubType) (err error) {
	switch clickSubType {
	case clicktypes.ClickSubTypeEquipStrength:
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeClickEquipmentStrengthenButton, 0, 1)
	}
	return
}

//进阶装备X次
func clickEquipmentUpgradeButton(pl player.Player, clickSubType clicktypes.ClickSubType) (err error) {
	switch clickSubType {
	case clicktypes.ClickSubTypeEquipUpgrade:
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeClickEquipmentUpgradeButton, 0, 1)
	}
	return
}

//装备升星的次数
func clickEquipmentUpgradeStarButton(pl player.Player, clickSubType clicktypes.ClickSubType) (err error) {
	switch clickSubType {
	case clicktypes.ClickSubTypeEquipUpgradeStar:
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeClickEquipmentUpgradeStarButton, 0, 1)
	}
	return
}
