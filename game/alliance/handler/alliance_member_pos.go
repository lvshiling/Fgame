package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_MEMBER_POS_TYPE), dispatch.HandlerFunc(handleAllianceMemberPos))
}

//处理盟友位置请求信息
func handleAllianceMemberPos(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理盟友位置请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = allianceMemberPos(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:处理盟友位置请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理盟友位置请求消息完成")
	return nil

}

//队友位置请求信息的逻辑
func allianceMemberPos(pl player.Player) (err error) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		return
	}

	scAllianceMemberPos := pbutil.BuildSCAllianceMemberPos(pl)
	pl.SendMsg(scAllianceMemberPos)
	return
}
