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
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_TRANSFER_TYPE), dispatch.HandlerFunc(handleAllianceTransfer))
}

//处理仙盟转让
func handleAllianceTransfer(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟转让")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csAllianceTransfer := msg.(*uipb.CSAllianceTransfer)
	transferMemberId := csAllianceTransfer.GetTransferMemberId()

	err = allianceTransfer(tpl, transferMemberId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"transferMemberId": transferMemberId,
				"error":            err,
			}).Error("alliance:处理仙盟转让,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":         pl.GetId(),
			"transferMemberId": transferMemberId,
		}).Debug("alliance:处理仙盟转让,完成")
	return nil

}

//仙盟转让
func allianceTransfer(pl player.Player, transferMemberId int64) (err error) {
	if activitylogic.IfActivityTime(activitytypes.ActivityTypeCoressTuLong) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:您所在的仙盟正在跨服屠龙,期间无法转让盟主")
		playerlogic.SendSystemMessage(pl, lang.AllianceTransferInCrossTuLong)
		return
	}

	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeYuXi)
	yuXiScene, _ := yuxiscene.GetYuXiScene(activityTemplate)
	if yuXiScene != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:您所在的仙盟正在玉玺之战,期间无法转让盟主")
		playerlogic.SendSystemMessage(pl, lang.AllianceTransferInYuXiWar)
		return
	}

	mem, transMem, err := alliance.GetAllianceService().Transfer(pl.GetId(), transferMemberId)
	if err != nil {
		return
	}

	//广播帮派
	format := lang.GetLangService().ReadLang(lang.AllianceTransferNotice)
	transfName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, mem.GetName())
	beTransfName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, transMem.GetName())
	content := fmt.Sprintf(format, transfName, beTransfName)
	chatlogic.SystemBroadcastAlliance(mem.GetAlliance(), chattypes.MsgTypeText, []byte(content))

	//推送仙盟所有在线玩家
	memberList := mem.GetAlliance().GetMemberList()
	for _, member := range memberList {
		if member.GetMemberId() == pl.GetId() {
			continue
		}

		memberPlayer := player.GetOnlinePlayerManager().GetPlayerById(member.GetMemberId())
		if memberPlayer != nil {
			transferBroadcast := pbutil.BuildSCAllianceTransferBroadcast(pl.GetId(), transferMemberId, mem.GetName(), transMem.GetName())
			memberPlayer.SendMsg(transferBroadcast)
		}
	}

	scAllianceTransfer := pbutil.BuildSCAllianceTransfer(transferMemberId)
	pl.SendMsg(scAllianceTransfer)

	return
}
