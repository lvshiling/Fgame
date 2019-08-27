package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/welfare/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_RANK_MY_GET_TYPE), dispatch.HandlerFunc(handleOpenActivityRankMyGet))
}

//处理我的排名信息
func handleOpenActivityRankMyGet(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取我的排名消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csOpenActivityRankMyGet := msg.(*uipb.CSOpenActivityRankMyGet)
	groupId := csOpenActivityRankMyGet.GetGroupId()

	err = openActivityRankMyGet(tpl, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"error":    err,
			}).Error("welfare:处理获取我的排名消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("welfare:处理获取我的排名消息完成")
	return nil

}

//获取我的排名界面信息的逻辑
func openActivityRankMyGet(pl player.Player, groupId int32) (err error) {
	pos := rank.GetRankService().GetMyRankPos(ranktypes.RankClassTypeLocalActivity, groupId, 0, pl.GetId())
	scRankMyGet := pbutil.BuildSCOpenActivityRankMyGet(groupId, pos)
	pl.SendMsg(scRankMyGet)
	return
}
