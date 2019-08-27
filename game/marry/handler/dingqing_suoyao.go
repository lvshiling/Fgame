package handler

import (
	"context"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
	"fgame/fgame/game/global"
	pbutil "fgame/fgame/game/marry/pbutil"
	marryplayer "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MarryDingQingSuoYao), dispatch.HandlerFunc(handleMarryDingQingSuoYao))
}

//索要发给伴侣的post数据结构
type marryDingQingSuoYaoRsp struct {
	SuitId     int32
	PosId      int32
	PlayerId   int64
	PlayerName string
	Content    string
}

//定情信物索要
func handleMarryDingQingSuoYao(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:定情信物索要")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csDingQing := msg.(*uipb.CSMarryDingQingSuoYaoMsg)
	suitId := csDingQing.GetSuitId()
	posId := csDingQing.GetPosId()
	content := csDingQing.GetContent()
	suoYaoDingQing(tpl, suitId, posId, content)
	log.Debug("marry:定情信物索要,成功")
	return
}

func suoYaoDingQing(pl player.Player, suitId int32, posId int32, content string) {
	wedManager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*marryplayer.PlayerMarryDataManager)

	cdTime := marrytemplate.GetMarryTemplateService().GetMarryConstTempalte().XinwuCd
	now := global.GetGame().GetTimeService().Now()
	wedTime := wedManager.GetXinWuTime()
	if wedTime+cdTime > now {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitId":   suitId,
				"posId":    posId,
			}).Warn("marry:cding")
		playerlogic.SendSystemMessage(pl, lang.MarryXinWuSuoYaoCd)
		return
	}

	flag := wedManager.ExistsDingQing(suitId, posId)
	if flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitId":   suitId,
				"posId":    posId,
			}).Warn("marry:定情信物索要,已经有订情信物了")
		playerlogic.SendSystemMessage(pl, lang.MarryXinWuAlreadyExists)
		return
	}

	spId := pl.GetSpouseId()
	if spId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitId":   suitId,
				"posId":    posId,
			}).Warn("marry:定情信物索要,配偶不存在")
		playerlogic.SendSystemMessage(pl, lang.MarryXinWuNotSpouse)
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
		playerlogic.SendSystemMessage(pl, lang.MarryXinWuNotSpouseNotOnLine)
		return
	}
	//TODO:cjy 结果好像没什么用
	responMsg := pbutil.BuildSCMarryDingQingSuoYaoMsg(int32(1))
	pl.SendMsg(responMsg)
	//发送给伴侣索要物品
	sendSpouseSuoYao(spl, suitId, posId, pl.GetId(), pl.GetName(), content)
	wedManager.SetXinWuTime(now)
}

func sendSpouseSuoYao(pl player.Player, suitId int32, posId int32, playerId int64, playerName string, content string) {
	splCtx := scene.WithPlayer(context.Background(), pl)
	postData := &marryDingQingSuoYaoRsp{
		SuitId:     suitId,
		PosId:      posId,
		PlayerId:   playerId,
		PlayerName: playerName,
		Content:    content,
	}
	msg := message.NewScheduleMessage(sendSpouseMsg, splCtx, postData, nil)
	pl.Post(msg)
}

func sendSpouseMsg(ctx context.Context, result interface{}, err error) error {
	info := result.(*marryDingQingSuoYaoRsp)
	msg := pbutil.BuildSCMarryDingQingSuoYaoRspMsg(info.PlayerId, info.PlayerName, info.SuitId, info.PosId, info.Content)
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	pl.SendMsg(msg)
	return nil
}
