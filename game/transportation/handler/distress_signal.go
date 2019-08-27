package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/transportation/pbutil"
	"fgame/fgame/game/transportation/transpotation"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DISTRESS_SIGNAL_TYPE), dispatch.HandlerFunc(handlerDistressSignal))
}

//穿云箭请求
func handlerDistressSignal(s session.Session, msg interface{}) (err error) {
	log.Debug("transport:处理穿云箭请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = distressSignal(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("transport:处理穿云箭请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("transport:处理穿云箭请求完成")

	return
}

func distressSignal(pl player.Player) (err error) {
	err = transpotation.GetTransportService().DistressSignal(pl.GetId())
	if err != nil {
		return
	}

	//广播求救
	al := alliance.GetAllianceService().GetAlliance(pl.GetAllianceId())
	if al == nil {
		return
	}
	scDistressSignalBroadcast := pbutil.BuildSCDistressSignalBroadcast(pl.GetId())
	for _, memberObj := range al.GetMemberList() {
		if memberObj.GetMemberId() == pl.GetId() {
			continue
		}
		//TODO:xzk:过滤跨服玩家

		memberPlayer := player.GetOnlinePlayerManager().GetPlayerById(memberObj.GetMemberId())
		if memberPlayer == nil {
			continue
		}

		memberPlayer.SendMsg(scDistressSignalBroadcast)
	}

	scDistressSignal := pbutil.BuildSCDistressSignal()
	pl.SendMsg(scDistressSignal)

	return
}
