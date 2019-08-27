package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	activitylogic "fgame/fgame/game/activity/logic"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	alliancetypes "fgame/fgame/game/alliance/types"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	yuxiscene "fgame/fgame/game/yuxi/scene"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_KICK_TYPE), dispatch.HandlerFunc(handleAllianceKick))
}

//处理仙盟踢人
func handleAllianceKick(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟踢人")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csAllianceKick := msg.(*uipb.CSAllianceKick)
	kickMemberId := csAllianceKick.GetKickMemberId()
	err = allianceKick(tpl, kickMemberId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),

				"error": err,
			}).Error("alliance:处理仙盟踢人,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理仙盟踢人,完成")
	return nil

}

//仙盟踢人
func allianceKick(pl player.Player, kickMemberId int64) (err error) {
	if activitylogic.IfActivityTime(activitytypes.ActivityTypeCoressTuLong) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"kickMemberId": kickMemberId,
			}).Warn("alliance:您的仙盟正在跨服屠龙,期间无法踢人")
		playerlogic.SendSystemMessage(pl, lang.AllianceTickInCrossTuLong)
		return
	}

	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeYuXi)
	yuXiScene, _ := yuxiscene.GetYuXiScene(activityTemplate)
	if yuXiScene != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:您所在的仙盟正在玉玺之战,期间无法踢人")
		playerlogic.SendSystemMessage(pl, lang.AllianceTickInYuXiWar)
		return
	}

	mem, kickMem, err := alliance.GetAllianceService().Kick(pl.GetId(), kickMemberId)
	if err != nil {
		return
	}

	//广播帮派
	format := lang.GetLangService().ReadLang(lang.AllianceKickNotice)
	beKickName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, kickMem.GetName())
	kickName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, mem.GetName())
	content := fmt.Sprintf(format, beKickName, kickName)
	chatlogic.SystemBroadcastAlliance(mem.GetAlliance(), chattypes.MsgTypeText, []byte(content))

	//踢人通知
	kickPlayer := player.GetOnlinePlayerManager().GetPlayerById(kickMemberId)
	if kickPlayer != nil {
		kickNotice := pbutil.BuildSCAllianceKickNotice(pl.GetId(), pl.GetName())
		kickPlayer.SendMsg(kickNotice)
	}

	scAllianceKick := pbutil.BuildSCAllianceKick(kickMemberId, kickMem.GetName())
	pl.SendMsg(scAllianceKick)
	return
}
