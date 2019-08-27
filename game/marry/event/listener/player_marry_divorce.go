package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/friend/friend"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

//玩家离婚
func playerMarryDivorce(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*marryeventtypes.MarryDivorceEventData)
	if !ok {
		return
	}
	spouseId := eventData.GetSpouseId()
	divorceType := eventData.GetDivorceType()
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	spouseName := marryInfo.SpouseName
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)

	//强制离婚
	if divorceType == marrytypes.MarryDivorceTypeForce {
		//清空亲密度
		friend.GetFriendService().DivorceSubPoint(pl, spouseId, marrytypes.MarryDivorceTypeForce, 0)

		manager.Divorce()
		scMarryInfoStatusChange := pbuitl.BuildSCMarryInfoStatusChange(int32(marrytypes.MarryStatusTypeDivorce))
		pl.SendMsg(scMarryInfoStatusChange)

		if spl != nil {
			splCtx := scene.WithPlayer(context.Background(), spl)
			playerDivorceForceMsg := message.NewScheduleMessage(onDivorceForce, splCtx, pl.GetId(), nil)
			spl.Post(playerDivorceForceMsg)
		}

		//强制离婚跑马灯
		name := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
		peerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(spouseName))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryDivorceForce), name, peerName, peerName)
		noticelogic.NoticeNumBroadcast([]byte(content), marrytypes.WeddingInterval, 3)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
		returnPreMarry := eventData.GetPreMarryList()
		if len(returnPreMarry) > 0 {
			returnMarryPreList(returnPreMarry, pl, spl, spouseId)
		}
	} else {
		if spl == nil {
			return
		}
		splCtx := scene.WithPlayer(context.Background(), spl)
		playerDivorceConsentMsg := message.NewScheduleMessage(onDivorceConsent, splCtx, pl.GetName(), nil)
		spl.Post(playerDivorceConsentMsg)
	}
	return
}

func onDivorceForce(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.Divorce()

	scMarryInfoStatusChange := pbuitl.BuildSCMarryInfoStatusChange(int32(marrytypes.MarryStatusTypeDivorce))
	pl.SendMsg(scMarryInfoStatusChange)
	return nil
}

func onDivorceConsent(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	name := result.(string)

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.ReceiveDivorce()
	scMarryPushDivorce := pbuitl.BuildSCMarryPushDivorce(name)
	pl.SendMsg(scMarryPushDivorce)
	return nil
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryDivorce, event.EventListenerFunc(playerMarryDivorce))
}
