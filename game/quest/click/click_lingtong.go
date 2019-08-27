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
	quest.RegisterClick(clicktypes.ClickTypeLingTong, quest.HandlerFunc(handleClickLingTong))
}

//处理灵童点击事件
func handleClickLingTong(pl player.Player, clickSubType clicktypes.ClickSubType) (err error) {
	log.Debug("quest:处理灵童点击事件")

	err = clickLingTongShengJiButton(pl, clickSubType)
	if err != nil {
		return
	}

	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"clickSubType": clickSubType,
				"error":        err,
			}).Error("quest:处理灵童点击事件,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":     pl.GetId(),
			"clickSubType": clickSubType,
		}).Debug("quest:处理灵童点击事件,完成")
	return nil
}

//灵童升级x次
func clickLingTongShengJiButton(pl player.Player, clickSubType clicktypes.ClickSubType) (err error) {
	switch clickSubType {
	case clicktypes.ClickSubTypeLingTongShengJi:
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeLingTongClick, int32(clickSubType.SubType()), 1)
	}
	return
}
