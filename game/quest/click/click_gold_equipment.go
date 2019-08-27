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
	quest.RegisterClick(clicktypes.ClickTypeGoldEquipment, quest.HandlerFunc(handleClickGoldEquipmentStrength))
}

//处理元神金装强化点击事件
func handleClickGoldEquipmentStrength(pl player.Player, clickSubType clicktypes.ClickSubType) (err error) {
	log.Debug("quest:处理元神金装强化点击事件")

	switch clickSubType {
	case clicktypes.ClickSubTypeGoldEquipmentWeapon,
		clicktypes.ClickSubTypeGoldEquipmentClothes,
		clicktypes.ClickSubTypeGoldEquipmentHelmet,
		clicktypes.ClickSubTypeGoldEquipmentCaliga,
		clicktypes.ClickSubTypeGoldEquipmentBelt,
		clicktypes.ClickSubTypeGoldEquipmentHand,
		clicktypes.ClickSubTypeGoldEquipmentJade,
		clicktypes.ClickSubTypeGoldEquipmentNecklace:
		{
			err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeGoldEquipmentStrength, int32(clickSubType.SubType()), 1)
			if err != nil {
				return
			}
			err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeGoldEquipmentStrength, int32(clicktypes.ClickSubTypeGoldEquipmentNecklace)+1, 1)
			break
		}
	}

	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"clickSubType": clickSubType,
				"error":        err,
			}).Error("quest:处理元神金装强化点击事件,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":     pl.GetId(),
			"clickSubType": clickSubType,
		}).Debug("quest:处理元神金装强化点击事件,完成")
	return nil
}
