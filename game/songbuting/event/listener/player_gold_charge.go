package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/songbuting/pbutil"
	playersongbuting "fgame/fgame/game/songbuting/player"
	songbutingtemplate "fgame/fgame/game/songbuting/template"
	"fmt"
)

//玩家充值元宝
func playerChargeGold(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	chargeGold := data.(int32)

	manager := pl.GetPlayerDataManager(playertypes.PlayerSongBuTingDataManagerType).(*playersongbuting.PlayerSongBuTingDataManager)
	songBuTingObj := manager.GetSongBuTingObj()
	if songBuTingObj.GetIsReceive() {
		return
	}

	songBuTingTempalte := songbutingtemplate.GetSongBuTingTemplateService().GetSongBuTingTemplate()
	if chargeGold < songBuTingTempalte.NeedGold {
		return
	}
	manager.SetIsReceive()
	scSongBuTingChanged := pbutil.BuildSCSongBuTingChanged(songBuTingObj)
	pl.SendMsg(scSongBuTingChanged)

	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	operationName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(songBuTingTempalte.Name))
	chargeGoldName := coreutils.FormatColor(chattypes.ColorTypePower, coreutils.FormatNoticeStr(fmt.Sprintf("%d", chargeGold)))

	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.SongBuTingNotice), playerName, operationName, chargeGoldName, songBuTingTempalte.RewBindGold)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeGold))
}
