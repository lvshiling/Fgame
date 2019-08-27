package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/friend/friend"
	"fgame/fgame/game/global"
	marryentity "fgame/fgame/game/marry/entity"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"math"
)

type marryDivorceDealReturn struct {
	Gold     int64
	BindGold int64
	Silver   int64
}

//玩家离婚决策
func playerMarryDivorceDeal(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*marryeventtypes.MarryDivorceDealEventData)
	if !ok {
		return
	}
	agree := eventData.GetAgree()
	spouseName := eventData.GetSpouseName()

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	spouseId := eventData.GetSpouseId()
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)

	percent := marrytemplate.GetMarryTemplateService().GetMarryDivorceLeftIntimacy()
	if agree {
		scMarryInfoStatusChange := pbuitl.BuildSCMarryInfoStatusChange(int32(marrytypes.MarryStatusTypeDivorce))
		manager.Divorce()
		pl.SendMsg(scMarryInfoStatusChange)
		friend.GetFriendService().DivorceSubPoint(pl, spouseId, marrytypes.MarryDivorceTypeConsent, percent)

		if spl != nil {
			splCtx := scene.WithPlayer(context.Background(), spl)
			divorceConsentAgreeMsg := message.NewScheduleMessage(onDivorceConsentArgee, splCtx, pl.GetId(), nil)
			spl.Post(divorceConsentAgreeMsg)
		} else {
			//写协议离婚成功离线日志
			if err = consentSucessOffonline(spouseId); err != nil {
				return
			}
		}

		//协议离婚跑马灯
		marryTemplate := marrytemplate.GetMarryTemplateService().GetMarryConstTempalte()
		if marryTemplate == nil {
			return
		}
		percent := float64(marryTemplate.DivorceQinmidu) / float64(common.MAX_RATE)
		percentInt := int32(math.Ceil(percent * float64(100)))
		name := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
		peerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(spouseName))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryDivorceConsent), name, peerName, percentInt)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
		noticelogic.NoticeNumBroadcast([]byte(content), marrytypes.WeddingInterval, 3)
		//离婚后通知，定情信物通知
		marrylogic.MarryPlayerDingQingPropertyChanged(pl)
		//离婚后回退各个预定的结婚
		returnPreMarry := eventData.GetPreMarryList()
		if len(returnPreMarry) > 0 {
			returnMarryPreList(returnPreMarry, pl, spl, spouseId)
		}
	} else {
		if spl == nil {
			return
		}
		scMarryDivorceDealPushPeer := pbuitl.BuildSCMarryDivorceDealPushPeer(pl.GetName())
		spl.SendMsg(scMarryDivorceDealPushPeer)
	}
	return
}

func consentSucessOffonline(playerId int64) (err error) {
	//玩家离线
	id, err := idutil.GetId()
	if err != nil {
		return err
	}

	now := global.GetGame().GetTimeService().Now()
	marryDivorceConsentEntity := &marryentity.MarryDivorceConsentEntity{
		Id:         id,
		ServerId:   global.GetGame().GetServerIndex(),
		PlayerId:   playerId,
		UpdateTime: now,
		CreateTime: now,
		DeleteTime: 0,
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(marryDivorceConsentEntity)
	return
}

func onDivorceConsentArgee(ctx context.Context, result interface{}, err error) error {

	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)

	scMarryInfoStatusChange := pbuitl.BuildSCMarryInfoStatusChange(int32(marrytypes.MarryStatusTypeDivorce))
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.Divorce()
	pl.SendMsg(scMarryInfoStatusChange)
	//离婚后通知，定情信物通知
	// marrylogic.MarryPlayerDingQingPropertyChanged(pl)
	return nil
}

func returnMarryPreList(returnPreMarry []*marryeventtypes.MarryDivorceDealEventDataItem, pl player.Player, spl player.Player, spouseId int64) {
	for _, value := range returnPreMarry {
		grade := value.GetMarryBanquetSubTypeWed()
		hunCheGrade := value.GetMarryBanquetSubTypeHunChe()
		sugarGrade := value.GetMarryBanquetSubTypeSugar()
		costBindGold, costGold, costSilver := marrytemplate.GetMarryTemplateService().GetMarryGradeCost(grade, hunCheGrade, sugarGrade)
		if value.GetReturnPlayerId() == pl.GetId() {
			if costGold == 0 && costBindGold == 0 && costSilver == 0 { //发送邮件回退东西
				continue
			}
			sendPreMarryReturnEmail(pl, int32(costGold), int32(costBindGold), int32(costSilver))
		} else {
			if costBindGold == 0 && costGold == 0 && costSilver == 0 {
				continue
			}

			if spl == nil { //发送离线邮件
				sendPreMarryReturnEmailOffline(spouseId, int32(costGold), int32(costBindGold), int32(costSilver))
			} else {
				splCtx := scene.WithPlayer(context.Background(), spl)
				returnEventData := &marryDivorceDealReturn{
					Gold:     int64(costGold),
					BindGold: int64(costBindGold),
					Silver:   int64(costSilver),
				}
				returnMsg := message.NewScheduleMessage(onPreMarryReturn, splCtx, returnEventData, nil)
				spl.Post(returnMsg)
			}
		}
	}
}

func onPreMarryReturn(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	data := result.(*marryDivorceDealReturn)
	sendPreMarryReturnEmail(pl, int32(data.Gold), int32(data.BindGold), int32(data.Silver))
	return nil
}

func sendPreMarryReturnEmail(pl player.Player, returnGold int32, returnBindGold int32, returnSilver int32) {
	emailItemMap := make(map[int32]int32)
	if returnGold > 0 {
		emailItemMap[constanttypes.GoldItem] = int32(returnGold)
	}
	if returnBindGold > 0 {
		emailItemMap[constanttypes.BindGoldItem] = int32(returnBindGold)
	}
	if returnSilver > 0 {
		emailItemMap[constanttypes.SilverItem] = int32(returnSilver)
	}
	title := lang.GetLangService().ReadLang(lang.MarryPreMarryReturn)
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryPreMarryReturnContent))
	emaillogic.AddEmail(pl, title, content, emailItemMap)
}

func sendPreMarryReturnEmailOffline(playerId int64, returnGold int32, returnBindGold int32, returnSilver int32) {
	emailItemMap := make(map[int32]int32)
	if returnGold > 0 {
		emailItemMap[constanttypes.GoldItem] = int32(returnGold)
	}
	if returnBindGold > 0 {
		emailItemMap[constanttypes.BindGoldItem] = int32(returnBindGold)
	}
	if returnSilver > 0 {
		emailItemMap[constanttypes.SilverItem] = int32(returnSilver)
	}
	title := lang.GetLangService().ReadLang(lang.MarryPreMarryReturn)
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryPreMarryReturnContent))
	emaillogic.AddOfflineEmail(playerId, title, content, emailItemMap)
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryDivorceDeal, event.EventListenerFunc(playerMarryDivorceDeal))
}
