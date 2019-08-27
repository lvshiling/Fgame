package handler

import (
	"context"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
	emaillogic "fgame/fgame/game/email/logic"
	marrylogic "fgame/fgame/game/marry/logic"
	marryservice "fgame/fgame/game/marry/marry"
	pbutil "fgame/fgame/game/marry/pbutil"
	marryplayer "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MarryDingQingZengSong), dispatch.HandlerFunc(handleMarryDingQingZengSong))
}

type marryDingqingZengSong struct {
	SuitId int32
	PosId  int32
}

//定情信物赠送
func handleMarryDingQingZengSong(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:定情信物赠送")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	suoYaoInfo := msg.(*uipb.CSMarryDingQingZengSongMsg)
	suitId := suoYaoInfo.GetSuitId()
	posId := suoYaoInfo.GetPosId()
	marryDingQingZengSongDealSuccess(tpl, suitId, posId)
	log.Debug("marry:定情信物赠送,成功")
	return
}

//赠送
func marryDingQingZengSongDealSuccess(pl player.Player, suitId int32, posId int32) {
	item := marrytemplate.GetMarryTemplateService().GetMarryXinWuItem(suitId, posId)
	spId := pl.GetSpouseId()
	spl := player.GetOnlinePlayerManager().GetPlayerById(spId)
	if spl == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitId":   suitId,
				"posId":    posId,
			}).Warn("marry:定情信物赠送,配偶不在线")
		playerlogic.SendSystemMessage(pl, lang.MarryXinWuNotSpouseNotOnLine)
		return
	}

	//伴侣已经拥有信物
	if marryservice.GetMarryService().ExistsSpouseDingQing(pl.GetId(), suitId, posId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitId":   suitId,
				"posId":    posId,
			}).Warn("marry:定情信物赠送,配偶已经有订情信物")
		playerlogic.SendSystemMessage(pl, lang.MarryXinWuAlreadyExists)
		return
	}

	flag := autoBuyDingQingItem(pl, item.GetItemId(), 1)
	if !flag {
		return
	}

	//发送给伴侣更改数据
	postData := &marryDingqingZengSong{}
	postData.PosId = posId
	postData.SuitId = suitId
	splCtx := scene.WithPlayer(context.Background(), spl)
	msg := message.NewScheduleMessage(sendMarryDingQingSpouseZengSong, splCtx, postData, nil)
	spl.Post(msg)

	responMsg := pbutil.BuildSCMarryDingQingZengSongDealMsg(int32(1))
	pl.SendMsg(responMsg)
}

//赠送定情
func sendMarryDingQingSpouseZengSong(ctx context.Context, result interface{}, err error) error {
	info := result.(*marryDingqingZengSong)
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	//TODO:cjy 有可能已经有定情信物
	item := marrytemplate.GetMarryTemplateService().GetMarryXinWuItem(info.SuitId, info.PosId)
	if item == nil {
		return nil
	}

	marryManager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*marryplayer.PlayerMarryDataManager)
	if marryManager.ExistsDingQing(info.SuitId, info.PosId) { //如果存在
		xinWuMap := make(map[int32]int32)
		xinWuMap[item.GetItemId()] = 1
		existsTitle := lang.GetLangService().ReadLang(lang.MarryXinWuDingQingZengSongSuccess)
		//TODO:cjy 信物名字是空的
		existsContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryXinWuDingQingZengSongSuccessContent), item.GetXinWuName(), item.Attack)
		emaillogic.AddEmail(pl, existsTitle, existsContent, xinWuMap)
		return nil
	}
	marrylogic.AddPlayerDingQing(pl, info.SuitId, info.PosId)

	title := lang.GetLangService().ReadLang(lang.MarryXinWuDingQingZengSongSuccess)
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryXinWuDingQingZengSongSuccessContent), item.GetXinWuName(), item.Attack)
	emaillogic.AddEmail(pl, title, content, nil)

	return nil
}

func newMarryDingqingZengSongPostData(suitId int32, posId int32) *marryDingqingZengSong {
	rst := &marryDingqingZengSong{
		SuitId: suitId,
		PosId:  posId,
	}
	return rst
}
