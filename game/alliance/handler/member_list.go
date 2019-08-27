package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_MEMBER_LIST_TYPE), dispatch.HandlerFunc(handleAllianceMemberList))
}

//处理仙盟成员列表
func handleAllianceMemberList(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟成员列表")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = allianceMemberList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:处理仙盟成员列表,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理仙盟成员列表,完成")
	return nil

}

//仙盟成员列表
func allianceMemberList(pl player.Player) (err error) {
	allianceMem := alliance.GetAllianceService().GetAllianceMember(pl.GetId())
	if allianceMem == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:获取仙盟成员列表失败，不在仙盟中")
		return
	}

	memList := allianceMem.GetAlliance().GetMemberList()
	scAllianceJoinApplyList := pbutil.BuildSCAllianceMemberList(memList)
	pl.SendMsg(scAllianceJoinApplyList)
	return
}
