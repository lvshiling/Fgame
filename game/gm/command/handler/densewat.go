package handler

import (
	"fgame/fgame/common/lang"
	activitytemplate "fgame/fgame/game/activity/template"
	activetypes "fgame/fgame/game/activity/types"
	densewatlogic "fgame/fgame/game/densewat/logic"
	"fgame/fgame/game/global"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeEnterDenseWat, command.CommandHandlerFunc(handleEnterDenseWat))

}

func handleEnterDenseWat(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)

	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activetypes.ActivityTypeDenseWat)

	//是否满足开启条件
	if !pl.IsFuncOpen(activityTemplate.GetFuncOpenType()) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("activity:进入活动请求，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	//判断活动是否在活动时间内
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	flag, err := activityTemplate.IsAtActivityTime(now, openTime, mergeTime)
	if err != nil {
		return
	}

	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"now":      now,
			}).Warn("activity:进入活动请求，当前活动未开启")
		playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
		return
	}

	densewatlogic.PlayerEnterDenseWat(pl, activityTemplate)
	return
}
