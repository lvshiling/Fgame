package handler

import (
	"fgame/fgame/common/lang"
	activitylogic "fgame/fgame/game/activity/logic"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	yuxiscene "fgame/fgame/game/yuxi/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeAllianceDismiss, command.CommandHandlerFunc(handleAllianceDismiss))
}

func handleAllianceDismiss(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:仙盟解散")

	err = allianceDismiss(pl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Error("gm:创建仙盟,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:创建仙盟,完成")
	return
}

func allianceDismiss(p scene.Player) (err error) {
	pl := p.(player.Player)
	if activitylogic.IfActivityTime(activitytypes.ActivityTypeCoressTuLong) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:您所在的仙盟正在跨服屠龙,期间无法解散仙盟")
		playerlogic.SendSystemMessage(pl, lang.AllianceDismissInCrossTuLong)
		return
	}

	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeYuXi)
	yuXiScene, _ := yuxiscene.GetYuXiScene(activityTemplate)
	if yuXiScene != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:您所在的仙盟正在玉玺之战,期间无法解散仙盟")
		playerlogic.SendSystemMessage(pl, lang.AllianceDismissInYuXiWar)
		return
	}

	al, memList, err := alliance.GetAllianceService().DismissAlliance(pl.GetId())
	if err != nil {
		return
	}

	for _, member := range memList {
		if member.GetMemberId() == pl.GetId() {
			continue
		}

		memberPlayer := player.GetOnlinePlayerManager().GetPlayerById(member.GetMemberId())
		if memberPlayer != nil {
			dismissBroadcast := pbutil.BuildSCAllianceDismissBroadcast(al.GetAllianceId())
			memberPlayer.SendMsg(dismissBroadcast)
		}
	}

	scAllianceDismiss := pbutil.BuildSCAllianceDismiss(al.GetAllianceId())
	pl.SendMsg(scAllianceDismiss)
	return
}
