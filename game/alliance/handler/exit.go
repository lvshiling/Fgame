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
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_EXIT_TYPE), dispatch.HandlerFunc(handleAllianceExit))
}

//处理仙盟退出
func handleAllianceExit(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟退出")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = allianceExit(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),

				"error": err,
			}).Error("alliance:处理仙盟退出,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理仙盟退出,完成")
	return nil

}

//仙盟退出
func allianceExit(pl player.Player) (err error) {
	if activitylogic.IfActivityTime(activitytypes.ActivityTypeCoressTuLong) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:您所在的仙盟正在跨服屠龙,期间无法退盟")
		playerlogic.SendSystemMessage(pl, lang.AllianceExitInCrossTuLong)
		return
	}

	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeYuXi)
	yuXiScene, _ := yuxiscene.GetYuXiScene(activityTemplate)
	if yuXiScene != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:您所在的仙盟正在玉玺之战,期间无法退盟")
		playerlogic.SendSystemMessage(pl, lang.AllianceExitInYuXiWar)
		return
	}

	al, err := alliance.GetAllianceService().ExitAlliance(pl.GetId())
	if err != nil {
		return
	}

	//广播帮派
	format := lang.GetLangService().ReadLang(lang.AllianceExitNotice)
	content := fmt.Sprintf(format, coreutils.FormatColor(alliancetypes.ColorTypeLogName, pl.GetName()))
	chatlogic.SystemBroadcastAlliance(al, chattypes.MsgTypeText, []byte(content))

	scAllianceExit := pbutil.BuildSCAllianceExit()
	pl.SendMsg(scAllianceExit)
	return
}
