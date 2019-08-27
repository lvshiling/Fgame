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
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_CHARM_TYPE), dispatch.HandlerFunc(handleAllianceCharm))
}

//处理仙盟一键招人
func handleAllianceCharm(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟一键招人")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = allianceCharm(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),

				"error": err,
			}).Error("alliance:处理仙盟一键招人,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理仙盟一键招人,完成")
	return nil

}

//仙盟一键招人
func allianceCharm(pl player.Player) (err error) {
	//TODO cd限制
	mem := alliance.GetAllianceService().GetAllianceMember(pl.GetId())
	if mem == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Debug("alliance:处理仙盟一键招人,不在仙盟")
		playerlogic.SendSystemMessage(pl, lang.AllianceUserNotInAlliance)
		return
	}
	if mem.IsPositionMember() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Debug("alliance:处理仙盟一键招人,权限不够")
		playerlogic.SendSystemMessage(pl, lang.AllianceUserPrivilegeNoEnough)
		return
	}

	//广播
	format := lang.GetLangService().ReadLang(lang.AllianceCharmSystemNotice)
	allianceColorName := coreutils.FormatColor(alliancetypes.ColorTypeCharmAllianceName, mem.GetAlliance().GetAllianceName())
	params := []int64{int64(chattypes.ChatLinkTypeXianMeng), mem.GetAlliance().GetAllianceId()}
	joinLink := coreutils.FormatLink(chattypes.ButtonTypeToJoin, params)
	content := fmt.Sprintf(format, allianceColorName, joinLink)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))

	scAllianceCharm := pbutil.BuildSCAllianceCharm()
	pl.SendMsg(scAllianceCharm)
	return
}
