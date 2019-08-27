package handler

import (
	"context"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
	pbutil "fgame/fgame/game/marry/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_DINGQING_YUE_TYPE), dispatch.HandlerFunc(handleMarryDingQingYue))
}

//定情信物约数据
type marryDingQingYuePostData struct {
	PlayerId int64
	SuitId   int32
	PosId    int32
}

func handleMarryDingQingYue(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:定情信物约")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csDingQing := msg.(*uipb.CSMarryDingQingYueMsg)
	suitId := csDingQing.GetSuitId()
	posId := csDingQing.GetPosId()
	playerId := tpl.GetId()

	spId := tpl.GetSpouseId()
	if spId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"suitId":   suitId,
				"posId":    posId,
			}).Warn("marry:定情信物索要,配偶不存在")
		playerlogic.SendSystemMessage(tpl, lang.MarryXinWuNotSpouse)
		return
	}
	spl := player.GetOnlinePlayerManager().GetPlayerById(spId)
	if spl == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitId":   suitId,
				"posId":    posId,
			}).Warn("marry:定情信物索要,配偶不在线")
		playerlogic.SendSystemMessage(tpl, lang.MarryXinWuNotSpouseNotOnLine)
		return
	}

	//发送伴侣
	posMsg := &marryDingQingYuePostData{
		PlayerId: playerId,
		SuitId:   suitId,
		PosId:    posId,
	}
	splCtx := scene.WithPlayer(context.Background(), spl)
	postMsg := message.NewScheduleMessage(marryDingQingYueSpouse, splCtx, posMsg, nil)
	spl.Post(postMsg)

	//回复信息
	rspMsg := pbutil.BuildSCMarryDingQingYueMsg()
	tpl.SendMsg(rspMsg)

	return
}

func marryDingQingYueSpouse(ctx context.Context, result interface{}, err error) error {
	info := result.(*marryDingQingYuePostData)
	msg := pbutil.BuildSCMarryDingQingYueSpouseMsg(info.PlayerId, info.SuitId, info.PosId)
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	pl.SendMsg(msg)
	return nil
}
