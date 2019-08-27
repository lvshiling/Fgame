package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	alliancetypes "fgame/fgame/game/alliance/types"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_COMMIT_TYPE), dispatch.HandlerFunc(handleAllianceCommit))
}

//处理仙盟任命
func handleAllianceCommit(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟任命")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csAllianceCommit := msg.(*uipb.CSAllianceCommit)
	commitMemberId := csAllianceCommit.GetCommitMemberId()
	position := csAllianceCommit.GetPosition()
	pos := alliancetypes.AlliancePosition(position)
	if !pos.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"commitMemberId": commitMemberId,
				"pos":            pos,
			}).Warn("alliance:处理仙盟任命,参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	err = allianceCommit(tpl, commitMemberId, pos)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"commitMemberId": commitMemberId,
				"pos":            pos,
				"error":          err,
			}).Error("alliance:处理仙盟任命,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":       pl.GetId(),
			"commitMemberId": commitMemberId,
			"pos":            pos,
		}).Debug("alliance:处理仙盟任命,完成")
	return nil

}

//仙盟任命
func allianceCommit(pl player.Player, commitMemberId int64, pos alliancetypes.AlliancePosition) (err error) {
	mem, commitMem, err := alliance.GetAllianceService().Commit(pl.GetId(), commitMemberId, pos)
	if err != nil {
		return
	}

	//广播帮派
	format := lang.GetLangService().ReadLang(lang.AllianceCommitNotice)
	commitName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, mem.GetName())
	beCommitName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, commitMem.GetName())
	content := fmt.Sprintf(format, commitName, beCommitName, pos)
	chatlogic.SystemBroadcastAlliance(mem.GetAlliance(), chattypes.MsgTypeText, []byte(content))

	//任命通知
	commitPlayer := player.GetOnlinePlayerManager().GetPlayerById(commitMemberId)
	if commitPlayer != nil {

		commitNotice := pbutil.BuildSCAllianceCommitNotice(pl.GetId(), pl.GetName(), pos)
		commitPlayer.SendMsg(commitNotice)
	}

	scAllianceCommit := pbutil.BuildSCAllianceCommit(commitMemberId, pos)
	pl.SendMsg(scAllianceCommit)

	return
}
