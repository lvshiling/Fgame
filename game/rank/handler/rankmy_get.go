package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/rank/pbutil"
	"fgame/fgame/game/rank/rank"
	"fgame/fgame/game/rank/types"
	ranktypes "fgame/fgame/game/rank/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RANK_MY_GET_TYPE), dispatch.HandlerFunc(handleRankMyGet))
}

//处理我的排名信息
func handleRankMyGet(s session.Session, msg interface{}) (err error) {
	log.Debug("rank:处理获取我的排名消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csRankMyGet := msg.(*uipb.CSRankMyGet)
	typ := csRankMyGet.GetRankType()
	isArea := csRankMyGet.GetIsArea()

	err = rankMyGet(tpl, typ, isArea)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"isArea":   isArea,
				"error":    err,
			}).Error("rank:处理获取我的排名消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("rank:处理获取我的排名消息完成")
	return nil

}

//获取我的排名界面信息的逻辑
func rankMyGet(pl player.Player, typ int32, isArea bool) (err error) {
	myRankType := types.MyRankReqType(typ)
	if !myRankType.Valid() {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
			"typ":      typ,
		}).Warn("rank:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	rankType := myRankType.GetRankType()
	var classType ranktypes.RankClassType
	if isArea {
		classType = ranktypes.RankClassTypeArea
	} else {
		classType = ranktypes.RankClassTypeLocal
	}

	id := pl.GetId()
	if rankType == types.RankTypeGang {
		id = pl.GetAllianceId()
	}
	pos := rank.GetRankService().GetMyRankPos(classType, 0, rankType, id)
	scRankMyGet := pbutil.BuildSCRankMyGet(isArea, typ, pos)
	pl.SendMsg(scRankMyGet)
	return
}
