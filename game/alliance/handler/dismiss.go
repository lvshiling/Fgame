package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	activitylogic "fgame/fgame/game/activity/logic"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	yuxiscene "fgame/fgame/game/yuxi/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_DISMISS_TYPE), dispatch.HandlerFunc(handleAllianceDismiss))
}

//处理仙盟解散
func handleAllianceDismiss(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟解散")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = allianceDismiss(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),

				"error": err,
			}).Error("alliance:处理仙盟解散,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理仙盟解散,完成")
	return nil

}

//仙盟解散
func allianceDismiss(pl player.Player) (err error) {
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

	//解散广播
	for _, member := range memList {
		if member.GetMemberId() == pl.GetId() {
			continue
		}

		memberPlayer := player.GetOnlinePlayerManager().GetPlayerById(member.GetMemberId())
		if memberPlayer == nil {
			continue
		}

		dismissBroadcast := pbutil.BuildSCAllianceDismissBroadcast(al.GetAllianceId())
		memberPlayer.SendMsg(dismissBroadcast)
	}

	scAllianceDismiss := pbutil.BuildSCAllianceDismiss(al.GetAllianceId())
	pl.SendMsg(scAllianceDismiss)
	return
}
